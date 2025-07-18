// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: bind/v1/bind.proto

package bindv1

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

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Typ string `protobuf:"bytes,1,opt,name=typ,proto3" json:"typ,omitempty"`
	Id  string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bind_v1_bind_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_bind_v1_bind_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_bind_v1_bind_proto_rawDescGZIP(), []int{0}
}

func (x *Id) GetTyp() string {
	if x != nil {
		return x.Typ
	}
	return ""
}

func (x *Id) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// request
type BindIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fd  int64 `protobuf:"varint,1,opt,name=fd,proto3" json:"fd,omitempty"`
	Ids []*Id `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *BindIdRequest) Reset() {
	*x = BindIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bind_v1_bind_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindIdRequest) ProtoMessage() {}

func (x *BindIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bind_v1_bind_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindIdRequest.ProtoReflect.Descriptor instead.
func (*BindIdRequest) Descriptor() ([]byte, []int) {
	return file_bind_v1_bind_proto_rawDescGZIP(), []int{1}
}

func (x *BindIdRequest) GetFd() int64 {
	if x != nil {
		return x.Fd
	}
	return 0
}

func (x *BindIdRequest) GetIds() []*Id {
	if x != nil {
		return x.Ids
	}
	return nil
}

// response
type BindIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BindIdResponse) Reset() {
	*x = BindIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bind_v1_bind_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindIdResponse) ProtoMessage() {}

func (x *BindIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bind_v1_bind_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindIdResponse.ProtoReflect.Descriptor instead.
func (*BindIdResponse) Descriptor() ([]byte, []int) {
	return file_bind_v1_bind_proto_rawDescGZIP(), []int{2}
}

// request
type BindExistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *Id `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BindExistRequest) Reset() {
	*x = BindExistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bind_v1_bind_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindExistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindExistRequest) ProtoMessage() {}

func (x *BindExistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bind_v1_bind_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindExistRequest.ProtoReflect.Descriptor instead.
func (*BindExistRequest) Descriptor() ([]byte, []int) {
	return file_bind_v1_bind_proto_rawDescGZIP(), []int{3}
}

func (x *BindExistRequest) GetId() *Id {
	if x != nil {
		return x.Id
	}
	return nil
}

// response
type BindExistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist bool `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
}

func (x *BindExistResponse) Reset() {
	*x = BindExistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bind_v1_bind_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindExistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindExistResponse) ProtoMessage() {}

func (x *BindExistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bind_v1_bind_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindExistResponse.ProtoReflect.Descriptor instead.
func (*BindExistResponse) Descriptor() ([]byte, []int) {
	return file_bind_v1_bind_proto_rawDescGZIP(), []int{4}
}

func (x *BindExistResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

// request
type UnBindIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fd    int64    `protobuf:"varint,1,opt,name=fd,proto3" json:"fd,omitempty"`
	Types []string `protobuf:"bytes,2,rep,name=types,proto3" json:"types,omitempty"`
}

func (x *UnBindIdRequest) Reset() {
	*x = UnBindIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bind_v1_bind_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnBindIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnBindIdRequest) ProtoMessage() {}

func (x *UnBindIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bind_v1_bind_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnBindIdRequest.ProtoReflect.Descriptor instead.
func (*UnBindIdRequest) Descriptor() ([]byte, []int) {
	return file_bind_v1_bind_proto_rawDescGZIP(), []int{5}
}

func (x *UnBindIdRequest) GetFd() int64 {
	if x != nil {
		return x.Fd
	}
	return 0
}

func (x *UnBindIdRequest) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

// response
type UnBindIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnBindIdResponse) Reset() {
	*x = UnBindIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bind_v1_bind_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnBindIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnBindIdResponse) ProtoMessage() {}

func (x *UnBindIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bind_v1_bind_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnBindIdResponse.ProtoReflect.Descriptor instead.
func (*UnBindIdResponse) Descriptor() ([]byte, []int) {
	return file_bind_v1_bind_proto_rawDescGZIP(), []int{6}
}

// request
type DisconnectTargetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DisconnectTargetRequest) Reset() {
	*x = DisconnectTargetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bind_v1_bind_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectTargetRequest) ProtoMessage() {}

func (x *DisconnectTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bind_v1_bind_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectTargetRequest.ProtoReflect.Descriptor instead.
func (*DisconnectTargetRequest) Descriptor() ([]byte, []int) {
	return file_bind_v1_bind_proto_rawDescGZIP(), []int{7}
}

func (x *DisconnectTargetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// response
type DisconnectTargetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisconnectTargetResponse) Reset() {
	*x = DisconnectTargetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bind_v1_bind_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisconnectTargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisconnectTargetResponse) ProtoMessage() {}

func (x *DisconnectTargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bind_v1_bind_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisconnectTargetResponse.ProtoReflect.Descriptor instead.
func (*DisconnectTargetResponse) Descriptor() ([]byte, []int) {
	return file_bind_v1_bind_proto_rawDescGZIP(), []int{8}
}

// request
type ProxyTargetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fd     int64    `protobuf:"varint,1,opt,name=fd,proto3" json:"fd,omitempty"`
	Target []string `protobuf:"bytes,2,rep,name=target,proto3" json:"target,omitempty"`
}

func (x *ProxyTargetRequest) Reset() {
	*x = ProxyTargetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bind_v1_bind_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyTargetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyTargetRequest) ProtoMessage() {}

func (x *ProxyTargetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bind_v1_bind_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyTargetRequest.ProtoReflect.Descriptor instead.
func (*ProxyTargetRequest) Descriptor() ([]byte, []int) {
	return file_bind_v1_bind_proto_rawDescGZIP(), []int{9}
}

func (x *ProxyTargetRequest) GetFd() int64 {
	if x != nil {
		return x.Fd
	}
	return 0
}

func (x *ProxyTargetRequest) GetTarget() []string {
	if x != nil {
		return x.Target
	}
	return nil
}

// response
type ProxyTargetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProxyTargetResponse) Reset() {
	*x = ProxyTargetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bind_v1_bind_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyTargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyTargetResponse) ProtoMessage() {}

func (x *ProxyTargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bind_v1_bind_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyTargetResponse.ProtoReflect.Descriptor instead.
func (*ProxyTargetResponse) Descriptor() ([]byte, []int) {
	return file_bind_v1_bind_proto_rawDescGZIP(), []int{10}
}

type TargetBindIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target   string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	BindType string `protobuf:"bytes,2,opt,name=bind_type,json=bindType,proto3" json:"bind_type,omitempty"`
}

func (x *TargetBindIdRequest) Reset() {
	*x = TargetBindIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bind_v1_bind_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetBindIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetBindIdRequest) ProtoMessage() {}

func (x *TargetBindIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bind_v1_bind_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetBindIdRequest.ProtoReflect.Descriptor instead.
func (*TargetBindIdRequest) Descriptor() ([]byte, []int) {
	return file_bind_v1_bind_proto_rawDescGZIP(), []int{11}
}

func (x *TargetBindIdRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *TargetBindIdRequest) GetBindType() string {
	if x != nil {
		return x.BindType
	}
	return ""
}

type TargetBindIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *Id `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TargetBindIdResponse) Reset() {
	*x = TargetBindIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bind_v1_bind_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetBindIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetBindIdResponse) ProtoMessage() {}

