// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: merlin/epochs/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type QueryEpochsInfoRequest struct {
}

func (m *QueryEpochsInfoRequest) Reset()         { *m = QueryEpochsInfoRequest{} }
func (m *QueryEpochsInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEpochsInfoRequest) ProtoMessage()    {}
func (*QueryEpochsInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_574bd176519c765f, []int{0}
}
func (m *QueryEpochsInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEpochsInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEpochsInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEpochsInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEpochsInfoRequest.Merge(m, src)
}
func (m *QueryEpochsInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEpochsInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEpochsInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEpochsInfoRequest proto.InternalMessageInfo

type QueryEpochsInfoResponse struct {
	Epochs []EpochInfo `protobuf:"bytes,1,rep,name=epochs,proto3" json:"epochs"`
}

func (m *QueryEpochsInfoResponse) Reset()         { *m = QueryEpochsInfoResponse{} }
func (m *QueryEpochsInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEpochsInfoResponse) ProtoMessage()    {}
func (*QueryEpochsInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_574bd176519c765f, []int{1}
}
func (m *QueryEpochsInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEpochsInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEpochsInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEpochsInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEpochsInfoResponse.Merge(m, src)
}
func (m *QueryEpochsInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEpochsInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEpochsInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEpochsInfoResponse proto.InternalMessageInfo

func (m *QueryEpochsInfoResponse) GetEpochs() []EpochInfo {
	if m != nil {
		return m.Epochs
	}
	return nil
}

type QueryCurrentEpochRequest struct {
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (m *QueryCurrentEpochRequest) Reset()         { *m = QueryCurrentEpochRequest{} }
func (m *QueryCurrentEpochRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCurrentEpochRequest) ProtoMessage()    {}
func (*QueryCurrentEpochRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_574bd176519c765f, []int{2}
}
func (m *QueryCurrentEpochRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrentEpochRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrentEpochRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrentEpochRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrentEpochRequest.Merge(m, src)
}
func (m *QueryCurrentEpochRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrentEpochRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrentEpochRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrentEpochRequest proto.InternalMessageInfo

func (m *QueryCurrentEpochRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type QueryCurrentEpochResponse struct {
	CurrentEpoch int64 `protobuf:"varint,1,opt,name=current_epoch,json=currentEpoch,proto3" json:"current_epoch,omitempty"`
}

func (m *QueryCurrentEpochResponse) Reset()         { *m = QueryCurrentEpochResponse{} }
func (m *QueryCurrentEpochResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCurrentEpochResponse) ProtoMessage()    {}
func (*QueryCurrentEpochResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_574bd176519c765f, []int{3}
}
func (m *QueryCurrentEpochResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCurrentEpochResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCurrentEpochResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCurrentEpochResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCurrentEpochResponse.Merge(m, src)
}
func (m *QueryCurrentEpochResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCurrentEpochResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCurrentEpochResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCurrentEpochResponse proto.InternalMessageInfo

func (m *QueryCurrentEpochResponse) GetCurrentEpoch() int64 {
	if m != nil {
		return m.CurrentEpoch
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryEpochsInfoRequest)(nil), "merlin.epochs.v1beta1.QueryEpochsInfoRequest")
	proto.RegisterType((*QueryEpochsInfoResponse)(nil), "merlin.epochs.v1beta1.QueryEpochsInfoResponse")
	proto.RegisterType((*QueryCurrentEpochRequest)(nil), "merlin.epochs.v1beta1.QueryCurrentEpochRequest")
	proto.RegisterType((*QueryCurrentEpochResponse)(nil), "merlin.epochs.v1beta1.QueryCurrentEpochResponse")
}

func init() { proto.RegisterFile("merlin/epochs/query.proto", fileDescriptor_574bd176519c765f) }

var fileDescriptor_574bd176519c765f = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcf, 0x6e, 0xd3, 0x30,
	0x18, 0x8f, 0x5b, 0xa8, 0x84, 0x29, 0x17, 0x0b, 0x95, 0x10, 0x21, 0x53, 0x82, 0x80, 0x0a, 0xa9,
	0x31, 0x29, 0x9c, 0xb8, 0x80, 0x8a, 0x38, 0x20, 0x71, 0x21, 0xc7, 0x5e, 0x50, 0x12, 0xdc, 0xd4,
	0x52, 0x6b, 0xa7, 0xb1, 0x53, 0xd1, 0xeb, 0x9e, 0x60, 0xd2, 0xb4, 0x17, 0xd8, 0xb3, 0xec, 0xd0,
	0x63, 0xa5, 0x5d, 0x76, 0x9a, 0xa6, 0x76, 0x0f, 0x32, 0xc5, 0x49, 0xaa, 0xae, 0x4b, 0xa7, 0xed,
	0x96, 0xf8, 0xf7, 0xf7, 0xfb, 0x6c, 0x68, 0x09, 0x39, 0x11, 0x92, 0x49, 0x42, 0x63, 0x11, 0x8e,
	0x24, 0x99, 0xa6, 0x34, 0x99, 0x3b, 0x71, 0x22, 0x94, 0x40, 0xad, 0x02, 0x73, 0x72, 0xcc, 0x99,
	0xb9, 0x01, 0x55, 0xbe, 0x6b, 0x3d, 0x8f, 0x44, 0x24, 0x34, 0x85, 0x64, 0x5f, 0x39, 0xdb, 0x7a,
	0x15, 0x09, 0x11, 0x8d, 0x29, 0xf1, 0x63, 0x46, 0x7c, 0xce, 0x85, 0xf2, 0x15, 0x13, 0x5c, 0x16,
	0xe8, 0xc7, 0x50, 0x9b, 0x91, 0xc0, 0x97, 0x34, 0x0f, 0x21, 0x85, 0x1d, 0x89, 0xfd, 0x88, 0x71,
	0x4d, 0x2e, 0x9d, 0x76, 0x3a, 0x45, 0x94, 0xd3, 0xac, 0x86, 0x46, 0x6d, 0x13, 0xb6, 0xfe, 0x64,
	0xfa, 0x9f, 0x1a, 0xfc, 0xc5, 0x87, 0xc2, 0xa3, 0xd3, 0x94, 0x4a, 0x65, 0x0f, 0xe0, 0x8b, 0x5b,
	0x88, 0x8c, 0x05, 0x97, 0x14, 0x7d, 0x83, 0x8d, 0xdc, 0xcc, 0x04, 0xed, 0x7a, 0xe7, 0x69, 0xef,
	0x8d, 0x53, 0x3d, 0x9b, 0xa3, 0xb5, 0x99, 0xb4, 0xff, 0x68, 0x71, 0xf1, 0xda, 0xf0, 0x0a, 0x99,
	0xfd, 0x15, 0x9a, 0xda, 0xfb, 0x47, 0x9a, 0x24, 0x94, 0x2b, 0x4d, 0x2b, 0x72, 0x11, 0x86, 0x90,
	0xfd, 0xa3, 0x5c, 0xb1, 0x21, 0xa3, 0x89, 0x09, 0xda, 0xa0, 0xf3, 0xc4, 0xdb, 0x3a, 0xb1, 0xbf,
	0xc3, 0x97, 0x15, 0xda, 0xa2, 0xd9, 0x5b, 0xf8, 0x2c, 0xcc, 0xcf, 0xff, 0xea, 0x28, 0xad, 0xaf,
	0x7b, 0xcd, 0x70, 0x8b, 0xdc, 0x3b, 0xad, 0xc1, 0xc7, 0xda, 0x02, 0x1d, 0x03, 0x08, 0x37, 0x1d,
	0x25, 0x72, 0xf6, 0xcd, 0x51, 0xbd, 0x22, 0x8b, 0xdc, 0x9b, 0x9f, 0xd7, 0xb3, 0xdf, 0x1f, 0x9c,
	0x5d, 0x1d, 0xd5, 0xda, 0x08, 0x93, 0x9d, 0x4b, 0x29, 0x6f, 0x2f, 0xff, 0x45, 0x27, 0x00, 0x36,
	0xb7, 0xe7, 0x43, 0x9f, 0xee, 0x4c, 0xaa, 0x58, 0xa3, 0xe5, 0x3e, 0x40, 0x51, 0xb4, 0xeb, 0xea,
	0x76, 0x1f, 0xd0, 0xbb, 0x7d, 0xed, 0x6e, 0xac, 0xb6, 0xff, 0x7b, 0xb1, 0xc2, 0x60, 0xb9, 0xc2,
	0xe0, 0x72, 0x85, 0xc1, 0xe1, 0x1a, 0x1b, 0xcb, 0x35, 0x36, 0xce, 0xd7, 0xd8, 0x18, 0xf4, 0x22,
	0xa6, 0x46, 0x69, 0xe0, 0x84, 0x62, 0x52, 0x5a, 0x75, 0xc7, 0x7e, 0x20, 0x37, 0xbe, 0x33, 0xf7,
	0x0b, 0xf9, 0x5f, 0xba, 0xab, 0x79, 0x4c, 0x65, 0xd0, 0xd0, 0xef, 0xf1, 0xf3, 0x75, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x05, 0xab, 0x92, 0xb5, 0x43, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// EpochInfos provide running epochInfos
	EpochInfos(ctx context.Context, in *QueryEpochsInfoRequest, opts ...grpc.CallOption) (*QueryEpochsInfoResponse, error)
	// CurrentEpoch provide current epoch of specified identifier
	CurrentEpoch(ctx context.Context, in *QueryCurrentEpochRequest, opts ...grpc.CallOption) (*QueryCurrentEpochResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) EpochInfos(ctx context.Context, in *QueryEpochsInfoRequest, opts ...grpc.CallOption) (*QueryEpochsInfoResponse, error) {
	out := new(QueryEpochsInfoResponse)
	err := c.cc.Invoke(ctx, "/merlin.epochs.v1beta1.Query/EpochInfos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CurrentEpoch(ctx context.Context, in *QueryCurrentEpochRequest, opts ...grpc.CallOption) (*QueryCurrentEpochResponse, error) {
	out := new(QueryCurrentEpochResponse)
	err := c.cc.Invoke(ctx, "/merlin.epochs.v1beta1.Query/CurrentEpoch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// EpochInfos provide running epochInfos
	EpochInfos(context.Context, *QueryEpochsInfoRequest) (*QueryEpochsInfoResponse, error)
	// CurrentEpoch provide current epoch of specified identifier
	CurrentEpoch(context.Context, *QueryCurrentEpochRequest) (*QueryCurrentEpochResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) EpochInfos(ctx context.Context, req *QueryEpochsInfoRequest) (*QueryEpochsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EpochInfos not implemented")
}
func (*UnimplementedQueryServer) CurrentEpoch(ctx context.Context, req *QueryCurrentEpochRequest) (*QueryCurrentEpochResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentEpoch not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_EpochInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEpochsInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EpochInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/merlin.epochs.v1beta1.Query/EpochInfos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EpochInfos(ctx, req.(*QueryEpochsInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CurrentEpoch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCurrentEpochRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CurrentEpoch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/merlin.epochs.v1beta1.Query/CurrentEpoch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CurrentEpoch(ctx, req.(*QueryCurrentEpochRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "merlin.epochs.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EpochInfos",
			Handler:    _Query_EpochInfos_Handler,
		},
		{
			MethodName: "CurrentEpoch",
			Handler:    _Query_CurrentEpoch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "merlin/epochs/query.proto",
}

func (m *QueryEpochsInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEpochsInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEpochsInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryEpochsInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEpochsInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEpochsInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Epochs) > 0 {
		for iNdEx := len(m.Epochs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Epochs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryCurrentEpochRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrentEpochRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrentEpochRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Identifier) > 0 {
		i -= len(m.Identifier)
		copy(dAtA[i:], m.Identifier)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Identifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCurrentEpochResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCurrentEpochResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCurrentEpochResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CurrentEpoch != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.CurrentEpoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryEpochsInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryEpochsInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Epochs) > 0 {
		for _, e := range m.Epochs {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryCurrentEpochRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCurrentEpochResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CurrentEpoch != 0 {
		n += 1 + sovQuery(uint64(m.CurrentEpoch))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryEpochsInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEpochsInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEpochsInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryEpochsInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEpochsInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEpochsInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epochs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Epochs = append(m.Epochs, EpochInfo{})
			if err := m.Epochs[len(m.Epochs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCurrentEpochRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCurrentEpochRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrentEpochRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCurrentEpochResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCurrentEpochResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCurrentEpochResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpoch", wireType)
			}
			m.CurrentEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
