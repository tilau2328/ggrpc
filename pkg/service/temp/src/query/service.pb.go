// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: temp/src/query/service.proto

package query

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_temp_src_query_service_proto protoreflect.FileDescriptor

var file_temp_src_query_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x74, 0x65, 0x6d, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x74, 0x65, 0x6d, 0x70, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x19, 0x74, 0x65, 0x6d, 0x70,
	0x2f, 0x73, 0x72, 0x63, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x74, 0x65, 0x6d, 0x70, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0x9a, 0x01, 0x0a, 0x12, 0x54, 0x65, 0x6d, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x66, 0x69, 0x6e, 0x64, 0x12,
	0x1b, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x54, 0x65, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x65,
	0x6d, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x1b, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x65, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6c, 0x61,
	0x75, 0x32, 0x33, 0x32, 0x38, 0x2f, 0x67, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_temp_src_query_service_proto_goTypes = []interface{}{
	(*FindTempRequest)(nil),  // 0: temp.query.FindTempRequest
	(*ListTempRequest)(nil),  // 1: temp.query.ListTempRequest
	(*FindTempResponse)(nil), // 2: temp.query.FindTempResponse
	(*ListTempResponse)(nil), // 3: temp.query.ListTempResponse
}
var file_temp_src_query_service_proto_depIdxs = []int32{
	0, // 0: temp.query.TempCommandService.find:input_type -> temp.query.FindTempRequest
	1, // 1: temp.query.TempCommandService.list:input_type -> temp.query.ListTempRequest
	2, // 2: temp.query.TempCommandService.find:output_type -> temp.query.FindTempResponse
	3, // 3: temp.query.TempCommandService.list:output_type -> temp.query.ListTempResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temp_src_query_service_proto_init() }
func file_temp_src_query_service_proto_init() {
	if File_temp_src_query_service_proto != nil {
		return
	}
	file_temp_src_query_find_proto_init()
	file_temp_src_query_list_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temp_src_query_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_temp_src_query_service_proto_goTypes,
		DependencyIndexes: file_temp_src_query_service_proto_depIdxs,
	}.Build()
	File_temp_src_query_service_proto = out.File
	file_temp_src_query_service_proto_rawDesc = nil
	file_temp_src_query_service_proto_goTypes = nil
	file_temp_src_query_service_proto_depIdxs = nil
}
