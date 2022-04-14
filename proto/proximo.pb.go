// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: proximo.proto

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

type Offset int32

const (
	Offset_OFFSET_DEFAULT Offset = 0
	Offset_OFFSET_NEWEST  Offset = 1
	Offset_OFFSET_OLDEST  Offset = 2
)

// Enum value maps for Offset.
var (
	Offset_name = map[int32]string{
		0: "OFFSET_DEFAULT",
		1: "OFFSET_NEWEST",
		2: "OFFSET_OLDEST",
	}
	Offset_value = map[string]int32{
		"OFFSET_DEFAULT": 0,
		"OFFSET_NEWEST":  1,
		"OFFSET_OLDEST":  2,
	}
)

func (x Offset) Enum() *Offset {
	p := new(Offset)
	*p = x
	return p
}

func (x Offset) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Offset) Descriptor() protoreflect.EnumDescriptor {
	return file_proximo_proto_enumTypes[0].Descriptor()
}

func (Offset) Type() protoreflect.EnumType {
	return &file_proximo_proto_enumTypes[0]
}

func (x Offset) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Offset.Descriptor instead.
func (Offset) EnumDescriptor() ([]byte, []int) {
	return file_proximo_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proximo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proximo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proximo_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ConsumerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// expected if this is a start request
	StartRequest *StartConsumeRequest `protobuf:"bytes,2,opt,name=startRequest,proto3" json:"startRequest,omitempty"`
	// expected if this is a confirmation
	Confirmation *Confirmation `protobuf:"bytes,3,opt,name=confirmation,proto3" json:"confirmation,omitempty"`
}

func (x *ConsumerRequest) Reset() {
	*x = ConsumerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proximo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerRequest) ProtoMessage() {}

func (x *ConsumerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proximo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerRequest.ProtoReflect.Descriptor instead.
func (*ConsumerRequest) Descriptor() ([]byte, []int) {
	return file_proximo_proto_rawDescGZIP(), []int{1}
}

func (x *ConsumerRequest) GetStartRequest() *StartConsumeRequest {
	if x != nil {
		return x.StartRequest
	}
	return nil
}

func (x *ConsumerRequest) GetConfirmation() *Confirmation {
	if x != nil {
		return x.Confirmation
	}
	return nil
}

type StartConsumeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic         string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Consumer      string `protobuf:"bytes,2,opt,name=consumer,proto3" json:"consumer,omitempty"`
	InitialOffset Offset `protobuf:"varint,3,opt,name=initial_offset,json=initialOffset,proto3,enum=proximo.Offset" json:"initial_offset,omitempty"`
}

func (x *StartConsumeRequest) Reset() {
	*x = StartConsumeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proximo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartConsumeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartConsumeRequest) ProtoMessage() {}

func (x *StartConsumeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proximo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartConsumeRequest.ProtoReflect.Descriptor instead.
func (*StartConsumeRequest) Descriptor() ([]byte, []int) {
	return file_proximo_proto_rawDescGZIP(), []int{2}
}

func (x *StartConsumeRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *StartConsumeRequest) GetConsumer() string {
	if x != nil {
		return x.Consumer
	}
	return ""
}

func (x *StartConsumeRequest) GetInitialOffset() Offset {
	if x != nil {
		return x.InitialOffset
	}
	return Offset_OFFSET_DEFAULT
}

type Confirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgID string `protobuf:"bytes,1,opt,name=msgID,proto3" json:"msgID,omitempty"`
}

func (x *Confirmation) Reset() {
	*x = Confirmation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proximo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmation) ProtoMessage() {}

