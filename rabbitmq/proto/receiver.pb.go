// Code generated by protoc-gen-go. DO NOT EDIT.
// source: receiver.proto

package proto

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

type MessageId struct {
	MessageId            int64    `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageId) Reset()         { *m = MessageId{} }
func (m *MessageId) String() string { return proto.CompactTextString(m) }
func (*MessageId) ProtoMessage()    {}
func (*MessageId) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7296e1d2b388c5, []int{0}
}

func (m *MessageId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageId.Unmarshal(m, b)
}
func (m *MessageId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageId.Marshal(b, m, deterministic)
}
func (m *MessageId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageId.Merge(m, src)
}
func (m *MessageId) XXX_Size() int {
	return xxx_messageInfo_MessageId.Size(m)
}
func (m *MessageId) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageId.DiscardUnknown(m)
}

var xxx_messageInfo_MessageId proto.InternalMessageInfo

func (m *MessageId) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type Reject struct {
	Id                   *MessageId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Requeue              bool       `protobuf:"varint,2,opt,name=requeue,proto3" json:"requeue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Reject) Reset()         { *m = Reject{} }
func (m *Reject) String() string { return proto.CompactTextString(m) }
func (*Reject) ProtoMessage()    {}
func (*Reject) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7296e1d2b388c5, []int{1}
}

func (m *Reject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reject.Unmarshal(m, b)
}
func (m *Reject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reject.Marshal(b, m, deterministic)
}
func (m *Reject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reject.Merge(m, src)
}
func (m *Reject) XXX_Size() int {
	return xxx_messageInfo_Reject.Size(m)
}
func (m *Reject) XXX_DiscardUnknown() {
	xxx_messageInfo_Reject.DiscardUnknown(m)
}

var xxx_messageInfo_Reject proto.InternalMessageInfo

func (m *Reject) GetId() *MessageId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Reject) GetRequeue() bool {
	if m != nil {
		return m.Requeue
	}
	return false
}

type ReceiverResponse struct {
	Filters              map[string]string `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Destination          string            `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Content              []byte            `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	IsARedelivery        bool              `protobuf:"varint,4,opt,name=is_a_redelivery,json=isARedelivery,proto3" json:"is_a_redelivery,omitempty"`
	MessageId            int64             `protobuf:"varint,5,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReceiverResponse) Reset()         { *m = ReceiverResponse{} }
func (m *ReceiverResponse) String() string { return proto.CompactTextString(m) }
func (*ReceiverResponse) ProtoMessage()    {}
func (*ReceiverResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7296e1d2b388c5, []int{2}
}

func (m *ReceiverResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiverResponse.Unmarshal(m, b)
}
func (m *ReceiverResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiverResponse.Marshal(b, m, deterministic)
}
func (m *ReceiverResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiverResponse.Merge(m, src)
}
func (m *ReceiverResponse) XXX_Size() int {
	return xxx_messageInfo_ReceiverResponse.Size(m)
}
func (m *ReceiverResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiverResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiverResponse proto.InternalMessageInfo

func (m *ReceiverResponse) GetFilters() map[string]string {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *ReceiverResponse) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *ReceiverResponse) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ReceiverResponse) GetIsARedelivery() bool {
	if m != nil {
		return m.IsARedelivery
	}
	return false
}

func (m *ReceiverResponse) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

