// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: gnomobiletypes.proto

package rpc

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

// messages
type SetRemote_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Remote string `protobuf:"bytes,1,opt,name=Remote,proto3" json:"Remote,omitempty"`
}

func (x *SetRemote_Request) Reset() {
	*x = SetRemote_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gnomobiletypes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRemote_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRemote_Request) ProtoMessage() {}

func (x *SetRemote_Request) ProtoReflect() protoreflect.Message {
	mi := &file_gnomobiletypes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRemote_Request.ProtoReflect.Descriptor instead.
func (*SetRemote_Request) Descriptor() ([]byte, []int) {
	return file_gnomobiletypes_proto_rawDescGZIP(), []int{0}
}

func (x *SetRemote_Request) GetRemote() string {
	if x != nil {
		return x.Remote
	}
	return ""
}

type SetRemote_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetRemote_Reply) Reset() {
	*x = SetRemote_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gnomobiletypes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRemote_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRemote_Reply) ProtoMessage() {}

func (x *SetRemote_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_gnomobiletypes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRemote_Reply.ProtoReflect.Descriptor instead.
func (*SetRemote_Reply) Descriptor() ([]byte, []int) {
	return file_gnomobiletypes_proto_rawDescGZIP(), []int{1}
}

type SetChainID_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainID string `protobuf:"bytes,1,opt,name=ChainID,proto3" json:"ChainID,omitempty"`
}

func (x *SetChainID_Request) Reset() {
	*x = SetChainID_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gnomobiletypes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetChainID_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetChainID_Request) ProtoMessage() {}

func (x *SetChainID_Request) ProtoReflect() protoreflect.Message {
	mi := &file_gnomobiletypes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetChainID_Request.ProtoReflect.Descriptor instead.
func (*SetChainID_Request) Descriptor() ([]byte, []int) {
	return file_gnomobiletypes_proto_rawDescGZIP(), []int{2}
}

func (x *SetChainID_Request) GetChainID() string {
	if x != nil {
		return x.ChainID
	}
	return ""
}

type SetChainID_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetChainID_Reply) Reset() {
	*x = SetChainID_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gnomobiletypes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetChainID_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetChainID_Reply) ProtoMessage() {}

func (x *SetChainID_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_gnomobiletypes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetChainID_Reply.ProtoReflect.Descriptor instead.
func (*SetChainID_Reply) Descriptor() ([]byte, []int) {
	return file_gnomobiletypes_proto_rawDescGZIP(), []int{3}
}

type SetNameOrBech32_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NameOrBech32 string `protobuf:"bytes,1,opt,name=NameOrBech32,proto3" json:"NameOrBech32,omitempty"`
}

func (x *SetNameOrBech32_Request) Reset() {
	*x = SetNameOrBech32_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gnomobiletypes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetNameOrBech32_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNameOrBech32_Request) ProtoMessage() {}

func (x *SetNameOrBech32_Request) ProtoReflect() protoreflect.Message {
	mi := &file_gnomobiletypes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNameOrBech32_Request.ProtoReflect.Descriptor instead.
func (*SetNameOrBech32_Request) Descriptor() ([]byte, []int) {
	return file_gnomobiletypes_proto_rawDescGZIP(), []int{4}
}

func (x *SetNameOrBech32_Request) GetNameOrBech32() string {
	if x != nil {
		return x.NameOrBech32
	}
	return ""
}

type SetNameOrBech32_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetNameOrBech32_Reply) Reset() {
	*x = SetNameOrBech32_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gnomobiletypes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetNameOrBech32_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNameOrBech32_Reply) ProtoMessage() {}

func (x *SetNameOrBech32_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_gnomobiletypes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNameOrBech32_Reply.ProtoReflect.Descriptor instead.
func (*SetNameOrBech32_Reply) Descriptor() ([]byte, []int) {
	return file_gnomobiletypes_proto_rawDescGZIP(), []int{5}
}

type SetPassword_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *SetPassword_Request) Reset() {
	*x = SetPassword_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gnomobiletypes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPassword_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPassword_Request) ProtoMessage() {}

func (x *SetPassword_Request) ProtoReflect() protoreflect.Message {
	mi := &file_gnomobiletypes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPassword_Request.ProtoReflect.Descriptor instead.
func (*SetPassword_Request) Descriptor() ([]byte, []int) {
	return file_gnomobiletypes_proto_rawDescGZIP(), []int{6}
}

func (x *SetPassword_Request) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SetPassword_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetPassword_Reply) Reset() {
	*x = SetPassword_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gnomobiletypes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPassword_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPassword_Reply) ProtoMessage() {}

