// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/streamswap/v1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params holds parameters for the streamswap module
type Params struct {
	// fee charged when creating a new sale. The fee will go to the
	// sale_fee_recipient unless it is not defined (empty).
	SaleCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=sale_creation_fee,json=saleCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"sale_creation_fee" yaml:"sale_creation_fee"`
	// bech32 address of the fee recipient
	SaleCreationFeeRecipient string `protobuf:"bytes,2,opt,name=sale_creation_fee_recipient,json=saleCreationFeeRecipient,proto3" json:"sale_creation_fee_recipient,omitempty"`
	// minimum amount duration of time between the sale creation and the sale
	// start time.
	MinDurationUntilStartTime time.Duration `protobuf:"bytes,3,opt,name=min_duration_until_start_time,json=minDurationUntilStartTime,proto3,stdduration" json:"min_duration_until_start_time"`
	// minimum duration for every new sale.
	MinSaleDuration time.Duration `protobuf:"bytes,4,opt,name=min_sale_duration,json=minSaleDuration,proto3,stdduration" json:"min_sale_duration"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_e87cd402886b5ef9, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetSaleCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.SaleCreationFee
	}
	return nil
}

func (m *Params) GetSaleCreationFeeRecipient() string {
	if m != nil {
		return m.SaleCreationFeeRecipient
	}
	return ""
}

func (m *Params) GetMinDurationUntilStartTime() time.Duration {
	if m != nil {
		return m.MinDurationUntilStartTime
	}
	return 0
}

func (m *Params) GetMinSaleDuration() time.Duration {
	if m != nil {
		return m.MinSaleDuration
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "osmosis.streamswap.v1.Params")
}

func init() {
	proto.RegisterFile("osmosis/streamswap/v1/params.proto", fileDescriptor_e87cd402886b5ef9)
}

