// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: course_item.proto

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

type CreateCourseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId      string  `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Description   string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price         float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	DaysPerWeek   int32   `protobuf:"varint,4,opt,name=days_per_week,json=daysPerWeek,proto3" json:"days_per_week,omitempty"`
	LessonHours   float32 `protobuf:"fixed32,5,opt,name=lesson_hours,json=lessonHours,proto3" json:"lesson_hours,omitempty"`
	DurationWeeks float32 `protobuf:"fixed32,6,opt,name=duration_weeks,json=durationWeeks,proto3" json:"duration_weeks,omitempty"`
}

func (x *CreateCourseItem) Reset() {
	*x = CreateCourseItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCourseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourseItem) ProtoMessage() {}

func (x *CreateCourseItem) ProtoReflect() protoreflect.Message {
	mi := &file_course_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourseItem.ProtoReflect.Descriptor instead.
func (*CreateCourseItem) Descriptor() ([]byte, []int) {
	return file_course_item_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCourseItem) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *CreateCourseItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateCourseItem) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateCourseItem) GetDaysPerWeek() int32 {
	if x != nil {
		return x.DaysPerWeek
	}
	return 0
}

func (x *CreateCourseItem) GetLessonHours() float32 {
	if x != nil {
		return x.LessonHours
	}
	return 0
}

func (x *CreateCourseItem) GetDurationWeeks() float32 {
	if x != nil {
		return x.DurationWeeks
	}
	return 0
}

type UpdateCourseItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CourseId      string  `protobuf:"bytes,2,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Description   string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	DaysPerWeek   int32   `protobuf:"varint,5,opt,name=days_per_week,json=daysPerWeek,proto3" json:"days_per_week,omitempty"`
	LessonHours   float32 `protobuf:"fixed32,6,opt,name=lesson_hours,json=lessonHours,proto3" json:"lesson_hours,omitempty"`
	DurationWeeks float32 `protobuf:"fixed32,7,opt,name=duration_weeks,json=durationWeeks,proto3" json:"duration_weeks,omitempty"`
}

func (x *UpdateCourseItem) Reset() {
	*x = UpdateCourseItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCourseItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCourseItem) ProtoMessage() {}

func (x *UpdateCourseItem) ProtoReflect() protoreflect.Message {
	mi := &file_course_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCourseItem.ProtoReflect.Descriptor instead.
func (*UpdateCourseItem) Descriptor() ([]byte, []int) {
	return file_course_item_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateCourseItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCourseItem) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *UpdateCourseItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateCourseItem) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateCourseItem) GetDaysPerWeek() int32 {
	if x != nil {
		return x.DaysPerWeek
	}
	return 0
}

func (x *UpdateCourseItem) GetLessonHours() float32 {
	if x != nil {
		return x.LessonHours
	}
	return 0
}

func (x *UpdateCourseItem) GetDurationWeeks() float32 {
	if x != nil {
		return x.DurationWeeks
	}
	return 0
}

type CourseItemRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CourseId      string  `protobuf:"bytes,2,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Description   string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	DaysPerWeek   int32   `protobuf:"varint,5,opt,name=days_per_week,json=daysPerWeek,proto3" json:"days_per_week,omitempty"`
	LessonHours   float32 `protobuf:"fixed32,6,opt,name=lesson_hours,json=lessonHours,proto3" json:"lesson_hours,omitempty"`
	DurationWeeks float32 `protobuf:"fixed32,7,opt,name=duration_weeks,json=durationWeeks,proto3" json:"duration_weeks,omitempty"`
	CreatedAt     string  `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *CourseItemRes) Reset() {
	*x = CourseItemRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseItemRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseItemRes) ProtoMessage() {}

func (x *CourseItemRes) ProtoReflect() protoreflect.Message {
	mi := &file_course_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseItemRes.ProtoReflect.Descriptor instead.
func (*CourseItemRes) Descriptor() ([]byte, []int) {
	return file_course_item_proto_rawDescGZIP(), []int{2}
}

func (x *CourseItemRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CourseItemRes) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *CourseItemRes) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CourseItemRes) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CourseItemRes) GetDaysPerWeek() int32 {
	if x != nil {
		return x.DaysPerWeek
	}
	return 0
}

func (x *CourseItemRes) GetLessonHours() float32 {
	if x != nil {
		return x.LessonHours
	}
	return 0
}

func (x *CourseItemRes) GetDurationWeeks() float32 {
	if x != nil {
		return x.DurationWeeks
	}
	return 0
}

func (x *CourseItemRes) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetListCourseItemReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId    string  `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	MaxPrice    float32 `protobuf:"fixed32,2,opt,name=max_price,json=maxPrice,proto3" json:"max_price,omitempty"`
	MinPrice    float32 `protobuf:"fixed32,3,opt,name=min_price,json=minPrice,proto3" json:"min_price,omitempty"`
	DaysPerWeek int32   `protobuf:"varint,4,opt,name=days_per_week,json=daysPerWeek,proto3" json:"days_per_week,omitempty"`
	LessonHours int32   `protobuf:"varint,5,opt,name=lesson_hours,json=lessonHours,proto3" json:"lesson_hours,omitempty"`
	Filter      *Filter `protobuf:"bytes,6,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *GetListCourseItemReq) Reset() {
	*x = GetListCourseItemReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListCourseItemReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListCourseItemReq) ProtoMessage() {}

func (x *GetListCourseItemReq) ProtoReflect() protoreflect.Message {
	mi := &file_course_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListCourseItemReq.ProtoReflect.Descriptor instead.
func (*GetListCourseItemReq) Descriptor() ([]byte, []int) {
	return file_course_item_proto_rawDescGZIP(), []int{3}
}

func (x *GetListCourseItemReq) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *GetListCourseItemReq) GetMaxPrice() float32 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

