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

	OwnerId   string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
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

func (x *ListStreamsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListStreamsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
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
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x46, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x36, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x43, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x22, 0x7e, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72,
	0x03, 0xb0, 0x01, 0x01, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x4b, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x72, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x73,
	0x6f, 0x6c, 0x73, 0x2f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x72, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_queuer_owners_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_queuer_owners_v1_messages_proto_goTypes = []interface{}{
	(*CreateOwnerRequest)(nil),  // 0: queuer.owners.v1.CreateOwnerRequest
	(*CreateOwnerResponse)(nil), // 1: queuer.owners.v1.CreateOwnerResponse
	(*GetOwnerRequest)(nil),     // 2: queuer.owners.v1.GetOwnerRequest
	(*GetOwnerResponse)(nil),    // 3: queuer.owners.v1.GetOwnerResponse
	(*ListStreamsRequest)(nil),  // 4: queuer.owners.v1.ListStreamsRequest
	(*ListStreamsResponse)(nil), // 5: queuer.owners.v1.ListStreamsResponse
	(*v1.Owner)(nil),            // 6: queuer.entities.v1.Owner
	(*v1.Stream)(nil),           // 7: queuer.entities.v1.Stream
}
var file_queuer_owners_v1_messages_proto_depIdxs = []int32{
	6, // 0: queuer.owners.v1.CreateOwnerResponse.owner:type_name -> queuer.entities.v1.Owner
	6, // 1: queuer.owners.v1.GetOwnerResponse.owner:type_name -> queuer.entities.v1.Owner
	7, // 2: queuer.owners.v1.ListStreamsResponse.streams:type_name -> queuer.entities.v1.Stream
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_queuer_owners_v1_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
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
