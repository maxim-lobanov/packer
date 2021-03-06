// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/iam/v1/user_account_service.proto

package iam

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetUserAccountRequest struct {
	// ID of the UserAccount resource to return.
	UserAccountId        string   `protobuf:"bytes,1,opt,name=user_account_id,json=userAccountId,proto3" json:"user_account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserAccountRequest) Reset()         { *m = GetUserAccountRequest{} }
func (m *GetUserAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserAccountRequest) ProtoMessage()    {}
func (*GetUserAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68bca2edf26b1c68, []int{0}
}

func (m *GetUserAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserAccountRequest.Unmarshal(m, b)
}
func (m *GetUserAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserAccountRequest.Marshal(b, m, deterministic)
}
func (m *GetUserAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserAccountRequest.Merge(m, src)
}
func (m *GetUserAccountRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserAccountRequest.Size(m)
}
func (m *GetUserAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserAccountRequest proto.InternalMessageInfo

func (m *GetUserAccountRequest) GetUserAccountId() string {
	if m != nil {
		return m.UserAccountId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserAccountRequest)(nil), "yandex.cloud.iam.v1.GetUserAccountRequest")
}

func init() {
	proto.RegisterFile("yandex/cloud/iam/v1/user_account_service.proto", fileDescriptor_68bca2edf26b1c68)
}

var fileDescriptor_68bca2edf26b1c68 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xab, 0x4c, 0xcc, 0x4b,
	0x49, 0xad, 0xd0, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0xcf, 0x4c, 0xcc, 0xd5, 0x2f, 0x33, 0xd4,
	0x2f, 0x2d, 0x4e, 0x2d, 0x8a, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x89, 0x2f, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x86, 0xa8, 0xd7, 0x03,
	0xab, 0xd7, 0xcb, 0x4c, 0xcc, 0xd5, 0x2b, 0x33, 0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49,
	0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b,
	0x86, 0x68, 0x91, 0x52, 0x23, 0x64, 0x05, 0x54, 0x9d, 0x2c, 0x8a, 0xba, 0xb2, 0xc4, 0x9c, 0xcc,
	0x14, 0xb0, 0x39, 0x10, 0x69, 0x25, 0x5f, 0x2e, 0x51, 0xf7, 0xd4, 0x92, 0xd0, 0xe2, 0xd4, 0x22,
	0x47, 0x88, 0xb6, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x13, 0x2e, 0x7e, 0x14, 0x07,
	0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x3a, 0xf1, 0xbc, 0x38, 0x6e, 0xc8, 0xd8, 0x75,
	0xc2, 0x90, 0xc5, 0xc6, 0xd6, 0xd4, 0x20, 0x88, 0xb7, 0x14, 0xa1, 0xd5, 0x33, 0xc5, 0x68, 0x16,
	0x23, 0x97, 0x10, 0x92, 0x61, 0xc1, 0x10, 0x5f, 0x0a, 0x35, 0x33, 0x72, 0x31, 0xbb, 0xa7, 0x96,
	0x08, 0x69, 0xe9, 0x61, 0xf1, 0xa8, 0x1e, 0x56, 0x07, 0x48, 0x29, 0x60, 0x55, 0x8b, 0xa4, 0x50,
	0x49, 0xaf, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0x1a, 0x42, 0x6a, 0xc8, 0xde, 0x87, 0x4a, 0x16, 0xeb,
	0x57, 0xa3, 0x39, 0xbf, 0xd6, 0x29, 0x8c, 0x4b, 0x1c, 0xc5, 0xc8, 0xc4, 0x82, 0x4c, 0xa8, 0xb1,
	0x51, 0xd6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x10, 0x35, 0xba,
	0x90, 0x00, 0x4b, 0xcf, 0xd7, 0x4d, 0x4f, 0xcd, 0x03, 0x87, 0x95, 0x3e, 0x96, 0x10, 0xb7, 0xce,
	0x4c, 0xcc, 0x4d, 0x62, 0x03, 0x4b, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x42, 0x6c, 0x28,
	0xfe, 0xf6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserAccountServiceClient is the client API for UserAccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserAccountServiceClient interface {
	// Returns the specified UserAccount resource.
	Get(ctx context.Context, in *GetUserAccountRequest, opts ...grpc.CallOption) (*UserAccount, error)
}

type userAccountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAccountServiceClient(cc grpc.ClientConnInterface) UserAccountServiceClient {
	return &userAccountServiceClient{cc}
}

func (c *userAccountServiceClient) Get(ctx context.Context, in *GetUserAccountRequest, opts ...grpc.CallOption) (*UserAccount, error) {
	out := new(UserAccount)
	err := c.cc.Invoke(ctx, "/yandex.cloud.iam.v1.UserAccountService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAccountServiceServer is the server API for UserAccountService service.
type UserAccountServiceServer interface {
	// Returns the specified UserAccount resource.
	Get(context.Context, *GetUserAccountRequest) (*UserAccount, error)
}

// UnimplementedUserAccountServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserAccountServiceServer struct {
}

func (*UnimplementedUserAccountServiceServer) Get(ctx context.Context, req *GetUserAccountRequest) (*UserAccount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterUserAccountServiceServer(s *grpc.Server, srv UserAccountServiceServer) {
	s.RegisterService(&_UserAccountService_serviceDesc, srv)
}

func _UserAccountService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAccountServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.iam.v1.UserAccountService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAccountServiceServer).Get(ctx, req.(*GetUserAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserAccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.iam.v1.UserAccountService",
	HandlerType: (*UserAccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _UserAccountService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/iam/v1/user_account_service.proto",
}
