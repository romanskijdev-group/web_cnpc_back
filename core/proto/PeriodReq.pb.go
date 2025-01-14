// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: messages/PeriodReq.proto

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

type PeriodReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatesFrom *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=datesFrom,proto3" json:"datesFrom,omitempty"`
	DatesTo   *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=datesTo,proto3" json:"datesTo,omitempty"`
}

func (x *PeriodReq) Reset() {
	*x = PeriodReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_PeriodReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeriodReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeriodReq) ProtoMessage() {}

func (x *PeriodReq) ProtoReflect() protoreflect.Message {
	mi := &file_messages_PeriodReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeriodReq.ProtoReflect.Descriptor instead.
func (*PeriodReq) Descriptor() ([]byte, []int) {
	return file_messages_PeriodReq_proto_rawDescGZIP(), []int{0}
}

func (x *PeriodReq) GetDatesFrom() *wrapperspb.Int64Value {
	if x != nil {
		return x.DatesFrom
	}
	return nil
}

func (x *PeriodReq) GetDatesTo() *wrapperspb.Int64Value {
	if x != nil {
		return x.DatesTo
	}
	return nil
}

var File_messages_PeriodReq_proto protoreflect.FileDescriptor

var file_messages_PeriodReq_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7d, 0x0a, 0x09, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x12, 0x39, 0x0a, 0x09,
	0x64, 0x61, 0x74, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x35, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x54, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x64, 0x61, 0x74, 0x65, 0x73, 0x54, 0x6f, 0x42, 0x12,
	0x5a, 0x10, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6f,
	0x62, 0x6a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_PeriodReq_proto_rawDescOnce sync.Once
	file_messages_PeriodReq_proto_rawDescData = file_messages_PeriodReq_proto_rawDesc
)

func file_messages_PeriodReq_proto_rawDescGZIP() []byte {
	file_messages_PeriodReq_proto_rawDescOnce.Do(func() {
		file_messages_PeriodReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_PeriodReq_proto_rawDescData)
	})
	return file_messages_PeriodReq_proto_rawDescData
}

var file_messages_PeriodReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_messages_PeriodReq_proto_goTypes = []interface{}{
	(*PeriodReq)(nil),             // 0: msg.PeriodReq
	(*wrapperspb.Int64Value)(nil), // 1: google.protobuf.Int64Value
}
var file_messages_PeriodReq_proto_depIdxs = []int32{
	1, // 0: msg.PeriodReq.datesFrom:type_name -> google.protobuf.Int64Value
	1, // 1: msg.PeriodReq.datesTo:type_name -> google.protobuf.Int64Value
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_messages_PeriodReq_proto_init() }
func file_messages_PeriodReq_proto_init() {
	if File_messages_PeriodReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_PeriodReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeriodReq); i {
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
			RawDescriptor: file_messages_PeriodReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_PeriodReq_proto_goTypes,
		DependencyIndexes: file_messages_PeriodReq_proto_depIdxs,
		MessageInfos:      file_messages_PeriodReq_proto_msgTypes,
	}.Build()
	File_messages_PeriodReq_proto = out.File
	file_messages_PeriodReq_proto_rawDesc = nil
	file_messages_PeriodReq_proto_goTypes = nil
	file_messages_PeriodReq_proto_depIdxs = nil
}
