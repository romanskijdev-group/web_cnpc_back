// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: messages/FilteringParamsList.proto

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

type FilteringParamsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset         *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit          *wrapperspb.UInt64Value `protobuf:"bytes,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	LikeFields     map[string]string       `protobuf:"bytes,3,rep,name=LikeFields,proto3" json:"LikeFields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OrSearchFields map[string]bool         `protobuf:"bytes,4,rep,name=OrSearchFields,proto3" json:"OrSearchFields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *FilteringParamsList) Reset() {
	*x = FilteringParamsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_FilteringParamsList_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilteringParamsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilteringParamsList) ProtoMessage() {}

func (x *FilteringParamsList) ProtoReflect() protoreflect.Message {
	mi := &file_messages_FilteringParamsList_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilteringParamsList.ProtoReflect.Descriptor instead.
func (*FilteringParamsList) Descriptor() ([]byte, []int) {
	return file_messages_FilteringParamsList_proto_rawDescGZIP(), []int{0}
}

func (x *FilteringParamsList) GetOffset() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Offset
	}
	return nil
}

func (x *FilteringParamsList) GetLimit() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Limit
	}
	return nil
}

func (x *FilteringParamsList) GetLikeFields() map[string]string {
	if x != nil {
		return x.LikeFields
	}
	return nil
}

func (x *FilteringParamsList) GetOrSearchFields() map[string]bool {
	if x != nil {
		return x.OrSearchFields
	}
	return nil
}

var File_messages_FilteringParamsList_proto protoreflect.FileDescriptor

var file_messages_FilteringParamsList_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x03, 0x0a, 0x13, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x34, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x48, 0x0a, 0x0a, 0x4c,
	0x69, 0x6b, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x4c, 0x69, 0x6b, 0x65, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x54, 0x0a, 0x0e, 0x4f, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x6d, 0x73, 0x67, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x4f, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x4f, 0x72, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x4c,
	0x69, 0x6b, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x41, 0x0a, 0x13, 0x4f, 0x72,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x12, 0x5a,
	0x10, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6f, 0x62,
	0x6a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_FilteringParamsList_proto_rawDescOnce sync.Once
	file_messages_FilteringParamsList_proto_rawDescData = file_messages_FilteringParamsList_proto_rawDesc
)

func file_messages_FilteringParamsList_proto_rawDescGZIP() []byte {
	file_messages_FilteringParamsList_proto_rawDescOnce.Do(func() {
		file_messages_FilteringParamsList_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_FilteringParamsList_proto_rawDescData)
	})
	return file_messages_FilteringParamsList_proto_rawDescData
}

var file_messages_FilteringParamsList_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_messages_FilteringParamsList_proto_goTypes = []interface{}{
	(*FilteringParamsList)(nil),    // 0: msg.FilteringParamsList
	nil,                            // 1: msg.FilteringParamsList.LikeFieldsEntry
	nil,                            // 2: msg.FilteringParamsList.OrSearchFieldsEntry
	(*wrapperspb.UInt64Value)(nil), // 3: google.protobuf.UInt64Value
}
var file_messages_FilteringParamsList_proto_depIdxs = []int32{
	3, // 0: msg.FilteringParamsList.Offset:type_name -> google.protobuf.UInt64Value
	3, // 1: msg.FilteringParamsList.Limit:type_name -> google.protobuf.UInt64Value
	1, // 2: msg.FilteringParamsList.LikeFields:type_name -> msg.FilteringParamsList.LikeFieldsEntry
	2, // 3: msg.FilteringParamsList.OrSearchFields:type_name -> msg.FilteringParamsList.OrSearchFieldsEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_messages_FilteringParamsList_proto_init() }
func file_messages_FilteringParamsList_proto_init() {
	if File_messages_FilteringParamsList_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_FilteringParamsList_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilteringParamsList); i {
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
			RawDescriptor: file_messages_FilteringParamsList_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_FilteringParamsList_proto_goTypes,
		DependencyIndexes: file_messages_FilteringParamsList_proto_depIdxs,
		MessageInfos:      file_messages_FilteringParamsList_proto_msgTypes,
	}.Build()
	File_messages_FilteringParamsList_proto = out.File
	file_messages_FilteringParamsList_proto_rawDesc = nil
	file_messages_FilteringParamsList_proto_goTypes = nil
	file_messages_FilteringParamsList_proto_depIdxs = nil
}