// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: certificate/certificate.proto

package certificate

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

type CertificateMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Serial         string  `protobuf:"bytes,2,opt,name=serial,proto3" json:"serial,omitempty"`
	Description    string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt      uint64  `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Ttl            *uint64 `protobuf:"varint,5,opt,name=ttl,proto3,oneof" json:"ttl,omitempty"`
	CertUuid       []byte  `protobuf:"bytes,6,opt,name=cert_uuid,json=certUuid,proto3" json:"cert_uuid,omitempty"`
	IsRoot         bool    `protobuf:"varint,7,opt,name=is_root,json=isRoot,proto3" json:"is_root,omitempty"`
	IsIntermediate bool    `protobuf:"varint,8,opt,name=is_intermediate,json=isIntermediate,proto3" json:"is_intermediate,omitempty"`
	Username       string  `protobuf:"bytes,9,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *CertificateMetadata) Reset() {
	*x = CertificateMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_certificate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateMetadata) ProtoMessage() {}

func (x *CertificateMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_certificate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateMetadata.ProtoReflect.Descriptor instead.
func (*CertificateMetadata) Descriptor() ([]byte, []int) {
	return file_certificate_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *CertificateMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CertificateMetadata) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *CertificateMetadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CertificateMetadata) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *CertificateMetadata) GetTtl() uint64 {
	if x != nil && x.Ttl != nil {
		return *x.Ttl
	}
	return 0
}

func (x *CertificateMetadata) GetCertUuid() []byte {
	if x != nil {
		return x.CertUuid
	}
	return nil
}

func (x *CertificateMetadata) GetIsRoot() bool {
	if x != nil {
		return x.IsRoot
	}
	return false
}

func (x *CertificateMetadata) GetIsIntermediate() bool {
	if x != nil {
		return x.IsIntermediate
	}
	return false
}

func (x *CertificateMetadata) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ServerPublicKeyMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentCertId []byte `protobuf:"bytes,1,opt,name=parent_cert_id,json=parentCertId,proto3" json:"parent_cert_id,omitempty"`
}

func (x *ServerPublicKeyMetadata) Reset() {
	*x = ServerPublicKeyMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_certificate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerPublicKeyMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerPublicKeyMetadata) ProtoMessage() {}

func (x *ServerPublicKeyMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_certificate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerPublicKeyMetadata.ProtoReflect.Descriptor instead.
func (*ServerPublicKeyMetadata) Descriptor() ([]byte, []int) {
	return file_certificate_certificate_proto_rawDescGZIP(), []int{1}
}

func (x *ServerPublicKeyMetadata) GetParentCertId() []byte {
	if x != nil {
		return x.ParentCertId
	}
	return nil
}

type AuthorizedUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName      string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserPublicKey []byte `protobuf:"bytes,2,opt,name=user_public_key,json=userPublicKey,proto3" json:"user_public_key,omitempty"`
	CreatedAt     int64  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CertId        []byte `protobuf:"bytes,4,opt,name=cert_id,json=certId,proto3" json:"cert_id,omitempty"`
}

func (x *AuthorizedUser) Reset() {
	*x = AuthorizedUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_certificate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizedUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizedUser) ProtoMessage() {}

func (x *AuthorizedUser) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_certificate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizedUser.ProtoReflect.Descriptor instead.
func (*AuthorizedUser) Descriptor() ([]byte, []int) {
	return file_certificate_certificate_proto_rawDescGZIP(), []int{2}
}

func (x *AuthorizedUser) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *AuthorizedUser) GetUserPublicKey() []byte {
	if x != nil {
		return x.UserPublicKey
	}
	return nil
}

func (x *AuthorizedUser) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *AuthorizedUser) GetCertId() []byte {
	if x != nil {
		return x.CertId
	}
	return nil
}

