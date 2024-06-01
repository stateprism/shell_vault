// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: adminproto/admin.proto

package adminproto

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adminproto_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_adminproto_admin_proto_msgTypes[0]
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
	return file_adminproto_admin_proto_rawDescGZIP(), []int{0}
}

type StopServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason string `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
	Token  []byte `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *StopServerRequest) Reset() {
	*x = StopServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adminproto_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopServerRequest) ProtoMessage() {}

func (x *StopServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adminproto_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopServerRequest.ProtoReflect.Descriptor instead.
func (*StopServerRequest) Descriptor() ([]byte, []int) {
	return file_adminproto_admin_proto_rawDescGZIP(), []int{1}
}

func (x *StopServerRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *StopServerRequest) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

type AdminAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeResponse []byte `protobuf:"bytes,1,opt,name=challenge_response,json=challengeResponse,proto3" json:"challenge_response,omitempty"`
	ChallengeId       []byte `protobuf:"bytes,2,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	ChallengeNonce    []byte `protobuf:"bytes,3,opt,name=challenge_nonce,json=challengeNonce,proto3" json:"challenge_nonce,omitempty"`
}

func (x *AdminAuthRequest) Reset() {
	*x = AdminAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adminproto_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminAuthRequest) ProtoMessage() {}

func (x *AdminAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adminproto_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminAuthRequest.ProtoReflect.Descriptor instead.
func (*AdminAuthRequest) Descriptor() ([]byte, []int) {
	return file_adminproto_admin_proto_rawDescGZIP(), []int{2}
}

func (x *AdminAuthRequest) GetChallengeResponse() []byte {
	if x != nil {
		return x.ChallengeResponse
	}
	return nil
}

func (x *AdminAuthRequest) GetChallengeId() []byte {
	if x != nil {
		return x.ChallengeId
	}
	return nil
}

func (x *AdminAuthRequest) GetChallengeNonce() []byte {
	if x != nil {
		return x.ChallengeNonce
	}
	return nil
}

type AdminAuthInitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Challenge         []byte `protobuf:"bytes,1,opt,name=challenge,proto3" json:"challenge,omitempty"`
	ChallengeId       []byte `protobuf:"bytes,2,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	ChallengeValidity uint64 `protobuf:"varint,3,opt,name=challenge_validity,json=challengeValidity,proto3" json:"challenge_validity,omitempty"`
	ChallengeNonce    uint64 `protobuf:"varint,4,opt,name=challenge_nonce,json=challengeNonce,proto3" json:"challenge_nonce,omitempty"`
}

func (x *AdminAuthInitResponse) Reset() {
	*x = AdminAuthInitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adminproto_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminAuthInitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminAuthInitResponse) ProtoMessage() {}

