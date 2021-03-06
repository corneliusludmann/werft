// Code generated by protoc-gen-go. DO NOT EDIT.
// source: werft-ui.proto

package v1

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

type ListJobSpecsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListJobSpecsRequest) Reset()         { *m = ListJobSpecsRequest{} }
func (m *ListJobSpecsRequest) String() string { return proto.CompactTextString(m) }
func (*ListJobSpecsRequest) ProtoMessage()    {}
func (*ListJobSpecsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d41ca2a021dc92d, []int{0}
}

func (m *ListJobSpecsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobSpecsRequest.Unmarshal(m, b)
}
func (m *ListJobSpecsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobSpecsRequest.Marshal(b, m, deterministic)
}
func (m *ListJobSpecsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobSpecsRequest.Merge(m, src)
}
func (m *ListJobSpecsRequest) XXX_Size() int {
	return xxx_messageInfo_ListJobSpecsRequest.Size(m)
}
func (m *ListJobSpecsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobSpecsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobSpecsRequest proto.InternalMessageInfo

type ListJobSpecsResponse struct {
	Repo                 *Repository          `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Path                 string               `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Description          string               `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Arguments            []*DesiredAnnotation `protobuf:"bytes,5,rep,name=arguments,proto3" json:"arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListJobSpecsResponse) Reset()         { *m = ListJobSpecsResponse{} }
func (m *ListJobSpecsResponse) String() string { return proto.CompactTextString(m) }
func (*ListJobSpecsResponse) ProtoMessage()    {}
func (*ListJobSpecsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d41ca2a021dc92d, []int{1}
}

func (m *ListJobSpecsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobSpecsResponse.Unmarshal(m, b)
}
func (m *ListJobSpecsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobSpecsResponse.Marshal(b, m, deterministic)
}
func (m *ListJobSpecsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobSpecsResponse.Merge(m, src)
}
func (m *ListJobSpecsResponse) XXX_Size() int {
	return xxx_messageInfo_ListJobSpecsResponse.Size(m)
}
func (m *ListJobSpecsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobSpecsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobSpecsResponse proto.InternalMessageInfo

func (m *ListJobSpecsResponse) GetRepo() *Repository {
	if m != nil {
		return m.Repo
	}
	return nil
}

func (m *ListJobSpecsResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListJobSpecsResponse) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ListJobSpecsResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ListJobSpecsResponse) GetArguments() []*DesiredAnnotation {
	if m != nil {
		return m.Arguments
	}
	return nil
}

// DesiredAnnotation describes an annotation a job should have
type DesiredAnnotation struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Required             bool     `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DesiredAnnotation) Reset()         { *m = DesiredAnnotation{} }
func (m *DesiredAnnotation) String() string { return proto.CompactTextString(m) }
func (*DesiredAnnotation) ProtoMessage()    {}
func (*DesiredAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d41ca2a021dc92d, []int{2}
}

func (m *DesiredAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DesiredAnnotation.Unmarshal(m, b)
}
func (m *DesiredAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DesiredAnnotation.Marshal(b, m, deterministic)
}
func (m *DesiredAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DesiredAnnotation.Merge(m, src)
}
func (m *DesiredAnnotation) XXX_Size() int {
	return xxx_messageInfo_DesiredAnnotation.Size(m)
}
func (m *DesiredAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_DesiredAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_DesiredAnnotation proto.InternalMessageInfo

func (m *DesiredAnnotation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DesiredAnnotation) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (m *DesiredAnnotation) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*ListJobSpecsRequest)(nil), "v1.ListJobSpecsRequest")
	proto.RegisterType((*ListJobSpecsResponse)(nil), "v1.ListJobSpecsResponse")
	proto.RegisterType((*DesiredAnnotation)(nil), "v1.DesiredAnnotation")
}

func init() { proto.RegisterFile("werft-ui.proto", fileDescriptor_8d41ca2a021dc92d) }

var fileDescriptor_8d41ca2a021dc92d = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xcd, 0x5a, 0x75, 0xfb, 0x2a, 0x03, 0xa3, 0xc3, 0xd0, 0x53, 0xc9, 0xa9, 0x17, 0x8b,
	0xeb, 0x7e, 0x81, 0xa0, 0x07, 0xc5, 0x83, 0x44, 0xc4, 0x73, 0xb7, 0x7d, 0x6a, 0x0e, 0x4b, 0xb2,
	0x24, 0x9d, 0xf8, 0xd3, 0xfc, 0x77, 0x92, 0x14, 0x66, 0xb5, 0xde, 0xbe, 0x3e, 0xef, 0x0b, 0x7d,
	0x78, 0x03, 0xd3, 0x0f, 0xb4, 0xaf, 0xfe, 0xb2, 0x95, 0x95, 0xb1, 0xda, 0x6b, 0x3a, 0xda, 0xcd,
	0xf3, 0x2c, 0xb2, 0x0e, 0xf0, 0x19, 0x9c, 0x3d, 0x48, 0xe7, 0xef, 0xf5, 0xf2, 0xc9, 0xe0, 0xca,
	0x09, 0xdc, 0xb6, 0xe8, 0x3c, 0xff, 0x22, 0x70, 0xfe, 0x9b, 0x3b, 0xa3, 0x95, 0x43, 0xca, 0x21,
	0xb5, 0x68, 0x34, 0x23, 0x05, 0x29, 0xb3, 0x7a, 0x5a, 0xed, 0xe6, 0x95, 0x40, 0xa3, 0x9d, 0xf4,
	0xda, 0x7e, 0x8a, 0x98, 0x51, 0x0a, 0xa9, 0x6a, 0x36, 0xc8, 0x46, 0x05, 0x29, 0x27, 0x22, 0xde,
	0x81, 0x99, 0xc6, 0xbf, 0xb3, 0xa4, 0x63, 0xe1, 0xa6, 0x05, 0x64, 0x6b, 0x74, 0x2b, 0x2b, 0x8d,
	0x97, 0x5a, 0xb1, 0x34, 0x46, 0x7d, 0x44, 0x17, 0x30, 0x69, 0xec, 0x5b, 0xbb, 0x41, 0xe5, 0x1d,
	0x3b, 0x2c, 0x92, 0x32, 0xab, 0x67, 0xe1, 0x97, 0x37, 0xe8, 0xa4, 0xc5, 0xf5, 0xb5, 0x52, 0xda,
	0x37, 0xa1, 0x29, 0x7e, 0x7a, 0x1c, 0xe1, 0x74, 0x90, 0xef, 0x9d, 0x48, 0xcf, 0x29, 0x87, 0xb1,
	0xc5, 0x6d, 0x1b, 0x9a, 0xd1, 0x75, 0x2c, 0xf6, 0xdf, 0x7f, 0xdd, 0x92, 0x81, 0x5b, 0xfd, 0x08,
	0xc7, 0x2f, 0x61, 0xc8, 0xe7, 0x3b, 0x7a, 0x0b, 0x27, 0xfd, 0xb1, 0xe8, 0x45, 0x70, 0xfc, 0x67,
	0xd6, 0x9c, 0x0d, 0x83, 0x6e, 0x57, 0x7e, 0x70, 0x45, 0x96, 0x47, 0xf1, 0x49, 0x16, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x14, 0xb8, 0xdd, 0x0d, 0xb5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WerftUIClient is the client API for WerftUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WerftUIClient interface {
	// ListJobSpecs returns a list of jobs that can be started through the UI.
	ListJobSpecs(ctx context.Context, in *ListJobSpecsRequest, opts ...grpc.CallOption) (WerftUI_ListJobSpecsClient, error)
}

type werftUIClient struct {
	cc *grpc.ClientConn
}

func NewWerftUIClient(cc *grpc.ClientConn) WerftUIClient {
	return &werftUIClient{cc}
}

func (c *werftUIClient) ListJobSpecs(ctx context.Context, in *ListJobSpecsRequest, opts ...grpc.CallOption) (WerftUI_ListJobSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WerftUI_serviceDesc.Streams[0], "/v1.WerftUI/ListJobSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &werftUIListJobSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WerftUI_ListJobSpecsClient interface {
	Recv() (*ListJobSpecsResponse, error)
	grpc.ClientStream
}

type werftUIListJobSpecsClient struct {
	grpc.ClientStream
}

func (x *werftUIListJobSpecsClient) Recv() (*ListJobSpecsResponse, error) {
	m := new(ListJobSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WerftUIServer is the server API for WerftUI service.
type WerftUIServer interface {
	// ListJobSpecs returns a list of jobs that can be started through the UI.
	ListJobSpecs(*ListJobSpecsRequest, WerftUI_ListJobSpecsServer) error
}

// UnimplementedWerftUIServer can be embedded to have forward compatible implementations.
type UnimplementedWerftUIServer struct {
}

func (*UnimplementedWerftUIServer) ListJobSpecs(req *ListJobSpecsRequest, srv WerftUI_ListJobSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListJobSpecs not implemented")
}

func RegisterWerftUIServer(s *grpc.Server, srv WerftUIServer) {
	s.RegisterService(&_WerftUI_serviceDesc, srv)
}

func _WerftUI_ListJobSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListJobSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WerftUIServer).ListJobSpecs(m, &werftUIListJobSpecsServer{stream})
}

type WerftUI_ListJobSpecsServer interface {
	Send(*ListJobSpecsResponse) error
	grpc.ServerStream
}

type werftUIListJobSpecsServer struct {
	grpc.ServerStream
}

func (x *werftUIListJobSpecsServer) Send(m *ListJobSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _WerftUI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.WerftUI",
	HandlerType: (*WerftUIServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListJobSpecs",
			Handler:       _WerftUI_ListJobSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "werft-ui.proto",
}
