// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package protogen

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// HelloMessage is a protobuf message for automatic tests.
type HelloMessage struct {
	HelloString          string   `protobuf:"bytes,1,opt,name=hello_string,json=helloString,proto3" json:"hello_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloMessage) Reset()         { *m = HelloMessage{} }
func (m *HelloMessage) String() string { return proto.CompactTextString(m) }
func (*HelloMessage) ProtoMessage()    {}
func (*HelloMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *HelloMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloMessage.Unmarshal(m, b)
}
func (m *HelloMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloMessage.Marshal(b, m, deterministic)
}
func (m *HelloMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloMessage.Merge(m, src)
}
func (m *HelloMessage) XXX_Size() int {
	return xxx_messageInfo_HelloMessage.Size(m)
}
func (m *HelloMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloMessage.DiscardUnknown(m)
}

var xxx_messageInfo_HelloMessage proto.InternalMessageInfo

func (m *HelloMessage) GetHelloString() string {
	if m != nil {
		return m.HelloString
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloMessage)(nil), "fd8_judge_grpc_test_api.HelloMessage")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4f, 0x4b, 0xb1, 0x88, 0xcf, 0x2a, 0x4d, 0x49,
	0x4f, 0x8d, 0x4f, 0x2f, 0x2a, 0x48, 0x8e, 0x07, 0x49, 0xc5, 0x27, 0x16, 0x64, 0x4a, 0xc9, 0xa4,
	0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24,
	0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0xb4, 0x29, 0x19, 0x72, 0xf1, 0x78, 0xa4, 0xe6, 0xe4, 0xe4,
	0xfb, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x0a, 0x29, 0x72, 0xf1, 0x64, 0x80, 0xf8, 0xf1, 0xc5,
	0x25, 0x45, 0x99, 0x79, 0xe9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xdc, 0x60, 0xb1, 0x60,
	0xb0, 0x90, 0x51, 0x39, 0x17, 0x77, 0x48, 0x6a, 0x71, 0x49, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72,
	0xaa, 0x50, 0x06, 0x17, 0x47, 0x70, 0x62, 0x25, 0xd8, 0x10, 0x21, 0x55, 0x3d, 0x1c, 0xae, 0xd0,
	0x43, 0xb6, 0x44, 0x8a, 0x38, 0x65, 0x4a, 0x7c, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0xe2, 0x10, 0x62,
	0xd3, 0x07, 0x5b, 0xef, 0xc4, 0x15, 0xc5, 0x01, 0x76, 0x74, 0x7a, 0x6a, 0x5e, 0x12, 0x1b, 0x98,
	0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x4e, 0x0b, 0xba, 0x03, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestServiceClient interface {
	// SayHello receives a HelloMessage and replies another HelloMessage with its string doubled.
	SayHello(ctx context.Context, in *HelloMessage, opts ...grpc.CallOption) (*HelloMessage, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) SayHello(ctx context.Context, in *HelloMessage, opts ...grpc.CallOption) (*HelloMessage, error) {
	out := new(HelloMessage)
	err := c.cc.Invoke(ctx, "/fd8_judge_grpc_test_api.TestService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
type TestServiceServer interface {
	// SayHello receives a HelloMessage and replies another HelloMessage with its string doubled.
	SayHello(context.Context, *HelloMessage) (*HelloMessage, error)
}

// UnimplementedTestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (*UnimplementedTestServiceServer) SayHello(ctx context.Context, req *HelloMessage) (*HelloMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fd8_judge_grpc_test_api.TestService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).SayHello(ctx, req.(*HelloMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fd8_judge_grpc_test_api.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _TestService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