func (x *AdminAuthInitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_adminproto_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminAuthInitResponse.ProtoReflect.Descriptor instead.
func (*AdminAuthInitResponse) Descriptor() ([]byte, []int) {
	return file_adminproto_admin_proto_rawDescGZIP(), []int{3}
}

func (x *AdminAuthInitResponse) GetChallenge() []byte {
	if x != nil {
		return x.Challenge
	}
	return nil
}

func (x *AdminAuthInitResponse) GetChallengeId() []byte {
	if x != nil {
		return x.ChallengeId
	}
	return nil
}

func (x *AdminAuthInitResponse) GetChallengeValidity() uint64 {
	if x != nil {
		return x.ChallengeValidity
	}
	return 0
}

func (x *AdminAuthInitResponse) GetChallengeNonce() uint64 {
	if x != nil {
		return x.ChallengeNonce
	}
	return 0
}

type AdminAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The username of the admin as will be logged in the server
	Username *string `protobuf:"bytes,1,opt,name=username,proto3,oneof" json:"username,omitempty"`
	// The token to be used for further requests
	Token []byte `protobuf:"bytes,2,opt,name=token,proto3,oneof" json:"token,omitempty"`
	// The validity of the token in seconds
	TokenValidity *uint64 `protobuf:"varint,3,opt,name=token_validity,json=tokenValidity,proto3,oneof" json:"token_validity,omitempty"`
	Success       bool    `protobuf:"varint,4,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AdminAuthResponse) Reset() {
	*x = AdminAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adminproto_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminAuthResponse) ProtoMessage() {}

func (x *AdminAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_adminproto_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminAuthResponse.ProtoReflect.Descriptor instead.
func (*AdminAuthResponse) Descriptor() ([]byte, []int) {
	return file_adminproto_admin_proto_rawDescGZIP(), []int{4}
}

func (x *AdminAuthResponse) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *AdminAuthResponse) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *AdminAuthResponse) GetTokenValidity() uint64 {
	if x != nil && x.TokenValidity != nil {
		return *x.TokenValidity
	}
	return 0
}

func (x *AdminAuthResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ChangeRootCertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CertType     string `protobuf:"bytes,1,opt,name=cert_type,json=certType,proto3" json:"cert_type,omitempty"`
	CertComent   string `protobuf:"bytes,2,opt,name=cert_coment,json=certComent,proto3" json:"cert_coment,omitempty"`
	CertValidity uint64 `protobuf:"varint,3,opt,name=cert_validity,json=certValidity,proto3" json:"cert_validity,omitempty"`
	Token        []byte `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ChangeRootCertRequest) Reset() {
	*x = ChangeRootCertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adminproto_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeRootCertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeRootCertRequest) ProtoMessage() {}

func (x *ChangeRootCertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_adminproto_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeRootCertRequest.ProtoReflect.Descriptor instead.
func (*ChangeRootCertRequest) Descriptor() ([]byte, []int) {
	return file_adminproto_admin_proto_rawDescGZIP(), []int{5}
}

func (x *ChangeRootCertRequest) GetCertType() string {
	if x != nil {
		return x.CertType
	}
	return ""
}

func (x *ChangeRootCertRequest) GetCertComent() string {
	if x != nil {
		return x.CertComent
	}
	return ""
}

func (x *ChangeRootCertRequest) GetCertValidity() uint64 {
	if x != nil {
		return x.CertValidity
	}
	return 0
}

func (x *ChangeRootCertRequest) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

// Client should check the response to see if the certificate was changed
// success should be true and the returned values should be the same as the requested ones
type ChangeRootCertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success      bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	CertType     *string `protobuf:"bytes,2,opt,name=cert_type,json=certType,proto3,oneof" json:"cert_type,omitempty"`
	CertComent   *string `protobuf:"bytes,3,opt,name=cert_coment,json=certComent,proto3,oneof" json:"cert_coment,omitempty"`
	CertValidity *uint64 `protobuf:"varint,4,opt,name=cert_validity,json=certValidity,proto3,oneof" json:"cert_validity,omitempty"`
}

func (x *ChangeRootCertResponse) Reset() {
	*x = ChangeRootCertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_adminproto_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeRootCertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeRootCertResponse) ProtoMessage() {}

