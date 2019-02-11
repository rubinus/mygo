// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Im.helloworld.proto

package example

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
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

type FOO int32

const (
	FOO_X FOO = 0
)

var FOO_name = map[int32]string{
	0: "X",
}

var FOO_value = map[string]int32{
	"X": 0,
}

func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}

func (FOO) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_34c7706cfdc151dc, []int{0}
}

type Helloworld struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Str                  string   `protobuf:"bytes,2,opt,name=str,proto3" json:"str,omitempty"`
	Opt                  int32    `protobuf:"varint,3,opt,name=opt,proto3" json:"opt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Helloworld) Reset()         { *m = Helloworld{} }
func (m *Helloworld) String() string { return proto.CompactTextString(m) }
func (*Helloworld) ProtoMessage()    {}
func (*Helloworld) Descriptor() ([]byte, []int) {
	return fileDescriptor_34c7706cfdc151dc, []int{0}
}
func (m *Helloworld) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Helloworld) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Helloworld.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Helloworld) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Helloworld.Merge(m, src)
}
func (m *Helloworld) XXX_Size() int {
	return m.Size()
}
func (m *Helloworld) XXX_DiscardUnknown() {
	xxx_messageInfo_Helloworld.DiscardUnknown(m)
}

var xxx_messageInfo_Helloworld proto.InternalMessageInfo

func (m *Helloworld) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Helloworld) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *Helloworld) GetOpt() int32 {
	if m != nil {
		return m.Opt
	}
	return 0
}

func init() {
	proto.RegisterEnum("example.FOO", FOO_name, FOO_value)
	proto.RegisterType((*Helloworld)(nil), "example.helloworld")
}

func init() { proto.RegisterFile("Im.helloworld.proto", fileDescriptor_34c7706cfdc151dc) }

var fileDescriptor_34c7706cfdc151dc = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xf6, 0xcc, 0xd5, 0xcb,
	0x48, 0xcd, 0xc9, 0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x55, 0x72, 0xe0, 0xe2, 0x42, 0x48, 0x0a, 0xf1,
	0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x09,
	0x70, 0x31, 0x17, 0x97, 0x14, 0x49, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x20, 0x91,
	0xfc, 0x82, 0x12, 0x09, 0x66, 0xb0, 0x12, 0x10, 0x53, 0x8b, 0x87, 0x8b, 0xd9, 0xcd, 0xdf, 0x5f,
	0x88, 0x95, 0x8b, 0x31, 0x42, 0x80, 0xc1, 0x49, 0xe0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4,
	0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0x6c, 0xa3, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x28, 0x17, 0xf1, 0xb6, 0x88, 0x00, 0x00, 0x00,
}

func (m *Helloworld) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Helloworld) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintImHelloworld(dAtA, i, uint64(m.Id))
	}
	if len(m.Str) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintImHelloworld(dAtA, i, uint64(len(m.Str)))
		i += copy(dAtA[i:], m.Str)
	}
	if m.Opt != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintImHelloworld(dAtA, i, uint64(m.Opt))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintImHelloworld(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Helloworld) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovImHelloworld(uint64(m.Id))
	}
	l = len(m.Str)
	if l > 0 {
		n += 1 + l + sovImHelloworld(uint64(l))
	}
	if m.Opt != 0 {
		n += 1 + sovImHelloworld(uint64(m.Opt))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovImHelloworld(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozImHelloworld(x uint64) (n int) {
	return sovImHelloworld(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Helloworld) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowImHelloworld
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
			return fmt.Errorf("proto: helloworld: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: helloworld: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImHelloworld
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Str", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImHelloworld
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
				return ErrInvalidLengthImHelloworld
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Str = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Opt", wireType)
			}
			m.Opt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowImHelloworld
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Opt |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipImHelloworld(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthImHelloworld
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
func skipImHelloworld(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowImHelloworld
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
					return 0, ErrIntOverflowImHelloworld
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
					return 0, ErrIntOverflowImHelloworld
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
				return 0, ErrInvalidLengthImHelloworld
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowImHelloworld
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
				next, err := skipImHelloworld(dAtA[start:])
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
	ErrInvalidLengthImHelloworld = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowImHelloworld   = fmt.Errorf("proto: integer overflow")
)