func (x *SetPassword_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_gnomobiletypes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPassword_Reply.ProtoReflect.Descriptor instead.
func (*SetPassword_Reply) Descriptor() ([]byte, []int) {
	return file_gnomobiletypes_proto_rawDescGZIP(), []int{7}
}

type Query_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Query_Request) Reset() {
	*x = Query_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gnomobiletypes_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query_Request) ProtoMessage() {}

func (x *Query_Request) ProtoReflect() protoreflect.Message {
	mi := &file_gnomobiletypes_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query_Request.ProtoReflect.Descriptor instead.
func (*Query_Request) Descriptor() ([]byte, []int) {
	return file_gnomobiletypes_proto_rawDescGZIP(), []int{8}
}

func (x *Query_Request) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Query_Request) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Query_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []byte `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *Query_Reply) Reset() {
	*x = Query_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gnomobiletypes_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query_Reply) ProtoMessage() {}

func (x *Query_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_gnomobiletypes_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query_Reply.ProtoReflect.Descriptor instead.
func (*Query_Reply) Descriptor() ([]byte, []int) {
	return file_gnomobiletypes_proto_rawDescGZIP(), []int{9}
}

func (x *Query_Reply) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

type Call_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackagePath string   `protobuf:"bytes,1,opt,name=PackagePath,proto3" json:"PackagePath,omitempty"`
	Fnc         string   `protobuf:"bytes,2,opt,name=Fnc,proto3" json:"Fnc,omitempty"`
	Args        []string `protobuf:"bytes,3,rep,name=Args,proto3" json:"Args,omitempty"`
	GasFee      string   `protobuf:"bytes,4,opt,name=GasFee,proto3" json:"GasFee,omitempty"`
	GasWanted   int64    `protobuf:"zigzag64,5,opt,name=GasWanted,proto3" json:"GasWanted,omitempty"`
	Send        string   `protobuf:"bytes,6,opt,name=Send,proto3" json:"Send,omitempty"`
	Memo        string   `protobuf:"bytes,7,opt,name=Memo,proto3" json:"Memo,omitempty"`
}

func (x *Call_Request) Reset() {
	*x = Call_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gnomobiletypes_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Call_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Call_Request) ProtoMessage() {}

func (x *Call_Request) ProtoReflect() protoreflect.Message {
	mi := &file_gnomobiletypes_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Call_Request.ProtoReflect.Descriptor instead.
func (*Call_Request) Descriptor() ([]byte, []int) {
	return file_gnomobiletypes_proto_rawDescGZIP(), []int{10}
}

func (x *Call_Request) GetPackagePath() string {
	if x != nil {
		return x.PackagePath
	}
	return ""
}

func (x *Call_Request) GetFnc() string {
	if x != nil {
		return x.Fnc
	}
	return ""
}

func (x *Call_Request) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *Call_Request) GetGasFee() string {
	if x != nil {
		return x.GasFee
	}
	return ""
}

func (x *Call_Request) GetGasWanted() int64 {
	if x != nil {
		return x.GasWanted
	}
	return 0
}

func (x *Call_Request) GetSend() string {
	if x != nil {
		return x.Send
	}
	return ""
}

func (x *Call_Request) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

type Call_Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []byte `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *Call_Reply) Reset() {
	*x = Call_Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gnomobiletypes_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Call_Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Call_Reply) ProtoMessage() {}

