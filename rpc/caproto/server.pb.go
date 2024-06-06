// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: caproto/server.proto

package caproto

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

type ExtensionType int32

const (
	ExtensionType_EMPTY     ExtensionType = 0
	ExtensionType_STRING    ExtensionType = 1
	ExtensionType_INTEGER   ExtensionType = 2
	ExtensionType_BOOLEAN   ExtensionType = 3
	ExtensionType_BYTES     ExtensionType = 4
	ExtensionType_EXTENSION ExtensionType = 5
	ExtensionType_ARRAY     ExtensionType = 6
)

// Enum value maps for ExtensionType.
var (
	ExtensionType_name = map[int32]string{
		0: "EMPTY",
		1: "STRING",
		2: "INTEGER",
		3: "BOOLEAN",
		4: "BYTES",
		5: "EXTENSION",
		6: "ARRAY",
	}
	ExtensionType_value = map[string]int32{
		"EMPTY":     0,
		"STRING":    1,
		"INTEGER":   2,
		"BOOLEAN":   3,
		"BYTES":     4,
		"EXTENSION": 5,
		"ARRAY":     6,
	}
)

func (x ExtensionType) Enum() *ExtensionType {
	p := new(ExtensionType)
	*p = x
	return p
}

func (x ExtensionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExtensionType) Descriptor() protoreflect.EnumDescriptor {
	return file_caproto_server_proto_enumTypes[0].Descriptor()
}

func (ExtensionType) Type() protoreflect.EnumType {
	return &file_caproto_server_proto_enumTypes[0]
}

func (x ExtensionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExtensionType.Descriptor instead.
func (ExtensionType) EnumDescriptor() ([]byte, []int) {
	return file_caproto_server_proto_rawDescGZIP(), []int{0}
}

type Errors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Errors map[string]string `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Errors) Reset() {
	*x = Errors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caproto_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Errors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Errors) ProtoMessage() {}

func (x *Errors) ProtoReflect() protoreflect.Message {
	mi := &file_caproto_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Errors.ProtoReflect.Descriptor instead.
func (*Errors) Descriptor() ([]byte, []int) {
	return file_caproto_server_proto_rawDescGZIP(), []int{0}
}

func (x *Errors) GetErrors() map[string]string {
	if x != nil {
		return x.Errors
	}
	return nil
}

type ConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientVersion *Version `protobuf:"bytes,1,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	// The version of the client understood protocol
	ClientProtocolVersion *Version `protobuf:"bytes,2,opt,name=client_protocol_version,json=clientProtocolVersion,proto3" json:"client_protocol_version,omitempty"`
}

func (x *ConfigRequest) Reset() {
	*x = ConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caproto_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest) ProtoMessage() {}

func (x *ConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_caproto_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest.ProtoReflect.Descriptor instead.
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return file_caproto_server_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigRequest) GetClientVersion() *Version {
	if x != nil {
		return x.ClientVersion
	}
	return nil
}

func (x *ConfigRequest) GetClientProtocolVersion() *Version {
	if x != nil {
		return x.ClientProtocolVersion
	}
	return nil
}

type ConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The version of the server understood protocol
	ServerProtocolVersion *Version `protobuf:"bytes,1,opt,name=server_protocol_version,json=serverProtocolVersion,proto3" json:"server_protocol_version,omitempty"`
	// The time at which the server replied
	ReplyTime uint64 `protobuf:"varint,2,opt,name=reply_time,json=replyTime,proto3" json:"reply_time,omitempty"`
	ServerId  string `protobuf:"bytes,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// The server's certification policy
	// and the allowed extensions
	Policy *Extensions `protobuf:"bytes,4,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *ConfigReply) Reset() {
	*x = ConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caproto_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigReply) ProtoMessage() {}

func (x *ConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_caproto_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigReply.ProtoReflect.Descriptor instead.
func (*ConfigReply) Descriptor() ([]byte, []int) {
	return file_caproto_server_proto_rawDescGZIP(), []int{2}
}

func (x *ConfigReply) GetServerProtocolVersion() *Version {
	if x != nil {
		return x.ServerProtocolVersion
	}
	return nil
}

func (x *ConfigReply) GetReplyTime() uint64 {
	if x != nil {
		return x.ReplyTime
	}
	return 0
}

func (x *ConfigReply) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *ConfigReply) GetPolicy() *Extensions {
	if x != nil {
		return x.Policy
	}
	return nil
}

type EmptyMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMsg) Reset() {
	*x = EmptyMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caproto_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMsg) ProtoMessage() {}

