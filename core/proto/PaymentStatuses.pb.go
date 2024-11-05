// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: enums/PaymentStatuses.proto

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

type PaymentStatuses int32

const (
	PaymentStatuses_PaymentStatuses_NULL      PaymentStatuses = 0
	PaymentStatuses_PaymentStatuses_completed PaymentStatuses = 1
	PaymentStatuses_PaymentStatuses_error     PaymentStatuses = 2
)

// Enum value maps for PaymentStatuses.
var (
	PaymentStatuses_name = map[int32]string{
		0: "PaymentStatuses_NULL",
		1: "PaymentStatuses_completed",
		2: "PaymentStatuses_error",
	}
	PaymentStatuses_value = map[string]int32{
		"PaymentStatuses_NULL":      0,
		"PaymentStatuses_completed": 1,
		"PaymentStatuses_error":     2,
	}
)

func (x PaymentStatuses) Enum() *PaymentStatuses {
	p := new(PaymentStatuses)
	*p = x
	return p
}

func (x PaymentStatuses) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentStatuses) Descriptor() protoreflect.EnumDescriptor {
	return file_enums_PaymentStatuses_proto_enumTypes[0].Descriptor()
}

func (PaymentStatuses) Type() protoreflect.EnumType {
	return &file_enums_PaymentStatuses_proto_enumTypes[0]
}

func (x PaymentStatuses) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentStatuses.Descriptor instead.
func (PaymentStatuses) EnumDescriptor() ([]byte, []int) {
	return file_enums_PaymentStatuses_proto_rawDescGZIP(), []int{0}
}

var File_enums_PaymentStatuses_proto protoreflect.FileDescriptor

var file_enums_PaymentStatuses_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d,
	0x73, 0x67, 0x2a, 0x65, 0x0a, 0x0f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12,
	0x1d, 0x0a, 0x19, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x19,
	0x0a, 0x15, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x02, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6f, 0x62, 0x6a, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_enums_PaymentStatuses_proto_rawDescOnce sync.Once
	file_enums_PaymentStatuses_proto_rawDescData = file_enums_PaymentStatuses_proto_rawDesc
)

func file_enums_PaymentStatuses_proto_rawDescGZIP() []byte {
	file_enums_PaymentStatuses_proto_rawDescOnce.Do(func() {
		file_enums_PaymentStatuses_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_PaymentStatuses_proto_rawDescData)
	})
	return file_enums_PaymentStatuses_proto_rawDescData
}

var file_enums_PaymentStatuses_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_enums_PaymentStatuses_proto_goTypes = []interface{}{
	(PaymentStatuses)(0), // 0: msg.PaymentStatuses
}
var file_enums_PaymentStatuses_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_enums_PaymentStatuses_proto_init() }
func file_enums_PaymentStatuses_proto_init() {
	if File_enums_PaymentStatuses_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enums_PaymentStatuses_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_enums_PaymentStatuses_proto_goTypes,
		DependencyIndexes: file_enums_PaymentStatuses_proto_depIdxs,
		EnumInfos:         file_enums_PaymentStatuses_proto_enumTypes,
	}.Build()
	File_enums_PaymentStatuses_proto = out.File
	file_enums_PaymentStatuses_proto_rawDesc = nil
	file_enums_PaymentStatuses_proto_goTypes = nil
	file_enums_PaymentStatuses_proto_depIdxs = nil
}