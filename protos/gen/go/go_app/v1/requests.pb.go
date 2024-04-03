// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: go_app/v1/requests.proto

package go_appv1

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

// Do not delete these unless you know what you're doing
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_app_v1_requests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_go_app_v1_requests_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_go_app_v1_requests_proto_rawDescGZIP(), []int{0}
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit   int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset  int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	OrderBy string `protobuf:"bytes,3,opt,name=orderBy,proto3" json:"orderBy,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_app_v1_requests_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_app_v1_requests_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_go_app_v1_requests_proto_rawDescGZIP(), []int{1}
}

func (x *ListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_app_v1_requests_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_app_v1_requests_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_go_app_v1_requests_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []string `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_app_v1_requests_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_app_v1_requests_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_go_app_v1_requests_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_app_v1_requests_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_app_v1_requests_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_go_app_v1_requests_proto_rawDescGZIP(), []int{4}
}

func (x *GetRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

// Below here are the messages that aren't generic
type Hellos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hellos []*Hello `protobuf:"bytes,1,rep,name=hellos,proto3" json:"hellos,omitempty"`
}

func (x *Hellos) Reset() {
	*x = Hellos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_app_v1_requests_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hellos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hellos) ProtoMessage() {}

func (x *Hellos) ProtoReflect() protoreflect.Message {
	mi := &file_go_app_v1_requests_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hellos.ProtoReflect.Descriptor instead.
func (*Hellos) Descriptor() ([]byte, []int) {
	return file_go_app_v1_requests_proto_rawDescGZIP(), []int{5}
}

func (x *Hellos) GetHellos() []*Hello {
	if x != nil {
		return x.Hellos
	}
	return nil
}

type Todos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Todo `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Todos) Reset() {
	*x = Todos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_app_v1_requests_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todos) ProtoMessage() {}

func (x *Todos) ProtoReflect() protoreflect.Message {
	mi := &file_go_app_v1_requests_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todos.ProtoReflect.Descriptor instead.
func (*Todos) Descriptor() ([]byte, []int) {
	return file_go_app_v1_requests_proto_rawDescGZIP(), []int{6}
}

func (x *Todos) GetData() []*Todo {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpsertHellosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hellos []*Hello `protobuf:"bytes,1,rep,name=hellos,proto3" json:"hellos,omitempty"`
	Fields []string `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *UpsertHellosRequest) Reset() {
	*x = UpsertHellosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_app_v1_requests_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertHellosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertHellosRequest) ProtoMessage() {}