func (x *EmptyMsg) ProtoReflect() protoreflect.Message {
	mi := &file_caproto_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMsg.ProtoReflect.Descriptor instead.
func (*EmptyMsg) Descriptor() ([]byte, []int) {
	return file_caproto_server_proto_rawDescGZIP(), []int{3}
}

type AuthReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The time at which the authentication was performed
	AuthTime int64 `protobuf:"varint,1,opt,name=auth_time,json=authTime,proto3" json:"auth_time,omitempty"`
	// The time at which the authentication expires, or 0 if failure
	AuthUntil int64 `protobuf:"varint,2,opt,name=auth_until,json=authUntil,proto3" json:"auth_until,omitempty"`
	// The authentication token to be used for further requests
	AuthToken string `protobuf:"bytes,3,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	// Whether the authentication was successful
	Success bool `protobuf:"varint,4,opt,name=success,proto3" json:"success,omitempty"`
	// Any errors that occurred during authentication
	Errors *Errors `protobuf:"bytes,5,opt,name=errors,proto3" json:"errors,omitempty"`
}

func (x *AuthReply) Reset() {
	*x = AuthReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caproto_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthReply) ProtoMessage() {}

func (x *AuthReply) ProtoReflect() protoreflect.Message {
	mi := &file_caproto_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthReply.ProtoReflect.Descriptor instead.
func (*AuthReply) Descriptor() ([]byte, []int) {
	return file_caproto_server_proto_rawDescGZIP(), []int{4}
}

func (x *AuthReply) GetAuthTime() int64 {
	if x != nil {
		return x.AuthTime
	}
	return 0
}

func (x *AuthReply) GetAuthUntil() int64 {
	if x != nil {
		return x.AuthUntil
	}
	return 0
}

func (x *AuthReply) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *AuthReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AuthReply) GetErrors() *Errors {
	if x != nil {
		return x.Errors
	}
	return nil
}

type CertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PublicKey                     []byte  `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	RequestedValidity             uint64  `protobuf:"fixed64,2,opt,name=requested_validity,json=requestedValidity,proto3" json:"requested_validity,omitempty"`
	ExtendedValidityJustification *string `protobuf:"bytes,4,opt,name=extendedValidityJustification,proto3,oneof" json:"extendedValidityJustification,omitempty"`
}

func (x *CertRequest) Reset() {
	*x = CertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caproto_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertRequest) ProtoMessage() {}

func (x *CertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_caproto_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertRequest.ProtoReflect.Descriptor instead.
func (*CertRequest) Descriptor() ([]byte, []int) {
	return file_caproto_server_proto_rawDescGZIP(), []int{5}
}

func (x *CertRequest) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *CertRequest) GetRequestedValidity() uint64 {
	if x != nil {
		return x.RequestedValidity
	}
	return 0
}

func (x *CertRequest) GetExtendedValidityJustification() string {
	if x != nil && x.ExtendedValidityJustification != nil {
		return *x.ExtendedValidityJustification
	}
	return ""
}

type CertReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cert       string `protobuf:"bytes,1,opt,name=cert,proto3" json:"cert,omitempty"`
	ValidUntil uint64 `protobuf:"fixed64,2,opt,name=validUntil,proto3" json:"validUntil,omitempty"`
}

func (x *CertReply) Reset() {
	*x = CertReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caproto_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertReply) ProtoMessage() {}

func (x *CertReply) ProtoReflect() protoreflect.Message {
	mi := &file_caproto_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertReply.ProtoReflect.Descriptor instead.
func (*CertReply) Descriptor() ([]byte, []int) {
	return file_caproto_server_proto_rawDescGZIP(), []int{6}
}

func (x *CertReply) GetCert() string {
	if x != nil {
		return x.Cert
	}
	return ""
}

func (x *CertReply) GetValidUntil() uint64 {
	if x != nil {
		return x.ValidUntil
	}
	return 0
}

type ExtensionArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Extension `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ExtensionArray) Reset() {
	*x = ExtensionArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caproto_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionArray) ProtoMessage() {}

func (x *ExtensionArray) ProtoReflect() protoreflect.Message {
	mi := &file_caproto_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionArray.ProtoReflect.Descriptor instead.
func (*ExtensionArray) Descriptor() ([]byte, []int) {
	return file_caproto_server_proto_rawDescGZIP(), []int{7}
}

func (x *ExtensionArray) GetValues() []*Extension {
	if x != nil {
		return x.Values
	}
	return nil
}

type Extensions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Extensions map[string]*Extension `protobuf:"bytes,1,rep,name=extensions,proto3" json:"extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IsEmpty    bool                  `protobuf:"varint,2,opt,name=is_empty,json=isEmpty,proto3" json:"is_empty,omitempty"`
}

