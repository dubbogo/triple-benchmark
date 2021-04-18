// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package protobuf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

import (
	dubboConstant "github.com/apache/dubbo-go/common/constant"
	"github.com/apache/dubbo-go/protocol"
	dgrpc "github.com/apache/dubbo-go/protocol/grpc"
	"github.com/apache/dubbo-go/protocol/invocation"
	dubbo3 "github.com/dubbogo/triple/pkg/triple"
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

// The request message containing the user's name.
type Dubbo3HelloRequest struct {
	Myname               string   `protobuf:"bytes,1,opt,name=myname,proto3" json:"myname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dubbo3HelloRequest) Reset()         { *m = Dubbo3HelloRequest{} }
func (m *Dubbo3HelloRequest) String() string { return proto.CompactTextString(m) }
func (*Dubbo3HelloRequest) ProtoMessage()    {}
func (*Dubbo3HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *Dubbo3HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dubbo3HelloRequest.Unmarshal(m, b)
}
func (m *Dubbo3HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dubbo3HelloRequest.Marshal(b, m, deterministic)
}
func (m *Dubbo3HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dubbo3HelloRequest.Merge(m, src)
}
func (m *Dubbo3HelloRequest) XXX_Size() int {
	return xxx_messageInfo_Dubbo3HelloRequest.Size(m)
}
func (m *Dubbo3HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Dubbo3HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_Dubbo3HelloRequest proto.InternalMessageInfo

func (m *Dubbo3HelloRequest) GetMyname() string {
	if m != nil {
		return m.Myname
	}
	return ""
}

// The response message containing the greetings
type Dubbo3HelloReply struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dubbo3HelloReply) Reset()         { *m = Dubbo3HelloReply{} }
func (m *Dubbo3HelloReply) String() string { return proto.CompactTextString(m) }
func (*Dubbo3HelloReply) ProtoMessage()    {}
func (*Dubbo3HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *Dubbo3HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dubbo3HelloReply.Unmarshal(m, b)
}
func (m *Dubbo3HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dubbo3HelloReply.Marshal(b, m, deterministic)
}
func (m *Dubbo3HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dubbo3HelloReply.Merge(m, src)
}
func (m *Dubbo3HelloReply) XXX_Size() int {
	return xxx_messageInfo_Dubbo3HelloReply.Size(m)
}
func (m *Dubbo3HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_Dubbo3HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_Dubbo3HelloReply proto.InternalMessageInfo

func (m *Dubbo3HelloReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type BigData struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	WantSize             int32    `protobuf:"varint,2,opt,name=wantSize,proto3" json:"wantSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BigData) Reset()         { *m = BigData{} }
func (m *BigData) String() string { return proto.CompactTextString(m) }
func (*BigData) ProtoMessage()    {}
func (*BigData) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{2}
}

func (m *BigData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BigData.Unmarshal(m, b)
}
func (m *BigData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BigData.Marshal(b, m, deterministic)
}
func (m *BigData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BigData.Merge(m, src)
}
func (m *BigData) XXX_Size() int {
	return xxx_messageInfo_BigData.Size(m)
}
func (m *BigData) XXX_DiscardUnknown() {
	xxx_messageInfo_BigData.DiscardUnknown(m)
}

var xxx_messageInfo_BigData proto.InternalMessageInfo

func (m *BigData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BigData) GetWantSize() int32 {
	if m != nil {
		return m.WantSize
	}
	return 0
}

func init() {
	proto.RegisterType((*Dubbo3HelloRequest)(nil), "protobuf.Dubbo3HelloRequest")
	proto.RegisterType((*Dubbo3HelloReply)(nil), "protobuf.Dubbo3HelloReply")
	proto.RegisterType((*BigData)(nil), "protobuf.BigData")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xdd, 0x56, 0x6b, 0x1d, 0x5a, 0x1b, 0xe7, 0x20, 0x25, 0x78, 0x28, 0xc1, 0x43, 0x41,
	0x09, 0xd2, 0x82, 0x20, 0xde, 0x42, 0xc1, 0x1e, 0x14, 0x42, 0xa2, 0xf4, 0x3c, 0x31, 0x6b, 0x1a,
	0x48, 0xb2, 0x71, 0xbb, 0xa1, 0xc4, 0xc7, 0x11, 0x7c, 0x4f, 0xd9, 0x6d, 0x54, 0x4a, 0xf1, 0x20,
	0x9e, 0x76, 0xfe, 0x7c, 0xf3, 0xcd, 0x8f, 0x59, 0xb0, 0x96, 0x3c, 0xcb, 0xc4, 0x5a, 0xc8, 0x2c,
	0x76, 0x4b, 0x29, 0x94, 0xc0, 0xae, 0x79, 0xa2, 0xea, 0xc5, 0xb9, 0x04, 0x9c, 0x55, 0x51, 0x24,
	0xa6, 0x73, 0xad, 0x09, 0xf8, 0x6b, 0xc5, 0x57, 0x0a, 0x4f, 0xa1, 0x93, 0xd7, 0x05, 0xe5, 0x7c,
	0xc8, 0x46, 0x6c, 0x7c, 0x14, 0x34, 0x99, 0x73, 0x0e, 0xd6, 0x96, 0xba, 0xcc, 0x6a, 0xb4, 0xa0,
	0x9d, 0xaf, 0x92, 0x46, 0xa8, 0x43, 0xe7, 0x06, 0x0e, 0xbd, 0x34, 0x99, 0x91, 0x22, 0x44, 0xd8,
	0x8f, 0x49, 0x91, 0xe9, 0xf6, 0x02, 0x13, 0xa3, 0x0d, 0xdd, 0x35, 0x15, 0x2a, 0x4c, 0xdf, 0xf8,
	0xb0, 0x35, 0x62, 0xe3, 0x83, 0xe0, 0x3b, 0x9f, 0x7c, 0xb4, 0xa0, 0xbf, 0xd9, 0x70, 0x27, 0x39,
	0x57, 0x5c, 0xa2, 0x0f, 0xc7, 0x9b, 0x42, 0x48, 0xb5, 0xd9, 0x8a, 0x67, 0xee, 0x17, 0xbd, 0xbb,
	0x8b, 0x6e, 0xdb, 0xbf, 0x74, 0xcb, 0xac, 0x76, 0xf6, 0xc6, 0xec, 0x8a, 0xe1, 0x03, 0x0c, 0xb6,
	0x1d, 0x27, 0xff, 0xb1, 0xc4, 0x5b, 0xe8, 0x7b, 0x69, 0x12, 0x2a, 0xc9, 0x29, 0x7f, 0xd4, 0xc7,
	0x3b, 0xf9, 0x91, 0x37, 0x67, 0xb0, 0x77, 0x4b, 0x0d, 0xcb, 0x35, 0xf4, 0xbc, 0x34, 0x79, 0x2a,
	0x48, 0xd6, 0x7f, 0x99, 0xf5, 0x2e, 0xc0, 0x12, 0x32, 0x71, 0xa9, 0xa4, 0xe7, 0x25, 0x77, 0x63,
	0x4d, 0xe5, 0x0d, 0x0c, 0xd6, 0x42, 0x7f, 0xb3, 0xaf, 0x27, 0x7c, 0xf6, 0xde, 0x6a, 0xcf, 0xef,
	0x17, 0x51, 0xc7, 0x18, 0x4c, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x68, 0x42, 0xd4, 0x23, 0x08,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// Dubbo3GreeterClient is the client API for Dubbo3Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Dubbo3GreeterClient interface {
	// Sends a greeting
	Dubbo3SayHello(ctx context.Context, opts ...grpc.CallOption) (Dubbo3Greeter_Dubbo3SayHelloClient, error)
	Dubbo3SayHello2(ctx context.Context, in *Dubbo3HelloRequest, opts ...grpc.CallOption) (*Dubbo3HelloReply, error)
	BigStreamTest(ctx context.Context, opts ...grpc.CallOption) (Dubbo3Greeter_BigStreamTestClient, error)
	BigUnaryTest(ctx context.Context, in *BigData, opts ...grpc.CallOption) (*BigData, error)
}

type dubbo3GreeterClient struct {
	cc grpc.ClientConnInterface
}

func NewDubbo3GreeterClient(cc grpc.ClientConnInterface) Dubbo3GreeterClient {
	return &dubbo3GreeterClient{cc}
}

func (c *dubbo3GreeterClient) Dubbo3SayHello(ctx context.Context, opts ...grpc.CallOption) (Dubbo3Greeter_Dubbo3SayHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Dubbo3Greeter_serviceDesc.Streams[0], "/protobuf.Dubbo3Greeter/Dubbo3SayHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &dubbo3GreeterDubbo3SayHelloClient{stream}
	return x, nil
}

type Dubbo3Greeter_Dubbo3SayHelloClient interface {
	Send(*Dubbo3HelloRequest) error
	Recv() (*Dubbo3HelloReply, error)
	grpc.ClientStream
}

type dubbo3GreeterDubbo3SayHelloClient struct {
	grpc.ClientStream
}

func (x *dubbo3GreeterDubbo3SayHelloClient) Send(m *Dubbo3HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dubbo3GreeterDubbo3SayHelloClient) Recv() (*Dubbo3HelloReply, error) {
	m := new(Dubbo3HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dubbo3GreeterClient) Dubbo3SayHello2(ctx context.Context, in *Dubbo3HelloRequest, opts ...grpc.CallOption) (*Dubbo3HelloReply, error) {
	out := new(Dubbo3HelloReply)
	err := c.cc.Invoke(ctx, "/protobuf.Dubbo3Greeter/Dubbo3SayHello2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dubbo3GreeterClient) BigStreamTest(ctx context.Context, opts ...grpc.CallOption) (Dubbo3Greeter_BigStreamTestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Dubbo3Greeter_serviceDesc.Streams[1], "/protobuf.Dubbo3Greeter/BigStreamTest", opts...)
	if err != nil {
		return nil, err
	}
	x := &dubbo3GreeterBigStreamTestClient{stream}
	return x, nil
}

type Dubbo3Greeter_BigStreamTestClient interface {
	Send(*BigData) error
	Recv() (*BigData, error)
	grpc.ClientStream
}

type dubbo3GreeterBigStreamTestClient struct {
	grpc.ClientStream
}

func (x *dubbo3GreeterBigStreamTestClient) Send(m *BigData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dubbo3GreeterBigStreamTestClient) Recv() (*BigData, error) {
	m := new(BigData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dubbo3GreeterClient) BigUnaryTest(ctx context.Context, in *BigData, opts ...grpc.CallOption) (*BigData, error) {
	out := new(BigData)
	err := c.cc.Invoke(ctx, "/protobuf.Dubbo3Greeter/BigUnaryTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Dubbo3GreeterServer is the server API for Dubbo3Greeter service.
type Dubbo3GreeterServer interface {
	// Sends a greeting
	Dubbo3SayHello(Dubbo3Greeter_Dubbo3SayHelloServer) error
	Dubbo3SayHello2(context.Context, *Dubbo3HelloRequest) (*Dubbo3HelloReply, error)
	BigStreamTest(Dubbo3Greeter_BigStreamTestServer) error
	BigUnaryTest(context.Context, *BigData) (*BigData, error)
}

// UnimplementedDubbo3GreeterServer can be embedded to have forward compatible implementations.
type UnimplementedDubbo3GreeterServer struct {
}

func (*UnimplementedDubbo3GreeterServer) Dubbo3SayHello(srv Dubbo3Greeter_Dubbo3SayHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method Dubbo3SayHello not implemented")
}
func (*UnimplementedDubbo3GreeterServer) Dubbo3SayHello2(ctx context.Context, req *Dubbo3HelloRequest) (*Dubbo3HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dubbo3SayHello2 not implemented")
}
func (*UnimplementedDubbo3GreeterServer) BigStreamTest(srv Dubbo3Greeter_BigStreamTestServer) error {
	return status.Errorf(codes.Unimplemented, "method BigStreamTest not implemented")
}
func (*UnimplementedDubbo3GreeterServer) BigUnaryTest(ctx context.Context, req *BigData) (*BigData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BigUnaryTest not implemented")
}

func RegisterDubbo3GreeterServer(s *grpc.Server, srv Dubbo3GreeterServer) {
	s.RegisterService(&_Dubbo3Greeter_serviceDesc, srv)
}

func _Dubbo3Greeter_Dubbo3SayHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(Dubbo3GreeterServer).Dubbo3SayHello(&dubbo3GreeterDubbo3SayHelloServer{stream})
}

type Dubbo3Greeter_Dubbo3SayHelloServer interface {
	Send(*Dubbo3HelloReply) error
	Recv() (*Dubbo3HelloRequest, error)
	grpc.ServerStream
}

type dubbo3GreeterDubbo3SayHelloServer struct {
	grpc.ServerStream
}

func (x *dubbo3GreeterDubbo3SayHelloServer) Send(m *Dubbo3HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dubbo3GreeterDubbo3SayHelloServer) Recv() (*Dubbo3HelloRequest, error) {
	m := new(Dubbo3HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Dubbo3Greeter_Dubbo3SayHello2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Dubbo3HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Dubbo3GreeterServer).Dubbo3SayHello2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Dubbo3Greeter/Dubbo3SayHello2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Dubbo3GreeterServer).Dubbo3SayHello2(ctx, req.(*Dubbo3HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dubbo3Greeter_BigStreamTest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(Dubbo3GreeterServer).BigStreamTest(&dubbo3GreeterBigStreamTestServer{stream})
}

type Dubbo3Greeter_BigStreamTestServer interface {
	Send(*BigData) error
	Recv() (*BigData, error)
	grpc.ServerStream
}

type dubbo3GreeterBigStreamTestServer struct {
	grpc.ServerStream
}

func (x *dubbo3GreeterBigStreamTestServer) Send(m *BigData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dubbo3GreeterBigStreamTestServer) Recv() (*BigData, error) {
	m := new(BigData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Dubbo3Greeter_BigUnaryTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BigData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Dubbo3GreeterServer).BigUnaryTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Dubbo3Greeter/BigUnaryTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Dubbo3GreeterServer).BigUnaryTest(ctx, req.(*BigData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dubbo3Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Dubbo3Greeter",
	HandlerType: (*Dubbo3GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dubbo3SayHello2",
			Handler:    _Dubbo3Greeter_Dubbo3SayHello2_Handler,
		},
		{
			MethodName: "BigUnaryTest",
			Handler:    _Dubbo3Greeter_BigUnaryTest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Dubbo3SayHello",
			Handler:       _Dubbo3Greeter_Dubbo3SayHello_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "BigStreamTest",
			Handler:       _Dubbo3Greeter_BigStreamTest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}

type dubbo3greeterDubbo3Client struct {
	cc *dubbo3.TripleConn
}

func NewDubbo3GreeterDubbo3Client(cc *dubbo3.TripleConn) Dubbo3GreeterClient {
	return &dubbo3greeterDubbo3Client{cc}
}
func (c *dubbo3greeterDubbo3Client) Dubbo3SayHello(ctx context.Context, opt ...grpc.CallOption) (Dubbo3Greeter_Dubbo3SayHelloClient, error) {
	interfaceKey := ctx.Value(dubboConstant.DubboCtxKey(dubboConstant.INTERFACE_KEY)).(string)
	stream, err := c.cc.NewStream(ctx, "/"+interfaceKey+"/Dubbo3SayHello", opt...)
	if err != nil {
		return nil, err
	}
	x := &dubbo3GreeterDubbo3SayHelloClient{stream}
	return x, nil
}
func (c *dubbo3greeterDubbo3Client) Dubbo3SayHello2(ctx context.Context, in *Dubbo3HelloRequest, opt ...grpc.CallOption) (*Dubbo3HelloReply, error) {
	out := new(Dubbo3HelloReply)
	interfaceKey := ctx.Value(dubboConstant.DubboCtxKey(dubboConstant.INTERFACE_KEY)).(string)
	err := c.cc.Invoke(ctx, "/"+interfaceKey+"/Dubbo3SayHello2", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (c *dubbo3greeterDubbo3Client) BigStreamTest(ctx context.Context, opt ...grpc.CallOption) (Dubbo3Greeter_BigStreamTestClient, error) {
	interfaceKey := ctx.Value(dubboConstant.DubboCtxKey(dubboConstant.INTERFACE_KEY)).(string)
	stream, err := c.cc.NewStream(ctx, "/"+interfaceKey+"/BigStreamTest", opt...)
	if err != nil {
		return nil, err
	}
	x := &dubbo3GreeterBigStreamTestClient{stream}
	return x, nil
}
func (c *dubbo3greeterDubbo3Client) BigUnaryTest(ctx context.Context, in *BigData, opt ...grpc.CallOption) (*BigData, error) {
	out := new(BigData)
	interfaceKey := ctx.Value(dubboConstant.DubboCtxKey(dubboConstant.INTERFACE_KEY)).(string)
	err := c.cc.Invoke(ctx, "/"+interfaceKey+"/BigUnaryTest", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Dubbo3GreeterClientImpl is the client API for Dubbo3Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Dubbo3GreeterClientImpl struct {
	// Sends a greeting
	Dubbo3SayHello  func(ctx context.Context) (Dubbo3Greeter_Dubbo3SayHelloClient, error)
	Dubbo3SayHello2 func(ctx context.Context, in *Dubbo3HelloRequest, out *Dubbo3HelloReply) error
	BigStreamTest   func(ctx context.Context) (Dubbo3Greeter_BigStreamTestClient, error)
	BigUnaryTest    func(ctx context.Context, in *BigData, out *BigData) error
}

func (c *Dubbo3GreeterClientImpl) Reference() string {
	return "dubbo3GreeterImpl"
}

func (c *Dubbo3GreeterClientImpl) GetDubboStub(cc *dubbo3.TripleConn) Dubbo3GreeterClient {
	return NewDubbo3GreeterDubbo3Client(cc)
}

type Dubbo3GreeterProviderBase struct {
	proxyImpl protocol.Invoker
}

func (s *Dubbo3GreeterProviderBase) SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *Dubbo3GreeterProviderBase) GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func _DUBBO_Dubbo3Greeter_Dubbo3SayHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	_, ok := srv.(dgrpc.DubboGrpcService)
	invo := invocation.NewRPCInvocation("Dubbo3SayHello", nil, nil)
	if !ok {
		fmt.Println(invo)
	}
	return srv.(Dubbo3GreeterServer).Dubbo3SayHello(&dubbo3GreeterDubbo3SayHelloServer{stream})
}

func _DUBBO_Dubbo3Greeter_Dubbo3SayHello2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Dubbo3HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dgrpc.DubboGrpcService)
	args := []interface{}{}
	args = append(args, in)
	invo := invocation.NewRPCInvocation("Dubbo3SayHello2", args, nil)
	if interceptor == nil {
		result := base.GetProxyImpl().Invoke(ctx, invo)
		return result.Result(), result.Error()
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Dubbo3Greeter/Dubbo3SayHello2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.GetProxyImpl().Invoke(context.Background(), invo)
		return result.Result(), result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _DUBBO_Dubbo3Greeter_BigStreamTest_Handler(srv interface{}, stream grpc.ServerStream) error {
	_, ok := srv.(dgrpc.DubboGrpcService)
	invo := invocation.NewRPCInvocation("BigStreamTest", nil, nil)
	if !ok {
		fmt.Println(invo)
	}
	return srv.(Dubbo3GreeterServer).BigStreamTest(&dubbo3GreeterBigStreamTestServer{stream})
}

func _DUBBO_Dubbo3Greeter_BigUnaryTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BigData)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dgrpc.DubboGrpcService)
	args := []interface{}{}
	args = append(args, in)
	invo := invocation.NewRPCInvocation("BigUnaryTest", args, nil)
	if interceptor == nil {
		result := base.GetProxyImpl().Invoke(ctx, invo)
		return result.Result(), result.Error()
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Dubbo3Greeter/BigUnaryTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.GetProxyImpl().Invoke(context.Background(), invo)
		return result.Result(), result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func (s *Dubbo3GreeterProviderBase) ServiceDesc() *grpc.ServiceDesc {
	return &grpc.ServiceDesc{
		ServiceName: "protobuf.Dubbo3Greeter",
		HandlerType: (*Dubbo3GreeterServer)(nil),
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Dubbo3SayHello2",
				Handler:    _DUBBO_Dubbo3Greeter_Dubbo3SayHello2_Handler,
			},
			{
				MethodName: "BigUnaryTest",
				Handler:    _DUBBO_Dubbo3Greeter_BigUnaryTest_Handler,
			},
		},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "Dubbo3SayHello",
				Handler:       _DUBBO_Dubbo3Greeter_Dubbo3SayHello_Handler,
				ServerStreams: true,
				ClientStreams: true,
			},
			{
				StreamName:    "BigStreamTest",
				Handler:       _DUBBO_Dubbo3Greeter_BigStreamTest_Handler,
				ServerStreams: true,
				ClientStreams: true,
			},
		},
		Metadata: "helloworld.proto",
	}
}
