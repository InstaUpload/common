// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: api/user.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_api_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Msg           string                 `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_api_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type LoginUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginUserRequest) Reset() {
	*x = LoginUserRequest{}
	mi := &file_api_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserRequest) ProtoMessage() {}

func (x *LoginUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserRequest.ProtoReflect.Descriptor instead.
func (*LoginUserRequest) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{2}
}

func (x *LoginUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginUserResponse) Reset() {
	*x = LoginUserResponse{}
	mi := &file_api_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserResponse) ProtoMessage() {}

func (x *LoginUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserResponse.ProtoReflect.Descriptor instead.
func (*LoginUserResponse) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{3}
}

func (x *LoginUserResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthUserRequest) Reset() {
	*x = AuthUserRequest{}
	mi := &file_api_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserRequest) ProtoMessage() {}

func (x *AuthUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserRequest.ProtoReflect.Descriptor instead.
func (*AuthUserRequest) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{4}
}

func (x *AuthUserRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AuthUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	IsVerified    bool                   `protobuf:"varint,5,opt,name=is_verified,json=isVerified,proto3" json:"is_verified,omitempty"`
	CreatedOn     string                 `protobuf:"bytes,6,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthUserResponse) Reset() {
	*x = AuthUserResponse{}
	mi := &file_api_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthUserResponse) ProtoMessage() {}

func (x *AuthUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthUserResponse.ProtoReflect.Descriptor instead.
func (*AuthUserResponse) Descriptor() ([]byte, []int) {
	return file_api_user_proto_rawDescGZIP(), []int{5}
}

func (x *AuthUserResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AuthUserResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthUserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthUserResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AuthUserResponse) GetIsVerified() bool {
	if x != nil {
		return x.IsVerified
	}
	return false
}

func (x *AuthUserResponse) GetCreatedOn() string {
	if x != nil {
		return x.CreatedOn
	}
	return ""
}

var File_api_user_proto protoreflect.FileDescriptor

const file_api_user_proto_rawDesc = "" +
	"\n" +
	"\x0eapi/user.proto\x12\x03api\"Y\n" +
	"\x11CreateUserRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\"&\n" +
	"\x12CreateUserResponse\x12\x10\n" +
	"\x03msg\x18\x01 \x01(\tR\x03msg\"D\n" +
	"\x10LoginUserRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\")\n" +
	"\x11LoginUserResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"'\n" +
	"\x0fAuthUserRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"\xa8\x01\n" +
	"\x10AuthUserResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x04 \x01(\tR\bpassword\x12\x1f\n" +
	"\vis_verified\x18\x05 \x01(\bR\n" +
	"isVerified\x12\x1d\n" +
	"\n" +
	"created_on\x18\x06 \x01(\tR\tcreatedOn2\xc1\x01\n" +
	"\vUserService\x12=\n" +
	"\n" +
	"CreateUser\x12\x16.api.CreateUserRequest\x1a\x17.api.CreateUserResponse\x12:\n" +
	"\tLoginUser\x12\x15.api.LoginUserRequest\x1a\x16.api.LoginUserResponse\x127\n" +
	"\bAuthUser\x12\x14.api.AuthUserRequest\x1a\x15.api.AuthUserResponseB#Z!github.com/InstaUpload/common/apib\x06proto3"

var (
	file_api_user_proto_rawDescOnce sync.Once
	file_api_user_proto_rawDescData []byte
)

func file_api_user_proto_rawDescGZIP() []byte {
	file_api_user_proto_rawDescOnce.Do(func() {
		file_api_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_user_proto_rawDesc), len(file_api_user_proto_rawDesc)))
	})
	return file_api_user_proto_rawDescData
}

var file_api_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_user_proto_goTypes = []any{
	(*CreateUserRequest)(nil),  // 0: api.CreateUserRequest
	(*CreateUserResponse)(nil), // 1: api.CreateUserResponse
	(*LoginUserRequest)(nil),   // 2: api.LoginUserRequest
	(*LoginUserResponse)(nil),  // 3: api.LoginUserResponse
	(*AuthUserRequest)(nil),    // 4: api.AuthUserRequest
	(*AuthUserResponse)(nil),   // 5: api.AuthUserResponse
}
var file_api_user_proto_depIdxs = []int32{
	0, // 0: api.UserService.CreateUser:input_type -> api.CreateUserRequest
	2, // 1: api.UserService.LoginUser:input_type -> api.LoginUserRequest
	4, // 2: api.UserService.AuthUser:input_type -> api.AuthUserRequest
	1, // 3: api.UserService.CreateUser:output_type -> api.CreateUserResponse
	3, // 4: api.UserService.LoginUser:output_type -> api.LoginUserResponse
	5, // 5: api.UserService.AuthUser:output_type -> api.AuthUserResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_user_proto_init() }
func file_api_user_proto_init() {
	if File_api_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_user_proto_rawDesc), len(file_api_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_user_proto_goTypes,
		DependencyIndexes: file_api_user_proto_depIdxs,
		MessageInfos:      file_api_user_proto_msgTypes,
	}.Build()
	File_api_user_proto = out.File
	file_api_user_proto_goTypes = nil
	file_api_user_proto_depIdxs = nil
}
