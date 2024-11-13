// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.23.3
// source: certificate.proto

package proto

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

type CreateCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IeltsScore     float32 `protobuf:"fixed32,2,opt,name=ielts_score,json=ieltsScore,proto3" json:"ielts_score,omitempty"`
	CefrLevel      string  `protobuf:"bytes,3,opt,name=cefr_level,json=cefrLevel,proto3" json:"cefr_level,omitempty"`
	Description    string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CertificateUrl string  `protobuf:"bytes,5,opt,name=certificate_url,json=certificateUrl,proto3" json:"certificate_url,omitempty"`
}

func (x *CreateCertificate) Reset() {
	*x = CreateCertificate{}
	mi := &file_certificate_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCertificate) ProtoMessage() {}

func (x *CreateCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCertificate.ProtoReflect.Descriptor instead.
func (*CreateCertificate) Descriptor() ([]byte, []int) {
	return file_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCertificate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCertificate) GetIeltsScore() float32 {
	if x != nil {
		return x.IeltsScore
	}
	return 0
}

func (x *CreateCertificate) GetCefrLevel() string {
	if x != nil {
		return x.CefrLevel
	}
	return ""
}

func (x *CreateCertificate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateCertificate) GetCertificateUrl() string {
	if x != nil {
		return x.CertificateUrl
	}
	return ""
}

type UpdateCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IeltsScore     float32 `protobuf:"fixed32,3,opt,name=ielts_score,json=ieltsScore,proto3" json:"ielts_score,omitempty"`
	CefrLevel      string  `protobuf:"bytes,4,opt,name=cefr_level,json=cefrLevel,proto3" json:"cefr_level,omitempty"`
	Description    string  `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	CertificateUrl string  `protobuf:"bytes,6,opt,name=certificate_url,json=certificateUrl,proto3" json:"certificate_url,omitempty"`
}

func (x *UpdateCertificate) Reset() {
	*x = UpdateCertificate{}
	mi := &file_certificate_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCertificate) ProtoMessage() {}

func (x *UpdateCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCertificate.ProtoReflect.Descriptor instead.
func (*UpdateCertificate) Descriptor() ([]byte, []int) {
	return file_certificate_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateCertificate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCertificate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCertificate) GetIeltsScore() float32 {
	if x != nil {
		return x.IeltsScore
	}
	return 0
}

func (x *UpdateCertificate) GetCefrLevel() string {
	if x != nil {
		return x.CefrLevel
	}
	return ""
}

func (x *UpdateCertificate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateCertificate) GetCertificateUrl() string {
	if x != nil {
		return x.CertificateUrl
	}
	return ""
}

type CertificateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IeltsScore     float32 `protobuf:"fixed32,3,opt,name=ielts_score,json=ieltsScore,proto3" json:"ielts_score,omitempty"`
	CefrLevel      string  `protobuf:"bytes,4,opt,name=cefr_level,json=cefrLevel,proto3" json:"cefr_level,omitempty"`
	Description    string  `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	CertificateUrl string  `protobuf:"bytes,6,opt,name=certificate_url,json=certificateUrl,proto3" json:"certificate_url,omitempty"`
	CreatedAt      string  `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *CertificateRes) Reset() {
	*x = CertificateRes{}
	mi := &file_certificate_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CertificateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateRes) ProtoMessage() {}

func (x *CertificateRes) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateRes.ProtoReflect.Descriptor instead.
func (*CertificateRes) Descriptor() ([]byte, []int) {
	return file_certificate_proto_rawDescGZIP(), []int{2}
}

func (x *CertificateRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CertificateRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CertificateRes) GetIeltsScore() float32 {
	if x != nil {
		return x.IeltsScore
	}
	return 0
}

func (x *CertificateRes) GetCefrLevel() string {
	if x != nil {
		return x.CefrLevel
	}
	return ""
}

func (x *CertificateRes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CertificateRes) GetCertificateUrl() string {
	if x != nil {
		return x.CertificateUrl
	}
	return ""
}

func (x *CertificateRes) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetListCertificateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IeltsScore float32 `protobuf:"fixed32,2,opt,name=ielts_score,json=ieltsScore,proto3" json:"ielts_score,omitempty"`
	CefrLevel  string  `protobuf:"bytes,3,opt,name=cefr_level,json=cefrLevel,proto3" json:"cefr_level,omitempty"`
	Filter     *Filter `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetListCertificateReq) Reset() {
	*x = GetListCertificateReq{}
	mi := &file_certificate_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListCertificateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListCertificateReq) ProtoMessage() {}

func (x *GetListCertificateReq) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListCertificateReq.ProtoReflect.Descriptor instead.
func (*GetListCertificateReq) Descriptor() ([]byte, []int) {
	return file_certificate_proto_rawDescGZIP(), []int{3}
}

