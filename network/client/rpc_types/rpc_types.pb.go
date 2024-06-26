// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: rpc_types.proto

package rpc_types

import (
	field_options "bnet-mock/network/client/global_extensions/field_options"
	method_options "bnet-mock/network/client/global_extensions/method_options"
	service_options "bnet-mock/network/client/global_extensions/service_options"
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

// Symbols defined in public import of global_extensions/method_options.proto.

var E_MethodId = method_options.E_MethodId

// Symbols defined in public import of global_extensions/service_options.proto.

var E_OriginalFullyQualifiedDescriptorName = service_options.E_OriginalFullyQualifiedDescriptorName
var E_ServiceId = service_options.E_ServiceId

// Symbols defined in public import of global_extensions/field_options.proto.

type LogOption = field_options.LogOption

const LogOption_HIDDEN = field_options.LogOption_HIDDEN
const LogOption_HEX = field_options.LogOption_HEX

var LogOption_name = field_options.LogOption_name
var LogOption_value = field_options.LogOption_value
var E_Log = field_options.E_Log

type NO_RESPONSE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NO_RESPONSE) Reset() {
	*x = NO_RESPONSE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NO_RESPONSE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NO_RESPONSE) ProtoMessage() {}

func (x *NO_RESPONSE) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NO_RESPONSE.ProtoReflect.Descriptor instead.
func (*NO_RESPONSE) Descriptor() ([]byte, []int) {
	return file_rpc_types_proto_rawDescGZIP(), []int{0}
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *string `protobuf:"bytes,1,req,name=address" json:"address,omitempty"`
	Port    *uint32 `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_rpc_types_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *Address) GetPort() uint32 {
	if x != nil && x.Port != nil {
		return *x.Port
	}
	return 0
}

type ProcessId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label *uint32 `protobuf:"varint,1,req,name=label" json:"label,omitempty"`
	Epoch *uint32 `protobuf:"varint,2,req,name=epoch" json:"epoch,omitempty"`
}

func (x *ProcessId) Reset() {
	*x = ProcessId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessId) ProtoMessage() {}

func (x *ProcessId) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessId.ProtoReflect.Descriptor instead.
func (*ProcessId) Descriptor() ([]byte, []int) {
	return file_rpc_types_proto_rawDescGZIP(), []int{2}
}

func (x *ProcessId) GetLabel() uint32 {
	if x != nil && x.Label != nil {
		return *x.Label
	}
	return 0
}

func (x *ProcessId) GetEpoch() uint32 {
	if x != nil && x.Epoch != nil {
		return *x.Epoch
	}
	return 0
}

type ObjectAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host     *ProcessId `protobuf:"bytes,1,req,name=host" json:"host,omitempty"`
	ObjectId *uint64    `protobuf:"varint,2,opt,name=object_id,json=objectId,def=0" json:"object_id,omitempty"`
}

// Default values for ObjectAddress fields.
const (
	Default_ObjectAddress_ObjectId = uint64(0)
)

func (x *ObjectAddress) Reset() {
	*x = ObjectAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectAddress) ProtoMessage() {}

func (x *ObjectAddress) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectAddress.ProtoReflect.Descriptor instead.
func (*ObjectAddress) Descriptor() ([]byte, []int) {
	return file_rpc_types_proto_rawDescGZIP(), []int{3}
}

func (x *ObjectAddress) GetHost() *ProcessId {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *ObjectAddress) GetObjectId() uint64 {
	if x != nil && x.ObjectId != nil {
		return *x.ObjectId
	}
	return Default_ObjectAddress_ObjectId
}

type NoData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoData) Reset() {
	*x = NoData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoData) ProtoMessage() {}

func (x *NoData) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoData.ProtoReflect.Descriptor instead.
func (*NoData) Descriptor() ([]byte, []int) {
	return file_rpc_types_proto_rawDescGZIP(), []int{4}
}

type ErrorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectAddress *ObjectAddress `protobuf:"bytes,1,req,name=object_address,json=objectAddress" json:"object_address,omitempty"`
	Status        *uint32        `protobuf:"varint,2,req,name=status" json:"status,omitempty"`
	ServiceHash   *uint32        `protobuf:"varint,3,req,name=service_hash,json=serviceHash" json:"service_hash,omitempty"`
	MethodId      *uint32        `protobuf:"varint,4,req,name=method_id,json=methodId" json:"method_id,omitempty"`
}

func (x *ErrorInfo) Reset() {
	*x = ErrorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorInfo) ProtoMessage() {}

func (x *ErrorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorInfo.ProtoReflect.Descriptor instead.
func (*ErrorInfo) Descriptor() ([]byte, []int) {
	return file_rpc_types_proto_rawDescGZIP(), []int{5}
}

func (x *ErrorInfo) GetObjectAddress() *ObjectAddress {
	if x != nil {
		return x.ObjectAddress
	}
	return nil
}

func (x *ErrorInfo) GetStatus() uint32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *ErrorInfo) GetServiceHash() uint32 {
	if x != nil && x.ServiceHash != nil {
		return *x.ServiceHash
	}
	return 0
}

func (x *ErrorInfo) GetMethodId() uint32 {
	if x != nil && x.MethodId != nil {
		return *x.MethodId
	}
	return 0
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceId      *uint32      `protobuf:"varint,1,req,name=service_id,json=serviceId" json:"service_id,omitempty"`
	MethodId       *uint32      `protobuf:"varint,2,opt,name=method_id,json=methodId" json:"method_id,omitempty"`
	Token          *uint32      `protobuf:"varint,3,req,name=token" json:"token,omitempty"`
	ObjectId       *uint64      `protobuf:"varint,4,opt,name=object_id,json=objectId,def=0" json:"object_id,omitempty"`
	Size           *uint32      `protobuf:"varint,5,opt,name=size,def=0" json:"size,omitempty"`
	Status         *uint32      `protobuf:"varint,6,opt,name=status,def=0" json:"status,omitempty"`
	Error          []*ErrorInfo `protobuf:"bytes,7,rep,name=error" json:"error,omitempty"`
	Timeout        *uint64      `protobuf:"varint,8,opt,name=timeout" json:"timeout,omitempty"`
	IsResponse     *bool        `protobuf:"varint,9,opt,name=is_response,json=isResponse" json:"is_response,omitempty"`
	ForwardTargets []*ProcessId `protobuf:"bytes,10,rep,name=forward_targets,json=forwardTargets" json:"forward_targets,omitempty"`
	ServiceHash    *uint32      `protobuf:"fixed32,11,opt,name=service_hash,json=serviceHash" json:"service_hash,omitempty"`
}

// Default values for Header fields.
const (
	Default_Header_ObjectId = uint64(0)
	Default_Header_Size     = uint32(0)
	Default_Header_Status   = uint32(0)
)

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_rpc_types_proto_rawDescGZIP(), []int{6}
}

func (x *Header) GetServiceId() uint32 {
	if x != nil && x.ServiceId != nil {
		return *x.ServiceId
	}
	return 0
}

func (x *Header) GetMethodId() uint32 {
	if x != nil && x.MethodId != nil {
		return *x.MethodId
	}
	return 0
}

func (x *Header) GetToken() uint32 {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return 0
}

func (x *Header) GetObjectId() uint64 {
	if x != nil && x.ObjectId != nil {
		return *x.ObjectId
	}
	return Default_Header_ObjectId
}

func (x *Header) GetSize() uint32 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return Default_Header_Size
}

func (x *Header) GetStatus() uint32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return Default_Header_Status
}

func (x *Header) GetError() []*ErrorInfo {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *Header) GetTimeout() uint64 {
	if x != nil && x.Timeout != nil {
		return *x.Timeout
	}
	return 0
}

func (x *Header) GetIsResponse() bool {
	if x != nil && x.IsResponse != nil {
		return *x.IsResponse
	}
	return false
}

func (x *Header) GetForwardTargets() []*ProcessId {
	if x != nil {
		return x.ForwardTargets
	}
	return nil
}

func (x *Header) GetServiceHash() uint32 {
	if x != nil && x.ServiceHash != nil {
		return *x.ServiceHash
	}
	return 0
}

var File_rpc_types_proto protoreflect.FileDescriptor

var file_rpc_types_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a,
	0x26, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x25, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x4e, 0x4f, 0x5f, 0x52, 0x45,
	0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x22, 0x37, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22,
	0x37, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x02, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x22, 0x5c, 0x0a, 0x0d, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x3a, 0x01, 0x30, 0x52, 0x08, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x08, 0x0a, 0x06, 0x4e, 0x6f, 0x44, 0x61, 0x74, 0x61,
	0x22, 0xa7, 0x01, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42,
	0x0a, 0x0e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1b, 0x0a,
	0x09, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x22, 0xfb, 0x02, 0x0a, 0x06, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x3a, 0x01, 0x30, 0x52, 0x08, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x01, 0x30, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x19,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x3a, 0x01,
	0x30, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62,
	0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x49, 0x64, 0x52, 0x0e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x48, 0x61, 0x73, 0x68, 0x42, 0x3f, 0x0a, 0x0d, 0x62, 0x6e, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x08, 0x52, 0x70, 0x63, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x48, 0x01, 0x5a, 0x22, 0x62, 0x6e, 0x65, 0x74, 0x2d, 0x6d, 0x6f, 0x63, 0x6b,
	0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x50, 0x00, 0x50, 0x01, 0x50, 0x02,
}

var (
	file_rpc_types_proto_rawDescOnce sync.Once
	file_rpc_types_proto_rawDescData = file_rpc_types_proto_rawDesc
)

func file_rpc_types_proto_rawDescGZIP() []byte {
	file_rpc_types_proto_rawDescOnce.Do(func() {
		file_rpc_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_types_proto_rawDescData)
	})
	return file_rpc_types_proto_rawDescData
}

var file_rpc_types_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rpc_types_proto_goTypes = []interface{}{
	(*NO_RESPONSE)(nil),   // 0: bgs.protocol.NO_RESPONSE
	(*Address)(nil),       // 1: bgs.protocol.Address
	(*ProcessId)(nil),     // 2: bgs.protocol.ProcessId
	(*ObjectAddress)(nil), // 3: bgs.protocol.ObjectAddress
	(*NoData)(nil),        // 4: bgs.protocol.NoData
	(*ErrorInfo)(nil),     // 5: bgs.protocol.ErrorInfo
	(*Header)(nil),        // 6: bgs.protocol.Header
}
var file_rpc_types_proto_depIdxs = []int32{
	2, // 0: bgs.protocol.ObjectAddress.host:type_name -> bgs.protocol.ProcessId
	3, // 1: bgs.protocol.ErrorInfo.object_address:type_name -> bgs.protocol.ObjectAddress
	5, // 2: bgs.protocol.Header.error:type_name -> bgs.protocol.ErrorInfo
	2, // 3: bgs.protocol.Header.forward_targets:type_name -> bgs.protocol.ProcessId
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_types_proto_init() }
func file_rpc_types_proto_init() {
	if File_rpc_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NO_RESPONSE); i {
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
		file_rpc_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_rpc_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessId); i {
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
		file_rpc_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectAddress); i {
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
		file_rpc_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoData); i {
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
		file_rpc_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorInfo); i {
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
		file_rpc_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
			RawDescriptor: file_rpc_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_types_proto_goTypes,
		DependencyIndexes: file_rpc_types_proto_depIdxs,
		MessageInfos:      file_rpc_types_proto_msgTypes,
	}.Build()
	File_rpc_types_proto = out.File
	file_rpc_types_proto_rawDesc = nil
	file_rpc_types_proto_goTypes = nil
	file_rpc_types_proto_depIdxs = nil
}
