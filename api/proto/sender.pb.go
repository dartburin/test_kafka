// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/sender.proto

package sender

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Result struct {
	Rez                  string   `protobuf:"bytes,1,opt,name=rez,proto3" json:"rez,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_77db3bc0961b3713, []int{0}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetRez() string {
	if m != nil {
		return m.Rez
	}
	return ""
}

type SendMsg struct {
	Msg                  []byte   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMsg) Reset()         { *m = SendMsg{} }
func (m *SendMsg) String() string { return proto.CompactTextString(m) }
func (*SendMsg) ProtoMessage()    {}
func (*SendMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_77db3bc0961b3713, []int{1}
}

func (m *SendMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMsg.Unmarshal(m, b)
}
func (m *SendMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMsg.Marshal(b, m, deterministic)
}
func (m *SendMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMsg.Merge(m, src)
}
func (m *SendMsg) XXX_Size() int {
	return xxx_messageInfo_SendMsg.Size(m)
}
func (m *SendMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SendMsg proto.InternalMessageInfo

func (m *SendMsg) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterType((*Result)(nil), "sender.Result")
	proto.RegisterType((*SendMsg)(nil), "sender.SendMsg")
}

func init() { proto.RegisterFile("api/proto/sender.proto", fileDescriptor_77db3bc0961b3713) }

var fileDescriptor_77db3bc0961b3713 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x2c, 0xc8, 0xd4,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x4e, 0xcd, 0x4b, 0x49, 0x2d, 0xd2, 0x03, 0x73, 0x84,
	0xd8, 0x20, 0x3c, 0x25, 0x29, 0x2e, 0xb6, 0xa0, 0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0x21, 0x01, 0x2e,
	0xe6, 0xa2, 0xd4, 0x2a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x49, 0x9a, 0x8b,
	0x3d, 0x38, 0x35, 0x2f, 0xc5, 0xb7, 0x38, 0x1d, 0x24, 0x99, 0x5b, 0x9c, 0x0e, 0x96, 0xe4, 0x09,
	0x02, 0x31, 0x8d, 0xac, 0xb8, 0xd8, 0x82, 0xc1, 0x46, 0x08, 0x19, 0x70, 0x71, 0x83, 0x95, 0xa5,
	0x16, 0x17, 0x27, 0xa6, 0xa7, 0x0a, 0xf1, 0xeb, 0x41, 0x2d, 0x82, 0xea, 0x95, 0xe2, 0x83, 0x09,
	0x40, 0x2c, 0x52, 0x62, 0x48, 0x62, 0x03, 0xbb, 0xc1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa8,
	0x08, 0x34, 0x81, 0x9d, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SenderClient is the client API for Sender service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SenderClient interface {
	SendMessage(ctx context.Context, in *SendMsg, opts ...grpc.CallOption) (*Result, error)
}

type senderClient struct {
	cc *grpc.ClientConn
}

func NewSenderClient(cc *grpc.ClientConn) SenderClient {
	return &senderClient{cc}
}

func (c *senderClient) SendMessage(ctx context.Context, in *SendMsg, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/sender.Sender/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SenderServer is the server API for Sender service.
type SenderServer interface {
	SendMessage(context.Context, *SendMsg) (*Result, error)
}

// UnimplementedSenderServer can be embedded to have forward compatible implementations.
type UnimplementedSenderServer struct {
}

func (*UnimplementedSenderServer) SendMessage(ctx context.Context, req *SendMsg) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

func RegisterSenderServer(s *grpc.Server, srv SenderServer) {
	s.RegisterService(&_Sender_serviceDesc, srv)
}

func _Sender_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SenderServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sender.Sender/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SenderServer).SendMessage(ctx, req.(*SendMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sender_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sender.Sender",
	HandlerType: (*SenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _Sender_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/sender.proto",
}