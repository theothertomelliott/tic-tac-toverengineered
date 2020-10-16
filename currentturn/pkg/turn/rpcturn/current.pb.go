// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: current.proto

package rpcturn

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId string `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (x *PlayerRequest) Reset() {
	*x = PlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_current_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerRequest) ProtoMessage() {}

func (x *PlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_current_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerRequest.ProtoReflect.Descriptor instead.
func (*PlayerRequest) Descriptor() ([]byte, []int) {
	return file_current_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

type PlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mark string `protobuf:"bytes,1,opt,name=mark,proto3" json:"mark,omitempty"`
}

func (x *PlayerResponse) Reset() {
	*x = PlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_current_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerResponse) ProtoMessage() {}

func (x *PlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_current_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerResponse.ProtoReflect.Descriptor instead.
func (*PlayerResponse) Descriptor() ([]byte, []int) {
	return file_current_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerResponse) GetMark() string {
	if x != nil {
		return x.Mark
	}
	return ""
}

type NextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId string `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (x *NextRequest) Reset() {
	*x = NextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_current_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextRequest) ProtoMessage() {}

func (x *NextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_current_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextRequest.ProtoReflect.Descriptor instead.
func (*NextRequest) Descriptor() ([]byte, []int) {
	return file_current_proto_rawDescGZIP(), []int{2}
}

func (x *NextRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

type NextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NextResponse) Reset() {
	*x = NextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_current_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextResponse) ProtoMessage() {}

func (x *NextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_current_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextResponse.ProtoReflect.Descriptor instead.
func (*NextResponse) Descriptor() ([]byte, []int) {
	return file_current_proto_rawDescGZIP(), []int{3}
}

var File_current_proto protoreflect.FileDescriptor

var file_current_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x28, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x22, 0x24, 0x0a, 0x0e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d,
	0x61, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x22,
	0x26, 0x0a, 0x0b, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x4e, 0x65, 0x78, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x5d, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x25, 0x0a, 0x04, 0x4e, 0x65, 0x78, 0x74, 0x12, 0x0c, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x74, 0x6f, 0x6d,
	0x65, 0x6c, 0x6c, 0x69, 0x6f, 0x74, 0x74, 0x2f, 0x74, 0x69, 0x63, 0x2d, 0x74, 0x61, 0x63, 0x2d,
	0x74, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x65, 0x72, 0x65, 0x64, 0x2f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x74, 0x75, 0x72, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x74, 0x75, 0x72, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x74, 0x75, 0x72, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_current_proto_rawDescOnce sync.Once
	file_current_proto_rawDescData = file_current_proto_rawDesc
)

func file_current_proto_rawDescGZIP() []byte {
	file_current_proto_rawDescOnce.Do(func() {
		file_current_proto_rawDescData = protoimpl.X.CompressGZIP(file_current_proto_rawDescData)
	})
	return file_current_proto_rawDescData
}

var file_current_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_current_proto_goTypes = []interface{}{
	(*PlayerRequest)(nil),  // 0: PlayerRequest
	(*PlayerResponse)(nil), // 1: PlayerResponse
	(*NextRequest)(nil),    // 2: NextRequest
	(*NextResponse)(nil),   // 3: NextResponse
}
var file_current_proto_depIdxs = []int32{
	0, // 0: Current.Player:input_type -> PlayerRequest
	2, // 1: Current.Next:input_type -> NextRequest
	1, // 2: Current.Player:output_type -> PlayerResponse
	3, // 3: Current.Next:output_type -> NextResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_current_proto_init() }
func file_current_proto_init() {
	if File_current_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_current_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerRequest); i {
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
		file_current_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerResponse); i {
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
		file_current_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextRequest); i {
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
		file_current_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextResponse); i {
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
			RawDescriptor: file_current_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_current_proto_goTypes,
		DependencyIndexes: file_current_proto_depIdxs,
		MessageInfos:      file_current_proto_msgTypes,
	}.Build()
	File_current_proto = out.File
	file_current_proto_rawDesc = nil
	file_current_proto_goTypes = nil
	file_current_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CurrentClient is the client API for Current service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CurrentClient interface {
	Player(ctx context.Context, in *PlayerRequest, opts ...grpc.CallOption) (*PlayerResponse, error)
	Next(ctx context.Context, in *NextRequest, opts ...grpc.CallOption) (*NextResponse, error)
}

type currentClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrentClient(cc grpc.ClientConnInterface) CurrentClient {
	return &currentClient{cc}
}

func (c *currentClient) Player(ctx context.Context, in *PlayerRequest, opts ...grpc.CallOption) (*PlayerResponse, error) {
	out := new(PlayerResponse)
	err := c.cc.Invoke(ctx, "/Current/Player", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currentClient) Next(ctx context.Context, in *NextRequest, opts ...grpc.CallOption) (*NextResponse, error) {
	out := new(NextResponse)
	err := c.cc.Invoke(ctx, "/Current/Next", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrentServer is the server API for Current service.
type CurrentServer interface {
	Player(context.Context, *PlayerRequest) (*PlayerResponse, error)
	Next(context.Context, *NextRequest) (*NextResponse, error)
}

// UnimplementedCurrentServer can be embedded to have forward compatible implementations.
type UnimplementedCurrentServer struct {
}

func (*UnimplementedCurrentServer) Player(context.Context, *PlayerRequest) (*PlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Player not implemented")
}
func (*UnimplementedCurrentServer) Next(context.Context, *NextRequest) (*NextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Next not implemented")
}

func RegisterCurrentServer(s *grpc.Server, srv CurrentServer) {
	s.RegisterService(&_Current_serviceDesc, srv)
}

func _Current_Player_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentServer).Player(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Current/Player",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentServer).Player(ctx, req.(*PlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Current_Next_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrentServer).Next(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Current/Next",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrentServer).Next(ctx, req.(*NextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Current_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Current",
	HandlerType: (*CurrentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Player",
			Handler:    _Current_Player_Handler,
		},
		{
			MethodName: "Next",
			Handler:    _Current_Next_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "current.proto",
}
