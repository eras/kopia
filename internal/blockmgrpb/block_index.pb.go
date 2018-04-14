// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal/blockmgrpb/block_index.proto

/*
	Package blockmgrpb is a generated protocol buffer package.

	It is generated from these files:
		internal/blockmgrpb/block_index.proto

	It has these top-level messages:
		IndexV1
		Indexes
*/
package blockmgrpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IndexV1 struct {
	PackBlockId     string            `protobuf:"bytes,1,opt,name=pack_block_id,json=packBlockId,proto3" json:"pack_block_id,omitempty"`
	PackLength      uint64            `protobuf:"varint,2,opt,name=pack_length,json=packLength,proto3" json:"pack_length,omitempty"`
	CreateTimeNanos uint64            `protobuf:"varint,3,opt,name=create_time_nanos,json=createTimeNanos,proto3" json:"create_time_nanos,omitempty"`
	Items           map[string]uint64 `protobuf:"bytes,4,rep,name=items" json:"items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	DeletedItems    []string          `protobuf:"bytes,5,rep,name=deleted_items,json=deletedItems" json:"deleted_items,omitempty"`
	InlineItems     map[string][]byte `protobuf:"bytes,6,rep,name=inline_items,json=inlineItems" json:"inline_items,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *IndexV1) Reset()                    { *m = IndexV1{} }
func (m *IndexV1) String() string            { return proto.CompactTextString(m) }
func (*IndexV1) ProtoMessage()               {}
func (*IndexV1) Descriptor() ([]byte, []int) { return fileDescriptorBlockIndex, []int{0} }

func (m *IndexV1) GetPackBlockId() string {
	if m != nil {
		return m.PackBlockId
	}
	return ""
}

func (m *IndexV1) GetPackLength() uint64 {
	if m != nil {
		return m.PackLength
	}
	return 0
}

func (m *IndexV1) GetCreateTimeNanos() uint64 {
	if m != nil {
		return m.CreateTimeNanos
	}
	return 0
}

func (m *IndexV1) GetItems() map[string]uint64 {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *IndexV1) GetDeletedItems() []string {
	if m != nil {
		return m.DeletedItems
	}
	return nil
}

func (m *IndexV1) GetInlineItems() map[string][]byte {
	if m != nil {
		return m.InlineItems
	}
	return nil
}

type Indexes struct {
	IndexesV1 []*IndexV1 `protobuf:"bytes,1,rep,name=indexes_v1,json=indexesV1" json:"indexes_v1,omitempty"`
}

func (m *Indexes) Reset()                    { *m = Indexes{} }
func (m *Indexes) String() string            { return proto.CompactTextString(m) }
func (*Indexes) ProtoMessage()               {}
func (*Indexes) Descriptor() ([]byte, []int) { return fileDescriptorBlockIndex, []int{1} }

func (m *Indexes) GetIndexesV1() []*IndexV1 {
	if m != nil {
		return m.IndexesV1
	}
	return nil
}

func init() {
	proto.RegisterType((*IndexV1)(nil), "IndexV1")
	proto.RegisterType((*Indexes)(nil), "Indexes")
}
func (m *IndexV1) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexV1) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PackBlockId) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBlockIndex(dAtA, i, uint64(len(m.PackBlockId)))
		i += copy(dAtA[i:], m.PackBlockId)
	}
	if m.PackLength != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBlockIndex(dAtA, i, uint64(m.PackLength))
	}
	if m.CreateTimeNanos != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintBlockIndex(dAtA, i, uint64(m.CreateTimeNanos))
	}
	if len(m.Items) > 0 {
		for k, _ := range m.Items {
			dAtA[i] = 0x22
			i++
			v := m.Items[k]
			mapSize := 1 + len(k) + sovBlockIndex(uint64(len(k))) + 1 + sovBlockIndex(uint64(v))
			i = encodeVarintBlockIndex(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintBlockIndex(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintBlockIndex(dAtA, i, uint64(v))
		}
	}
	if len(m.DeletedItems) > 0 {
		for _, s := range m.DeletedItems {
			dAtA[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.InlineItems) > 0 {
		for k, _ := range m.InlineItems {
			dAtA[i] = 0x32
			i++
			v := m.InlineItems[k]
			byteSize := 0
			if len(v) > 0 {
				byteSize = 1 + len(v) + sovBlockIndex(uint64(len(v)))
			}
			mapSize := 1 + len(k) + sovBlockIndex(uint64(len(k))) + byteSize
			i = encodeVarintBlockIndex(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintBlockIndex(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if len(v) > 0 {
				dAtA[i] = 0x12
				i++
				i = encodeVarintBlockIndex(dAtA, i, uint64(len(v)))
				i += copy(dAtA[i:], v)
			}
		}
	}
	return i, nil
}

func (m *Indexes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Indexes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.IndexesV1) > 0 {
		for _, msg := range m.IndexesV1 {
			dAtA[i] = 0xa
			i++
			i = encodeVarintBlockIndex(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintBlockIndex(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *IndexV1) Size() (n int) {
	var l int
	_ = l
	l = len(m.PackBlockId)
	if l > 0 {
		n += 1 + l + sovBlockIndex(uint64(l))
	}
	if m.PackLength != 0 {
		n += 1 + sovBlockIndex(uint64(m.PackLength))
	}
	if m.CreateTimeNanos != 0 {
		n += 1 + sovBlockIndex(uint64(m.CreateTimeNanos))
	}
	if len(m.Items) > 0 {
		for k, v := range m.Items {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovBlockIndex(uint64(len(k))) + 1 + sovBlockIndex(uint64(v))
			n += mapEntrySize + 1 + sovBlockIndex(uint64(mapEntrySize))
		}
	}
	if len(m.DeletedItems) > 0 {
		for _, s := range m.DeletedItems {
			l = len(s)
			n += 1 + l + sovBlockIndex(uint64(l))
		}
	}
	if len(m.InlineItems) > 0 {
		for k, v := range m.InlineItems {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovBlockIndex(uint64(len(v)))
			}
			mapEntrySize := 1 + len(k) + sovBlockIndex(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovBlockIndex(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *Indexes) Size() (n int) {
	var l int
	_ = l
	if len(m.IndexesV1) > 0 {
		for _, e := range m.IndexesV1 {
			l = e.Size()
			n += 1 + l + sovBlockIndex(uint64(l))
		}
	}
	return n
}

func sovBlockIndex(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBlockIndex(x uint64) (n int) {
	return sovBlockIndex(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IndexV1) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockIndex
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IndexV1: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexV1: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PackBlockId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlockIndex
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PackBlockId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PackLength", wireType)
			}
			m.PackLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PackLength |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateTimeNanos", wireType)
			}
			m.CreateTimeNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateTimeNanos |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlockIndex
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Items == nil {
				m.Items = make(map[string]uint64)
			}
			var mapkey string
			var mapvalue uint64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBlockIndex
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBlockIndex
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthBlockIndex
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBlockIndex
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipBlockIndex(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthBlockIndex
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Items[mapkey] = mapvalue
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedItems", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBlockIndex
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeletedItems = append(m.DeletedItems, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InlineItems", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlockIndex
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.InlineItems == nil {
				m.InlineItems = make(map[string][]byte)
			}
			var mapkey string
			mapvalue := []byte{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowBlockIndex
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBlockIndex
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthBlockIndex
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapbyteLen uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowBlockIndex
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapbyteLen |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intMapbyteLen := int(mapbyteLen)
					if intMapbyteLen < 0 {
						return ErrInvalidLengthBlockIndex
					}
					postbytesIndex := iNdEx + intMapbyteLen
					if postbytesIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = make([]byte, mapbyteLen)
					copy(mapvalue, dAtA[iNdEx:postbytesIndex])
					iNdEx = postbytesIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipBlockIndex(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthBlockIndex
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.InlineItems[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockIndex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlockIndex
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Indexes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockIndex
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Indexes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Indexes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexesV1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockIndex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlockIndex
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IndexesV1 = append(m.IndexesV1, &IndexV1{})
			if err := m.IndexesV1[len(m.IndexesV1)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockIndex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlockIndex
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBlockIndex(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlockIndex
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBlockIndex
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBlockIndex
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthBlockIndex
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBlockIndex
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipBlockIndex(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthBlockIndex = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlockIndex   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("internal/blockmgrpb/block_index.proto", fileDescriptorBlockIndex) }

var fileDescriptorBlockIndex = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x5d, 0x0a, 0x28, 0x03, 0x44, 0x5c, 0x3d, 0x54, 0x62, 0x6a, 0x83, 0x31, 0x56, 0x0f,
	0x18, 0xf0, 0x42, 0x8c, 0xf1, 0x40, 0xe2, 0x81, 0xc4, 0x78, 0x68, 0x0c, 0x07, 0x2f, 0x9b, 0x42,
	0x27, 0xb8, 0xa1, 0xdd, 0x92, 0xb2, 0x12, 0xb9, 0xfa, 0x14, 0x3e, 0x92, 0x47, 0x1f, 0xc1, 0xd4,
	0x17, 0x31, 0xbb, 0x5d, 0x53, 0xe2, 0xc5, 0xdb, 0xcc, 0x37, 0xff, 0xff, 0x67, 0x67, 0x07, 0x4e,
	0xb9, 0x90, 0x98, 0x8a, 0x20, 0xba, 0x9c, 0x44, 0xc9, 0x74, 0x1e, 0xcf, 0xd2, 0xc5, 0x24, 0x2f,
	0x19, 0x17, 0x21, 0xbe, 0x76, 0x17, 0x69, 0x22, 0x93, 0xce, 0x9b, 0x05, 0xdb, 0x23, 0xd5, 0x8f,
	0x7b, 0xb4, 0x03, 0xcd, 0x45, 0x30, 0x9d, 0x33, 0xa3, 0x0a, 0x6d, 0xe2, 0x12, 0xaf, 0xe6, 0xd7,
	0x15, 0x1c, 0x2a, 0x36, 0x0a, 0xe9, 0x31, 0xe8, 0x96, 0x45, 0x28, 0x66, 0xf2, 0xd9, 0x2e, 0xb9,
	0xc4, 0x2b, 0xfb, 0xa0, 0xd0, 0xbd, 0x26, 0xf4, 0x02, 0xf6, 0xa6, 0x29, 0x06, 0x12, 0x99, 0xe4,
	0x31, 0x32, 0x11, 0x88, 0x64, 0x69, 0x5b, 0x5a, 0xb6, 0x9b, 0x0f, 0x1e, 0x79, 0x8c, 0x0f, 0x0a,
	0xd3, 0x73, 0xa8, 0x70, 0x89, 0xf1, 0xd2, 0x2e, 0xbb, 0x96, 0x57, 0xef, 0xef, 0x77, 0xcd, 0x4b,
	0xba, 0x23, 0x45, 0xef, 0x84, 0x4c, 0xd7, 0x7e, 0xae, 0xa0, 0x27, 0xd0, 0x0c, 0x31, 0x42, 0x89,
	0x21, 0xcb, 0x2d, 0x15, 0xd7, 0xf2, 0x6a, 0x7e, 0xc3, 0x40, 0x6d, 0xa0, 0x37, 0xd0, 0xe0, 0x22,
	0xe2, 0x02, 0x8d, 0xa6, 0xaa, 0x63, 0x0f, 0x8b, 0x58, 0x3d, 0xdc, 0x08, 0xaf, 0xf3, 0x82, 0xb4,
	0x07, 0x00, 0xc5, 0x88, 0xb6, 0xc0, 0x9a, 0xe3, 0xda, 0x7c, 0x81, 0x2a, 0xe9, 0x01, 0x54, 0x56,
	0x41, 0xf4, 0x82, 0x66, 0xe9, 0xbc, 0xb9, 0x2e, 0x0d, 0x48, 0xfb, 0x16, 0x5a, 0x7f, 0xa3, 0xff,
	0xf3, 0x37, 0x36, 0xfc, 0x9d, 0xbe, 0xb9, 0x01, 0x2e, 0xe9, 0x19, 0x00, 0xcf, 0x4b, 0xb6, 0xea,
	0xd9, 0x44, 0x2f, 0xb0, 0xf3, 0xbb, 0x80, 0x5f, 0x33, 0xb3, 0x71, 0x6f, 0x78, 0xf4, 0x91, 0x39,
	0xe4, 0x33, 0x73, 0xc8, 0x57, 0xe6, 0x90, 0xf7, 0x6f, 0x67, 0xeb, 0x09, 0x8a, 0x43, 0x4f, 0xaa,
	0xfa, 0xba, 0x57, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xec, 0xc2, 0xbe, 0x06, 0x02, 0x00,
	0x00,
}