// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dydxprotocol/vest/query.proto

package types

import (
	context "context"
	fmt "fmt"
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

// QueryVestEntryRequest is a request type for the VestEntry RPC method.
type QueryVestEntryRequest struct {
	VesterAccount string `protobuf:"bytes,1,opt,name=vester_account,json=vesterAccount,proto3" json:"vester_account,omitempty"`
}

func (m *QueryVestEntryRequest) Reset()         { *m = QueryVestEntryRequest{} }
func (m *QueryVestEntryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryVestEntryRequest) ProtoMessage()    {}
func (*QueryVestEntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a660be800e547c7, []int{0}
}
func (m *QueryVestEntryRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVestEntryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVestEntryRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVestEntryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVestEntryRequest.Merge(m, src)
}
func (m *QueryVestEntryRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryVestEntryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVestEntryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVestEntryRequest proto.InternalMessageInfo

func (m *QueryVestEntryRequest) GetVesterAccount() string {
	if m != nil {
		return m.VesterAccount
	}
	return ""
}

// QueryVestEntryResponse is a response type for the VestEntry RPC method.
type QueryVestEntryResponse struct {
	Entry VestEntry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry"`
}

func (m *QueryVestEntryResponse) Reset()         { *m = QueryVestEntryResponse{} }
func (m *QueryVestEntryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryVestEntryResponse) ProtoMessage()    {}
func (*QueryVestEntryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a660be800e547c7, []int{1}
}
func (m *QueryVestEntryResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryVestEntryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryVestEntryResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryVestEntryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryVestEntryResponse.Merge(m, src)
}
func (m *QueryVestEntryResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryVestEntryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryVestEntryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryVestEntryResponse proto.InternalMessageInfo

func (m *QueryVestEntryResponse) GetEntry() VestEntry {
	if m != nil {
		return m.Entry
	}
	return VestEntry{}
}

func init() {
	proto.RegisterType((*QueryVestEntryRequest)(nil), "dydxprotocol.vest.QueryVestEntryRequest")
	proto.RegisterType((*QueryVestEntryResponse)(nil), "dydxprotocol.vest.QueryVestEntryResponse")
}

func init() { proto.RegisterFile("dydxprotocol/vest/query.proto", fileDescriptor_3a660be800e547c7) }

var fileDescriptor_3a660be800e547c7 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xa9, 0x4c, 0xa9,
	0x28, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x4b, 0x2d, 0x2e, 0xd1, 0x2f, 0x2c,
	0x4d, 0x2d, 0xaa, 0xd4, 0x03, 0x8b, 0x09, 0x09, 0x22, 0x4b, 0xeb, 0x81, 0xa4, 0xa5, 0x64, 0xd2,
	0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12,
	0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0x1a, 0xa4, 0x94, 0x30, 0xcd, 0x03, 0x11, 0xf1, 0xa9, 0x79,
	0x25, 0x30, 0x43, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x4c, 0x7d, 0x10, 0x0b, 0x22, 0xaa,
	0x64, 0xc7, 0x25, 0x1a, 0x08, 0xb2, 0x39, 0x2c, 0xb5, 0xb8, 0xc4, 0x15, 0xa4, 0x3a, 0x28, 0xb5,
	0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x95, 0x8b, 0x0f, 0x64, 0x44, 0x6a, 0x51, 0x7c, 0x62, 0x72,
	0x72, 0x7e, 0x69, 0x5e, 0x89, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x2f, 0x44, 0xd4, 0x11,
	0x22, 0xa8, 0x14, 0xc4, 0x25, 0x86, 0xae, 0xbf, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xc8, 0x82,
	0x8b, 0x15, 0x6c, 0x3d, 0x58, 0x1f, 0xb7, 0x91, 0x8c, 0x1e, 0x86, 0xa7, 0xf4, 0xe0, 0x9a, 0x9c,
	0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x82, 0x68, 0x30, 0x9a, 0xc2, 0xc8, 0xc5, 0x0a, 0x36, 0x54,
	0xa8, 0x8b, 0x91, 0x8b, 0x13, 0xae, 0x48, 0x48, 0x03, 0x8b, 0x11, 0x58, 0x1d, 0x2f, 0xa5, 0x49,
	0x84, 0x4a, 0x88, 0x33, 0x95, 0x34, 0x9a, 0x2e, 0x3f, 0x99, 0xcc, 0xa4, 0x24, 0xa4, 0xa0, 0x8f,
	0x1a, 0x86, 0x26, 0xe8, 0xc1, 0xe8, 0x14, 0x77, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c,
	0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72,
	0x0c, 0x51, 0x2e, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc1, 0x25,
	0x45, 0xa9, 0x89, 0xb9, 0x6e, 0x99, 0x79, 0x89, 0x79, 0xc9, 0xa9, 0xba, 0x01, 0x30, 0xf3, 0x8a,
	0xc1, 0xc2, 0xba, 0xc9, 0x19, 0x89, 0x99, 0x79, 0xfa, 0x70, 0x5b, 0x2a, 0x20, 0x96, 0x94, 0x54,
	0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x85, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x94,
	0xc6, 0xed, 0x1d, 0x02, 0x00, 0x00,
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
	// Queries the VestEntry.
	VestEntry(ctx context.Context, in *QueryVestEntryRequest, opts ...grpc.CallOption) (*QueryVestEntryResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) VestEntry(ctx context.Context, in *QueryVestEntryRequest, opts ...grpc.CallOption) (*QueryVestEntryResponse, error) {
	out := new(QueryVestEntryResponse)
	err := c.cc.Invoke(ctx, "/dydxprotocol.vest.Query/VestEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries the VestEntry.
	VestEntry(context.Context, *QueryVestEntryRequest) (*QueryVestEntryResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) VestEntry(ctx context.Context, req *QueryVestEntryRequest) (*QueryVestEntryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VestEntry not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_VestEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVestEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).VestEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dydxprotocol.vest.Query/VestEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).VestEntry(ctx, req.(*QueryVestEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dydxprotocol.vest.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VestEntry",
			Handler:    _Query_VestEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dydxprotocol/vest/query.proto",
}

func (m *QueryVestEntryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVestEntryRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVestEntryRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VesterAccount) > 0 {
		i -= len(m.VesterAccount)
		copy(dAtA[i:], m.VesterAccount)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.VesterAccount)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryVestEntryResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryVestEntryResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryVestEntryResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Entry.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryVestEntryRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.VesterAccount)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryVestEntryResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Entry.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryVestEntryRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryVestEntryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVestEntryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VesterAccount", wireType)
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
			m.VesterAccount = string(dAtA[iNdEx:postIndex])
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
func (m *QueryVestEntryResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryVestEntryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryVestEntryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entry", wireType)
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
			if err := m.Entry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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