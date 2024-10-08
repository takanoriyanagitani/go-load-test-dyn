// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: loadtest_dyn/raw/seed/integer/v1/seed2bytes.proto

package v1

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

type BytesFrom64IRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seed int64 `protobuf:"fixed64,1,opt,name=seed,proto3" json:"seed,omitempty"`
}

func (x *BytesFrom64IRequest) Reset() {
	*x = BytesFrom64IRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesFrom64IRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesFrom64IRequest) ProtoMessage() {}

func (x *BytesFrom64IRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesFrom64IRequest.ProtoReflect.Descriptor instead.
func (*BytesFrom64IRequest) Descriptor() ([]byte, []int) {
	return file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_rawDescGZIP(), []int{0}
}

func (x *BytesFrom64IRequest) GetSeed() int64 {
	if x != nil {
		return x.Seed
	}
	return 0
}

type BytesFrom64IResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Generated []byte `protobuf:"bytes,1,opt,name=generated,proto3" json:"generated,omitempty"`
}

func (x *BytesFrom64IResponse) Reset() {
	*x = BytesFrom64IResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesFrom64IResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesFrom64IResponse) ProtoMessage() {}

func (x *BytesFrom64IResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesFrom64IResponse.ProtoReflect.Descriptor instead.
func (*BytesFrom64IResponse) Descriptor() ([]byte, []int) {
	return file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_rawDescGZIP(), []int{1}
}

func (x *BytesFrom64IResponse) GetGenerated() []byte {
	if x != nil {
		return x.Generated
	}
	return nil
}

var File_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto protoreflect.FileDescriptor

var file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x79, 0x6e, 0x2f, 0x72,
	0x61, 0x77, 0x2f, 0x73, 0x65, 0x65, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x65, 0x64, 0x32, 0x62, 0x79, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x20, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x79,
	0x6e, 0x2e, 0x72, 0x61, 0x77, 0x2e, 0x73, 0x65, 0x65, 0x64, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x29, 0x0a, 0x13, 0x42, 0x79, 0x74, 0x65, 0x73, 0x46, 0x72,
	0x6f, 0x6d, 0x36, 0x34, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x10, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64,
	0x22, 0x34, 0x0a, 0x14, 0x42, 0x79, 0x74, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x36, 0x34, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x32, 0x8f, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x65, 0x64, 0x36,
	0x34, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7d, 0x0a, 0x0c, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x36, 0x34, 0x69, 0x12, 0x35, 0x2e, 0x6c, 0x6f, 0x61, 0x64,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x79, 0x6e, 0x2e, 0x72, 0x61, 0x77, 0x2e, 0x73, 0x65, 0x65,
	0x64, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x36, 0x34, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x36, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x79, 0x6e, 0x2e,
	0x72, 0x61, 0x77, 0x2e, 0x73, 0x65, 0x65, 0x64, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x36, 0x34, 0x69,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x6b, 0x61, 0x6e, 0x6f, 0x72, 0x69, 0x79,
	0x61, 0x6e, 0x61, 0x67, 0x69, 0x74, 0x61, 0x6e, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x6c, 0x6f, 0x61,
	0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x64, 0x79, 0x6e, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x64, 0x79, 0x6e, 0x2f, 0x72, 0x61, 0x77, 0x2f, 0x73, 0x65, 0x65, 0x64,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_rawDescOnce sync.Once
	file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_rawDescData = file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_rawDesc
)

func file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_rawDescGZIP() []byte {
	file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_rawDescOnce.Do(func() {
		file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_rawDescData = protoimpl.X.CompressGZIP(file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_rawDescData)
	})
	return file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_rawDescData
}

var file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_goTypes = []any{
	(*BytesFrom64IRequest)(nil),  // 0: loadtest_dyn.raw.seed.integer.v1.BytesFrom64iRequest
	(*BytesFrom64IResponse)(nil), // 1: loadtest_dyn.raw.seed.integer.v1.BytesFrom64iResponse
}
var file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_depIdxs = []int32{
	0, // 0: loadtest_dyn.raw.seed.integer.v1.Seed64iService.BytesFrom64i:input_type -> loadtest_dyn.raw.seed.integer.v1.BytesFrom64iRequest
	1, // 1: loadtest_dyn.raw.seed.integer.v1.Seed64iService.BytesFrom64i:output_type -> loadtest_dyn.raw.seed.integer.v1.BytesFrom64iResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_init() }
func file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_init() {
	if File_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BytesFrom64IRequest); i {
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
		file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BytesFrom64IResponse); i {
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
			RawDescriptor: file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_goTypes,
		DependencyIndexes: file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_depIdxs,
		MessageInfos:      file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_msgTypes,
	}.Build()
	File_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto = out.File
	file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_rawDesc = nil
	file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_goTypes = nil
	file_loadtest_dyn_raw_seed_integer_v1_seed2bytes_proto_depIdxs = nil
}
