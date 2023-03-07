// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: producer.proto

package proto

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

type MainProducerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProducerId           string  `protobuf:"bytes,1,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	ProducerServiceArea  string  `protobuf:"bytes,2,opt,name=producer_service_area,json=producerServiceArea,proto3" json:"producer_service_area,omitempty"`
	ProducerCreatedAt    string  `protobuf:"bytes,3,opt,name=producer_created_at,json=producerCreatedAt,proto3" json:"producer_created_at,omitempty"`
	ProducerMessageId    string  `protobuf:"bytes,4,opt,name=producer_message_id,json=producerMessageId,proto3" json:"producer_message_id,omitempty"`
	Host                 string  `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	Client               string  `protobuf:"bytes,6,opt,name=client,proto3" json:"client,omitempty"`
	Ip                   string  `protobuf:"bytes,7,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 string  `protobuf:"bytes,8,opt,name=port,proto3" json:"port,omitempty"`
	InformationCreatedAt string  `protobuf:"bytes,9,opt,name=information_created_at,json=informationCreatedAt,proto3" json:"information_created_at,omitempty"`
	DataProducerId       string  `protobuf:"bytes,10,opt,name=data_producer_id,json=dataProducerId,proto3" json:"data_producer_id,omitempty"`
	Product              string  `protobuf:"bytes,11,opt,name=product,proto3" json:"product,omitempty"`
	Name                 string  `protobuf:"bytes,12,opt,name=name,proto3" json:"name,omitempty"`
	Category             string  `protobuf:"bytes,13,opt,name=category,proto3" json:"category,omitempty"`
	SubCategory          string  `protobuf:"bytes,14,opt,name=sub_category,json=subCategory,proto3" json:"sub_category,omitempty"`
	Price                float64 `protobuf:"fixed64,15,opt,name=price,proto3" json:"price,omitempty"`
	Quantity             int32   `protobuf:"varint,16,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Supplier             string  `protobuf:"bytes,17,opt,name=supplier,proto3" json:"supplier,omitempty"`
	Description          string  `protobuf:"bytes,18,opt,name=description,proto3" json:"description,omitempty"`
	Gender               string  `protobuf:"bytes,19,opt,name=gender,proto3" json:"gender,omitempty"`
}

func (x *MainProducerRequest) Reset() {
	*x = MainProducerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainProducerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainProducerRequest) ProtoMessage() {}

func (x *MainProducerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_producer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainProducerRequest.ProtoReflect.Descriptor instead.
func (*MainProducerRequest) Descriptor() ([]byte, []int) {
	return file_producer_proto_rawDescGZIP(), []int{0}
}

func (x *MainProducerRequest) GetProducerId() string {
	if x != nil {
		return x.ProducerId
	}
	return ""
}

func (x *MainProducerRequest) GetProducerServiceArea() string {
	if x != nil {
		return x.ProducerServiceArea
	}
	return ""
}

func (x *MainProducerRequest) GetProducerCreatedAt() string {
	if x != nil {
		return x.ProducerCreatedAt
	}
	return ""
}

func (x *MainProducerRequest) GetProducerMessageId() string {
	if x != nil {
		return x.ProducerMessageId
	}
	return ""
}

func (x *MainProducerRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *MainProducerRequest) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *MainProducerRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *MainProducerRequest) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *MainProducerRequest) GetInformationCreatedAt() string {
	if x != nil {
		return x.InformationCreatedAt
	}
	return ""
}

func (x *MainProducerRequest) GetDataProducerId() string {
	if x != nil {
		return x.DataProducerId
	}
	return ""
}

func (x *MainProducerRequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *MainProducerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MainProducerRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *MainProducerRequest) GetSubCategory() string {
	if x != nil {
		return x.SubCategory
	}
	return ""
}

func (x *MainProducerRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *MainProducerRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *MainProducerRequest) GetSupplier() string {
	if x != nil {
		return x.Supplier
	}
	return ""
}

func (x *MainProducerRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MainProducerRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

type MainProducerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProducerId       string `protobuf:"bytes,1,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	HistoryMessage   string `protobuf:"bytes,2,opt,name=history_message,json=historyMessage,proto3" json:"history_message,omitempty"`
	HistoryCreatedAt string `protobuf:"bytes,3,opt,name=history_created_at,json=historyCreatedAt,proto3" json:"history_created_at,omitempty"`
}

func (x *MainProducerResponse) Reset() {
	*x = MainProducerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_producer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MainProducerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MainProducerResponse) ProtoMessage() {}

func (x *MainProducerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_producer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MainProducerResponse.ProtoReflect.Descriptor instead.
func (*MainProducerResponse) Descriptor() ([]byte, []int) {
	return file_producer_proto_rawDescGZIP(), []int{1}
}

func (x *MainProducerResponse) GetProducerId() string {
	if x != nil {
		return x.ProducerId
	}
	return ""
}

func (x *MainProducerResponse) GetHistoryMessage() string {
	if x != nil {
		return x.HistoryMessage
	}
	return ""
}

func (x *MainProducerResponse) GetHistoryCreatedAt() string {
	if x != nil {
		return x.HistoryCreatedAt
	}
	return ""
}

var File_producer_proto protoreflect.FileDescriptor

var file_producer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x22, 0xef, 0x04, 0x0a, 0x13, 0x4d,
	0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x41, 0x72, 0x65, 0x61, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a,
	0x10, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x8e, 0x01, 0x0a,
	0x14, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2c, 0x0a, 0x12, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x66, 0x0a,
	0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x53, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x72, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x2e, 0x4d, 0x61,
	0x69, 0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x69,
	0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x27, 0x5a, 0x25, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x33, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_producer_proto_rawDescOnce sync.Once
	file_producer_proto_rawDescData = file_producer_proto_rawDesc
)

func file_producer_proto_rawDescGZIP() []byte {
	file_producer_proto_rawDescOnce.Do(func() {
		file_producer_proto_rawDescData = protoimpl.X.CompressGZIP(file_producer_proto_rawDescData)
	})
	return file_producer_proto_rawDescData
}

var file_producer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_producer_proto_goTypes = []interface{}{
	(*MainProducerRequest)(nil),  // 0: producer.MainProducerRequest
	(*MainProducerResponse)(nil), // 1: producer.MainProducerResponse
}
var file_producer_proto_depIdxs = []int32{
	0, // 0: producer.ProducerService.StreamProducer:input_type -> producer.MainProducerRequest
	1, // 1: producer.ProducerService.StreamProducer:output_type -> producer.MainProducerResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_producer_proto_init() }
func file_producer_proto_init() {
	if File_producer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_producer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainProducerRequest); i {
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
		file_producer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MainProducerResponse); i {
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
			RawDescriptor: file_producer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_producer_proto_goTypes,
		DependencyIndexes: file_producer_proto_depIdxs,
		MessageInfos:      file_producer_proto_msgTypes,
	}.Build()
	File_producer_proto = out.File
	file_producer_proto_rawDesc = nil
	file_producer_proto_goTypes = nil
	file_producer_proto_depIdxs = nil
}
