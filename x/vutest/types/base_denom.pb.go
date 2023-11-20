// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vutest/vutest/base_denom.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type BaseDenom struct {
	Denom   types.Coin `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom"`
	Creator string     `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *BaseDenom) Reset()         { *m = BaseDenom{} }
func (m *BaseDenom) String() string { return proto.CompactTextString(m) }
func (*BaseDenom) ProtoMessage()    {}
func (*BaseDenom) Descriptor() ([]byte, []int) {
	return fileDescriptor_912d728ab04c6427, []int{0}
}
func (m *BaseDenom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseDenom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseDenom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseDenom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseDenom.Merge(m, src)
}
func (m *BaseDenom) XXX_Size() int {
	return m.Size()
}
func (m *BaseDenom) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseDenom.DiscardUnknown(m)
}

var xxx_messageInfo_BaseDenom proto.InternalMessageInfo

func (m *BaseDenom) GetDenom() types.Coin {
	if m != nil {
		return m.Denom
	}
	return types.Coin{}
}

func (m *BaseDenom) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*BaseDenom)(nil), "vutest.vutest.BaseDenom")
}

func init() { proto.RegisterFile("vutest/vutest/base_denom.proto", fileDescriptor_912d728ab04c6427) }

var fileDescriptor_912d728ab04c6427 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x2b, 0x2d, 0x49,
	0x2d, 0x2e, 0xd1, 0x87, 0x52, 0x49, 0x89, 0xc5, 0xa9, 0xf1, 0x29, 0xa9, 0x79, 0xf9, 0xb9, 0x7a,
	0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xbc, 0x10, 0x09, 0x3d, 0x08, 0x25, 0x25, 0x97, 0x9c, 0x5f,
	0x9c, 0x9b, 0x5f, 0x0c, 0x56, 0xa7, 0x5f, 0x66, 0x98, 0x94, 0x5a, 0x92, 0x68, 0xa8, 0x9f, 0x9c,
	0x9f, 0x99, 0x07, 0x51, 0x2e, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x66, 0xea, 0x83, 0x58, 0x10,
	0x51, 0xa5, 0x18, 0x2e, 0x4e, 0xa7, 0xc4, 0xe2, 0x54, 0x17, 0x90, 0xb9, 0x42, 0xa6, 0x5c, 0xac,
	0x60, 0x0b, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x24, 0xf5, 0x20, 0x46, 0xea, 0x81, 0x8c,
	0xd4, 0x83, 0x1a, 0xa9, 0xe7, 0x9c, 0x9f, 0x99, 0xe7, 0xc4, 0x72, 0xe2, 0x9e, 0x3c, 0x43, 0x10,
	0x44, 0xb5, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x51, 0x6a, 0x62, 0x49, 0x7e, 0x91, 0x04, 0x93, 0x02,
	0xa3, 0x06, 0x67, 0x10, 0x8c, 0xeb, 0xa4, 0x7f, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72,
	0x0c, 0x51, 0xa2, 0x50, 0x5f, 0x55, 0xc0, 0xbc, 0x57, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06,
	0x76, 0x95, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x57, 0xda, 0xca, 0xa5, 0xfc, 0x00, 0x00, 0x00,
}

func (m *BaseDenom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseDenom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseDenom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintBaseDenom(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Denom.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintBaseDenom(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintBaseDenom(dAtA []byte, offset int, v uint64) int {
	offset -= sovBaseDenom(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseDenom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Denom.Size()
	n += 1 + l + sovBaseDenom(uint64(l))
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovBaseDenom(uint64(l))
	}
	return n
}

func sovBaseDenom(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBaseDenom(x uint64) (n int) {
	return sovBaseDenom(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseDenom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBaseDenom
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
			return fmt.Errorf("proto: BaseDenom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseDenom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseDenom
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBaseDenom
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBaseDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Denom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseDenom
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
				return ErrInvalidLengthBaseDenom
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBaseDenom
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBaseDenom(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBaseDenom
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
func skipBaseDenom(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBaseDenom
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
					return 0, ErrIntOverflowBaseDenom
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBaseDenom
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
				return 0, ErrInvalidLengthBaseDenom
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBaseDenom
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBaseDenom
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBaseDenom        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBaseDenom          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBaseDenom = fmt.Errorf("proto: unexpected end of group")
)
