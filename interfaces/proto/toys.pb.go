// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: interfaces/proto/toys.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetPixivsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextMarker string `protobuf:"bytes,1,opt,name=next_marker,json=nextMarker,proto3" json:"next_marker,omitempty"`
	Page       int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetPixivsReq) Reset() {
	*x = GetPixivsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interfaces_proto_toys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPixivsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPixivsReq) ProtoMessage() {}

func (x *GetPixivsReq) ProtoReflect() protoreflect.Message {
	mi := &file_interfaces_proto_toys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPixivsReq.ProtoReflect.Descriptor instead.
func (*GetPixivsReq) Descriptor() ([]byte, []int) {
	return file_interfaces_proto_toys_proto_rawDescGZIP(), []int{0}
}

func (x *GetPixivsReq) GetNextMarker() string {
	if x != nil {
		return x.NextMarker
	}
	return ""
}

func (x *GetPixivsReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

var File_interfaces_proto_toys_proto protoreflect.FileDescriptor

var file_interfaces_proto_toys_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x43, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x69, 0x78, 0x69, 0x76, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x4d, 0x61, 0x72, 0x6b, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x32, 0x74, 0x0a, 0x04, 0x54, 0x6f, 0x79, 0x73, 0x12, 0x33, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x50, 0x69, 0x78, 0x69, 0x76, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x78, 0x69, 0x76, 0x73, 0x52, 0x65, 0x71, 0x1a,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interfaces_proto_toys_proto_rawDescOnce sync.Once
	file_interfaces_proto_toys_proto_rawDescData = file_interfaces_proto_toys_proto_rawDesc
)

func file_interfaces_proto_toys_proto_rawDescGZIP() []byte {
	file_interfaces_proto_toys_proto_rawDescOnce.Do(func() {
		file_interfaces_proto_toys_proto_rawDescData = protoimpl.X.CompressGZIP(file_interfaces_proto_toys_proto_rawDescData)
	})
	return file_interfaces_proto_toys_proto_rawDescData
}

var file_interfaces_proto_toys_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_interfaces_proto_toys_proto_goTypes = []interface{}{
	(*GetPixivsReq)(nil),  // 0: proto.GetPixivsReq
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
	(*Response)(nil),      // 2: proto.Response
}
var file_interfaces_proto_toys_proto_depIdxs = []int32{
	0, // 0: proto.Toys.GetPixivs:input_type -> proto.GetPixivsReq
	1, // 1: proto.Toys.GetFriends:input_type -> google.protobuf.Empty
	2, // 2: proto.Toys.GetPixivs:output_type -> proto.Response
	2, // 3: proto.Toys.GetFriends:output_type -> proto.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_interfaces_proto_toys_proto_init() }
func file_interfaces_proto_toys_proto_init() {
	if File_interfaces_proto_toys_proto != nil {
		return
	}
	file_interfaces_proto_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_interfaces_proto_toys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPixivsReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_interfaces_proto_toys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interfaces_proto_toys_proto_goTypes,
		DependencyIndexes: file_interfaces_proto_toys_proto_depIdxs,
		MessageInfos:      file_interfaces_proto_toys_proto_msgTypes,
	}.Build()
	File_interfaces_proto_toys_proto = out.File
	file_interfaces_proto_toys_proto_rawDesc = nil
	file_interfaces_proto_toys_proto_goTypes = nil
	file_interfaces_proto_toys_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ToysClient is the client API for Toys service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ToysClient interface {
	GetPixivs(ctx context.Context, in *GetPixivsReq, opts ...grpc.CallOption) (*Response, error)
	GetFriends(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Response, error)
}

type toysClient struct {
	cc grpc.ClientConnInterface
}

func NewToysClient(cc grpc.ClientConnInterface) ToysClient {
	return &toysClient{cc}
}

func (c *toysClient) GetPixivs(ctx context.Context, in *GetPixivsReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Toys/GetPixivs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toysClient) GetFriends(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.Toys/GetFriends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToysServer is the server API for Toys service.
type ToysServer interface {
	GetPixivs(context.Context, *GetPixivsReq) (*Response, error)
	GetFriends(context.Context, *emptypb.Empty) (*Response, error)
}

// UnimplementedToysServer can be embedded to have forward compatible implementations.
type UnimplementedToysServer struct {
}

func (*UnimplementedToysServer) GetPixivs(context.Context, *GetPixivsReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPixivs not implemented")
}
func (*UnimplementedToysServer) GetFriends(context.Context, *emptypb.Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriends not implemented")
}

func RegisterToysServer(s *grpc.Server, srv ToysServer) {
	s.RegisterService(&_Toys_serviceDesc, srv)
}

func _Toys_GetPixivs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPixivsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToysServer).GetPixivs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Toys/GetPixivs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToysServer).GetPixivs(ctx, req.(*GetPixivsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Toys_GetFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToysServer).GetFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Toys/GetFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToysServer).GetFriends(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Toys_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Toys",
	HandlerType: (*ToysServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPixivs",
			Handler:    _Toys_GetPixivs_Handler,
		},
		{
			MethodName: "GetFriends",
			Handler:    _Toys_GetFriends_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interfaces/proto/toys.proto",
}
