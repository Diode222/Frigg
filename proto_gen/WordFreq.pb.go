// Code generated by protoc-gen-go. DO NOT EDIT.
// source: WordFreq.proto

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

type PartOfSpeech_POSType int32

const (
	PartOfSpeech_NOUN      PartOfSpeech_POSType = 0
	PartOfSpeech_VERB      PartOfSpeech_POSType = 1
	PartOfSpeech_ADJECTIVE PartOfSpeech_POSType = 2
	PartOfSpeech_PHRASE    PartOfSpeech_POSType = 3
)

var PartOfSpeech_POSType_name = map[int32]string{
	0: "NOUN",
	1: "VERB",
	2: "ADJECTIVE",
	3: "PHRASE",
}

var PartOfSpeech_POSType_value = map[string]int32{
	"NOUN":      0,
	"VERB":      1,
	"ADJECTIVE": 2,
	"PHRASE":    3,
}

func (x PartOfSpeech_POSType) Enum() *PartOfSpeech_POSType {
	p := new(PartOfSpeech_POSType)
	*p = x
	return p
}

func (x PartOfSpeech_POSType) String() string {
	return proto.EnumName(PartOfSpeech_POSType_name, int32(x))
}

func (x *PartOfSpeech_POSType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(PartOfSpeech_POSType_value, data, "PartOfSpeech_POSType")
	if err != nil {
		return err
	}
	*x = PartOfSpeech_POSType(value)
	return nil
}

func (PartOfSpeech_POSType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3cb02e80b21beb96, []int{2, 0}
}

type WordFreq struct {
	Word                 *string  `protobuf:"bytes,1,req,name=word" json:"word,omitempty"`
	Count                *int32   `protobuf:"varint,2,req,name=count" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WordFreq) Reset()         { *m = WordFreq{} }
func (m *WordFreq) String() string { return proto.CompactTextString(m) }
func (*WordFreq) ProtoMessage()    {}
func (*WordFreq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cb02e80b21beb96, []int{0}
}

func (m *WordFreq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WordFreq.Unmarshal(m, b)
}
func (m *WordFreq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WordFreq.Marshal(b, m, deterministic)
}
func (m *WordFreq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WordFreq.Merge(m, src)
}
func (m *WordFreq) XXX_Size() int {
	return xxx_messageInfo_WordFreq.Size(m)
}
func (m *WordFreq) XXX_DiscardUnknown() {
	xxx_messageInfo_WordFreq.DiscardUnknown(m)
}

var xxx_messageInfo_WordFreq proto.InternalMessageInfo

func (m *WordFreq) GetWord() string {
	if m != nil && m.Word != nil {
		return *m.Word
	}
	return ""
}

func (m *WordFreq) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

type WordFreqList struct {
	WordFreqs            []*WordFreq `protobuf:"bytes,1,rep,name=wordFreqs" json:"wordFreqs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *WordFreqList) Reset()         { *m = WordFreqList{} }
func (m *WordFreqList) String() string { return proto.CompactTextString(m) }
func (*WordFreqList) ProtoMessage()    {}
func (*WordFreqList) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cb02e80b21beb96, []int{1}
}

func (m *WordFreqList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WordFreqList.Unmarshal(m, b)
}
func (m *WordFreqList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WordFreqList.Marshal(b, m, deterministic)
}
func (m *WordFreqList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WordFreqList.Merge(m, src)
}
func (m *WordFreqList) XXX_Size() int {
	return xxx_messageInfo_WordFreqList.Size(m)
}
func (m *WordFreqList) XXX_DiscardUnknown() {
	xxx_messageInfo_WordFreqList.DiscardUnknown(m)
}

var xxx_messageInfo_WordFreqList proto.InternalMessageInfo

func (m *WordFreqList) GetWordFreqs() []*WordFreq {
	if m != nil {
		return m.WordFreqs
	}
	return nil
}