func (x *ChangeRootCertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_adminproto_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeRootCertResponse.ProtoReflect.Descriptor instead.
func (*ChangeRootCertResponse) Descriptor() ([]byte, []int) {
	return file_adminproto_admin_proto_rawDescGZIP(), []int{6}
}

func (x *ChangeRootCertResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ChangeRootCertResponse) GetCertType() string {
	if x != nil && x.CertType != nil {
		return *x.CertType
	}
	return ""
}

func (x *ChangeRootCertResponse) GetCertComent() string {
	if x != nil && x.CertComent != nil {
		return *x.CertComent
	}
	return ""
}

func (x *ChangeRootCertResponse) GetCertValidity() uint64 {
	if x != nil && x.CertValidity != nil {
		return *x.CertValidity
	}
	return 0
}

var File_adminproto_admin_proto protoreflect.FileDescriptor

var file_adminproto_admin_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x41, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8d, 0x01, 0x0a, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4e,
	0x6f, 0x6e, 0x63, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x15, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x75,
	0x74, 0x68, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x2d, 0x0a, 0x12, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x63, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0xbf, 0x01, 0x0a, 0x11, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x01, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x48, 0x02, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x22, 0x90, 0x01, 0x0a, 0x15, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x63, 0x65, 0x72, 0x74, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xd4, 0x01, 0x0a,
	0x16, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x20, 0x0a, 0x09, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x63, 0x65, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x63, 0x65, 0x72, 0x74,
	0x43, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x63, 0x65, 0x72,
	0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x48, 0x02, 0x52, 0x0c, 0x63, 0x65, 0x72, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x32, 0xe7, 0x01, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x0d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x75, 0x74,
	0x68, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x12, 0x11, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x12,
	0x2e, 0x53, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x41, 0x0a, 0x0e, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x6f, 0x6f,
	0x74, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x70, 0x72, 0x69, 0x73, 0x6d, 0x2f, 0x70, 0x72, 0x69, 0x73, 0x6d, 0x61, 0x5f, 0x63, 0x61,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_adminproto_admin_proto_rawDescOnce sync.Once
	file_adminproto_admin_proto_rawDescData = file_adminproto_admin_proto_rawDesc
)

func file_adminproto_admin_proto_rawDescGZIP() []byte {
	file_adminproto_admin_proto_rawDescOnce.Do(func() {
		file_adminproto_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_adminproto_admin_proto_rawDescData)
	})
	return file_adminproto_admin_proto_rawDescData
}

var file_adminproto_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_adminproto_admin_proto_goTypes = []interface{}{
	(*Empty)(nil),                  // 0: Empty
	(*StopServerRequest)(nil),      // 1: StopServerRequest
	(*AdminAuthRequest)(nil),       // 2: AdminAuthRequest
	(*AdminAuthInitResponse)(nil),  // 3: AdminAuthInitResponse
	(*AdminAuthResponse)(nil),      // 4: AdminAuthResponse
	(*ChangeRootCertRequest)(nil),  // 5: ChangeRootCertRequest
	(*ChangeRootCertResponse)(nil), // 6: ChangeRootCertResponse
}
var file_adminproto_admin_proto_depIdxs = []int32{
	0, // 0: AdminService.AdminAuthInit:input_type -> Empty
	2, // 1: AdminService.AdminAuthRespond:input_type -> AdminAuthRequest
	1, // 2: AdminService.StopServer:input_type -> StopServerRequest
	5, // 3: AdminService.ChangeRootCert:input_type -> ChangeRootCertRequest
	3, // 4: AdminService.AdminAuthInit:output_type -> AdminAuthInitResponse
	4, // 5: AdminService.AdminAuthRespond:output_type -> AdminAuthResponse
	0, // 6: AdminService.StopServer:output_type -> Empty
	6, // 7: AdminService.ChangeRootCert:output_type -> ChangeRootCertResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_adminproto_admin_proto_init() }
func file_adminproto_admin_proto_init() {
	if File_adminproto_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_adminproto_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_adminproto_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopServerRequest); i {
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
		file_adminproto_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminAuthRequest); i {
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
		file_adminproto_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminAuthInitResponse); i {
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
		file_adminproto_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminAuthResponse); i {
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
		file_adminproto_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeRootCertRequest); i {
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
		file_adminproto_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeRootCertResponse); i {
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
	file_adminproto_admin_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_adminproto_admin_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_adminproto_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_adminproto_admin_proto_goTypes,
		DependencyIndexes: file_adminproto_admin_proto_depIdxs,
		MessageInfos:      file_adminproto_admin_proto_msgTypes,
	}.Build()
	File_adminproto_admin_proto = out.File
	file_adminproto_admin_proto_rawDesc = nil
	file_adminproto_admin_proto_goTypes = nil
	file_adminproto_admin_proto_depIdxs = nil
}
