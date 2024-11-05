// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: messages/UserAuthReqAccountReq.proto

package protoobj

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserAuthReqAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email             *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	TemporaryPassword *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=TemporaryPassword,proto3" json:"TemporaryPassword,omitempty"`
	TelegramID        *wrapperspb.Int64Value  `protobuf:"bytes,3,opt,name=TelegramID,proto3" json:"TelegramID,omitempty"`
	EmailCode         *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=EmailCode,proto3" json:"EmailCode,omitempty"`
	SystemID          *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=SystemID,proto3" json:"SystemID,omitempty"`
	Code              *wrapperspb.StringValue `protobuf:"bytes,7,opt,name=Code,proto3" json:"Code,omitempty"`
	Secret            *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=Secret,proto3" json:"Secret,omitempty"`
	AuthType          TypeAuth                `protobuf:"varint,10,opt,name=AuthType,proto3,enum=msg.TypeAuth" json:"AuthType,omitempty"`
	DetectorIPStruct  *DetectorIPStruct       `protobuf:"bytes,12,opt,name=DetectorIPStruct,proto3" json:"DetectorIPStruct,omitempty"`
}

func (x *UserAuthReqAccountReq) Reset() {
	*x = UserAuthReqAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_UserAuthReqAccountReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAuthReqAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAuthReqAccountReq) ProtoMessage() {}

func (x *UserAuthReqAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_messages_UserAuthReqAccountReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAuthReqAccountReq.ProtoReflect.Descriptor instead.
func (*UserAuthReqAccountReq) Descriptor() ([]byte, []int) {
	return file_messages_UserAuthReqAccountReq_proto_rawDescGZIP(), []int{0}
}

func (x *UserAuthReqAccountReq) GetEmail() *wrapperspb.StringValue {
	if x != nil {
		return x.Email
	}
	return nil
}

func (x *UserAuthReqAccountReq) GetTemporaryPassword() *wrapperspb.StringValue {
	if x != nil {
		return x.TemporaryPassword
	}
	return nil
}

func (x *UserAuthReqAccountReq) GetTelegramID() *wrapperspb.Int64Value {
	if x != nil {
		return x.TelegramID
	}
	return nil
}

func (x *UserAuthReqAccountReq) GetEmailCode() *wrapperspb.StringValue {
	if x != nil {
		return x.EmailCode
	}
	return nil
}

func (x *UserAuthReqAccountReq) GetSystemID() *wrapperspb.StringValue {
	if x != nil {
		return x.SystemID
	}
	return nil
}

func (x *UserAuthReqAccountReq) GetCode() *wrapperspb.StringValue {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *UserAuthReqAccountReq) GetSecret() *wrapperspb.StringValue {
	if x != nil {
		return x.Secret
	}
	return nil
}

func (x *UserAuthReqAccountReq) GetAuthType() TypeAuth {
	if x != nil {
		return x.AuthType
	}
	return TypeAuth_TypeAuth_NULL
}

func (x *UserAuthReqAccountReq) GetDetectorIPStruct() *DetectorIPStruct {
	if x != nil {
		return x.DetectorIPStruct
	}
	return nil
}

var File_messages_UserAuthReqAccountReq_proto protoreflect.FileDescriptor

var file_messages_UserAuthReqAccountReq_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2f, 0x54, 0x79, 0x70, 0x65, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x44, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x49, 0x50, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa0, 0x04, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x32, 0x0a, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x4a, 0x0a, 0x11, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x54, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x72, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x3b, 0x0a, 0x0a,
	0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x54,
	0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x09, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49,
	0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x44, 0x12,
	0x30, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x34, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x29, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x08, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x41, 0x0a, 0x10, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x50,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x50, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x10, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x50, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6f, 0x62, 0x6a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_messages_UserAuthReqAccountReq_proto_rawDescOnce sync.Once
	file_messages_UserAuthReqAccountReq_proto_rawDescData = file_messages_UserAuthReqAccountReq_proto_rawDesc
)

func file_messages_UserAuthReqAccountReq_proto_rawDescGZIP() []byte {
	file_messages_UserAuthReqAccountReq_proto_rawDescOnce.Do(func() {
		file_messages_UserAuthReqAccountReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_UserAuthReqAccountReq_proto_rawDescData)
	})
	return file_messages_UserAuthReqAccountReq_proto_rawDescData
}

var file_messages_UserAuthReqAccountReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_messages_UserAuthReqAccountReq_proto_goTypes = []interface{}{
	(*UserAuthReqAccountReq)(nil),  // 0: msg.UserAuthReqAccountReq
	(*wrapperspb.StringValue)(nil), // 1: google.protobuf.StringValue
	(*wrapperspb.Int64Value)(nil),  // 2: google.protobuf.Int64Value
	(TypeAuth)(0),                  // 3: msg.TypeAuth
	(*DetectorIPStruct)(nil),       // 4: msg.DetectorIPStruct
}
var file_messages_UserAuthReqAccountReq_proto_depIdxs = []int32{
	1, // 0: msg.UserAuthReqAccountReq.Email:type_name -> google.protobuf.StringValue
	1, // 1: msg.UserAuthReqAccountReq.TemporaryPassword:type_name -> google.protobuf.StringValue
	2, // 2: msg.UserAuthReqAccountReq.TelegramID:type_name -> google.protobuf.Int64Value
	1, // 3: msg.UserAuthReqAccountReq.EmailCode:type_name -> google.protobuf.StringValue
	1, // 4: msg.UserAuthReqAccountReq.SystemID:type_name -> google.protobuf.StringValue
	1, // 5: msg.UserAuthReqAccountReq.Code:type_name -> google.protobuf.StringValue
	1, // 6: msg.UserAuthReqAccountReq.Secret:type_name -> google.protobuf.StringValue
	3, // 7: msg.UserAuthReqAccountReq.AuthType:type_name -> msg.TypeAuth
	4, // 8: msg.UserAuthReqAccountReq.DetectorIPStruct:type_name -> msg.DetectorIPStruct
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_messages_UserAuthReqAccountReq_proto_init() }
func file_messages_UserAuthReqAccountReq_proto_init() {
	if File_messages_UserAuthReqAccountReq_proto != nil {
		return
	}
	file_enums_TypeAuth_proto_init()
	file_messages_DetectorIPStruct_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_messages_UserAuthReqAccountReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAuthReqAccountReq); i {
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
			RawDescriptor: file_messages_UserAuthReqAccountReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_UserAuthReqAccountReq_proto_goTypes,
		DependencyIndexes: file_messages_UserAuthReqAccountReq_proto_depIdxs,
		MessageInfos:      file_messages_UserAuthReqAccountReq_proto_msgTypes,
	}.Build()
	File_messages_UserAuthReqAccountReq_proto = out.File
	file_messages_UserAuthReqAccountReq_proto_rawDesc = nil
	file_messages_UserAuthReqAccountReq_proto_goTypes = nil
	file_messages_UserAuthReqAccountReq_proto_depIdxs = nil
}