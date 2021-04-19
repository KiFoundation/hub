// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sentinel/plan/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/sentinel-official/hub/x/node/types"
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

type GenesisPlan struct {
	Plan  Plan     `protobuf:"bytes,1,opt,name=plan,proto3" json:"_,omitempty"`
	Nodes []string `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (m *GenesisPlan) Reset()         { *m = GenesisPlan{} }
func (m *GenesisPlan) String() string { return proto.CompactTextString(m) }
func (*GenesisPlan) ProtoMessage()    {}
func (*GenesisPlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_12352cbc0b89b5a8, []int{0}
}
func (m *GenesisPlan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisPlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisPlan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisPlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisPlan.Merge(m, src)
}
func (m *GenesisPlan) XXX_Size() int {
	return m.Size()
}
func (m *GenesisPlan) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisPlan.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisPlan proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisPlan)(nil), "sentinel.plan.v1.GenesisPlan")
}

func init() { proto.RegisterFile("sentinel/plan/v1/genesis.proto", fileDescriptor_12352cbc0b89b5a8) }

var fileDescriptor_12352cbc0b89b5a8 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x4e, 0xcd, 0x2b,
	0xc9, 0xcc, 0x4b, 0xcd, 0xd1, 0x2f, 0xc8, 0x49, 0xcc, 0xd3, 0x2f, 0x33, 0xd4, 0x4f, 0x4f, 0xcd,
	0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0xc9, 0xeb, 0x81,
	0xe4, 0xf5, 0xca, 0x0c, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x92, 0xfa, 0x20, 0x16, 0x44,
	0x9d, 0x94, 0x34, 0xdc, 0x9c, 0xbc, 0xfc, 0x94, 0x54, 0x90, 0x39, 0x20, 0x1a, 0x43, 0x12, 0x66,
	0x09, 0xd8, 0x30, 0xb0, 0xa4, 0x52, 0x11, 0x17, 0xb7, 0x3b, 0xc4, 0xca, 0x80, 0x9c, 0xc4, 0x3c,
	0x21, 0x7b, 0x2e, 0x16, 0x90, 0xa4, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x98, 0x1e, 0xba,
	0xfd, 0x7a, 0x20, 0x55, 0x4e, 0xc2, 0x27, 0xee, 0xc9, 0x33, 0xbc, 0xba, 0x27, 0xcf, 0x1d, 0xaf,
	0x93, 0x9f, 0x9b, 0x59, 0x92, 0x9a, 0x5b, 0x50, 0x52, 0x19, 0x04, 0xd6, 0x28, 0x24, 0xc2, 0xc5,
	0x0a, 0xb2, 0xba, 0x58, 0x82, 0x49, 0x81, 0x59, 0x83, 0x33, 0x08, 0xc2, 0xb1, 0xe2, 0xe9, 0x58,
	0x20, 0xcf, 0x30, 0x63, 0x81, 0x3c, 0xe3, 0x8b, 0x05, 0xf2, 0x0c, 0x4e, 0xee, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e,
	0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x9b, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97,
	0x9c, 0x9f, 0xab, 0x0f, 0xb3, 0x5a, 0x37, 0x3f, 0x2d, 0x2d, 0x33, 0x39, 0x33, 0x31, 0x47, 0x3f,
	0xa3, 0x34, 0x49, 0xbf, 0x02, 0xe2, 0x89, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x1f,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x18, 0xe7, 0xc8, 0x10, 0x47, 0x01, 0x00, 0x00,
}

func (m *GenesisPlan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisPlan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisPlan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for iNdEx := len(m.Nodes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Nodes[iNdEx])
			copy(dAtA[i:], m.Nodes[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.Nodes[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Plan.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisPlan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Plan.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Nodes) > 0 {
		for _, s := range m.Nodes {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisPlan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisPlan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisPlan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Plan.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)