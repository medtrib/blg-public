// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: sms/v1/sms.proto

package v1

import (
	v1 "blg-ext/api/ext/v1"
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

// 发送短信参数
type SendSmsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 手机号
	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	// 发送类型
	SmsType string `protobuf:"bytes,2,opt,name=sms_type,json=smsType,proto3" json:"sms_type,omitempty"`
}

func (x *SendSmsRequest) Reset() {
	*x = SendSmsRequest{}
	mi := &file_sms_v1_sms_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendSmsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsRequest) ProtoMessage() {}

func (x *SendSmsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sms_v1_sms_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsRequest.ProtoReflect.Descriptor instead.
func (*SendSmsRequest) Descriptor() ([]byte, []int) {
	return file_sms_v1_sms_proto_rawDescGZIP(), []int{0}
}

func (x *SendSmsRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SendSmsRequest) GetSmsType() string {
	if x != nil {
		return x.SmsType
	}
	return ""
}

var File_sms_v1_sms_proto protoreflect.FileDescriptor

var file_sms_v1_sms_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x13,
	0x73, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x41, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6d, 0x73,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6d, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x32, 0x5b, 0x0a, 0x03, 0x53, 0x6d, 0x73, 0x12, 0x54, 0x0a, 0x07, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a,
	0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x53, 0x6d,
	0x73, 0x42, 0x37, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x5a, 0x51,
	0x43, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x62, 0x6b, 0x2d, 0x65, 0x78, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sms_v1_sms_proto_rawDescOnce sync.Once
	file_sms_v1_sms_proto_rawDescData = file_sms_v1_sms_proto_rawDesc
)

func file_sms_v1_sms_proto_rawDescGZIP() []byte {
	file_sms_v1_sms_proto_rawDescOnce.Do(func() {
		file_sms_v1_sms_proto_rawDescData = protoimpl.X.CompressGZIP(file_sms_v1_sms_proto_rawDescData)
	})
	return file_sms_v1_sms_proto_rawDescData
}

var file_sms_v1_sms_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sms_v1_sms_proto_goTypes = []any{
	(*SendSmsRequest)(nil), // 0: api.sms.v1.SendSmsRequest
	(*v1.Reply)(nil),       // 1: api.sms.v1.Reply
}
var file_sms_v1_sms_proto_depIdxs = []int32{
	0, // 0: api.sms.v1.Sms.SendSms:input_type -> api.sms.v1.SendSmsRequest
	1, // 1: api.sms.v1.Sms.SendSms:output_type -> api.sms.v1.Reply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sms_v1_sms_proto_init() }
func file_sms_v1_sms_proto_init() {
	if File_sms_v1_sms_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sms_v1_sms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sms_v1_sms_proto_goTypes,
		DependencyIndexes: file_sms_v1_sms_proto_depIdxs,
		MessageInfos:      file_sms_v1_sms_proto_msgTypes,
	}.Build()
	File_sms_v1_sms_proto = out.File
	file_sms_v1_sms_proto_rawDesc = nil
	file_sms_v1_sms_proto_goTypes = nil
	file_sms_v1_sms_proto_depIdxs = nil
}