func (x *GetListCertificateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetListCertificateReq) GetIeltsScore() float32 {
	if x != nil {
		return x.IeltsScore
	}
	return 0
}

func (x *GetListCertificateReq) GetCefrLevel() string {
	if x != nil {
		return x.CefrLevel
	}
	return ""
}

func (x *GetListCertificateReq) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetListCertificateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certificates []*CertificateRes `protobuf:"bytes,1,rep,name=certificates,proto3" json:"certificates,omitempty"`
	TotalCount   int32             `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *GetListCertificateRes) Reset() {
	*x = GetListCertificateRes{}
	mi := &file_certificate_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListCertificateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListCertificateRes) ProtoMessage() {}

func (x *GetListCertificateRes) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListCertificateRes.ProtoReflect.Descriptor instead.
func (*GetListCertificateRes) Descriptor() ([]byte, []int) {
	return file_certificate_proto_rawDescGZIP(), []int{4}
}

func (x *GetListCertificateRes) GetCertificates() []*CertificateRes {
	if x != nil {
		return x.Certificates
	}
	return nil
}

func (x *GetListCertificateRes) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_certificate_proto protoreflect.FileDescriptor

var file_certificate_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x5f, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x65, 0x66, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x65, 0x66, 0x72, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x22, 0xc2, 0x01,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x65, 0x6c, 0x74, 0x73,
	0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x69, 0x65,
	0x6c, 0x74, 0x73, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x65, 0x66, 0x72,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x65,
	0x66, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x55,
	0x72, 0x6c, 0x22, 0xde, 0x01, 0x0a, 0x0e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x65, 0x6c,
	0x74, 0x73, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a,
	0x69, 0x65, 0x6c, 0x74, 0x73, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x65,
	0x66, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x65, 0x66, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x92, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x69, 0x65, 0x6c, 0x74, 0x73, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x65, 0x66, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x65, 0x66, 0x72, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x25, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x73, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x12, 0x39, 0x0a, 0x0c, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x52, 0x0c,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x90, 0x02,
	0x0a, 0x12, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x79,
	0x49, 0x64, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x45, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x42, 0x19, 0x5a, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_certificate_proto_rawDescOnce sync.Once
	file_certificate_proto_rawDescData = file_certificate_proto_rawDesc
)

func file_certificate_proto_rawDescGZIP() []byte {
	file_certificate_proto_rawDescOnce.Do(func() {
		file_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_certificate_proto_rawDescData)
	})
	return file_certificate_proto_rawDescData
}

var file_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_certificate_proto_goTypes = []any{
	(*CreateCertificate)(nil),     // 0: proto.CreateCertificate
	(*UpdateCertificate)(nil),     // 1: proto.UpdateCertificate
	(*CertificateRes)(nil),        // 2: proto.CertificateRes
	(*GetListCertificateReq)(nil), // 3: proto.GetListCertificateReq
	(*GetListCertificateRes)(nil), // 4: proto.GetListCertificateRes
	(*Filter)(nil),                // 5: proto.Filter
	(*ById)(nil),                  // 6: proto.ById
	(*Void)(nil),                  // 7: proto.Void
}
var file_certificate_proto_depIdxs = []int32{
	5, // 0: proto.GetListCertificateReq.filter:type_name -> proto.Filter
	2, // 1: proto.GetListCertificateRes.certificates:type_name -> proto.CertificateRes
	0, // 2: proto.CertificateService.Create:input_type -> proto.CreateCertificate
	1, // 3: proto.CertificateService.Update:input_type -> proto.UpdateCertificate
	6, // 4: proto.CertificateService.Delete:input_type -> proto.ById
	6, // 5: proto.CertificateService.GetById:input_type -> proto.ById
	3, // 6: proto.CertificateService.GetList:input_type -> proto.GetListCertificateReq
	7, // 7: proto.CertificateService.Create:output_type -> proto.Void
	7, // 8: proto.CertificateService.Update:output_type -> proto.Void
	7, // 9: proto.CertificateService.Delete:output_type -> proto.Void
	2, // 10: proto.CertificateService.GetById:output_type -> proto.CertificateRes
	4, // 11: proto.CertificateService.GetList:output_type -> proto.GetListCertificateRes
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_certificate_proto_init() }
func file_certificate_proto_init() {
	if File_certificate_proto != nil {
		return
	}
	file_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_certificate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_certificate_proto_goTypes,
		DependencyIndexes: file_certificate_proto_depIdxs,
		MessageInfos:      file_certificate_proto_msgTypes,
	}.Build()
	File_certificate_proto = out.File
	file_certificate_proto_rawDesc = nil
	file_certificate_proto_goTypes = nil
	file_certificate_proto_depIdxs = nil
}
