package fs

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/kopia/kopia/cas"
)

// import (
// 	"os"
// 	"strings"
// 	"time"

// 	"github.com/kopia/kopia/cas"
// )

// Entry stores attributes of a single entry in a directory.
type Entry struct {
	Name     string
	FileMode os.FileMode
	FileSize int64
	ModTime  time.Time
	UserID   uint32
	GroupID  uint32
	ObjectID cas.ObjectID
}

func isLess(name1, name2 string) bool {
	if name1 == name2 {
		return false
	}

	return isLessOrEqual(name1, name2)
}

func split1(name string) (head, tail string) {
	n := strings.IndexByte(name, '/')
	if n >= 0 {
		return name[0 : n+1], name[n+1:]
	}

	return name, ""
}

func isLessOrEqual(name1, name2 string) bool {
	parts1 := strings.Split(name1, "/")
	parts2 := strings.Split(name2, "/")

	i := 0
	for i < len(parts1) && i < len(parts2) {
		if parts1[i] == parts2[i] {
			i++
			continue
		}
		if parts1[i] == "" {
			return false
		}
		if parts2[i] == "" {
			return true
		}
		return parts1[i] < parts2[i]
	}

	return len(parts1) <= len(parts2)
}

func (e *Entry) IsDir() bool {
	return e.FileMode.IsDir()
}

func metadataEquals(e1, e2 *Entry) bool {
	if (e1 != nil) != (e2 != nil) {
		return false
	}

	if e1.FileMode != e2.FileMode {
		log.Printf("a1")
		return false
	}

	if e1.ModTime.UnixNano() != e2.ModTime.UnixNano() {
		log.Printf("a2")
		return false
	}

	if e1.FileSize != e2.FileSize {
		log.Printf("a3")
		return false
	}

	if e1.UserID != e2.UserID {
		log.Printf("a4")
		return false
	}

	if e1.GroupID != e2.GroupID {
		log.Printf("a5")
		return false
	}

	return true
}