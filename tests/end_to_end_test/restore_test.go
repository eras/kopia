package endtoend_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/kopia/kopia/fs/localfs"
	"github.com/kopia/kopia/internal/diff"
	"github.com/kopia/kopia/internal/fshasher"
	"github.com/kopia/kopia/tests/testenv"
)

func TestRestoreCommand(t *testing.T) {
	t.Parallel()

	e := testenv.NewCLITest(t)
	defer e.Cleanup(t)

	e.RunAndExpectSuccess(t, "repo", "create", "filesystem", "--path", e.RepoDir)

	source := makeScratchDir(t)
	testenv.MustCreateDirectoryTree(t, source, testenv.DirectoryTreeOptions{
		Depth:                1,
		MaxFilesPerDirectory: 10,
	})

	restoreDir := makeScratchDir(t)
	r1 := makeScratchDir(t)
	// Attempt to restore a snapshot from an empty repo.
	e.RunAndExpectFailure(t, "restore", "kffbb7c28ea6c34d6cbe555d1cf80faa9", r1)
	e.RunAndExpectSuccess(t, "snapshot", "create", source)

	// obtain snapshot root id and use it for restore
	si := e.ListSnapshotsAndExpectSuccess(t, source)
	if got, want := len(si), 1; got != want {
		t.Fatalf("got %v sources, wanted %v", got, want)
	}

	if got, want := len(si[0].Snapshots), 1; got != want {
		t.Fatalf("got %v snapshots, wanted %v", got, want)
	}

	snapID := si[0].Snapshots[0].SnapshotID
	rootID := si[0].Snapshots[0].ObjectID

	r2 := makeScratchDir(t)
	// Attempt to restore a non-existing snapshot.
	e.RunAndExpectFailure(t, "restore", "kffbb7c28ea6c34d6cbe555d1cf80fdd9", r2)

	// Ensure restored files are created with a different ModTime
	time.Sleep(time.Second)

	// Attempt to restore using snapshot ID
	restoreFailDir := makeScratchDir(t)
	e.RunAndExpectFailure(t, "restore", snapID, restoreFailDir)

	// Restore last snapshot
	e.RunAndExpectSuccess(t, "restore", rootID, restoreDir)

	// Note: restore does not reset the permissions for the top directory due to
	// the way the top FS entry is created in snapshotfs. Force the permissions
	// of the top directory to match those of the source so the recursive
	// directory comparison has a chance of succeeding.
	testenv.AssertNoError(t, os.Chmod(restoreDir, 0700))
	compareDirs(t, source, restoreDir)

	// Attempt to restore into a target directory that already exists
	e.RunAndExpectFailure(t, "restore", rootID, restoreDir, "--no-overwrite-directories")

	// Attempt to restore into a target directory that already exists
	e.RunAndExpectFailure(t, "restore", rootID, restoreDir, "--no-overwrite-files")
}

func compareDirs(t *testing.T, source, restoreDir string) {
	t.Helper()

	// Restored contents should match source
	s, err := localfs.Directory(source)
	testenv.AssertNoError(t, err)
	wantHash, err := fshasher.Hash(context.Background(), s)
	testenv.AssertNoError(t, err)

	// check restored contents
	r, err := localfs.Directory(restoreDir)
	testenv.AssertNoError(t, err)

	ctx := context.Background()
	gotHash, err := fshasher.Hash(ctx, r)
	testenv.AssertNoError(t, err)

	if !assert.Equal(t, wantHash, gotHash, "restored directory hash does not match source's hash") {
		cmp, err := diff.NewComparer(os.Stderr)
		testenv.AssertNoError(t, err)

		cmp.DiffCommand = "cmp"
		_ = cmp.Compare(ctx, s, r)
	}
}

func TestSnapshotRestore(t *testing.T) {
	e := testenv.NewCLITest(t)
	defer e.Cleanup(t)
	defer e.RunAndExpectSuccess(t, "repo", "disconnect")

	e.RunAndExpectSuccess(t, "repo", "create", "filesystem", "--path", e.RepoDir)

	source := makeScratchDir(t)
	testenv.MustCreateDirectoryTree(t, source, testenv.DirectoryTreeOptions{
		Depth:                  5,
		MaxSubdirsPerDirectory: 5,
		MaxFilesPerDirectory:   5,
	})

	restoreDir := makeScratchDir(t)
	r1 := makeScratchDir(t)
	// Attempt to restore a snapshot from an empty repo.
	e.RunAndExpectFailure(t, "snapshot", "restore", "kffbb7c28ea6c34d6cbe555d1cf80faa9", r1)
	e.RunAndExpectSuccess(t, "snapshot", "create", source)

	// obtain snapshot root id and use it for restore
	si := e.ListSnapshotsAndExpectSuccess(t, source)
	if got, want := len(si), 1; got != want {
		t.Fatalf("got %v sources, wanted %v", got, want)
	}

	if got, want := len(si[0].Snapshots), 1; got != want {
		t.Fatalf("got %v snapshots, wanted %v", got, want)
	}

	snapID := si[0].Snapshots[0].SnapshotID
	rootID := si[0].Snapshots[0].ObjectID

	// Attempt to restore a non-existing snapshot.
	r2 := makeScratchDir(t)
	e.RunAndExpectFailure(t, "snapshot", "restore", "kffbb7c28ea6c34d6cbe555d1cf80fdd9", r2)

	// Ensure restored files are created with a different ModTime
	time.Sleep(time.Second)

	// Attempt to restore snapshot with root ID
	restoreFailDir := makeScratchDir(t)
	e.RunAndExpectFailure(t, "snapshot", "restore", rootID, restoreFailDir)

	// Restore last snapshot using the snapshot ID
	e.RunAndExpectSuccess(t, "snapshot", "restore", snapID, restoreDir)

	// Restored contents should match source
	compareDirs(t, source, restoreDir)

	// Attempt to restore snapshot with an already-existing target directory
	// It should fail because the directory is not empty
	_ = os.MkdirAll(restoreFailDir, 0700)

	e.RunAndExpectFailure(t, "snapshot", "restore", "--no-overwrite-directories", snapID, restoreDir)

	// Attempt to restore snapshot with an already-existing target directory
	// It should fail because target files already exist
	_ = os.MkdirAll(restoreFailDir, 0700)

	e.RunAndExpectFailure(t, "snapshot", "restore", "--no-overwrite-files", snapID, restoreDir)
}
