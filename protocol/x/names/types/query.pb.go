// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: klyraprotocol/names/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// Queries a Name by id.
type QueryNameRequest struct {
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryNameRequest) Reset()         { *m = QueryNameRequest{} }
func (m *QueryNameRequest) String() string { return proto.CompactTextString(m) }
func (*QueryNameRequest) ProtoMessage()    {}
func (*QueryNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7410d881f070f6e, []int{0}
}
func (m *QueryNameRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryNameRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryNameRequest.Merge(m, src)
}
func (m *QueryNameRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryNameRequest proto.InternalMessageInfo

func (m *QueryNameRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// QueryNameResponse is response type for the Name RPC method.
type QueryNameResponse struct {
	Name Name `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
}

func (m *QueryNameResponse) Reset()         { *m = QueryNameResponse{} }
func (m *QueryNameResponse) String() string { return proto.CompactTextString(m) }
func (*QueryNameResponse) ProtoMessage()    {}
func (*QueryNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7410d881f070f6e, []int{1}
}
func (m *QueryNameResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryNameResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryNameResponse.Merge(m, src)
}
func (m *QueryNameResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryNameResponse proto.InternalMessageInfo

func (m *QueryNameResponse) GetName() Name {
	if m != nil {
		return m.Name
	}
	return Name{}
}

// Queries a list of Name items.
type QueryAllNamesRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllNamesRequest) Reset()         { *m = QueryAllNamesRequest{} }
func (m *QueryAllNamesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllNamesRequest) ProtoMessage()    {}
func (*QueryAllNamesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7410d881f070f6e, []int{2}
}
func (m *QueryAllNamesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllNamesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllNamesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllNamesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllNamesRequest.Merge(m, src)
}
func (m *QueryAllNamesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllNamesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllNamesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllNamesRequest proto.InternalMessageInfo

func (m *QueryAllNamesRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryAllNamesResponse is response type for the AllNames RPC method.
type QueryAllNamesResponse struct {
	Name       []Name              `protobuf:"bytes,1,rep,name=name,proto3" json:"name"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllNamesResponse) Reset()         { *m = QueryAllNamesResponse{} }
func (m *QueryAllNamesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllNamesResponse) ProtoMessage()    {}
func (*QueryAllNamesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7410d881f070f6e, []int{3}
}
func (m *QueryAllNamesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllNamesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllNamesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllNamesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllNamesResponse.Merge(m, src)
}
func (m *QueryAllNamesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllNamesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllNamesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllNamesResponse proto.InternalMessageInfo

func (m *QueryAllNamesResponse) GetName() []Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *QueryAllNamesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryNameRequest)(nil), "klyraprotocol.names.QueryNameRequest")
	proto.RegisterType((*QueryNameResponse)(nil), "klyraprotocol.names.QueryNameResponse")
	proto.RegisterType((*QueryAllNamesRequest)(nil), "klyraprotocol.names.QueryAllNamesRequest")
	proto.RegisterType((*QueryAllNamesResponse)(nil), "klyraprotocol.names.QueryAllNamesResponse")
}

func init() { proto.RegisterFile("klyraprotocol/names/query.proto", fileDescriptor_d7410d881f070f6e) }

