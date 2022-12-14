// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: consensus/config.proto

package consensus

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ConfOp int32

const (
	AddVoter        ConfOp = 0
	AddLearner      ConfOp = 1
	Remove          ConfOp = 2
	UpdateToLearner ConfOp = 3
	UpdateToVoter   ConfOp = 4
)

var ConfOp_name = map[int32]string{
	0: "AddVoter",
	1: "AddLearner",
	2: "Remove",
	3: "UpdateToLearner",
	4: "UpdateToVoter",
}

var ConfOp_value = map[string]int32{
	"AddVoter":        0,
	"AddLearner":      1,
	"Remove":          2,
	"UpdateToLearner": 3,
	"UpdateToVoter":   4,
}

func (x ConfOp) String() string {
	return proto.EnumName(ConfOp_name, int32(x))
}

func (ConfOp) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e2dfec80d10c6abe, []int{0}
}

type Consenter struct {
	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Org      string `protobuf:"bytes,2,opt,name=org,proto3" json:"org,omitempty"`
	Host     string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Port     uint64 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	ReqPort  uint32 `protobuf:"varint,7,opt,name=req_port,json=reqPort,proto3" json:"req_port,omitempty"`
	RestPort uint32 `protobuf:"varint,8,opt,name=rest_port,json=restPort,proto3" json:"rest_port,omitempty"`
	ReeCert  []byte `protobuf:"bytes,5,opt,name=ree_cert,json=reeCert,proto3" json:"ree_cert,omitempty"`
	TeeCert  []byte `protobuf:"bytes,6,opt,name=tee_cert,json=teeCert,proto3" json:"tee_cert,omitempty"`
}

func (m *Consenter) Reset()         { *m = Consenter{} }
func (m *Consenter) String() string { return proto.CompactTextString(m) }
func (*Consenter) ProtoMessage()    {}
func (*Consenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2dfec80d10c6abe, []int{0}
}
func (m *Consenter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Consenter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Consenter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Consenter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consenter.Merge(m, src)
}
func (m *Consenter) XXX_Size() int {
	return m.Size()
}
func (m *Consenter) XXX_DiscardUnknown() {
	xxx_messageInfo_Consenter.DiscardUnknown(m)
}

var xxx_messageInfo_Consenter proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("consensus.ConfOp", ConfOp_name, ConfOp_value)
	proto.RegisterType((*Consenter)(nil), "consensus.Consenter")
}

func init() { proto.RegisterFile("consensus/config.proto", fileDescriptor_e2dfec80d10c6abe) }

