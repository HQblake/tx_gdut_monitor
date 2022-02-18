// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: manage2send.proto

package sendpb

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

type Type int32

const (
	Type_EMAIL    Type = 0
	Type_KAFKA    Type = 1
	Type_NSQ      Type = 2
	Type_HTTP_LOG Type = 3
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "EMAIL",
		1: "KAFKA",
		2: "NSQ",
		3: "HTTP_LOG",
	}
	Type_value = map[string]int32{
		"EMAIL":    0,
		"KAFKA":    1,
		"NSQ":      2,
		"HTTP_LOG": 3,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_manage2send_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_manage2send_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_manage2send_proto_rawDescGZIP(), []int{0}
}

type SendResponse_ResponseCode int32

const (
	SendResponse_SUCCESS SendResponse_ResponseCode = 0
	SendResponse_ERROR   SendResponse_ResponseCode = 1
)

// Enum value maps for SendResponse_ResponseCode.
var (
	SendResponse_ResponseCode_name = map[int32]string{
		0: "SUCCESS",
		1: "ERROR",
	}
	SendResponse_ResponseCode_value = map[string]int32{
		"SUCCESS": 0,
		"ERROR":   1,
	}
)

func (x SendResponse_ResponseCode) Enum() *SendResponse_ResponseCode {
	p := new(SendResponse_ResponseCode)
	*p = x
	return p
}

func (x SendResponse_ResponseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SendResponse_ResponseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_manage2send_proto_enumTypes[1].Descriptor()
}

func (SendResponse_ResponseCode) Type() protoreflect.EnumType {
	return &file_manage2send_proto_enumTypes[1]
}

func (x SendResponse_ResponseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SendResponse_ResponseCode.Descriptor instead.
func (SendResponse_ResponseCode) EnumDescriptor() ([]byte, []int) {
	return file_manage2send_proto_rawDescGZIP(), []int{8, 0}
}

// 发送配置，包括配置类型，配置内容及配置等级
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendType Type   `protobuf:"varint,1,opt,name=SendType,proto3,enum=api.Type" json:"SendType,omitempty"`
	Config   string `protobuf:"bytes,2,opt,name=Config,proto3" json:"Config,omitempty"`
	Level    int32  `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage2send_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_manage2send_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_manage2send_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetSendType() Type {
	if x != nil {
		return x.SendType
	}
	return Type_EMAIL
}

func (x *Config) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Config) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type ConfigEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigID int32   `protobuf:"varint,1,opt,name=ConfigID,proto3" json:"ConfigID,omitempty"`
	Conf     *Config `protobuf:"bytes,2,opt,name=Conf,proto3" json:"Conf,omitempty"`
}

func (x *ConfigEntry) Reset() {
	*x = ConfigEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage2send_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigEntry) ProtoMessage() {}

func (x *ConfigEntry) ProtoReflect() protoreflect.Message {
	mi := &file_manage2send_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigEntry.ProtoReflect.Descriptor instead.
func (*ConfigEntry) Descriptor() ([]byte, []int) {
	return file_manage2send_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigEntry) GetConfigID() int32 {
	if x != nil {
		return x.ConfigID
	}
	return 0
}

func (x *ConfigEntry) GetConf() *Config {
	if x != nil {
		return x.Conf
	}
	return nil
}

type InitConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IP     string         `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"` // ip和local组成agent的唯一标识
	Local  string         `protobuf:"bytes,2,opt,name=Local,proto3" json:"Local,omitempty"`
	Config []*ConfigEntry `protobuf:"bytes,3,rep,name=Config,proto3" json:"Config,omitempty"`
}

func (x *InitConfig) Reset() {
	*x = InitConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage2send_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitConfig) ProtoMessage() {}

func (x *InitConfig) ProtoReflect() protoreflect.Message {
	mi := &file_manage2send_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitConfig.ProtoReflect.Descriptor instead.
func (*InitConfig) Descriptor() ([]byte, []int) {
	return file_manage2send_proto_rawDescGZIP(), []int{2}
}

func (x *InitConfig) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *InitConfig) GetLocal() string {
	if x != nil {
		return x.Local
	}
	return ""
}