func (x *Extensions) Reset() {
	*x = Extensions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caproto_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extensions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extensions) ProtoMessage() {}

func (x *Extensions) ProtoReflect() protoreflect.Message {
	mi := &file_caproto_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extensions.ProtoReflect.Descriptor instead.
func (*Extensions) Descriptor() ([]byte, []int) {
	return file_caproto_server_proto_rawDescGZIP(), []int{8}
}

func (x *Extensions) GetExtensions() map[string]*Extension {
	if x != nil {
		return x.Extensions
	}
	return nil
}

func (x *Extensions) GetIsEmpty() bool {
	if x != nil {
		return x.IsEmpty
	}
	return false
}

type Extension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type           ExtensionType   `protobuf:"varint,1,opt,name=type,proto3,enum=ExtensionType" json:"type,omitempty"`
	StringValue    *string         `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof" json:"string_value,omitempty"`
	IntegerValue   *int64          `protobuf:"varint,3,opt,name=integer_value,json=integerValue,proto3,oneof" json:"integer_value,omitempty"`
	BooleanValue   *bool           `protobuf:"varint,4,opt,name=boolean_value,json=booleanValue,proto3,oneof" json:"boolean_value,omitempty"`
	BytesValue     []byte          `protobuf:"bytes,5,opt,name=bytes_value,json=bytesValue,proto3,oneof" json:"bytes_value,omitempty"`
	ExtensionValue *Extensions     `protobuf:"bytes,6,opt,name=extension_value,json=extensionValue,proto3,oneof" json:"extension_value,omitempty"`
	ArrayValue     *ExtensionArray `protobuf:"bytes,7,opt,name=array_value,json=arrayValue,proto3,oneof" json:"array_value,omitempty"`
}

func (x *Extension) Reset() {
	*x = Extension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caproto_server_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extension) ProtoMessage() {}