var fileDescriptor_e2dfec80d10c6abe = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xb1, 0xce, 0xd3, 0x30,
	0x14, 0x85, 0xe3, 0xbf, 0x21, 0x4d, 0xac, 0x16, 0x82, 0x41, 0x28, 0x80, 0x64, 0x45, 0x0c, 0x28,
	0x62, 0x48, 0x06, 0x16, 0xd6, 0xd2, 0x15, 0x89, 0x2a, 0x02, 0x86, 0x2e, 0x55, 0x48, 0x6e, 0x93,
	0x0c, 0xf1, 0x4d, 0x1d, 0x97, 0xee, 0x3c, 0x01, 0x6f, 0x45, 0xc7, 0x8e, 0x8c, 0xb4, 0x7d, 0x11,
	0x64, 0x5b, 0x89, 0xd8, 0x8e, 0xbe, 0xef, 0x1c, 0x0f, 0xbe, 0xf4, 0x45, 0x89, 0x62, 0x00, 0x31,
	0x1c, 0x87, 0xac, 0x44, 0xb1, 0x6f, 0xeb, 0xb4, 0x97, 0xa8, 0x90, 0x05, 0x13, 0x7f, 0xf5, 0xbc,
	0xc6, 0x1a, 0x0d, 0xcd, 0x74, 0xb2, 0x85, 0x37, 0xbf, 0x09, 0x0d, 0xd6, 0xa6, 0xa3, 0x40, 0x32,
	0x46, 0x5d, 0x51, 0x74, 0x10, 0x91, 0x98, 0x24, 0x41, 0x6e, 0x32, 0x0b, 0xe9, 0x0c, 0x65, 0x1d,
	0x3d, 0x18, 0xa4, 0xa3, 0x6e, 0x35, 0x38, 0xa8, 0x68, 0x66, 0x5b, 0x3a, 0x6b, 0xd6, 0xa3, 0x54,
	0x91, 0x1b, 0x93, 0xc4, 0xcd, 0x4d, 0x66, 0x2f, 0xa9, 0x2f, 0xe1, 0xb0, 0x33, 0x7c, 0x1e, 0x93,
	0x64, 0x99, 0xcf, 0x25, 0x1c, 0x36, 0x5a, 0xbd, 0xa6, 0x81, 0x84, 0x41, 0x59, 0xe7, 0x1b, 0xe7,
	0x6b, 0xb0, 0x99, 0x76, 0xb0, 0x2b, 0x41, 0xaa, 0xe8, 0x51, 0x4c, 0x92, 0x85, 0xde, 0xc1, 0x1a,
	0xac, 0x52, 0xa3, 0xf2, 0xac, 0x52, 0x56, 0xbd, 0xdb, 0x52, 0x6f, 0x8d, 0x62, 0xff, 0xb9, 0x67,
	0x0b, 0xea, 0xaf, 0xaa, 0xea, 0x1b, 0x2a, 0x90, 0xa1, 0xc3, 0x1e, 0x53, 0xba, 0xaa, 0xaa, 0x4f,
	0x50, 0x48, 0x01, 0x32, 0x24, 0x8c, 0x52, 0x2f, 0x87, 0x0e, 0x7f, 0x40, 0xf8, 0xc0, 0x9e, 0xd1,
	0x27, 0x5f, 0xfb, 0xaa, 0x50, 0xf0, 0x05, 0xc7, 0xc2, 0x8c, 0x3d, 0xa5, 0xcb, 0x11, 0xda, 0x37,
	0xdc, 0x8f, 0x3f, 0xc9, 0xf9, 0xca, 0x9d, 0xcb, 0x95, 0x3b, 0xe7, 0x1b, 0x27, 0x97, 0x1b, 0x27,
	0x7f, 0x6f, 0x9c, 0xfc, 0xba, 0x73, 0xe7, 0x72, 0xe7, 0xce, 0x9f, 0x3b, 0x77, 0xe8, 0xdb, 0x12,
	0xbb, 0xb4, 0x39, 0x16, 0x27, 0x68, 0xd3, 0x53, 0x0b, 0x02, 0x64, 0xd9, 0x14, 0xad, 0xb0, 0xdf,
	0x9c, 0x4e, 0x67, 0xd8, 0x7e, 0xa8, 0x5b, 0x35, 0xf6, 0x4a, 0xec, 0xb2, 0x1e, 0xdb, 0x61, 0xd0,
	0xb6, 0x90, 0x65, 0x93, 0xfd, 0xb7, 0xcc, 0xec, 0xad, 0xa6, 0xe5, 0x77, 0xcf, 0x80, 0xf7, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x61, 0x61, 0x36, 0x9e, 0xec, 0x01, 0x00, 0x00,
}

func (m *Consenter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Consenter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Consenter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RestPort != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.RestPort))
		i--
		dAtA[i] = 0x40
	}
	if m.ReqPort != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.ReqPort))
		i--
		dAtA[i] = 0x38
	}
	if len(m.TeeCert) > 0 {
		i -= len(m.TeeCert)
		copy(dAtA[i:], m.TeeCert)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.TeeCert)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.ReeCert) > 0 {
		i -= len(m.ReeCert)
		copy(dAtA[i:], m.ReeCert)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ReeCert)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Port != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Host) > 0 {
		i -= len(m.Host)
		copy(dAtA[i:], m.Host)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Host)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Org) > 0 {
		i -= len(m.Org)
		copy(dAtA[i:], m.Org)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Org)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintConfig(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Consenter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.Org)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.Host)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovConfig(uint64(m.Port))
	}
	l = len(m.ReeCert)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	l = len(m.TeeCert)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.ReqPort != 0 {
		n += 1 + sovConfig(uint64(m.ReqPort))
	}
	if m.RestPort != 0 {
		n += 1 + sovConfig(uint64(m.RestPort))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Consenter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: Consenter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Consenter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Org", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Org = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Host", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Host = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReeCert", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReeCert = append(m.ReeCert[:0], dAtA[iNdEx:postIndex]...)
			if m.ReeCert == nil {
				m.ReeCert = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TeeCert", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TeeCert = append(m.TeeCert[:0], dAtA[iNdEx:postIndex]...)
			if m.TeeCert == nil {
				m.TeeCert = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqPort", wireType)
			}
			m.ReqPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReqPort |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RestPort", wireType)
			}
			m.RestPort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RestPort |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)