func (x *UpsertHellosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_app_v1_requests_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertHellosRequest.ProtoReflect.Descriptor instead.
func (*UpsertHellosRequest) Descriptor() ([]byte, []int) {
	return file_go_app_v1_requests_proto_rawDescGZIP(), []int{7}
}

func (x *UpsertHellosRequest) GetHellos() []*Hello {
	if x != nil {
		return x.Hellos
	}
	return nil
}

func (x *UpsertHellosRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

type UpsertTodosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   []*Todo  `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Fields []string `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *UpsertTodosRequest) Reset() {
	*x = UpsertTodosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_app_v1_requests_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertTodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertTodosRequest) ProtoMessage() {}

func (x *UpsertTodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_app_v1_requests_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertTodosRequest.ProtoReflect.Descriptor instead.
func (*UpsertTodosRequest) Descriptor() ([]byte, []int) {
	return file_go_app_v1_requests_proto_rawDescGZIP(), []int{8}
}

func (x *UpsertTodosRequest) GetData() []*Todo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UpsertTodosRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

type DeleteTodosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteTodosRequest) Reset() {
	*x = DeleteTodosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_app_v1_requests_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodosRequest) ProtoMessage() {}

func (x *DeleteTodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_app_v1_requests_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodosRequest.ProtoReflect.Descriptor instead.
func (*DeleteTodosRequest) Descriptor() ([]byte, []int) {
	return file_go_app_v1_requests_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteTodosRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeleteTodosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors []string `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *DeleteTodosResponse) Reset() {
	*x = DeleteTodosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_app_v1_requests_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTodosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodosResponse) ProtoMessage() {}

func (x *DeleteTodosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_go_app_v1_requests_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodosResponse.ProtoReflect.Descriptor instead.
func (*DeleteTodosResponse) Descriptor() ([]byte, []int) {
	return file_go_app_v1_requests_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteTodosResponse) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

type ListTodosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit   int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset  int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	OrderBy string `protobuf:"bytes,3,opt,name=orderBy,proto3" json:"orderBy,omitempty"`
}

func (x *ListTodosRequest) Reset() {
	*x = ListTodosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_app_v1_requests_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTodosRequest) ProtoMessage() {}

func (x *ListTodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_app_v1_requests_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTodosRequest.ProtoReflect.Descriptor instead.
func (*ListTodosRequest) Descriptor() ([]byte, []int) {
	return file_go_app_v1_requests_proto_rawDescGZIP(), []int{11}
}

func (x *ListTodosRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTodosRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListTodosRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

type GetTodosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetTodosRequest) Reset() {
	*x = GetTodosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_app_v1_requests_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTodosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodosRequest) ProtoMessage() {}

func (x *GetTodosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_go_app_v1_requests_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodosRequest.ProtoReflect.Descriptor instead.
func (*GetTodosRequest) Descriptor() ([]byte, []int) {
	return file_go_app_v1_requests_proto_rawDescGZIP(), []int{12}
}

func (x *GetTodosRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

var File_go_app_v1_requests_proto protoreflect.FileDescriptor

var file_go_app_v1_requests_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x6f, 0x5f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x6f, 0x5f, 0x61,
	0x70, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x15, 0x67, 0x6f, 0x5f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x6f,
	0x5f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x55, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x79, 0x22, 0x21, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x28, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x22, 0x1e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0x32, 0x0a, 0x06, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x5f,
	0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x06, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x73, 0x22, 0x2c, 0x0a, 0x05, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x12, 0x23, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x6f,
	0x5f, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x57, 0x0a, 0x13, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x5f, 0x61,
	0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x06, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x51, 0x0a, 0x12, 0x55,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x67, 0x6f, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x26,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x2d, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x22, 0x5a, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64,
	0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x42, 0x9b, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x63, 0x63, 0x34, 0x62, 0x33, 0x72, 0x2f, 0x67, 0x6f,
	0x5f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x67, 0x6f, 0x5f, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x6f, 0x5f,
	0x61, 0x70, 0x70, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x47, 0x6f,
	0x41, 0x70, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x47, 0x6f, 0x41, 0x70, 0x70, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x14, 0x47, 0x6f, 0x41, 0x70, 0x70, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x47, 0x6f, 0x41, 0x70, 0x70,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_app_v1_requests_proto_rawDescOnce sync.Once
	file_go_app_v1_requests_proto_rawDescData = file_go_app_v1_requests_proto_rawDesc
)

func file_go_app_v1_requests_proto_rawDescGZIP() []byte {
	file_go_app_v1_requests_proto_rawDescOnce.Do(func() {
		file_go_app_v1_requests_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_app_v1_requests_proto_rawDescData)
	})
	return file_go_app_v1_requests_proto_rawDescData
}

var file_go_app_v1_requests_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_go_app_v1_requests_proto_goTypes = []interface{}{
	(*Empty)(nil),               // 0: go_app.v1.Empty
	(*ListRequest)(nil),         // 1: go_app.v1.ListRequest
	(*DeleteRequest)(nil),       // 2: go_app.v1.DeleteRequest
	(*DeleteResponse)(nil),      // 3: go_app.v1.DeleteResponse
	(*GetRequest)(nil),          // 4: go_app.v1.GetRequest
	(*Hellos)(nil),              // 5: go_app.v1.Hellos
	(*Todos)(nil),               // 6: go_app.v1.Todos
	(*UpsertHellosRequest)(nil), // 7: go_app.v1.UpsertHellosRequest
	(*UpsertTodosRequest)(nil),  // 8: go_app.v1.UpsertTodosRequest
	(*DeleteTodosRequest)(nil),  // 9: go_app.v1.DeleteTodosRequest
	(*DeleteTodosResponse)(nil), // 10: go_app.v1.DeleteTodosResponse
	(*ListTodosRequest)(nil),    // 11: go_app.v1.ListTodosRequest
	(*GetTodosRequest)(nil),     // 12: go_app.v1.GetTodosRequest
	(*Hello)(nil),               // 13: go_app.v1.Hello
	(*Todo)(nil),                // 14: go_app.v1.Todo
}
var file_go_app_v1_requests_proto_depIdxs = []int32{
	13, // 0: go_app.v1.Hellos.hellos:type_name -> go_app.v1.Hello
	14, // 1: go_app.v1.Todos.data:type_name -> go_app.v1.Todo
	13, // 2: go_app.v1.UpsertHellosRequest.hellos:type_name -> go_app.v1.Hello
	14, // 3: go_app.v1.UpsertTodosRequest.data:type_name -> go_app.v1.Todo
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_go_app_v1_requests_proto_init() }
func file_go_app_v1_requests_proto_init() {
	if File_go_app_v1_requests_proto != nil {
		return
	}
	file_go_app_v1_types_proto_init()
	file_go_app_v1_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_go_app_v1_requests_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_go_app_v1_requests_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_go_app_v1_requests_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_go_app_v1_requests_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_go_app_v1_requests_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_go_app_v1_requests_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hellos); i {
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
		file_go_app_v1_requests_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todos); i {
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
		file_go_app_v1_requests_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertHellosRequest); i {
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
		file_go_app_v1_requests_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertTodosRequest); i {
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
		file_go_app_v1_requests_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTodosRequest); i {
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
		file_go_app_v1_requests_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTodosResponse); i {
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
		file_go_app_v1_requests_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTodosRequest); i {
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
		file_go_app_v1_requests_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTodosRequest); i {
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
			RawDescriptor: file_go_app_v1_requests_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_app_v1_requests_proto_goTypes,
		DependencyIndexes: file_go_app_v1_requests_proto_depIdxs,
		MessageInfos:      file_go_app_v1_requests_proto_msgTypes,
	}.Build()
	File_go_app_v1_requests_proto = out.File
	file_go_app_v1_requests_proto_rawDesc = nil
	file_go_app_v1_requests_proto_goTypes = nil
	file_go_app_v1_requests_proto_depIdxs = nil
}
