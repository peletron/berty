// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: entity/local_record.proto

package entity

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LocalRecord struct {
	ContactId            string   `protobuf:"bytes,1,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalRecord) Reset()         { *m = LocalRecord{} }
func (m *LocalRecord) String() string { return proto.CompactTextString(m) }
func (*LocalRecord) ProtoMessage()    {}
func (*LocalRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_b561157065c2b28c, []int{0}
}
func (m *LocalRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LocalRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LocalRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LocalRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalRecord.Merge(m, src)
}
func (m *LocalRecord) XXX_Size() int {
	return m.Size()
}
func (m *LocalRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalRecord.DiscardUnknown(m)
}

var xxx_messageInfo_LocalRecord proto.InternalMessageInfo

func (m *LocalRecord) GetContactId() string {
	if m != nil {
		return m.ContactId
	}
	return ""
}

func init() {
	proto.RegisterType((*LocalRecord)(nil), "berty.entity.LocalRecord")
}

func init() { proto.RegisterFile("entity/local_record.proto", fileDescriptor_b561157065c2b28c) }

var fileDescriptor_b561157065c2b28c = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0xcd, 0x2b, 0xc9,
	0x2c, 0xa9, 0xd4, 0xcf, 0xc9, 0x4f, 0x4e, 0xcc, 0x89, 0x2f, 0x4a, 0x4d, 0xce, 0x2f, 0x4a, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x49, 0x4a, 0x2d, 0x2a, 0xa9, 0xd4, 0x83, 0x28, 0x50,
	0xd2, 0xe1, 0xe2, 0xf6, 0x01, 0xa9, 0x09, 0x02, 0x2b, 0x11, 0x92, 0xe5, 0xe2, 0x4a, 0xce, 0xcf,
	0x2b, 0x49, 0x4c, 0x2e, 0x89, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x84,
	0x8a, 0x78, 0xa6, 0x38, 0x69, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47,
	0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x44, 0x89, 0x41, 0x4c, 0x2b, 0x49, 0x4d, 0xce, 0xd0, 0x4f,
	0xce, 0x2f, 0x4a, 0xd5, 0x87, 0x98, 0x9b, 0xc4, 0x06, 0xb6, 0xcc, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x10, 0xc8, 0x42, 0xb4, 0x89, 0x00, 0x00, 0x00,
}

func (m *LocalRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LocalRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ContactId) > 0 {
		i -= len(m.ContactId)
		copy(dAtA[i:], m.ContactId)
		i = encodeVarintLocalRecord(dAtA, i, uint64(len(m.ContactId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLocalRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovLocalRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LocalRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContactId)
	if l > 0 {
		n += 1 + l + sovLocalRecord(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLocalRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLocalRecord(x uint64) (n int) {
	return sovLocalRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LocalRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocalRecord
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LocalRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocalRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContactId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLocalRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLocalRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLocalRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContactId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLocalRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLocalRecord
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLocalRecord
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLocalRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLocalRecord
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
					return 0, ErrIntOverflowLocalRecord
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
					return 0, ErrIntOverflowLocalRecord
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
			if length < 0 {
				return 0, ErrInvalidLengthLocalRecord
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthLocalRecord
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLocalRecord
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
				next, err := skipLocalRecord(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthLocalRecord
				}
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
	ErrInvalidLengthLocalRecord = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLocalRecord   = fmt.Errorf("proto: integer overflow")
)
