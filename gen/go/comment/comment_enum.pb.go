// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: comment_enum.proto

package comment

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

type SortOrder int32

const (
	SortOrder_NEWEST SortOrder = 0
	SortOrder_OLDEST SortOrder = 1
)

// Enum value maps for SortOrder.
var (
	SortOrder_name = map[int32]string{
		0: "NEWEST",
		1: "OLDEST",
	}
	SortOrder_value = map[string]int32{
		"NEWEST": 0,
		"OLDEST": 1,
	}
)

func (x SortOrder) Enum() *SortOrder {
	p := new(SortOrder)
	*p = x
	return p
}

func (x SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_comment_enum_proto_enumTypes[0].Descriptor()
}

func (SortOrder) Type() protoreflect.EnumType {
	return &file_comment_enum_proto_enumTypes[0]
}

func (x SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortOrder.Descriptor instead.
func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_comment_enum_proto_rawDescGZIP(), []int{0}
}

var File_comment_enum_proto protoreflect.FileDescriptor

var file_comment_enum_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2a, 0x23, 0x0a,
	0x09, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x45,
	0x57, 0x45, 0x53, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x4c, 0x44, 0x45, 0x53, 0x54,
	0x10, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comment_enum_proto_rawDescOnce sync.Once
	file_comment_enum_proto_rawDescData = file_comment_enum_proto_rawDesc
)

func file_comment_enum_proto_rawDescGZIP() []byte {
	file_comment_enum_proto_rawDescOnce.Do(func() {
		file_comment_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_enum_proto_rawDescData)
	})
	return file_comment_enum_proto_rawDescData
}

var file_comment_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_comment_enum_proto_goTypes = []any{
	(SortOrder)(0), // 0: comment.SortOrder
}
var file_comment_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_comment_enum_proto_init() }
func file_comment_enum_proto_init() {
	if File_comment_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_comment_enum_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_comment_enum_proto_goTypes,
		DependencyIndexes: file_comment_enum_proto_depIdxs,
		EnumInfos:         file_comment_enum_proto_enumTypes,
	}.Build()
	File_comment_enum_proto = out.File
	file_comment_enum_proto_rawDesc = nil
	file_comment_enum_proto_goTypes = nil
	file_comment_enum_proto_depIdxs = nil
}