var fileDescriptor_e87cd402886b5ef9 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0xcf, 0xb8, 0x52, 0x34, 0x1e, 0x4a, 0x83, 0x42, 0xb6, 0x62, 0xb2, 0xe4, 0xb4, 0x97, 0xce,
	0x10, 0x05, 0x0f, 0x82, 0x97, 0x56, 0x3c, 0x09, 0x96, 0x54, 0x2f, 0x5e, 0x86, 0x49, 0xfa, 0x1a,
	0x07, 0x33, 0x33, 0x21, 0x33, 0x89, 0xee, 0xb7, 0x10, 0x44, 0xf0, 0x33, 0xf8, 0x49, 0x7a, 0xec,
	0xd1, 0x53, 0x2b, 0xbb, 0xdf, 0xc0, 0x4f, 0x20, 0x99, 0xcc, 0x68, 0xe9, 0x5e, 0x7a, 0x4a, 0xde,
	0xfc, 0xfe, 0xbd, 0xf7, 0x78, 0x61, 0xa6, 0xb4, 0x50, 0x9a, 0x6b, 0xa2, 0x4d, 0x07, 0x4c, 0xe8,
	0xcf, 0xac, 0x25, 0x43, 0x4e, 0x5a, 0xd6, 0x31, 0xa1, 0x71, 0xdb, 0x29, 0xa3, 0xa2, 0x47, 0x8e,
	0x83, 0xff, 0x73, 0xf0, 0x90, 0xef, 0x3f, 0xac, 0x55, 0xad, 0x2c, 0x83, 0x8c, 0x7f, 0x13, 0x79,
	0x7f, 0x5e, 0x2b, 0x55, 0x37, 0x40, 0x6c, 0x55, 0xf6, 0x67, 0x84, 0xc9, 0x95, 0x87, 0x2a, 0x6b,
	0x44, 0x27, 0xcd, 0x54, 0x38, 0x28, 0x99, 0x2a, 0x52, 0x32, 0x0d, 0x64, 0xc8, 0x4b, 0x30, 0x2c,
	0x27, 0x95, 0xe2, 0xd2, 0xe3, 0x37, 0x5d, 0x4f, 0xfb, 0x8e, 0x19, 0xae, 0x1c, 0x9e, 0x7d, 0x9f,
	0x85, 0x3b, 0xc7, 0xb6, 0xe7, 0xe8, 0x1b, 0x0a, 0xf7, 0x34, 0x6b, 0x80, 0x56, 0x1d, 0x58, 0x0a,
	0x3d, 0x03, 0x88, 0xd1, 0x62, 0xb6, 0x7c, 0xf0, 0x74, 0x8e, 0x5d, 0xea, 0x98, 0x83, 0x5d, 0x0e,
	0x3e, 0x52, 0x5c, 0x1e, 0xbe, 0x39, 0xbf, 0x4c, 0x83, 0x3f, 0x97, 0x69, 0xbc, 0x62, 0xa2, 0x79,
	0x91, 0x6d, 0x39, 0x64, 0x3f, 0xaf, 0xd2, 0x65, 0xcd, 0xcd, 0xc7, 0xbe, 0xc4, 0x95, 0x12, 0xae,
	0x7d, 0xf7, 0x39, 0xd0, 0xa7, 0x9f, 0x88, 0x59, 0xb5, 0xa0, 0xad, 0x99, 0x2e, 0x76, 0x47, 0xfd,
	0x91, 0x93, 0xbf, 0x06, 0x88, 0x5e, 0x86, 0x8f, 0xb7, 0x2c, 0x69, 0x07, 0x15, 0x6f, 0x39, 0x48,
	0x13, 0xdf, 0x59, 0xa0, 0xe5, 0xfd, 0x22, 0xbe, 0xa1, 0x2a, 0x3c, 0x1e, 0x41, 0xf8, 0x44, 0x70,
	0x49, 0xfd, 0xd4, 0xb4, 0x97, 0x86, 0x37, 0x54, 0x1b, 0xd6, 0x19, 0x6a, 0xb8, 0x80, 0x78, 0xb6,
	0x40, 0x76, 0xbe, 0x69, 0x4f, 0xd8, 0xef, 0x09, 0xbf, 0x72, 0x8a, 0xc3, 0x7b, 0xe3, 0x7c, 0x3f,
	0xae, 0x52, 0x54, 0xcc, 0x05, 0x97, 0xfe, 0xf9, 0xfd, 0xe8, 0x73, 0x32, 0xda, 0xbc, 0xe3, 0x02,
	0xa2, 0xb7, 0xe1, 0xde, 0x18, 0x63, 0x3b, 0xf5, 0x59, 0xf1, 0xdd, 0xdb, 0x5b, 0xef, 0x0a, 0x2e,
	0x4f, 0x58, 0x03, 0xff, 0xa0, 0xe3, 0xf3, 0x75, 0x82, 0x2e, 0xd6, 0x09, 0xfa, 0xbd, 0x4e, 0xd0,
	0xd7, 0x4d, 0x12, 0x5c, 0x6c, 0x92, 0xe0, 0xd7, 0x26, 0x09, 0x3e, 0x3c, 0xbf, 0xb6, 0x4b, 0x77,
	0x5f, 0x07, 0x0d, 0x2b, 0xb5, 0x2f, 0xc8, 0x90, 0xe7, 0xe4, 0xcb, 0xf5, 0xb3, 0xb4, 0xfb, 0x2d,
	0x77, 0x6c, 0xfe, 0xb3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x6b, 0x0a, 0x7a, 0xb9, 0x02,
	0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.MinSaleDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.MinSaleDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.MinDurationUntilStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.MinDurationUntilStartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintParams(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	if len(m.SaleCreationFeeRecipient) > 0 {
		i -= len(m.SaleCreationFeeRecipient)
		copy(dAtA[i:], m.SaleCreationFeeRecipient)
		i = encodeVarintParams(dAtA, i, uint64(len(m.SaleCreationFeeRecipient)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SaleCreationFee) > 0 {
		for iNdEx := len(m.SaleCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SaleCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SaleCreationFee) > 0 {
		for _, e := range m.SaleCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	l = len(m.SaleCreationFeeRecipient)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.MinDurationUntilStartTime)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.MinSaleDuration)
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SaleCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SaleCreationFee = append(m.SaleCreationFee, types.Coin{})
			if err := m.SaleCreationFee[len(m.SaleCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SaleCreationFeeRecipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SaleCreationFeeRecipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDurationUntilStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.MinDurationUntilStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSaleDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.MinSaleDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