func (x *Confirmation) ProtoReflect() protoreflect.Message {
	mi := &file_proximo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Confirmation.ProtoReflect.Descriptor instead.
func (*Confirmation) Descriptor() ([]byte, []int) {
	return file_proximo_proto_rawDescGZIP(), []int{3}
}

func (x *Confirmation) GetMsgID() string {
	if x != nil {
		return x.MsgID
	}
	return ""
}

type PublisherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// expected if this is a start request
	StartRequest *StartPublishRequest `protobuf:"bytes,2,opt,name=startRequest,proto3" json:"startRequest,omitempty"`
	// expected if this is a message
	Msg *Message `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *PublisherRequest) Reset() {
	*x = PublisherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proximo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublisherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublisherRequest) ProtoMessage() {}

func (x *PublisherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proximo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublisherRequest.ProtoReflect.Descriptor instead.
func (*PublisherRequest) Descriptor() ([]byte, []int) {
	return file_proximo_proto_rawDescGZIP(), []int{4}
}

func (x *PublisherRequest) GetStartRequest() *StartPublishRequest {
	if x != nil {
		return x.StartRequest
	}
	return nil
}

func (x *PublisherRequest) GetMsg() *Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

type StartPublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *StartPublishRequest) Reset() {
	*x = StartPublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proximo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartPublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartPublishRequest) ProtoMessage() {}

func (x *StartPublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proximo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartPublishRequest.ProtoReflect.Descriptor instead.
func (*StartPublishRequest) Descriptor() ([]byte, []int) {
	return file_proximo_proto_rawDescGZIP(), []int{5}
}

func (x *StartPublishRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

var File_proximo_proto protoreflect.FileDescriptor

var file_proximo_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x6f, 0x22, 0x2d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0c, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x6f, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7f, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x12, 0x36, 0x0a, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x78,
	0x69, 0x6d, 0x6f, 0x2e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x52, 0x0d, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x24, 0x0a, 0x0c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x44, 0x22,
	0x78, 0x0a, 0x10, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x78,
	0x69, 0x6d, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x6f, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2b, 0x0a, 0x13, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2a, 0x42, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x46, 0x46, 0x53, 0x45, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55,
	0x4c, 0x54, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x46, 0x46, 0x53, 0x45, 0x54, 0x5f, 0x4e,
	0x45, 0x57, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x46, 0x46, 0x53, 0x45,
	0x54, 0x5f, 0x4f, 0x4c, 0x44, 0x45, 0x53, 0x54, 0x10, 0x02, 0x32, 0x4c, 0x0a, 0x0d, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x6f,
	0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32, 0x50, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x6e, 0x6b, 0x12, 0x41, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x6f, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x77, 0x2d, 0x6c, 0x61, 0x62, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proximo_proto_rawDescOnce sync.Once
	file_proximo_proto_rawDescData = file_proximo_proto_rawDesc
)

func file_proximo_proto_rawDescGZIP() []byte {
	file_proximo_proto_rawDescOnce.Do(func() {
		file_proximo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proximo_proto_rawDescData)
	})
	return file_proximo_proto_rawDescData
}

var file_proximo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proximo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proximo_proto_goTypes = []interface{}{
	(Offset)(0),                 // 0: proximo.Offset
	(*Message)(nil),             // 1: proximo.Message
	(*ConsumerRequest)(nil),     // 2: proximo.ConsumerRequest
	(*StartConsumeRequest)(nil), // 3: proximo.StartConsumeRequest
	(*Confirmation)(nil),        // 4: proximo.Confirmation
	(*PublisherRequest)(nil),    // 5: proximo.PublisherRequest
	(*StartPublishRequest)(nil), // 6: proximo.StartPublishRequest
}
var file_proximo_proto_depIdxs = []int32{
	3, // 0: proximo.ConsumerRequest.startRequest:type_name -> proximo.StartConsumeRequest
	4, // 1: proximo.ConsumerRequest.confirmation:type_name -> proximo.Confirmation
	0, // 2: proximo.StartConsumeRequest.initial_offset:type_name -> proximo.Offset
	6, // 3: proximo.PublisherRequest.startRequest:type_name -> proximo.StartPublishRequest
	1, // 4: proximo.PublisherRequest.msg:type_name -> proximo.Message
	2, // 5: proximo.MessageSource.Consume:input_type -> proximo.ConsumerRequest
	5, // 6: proximo.MessageSink.Publish:input_type -> proximo.PublisherRequest
	1, // 7: proximo.MessageSource.Consume:output_type -> proximo.Message
	4, // 8: proximo.MessageSink.Publish:output_type -> proximo.Confirmation
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proximo_proto_init() }
func file_proximo_proto_init() {
	if File_proximo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proximo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_proximo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerRequest); i {
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
		file_proximo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartConsumeRequest); i {
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
		file_proximo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Confirmation); i {
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
		file_proximo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublisherRequest); i {
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
		file_proximo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartPublishRequest); i {
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
			RawDescriptor: file_proximo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proximo_proto_goTypes,
		DependencyIndexes: file_proximo_proto_depIdxs,
		EnumInfos:         file_proximo_proto_enumTypes,
		MessageInfos:      file_proximo_proto_msgTypes,
	}.Build()
	File_proximo_proto = out.File
	file_proximo_proto_rawDesc = nil
	file_proximo_proto_goTypes = nil
	file_proximo_proto_depIdxs = nil
}
