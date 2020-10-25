// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: proto/src/command/create.proto

package command

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateProtoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateProtoRequest) Reset() {
	*x = CreateProtoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_src_command_create_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProtoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProtoRequest) ProtoMessage() {}

func (x *CreateProtoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_src_command_create_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProtoRequest.ProtoReflect.Descriptor instead.
func (*CreateProtoRequest) Descriptor() ([]byte, []int) {
	return file_proto_src_command_create_proto_rawDescGZIP(), []int{0}
}

type CreateProtoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateProtoResponse) Reset() {
	*x = CreateProtoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_src_command_create_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProtoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProtoResponse) ProtoMessage() {}

func (x *CreateProtoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_src_command_create_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProtoResponse.ProtoReflect.Descriptor instead.
func (*CreateProtoResponse) Descriptor() ([]byte, []int) {
	return file_proto_src_command_create_proto_rawDescGZIP(), []int{1}
}

var File_proto_src_command_create_proto protoreflect.FileDescriptor

var file_proto_src_command_create_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22,
	0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6c, 0x61, 0x75,
	0x32, 0x33, 0x32, 0x38, 0x2f, 0x67, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_src_command_create_proto_rawDescOnce sync.Once
	file_proto_src_command_create_proto_rawDescData = file_proto_src_command_create_proto_rawDesc
)

func file_proto_src_command_create_proto_rawDescGZIP() []byte {
	file_proto_src_command_create_proto_rawDescOnce.Do(func() {
		file_proto_src_command_create_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_src_command_create_proto_rawDescData)
	})
	return file_proto_src_command_create_proto_rawDescData
}

var file_proto_src_command_create_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_src_command_create_proto_goTypes = []interface{}{
	(*CreateProtoRequest)(nil),  // 0: proto.command.CreateProtoRequest
	(*CreateProtoResponse)(nil), // 1: proto.command.CreateProtoResponse
}
var file_proto_src_command_create_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_src_command_create_proto_init() }
func file_proto_src_command_create_proto_init() {
	if File_proto_src_command_create_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_src_command_create_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProtoRequest); i {
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
		file_proto_src_command_create_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProtoResponse); i {
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
			RawDescriptor: file_proto_src_command_create_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_src_command_create_proto_goTypes,
		DependencyIndexes: file_proto_src_command_create_proto_depIdxs,
		MessageInfos:      file_proto_src_command_create_proto_msgTypes,
	}.Build()
	File_proto_src_command_create_proto = out.File
	file_proto_src_command_create_proto_rawDesc = nil
	file_proto_src_command_create_proto_goTypes = nil
	file_proto_src_command_create_proto_depIdxs = nil
}