func (x *InitConfig) GetConfig() []*ConfigEntry {
	if x != nil {
		return x.Config
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IP     string       `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"` // ip和local组成agent的唯一标识
	Local  string       `protobuf:"bytes,2,opt,name=Local,proto3" json:"Local,omitempty"`
	Config *ConfigEntry `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage2send_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manage2send_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_manage2send_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequest) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *UpdateRequest) GetLocal() string {
	if x != nil {
		return x.Local
	}
	return ""
}

func (x *UpdateRequest) GetConfig() *ConfigEntry {
	if x != nil {
		return x.Config
	}
	return nil
}

type DelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IP       string `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"` // ip和local组成agent的唯一标识
	Local    string `protobuf:"bytes,2,opt,name=Local,proto3" json:"Local,omitempty"`
	ConfigID int32  `protobuf:"varint,3,opt,name=ConfigID,proto3" json:"ConfigID,omitempty"`
}

func (x *DelRequest) Reset() {
	*x = DelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage2send_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelRequest) ProtoMessage() {}

func (x *DelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manage2send_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelRequest.ProtoReflect.Descriptor instead.
func (*DelRequest) Descriptor() ([]byte, []int) {
	return file_manage2send_proto_rawDescGZIP(), []int{4}
}

func (x *DelRequest) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *DelRequest) GetLocal() string {
	if x != nil {
		return x.Local
	}
	return ""
}

func (x *DelRequest) GetConfigID() int32 {
	if x != nil {
		return x.ConfigID
	}
	return 0
}

// 以agentID为键值，返回配置列表
type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config map[string]*InitConfig `protobuf:"bytes,1,rep,name=Config,proto3" json:"Config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage2send_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manage2send_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_manage2send_proto_rawDescGZIP(), []int{5}
}

func (x *InitRequest) GetConfig() map[string]*InitConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IP     string       `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"` // ip和local组成agent的唯一标识
	Local  string       `protobuf:"bytes,2,opt,name=Local,proto3" json:"Local,omitempty"`
	Config *CheckConfig `protobuf:"bytes,3,opt,name=Config,proto3" json:"Config,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage2send_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manage2send_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_manage2send_proto_rawDescGZIP(), []int{6}
}

func (x *CheckRequest) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *CheckRequest) GetLocal() string {
	if x != nil {
		return x.Local
	}
	return ""
}

func (x *CheckRequest) GetConfig() *CheckConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type CheckConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SendType Type   `protobuf:"varint,1,opt,name=SendType,proto3,enum=api.Type" json:"SendType,omitempty"`
	Config   string `protobuf:"bytes,2,opt,name=Config,proto3" json:"Config,omitempty"`
}

func (x *CheckConfig) Reset() {
	*x = CheckConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage2send_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckConfig) ProtoMessage() {}

func (x *CheckConfig) ProtoReflect() protoreflect.Message {
	mi := &file_manage2send_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckConfig.ProtoReflect.Descriptor instead.
func (*CheckConfig) Descriptor() ([]byte, []int) {
	return file_manage2send_proto_rawDescGZIP(), []int{7}
}

func (x *CheckConfig) GetSendType() Type {
	if x != nil {
		return x.SendType
	}
	return Type_EMAIL
}

func (x *CheckConfig) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

type SendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code SendResponse_ResponseCode `protobuf:"varint,1,opt,name=Code,proto3,enum=api.SendResponse_ResponseCode" json:"Code,omitempty"`
	Msg  string                    `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *SendResponse) Reset() {
	*x = SendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage2send_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendResponse) ProtoMessage() {}

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manage2send_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendResponse.ProtoReflect.Descriptor instead.
func (*SendResponse) Descriptor() ([]byte, []int) {
	return file_manage2send_proto_rawDescGZIP(), []int{8}
}

func (x *SendResponse) GetCode() SendResponse_ResponseCode {
	if x != nil {
		return x.Code
	}
	return SendResponse_SUCCESS
}