type PartOfSpeech struct {
	Type                 *PartOfSpeech_POSType `protobuf:"varint,1,req,name=type,enum=proto_gen.PartOfSpeech_POSType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PartOfSpeech) Reset()         { *m = PartOfSpeech{} }
func (m *PartOfSpeech) String() string { return proto.CompactTextString(m) }
func (*PartOfSpeech) ProtoMessage()    {}
func (*PartOfSpeech) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cb02e80b21beb96, []int{2}
}

func (m *PartOfSpeech) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartOfSpeech.Unmarshal(m, b)
}
func (m *PartOfSpeech) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartOfSpeech.Marshal(b, m, deterministic)
}
func (m *PartOfSpeech) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartOfSpeech.Merge(m, src)
}
func (m *PartOfSpeech) XXX_Size() int {
	return xxx_messageInfo_PartOfSpeech.Size(m)
}
func (m *PartOfSpeech) XXX_DiscardUnknown() {
	xxx_messageInfo_PartOfSpeech.DiscardUnknown(m)
}

var xxx_messageInfo_PartOfSpeech proto.InternalMessageInfo

func (m *PartOfSpeech) GetType() PartOfSpeech_POSType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return PartOfSpeech_NOUN
}

func init() {
	proto.RegisterEnum("proto_gen.PartOfSpeech_POSType", PartOfSpeech_POSType_name, PartOfSpeech_POSType_value)
	proto.RegisterType((*WordFreq)(nil), "proto_gen.WordFreq")
	proto.RegisterType((*WordFreqList)(nil), "proto_gen.WordFreqList")
	proto.RegisterType((*PartOfSpeech)(nil), "proto_gen.PartOfSpeech")
}

func init() { proto.RegisterFile("WordFreq.proto", fileDescriptor_3cb02e80b21beb96) }

var fileDescriptor_3cb02e80b21beb96 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x0b, 0xcf, 0x2f, 0x4a,
	0x71, 0x2b, 0x4a, 0x2d, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x53, 0xf1, 0xe9,
	0xa9, 0x79, 0x4a, 0x26, 0x5c, 0x1c, 0x30, 0x49, 0x21, 0x21, 0x2e, 0x96, 0xf2, 0xfc, 0xa2, 0x14,
	0x09, 0x46, 0x05, 0x26, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x84, 0x8b, 0x35, 0x39, 0xbf, 0x34,
	0xaf, 0x44, 0x82, 0x49, 0x81, 0x49, 0x83, 0x35, 0x08, 0xc2, 0x51, 0x72, 0xe4, 0xe2, 0x81, 0xe9,
	0xf2, 0xc9, 0x2c, 0x2e, 0x11, 0x32, 0xe4, 0xe2, 0x2c, 0x87, 0xf2, 0x8b, 0x25, 0x18, 0x15, 0x98,
	0x35, 0xb8, 0x8d, 0x84, 0xf5, 0xe0, 0x96, 0xe8, 0xc1, 0xd4, 0x06, 0x21, 0x54, 0x29, 0xd5, 0x72,
	0xf1, 0x04, 0x24, 0x16, 0x95, 0xf8, 0xa7, 0x05, 0x17, 0xa4, 0xa6, 0x26, 0x67, 0x08, 0x19, 0x73,
	0xb1, 0x94, 0x54, 0x16, 0xa4, 0x82, 0x2d, 0xe7, 0x33, 0x92, 0x47, 0xd2, 0x8d, 0xac, 0x4c, 0x2f,
	0xc0, 0x3f, 0x38, 0xa4, 0xb2, 0x20, 0x35, 0x08, 0xac, 0x58, 0xc9, 0x82, 0x8b, 0x1d, 0x2a, 0x20,
	0xc4, 0xc1, 0xc5, 0xe2, 0xe7, 0x1f, 0xea, 0x27, 0xc0, 0x00, 0x62, 0x85, 0xb9, 0x06, 0x39, 0x09,
	0x30, 0x0a, 0xf1, 0x72, 0x71, 0x3a, 0xba, 0x78, 0xb9, 0x3a, 0x87, 0x78, 0x86, 0xb9, 0x0a, 0x30,
	0x09, 0x71, 0x71, 0xb1, 0x05, 0x78, 0x04, 0x39, 0x06, 0xbb, 0x0a, 0x30, 0x1b, 0xc5, 0x70, 0x09,
	0x23, 0xfb, 0x20, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x95, 0x8b, 0xdf, 0x3d, 0xb5,
	0x04, 0xc5, 0x6f, 0xe2, 0x38, 0x9c, 0x22, 0x25, 0x8e, 0xc5, 0x87, 0x20, 0x1d, 0x4a, 0x0c, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0xfb, 0xcd, 0x7b, 0x71, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WordFreqListServiceClient is the client API for WordFreqListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WordFreqListServiceClient interface {
	GetWordFreqList(ctx context.Context, in *PartOfSpeech, opts ...grpc.CallOption) (*WordFreqList, error)
}

type wordFreqListServiceClient struct {
	cc *grpc.ClientConn
}

func NewWordFreqListServiceClient(cc *grpc.ClientConn) WordFreqListServiceClient {
	return &wordFreqListServiceClient{cc}
}

func (c *wordFreqListServiceClient) GetWordFreqList(ctx context.Context, in *PartOfSpeech, opts ...grpc.CallOption) (*WordFreqList, error) {
	out := new(WordFreqList)
	err := c.cc.Invoke(ctx, "/proto_gen.WordFreqListService/GetWordFreqList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WordFreqListServiceServer is the server API for WordFreqListService service.
type WordFreqListServiceServer interface {
	GetWordFreqList(context.Context, *PartOfSpeech) (*WordFreqList, error)
}

// UnimplementedWordFreqListServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWordFreqListServiceServer struct {
}

func (*UnimplementedWordFreqListServiceServer) GetWordFreqList(ctx context.Context, req *PartOfSpeech) (*WordFreqList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWordFreqList not implemented")
}

func RegisterWordFreqListServiceServer(s *grpc.Server, srv WordFreqListServiceServer) {
	s.RegisterService(&_WordFreqListService_serviceDesc, srv)
}

func _WordFreqListService_GetWordFreqList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PartOfSpeech)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordFreqListServiceServer).GetWordFreqList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_gen.WordFreqListService/GetWordFreqList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordFreqListServiceServer).GetWordFreqList(ctx, req.(*PartOfSpeech))
	}
	return interceptor(ctx, in, info, handler)
}

var _WordFreqListService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto_gen.WordFreqListService",
	HandlerType: (*WordFreqListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWordFreqList",
			Handler:    _WordFreqListService_GetWordFreqList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "WordFreq.proto",
}