func (x *Call_Reply) ProtoReflect() protoreflect.Message {
	mi := &file_gnomobiletypes_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Call_Reply.ProtoReflect.Descriptor instead.
func (*Call_Reply) Descriptor() ([]byte, []int) {
	return file_gnomobiletypes_proto_rawDescGZIP(), []int{11}
}

func (x *Call_Reply) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_gnomobiletypes_proto protoreflect.FileDescriptor

var file_gnomobiletypes_proto_rawDesc = []byte{
	0x0a, 0x14, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x6c, 0x61, 0x6e, 0x64, 0x2e, 0x67, 0x6e, 0x6f,
	0x2e, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x2b, 0x0a,
	0x11, 0x53, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x65,
	0x74, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2e, 0x0a,
	0x12, 0x53, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x22, 0x12, 0x0a,
	0x10, 0x53, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x3d, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x72, 0x42, 0x65,
	0x63, 0x68, 0x33, 0x32, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x72, 0x42, 0x65, 0x63, 0x68, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x72, 0x42, 0x65, 0x63, 0x68, 0x33, 0x32,
	0x22, 0x17, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x4f, 0x72, 0x42, 0x65, 0x63,
	0x68, 0x33, 0x32, 0x5f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x31, 0x0a, 0x13, 0x53, 0x65, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x13, 0x0a, 0x11,
	0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x37, 0x0a, 0x0d, 0x51, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x25, 0x0a, 0x0b, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0xb4, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x46, 0x6e, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x46, 0x6e, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x72, 0x67, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x41, 0x72, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x61,
	0x73, 0x46, 0x65, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x61, 0x73, 0x46,
	0x65, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x61, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x12, 0x52, 0x09, 0x47, 0x61, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x65, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x53, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x65, 0x6d, 0x6f, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4d, 0x65, 0x6d, 0x6f, 0x22, 0x24, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x6c,
	0x5f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x2a,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6e, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x6e, 0x6f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_gnomobiletypes_proto_rawDescOnce sync.Once
	file_gnomobiletypes_proto_rawDescData = file_gnomobiletypes_proto_rawDesc
)

func file_gnomobiletypes_proto_rawDescGZIP() []byte {
	file_gnomobiletypes_proto_rawDescOnce.Do(func() {
		file_gnomobiletypes_proto_rawDescData = protoimpl.X.CompressGZIP(file_gnomobiletypes_proto_rawDescData)
	})
	return file_gnomobiletypes_proto_rawDescData
}

var file_gnomobiletypes_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_gnomobiletypes_proto_goTypes = []interface{}{
	(*SetRemote_Request)(nil),       // 0: land.gno.gnomobile.v1.SetRemote_Request
	(*SetRemote_Reply)(nil),         // 1: land.gno.gnomobile.v1.SetRemote_Reply
	(*SetChainID_Request)(nil),      // 2: land.gno.gnomobile.v1.SetChainID_Request
	(*SetChainID_Reply)(nil),        // 3: land.gno.gnomobile.v1.SetChainID_Reply
	(*SetNameOrBech32_Request)(nil), // 4: land.gno.gnomobile.v1.SetNameOrBech32_Request
	(*SetNameOrBech32_Reply)(nil),   // 5: land.gno.gnomobile.v1.SetNameOrBech32_Reply
	(*SetPassword_Request)(nil),     // 6: land.gno.gnomobile.v1.SetPassword_Request
	(*SetPassword_Reply)(nil),       // 7: land.gno.gnomobile.v1.SetPassword_Reply
	(*Query_Request)(nil),           // 8: land.gno.gnomobile.v1.Query_Request
	(*Query_Reply)(nil),             // 9: land.gno.gnomobile.v1.Query_Reply
	(*Call_Request)(nil),            // 10: land.gno.gnomobile.v1.Call_Request
	(*Call_Reply)(nil),              // 11: land.gno.gnomobile.v1.Call_Reply
}
var file_gnomobiletypes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gnomobiletypes_proto_init() }
func file_gnomobiletypes_proto_init() {
	if File_gnomobiletypes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gnomobiletypes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRemote_Request); i {
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
		file_gnomobiletypes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRemote_Reply); i {
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
		file_gnomobiletypes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetChainID_Request); i {
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
		file_gnomobiletypes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetChainID_Reply); i {
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
		file_gnomobiletypes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetNameOrBech32_Request); i {
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
		file_gnomobiletypes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetNameOrBech32_Reply); i {
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
		file_gnomobiletypes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPassword_Request); i {
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
		file_gnomobiletypes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPassword_Reply); i {
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
		file_gnomobiletypes_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query_Request); i {
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
		file_gnomobiletypes_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query_Reply); i {
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
		file_gnomobiletypes_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Call_Request); i {
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
		file_gnomobiletypes_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Call_Reply); i {
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
			RawDescriptor: file_gnomobiletypes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gnomobiletypes_proto_goTypes,
		DependencyIndexes: file_gnomobiletypes_proto_depIdxs,
		MessageInfos:      file_gnomobiletypes_proto_msgTypes,
	}.Build()
	File_gnomobiletypes_proto = out.File
	file_gnomobiletypes_proto_rawDesc = nil
	file_gnomobiletypes_proto_goTypes = nil
	file_gnomobiletypes_proto_depIdxs = nil
}
