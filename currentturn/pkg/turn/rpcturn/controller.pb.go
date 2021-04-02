// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: controller.proto

package rpcturn

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	rpcgrid "github.com/theothertomelliott/tic-tac-toverengineered/grid/pkg/grid/rpcgrid"
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

type TakeTurnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId   string            `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	Mark     string            `protobuf:"bytes,2,opt,name=mark,proto3" json:"mark,omitempty"`
	Position *rpcgrid.Position `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *TakeTurnRequest) Reset() {
	*x = TakeTurnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeTurnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeTurnRequest) ProtoMessage() {}

func (x *TakeTurnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeTurnRequest.ProtoReflect.Descriptor instead.
func (*TakeTurnRequest) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{0}
}

func (x *TakeTurnRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *TakeTurnRequest) GetMark() string {
	if x != nil {
		return x.Mark
	}
	return ""
}

func (x *TakeTurnRequest) GetPosition() *rpcgrid.Position {
	if x != nil {
		return x.Position
	}
	return nil
}

type TakeTurnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TakeTurnResponse) Reset() {
	*x = TakeTurnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeTurnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeTurnResponse) ProtoMessage() {}

func (x *TakeTurnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeTurnResponse.ProtoReflect.Descriptor instead.
func (*TakeTurnResponse) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{1}
}

type NextPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId string `protobuf:"bytes,1,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
}

func (x *NextPlayerRequest) Reset() {
	*x = NextPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextPlayerRequest) ProtoMessage() {}

func (x *NextPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextPlayerRequest.ProtoReflect.Descriptor instead.
func (*NextPlayerRequest) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{2}
}

func (x *NextPlayerRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

type NextPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mark string `protobuf:"bytes,1,opt,name=mark,proto3" json:"mark,omitempty"`
}

func (x *NextPlayerResponse) Reset() {
	*x = NextPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NextPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NextPlayerResponse) ProtoMessage() {}

func (x *NextPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NextPlayerResponse.ProtoReflect.Descriptor instead.
func (*NextPlayerResponse) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{3}
}

func (x *NextPlayerResponse) GetMark() string {
	if x != nil {
		return x.Mark
	}
	return ""
}

var File_controller_proto protoreflect.FileDescriptor

var file_controller_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0a, 0x67, 0x72, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x65,
	0x0a, 0x0f, 0x54, 0x61, 0x6b, 0x65, 0x54, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61,
	0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x25,
	0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x12, 0x0a, 0x10, 0x54, 0x61, 0x6b, 0x65, 0x54, 0x75, 0x72,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x11, 0x4e, 0x65, 0x78,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x12, 0x4e, 0x65, 0x78, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x72,
	0x6b, 0x32, 0x78, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12,
	0x31, 0x0a, 0x08, 0x54, 0x61, 0x6b, 0x65, 0x54, 0x75, 0x72, 0x6e, 0x12, 0x10, 0x2e, 0x54, 0x61,
	0x6b, 0x65, 0x54, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x54, 0x61, 0x6b, 0x65, 0x54, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x12, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x54, 0x5a, 0x52, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x74, 0x6f, 0x6d, 0x65, 0x6c, 0x6c, 0x69, 0x6f, 0x74, 0x74, 0x2f, 0x74, 0x69, 0x63,
	0x2d, 0x74, 0x61, 0x63, 0x2d, 0x74, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x65, 0x72, 0x65, 0x64, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x74, 0x75, 0x72, 0x6e,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x75, 0x72, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x74, 0x75, 0x72,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_proto_rawDescOnce sync.Once
	file_controller_proto_rawDescData = file_controller_proto_rawDesc
)

func file_controller_proto_rawDescGZIP() []byte {
	file_controller_proto_rawDescOnce.Do(func() {
		file_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_proto_rawDescData)
	})
	return file_controller_proto_rawDescData
}

var file_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_controller_proto_goTypes = []interface{}{
	(*TakeTurnRequest)(nil),    // 0: TakeTurnRequest
	(*TakeTurnResponse)(nil),   // 1: TakeTurnResponse
	(*NextPlayerRequest)(nil),  // 2: NextPlayerRequest
	(*NextPlayerResponse)(nil), // 3: NextPlayerResponse
	(*rpcgrid.Position)(nil),   // 4: Position
}
var file_controller_proto_depIdxs = []int32{
	4, // 0: TakeTurnRequest.position:type_name -> Position
	0, // 1: Controller.TakeTurn:input_type -> TakeTurnRequest
	2, // 2: Controller.NextPlayer:input_type -> NextPlayerRequest
	1, // 3: Controller.TakeTurn:output_type -> TakeTurnResponse
	3, // 4: Controller.NextPlayer:output_type -> NextPlayerResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_controller_proto_init() }
func file_controller_proto_init() {
	if File_controller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeTurnRequest); i {
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
		file_controller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeTurnResponse); i {
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
		file_controller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextPlayerRequest); i {
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
		file_controller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NextPlayerResponse); i {
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
			RawDescriptor: file_controller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controller_proto_goTypes,
		DependencyIndexes: file_controller_proto_depIdxs,
		MessageInfos:      file_controller_proto_msgTypes,
	}.Build()
	File_controller_proto = out.File
	file_controller_proto_rawDesc = nil
	file_controller_proto_goTypes = nil
	file_controller_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ControllerClient is the client API for Controller service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControllerClient interface {
	TakeTurn(ctx context.Context, in *TakeTurnRequest, opts ...grpc.CallOption) (*TakeTurnResponse, error)
	NextPlayer(ctx context.Context, in *NextPlayerRequest, opts ...grpc.CallOption) (*NextPlayerResponse, error)
}

type controllerClient struct {
	cc grpc.ClientConnInterface
}

func NewControllerClient(cc grpc.ClientConnInterface) ControllerClient {
	return &controllerClient{cc}
}

func (c *controllerClient) TakeTurn(ctx context.Context, in *TakeTurnRequest, opts ...grpc.CallOption) (*TakeTurnResponse, error) {
	out := new(TakeTurnResponse)
	err := c.cc.Invoke(ctx, "/Controller/TakeTurn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controllerClient) NextPlayer(ctx context.Context, in *NextPlayerRequest, opts ...grpc.CallOption) (*NextPlayerResponse, error) {
	out := new(NextPlayerResponse)
	err := c.cc.Invoke(ctx, "/Controller/NextPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControllerServer is the server API for Controller service.
type ControllerServer interface {
	TakeTurn(context.Context, *TakeTurnRequest) (*TakeTurnResponse, error)
	NextPlayer(context.Context, *NextPlayerRequest) (*NextPlayerResponse, error)
}

// UnimplementedControllerServer can be embedded to have forward compatible implementations.
type UnimplementedControllerServer struct {
}

func (*UnimplementedControllerServer) TakeTurn(context.Context, *TakeTurnRequest) (*TakeTurnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TakeTurn not implemented")
}
func (*UnimplementedControllerServer) NextPlayer(context.Context, *NextPlayerRequest) (*NextPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextPlayer not implemented")
}

func RegisterControllerServer(s *grpc.Server, srv ControllerServer) {
	s.RegisterService(&_Controller_serviceDesc, srv)
}

func _Controller_TakeTurn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TakeTurnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServer).TakeTurn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Controller/TakeTurn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServer).TakeTurn(ctx, req.(*TakeTurnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Controller_NextPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControllerServer).NextPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Controller/NextPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControllerServer).NextPlayer(ctx, req.(*NextPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Controller_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Controller",
	HandlerType: (*ControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TakeTurn",
			Handler:    _Controller_TakeTurn_Handler,
		},
		{
			MethodName: "NextPlayer",
			Handler:    _Controller_NextPlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller.proto",
}
