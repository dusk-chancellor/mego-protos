// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: resource_service.proto

package resource

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

var File_resource_service_proto protoreflect.FileDescriptor

var file_resource_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x1a, 0x16, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb8, 0x01, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x46, 0x69, 0x6e, 0x64, 0x12,
	0x0f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x2a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2a, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_resource_service_proto_goTypes = []any{
	(*Empty)(nil), // 0: resource.Empty
}
var file_resource_service_proto_depIdxs = []int32{
	0, // 0: resource.Resource.Find:input_type -> resource.Empty
	0, // 1: resource.Resource.Create:input_type -> resource.Empty
	0, // 2: resource.Resource.Update:input_type -> resource.Empty
	0, // 3: resource.Resource.Delete:input_type -> resource.Empty
	0, // 4: resource.Resource.Find:output_type -> resource.Empty
	0, // 5: resource.Resource.Create:output_type -> resource.Empty
	0, // 6: resource.Resource.Update:output_type -> resource.Empty
	0, // 7: resource.Resource.Delete:output_type -> resource.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_resource_service_proto_init() }
func file_resource_service_proto_init() {
	if File_resource_service_proto != nil {
		return
	}
	file_resource_message_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resource_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_resource_service_proto_goTypes,
		DependencyIndexes: file_resource_service_proto_depIdxs,
	}.Build()
	File_resource_service_proto = out.File
	file_resource_service_proto_rawDesc = nil
	file_resource_service_proto_goTypes = nil
	file_resource_service_proto_depIdxs = nil
}