type ReceiverArgs struct {
	QueueName            string   `protobuf:"bytes,1,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiverArgs) Reset()         { *m = ReceiverArgs{} }
func (m *ReceiverArgs) String() string { return proto.CompactTextString(m) }
func (*ReceiverArgs) ProtoMessage()    {}
func (*ReceiverArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7296e1d2b388c5, []int{3}
}

func (m *ReceiverArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiverArgs.Unmarshal(m, b)
}
func (m *ReceiverArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiverArgs.Marshal(b, m, deterministic)
}
func (m *ReceiverArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiverArgs.Merge(m, src)
}
func (m *ReceiverArgs) XXX_Size() int {
	return xxx_messageInfo_ReceiverArgs.Size(m)
}
func (m *ReceiverArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiverArgs.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiverArgs proto.InternalMessageInfo

func (m *ReceiverArgs) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

type ActionStatus struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActionStatus) Reset()         { *m = ActionStatus{} }
func (m *ActionStatus) String() string { return proto.CompactTextString(m) }
func (*ActionStatus) ProtoMessage()    {}
func (*ActionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b7296e1d2b388c5, []int{4}
}

func (m *ActionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionStatus.Unmarshal(m, b)
}
func (m *ActionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionStatus.Marshal(b, m, deterministic)
}
func (m *ActionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionStatus.Merge(m, src)
}
func (m *ActionStatus) XXX_Size() int {
	return xxx_messageInfo_ActionStatus.Size(m)
}
func (m *ActionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ActionStatus proto.InternalMessageInfo

func (m *ActionStatus) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*MessageId)(nil), "receiver.MessageId")
	proto.RegisterType((*Reject)(nil), "receiver.Reject")
	proto.RegisterType((*ReceiverResponse)(nil), "receiver.ReceiverResponse")
	proto.RegisterMapType((map[string]string)(nil), "receiver.ReceiverResponse.FiltersEntry")
	proto.RegisterType((*ReceiverArgs)(nil), "receiver.ReceiverArgs")
	proto.RegisterType((*ActionStatus)(nil), "receiver.ActionStatus")
}

func init() {
	proto.RegisterFile("receiver.proto", fileDescriptor_4b7296e1d2b388c5)
}

var fileDescriptor_4b7296e1d2b388c5 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x52, 0xdb, 0x6a, 0xdb, 0x40,
	0x10, 0xb5, 0xa4, 0xda, 0x96, 0xc6, 0x72, 0x6b, 0xb6, 0xc5, 0x08, 0x41, 0xc1, 0xa8, 0xe0, 0x96,
	0x42, 0x4d, 0x71, 0x5f, 0x8a, 0xfb, 0x24, 0x4a, 0x1b, 0xf2, 0x90, 0x3c, 0x6c, 0x3e, 0x40, 0x28,
	0xd2, 0xc4, 0x28, 0x96, 0x57, 0xce, 0xee, 0xda, 0xc1, 0xff, 0x91, 0xff, 0xca, 0x2f, 0x45, 0x5a,
	0x69, 0x6d, 0xe1, 0x5c, 0xde, 0xf6, 0xcc, 0xe5, 0xcc, 0xcc, 0x39, 0x0b, 0xef, 0x39, 0x26, 0x98,
	0xed, 0x90, 0xcf, 0x36, 0xbc, 0x90, 0x05, 0xb1, 0x35, 0x0e, 0xbe, 0x83, 0x73, 0x81, 0x42, 0xc4,
	0x4b, 0x3c, 0x4f, 0xc9, 0x67, 0x80, 0x75, 0x0d, 0xa2, 0x2c, 0xf5, 0x8c, 0x89, 0xf1, 0xcd, 0xa2,
	0xce, 0x5a, 0xa7, 0x83, 0x33, 0xe8, 0x51, 0xbc, 0xc5, 0x44, 0x92, 0x2f, 0x60, 0x36, 0x05, 0x83,
	0xf9, 0xc7, 0xd9, 0x81, 0xfc, 0xc0, 0x44, 0xcb, 0x34, 0xf1, 0xa0, 0xcf, 0xf1, 0x6e, 0x8b, 0x5b,
	0xf4, 0xcc, 0xb2, 0xd2, 0xa6, 0x1a, 0x06, 0x0f, 0x26, 0x8c, 0x68, 0xd3, 0x44, 0x51, 0x6c, 0x0a,
	0x26, 0x90, 0x84, 0xd0, 0xbf, 0xc9, 0x72, 0x89, 0x5c, 0x94, 0xc4, 0x56, 0x49, 0xfc, 0xf5, 0x48,
	0x7c, 0x5a, 0x3c, 0xfb, 0x5f, 0x57, 0xfe, 0x63, 0x92, 0xef, 0xa9, 0xee, 0x23, 0x13, 0x18, 0xa4,
	0x28, 0x64, 0xc6, 0x62, 0x99, 0x15, 0x4c, 0x4d, 0x75, 0x68, 0x3b, 0x54, 0xed, 0x94, 0x14, 0x4c,
	0x22, 0x93, 0x9e, 0x55, 0x66, 0x5d, 0xaa, 0x21, 0x99, 0xc2, 0x87, 0x4c, 0x44, 0x71, 0xc4, 0x31,
	0xc5, 0xbc, 0x9a, 0xb5, 0xf7, 0xde, 0xa9, 0xad, 0x87, 0x99, 0x08, 0xe9, 0x21, 0x78, 0xa2, 0x51,
	0xf7, 0x44, 0x23, 0x7f, 0x01, 0x6e, 0x7b, 0x37, 0x32, 0x02, 0x6b, 0x85, 0x7b, 0x25, 0x95, 0x43,
	0xab, 0x27, 0xf9, 0x04, 0xdd, 0x5d, 0x9c, 0x37, 0xa2, 0x38, 0xb4, 0x06, 0x0b, 0xf3, 0xb7, 0x11,
	0xfc, 0x00, 0x57, 0x1f, 0x1a, 0xf2, 0xa5, 0xa8, 0x46, 0x29, 0xbd, 0x22, 0x16, 0xaf, 0xb1, 0xa1,
	0x70, 0x54, 0xe4, 0xb2, 0x0c, 0x04, 0x53, 0x70, 0xc3, 0xa4, 0xba, 0xea, 0x4a, 0xc6, 0x72, 0x2b,
	0xc8, 0x18, 0x7a, 0x42, 0xbd, 0x54, 0xa9, 0x4d, 0x1b, 0x34, 0x7f, 0x34, 0xc0, 0xd6, 0xbc, 0xe4,
	0x2f, 0x90, 0x30, 0x59, 0xb1, 0xe2, 0x3e, 0xc7, 0x74, 0x89, 0x8d, 0x61, 0xe4, 0x25, 0x0f, 0xfd,
	0xf1, 0x31, 0xd8, 0x9e, 0x13, 0x74, 0xc8, 0x1f, 0x18, 0xd6, 0x1f, 0x41, 0xf7, 0x8f, 0xda, 0x56,
	0x55, 0x89, 0x37, 0x9a, 0x4b, 0x9f, 0x9b, 0x6d, 0xc8, 0xf8, 0xb9, 0xc3, 0xd5, 0xe1, 0xbe, 0xff,
	0xba, 0xf3, 0x41, 0xe7, 0xa7, 0x71, 0xdd, 0x53, 0xbf, 0xf8, 0xd7, 0x53, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xcc, 0x43, 0xbe, 0xab, 0xd7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReceiverClient is the client API for Receiver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReceiverClient interface {
	AcknowledgeMessage(ctx context.Context, in *MessageId, opts ...grpc.CallOption) (*ActionStatus, error)
	RejectMessage(ctx context.Context, in *Reject, opts ...grpc.CallOption) (*ActionStatus, error)
	Receive(ctx context.Context, in *ReceiverArgs, opts ...grpc.CallOption) (Receiver_ReceiveClient, error)
}

type receiverClient struct {
	cc grpc.ClientConnInterface
}

func NewReceiverClient(cc grpc.ClientConnInterface) ReceiverClient {
	return &receiverClient{cc}
}

func (c *receiverClient) AcknowledgeMessage(ctx context.Context, in *MessageId, opts ...grpc.CallOption) (*ActionStatus, error) {
	out := new(ActionStatus)
	err := c.cc.Invoke(ctx, "/receiver.Receiver/AcknowledgeMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *receiverClient) RejectMessage(ctx context.Context, in *Reject, opts ...grpc.CallOption) (*ActionStatus, error) {
	out := new(ActionStatus)
	err := c.cc.Invoke(ctx, "/receiver.Receiver/RejectMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *receiverClient) Receive(ctx context.Context, in *ReceiverArgs, opts ...grpc.CallOption) (Receiver_ReceiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Receiver_serviceDesc.Streams[0], "/receiver.Receiver/Receive", opts...)
	if err != nil {
		return nil, err
	}
	x := &receiverReceiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Receiver_ReceiveClient interface {
	Recv() (*ReceiverResponse, error)
	grpc.ClientStream
}

type receiverReceiveClient struct {
	grpc.ClientStream
}

func (x *receiverReceiveClient) Recv() (*ReceiverResponse, error) {
	m := new(ReceiverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReceiverServer is the server API for Receiver service.
type ReceiverServer interface {
	AcknowledgeMessage(context.Context, *MessageId) (*ActionStatus, error)
	RejectMessage(context.Context, *Reject) (*ActionStatus, error)
	Receive(*ReceiverArgs, Receiver_ReceiveServer) error
}

// UnimplementedReceiverServer can be embedded to have forward compatible implementations.
type UnimplementedReceiverServer struct {
}

func (*UnimplementedReceiverServer) AcknowledgeMessage(ctx context.Context, req *MessageId) (*ActionStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcknowledgeMessage not implemented")
}
func (*UnimplementedReceiverServer) RejectMessage(ctx context.Context, req *Reject) (*ActionStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectMessage not implemented")
}
func (*UnimplementedReceiverServer) Receive(req *ReceiverArgs, srv Receiver_ReceiveServer) error {
	return status.Errorf(codes.Unimplemented, "method Receive not implemented")
}

func RegisterReceiverServer(s *grpc.Server, srv ReceiverServer) {
	s.RegisterService(&_Receiver_serviceDesc, srv)
}

func _Receiver_AcknowledgeMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiverServer).AcknowledgeMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/receiver.Receiver/AcknowledgeMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiverServer).AcknowledgeMessage(ctx, req.(*MessageId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Receiver_RejectMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Reject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiverServer).RejectMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/receiver.Receiver/RejectMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiverServer).RejectMessage(ctx, req.(*Reject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Receiver_Receive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiverArgs)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReceiverServer).Receive(m, &receiverReceiveServer{stream})
}

type Receiver_ReceiveServer interface {
	Send(*ReceiverResponse) error
	grpc.ServerStream
}

type receiverReceiveServer struct {
	grpc.ServerStream
}

func (x *receiverReceiveServer) Send(m *ReceiverResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Receiver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "receiver.Receiver",
	HandlerType: (*ReceiverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AcknowledgeMessage",
			Handler:    _Receiver_AcknowledgeMessage_Handler,
		},
		{
			MethodName: "RejectMessage",
			Handler:    _Receiver_RejectMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Receive",
			Handler:       _Receiver_Receive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "receiver.proto",
}