func (x *SendResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_manage2send_proto protoreflect.FileDescriptor

var file_manage2send_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x32, 0x73, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x5d, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x25, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x4a, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x49, 0x44, 0x12, 0x1f, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x43,
	0x6f, 0x6e, 0x66, 0x22, 0x5c, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x50, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x5f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x50, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x4e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50,
	0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x49, 0x44, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x4a, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49,
	0x6e, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x5e, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x50, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0x4c, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x22, 0x7c, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x26, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01,
	0x2a, 0x33, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49,
	0x4c, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4b, 0x41, 0x46, 0x4b, 0x41, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x4e, 0x53, 0x51, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x54, 0x54, 0x50, 0x5f,
	0x4c, 0x4f, 0x47, 0x10, 0x03, 0x32, 0xc2, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x03, 0x44, 0x65, 0x6c, 0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b,
	0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x6e, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69,
	0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x65, 0x6b, 0x65, 0x47, 0x69, 0x74, 0x65,
	0x65, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x74, 0x78, 0x5f, 0x67, 0x64, 0x75, 0x74, 0x5f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2d,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73,
	0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x3b, 0x73, 0x65, 0x6e, 0x64,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manage2send_proto_rawDescOnce sync.Once
	file_manage2send_proto_rawDescData = file_manage2send_proto_rawDesc
)

func file_manage2send_proto_rawDescGZIP() []byte {
	file_manage2send_proto_rawDescOnce.Do(func() {
		file_manage2send_proto_rawDescData = protoimpl.X.CompressGZIP(file_manage2send_proto_rawDescData)
	})
	return file_manage2send_proto_rawDescData
}

var file_manage2send_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_manage2send_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_manage2send_proto_goTypes = []interface{}{
	(Type)(0),                      // 0: api.Type
	(SendResponse_ResponseCode)(0), // 1: api.SendResponse.ResponseCode
	(*Config)(nil),                 // 2: api.Config
	(*ConfigEntry)(nil),            // 3: api.ConfigEntry
	(*InitConfig)(nil),             // 4: api.InitConfig
	(*UpdateRequest)(nil),          // 5: api.UpdateRequest
	(*DelRequest)(nil),             // 6: api.DelRequest
	(*InitRequest)(nil),            // 7: api.InitRequest
	(*CheckRequest)(nil),           // 8: api.CheckRequest
	(*CheckConfig)(nil),            // 9: api.CheckConfig
	(*SendResponse)(nil),           // 10: api.SendResponse
	nil,                            // 11: api.InitRequest.ConfigEntry
}
var file_manage2send_proto_depIdxs = []int32{
	0,  // 0: api.Config.SendType:type_name -> api.Type
	2,  // 1: api.ConfigEntry.Conf:type_name -> api.Config
	3,  // 2: api.InitConfig.Config:type_name -> api.ConfigEntry
	3,  // 3: api.UpdateRequest.config:type_name -> api.ConfigEntry
	11, // 4: api.InitRequest.Config:type_name -> api.InitRequest.ConfigEntry
	9,  // 5: api.CheckRequest.Config:type_name -> api.CheckConfig
	0,  // 6: api.CheckConfig.SendType:type_name -> api.Type
	1,  // 7: api.SendResponse.Code:type_name -> api.SendResponse.ResponseCode
	4,  // 8: api.InitRequest.ConfigEntry.value:type_name -> api.InitConfig
	5,  // 9: api.SendService.Set:input_type -> api.UpdateRequest
	6,  // 10: api.SendService.Del:input_type -> api.DelRequest
	7,  // 11: api.SendService.Init:input_type -> api.InitRequest
	8,  // 12: api.SendService.Check:input_type -> api.CheckRequest
	10, // 13: api.SendService.Set:output_type -> api.SendResponse
	10, // 14: api.SendService.Del:output_type -> api.SendResponse
	10, // 15: api.SendService.Init:output_type -> api.SendResponse
	10, // 16: api.SendService.Check:output_type -> api.SendResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_manage2send_proto_init() }
func file_manage2send_proto_init() {
	if File_manage2send_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manage2send_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_manage2send_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigEntry); i {
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
		file_manage2send_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitConfig); i {
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
		file_manage2send_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_manage2send_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelRequest); i {
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
		file_manage2send_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
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
		file_manage2send_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRequest); i {
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
		file_manage2send_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckConfig); i {
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
		file_manage2send_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendResponse); i {
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
			RawDescriptor: file_manage2send_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_manage2send_proto_goTypes,
		DependencyIndexes: file_manage2send_proto_depIdxs,
		EnumInfos:         file_manage2send_proto_enumTypes,
		MessageInfos:      file_manage2send_proto_msgTypes,
	}.Build()
	File_manage2send_proto = out.File
	file_manage2send_proto_rawDesc = nil
	file_manage2send_proto_goTypes = nil
	file_manage2send_proto_depIdxs = nil
}
