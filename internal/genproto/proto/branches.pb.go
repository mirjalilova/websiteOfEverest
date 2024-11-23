// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: branches.proto

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

type CreateBranches struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	GoogleUrl   string `protobuf:"bytes,3,opt,name=google_url,json=googleUrl,proto3" json:"google_url,omitempty"`
	YandexUrl   string `protobuf:"bytes,4,opt,name=yandex_url,json=yandexUrl,proto3" json:"yandex_url,omitempty"`
	Contact     string `protobuf:"bytes,5,opt,name=contact,proto3" json:"contact,omitempty"`
}

func (x *CreateBranches) Reset() {
	*x = CreateBranches{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBranches) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBranches) ProtoMessage() {}

func (x *CreateBranches) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBranches.ProtoReflect.Descriptor instead.
func (*CreateBranches) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBranches) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBranches) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateBranches) GetGoogleUrl() string {
	if x != nil {
		return x.GoogleUrl
	}
	return ""
}

func (x *CreateBranches) GetYandexUrl() string {
	if x != nil {
		return x.YandexUrl
	}
	return ""
}

func (x *CreateBranches) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

type UpdateBranches struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	GoogleUrl   string `protobuf:"bytes,4,opt,name=google_url,json=googleUrl,proto3" json:"google_url,omitempty"`
	YandexUrl   string `protobuf:"bytes,5,opt,name=yandex_url,json=yandexUrl,proto3" json:"yandex_url,omitempty"`
	Contact     string `protobuf:"bytes,6,opt,name=contact,proto3" json:"contact,omitempty"`
}

func (x *UpdateBranches) Reset() {
	*x = UpdateBranches{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBranches) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBranches) ProtoMessage() {}

func (x *UpdateBranches) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBranches.ProtoReflect.Descriptor instead.
func (*UpdateBranches) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateBranches) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBranches) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBranches) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateBranches) GetGoogleUrl() string {
	if x != nil {
		return x.GoogleUrl
	}
	return ""
}

func (x *UpdateBranches) GetYandexUrl() string {
	if x != nil {
		return x.YandexUrl
	}
	return ""
}

func (x *UpdateBranches) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

type BranchesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	GoogleUrl   string `protobuf:"bytes,4,opt,name=google_url,json=googleUrl,proto3" json:"google_url,omitempty"`
	YandexUrl   string `protobuf:"bytes,5,opt,name=yandex_url,json=yandexUrl,proto3" json:"yandex_url,omitempty"`
	Contact     string `protobuf:"bytes,6,opt,name=contact,proto3" json:"contact,omitempty"`
	CreatedAt   string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *BranchesRes) Reset() {
	*x = BranchesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchesRes) ProtoMessage() {}

func (x *BranchesRes) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchesRes.ProtoReflect.Descriptor instead.
func (*BranchesRes) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{2}
}

func (x *BranchesRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BranchesRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BranchesRes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BranchesRes) GetGoogleUrl() string {
	if x != nil {
		return x.GoogleUrl
	}
	return ""
}

func (x *BranchesRes) GetYandexUrl() string {
	if x != nil {
		return x.YandexUrl
	}
	return ""
}

func (x *BranchesRes) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *BranchesRes) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetListBranchesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Filter *Filter `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetListBranchesReq) Reset() {
	*x = GetListBranchesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListBranchesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListBranchesReq) ProtoMessage() {}

func (x *GetListBranchesReq) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListBranchesReq.ProtoReflect.Descriptor instead.
func (*GetListBranchesReq) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{3}
}

func (x *GetListBranchesReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetListBranchesReq) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetListBranchesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Branches   []*BranchesRes `protobuf:"bytes,1,rep,name=branches,proto3" json:"branches,omitempty"`
	TotalCount int32          `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *GetListBranchesRes) Reset() {
	*x = GetListBranchesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branches_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListBranchesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListBranchesRes) ProtoMessage() {}

func (x *GetListBranchesRes) ProtoReflect() protoreflect.Message {
	mi := &file_branches_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListBranchesRes.ProtoReflect.Descriptor instead.
func (*GetListBranchesRes) Descriptor() ([]byte, []int) {
	return file_branches_proto_rawDescGZIP(), []int{4}
}

func (x *GetListBranchesRes) GetBranches() []*BranchesRes {
	if x != nil {
		return x.Branches
	}
	return nil
}

func (x *GetListBranchesRes) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_branches_proto protoreflect.FileDescriptor

var file_branches_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a,
	0x0a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0xae, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1d,
	0x0a, 0x0a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0xca, 0x01, 0x0a, 0x0b, 0x42, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x4f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x65, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xfe, 0x01, 0x0a,
	0x0f, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65,
	0x73, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2c,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x1a,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42,
	0x79, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64,
	0x12, 0x2a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x52, 0x65, 0x73, 0x42, 0x19, 0x5a,
	0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_branches_proto_rawDescOnce sync.Once
	file_branches_proto_rawDescData = file_branches_proto_rawDesc
)

func file_branches_proto_rawDescGZIP() []byte {
	file_branches_proto_rawDescOnce.Do(func() {
		file_branches_proto_rawDescData = protoimpl.X.CompressGZIP(file_branches_proto_rawDescData)
	})
	return file_branches_proto_rawDescData
}

var file_branches_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_branches_proto_goTypes = []interface{}{
	(*CreateBranches)(nil),     // 0: proto.CreateBranches
	(*UpdateBranches)(nil),     // 1: proto.UpdateBranches
	(*BranchesRes)(nil),        // 2: proto.BranchesRes
	(*GetListBranchesReq)(nil), // 3: proto.GetListBranchesReq
	(*GetListBranchesRes)(nil), // 4: proto.GetListBranchesRes
	(*Filter)(nil),             // 5: proto.Filter
	(*ById)(nil),               // 6: proto.ById
	(*Void)(nil),               // 7: proto.Void
}
var file_branches_proto_depIdxs = []int32{
	5, // 0: proto.GetListBranchesReq.filter:type_name -> proto.Filter
	2, // 1: proto.GetListBranchesRes.branches:type_name -> proto.BranchesRes
	0, // 2: proto.BranchesService.Create:input_type -> proto.CreateBranches
	1, // 3: proto.BranchesService.Update:input_type -> proto.UpdateBranches
	6, // 4: proto.BranchesService.Delete:input_type -> proto.ById
	6, // 5: proto.BranchesService.GetById:input_type -> proto.ById
	3, // 6: proto.BranchesService.GetList:input_type -> proto.GetListBranchesReq
	7, // 7: proto.BranchesService.Create:output_type -> proto.Void
	7, // 8: proto.BranchesService.Update:output_type -> proto.Void
	7, // 9: proto.BranchesService.Delete:output_type -> proto.Void
	2, // 10: proto.BranchesService.GetById:output_type -> proto.BranchesRes
	4, // 11: proto.BranchesService.GetList:output_type -> proto.GetListBranchesRes
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_branches_proto_init() }
func file_branches_proto_init() {
	if File_branches_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_branches_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBranches); i {
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
		file_branches_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBranches); i {
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
		file_branches_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchesRes); i {
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
		file_branches_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListBranchesReq); i {
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
		file_branches_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListBranchesRes); i {
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
			RawDescriptor: file_branches_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_branches_proto_goTypes,
		DependencyIndexes: file_branches_proto_depIdxs,
		MessageInfos:      file_branches_proto_msgTypes,
	}.Build()
	File_branches_proto = out.File
	file_branches_proto_rawDesc = nil
	file_branches_proto_goTypes = nil
	file_branches_proto_depIdxs = nil
}
