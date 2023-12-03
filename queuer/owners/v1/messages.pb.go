// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        (unknown)
// source: queuer/owners/v1/messages.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/gsols/goproto/queuer/entities/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateOwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateOwnerRequest) Reset() {
	*x = CreateOwnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_owners_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOwnerRequest) ProtoMessage() {}

func (x *CreateOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_owners_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOwnerRequest.ProtoReflect.Descriptor instead.
func (*CreateOwnerRequest) Descriptor() ([]byte, []int) {
	return file_queuer_owners_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOwnerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateOwnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner *v1.Owner `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *CreateOwnerResponse) Reset() {
	*x = CreateOwnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_owners_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOwnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOwnerResponse) ProtoMessage() {}

func (x *CreateOwnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_owners_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOwnerResponse.ProtoReflect.Descriptor instead.
func (*CreateOwnerResponse) Descriptor() ([]byte, []int) {
	return file_queuer_owners_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOwnerResponse) GetOwner() *v1.Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

type GetOwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *GetOwnerRequest) Reset() {
	*x = GetOwnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_owners_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOwnerRequest) ProtoMessage() {}

func (x *GetOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_owners_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOwnerRequest.ProtoReflect.Descriptor instead.
func (*GetOwnerRequest) Descriptor() ([]byte, []int) {
	return file_queuer_owners_v1_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GetOwnerRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type GetOwnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner *v1.Owner `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *GetOwnerResponse) Reset() {
	*x = GetOwnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_owners_v1_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOwnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOwnerResponse) ProtoMessage() {}

func (x *GetOwnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_owners_v1_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOwnerResponse.ProtoReflect.Descriptor instead.
func (*GetOwnerResponse) Descriptor() ([]byte, []int) {
	return file_queuer_owners_v1_messages_proto_rawDescGZIP(), []int{3}
}

func (x *GetOwnerResponse) GetOwner() *v1.Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

type ListStreamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *ListStreamsRequest) Reset() {
	*x = ListStreamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_owners_v1_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStreamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStreamsRequest) ProtoMessage() {}

func (x *ListStreamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_owners_v1_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStreamsRequest.ProtoReflect.Descriptor instead.
func (*ListStreamsRequest) Descriptor() ([]byte, []int) {
	return file_queuer_owners_v1_messages_proto_rawDescGZIP(), []int{4}
}

func (x *ListStreamsRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type ListStreamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Streams []*v1.Stream `protobuf:"bytes,1,rep,name=streams,proto3" json:"streams,omitempty"`
}

func (x *ListStreamsResponse) Reset() {
	*x = ListStreamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_owners_v1_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStreamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStreamsResponse) ProtoMessage() {}

func (x *ListStreamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_owners_v1_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStreamsResponse.ProtoReflect.Descriptor instead.
func (*ListStreamsResponse) Descriptor() ([]byte, []int) {
	return file_queuer_owners_v1_messages_proto_rawDescGZIP(), []int{5}
}

func (x *ListStreamsResponse) GetStreams() []*v1.Stream {
	if x != nil {
		return x.Streams
	}
	return nil
}

type RegisterConsumerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId  string       `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Consumer *v1.Consumer `protobuf:"bytes,2,opt,name=consumer,proto3" json:"consumer,omitempty"`
}

func (x *RegisterConsumerRequest) Reset() {
	*x = RegisterConsumerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_owners_v1_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterConsumerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterConsumerRequest) ProtoMessage() {}

func (x *RegisterConsumerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_owners_v1_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterConsumerRequest.ProtoReflect.Descriptor instead.
func (*RegisterConsumerRequest) Descriptor() ([]byte, []int) {
	return file_queuer_owners_v1_messages_proto_rawDescGZIP(), []int{6}
}

func (x *RegisterConsumerRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *RegisterConsumerRequest) GetConsumer() *v1.Consumer {
	if x != nil {
		return x.Consumer
	}
	return nil
}

type RegisterConsumerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *v1.Result `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *RegisterConsumerResponse) Reset() {
	*x = RegisterConsumerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_owners_v1_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterConsumerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterConsumerResponse) ProtoMessage() {}

func (x *RegisterConsumerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_owners_v1_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterConsumerResponse.ProtoReflect.Descriptor instead.
func (*RegisterConsumerResponse) Descriptor() ([]byte, []int) {
	return file_queuer_owners_v1_messages_proto_rawDescGZIP(), []int{7}
}

func (x *RegisterConsumerResponse) GetResult() *v1.Result {
	if x != nil {
		return x.Result
	}
	return nil
}

type GetConsumersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *GetConsumersRequest) Reset() {
	*x = GetConsumersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_owners_v1_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsumersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsumersRequest) ProtoMessage() {}

func (x *GetConsumersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_owners_v1_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsumersRequest.ProtoReflect.Descriptor instead.
func (*GetConsumersRequest) Descriptor() ([]byte, []int) {
	return file_queuer_owners_v1_messages_proto_rawDescGZIP(), []int{8}
}

func (x *GetConsumersRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type GetConsumersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consumers []*v1.Consumer `protobuf:"bytes,1,rep,name=consumers,proto3" json:"consumers,omitempty"`
}

func (x *GetConsumersResponse) Reset() {
	*x = GetConsumersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queuer_owners_v1_messages_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsumersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsumersResponse) ProtoMessage() {}

func (x *GetConsumersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_queuer_owners_v1_messages_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsumersResponse.ProtoReflect.Descriptor instead.
func (*GetConsumersResponse) Descriptor() ([]byte, []int) {
	return file_queuer_owners_v1_messages_proto_rawDescGZIP(), []int{9}
}

func (x *GetConsumersResponse) GetConsumers() []*v1.Consumer {
	if x != nil {
		return x.Consumers
	}
	return nil
}

var File_queuer_owners_v1_messages_proto protoreflect.FileDescriptor

var file_queuer_owners_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x46, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72,
	0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x43, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x39, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x4b, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x22, 0x78,
	0x0a, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38,
	0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x08,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x22, 0x4e, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x09, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x73, 0x6f, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2f, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_queuer_owners_v1_messages_proto_rawDescOnce sync.Once
	file_queuer_owners_v1_messages_proto_rawDescData = file_queuer_owners_v1_messages_proto_rawDesc
)

func file_queuer_owners_v1_messages_proto_rawDescGZIP() []byte {
	file_queuer_owners_v1_messages_proto_rawDescOnce.Do(func() {
		file_queuer_owners_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_queuer_owners_v1_messages_proto_rawDescData)
	})
	return file_queuer_owners_v1_messages_proto_rawDescData
}

var file_queuer_owners_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_queuer_owners_v1_messages_proto_goTypes = []interface{}{
	(*CreateOwnerRequest)(nil),       // 0: queuer.owners.v1.CreateOwnerRequest
	(*CreateOwnerResponse)(nil),      // 1: queuer.owners.v1.CreateOwnerResponse
	(*GetOwnerRequest)(nil),          // 2: queuer.owners.v1.GetOwnerRequest
	(*GetOwnerResponse)(nil),         // 3: queuer.owners.v1.GetOwnerResponse
	(*ListStreamsRequest)(nil),       // 4: queuer.owners.v1.ListStreamsRequest
	(*ListStreamsResponse)(nil),      // 5: queuer.owners.v1.ListStreamsResponse
	(*RegisterConsumerRequest)(nil),  // 6: queuer.owners.v1.RegisterConsumerRequest
	(*RegisterConsumerResponse)(nil), // 7: queuer.owners.v1.RegisterConsumerResponse
	(*GetConsumersRequest)(nil),      // 8: queuer.owners.v1.GetConsumersRequest
	(*GetConsumersResponse)(nil),     // 9: queuer.owners.v1.GetConsumersResponse
	(*v1.Owner)(nil),                 // 10: queuer.entities.v1.Owner
	(*v1.Stream)(nil),                // 11: queuer.entities.v1.Stream
	(*v1.Consumer)(nil),              // 12: queuer.entities.v1.Consumer
	(*v1.Result)(nil),                // 13: queuer.entities.v1.Result
}
var file_queuer_owners_v1_messages_proto_depIdxs = []int32{
	10, // 0: queuer.owners.v1.CreateOwnerResponse.owner:type_name -> queuer.entities.v1.Owner
	10, // 1: queuer.owners.v1.GetOwnerResponse.owner:type_name -> queuer.entities.v1.Owner
	11, // 2: queuer.owners.v1.ListStreamsResponse.streams:type_name -> queuer.entities.v1.Stream
	12, // 3: queuer.owners.v1.RegisterConsumerRequest.consumer:type_name -> queuer.entities.v1.Consumer
	13, // 4: queuer.owners.v1.RegisterConsumerResponse.result:type_name -> queuer.entities.v1.Result
	12, // 5: queuer.owners.v1.GetConsumersResponse.consumers:type_name -> queuer.entities.v1.Consumer
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_queuer_owners_v1_messages_proto_init() }
func file_queuer_owners_v1_messages_proto_init() {
	if File_queuer_owners_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_queuer_owners_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOwnerRequest); i {
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
		file_queuer_owners_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOwnerResponse); i {
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
		file_queuer_owners_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOwnerRequest); i {
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
		file_queuer_owners_v1_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOwnerResponse); i {
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
		file_queuer_owners_v1_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStreamsRequest); i {
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
		file_queuer_owners_v1_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListStreamsResponse); i {
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
		file_queuer_owners_v1_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterConsumerRequest); i {
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
		file_queuer_owners_v1_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterConsumerResponse); i {
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
		file_queuer_owners_v1_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsumersRequest); i {
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
		file_queuer_owners_v1_messages_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsumersResponse); i {
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
			RawDescriptor: file_queuer_owners_v1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_queuer_owners_v1_messages_proto_goTypes,
		DependencyIndexes: file_queuer_owners_v1_messages_proto_depIdxs,
		MessageInfos:      file_queuer_owners_v1_messages_proto_msgTypes,
	}.Build()
	File_queuer_owners_v1_messages_proto = out.File
	file_queuer_owners_v1_messages_proto_rawDesc = nil
	file_queuer_owners_v1_messages_proto_goTypes = nil
	file_queuer_owners_v1_messages_proto_depIdxs = nil
}
