// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: user/v1/user_auth.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// 用户合并账号
type UserMergeAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 要合并的手机号
	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	// 用户OpenID
	OpenId string `protobuf:"bytes,2,opt,name=openId,proto3" json:"openId,omitempty"`
	// 应用AppId，不能为空
	AppId string `protobuf:"bytes,3,opt,name=appId,proto3" json:"appId,omitempty"`
}

func (x *UserMergeAccountRequest) Reset() {
	*x = UserMergeAccountRequest{}
	mi := &file_user_v1_user_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserMergeAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMergeAccountRequest) ProtoMessage() {}

func (x *UserMergeAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMergeAccountRequest.ProtoReflect.Descriptor instead.
func (*UserMergeAccountRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_auth_proto_rawDescGZIP(), []int{0}
}

func (x *UserMergeAccountRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserMergeAccountRequest) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *UserMergeAccountRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

// 用户绑定手机号
type UserBingPhoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户标识，不能为空
	OpenId string `protobuf:"bytes,1,opt,name=openId,proto3" json:"openId,omitempty"`
	// 手机号获取凭证，不能为空
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// 应用AppId，不能为空
	AppId string `protobuf:"bytes,3,opt,name=appId,proto3" json:"appId,omitempty"`
}

func (x *UserBingPhoneRequest) Reset() {
	*x = UserBingPhoneRequest{}
	mi := &file_user_v1_user_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserBingPhoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBingPhoneRequest) ProtoMessage() {}

func (x *UserBingPhoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBingPhoneRequest.ProtoReflect.Descriptor instead.
func (*UserBingPhoneRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_auth_proto_rawDescGZIP(), []int{1}
}

func (x *UserBingPhoneRequest) GetOpenId() string {
	if x != nil {
		return x.OpenId
	}
	return ""
}

func (x *UserBingPhoneRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UserBingPhoneRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

// 注册用户必须参数
type UserRegRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 设备编号，不能为空
	Sn string `protobuf:"bytes,1,opt,name=sn,proto3" json:"sn,omitempty"`
	// 临时登录凭证，不能为空
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// 应用AppId，不能为空
	AppId string `protobuf:"bytes,3,opt,name=appId,proto3" json:"appId,omitempty"`
}

func (x *UserRegRequest) Reset() {
	*x = UserRegRequest{}
	mi := &file_user_v1_user_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRegRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegRequest) ProtoMessage() {}

func (x *UserRegRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegRequest.ProtoReflect.Descriptor instead.
func (*UserRegRequest) Descriptor() ([]byte, []int) {
	return file_user_v1_user_auth_proto_rawDescGZIP(), []int{2}
}

func (x *UserRegRequest) GetSn() string {
	if x != nil {
		return x.Sn
	}
	return ""
}

func (x *UserRegRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UserRegRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

var File_user_v1_user_auth_proto protoreflect.FileDescriptor

var file_user_v1_user_auth_proto_rawDesc = []byte{
	0x0a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x17, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x70, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x65,
	0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x14, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x69, 0x6e, 0x67, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x70, 0x70,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x32,
	0xc1, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x53, 0x0a, 0x07,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x65,
	0x67, 0x12, 0x65, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x67, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6e, 0x67, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x62,
	0x69, 0x6e, 0x64, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x79, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72,
	0x4d, 0x65, 0x72, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x65, 0x72, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01,
	0x2a, 0x22, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x4d, 0x65, 0x72, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x42, 0x28, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x17, 0x62, 0x6c, 0x67, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_v1_user_auth_proto_rawDescOnce sync.Once
	file_user_v1_user_auth_proto_rawDescData = file_user_v1_user_auth_proto_rawDesc
)

func file_user_v1_user_auth_proto_rawDescGZIP() []byte {
	file_user_v1_user_auth_proto_rawDescOnce.Do(func() {
		file_user_v1_user_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_v1_user_auth_proto_rawDescData)
	})
	return file_user_v1_user_auth_proto_rawDescData
}

var file_user_v1_user_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_user_v1_user_auth_proto_goTypes = []any{
	(*UserMergeAccountRequest)(nil), // 0: api.user.v1.UserMergeAccountRequest
	(*UserBingPhoneRequest)(nil),    // 1: api.user.v1.UserBingPhoneRequest
	(*UserRegRequest)(nil),          // 2: api.user.v1.UserRegRequest
	(*Reply)(nil),                   // 3: api.user.v1.Reply
}
var file_user_v1_user_auth_proto_depIdxs = []int32{
	2, // 0: api.user.v1.UserAuth.UserReg:input_type -> api.user.v1.UserRegRequest
	1, // 1: api.user.v1.UserAuth.UserBingPhone:input_type -> api.user.v1.UserBingPhoneRequest
	0, // 2: api.user.v1.UserAuth.UserMergeAccount:input_type -> api.user.v1.UserMergeAccountRequest
	3, // 3: api.user.v1.UserAuth.UserReg:output_type -> api.user.v1.Reply
	3, // 4: api.user.v1.UserAuth.UserBingPhone:output_type -> api.user.v1.Reply
	3, // 5: api.user.v1.UserAuth.UserMergeAccount:output_type -> api.user.v1.Reply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_v1_user_auth_proto_init() }
func file_user_v1_user_auth_proto_init() {
	if File_user_v1_user_auth_proto != nil {
		return
	}
	file_user_v1_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_v1_user_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_v1_user_auth_proto_goTypes,
		DependencyIndexes: file_user_v1_user_auth_proto_depIdxs,
		MessageInfos:      file_user_v1_user_auth_proto_msgTypes,
	}.Build()
	File_user_v1_user_auth_proto = out.File
	file_user_v1_user_auth_proto_rawDesc = nil
	file_user_v1_user_auth_proto_goTypes = nil
	file_user_v1_user_auth_proto_depIdxs = nil
}