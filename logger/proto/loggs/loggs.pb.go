// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: loggs.proto

package loggs

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

type MainLoggerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoggerId      string `protobuf:"bytes,1,opt,name=logger_id,json=loggerId,proto3" json:"logger_id,omitempty"`
	LoggerMessage string `protobuf:"bytes,2,opt,name=logger_message,json=loggerMessage,proto3" json:"logger_message,omitempty"`
}

func (x *MainLoggerRequest) Reset() {
	*x = MainLoggerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loggs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainLoggerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainLoggerRequest) ProtoMessage() {}

func (x *MainLoggerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_loggs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainLoggerRequest.ProtoReflect.Descriptor instead.
func (*MainLoggerRequest) Descriptor() ([]byte, []int) {
	return file_loggs_proto_rawDescGZIP(), []int{0}
}

func (x *MainLoggerRequest) GetLoggerId() string {
	if x != nil {
		return x.LoggerId
	}
	return ""
}

func (x *MainLoggerRequest) GetLoggerMessage() string {
	if x != nil {
		return x.LoggerMessage
	}
	return ""
}

type MainLoggerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseLogger string `protobuf:"bytes,1,opt,name=response_logger,json=responseLogger,proto3" json:"response_logger,omitempty"`
}

func (x *MainLoggerResponse) Reset() {
	*x = MainLoggerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loggs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainLoggerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainLoggerResponse) ProtoMessage() {}

func (x *MainLoggerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_loggs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainLoggerResponse.ProtoReflect.Descriptor instead.
func (*MainLoggerResponse) Descriptor() ([]byte, []int) {
	return file_loggs_proto_rawDescGZIP(), []int{1}
}

func (x *MainLoggerResponse) GetResponseLogger() string {
	if x != nil {
		return x.ResponseLogger
	}
	return ""
}

var File_loggs_proto protoreflect.FileDescriptor

var file_loggs_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x22, 0x57, 0x0a, 0x11, 0x4d, 0x61, 0x69, 0x6e, 0x4c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x3d, 0x0a, 0x12, 0x4d, 0x61, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x32,
	0x67, 0x0a, 0x14, 0x4c, 0x6f, 0x67, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65,
	0x72, 0x2e, 0x4d, 0x61, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x21, 0x5a, 0x1f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_loggs_proto_rawDescOnce sync.Once
	file_loggs_proto_rawDescData = file_loggs_proto_rawDesc
)

func file_loggs_proto_rawDescGZIP() []byte {
	file_loggs_proto_rawDescOnce.Do(func() {
		file_loggs_proto_rawDescData = protoimpl.X.CompressGZIP(file_loggs_proto_rawDescData)
	})
	return file_loggs_proto_rawDescData
}

var file_loggs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_loggs_proto_goTypes = []interface{}{
	(*MainLoggerRequest)(nil),  // 0: producer.MainLoggerRequest
	(*MainLoggerResponse)(nil), // 1: producer.MainLoggerResponse
}
var file_loggs_proto_depIdxs = []int32{
	0, // 0: producer.LoggsProducerService.StreamProducer:input_type -> producer.MainLoggerRequest
	1, // 1: producer.LoggsProducerService.StreamProducer:output_type -> producer.MainLoggerResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_loggs_proto_init() }
func file_loggs_proto_init() {
	if File_loggs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loggs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainLoggerRequest); i {
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
		file_loggs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainLoggerResponse); i {
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
			RawDescriptor: file_loggs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_loggs_proto_goTypes,
		DependencyIndexes: file_loggs_proto_depIdxs,
		MessageInfos:      file_loggs_proto_msgTypes,
	}.Build()
	File_loggs_proto = out.File
	file_loggs_proto_rawDesc = nil
	file_loggs_proto_goTypes = nil
	file_loggs_proto_depIdxs = nil
}
