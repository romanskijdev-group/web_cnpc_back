// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: service/UserService.proto

package protoobj

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_UserService_proto protoreflect.FileDescriptor

var file_service_UserService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67,
	0x1a, 0x17, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x4d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x4c, 0x6f,
	0x67, 0x49, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x24, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x55, 0x73, 0x65, 0x72, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe9, 0x04, 0x0a, 0x17, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x43, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0d, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x73, 0x67, 0x1a, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10,
	0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71,
	0x1a, 0x11, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x73, 0x67, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0d, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x73, 0x67, 0x1a, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x6d, 0x73,
	0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x73, 0x67, 0x1a, 0x0d, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x73, 0x67, 0x1a, 0x0a, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x19, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43,
	0x6f, 0x64, 0x65, 0x42, 0x6f, 0x74, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x12, 0x1a, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x11, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x52, 0x4c, 0x12, 0x1b, 0x2e, 0x6d, 0x73,
	0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x15, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x12, 0x11, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x15, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x6f, 0x62, 0x6a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_UserService_proto_goTypes = []interface{}{
	(*UserAuthReqAccountReq)(nil),  // 0: msg.UserAuthReqAccountReq
	(*UsersMsg)(nil),               // 1: msg.UsersMsg
	(*UsersMsgReq)(nil),            // 2: msg.UsersMsgReq
	(*UpdateUserAvatarURLReq)(nil), // 3: msg.UpdateUserAvatarURLReq
	(*UserAlertMsg)(nil),           // 4: msg.UserAlertMsg
	(*LogInInfoRes)(nil),           // 5: msg.LogInInfoRes
	(*UsersMsgList)(nil),           // 6: msg.UsersMsgList
	(*Empty)(nil),                  // 7: msg.Empty
	(*UserAlertMsgList)(nil),       // 8: msg.UserAlertMsgList
}
var file_service_UserService_proto_depIdxs = []int32{
	0,  // 0: msg.UserAccountServiceProto.UserLoginAccount:input_type -> msg.UserAuthReqAccountReq
	1,  // 1: msg.UserAccountServiceProto.GetUserProfile:input_type -> msg.UsersMsg
	2,  // 2: msg.UserAccountServiceProto.GetUsersInfoList:input_type -> msg.UsersMsgReq
	1,  // 3: msg.UserAccountServiceProto.UpdateUserProfile:input_type -> msg.UsersMsg
	1,  // 4: msg.UserAccountServiceProto.CreateNewUser:input_type -> msg.UsersMsg
	1,  // 5: msg.UserAccountServiceProto.DeleteUser:input_type -> msg.UsersMsg
	0,  // 6: msg.UserAccountServiceProto.CheckCodeBotAuthUserValid:input_type -> msg.UserAuthReqAccountReq
	3,  // 7: msg.UserAccountServiceProto.UpdateUserAvatarURL:input_type -> msg.UpdateUserAvatarURLReq
	4,  // 8: msg.UserAccountServiceProto.GetUserAlerts:input_type -> msg.UserAlertMsg
	4,  // 9: msg.UserAccountServiceProto.UpdateUserAlerts:input_type -> msg.UserAlertMsg
	5,  // 10: msg.UserAccountServiceProto.UserLoginAccount:output_type -> msg.LogInInfoRes
	1,  // 11: msg.UserAccountServiceProto.GetUserProfile:output_type -> msg.UsersMsg
	6,  // 12: msg.UserAccountServiceProto.GetUsersInfoList:output_type -> msg.UsersMsgList
	1,  // 13: msg.UserAccountServiceProto.UpdateUserProfile:output_type -> msg.UsersMsg
	1,  // 14: msg.UserAccountServiceProto.CreateNewUser:output_type -> msg.UsersMsg
	7,  // 15: msg.UserAccountServiceProto.DeleteUser:output_type -> msg.Empty
	5,  // 16: msg.UserAccountServiceProto.CheckCodeBotAuthUserValid:output_type -> msg.LogInInfoRes
	7,  // 17: msg.UserAccountServiceProto.UpdateUserAvatarURL:output_type -> msg.Empty
	8,  // 18: msg.UserAccountServiceProto.GetUserAlerts:output_type -> msg.UserAlertMsgList
	8,  // 19: msg.UserAccountServiceProto.UpdateUserAlerts:output_type -> msg.UserAlertMsgList
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_UserService_proto_init() }
func file_service_UserService_proto_init() {
	if File_service_UserService_proto != nil {
		return
	}
	file_messages_UsersMsg_proto_init()
	file_messages_UsersMsgList_proto_init()
	file_messages_UpdateUserAvatarURLReq_proto_init()
	file_messages_Empty_proto_init()
	file_messages_LogInInfoRes_proto_init()
	file_messages_UserAuthReqAccountReq_proto_init()
	file_messages_UserSystemAlerts_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_UserService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_UserService_proto_goTypes,
		DependencyIndexes: file_service_UserService_proto_depIdxs,
	}.Build()
	File_service_UserService_proto = out.File
	file_service_UserService_proto_rawDesc = nil
	file_service_UserService_proto_goTypes = nil
	file_service_UserService_proto_depIdxs = nil
}
