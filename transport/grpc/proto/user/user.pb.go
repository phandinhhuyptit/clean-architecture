// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/proto/user/user.proto

package user

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,json=_id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Fullname             string   `protobuf:"bytes,3,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Role                 string   `protobuf:"bytes,7,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca47c89e84f91d88, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type FindByIDRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindByIDRequest) Reset()         { *m = FindByIDRequest{} }
func (m *FindByIDRequest) String() string { return proto.CompactTextString(m) }
func (*FindByIDRequest) ProtoMessage()    {}
func (*FindByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca47c89e84f91d88, []int{1}
}

func (m *FindByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindByIDRequest.Unmarshal(m, b)
}
func (m *FindByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindByIDRequest.Marshal(b, m, deterministic)
}
func (m *FindByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindByIDRequest.Merge(m, src)
}
func (m *FindByIDRequest) XXX_Size() int {
	return xxx_messageInfo_FindByIDRequest.Size(m)
}
func (m *FindByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindByIDRequest proto.InternalMessageInfo

func (m *FindByIDRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type FindByIDResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindByIDResponse) Reset()         { *m = FindByIDResponse{} }
func (m *FindByIDResponse) String() string { return proto.CompactTextString(m) }
func (*FindByIDResponse) ProtoMessage()    {}
func (*FindByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ca47c89e84f91d88, []int{2}
}

func (m *FindByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindByIDResponse.Unmarshal(m, b)
}
func (m *FindByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindByIDResponse.Marshal(b, m, deterministic)
}
func (m *FindByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindByIDResponse.Merge(m, src)
}
func (m *FindByIDResponse) XXX_Size() int {
	return xxx_messageInfo_FindByIDResponse.Size(m)
}
func (m *FindByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindByIDResponse proto.InternalMessageInfo

func (m *FindByIDResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*FindByIDRequest)(nil), "user.FindByIDRequest")
	proto.RegisterType((*FindByIDResponse)(nil), "user.FindByIDResponse")
}

func init() { proto.RegisterFile("grpc/proto/user/user.proto", fileDescriptor_ca47c89e84f91d88) }

var fileDescriptor_ca47c89e84f91d88 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x4d, 0x9b, 0xc6, 0x3a, 0x3d, 0x54, 0x06, 0x2d, 0x4b, 0x0e, 0x22, 0x39, 0xe9, 0xa5,
	0x85, 0x78, 0xf4, 0x26, 0x45, 0xa8, 0xc7, 0x88, 0x67, 0x89, 0xcd, 0xa8, 0x0b, 0xe9, 0xee, 0xba,
	0xdb, 0x28, 0x7e, 0x26, 0xbf, 0xa4, 0xcc, 0x6c, 0xa3, 0x90, 0xcb, 0x32, 0xbf, 0xf7, 0x1e, 0x3b,
	0x7f, 0x20, 0x7f, 0xf3, 0x6e, 0xbb, 0x72, 0xde, 0xee, 0xed, 0xaa, 0x0b, 0xe4, 0xe5, 0x59, 0x0a,
	0x63, 0xca, 0x75, 0xf1, 0x93, 0x40, 0xfa, 0x14, 0xc8, 0xe3, 0x1c, 0x46, 0xba, 0x51, 0xc9, 0x65,
	0x72, 0x75, 0x52, 0x8d, 0x9f, 0x75, 0x83, 0x39, 0x4c, 0x39, 0x61, 0xea, 0x1d, 0xa9, 0x91, 0xc8,
	0x7f, 0xcc, 0xde, 0x6b, 0xd7, 0xb6, 0xe2, 0x8d, 0xa3, 0xd7, 0x33, 0x7b, 0xae, 0x0e, 0xe1, 0xcb,
	0xfa, 0x46, 0xa5, 0xd1, 0xeb, 0x19, 0xcf, 0x60, 0x42, 0xbb, 0x5a, 0xb7, 0x6a, 0x22, 0x46, 0x04,
	0x56, 0xdd, 0xbb, 0x35, 0xa4, 0xb2, 0xa8, 0x0a, 0x20, 0x42, 0xea, 0x6d, 0x4b, 0xea, 0x58, 0x44,
	0xa9, 0x8b, 0x6b, 0x98, 0xdf, 0x6b, 0xd3, 0xdc, 0x7d, 0x6f, 0xd6, 0x15, 0x7d, 0x74, 0x14, 0xf6,
	0xb8, 0x80, 0x8c, 0xc7, 0xda, 0xac, 0x0f, 0xb3, 0x1f, 0xa8, 0x28, 0xe1, 0xf4, 0x3f, 0x1a, 0x9c,
	0x35, 0x81, 0xf0, 0x02, 0x64, 0x69, 0x49, 0xce, 0x4a, 0x58, 0xca, 0x35, 0x78, 0xfb, 0x4a, 0xf4,
	0xf2, 0x01, 0x66, 0x4c, 0x8f, 0xe4, 0x3f, 0xf5, 0x96, 0xf0, 0x16, 0xa6, 0xfd, 0x17, 0x78, 0x1e,
	0xc3, 0x83, 0xee, 0xf9, 0x62, 0x28, 0xc7, 0x4e, 0xc5, 0xd1, 0x4b, 0x26, 0x57, 0xbe, 0xf9, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xc1, 0xae, 0x39, 0x32, 0x83, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	FindByID(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*FindByIDResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) FindByID(ctx context.Context, in *FindByIDRequest, opts ...grpc.CallOption) (*FindByIDResponse, error) {
	out := new(FindByIDResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/FindByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	FindByID(context.Context, *FindByIDRequest) (*FindByIDResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) FindByID(ctx context.Context, req *FindByIDRequest) (*FindByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByID not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_FindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FindByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindByID(ctx, req.(*FindByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindByID",
			Handler:    _UserService_FindByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto/user/user.proto",
}