func (x *TargetBindIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bind_v1_bind_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetBindIdResponse.ProtoReflect.Descriptor instead.
func (*TargetBindIdResponse) Descriptor() ([]byte, []int) {
	return file_bind_v1_bind_proto_rawDescGZIP(), []int{12}
}

func (x *TargetBindIdResponse) GetId() *Id {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_bind_v1_bind_proto protoreflect.FileDescriptor

var file_bind_v1_bind_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x69, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x26, 0x0a,
	0x02, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x79, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x74, 0x79, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x0d, 0x42, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x66, 0x64, 0x12, 0x1d, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x42, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x0a, 0x10, 0x42, 0x69, 0x6e, 0x64, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x11, 0x42, 0x69, 0x6e, 0x64,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x0f, 0x55, 0x6e, 0x42, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x66, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x66, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x12, 0x0a, 0x10,
	0x55, 0x6e, 0x42, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x29, 0x0a, 0x17, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x66, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x66, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4a, 0x0a, 0x13,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62,
	0x69, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x62, 0x69, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x33, 0x0a, 0x14, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62,
	0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x32, 0x91, 0x04,
	0x0a, 0x0b, 0x42, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a,
	0x06, 0x42, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x42, 0x69, 0x6e, 0x64,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x69, 0x6e, 0x64, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08,
	0x55, 0x6e, 0x42, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x6e, 0x42, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x42,
	0x69, 0x6e, 0x64, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a,
	0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x12, 0x20, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0f, 0x42, 0x69, 0x6e, 0x64, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x62, 0x69, 0x6e, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x11, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x62, 0x69, 0x6e, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x69,
	0x6e, 0x64, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x89, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x69, 0x6e, 0x64, 0x2e, 0x76,
	0x31, 0x42, 0x09, 0x42, 0x69, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x62, 0x6e, 0x61, 0x68,
	0x73, 0x67, 0x6e, 0x61, 0x77, 0x2f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x62, 0x69, 0x6e, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x69, 0x6e, 0x64,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x42, 0x69, 0x6e, 0x64, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x07, 0x42, 0x69, 0x6e, 0x64, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x42,
	0x69, 0x6e, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x08, 0x42, 0x69, 0x6e, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bind_v1_bind_proto_rawDescOnce sync.Once
	file_bind_v1_bind_proto_rawDescData = file_bind_v1_bind_proto_rawDesc
)

func file_bind_v1_bind_proto_rawDescGZIP() []byte {
	file_bind_v1_bind_proto_rawDescOnce.Do(func() {
		file_bind_v1_bind_proto_rawDescData = protoimpl.X.CompressGZIP(file_bind_v1_bind_proto_rawDescData)
	})
	return file_bind_v1_bind_proto_rawDescData
}

var file_bind_v1_bind_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_bind_v1_bind_proto_goTypes = []interface{}{
	(*Id)(nil),                       // 0: bind.v1.Id
	(*BindIdRequest)(nil),            // 1: bind.v1.BindIdRequest
	(*BindIdResponse)(nil),           // 2: bind.v1.BindIdResponse
	(*BindExistRequest)(nil),         // 3: bind.v1.BindExistRequest
	(*BindExistResponse)(nil),        // 4: bind.v1.BindExistResponse
	(*UnBindIdRequest)(nil),          // 5: bind.v1.UnBindIdRequest
	(*UnBindIdResponse)(nil),         // 6: bind.v1.UnBindIdResponse
	(*DisconnectTargetRequest)(nil),  // 7: bind.v1.DisconnectTargetRequest
	(*DisconnectTargetResponse)(nil), // 8: bind.v1.DisconnectTargetResponse
	(*ProxyTargetRequest)(nil),       // 9: bind.v1.ProxyTargetRequest
	(*ProxyTargetResponse)(nil),      // 10: bind.v1.ProxyTargetResponse
	(*TargetBindIdRequest)(nil),      // 11: bind.v1.TargetBindIdRequest
	(*TargetBindIdResponse)(nil),     // 12: bind.v1.TargetBindIdResponse
}
var file_bind_v1_bind_proto_depIdxs = []int32{
	0,  // 0: bind.v1.BindIdRequest.ids:type_name -> bind.v1.Id
	0,  // 1: bind.v1.BindExistRequest.id:type_name -> bind.v1.Id
	0,  // 2: bind.v1.TargetBindIdResponse.id:type_name -> bind.v1.Id
	1,  // 3: bind.v1.BindService.BindId:input_type -> bind.v1.BindIdRequest
	3,  // 4: bind.v1.BindService.BindExist:input_type -> bind.v1.BindExistRequest
	5,  // 5: bind.v1.BindService.UnBindId:input_type -> bind.v1.UnBindIdRequest
	7,  // 6: bind.v1.BindService.DisconnectTarget:input_type -> bind.v1.DisconnectTargetRequest
	9,  // 7: bind.v1.BindService.BindProxyTarget:input_type -> bind.v1.ProxyTargetRequest
	9,  // 8: bind.v1.BindService.UnbindProxyTarget:input_type -> bind.v1.ProxyTargetRequest
	11, // 9: bind.v1.BindService.TargetBindId:input_type -> bind.v1.TargetBindIdRequest
	2,  // 10: bind.v1.BindService.BindId:output_type -> bind.v1.BindIdResponse
	4,  // 11: bind.v1.BindService.BindExist:output_type -> bind.v1.BindExistResponse
	6,  // 12: bind.v1.BindService.UnBindId:output_type -> bind.v1.UnBindIdResponse
	8,  // 13: bind.v1.BindService.DisconnectTarget:output_type -> bind.v1.DisconnectTargetResponse
	10, // 14: bind.v1.BindService.BindProxyTarget:output_type -> bind.v1.ProxyTargetResponse
	10, // 15: bind.v1.BindService.UnbindProxyTarget:output_type -> bind.v1.ProxyTargetResponse
	12, // 16: bind.v1.BindService.TargetBindId:output_type -> bind.v1.TargetBindIdResponse
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_bind_v1_bind_proto_init() }
func file_bind_v1_bind_proto_init() {
	if File_bind_v1_bind_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bind_v1_bind_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_bind_v1_bind_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindIdRequest); i {
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
		file_bind_v1_bind_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindIdResponse); i {
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
		file_bind_v1_bind_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindExistRequest); i {
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
		file_bind_v1_bind_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindExistResponse); i {
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
		file_bind_v1_bind_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnBindIdRequest); i {
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
		file_bind_v1_bind_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnBindIdResponse); i {
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
		file_bind_v1_bind_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectTargetRequest); i {
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
		file_bind_v1_bind_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisconnectTargetResponse); i {
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
		file_bind_v1_bind_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyTargetRequest); i {
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
		file_bind_v1_bind_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyTargetResponse); i {
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
		file_bind_v1_bind_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetBindIdRequest); i {
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
		file_bind_v1_bind_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetBindIdResponse); i {
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
			RawDescriptor: file_bind_v1_bind_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bind_v1_bind_proto_goTypes,
		DependencyIndexes: file_bind_v1_bind_proto_depIdxs,
		MessageInfos:      file_bind_v1_bind_proto_msgTypes,
	}.Build()
	File_bind_v1_bind_proto = out.File
	file_bind_v1_bind_proto_rawDesc = nil
	file_bind_v1_bind_proto_goTypes = nil
	file_bind_v1_bind_proto_depIdxs = nil
}
