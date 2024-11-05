// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: enums/NotifyCategory.proto

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

// Типы исполнений предложений p2p
type NotifyCategory int32

const (
	NotifyCategory_NotifyCategory_NULL               NotifyCategory = 0
	NotifyCategory_NotifyCategory_bearer_cheque      NotifyCategory = 1 //  чек
	NotifyCategory_NotifyCategory_chats              NotifyCategory = 2 // Чаты
	NotifyCategory_NotifyCategory_info               NotifyCategory = 3 // Информационное уведомление(от админа)
	NotifyCategory_NotifyCategory_temporary_password NotifyCategory = 4 // Временный пароль
	NotifyCategory_NotifyCategory_device_new         NotifyCategory = 5 // Новое устройство
)

// Enum value maps for NotifyCategory.
var (
	NotifyCategory_name = map[int32]string{
		0: "NotifyCategory_NULL",
		1: "NotifyCategory_bearer_cheque",
		2: "NotifyCategory_chats",
		3: "NotifyCategory_info",
		4: "NotifyCategory_temporary_password",
		5: "NotifyCategory_device_new",
	}
	NotifyCategory_value = map[string]int32{
		"NotifyCategory_NULL":               0,
		"NotifyCategory_bearer_cheque":      1,
		"NotifyCategory_chats":              2,
		"NotifyCategory_info":               3,
		"NotifyCategory_temporary_password": 4,
		"NotifyCategory_device_new":         5,
	}
)

func (x NotifyCategory) Enum() *NotifyCategory {
	p := new(NotifyCategory)
	*p = x
	return p
}

func (x NotifyCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotifyCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_NotifyCategory_proto_enumTypes[0].Descriptor()
}

func (NotifyCategory) Type() protoreflect.EnumType {
	return &file_enums_NotifyCategory_proto_enumTypes[0]
}

func (x NotifyCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotifyCategory.Descriptor instead.
func (NotifyCategory) EnumDescriptor() ([]byte, []int) {
	return file_enums_NotifyCategory_proto_rawDescGZIP(), []int{0}
}

var File_enums_NotifyCategory_proto protoreflect.FileDescriptor

var file_enums_NotifyCategory_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73,
	0x67, 0x2a, 0xc4, 0x01, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x20, 0x0a,
	0x1c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
	0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x65, 0x71, 0x75, 0x65, 0x10, 0x01, 0x12,
	0x18, 0x0a, 0x14, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x73, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x10, 0x03, 0x12, 0x25, 0x0a, 0x21, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x65, 0x77, 0x10, 0x05, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6f, 0x62, 0x6a, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_NotifyCategory_proto_rawDescOnce sync.Once
	file_enums_NotifyCategory_proto_rawDescData = file_enums_NotifyCategory_proto_rawDesc
)

func file_enums_NotifyCategory_proto_rawDescGZIP() []byte {
	file_enums_NotifyCategory_proto_rawDescOnce.Do(func() {
		file_enums_NotifyCategory_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_NotifyCategory_proto_rawDescData)
	})
	return file_enums_NotifyCategory_proto_rawDescData
}

var file_enums_NotifyCategory_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_NotifyCategory_proto_goTypes = []interface{}{
	(NotifyCategory)(0), // 0: msg.NotifyCategory
}
var file_enums_NotifyCategory_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_NotifyCategory_proto_init() }
func file_enums_NotifyCategory_proto_init() {
	if File_enums_NotifyCategory_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enums_NotifyCategory_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_NotifyCategory_proto_goTypes,
		DependencyIndexes: file_enums_NotifyCategory_proto_depIdxs,
		EnumInfos:         file_enums_NotifyCategory_proto_enumTypes,
	}.Build()
	File_enums_NotifyCategory_proto = out.File
	file_enums_NotifyCategory_proto_rawDesc = nil
	file_enums_NotifyCategory_proto_goTypes = nil
	file_enums_NotifyCategory_proto_depIdxs = nil
}