func (x *GetListCourseItemReq) GetMinPrice() float32 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *GetListCourseItemReq) GetDaysPerWeek() int32 {
	if x != nil {
		return x.DaysPerWeek
	}
	return 0
}

func (x *GetListCourseItemReq) GetLessonHours() int32 {
	if x != nil {
		return x.LessonHours
	}
	return 0
}

func (x *GetListCourseItemReq) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type GetListCourseItemRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseItems []*CourseItemRes `protobuf:"bytes,1,rep,name=course_items,json=courseItems,proto3" json:"course_items,omitempty"`
	TotalCount  int32            `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *GetListCourseItemRes) Reset() {
	*x = GetListCourseItemRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListCourseItemRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListCourseItemRes) ProtoMessage() {}

func (x *GetListCourseItemRes) ProtoReflect() protoreflect.Message {
	mi := &file_course_item_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListCourseItemRes.ProtoReflect.Descriptor instead.
func (*GetListCourseItemRes) Descriptor() ([]byte, []int) {
	return file_course_item_proto_rawDescGZIP(), []int{4}
}

func (x *GetListCourseItemRes) GetCourseItems() []*CourseItemRes {
	if x != nil {
		return x.CourseItems
	}
	return nil
}

func (x *GetListCourseItemRes) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_course_item_proto protoreflect.FileDescriptor

var file_course_item_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x77,
	0x65, 0x65, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x61, 0x79, 0x73, 0x50,
	0x65, 0x72, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0d, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x65, 0x65, 0x6b, 0x73,
	0x22, 0xe5, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61,
	0x79, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x64, 0x61, 0x79, 0x73, 0x50, 0x65, 0x72, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x21,
	0x0a, 0x0c, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x48, 0x6f, 0x75, 0x72,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x65,
	0x65, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x57, 0x65, 0x65, 0x6b, 0x73, 0x22, 0x81, 0x02, 0x0a, 0x0d, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x22, 0x0a, 0x0d, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x77, 0x65, 0x65, 0x6b,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x61, 0x79, 0x73, 0x50, 0x65, 0x72, 0x57,
	0x65, 0x65, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x68, 0x6f,
	0x75, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x65, 0x65, 0x6b, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xdb, 0x01, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0d,
	0x64, 0x61, 0x79, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x61, 0x79, 0x73, 0x50, 0x65, 0x72, 0x57, 0x65, 0x65, 0x6b,
	0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x48, 0x6f,
	0x75, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x70, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x12, 0x37, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x52, 0x0b,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x8a, 0x02, 0x0a,
	0x11, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x2e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x73, 0x12, 0x43, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x42, 0x19, 0x5a, 0x17, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_course_item_proto_rawDescOnce sync.Once
	file_course_item_proto_rawDescData = file_course_item_proto_rawDesc
)

func file_course_item_proto_rawDescGZIP() []byte {
	file_course_item_proto_rawDescOnce.Do(func() {
		file_course_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_course_item_proto_rawDescData)
	})
	return file_course_item_proto_rawDescData
}

var file_course_item_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_course_item_proto_goTypes = []interface{}{
	(*CreateCourseItem)(nil),     // 0: proto.CreateCourseItem
	(*UpdateCourseItem)(nil),     // 1: proto.UpdateCourseItem
	(*CourseItemRes)(nil),        // 2: proto.CourseItemRes
	(*GetListCourseItemReq)(nil), // 3: proto.GetListCourseItemReq
	(*GetListCourseItemRes)(nil), // 4: proto.GetListCourseItemRes
	(*Filter)(nil),               // 5: proto.Filter
	(*ById)(nil),                 // 6: proto.ById
	(*Void)(nil),                 // 7: proto.Void
}
var file_course_item_proto_depIdxs = []int32{
	5, // 0: proto.GetListCourseItemReq.filter:type_name -> proto.Filter
	2, // 1: proto.GetListCourseItemRes.course_items:type_name -> proto.CourseItemRes
	0, // 2: proto.CourseItemService.Create:input_type -> proto.CreateCourseItem
	1, // 3: proto.CourseItemService.Update:input_type -> proto.UpdateCourseItem
	6, // 4: proto.CourseItemService.Delete:input_type -> proto.ById
	6, // 5: proto.CourseItemService.GetById:input_type -> proto.ById
	3, // 6: proto.CourseItemService.GetList:input_type -> proto.GetListCourseItemReq
	7, // 7: proto.CourseItemService.Create:output_type -> proto.Void
	7, // 8: proto.CourseItemService.Update:output_type -> proto.Void
	7, // 9: proto.CourseItemService.Delete:output_type -> proto.Void
	2, // 10: proto.CourseItemService.GetById:output_type -> proto.CourseItemRes
	4, // 11: proto.CourseItemService.GetList:output_type -> proto.GetListCourseItemRes
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_course_item_proto_init() }
func file_course_item_proto_init() {
	if File_course_item_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_course_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCourseItem); i {
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
		file_course_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCourseItem); i {
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
		file_course_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseItemRes); i {
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
		file_course_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListCourseItemReq); i {
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
		file_course_item_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListCourseItemRes); i {
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
			RawDescriptor: file_course_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_course_item_proto_goTypes,
		DependencyIndexes: file_course_item_proto_depIdxs,
		MessageInfos:      file_course_item_proto_msgTypes,
	}.Build()
	File_course_item_proto = out.File
	file_course_item_proto_rawDesc = nil
	file_course_item_proto_goTypes = nil
	file_course_item_proto_depIdxs = nil
}
