// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: messages/UsersMsgList.proto

package protoobj

import (
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

type UsersMsgList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsersMsg []*UsersMsg `protobuf:"bytes,1,rep,name=usersMsg,proto3" json:"usersMsg,omitempty"`
}

func (x *UsersMsgList) Reset() {
	*x = UsersMsgList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_UsersMsgList_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersMsgList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersMsgList) ProtoMessage() {}

func (x *UsersMsgList) ProtoReflect() protoreflect.Message {
	mi := &file_messages_UsersMsgList_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersMsgList.ProtoReflect.Descriptor instead.
func (*UsersMsgList) Descriptor() ([]byte, []int) {
	return file_messages_UsersMsgList_proto_rawDescGZIP(), []int{0}
}

func (x *UsersMsgList) GetUsersMsg() []*UsersMsg {
	if x != nil {
		return x.UsersMsg
	}
	return nil
}

var File_messages_UsersMsgList_proto protoreflect.FileDescriptor

var file_messages_UsersMsgList_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d,
	0x73, 0x67, 0x1a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x4d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0c, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x6d, 0x73, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x73, 0x67, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x4d, 0x73, 0x67, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6f, 0x62, 0x6a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_messages_UsersMsgList_proto_rawDescOnce sync.Once
	file_messages_UsersMsgList_proto_rawDescData = file_messages_UsersMsgList_proto_rawDesc
)

func file_messages_UsersMsgList_proto_rawDescGZIP() []byte {
	file_messages_UsersMsgList_proto_rawDescOnce.Do(func() {
		file_messages_UsersMsgList_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_UsersMsgList_proto_rawDescData)
	})
	return file_messages_UsersMsgList_proto_rawDescData
}

var file_messages_UsersMsgList_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_messages_UsersMsgList_proto_goTypes = []interface{}{
	(*UsersMsgList)(nil), // 0: msg.UsersMsgList
	(*UsersMsg)(nil),     // 1: msg.UsersMsg
}
var file_messages_UsersMsgList_proto_depIdxs = []int32{
	1, // 0: msg.UsersMsgList.usersMsg:type_name -> msg.UsersMsg
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_messages_UsersMsgList_proto_init() }
func file_messages_UsersMsgList_proto_init() {
	if File_messages_UsersMsgList_proto != nil {
		return
	}
	file_messages_UsersMsg_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_messages_UsersMsgList_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersMsgList); i {
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
			RawDescriptor: file_messages_UsersMsgList_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_UsersMsgList_proto_goTypes,
		DependencyIndexes: file_messages_UsersMsgList_proto_depIdxs,
		MessageInfos:      file_messages_UsersMsgList_proto_msgTypes,
	}.Build()
	File_messages_UsersMsgList_proto = out.File
	file_messages_UsersMsgList_proto_rawDesc = nil
	file_messages_UsersMsgList_proto_goTypes = nil
	file_messages_UsersMsgList_proto_depIdxs = nil
}