func (x *Extension) ProtoReflect() protoreflect.Message {
	mi := &file_caproto_server_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extension.ProtoReflect.Descriptor instead.
func (*Extension) Descriptor() ([]byte, []int) {
	return file_caproto_server_proto_rawDescGZIP(), []int{9}
}

func (x *Extension) GetType() ExtensionType {
	if x != nil {
		return x.Type
	}
	return ExtensionType_EMPTY
}

func (x *Extension) GetStringValue() string {
	if x != nil && x.StringValue != nil {
		return *x.StringValue
	}
	return ""
}

func (x *Extension) GetIntegerValue() int64 {
	if x != nil && x.IntegerValue != nil {
		return *x.IntegerValue
	}
	return 0
}

func (x *Extension) GetBooleanValue() bool {
	if x != nil && x.BooleanValue != nil {
		return *x.BooleanValue
	}
	return false
}

func (x *Extension) GetBytesValue() []byte {
	if x != nil {
		return x.BytesValue
	}
	return nil
}

func (x *Extension) GetExtensionValue() *Extensions {
	if x != nil {
		return x.ExtensionValue
	}
	return nil
}

func (x *Extension) GetArrayValue() *ExtensionArray {
	if x != nil {
		return x.ArrayValue
	}
	return nil
}

type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major uint32 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor uint32 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch uint32 `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_caproto_server_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_caproto_server_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_caproto_server_proto_rawDescGZIP(), []int{10}
}

func (x *Version) GetMajor() uint32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *Version) GetMinor() uint32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *Version) GetPatch() uint32 {
	if x != nil {
		return x.Patch
	}
	return 0
}

var File_caproto_server_proto protoreflect.FileDescriptor

var file_caproto_server_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x12, 0x2b, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x39, 0x0a,
	0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x82, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x0e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x17, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xb0, 0x01,
	0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x40, 0x0a,
	0x17, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x22, 0xa1, 0x01, 0x0a,
	0x09, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x75, 0x74,
	0x68, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x1f, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x22, 0xc8, 0x01, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x2d, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x52, 0x11, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x49,
	0x0a, 0x1d, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x1d, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x4a, 0x75, 0x73, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x20, 0x0a, 0x1e, 0x5f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x4a, 0x75,
	0x73, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x09, 0x43,
	0x65, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06,
	0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x22, 0x34, 0x0a, 0x0e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x22,
	0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x73, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x73, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x49, 0x0a, 0x0f, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xac, 0x03, 0x0a, 0x09, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0e, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x28,
	0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6c,
	0x65, 0x61, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x02, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x03, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x04,
	0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x0b, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48, 0x05, 0x52, 0x0a, 0x61, 0x72, 0x72,
	0x61, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x12, 0x0a, 0x10, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x4b, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d,
	0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x2a, 0x65, 0x0a, 0x0d, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x50, 0x54, 0x59, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x54, 0x45,
	0x47, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e,
	0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x04, 0x12, 0x0d, 0x0a,
	0x09, 0x45, 0x58, 0x54, 0x45, 0x4e, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05,
	0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x06, 0x32, 0x8b, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x73,
	0x6d, 0x61, 0x43, 0x61, 0x12, 0x2b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x0e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x27, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x09, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x73, 0x67, 0x1a, 0x0a, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x65, 0x72, 0x74, 0x12, 0x0c, 0x2e, 0x43, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x70, 0x72, 0x69, 0x73, 0x6d, 0x2f, 0x70,
	0x72, 0x69, 0x73, 0x6d, 0x61, 0x5f, 0x63, 0x61, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x61, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_caproto_server_proto_rawDescOnce sync.Once
	file_caproto_server_proto_rawDescData = file_caproto_server_proto_rawDesc
)

func file_caproto_server_proto_rawDescGZIP() []byte {
	file_caproto_server_proto_rawDescOnce.Do(func() {
		file_caproto_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_caproto_server_proto_rawDescData)
	})
	return file_caproto_server_proto_rawDescData
}

var file_caproto_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_caproto_server_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_caproto_server_proto_goTypes = []interface{}{
	(ExtensionType)(0),     // 0: ExtensionType
	(*Errors)(nil),         // 1: Errors
	(*ConfigRequest)(nil),  // 2: ConfigRequest
	(*ConfigReply)(nil),    // 3: ConfigReply
	(*EmptyMsg)(nil),       // 4: EmptyMsg
	(*AuthReply)(nil),      // 5: AuthReply
	(*CertRequest)(nil),    // 6: CertRequest
	(*CertReply)(nil),      // 7: CertReply
	(*ExtensionArray)(nil), // 8: ExtensionArray
	(*Extensions)(nil),     // 9: Extensions
	(*Extension)(nil),      // 10: Extension
	(*Version)(nil),        // 11: Version
	nil,                    // 12: Errors.ErrorsEntry
	nil,                    // 13: Extensions.ExtensionsEntry
}
var file_caproto_server_proto_depIdxs = []int32{
	12, // 0: Errors.errors:type_name -> Errors.ErrorsEntry
	11, // 1: ConfigRequest.client_version:type_name -> Version
	11, // 2: ConfigRequest.client_protocol_version:type_name -> Version
	11, // 3: ConfigReply.server_protocol_version:type_name -> Version
	9,  // 4: ConfigReply.policy:type_name -> Extensions
	1,  // 5: AuthReply.errors:type_name -> Errors
	10, // 6: ExtensionArray.values:type_name -> Extension
	13, // 7: Extensions.extensions:type_name -> Extensions.ExtensionsEntry
	0,  // 8: Extension.type:type_name -> ExtensionType
	9,  // 9: Extension.extension_value:type_name -> Extensions
	8,  // 10: Extension.array_value:type_name -> ExtensionArray
	10, // 11: Extensions.ExtensionsEntry.value:type_name -> Extension
	2,  // 12: PrismaCa.GetConfig:input_type -> ConfigRequest
	4,  // 13: PrismaCa.Authenticate:input_type -> EmptyMsg
	6,  // 14: PrismaCa.RequestCert:input_type -> CertRequest
	3,  // 15: PrismaCa.GetConfig:output_type -> ConfigReply
	5,  // 16: PrismaCa.Authenticate:output_type -> AuthReply
	7,  // 17: PrismaCa.RequestCert:output_type -> CertReply
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_caproto_server_proto_init() }
func file_caproto_server_proto_init() {
	if File_caproto_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_caproto_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Errors); i {
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
		file_caproto_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest); i {
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
		file_caproto_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigReply); i {
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
		file_caproto_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMsg); i {
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
		file_caproto_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthReply); i {
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
		file_caproto_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertRequest); i {
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
		file_caproto_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertReply); i {
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
		file_caproto_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionArray); i {
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
		file_caproto_server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extensions); i {
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
		file_caproto_server_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extension); i {
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
		file_caproto_server_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
	file_caproto_server_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_caproto_server_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_caproto_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_caproto_server_proto_goTypes,
		DependencyIndexes: file_caproto_server_proto_depIdxs,
		EnumInfos:         file_caproto_server_proto_enumTypes,
		MessageInfos:      file_caproto_server_proto_msgTypes,
	}.Build()
	File_caproto_server_proto = out.File
	file_caproto_server_proto_rawDesc = nil
	file_caproto_server_proto_goTypes = nil
	file_caproto_server_proto_depIdxs = nil
}