var fileDescriptor_d7410d881f070f6e = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x93, 0x58, 0x45, 0x46, 0x14, 0x1d, 0xaf, 0xe0, 0x8d, 0x32, 0xf7, 0x1a, 0xb0, 0xea,
	0x85, 0x3b, 0x43, 0xdb, 0x27, 0xb0, 0x60, 0x75, 0x25, 0x35, 0xee, 0x5c, 0x28, 0x93, 0x74, 0x48,
	0x07, 0x93, 0x99, 0x34, 0x33, 0x15, 0x8b, 0x74, 0x23, 0x6e, 0x05, 0x41, 0x7c, 0xa7, 0x2e, 0x0b,
	0x6e, 0x5c, 0x89, 0xb4, 0x3e, 0x88, 0xcc, 0x24, 0xd1, 0xb6, 0xb6, 0xb4, 0x9b, 0x10, 0xce, 0xf9,
	0xe7, 0xff, 0xbf, 0x73, 0x0e, 0x38, 0x79, 0x9b, 0x4e, 0x0a, 0x9a, 0x17, 0x52, 0xcb, 0x58, 0xa6,
	0x44, 0xd0, 0x8c, 0x29, 0x32, 0x1a, 0xb3, 0x62, 0x82, 0x6d, 0x11, 0xde, 0x5c, 0x13, 0x60, 0x2b,
	0xf0, 0x8f, 0x12, 0x99, 0x48, 0x5b, 0x23, 0xe6, 0xaf, 0x94, 0xfa, 0x77, 0x13, 0x29, 0x93, 0x94,
	0x11, 0x9a, 0x73, 0x42, 0x85, 0x90, 0x9a, 0x6a, 0x2e, 0x85, 0xaa, 0xba, 0x67, 0xb1, 0x54, 0x99,
	0x54, 0x24, 0xa2, 0x8a, 0x95, 0x09, 0xe4, 0x5d, 0x2b, 0x62, 0x9a, 0xb6, 0x48, 0x4e, 0x13, 0x2e,
	0xac, 0xb8, 0xd2, 0xa2, 0x6d, 0x54, 0xe6, 0x5b, 0xf6, 0x83, 0x00, 0x5c, 0x7f, 0x61, 0x1c, 0x9e,
	0xd3, 0x8c, 0x85, 0x6c, 0x34, 0x66, 0x4a, 0xc3, 0x6b, 0xc0, 0xe3, 0x83, 0xdb, 0xee, 0xa9, 0xfb,
	0xf0, 0x6a, 0xe8, 0xf1, 0x41, 0xf0, 0x0c, 0xdc, 0x58, 0xd1, 0xa8, 0x5c, 0x0a, 0xc5, 0x60, 0x07,
	0x34, 0x8c, 0x8d, 0x95, 0x5d, 0x69, 0x1f, 0xe3, 0x2d, 0xc3, 0x61, 0xf3, 0xa0, 0xdb, 0x98, 0xfd,
	0x3c, 0x71, 0x42, 0x2b, 0x0e, 0x5e, 0x83, 0x23, 0xeb, 0xf4, 0x38, 0x4d, 0x4d, 0x4f, 0xd5, 0x89,
	0x3d, 0x00, 0xfe, 0x91, 0x57, 0x96, 0x4d, 0x5c, 0x8e, 0x89, 0xcd, 0x98, 0xb8, 0x5c, 0x64, 0x35,
	0x26, 0xee, 0xd3, 0xa4, 0xa6, 0x0d, 0x57, 0x5e, 0x06, 0xdf, 0x5c, 0x70, 0x6b, 0x23, 0xe0, 0x3f,
	0xdc, 0x0b, 0x07, 0xe3, 0xc2, 0xa7, 0x6b, 0x58, 0x9e, 0xc5, 0x7a, 0xb0, 0x17, 0xab, 0x4c, 0x5c,
	0xe5, 0x6a, 0x7f, 0xf6, 0xc0, 0x45, 0xcb, 0x05, 0xa7, 0xa0, 0x61, 0x62, 0xe0, 0xfd, 0xad, 0x04,
	0x9b, 0xa7, 0xf0, 0x9b, 0xfb, 0x64, 0x65, 0x58, 0xd0, 0xfc, 0xf8, 0xfd, 0xf7, 0x57, 0xef, 0x14,
	0x22, 0xb2, 0xeb, 0xde, 0xe4, 0x03, 0x1f, 0x4c, 0xe1, 0x27, 0x17, 0x5c, 0xae, 0x77, 0x03, 0x1f,
	0xed, 0x36, 0xdf, 0x38, 0x90, 0x7f, 0x76, 0x88, 0xb4, 0x62, 0xb9, 0x67, 0x59, 0xee, 0xc0, 0xe3,
	0x9d, 0x2c, 0xdd, 0x37, 0xb3, 0x05, 0x72, 0xe7, 0x0b, 0xe4, 0xfe, 0x5a, 0x20, 0xf7, 0xcb, 0x12,
	0x39, 0xf3, 0x25, 0x72, 0x7e, 0x2c, 0x91, 0xf3, 0xea, 0x49, 0xc2, 0xf5, 0x70, 0x1c, 0xe1, 0x58,
	0x66, 0xe4, 0xa5, 0x2e, 0x18, 0xcd, 0x7a, 0x5c, 0x50, 0x11, 0xb3, 0xf3, 0x7e, 0xed, 0xa3, 0x6c,
	0xf9, 0x3c, 0x1e, 0x52, 0x2e, 0xc8, 0x5f, 0xf7, 0xf7, 0x95, 0xbf, 0x9e, 0xe4, 0x4c, 0x45, 0x97,
	0x6c, 0xbd, 0xf3, 0x27, 0x00, 0x00, 0xff, 0xff, 0x31, 0x82, 0x92, 0x90, 0x95, 0x03, 0x00, 0x00,
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
	// Queries a Name by id.
	Name(ctx context.Context, in *QueryNameRequest, opts ...grpc.CallOption) (*QueryNameResponse, error)
	// Queries a list of Name items.
	AllNames(ctx context.Context, in *QueryAllNamesRequest, opts ...grpc.CallOption) (*QueryAllNamesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Name(ctx context.Context, in *QueryNameRequest, opts ...grpc.CallOption) (*QueryNameResponse, error) {
	out := new(QueryNameResponse)
	err := c.cc.Invoke(ctx, "/klyraprotocol.names.Query/Name", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllNames(ctx context.Context, in *QueryAllNamesRequest, opts ...grpc.CallOption) (*QueryAllNamesResponse, error) {
	out := new(QueryAllNamesResponse)
	err := c.cc.Invoke(ctx, "/klyraprotocol.names.Query/AllNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a Name by id.
	Name(context.Context, *QueryNameRequest) (*QueryNameResponse, error)
	// Queries a list of Name items.
	AllNames(context.Context, *QueryAllNamesRequest) (*QueryAllNamesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Name(ctx context.Context, req *QueryNameRequest) (*QueryNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Name not implemented")
}
func (*UnimplementedQueryServer) AllNames(ctx context.Context, req *QueryAllNamesRequest) (*QueryAllNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllNames not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/klyraprotocol.names.Query/Name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Name(ctx, req.(*QueryNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/klyraprotocol.names.Query/AllNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllNames(ctx, req.(*QueryAllNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "klyraprotocol.names.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Name",
			Handler:    _Query_Name_Handler,
		},
		{
			MethodName: "AllNames",
			Handler:    _Query_AllNames_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "klyraprotocol/names/query.proto",
}

func (m *QueryNameRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryNameRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryNameRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryNameResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryNameResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryNameResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Name.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllNamesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllNamesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllNamesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllNamesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllNamesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllNamesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		for iNdEx := len(m.Name) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Name[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryNameRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryNameResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Name.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllNamesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllNamesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Name) > 0 {
		for _, e := range m.Name {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryNameRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryNameRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryNameRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
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
func (m *QueryNameResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryNameResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryNameResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			if err := m.Name.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllNamesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllNamesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllNamesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllNamesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllNamesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllNamesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = append(m.Name, Name{})
			if err := m.Name[len(m.Name)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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