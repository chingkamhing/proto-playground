// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.3
// source: gateway/gateway_service.proto

package tm2_proto_gateway

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

var File_gateway_gateway_service_proto protoreflect.FileDescriptor

var file_gateway_gateway_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x72, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xe4, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x69, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x1c, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x69, 0x63, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x71, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xd0, 0x02, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x4d, 0x47, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x68,
	0x6b, 0x2f, 0x74, 0x6d, 0x32, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2d, 0x67, 0x6f, 0x3b, 0x74, 0x6d, 0x32, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x92, 0x41, 0x8e, 0x02, 0x12, 0x7b, 0x0a, 0x14,
	0x54, 0x4d, 0x32, 0x20, 0x47, 0x61, 0x6d, 0x65, 0x20, 0x41, 0x50, 0x49, 0x20, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x22, 0x5c, 0x0a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x54, 0x4d, 0x47, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x68, 0x6b, 0x2f, 0x74, 0x6d,
	0x32, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12,
	0x1a, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x74, 0x6d, 0x67,
	0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x1a, 0x15, 0x6b, 0x61, 0x6d,
	0x2e, 0x64, 0x65, 0x76, 0x40, 0x74, 0x6d, 0x67, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x61,
	0x67, 0x65, 0x32, 0x05, 0x30, 0x2e, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x5a, 0x59, 0x0a, 0x57, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x4d, 0x08,
	0x02, 0x12, 0x38, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2c, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x20, 0x62, 0x79, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x3a, 0x20, 0x42, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x20, 0x3c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x3e, 0x1a, 0x0d, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a,
	0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_gateway_gateway_service_proto_goTypes = []interface{}{
	(*GetUserBasicRequest)(nil),   // 0: gateway.GetUserBasicRequest
	(*GetUserBalanceRequest)(nil), // 1: gateway.GetUserBalanceRequest
	(*GetUserBasicReply)(nil),     // 2: gateway.GetUserBasicReply
	(*GetUserBalanceReply)(nil),   // 3: gateway.GetUserBalanceReply
}
var file_gateway_gateway_service_proto_depIdxs = []int32{
	0, // 0: gateway.User.GetUserBasic:input_type -> gateway.GetUserBasicRequest
	1, // 1: gateway.User.GetUserBalance:input_type -> gateway.GetUserBalanceRequest
	2, // 2: gateway.User.GetUserBasic:output_type -> gateway.GetUserBasicReply
	3, // 3: gateway.User.GetUserBalance:output_type -> gateway.GetUserBalanceReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gateway_gateway_service_proto_init() }
func file_gateway_gateway_service_proto_init() {
	if File_gateway_gateway_service_proto != nil {
		return
	}
	file_gateway_gateway_rr_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gateway_gateway_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_gateway_service_proto_goTypes,
		DependencyIndexes: file_gateway_gateway_service_proto_depIdxs,
	}.Build()
	File_gateway_gateway_service_proto = out.File
	file_gateway_gateway_service_proto_rawDesc = nil
	file_gateway_gateway_service_proto_goTypes = nil
	file_gateway_gateway_service_proto_depIdxs = nil
}