type AuthorizedUsers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdatedAt     uint64                     `protobuf:"varint,1,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Users         map[string]*AuthorizedUser `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IntegrityHash []byte                     `protobuf:"bytes,3,opt,name=integrity_hash,json=integrityHash,proto3" json:"integrity_hash,omitempty"`
}

func (x *AuthorizedUsers) Reset() {
	*x = AuthorizedUsers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_certificate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizedUsers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizedUsers) ProtoMessage() {}

func (x *AuthorizedUsers) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_certificate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizedUsers.ProtoReflect.Descriptor instead.
func (*AuthorizedUsers) Descriptor() ([]byte, []int) {
	return file_certificate_certificate_proto_rawDescGZIP(), []int{3}
}

func (x *AuthorizedUsers) GetUpdatedAt() uint64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *AuthorizedUsers) GetUsers() map[string]*AuthorizedUser {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *AuthorizedUsers) GetIntegrityHash() []byte {
	if x != nil {
		return x.IntegrityHash
	}
	return nil
}

var File_certificate_certificate_proto protoreflect.FileDescriptor

var file_certificate_certificate_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9c, 0x02, 0x0a, 0x13, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x15, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x48, 0x00, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x65, 0x72, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x63, 0x65, 0x72, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x72,
	0x6f, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x52, 0x6f, 0x6f,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64,
	0x69, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x74, 0x74, 0x6c, 0x22, 0x3f,
	0x0a, 0x17, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x49, 0x64, 0x22,
	0x8d, 0x01, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x26, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x65, 0x72, 0x74, 0x49, 0x64, 0x22,
	0xd5, 0x01, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x69,
	0x74, 0x79, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x69, 0x74, 0x79, 0x48, 0x61, 0x73, 0x68, 0x1a, 0x49, 0x0a, 0x0a,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x70, 0x72, 0x69, 0x73, 0x6d,
	0x2f, 0x70, 0x72, 0x69, 0x73, 0x6d, 0x61, 0x5f, 0x63, 0x61, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_certificate_certificate_proto_rawDescOnce sync.Once
	file_certificate_certificate_proto_rawDescData = file_certificate_certificate_proto_rawDesc
)

func file_certificate_certificate_proto_rawDescGZIP() []byte {
	file_certificate_certificate_proto_rawDescOnce.Do(func() {
		file_certificate_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_certificate_certificate_proto_rawDescData)
	})
	return file_certificate_certificate_proto_rawDescData
}

var file_certificate_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_certificate_certificate_proto_goTypes = []interface{}{
	(*CertificateMetadata)(nil),     // 0: CertificateMetadata
	(*ServerPublicKeyMetadata)(nil), // 1: ServerPublicKeyMetadata
	(*AuthorizedUser)(nil),          // 2: AuthorizedUser
	(*AuthorizedUsers)(nil),         // 3: AuthorizedUsers
	nil,                             // 4: AuthorizedUsers.UsersEntry
}
var file_certificate_certificate_proto_depIdxs = []int32{
	4, // 0: AuthorizedUsers.users:type_name -> AuthorizedUsers.UsersEntry
	2, // 1: AuthorizedUsers.UsersEntry.value:type_name -> AuthorizedUser
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_certificate_certificate_proto_init() }
func file_certificate_certificate_proto_init() {
	if File_certificate_certificate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_certificate_certificate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateMetadata); i {
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
		file_certificate_certificate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerPublicKeyMetadata); i {
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
		file_certificate_certificate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizedUser); i {
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
		file_certificate_certificate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizedUsers); i {
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
	file_certificate_certificate_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_certificate_certificate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_certificate_certificate_proto_goTypes,
		DependencyIndexes: file_certificate_certificate_proto_depIdxs,
		MessageInfos:      file_certificate_certificate_proto_msgTypes,
	}.Build()
	File_certificate_certificate_proto = out.File
	file_certificate_certificate_proto_rawDesc = nil
	file_certificate_certificate_proto_goTypes = nil
	file_certificate_certificate_proto_depIdxs = nil
}
