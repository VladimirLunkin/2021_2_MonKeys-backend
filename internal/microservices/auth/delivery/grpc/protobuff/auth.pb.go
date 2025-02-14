//protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: auth.proto

package auth

import (
	context "context"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type Cookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cookie string `protobuf:"bytes,1,opt,name=Cookie,proto3" json:"Cookie,omitempty"`
}

func (x *Cookie) Reset() {
	*x = Cookie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cookie) ProtoMessage() {}

func (x *Cookie) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cookie.ProtoReflect.Descriptor instead.
func (*Cookie) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *Cookie) GetCookie() string {
	if x != nil {
		return x.Cookie
	}
	return ""
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Cookie string `protobuf:"bytes,2,opt,name=Cookie,proto3" json:"Cookie,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *Session) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Session) GetCookie() string {
	if x != nil {
		return x.Cookie
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Email        string   `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Password     string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	Name         string   `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Gender       string   `protobuf:"bytes,5,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Prefer       string   `protobuf:"bytes,6,opt,name=Prefer,proto3" json:"Prefer,omitempty"`
	FromAge      uint32   `protobuf:"varint,7,opt,name=FromAge,proto3" json:"FromAge,omitempty"`
	ToAge        uint32   `protobuf:"varint,8,opt,name=ToAge,proto3" json:"ToAge,omitempty"`
	Date         string   `protobuf:"bytes,9,opt,name=Date,proto3" json:"Date,omitempty"`
	Age          uint32   `protobuf:"varint,10,opt,name=Age,proto3" json:"Age,omitempty"`
	Description  string   `protobuf:"bytes,11,opt,name=Description,proto3" json:"Description,omitempty"`
	Imgs         []string `protobuf:"bytes,12,rep,name=Imgs,proto3" json:"Imgs,omitempty"`
	Tags         []string `protobuf:"bytes,13,rep,name=Tags,proto3" json:"Tags,omitempty"`
	ReportStatus string   `protobuf:"bytes,14,opt,name=ReportStatus,proto3" json:"ReportStatus,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *User) GetPrefer() string {
	if x != nil {
		return x.Prefer
	}
	return ""
}

func (x *User) GetFromAge() uint32 {
	if x != nil {
		return x.FromAge
	}
	return 0
}

func (x *User) GetToAge() uint32 {
	if x != nil {
		return x.ToAge
	}
	return 0
}

func (x *User) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *User) GetAge() uint32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *User) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *User) GetImgs() []string {
	if x != nil {
		return x.Imgs
	}
	return nil
}

func (x *User) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *User) GetReportStatus() string {
	if x != nil {
		return x.ReportStatus
	}
	return ""
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x06, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43,
	0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x22, 0x39, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x6f, 0x6b,
	0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65,
	0x22, 0xd0, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x41,
	0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x41, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6d, 0x67, 0x73, 0x18,
	0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x49, 0x6d, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x61, 0x67, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x54, 0x61, 0x67, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0x66, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x47, 0x72, 0x70, 0x63, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x1a, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0c,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x1a, 0x0d, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0x5a, 0x06, 0x2e,
	0x3b, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_auth_proto_goTypes = []interface{}{
	(*Cookie)(nil),  // 0: auth.Cookie
	(*Session)(nil), // 1: auth.Session
	(*User)(nil),    // 2: auth.User
}
var file_auth_proto_depIdxs = []int32{
	1, // 0: auth.AuthGrpcHandler.GetById:input_type -> auth.Session
	0, // 1: auth.AuthGrpcHandler.GetFromSession:input_type -> auth.Cookie
	2, // 2: auth.AuthGrpcHandler.GetById:output_type -> auth.User
	1, // 3: auth.AuthGrpcHandler.GetFromSession:output_type -> auth.Session
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cookie); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthGrpcHandlerClient is the client API for AuthGrpcHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthGrpcHandlerClient interface {
	GetById(ctx context.Context, in *Session, opts ...grpc.CallOption) (*User, error)
	GetFromSession(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*Session, error)
}

type authGrpcHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthGrpcHandlerClient(cc grpc.ClientConnInterface) AuthGrpcHandlerClient {
	return &authGrpcHandlerClient{cc}
}

func (c *authGrpcHandlerClient) GetById(ctx context.Context, in *Session, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/auth.AuthGrpcHandler/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authGrpcHandlerClient) GetFromSession(ctx context.Context, in *Cookie, opts ...grpc.CallOption) (*Session, error) {
	out := new(Session)
	err := c.cc.Invoke(ctx, "/auth.AuthGrpcHandler/GetFromSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthGrpcHandlerServer is the server API for AuthGrpcHandler service.
type AuthGrpcHandlerServer interface {
	GetById(context.Context, *Session) (*User, error)
	GetFromSession(context.Context, *Cookie) (*Session, error)
}

// UnimplementedAuthGrpcHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedAuthGrpcHandlerServer struct {
}

func (*UnimplementedAuthGrpcHandlerServer) GetById(context.Context, *Session) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (*UnimplementedAuthGrpcHandlerServer) GetFromSession(context.Context, *Cookie) (*Session, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFromSession not implemented")
}

func RegisterAuthGrpcHandlerServer(s *grpc.Server, srv AuthGrpcHandlerServer) {
	s.RegisterService(&_AuthGrpcHandler_serviceDesc, srv)
}

func _AuthGrpcHandler_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthGrpcHandlerServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthGrpcHandler/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthGrpcHandlerServer).GetById(ctx, req.(*Session))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthGrpcHandler_GetFromSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cookie)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthGrpcHandlerServer).GetFromSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthGrpcHandler/GetFromSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthGrpcHandlerServer).GetFromSession(ctx, req.(*Cookie))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthGrpcHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthGrpcHandler",
	HandlerType: (*AuthGrpcHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _AuthGrpcHandler_GetById_Handler,
		},
		{
			MethodName: "GetFromSession",
			Handler:    _AuthGrpcHandler_GetFromSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
