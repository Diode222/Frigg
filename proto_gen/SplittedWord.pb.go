// Code generated by protoc-gen-go. DO NOT EDIT.
// source: SplittedWord.proto

package proto_gen

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

// 每个聊天消息只有一个句子，这个句子被分词后会重新组装成一个SplittedMessage，其中wordList表示句子分词得到的词语切片
type SplittedMessage struct {
	WordList             []string `protobuf:"bytes,1,rep,name=wordList" json:"wordList,omitempty"`
	Time                 *int64   `protobuf:"varint,2,req,name=time" json:"time,omitempty"`
	ChatPerson           *string  `protobuf:"bytes,3,req,name=chatPerson" json:"chatPerson,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SplittedMessage) Reset()         { *m = SplittedMessage{} }
func (m *SplittedMessage) String() string { return proto.CompactTextString(m) }
func (*SplittedMessage) ProtoMessage()    {}
func (*SplittedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb11982d59375097, []int{0}
}

func (m *SplittedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SplittedMessage.Unmarshal(m, b)
}
func (m *SplittedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SplittedMessage.Marshal(b, m, deterministic)
}
func (m *SplittedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SplittedMessage.Merge(m, src)
}
func (m *SplittedMessage) XXX_Size() int {
	return xxx_messageInfo_SplittedMessage.Size(m)
}
func (m *SplittedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SplittedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SplittedMessage proto.InternalMessageInfo

func (m *SplittedMessage) GetWordList() []string {
	if m != nil {
		return m.WordList
	}
	return nil
}

func (m *SplittedMessage) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *SplittedMessage) GetChatPerson() string {
	if m != nil && m.ChatPerson != nil {
		return *m.ChatPerson
	}
	return ""
}

type SplittedMessageList struct {
	SplittedMessages     []*SplittedMessage `protobuf:"bytes,1,rep,name=splittedMessages" json:"splittedMessages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SplittedMessageList) Reset()         { *m = SplittedMessageList{} }
func (m *SplittedMessageList) String() string { return proto.CompactTextString(m) }
func (*SplittedMessageList) ProtoMessage()    {}
func (*SplittedMessageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb11982d59375097, []int{1}
}

func (m *SplittedMessageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SplittedMessageList.Unmarshal(m, b)
}
func (m *SplittedMessageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SplittedMessageList.Marshal(b, m, deterministic)
}
func (m *SplittedMessageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SplittedMessageList.Merge(m, src)
}
func (m *SplittedMessageList) XXX_Size() int {
	return xxx_messageInfo_SplittedMessageList.Size(m)
}
func (m *SplittedMessageList) XXX_DiscardUnknown() {
	xxx_messageInfo_SplittedMessageList.DiscardUnknown(m)
}

var xxx_messageInfo_SplittedMessageList proto.InternalMessageInfo

func (m *SplittedMessageList) GetSplittedMessages() []*SplittedMessage {
	if m != nil {
		return m.SplittedMessages
	}
	return nil
}

func init() {
	proto.RegisterType((*SplittedMessage)(nil), "proto_gen.SplittedMessage")
	proto.RegisterType((*SplittedMessageList)(nil), "proto_gen.SplittedMessageList")
}

func init() { proto.RegisterFile("SplittedWord.proto", fileDescriptor_bb11982d59375097) }

var fileDescriptor_bb11982d59375097 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0a, 0x2e, 0xc8, 0xc9,
	0x2c, 0x29, 0x49, 0x4d, 0x09, 0xcf, 0x2f, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x04, 0x53, 0xf1, 0xe9, 0xa9, 0x79, 0x52, 0x22, 0xce, 0x19, 0x89, 0x25, 0xbe, 0xa9, 0xc5, 0xc5,
	0x89, 0xe9, 0xa9, 0x21, 0xb9, 0x05, 0x10, 0x05, 0x4a, 0x89, 0x5c, 0xfc, 0x30, 0x6d, 0x50, 0x39,
	0x21, 0x29, 0x2e, 0x8e, 0xf2, 0xfc, 0xa2, 0x14, 0x9f, 0xcc, 0xe2, 0x12, 0x09, 0x46, 0x05, 0x66,
	0x0d, 0xce, 0x20, 0x38, 0x5f, 0x48, 0x88, 0x8b, 0xa5, 0x24, 0x33, 0x37, 0x55, 0x82, 0x49, 0x81,
	0x49, 0x83, 0x39, 0x08, 0xcc, 0x16, 0x92, 0xe3, 0xe2, 0x4a, 0xce, 0x48, 0x2c, 0x09, 0x48, 0x2d,
	0x2a, 0xce, 0xcf, 0x93, 0x60, 0x56, 0x60, 0xd2, 0xe0, 0x0c, 0x42, 0x12, 0x51, 0x8a, 0xe5, 0x12,
	0x46, 0xb3, 0x02, 0x6c, 0x94, 0x1b, 0x97, 0x40, 0x31, 0xaa, 0x70, 0x31, 0xd8, 0x3a, 0x6e, 0x23,
	0x29, 0x3d, 0xb8, 0xab, 0xf5, 0xd0, 0x74, 0x06, 0x61, 0xe8, 0x31, 0x4a, 0xe3, 0x12, 0x00, 0x79,
	0x18, 0xac, 0x30, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28, 0x88, 0x4b, 0xc8, 0x3d, 0xb5,
	0x04, 0xc3, 0x63, 0x48, 0xe6, 0x22, 0x05, 0x06, 0xc8, 0x35, 0x52, 0x72, 0xb8, 0xed, 0x04, 0xc9,
	0x2b, 0x31, 0x00, 0x02, 0x00, 0x00, 0xff, 0xff, 0x25, 0x45, 0x93, 0xbe, 0x5f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WordSplitServiceClient is the client API for WordSplitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WordSplitServiceClient interface {
	GetSplittedMessage(ctx context.Context, in *ChatMessageList, opts ...grpc.CallOption) (*SplittedMessageList, error)
}

type wordSplitServiceClient struct {
	cc *grpc.ClientConn
}

func NewWordSplitServiceClient(cc *grpc.ClientConn) WordSplitServiceClient {
	return &wordSplitServiceClient{cc}
}

func (c *wordSplitServiceClient) GetSplittedMessage(ctx context.Context, in *ChatMessageList, opts ...grpc.CallOption) (*SplittedMessageList, error) {
	out := new(SplittedMessageList)
	err := c.cc.Invoke(ctx, "/proto_gen.WordSplitService/GetSplittedMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WordSplitServiceServer is the server API for WordSplitService service.
type WordSplitServiceServer interface {
	GetSplittedMessage(context.Context, *ChatMessageList) (*SplittedMessageList, error)
}

// UnimplementedWordSplitServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWordSplitServiceServer struct {
}

func (*UnimplementedWordSplitServiceServer) GetSplittedMessage(ctx context.Context, req *ChatMessageList) (*SplittedMessageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSplittedMessage not implemented")
}

func RegisterWordSplitServiceServer(s *grpc.Server, srv WordSplitServiceServer) {
	s.RegisterService(&_WordSplitService_serviceDesc, srv)
}

func _WordSplitService_GetSplittedMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMessageList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordSplitServiceServer).GetSplittedMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_gen.WordSplitService/GetSplittedMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordSplitServiceServer).GetSplittedMessage(ctx, req.(*ChatMessageList))
	}
	return interceptor(ctx, in, info, handler)
}

var _WordSplitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto_gen.WordSplitService",
	HandlerType: (*WordSplitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSplittedMessage",
			Handler:    _WordSplitService_GetSplittedMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "SplittedWord.proto",
}