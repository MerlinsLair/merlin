// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: merlin/downtime-detector/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
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

type GenesisDowntimeEntry struct {
	Duration     Downtime  `protobuf:"varint,1,opt,name=duration,proto3,enum=merlin.downtimedetector.v1beta1.Downtime" json:"duration,omitempty" yaml:"duration"`
	LastDowntime time.Time `protobuf:"bytes,2,opt,name=last_downtime,json=lastDowntime,proto3,stdtime" json:"last_downtime" yaml:"last_downtime"`
}

func (m *GenesisDowntimeEntry) Reset()         { *m = GenesisDowntimeEntry{} }
func (m *GenesisDowntimeEntry) String() string { return proto.CompactTextString(m) }
func (*GenesisDowntimeEntry) ProtoMessage()    {}
func (*GenesisDowntimeEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_4581e137a44782af, []int{0}
}
func (m *GenesisDowntimeEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisDowntimeEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisDowntimeEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisDowntimeEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisDowntimeEntry.Merge(m, src)
}
func (m *GenesisDowntimeEntry) XXX_Size() int {
	return m.Size()
}
func (m *GenesisDowntimeEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisDowntimeEntry.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisDowntimeEntry proto.InternalMessageInfo

func (m *GenesisDowntimeEntry) GetDuration() Downtime {
	if m != nil {
		return m.Duration
	}
	return Downtime_DURATION_30S
}

func (m *GenesisDowntimeEntry) GetLastDowntime() time.Time {
	if m != nil {
		return m.LastDowntime
	}
	return time.Time{}
}

// GenesisState defines the twap module's genesis state.
type GenesisState struct {
	Downtimes     []GenesisDowntimeEntry `protobuf:"bytes,1,rep,name=downtimes,proto3" json:"downtimes"`
	LastBlockTime time.Time              `protobuf:"bytes,2,opt,name=last_block_time,json=lastBlockTime,proto3,stdtime" json:"last_block_time" yaml:"last_block_time"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4581e137a44782af, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetDowntimes() []GenesisDowntimeEntry {
	if m != nil {
		return m.Downtimes
	}
	return nil
}

func (m *GenesisState) GetLastBlockTime() time.Time {
	if m != nil {
		return m.LastBlockTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*GenesisDowntimeEntry)(nil), "merlin.downtimedetector.v1beta1.GenesisDowntimeEntry")
	proto.RegisterType((*GenesisState)(nil), "merlin.downtimedetector.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("merlin/downtime-detector/v1beta1/genesis.proto", fileDescriptor_4581e137a44782af)
}

var fileDescriptor_4581e137a44782af = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6a, 0xe2, 0x40,
	0x1c, 0xc6, 0x33, 0xbb, 0xcb, 0xb2, 0x1b, 0xdd, 0x15, 0xb2, 0xb2, 0x58, 0x0f, 0x49, 0xc8, 0x49,
	0x0a, 0xce, 0xa0, 0x2d, 0x85, 0x16, 0x7a, 0x09, 0x2d, 0xbd, 0xdb, 0x42, 0xc1, 0x1e, 0xc2, 0x24,
	0x8e, 0x69, 0x68, 0x92, 0x91, 0xcc, 0x68, 0x9b, 0xb7, 0xf0, 0xb1, 0x3c, 0x4a, 0x0f, 0xa5, 0x27,
	0x5b, 0xf4, 0x0d, 0x7c, 0x82, 0x92, 0x64, 0x46, 0xab, 0x08, 0xf6, 0x96, 0xcc, 0xff, 0xfb, 0x3e,
	0xbe, 0xdf, 0x9f, 0xbf, 0x8a, 0x28, 0x8b, 0x28, 0x0b, 0x18, 0xea, 0xd1, 0xc7, 0x98, 0x07, 0x11,
	0x69, 0xf6, 0x08, 0x27, 0x1e, 0xa7, 0x09, 0x1a, 0xb5, 0x5c, 0xc2, 0x71, 0x0b, 0xf9, 0x24, 0x26,
	0x2c, 0x60, 0x70, 0x90, 0x50, 0x4e, 0x35, 0x53, 0x18, 0xa0, 0x34, 0x48, 0x3d, 0x14, 0xfa, 0x7a,
	0xd5, 0xa7, 0x3e, 0xcd, 0xc5, 0x28, 0xfb, 0x2a, 0x7c, 0xf5, 0x03, 0x9f, 0x52, 0x3f, 0x24, 0x28,
	0xff, 0x73, 0x87, 0x7d, 0x84, 0xe3, 0x54, 0x8e, 0xbc, 0x3c, 0xd3, 0x29, 0x3c, 0xc5, 0x8f, 0x18,
	0xe9, 0xdb, 0xae, 0xde, 0x30, 0xc1, 0x3c, 0xa0, 0xb1, 0x98, 0x1b, 0xdb, 0xf3, 0xac, 0x11, 0xe3,
	0x38, 0x1a, 0x08, 0xc1, 0xe9, 0x7e, 0x3e, 0x39, 0x71, 0x36, 0xb3, 0xad, 0x17, 0xa0, 0x56, 0xaf,
	0x0a, 0xf6, 0x0b, 0x21, 0xb9, 0x8c, 0x79, 0x92, 0x6a, 0x77, 0xea, 0x2f, 0x29, 0xad, 0x01, 0x13,
	0x34, 0xfe, 0xb6, 0x0f, 0xe1, 0xbe, 0xad, 0x40, 0x19, 0x61, 0xff, 0x5b, 0xce, 0x8c, 0x4a, 0x8a,
	0xa3, 0xf0, 0xcc, 0x92, 0x29, 0x56, 0x67, 0x15, 0xa8, 0x61, 0xf5, 0x4f, 0x88, 0x19, 0x77, 0x64,
	0x50, 0xed, 0x9b, 0x09, 0x1a, 0xa5, 0x76, 0x1d, 0x16, 0xa4, 0x50, 0x92, 0xc2, 0x1b, 0x49, 0x6a,
	0x9b, 0x93, 0x99, 0xa1, 0x2c, 0x67, 0x46, 0xb5, 0x48, 0xdd, 0xb0, 0x5b, 0xe3, 0x37, 0x03, 0x74,
	0xca, 0xd9, 0x9b, 0x6c, 0x60, 0x3d, 0x03, 0xb5, 0x2c, 0xc0, 0xae, 0x39, 0xe6, 0x44, 0xeb, 0xaa,
	0xbf, 0xa5, 0x9e, 0xd5, 0x80, 0xf9, 0xbd, 0x51, 0x6a, 0x9f, 0xec, 0x27, 0xda, 0xb5, 0x1b, 0xfb,
	0x47, 0xd6, 0xa5, 0xb3, 0x8e, 0xd3, 0xfa, 0x6a, 0x25, 0x2f, 0xe4, 0x86, 0xd4, 0x7b, 0x70, 0xbe,
	0x48, 0x64, 0x09, 0xa2, 0xff, 0x9f, 0x88, 0xd6, 0x01, 0x05, 0x53, 0xbe, 0x26, 0x3b, 0x7b, 0xcc,
	0x7c, 0xf6, 0xed, 0x64, 0xae, 0x83, 0xe9, 0x5c, 0x07, 0xef, 0x73, 0x1d, 0x8c, 0x17, 0xba, 0x32,
	0x5d, 0xe8, 0xca, 0xeb, 0x42, 0x57, 0xba, 0xe7, 0x7e, 0xc0, 0xef, 0x87, 0x2e, 0xf4, 0x68, 0x24,
	0xaf, 0xbd, 0x19, 0x62, 0x97, 0xad, 0x4e, 0x7f, 0xd4, 0x3a, 0x46, 0x4f, 0x3b, 0x0e, 0x84, 0xa7,
	0x03, 0xc2, 0xdc, 0x9f, 0x79, 0xbf, 0xa3, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0xde, 0x42,
	0xac, 0x2a, 0x03, 0x00, 0x00,
}

func (m *GenesisDowntimeEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisDowntimeEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisDowntimeEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastDowntime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastDowntime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if m.Duration != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastBlockTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastBlockTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.Downtimes) > 0 {
		for iNdEx := len(m.Downtimes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Downtimes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
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
func (m *GenesisDowntimeEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Duration != 0 {
		n += 1 + sovGenesis(uint64(m.Duration))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastDowntime)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Downtimes) > 0 {
		for _, e := range m.Downtimes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastBlockTime)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisDowntimeEntry) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisDowntimeEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisDowntimeEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= Downtime(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastDowntime", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastDowntime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Downtimes", wireType)
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
			m.Downtimes = append(m.Downtimes, GenesisDowntimeEntry{})
			if err := m.Downtimes[len(m.Downtimes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockTime", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
