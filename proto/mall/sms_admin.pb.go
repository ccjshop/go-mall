// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: admin/sms_admin.proto

package mall

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 添加或更新首页轮播广告表参数
type AddOrUpdateHomeAdvertiseParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`     // 主键
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`  // 名称
	Pic  string `protobuf:"bytes,3,opt,name=pic,proto3" json:"pic,omitempty"`    // 图片地址
	Url  string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`    // 链接地址
	Sort uint32 `protobuf:"varint,5,opt,name=sort,proto3" json:"sort,omitempty"` // 排序
	Note string `protobuf:"bytes,6,opt,name=note,proto3" json:"note,omitempty"`  // 备注
	// 类型
	Type uint32 `protobuf:"varint,7,opt,name=type,proto3" json:"type,omitempty"` // 轮播位置：0->PC首页轮播；1->app首页轮播，注意：在proto中使用uint32代替uint8
	// 时间
	StartTime uint32 `protobuf:"varint,8,opt,name=startTime,proto3" json:"startTime,omitempty"` // 开始时间
	EndTime   uint32 `protobuf:"varint,9,opt,name=endTime,proto3" json:"endTime,omitempty"`     // 结束时间
	// 状态
	Status uint32 `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"` // 上下线状态：0->下线；1->上线，注意：在proto中使用uint32代替uint8
	// 统计
	ClickCount uint32 `protobuf:"varint,11,opt,name=clickCount,proto3" json:"clickCount,omitempty"` // 点击数
	OrderCount uint32 `protobuf:"varint,12,opt,name=orderCount,proto3" json:"orderCount,omitempty"` // 下单数
}

func (x *AddOrUpdateHomeAdvertiseParam) Reset() {
	*x = AddOrUpdateHomeAdvertiseParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_sms_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrUpdateHomeAdvertiseParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrUpdateHomeAdvertiseParam) ProtoMessage() {}

func (x *AddOrUpdateHomeAdvertiseParam) ProtoReflect() protoreflect.Message {
	mi := &file_admin_sms_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrUpdateHomeAdvertiseParam.ProtoReflect.Descriptor instead.
func (*AddOrUpdateHomeAdvertiseParam) Descriptor() ([]byte, []int) {
	return file_admin_sms_admin_proto_rawDescGZIP(), []int{0}
}

func (x *AddOrUpdateHomeAdvertiseParam) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddOrUpdateHomeAdvertiseParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddOrUpdateHomeAdvertiseParam) GetPic() string {
	if x != nil {
		return x.Pic
	}
	return ""
}

func (x *AddOrUpdateHomeAdvertiseParam) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AddOrUpdateHomeAdvertiseParam) GetSort() uint32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *AddOrUpdateHomeAdvertiseParam) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *AddOrUpdateHomeAdvertiseParam) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *AddOrUpdateHomeAdvertiseParam) GetStartTime() uint32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *AddOrUpdateHomeAdvertiseParam) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *AddOrUpdateHomeAdvertiseParam) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddOrUpdateHomeAdvertiseParam) GetClickCount() uint32 {
	if x != nil {
		return x.ClickCount
	}
	return 0
}

func (x *AddOrUpdateHomeAdvertiseParam) GetOrderCount() uint32 {
	if x != nil {
		return x.OrderCount
	}
	return 0
}

// 分页查询首页轮播广告表
type GetHomeAdvertisesParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  uint32 `protobuf:"varint,10,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`    // 页面大小
	PageSize uint32 `protobuf:"varint,11,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"` // 页码
}

func (x *GetHomeAdvertisesParam) Reset() {
	*x = GetHomeAdvertisesParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_sms_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeAdvertisesParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeAdvertisesParam) ProtoMessage() {}

func (x *GetHomeAdvertisesParam) ProtoReflect() protoreflect.Message {
	mi := &file_admin_sms_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeAdvertisesParam.ProtoReflect.Descriptor instead.
func (*GetHomeAdvertisesParam) Descriptor() ([]byte, []int) {
	return file_admin_sms_admin_proto_rawDescGZIP(), []int{1}
}

func (x *GetHomeAdvertisesParam) GetPageNum() uint32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *GetHomeAdvertisesParam) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// 分页查询首页轮播广告表
type HomeAdvertisesData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []*HomeAdvertise `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`                             // 首页轮播广告表
	PageTotal uint32           `protobuf:"varint,2,opt,name=page_total,json=pageTotal,proto3" json:"page_total,omitempty"` // 数据总数
	PageSize  uint32           `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`    // 页码
	PageNum   uint32           `protobuf:"varint,4,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`       // 页面大小
}

func (x *HomeAdvertisesData) Reset() {
	*x = HomeAdvertisesData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_sms_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeAdvertisesData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeAdvertisesData) ProtoMessage() {}

func (x *HomeAdvertisesData) ProtoReflect() protoreflect.Message {
	mi := &file_admin_sms_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeAdvertisesData.ProtoReflect.Descriptor instead.
func (*HomeAdvertisesData) Descriptor() ([]byte, []int) {
	return file_admin_sms_admin_proto_rawDescGZIP(), []int{2}
}

func (x *HomeAdvertisesData) GetData() []*HomeAdvertise {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *HomeAdvertisesData) GetPageTotal() uint32 {
	if x != nil {
		return x.PageTotal
	}
	return 0
}

func (x *HomeAdvertisesData) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *HomeAdvertisesData) GetPageNum() uint32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

type GetHomeAdvertisesRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32              `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 状态码
	Message string              `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 提示信息
	Data    *HomeAdvertisesData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`       //
}

func (x *GetHomeAdvertisesRsp) Reset() {
	*x = GetHomeAdvertisesRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_sms_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeAdvertisesRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeAdvertisesRsp) ProtoMessage() {}

func (x *GetHomeAdvertisesRsp) ProtoReflect() protoreflect.Message {
	mi := &file_admin_sms_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeAdvertisesRsp.ProtoReflect.Descriptor instead.
func (*GetHomeAdvertisesRsp) Descriptor() ([]byte, []int) {
	return file_admin_sms_admin_proto_rawDescGZIP(), []int{3}
}

func (x *GetHomeAdvertisesRsp) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetHomeAdvertisesRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetHomeAdvertisesRsp) GetData() *HomeAdvertisesData {
	if x != nil {
		return x.Data
	}
	return nil
}

// 根据id获取首页轮播广告表
type GetHomeAdvertiseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetHomeAdvertiseReq) Reset() {
	*x = GetHomeAdvertiseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_sms_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeAdvertiseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeAdvertiseReq) ProtoMessage() {}

func (x *GetHomeAdvertiseReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_sms_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeAdvertiseReq.ProtoReflect.Descriptor instead.
func (*GetHomeAdvertiseReq) Descriptor() ([]byte, []int) {
	return file_admin_sms_admin_proto_rawDescGZIP(), []int{4}
}

func (x *GetHomeAdvertiseReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 根据id获取首页轮播广告表
type GetHomeAdvertiseRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32         `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 状态码
	Message string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 提示信息
	Data    *HomeAdvertise `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`       // 数据
}

func (x *GetHomeAdvertiseRsp) Reset() {
	*x = GetHomeAdvertiseRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_sms_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHomeAdvertiseRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHomeAdvertiseRsp) ProtoMessage() {}

func (x *GetHomeAdvertiseRsp) ProtoReflect() protoreflect.Message {
	mi := &file_admin_sms_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHomeAdvertiseRsp.ProtoReflect.Descriptor instead.
func (*GetHomeAdvertiseRsp) Descriptor() ([]byte, []int) {
	return file_admin_sms_admin_proto_rawDescGZIP(), []int{5}
}

func (x *GetHomeAdvertiseRsp) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetHomeAdvertiseRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetHomeAdvertiseRsp) GetData() *HomeAdvertise {
	if x != nil {
		return x.Data
	}
	return nil
}

// 删除首页轮播广告表
type DeleteHomeAdvertiseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteHomeAdvertiseReq) Reset() {
	*x = DeleteHomeAdvertiseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_sms_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHomeAdvertiseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHomeAdvertiseReq) ProtoMessage() {}

func (x *DeleteHomeAdvertiseReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_sms_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHomeAdvertiseReq.ProtoReflect.Descriptor instead.
func (*DeleteHomeAdvertiseReq) Descriptor() ([]byte, []int) {
	return file_admin_sms_admin_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteHomeAdvertiseReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_admin_sms_admin_proto protoreflect.FileDescriptor

var file_admin_sms_admin_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x6d, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x73, 0x6d, 0x73, 0x5f, 0x68, 0x6f, 0x6d, 0x65, 0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x02, 0x0a, 0x1d, 0x41, 0x64, 0x64,
	0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x69, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x63,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x62,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69,
	0x73, 0x65, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a,
	0x02, 0x28, 0x00, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x24, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x12, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x73, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0x73, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x73, 0x52,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x73, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d,
	0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48,
	0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x32,
	0xa2, 0x04, 0x0a, 0x0b, 0x53, 0x6d, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x70, 0x69, 0x12,
	0x69, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76,
	0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x12, 0x24, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41,
	0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x10, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x68, 0x6f, 0x6d, 0x65,
	0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x73, 0x12, 0x6e, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73,
	0x65, 0x12, 0x24, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69,
	0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x10, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x3a, 0x01, 0x2a, 0x1a, 0x14, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x68, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x73, 0x12,
	0x1d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x41,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x1b,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64,
	0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x73, 0x52, 0x73, 0x70, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x69, 0x73, 0x65, 0x73, 0x12, 0x68, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x41,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x12, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x52, 0x73, 0x70,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x41,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x64,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65,
	0x72, 0x74, 0x69, 0x73, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14,
	0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_sms_admin_proto_rawDescOnce sync.Once
	file_admin_sms_admin_proto_rawDescData = file_admin_sms_admin_proto_rawDesc
)

func file_admin_sms_admin_proto_rawDescGZIP() []byte {
	file_admin_sms_admin_proto_rawDescOnce.Do(func() {
		file_admin_sms_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_sms_admin_proto_rawDescData)
	})
	return file_admin_sms_admin_proto_rawDescData
}

var file_admin_sms_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_admin_sms_admin_proto_goTypes = []any{
	(*AddOrUpdateHomeAdvertiseParam)(nil), // 0: admin.AddOrUpdateHomeAdvertiseParam
	(*GetHomeAdvertisesParam)(nil),        // 1: admin.GetHomeAdvertisesParam
	(*HomeAdvertisesData)(nil),            // 2: admin.HomeAdvertisesData
	(*GetHomeAdvertisesRsp)(nil),          // 3: admin.GetHomeAdvertisesRsp
	(*GetHomeAdvertiseReq)(nil),           // 4: admin.GetHomeAdvertiseReq
	(*GetHomeAdvertiseRsp)(nil),           // 5: admin.GetHomeAdvertiseRsp
	(*DeleteHomeAdvertiseReq)(nil),        // 6: admin.DeleteHomeAdvertiseReq
	(*HomeAdvertise)(nil),                 // 7: model.HomeAdvertise
	(*CommonRsp)(nil),                     // 8: admin.CommonRsp
}
var file_admin_sms_admin_proto_depIdxs = []int32{
	7, // 0: admin.HomeAdvertisesData.data:type_name -> model.HomeAdvertise
	2, // 1: admin.GetHomeAdvertisesRsp.data:type_name -> admin.HomeAdvertisesData
	7, // 2: admin.GetHomeAdvertiseRsp.data:type_name -> model.HomeAdvertise
	0, // 3: admin.SmsAdminApi.CreateHomeAdvertise:input_type -> admin.AddOrUpdateHomeAdvertiseParam
	0, // 4: admin.SmsAdminApi.UpdateHomeAdvertise:input_type -> admin.AddOrUpdateHomeAdvertiseParam
	1, // 5: admin.SmsAdminApi.GetHomeAdvertises:input_type -> admin.GetHomeAdvertisesParam
	4, // 6: admin.SmsAdminApi.GetHomeAdvertise:input_type -> admin.GetHomeAdvertiseReq
	6, // 7: admin.SmsAdminApi.DeleteHomeAdvertise:input_type -> admin.DeleteHomeAdvertiseReq
	8, // 8: admin.SmsAdminApi.CreateHomeAdvertise:output_type -> admin.CommonRsp
	8, // 9: admin.SmsAdminApi.UpdateHomeAdvertise:output_type -> admin.CommonRsp
	3, // 10: admin.SmsAdminApi.GetHomeAdvertises:output_type -> admin.GetHomeAdvertisesRsp
	5, // 11: admin.SmsAdminApi.GetHomeAdvertise:output_type -> admin.GetHomeAdvertiseRsp
	8, // 12: admin.SmsAdminApi.DeleteHomeAdvertise:output_type -> admin.CommonRsp
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_admin_sms_admin_proto_init() }
func file_admin_sms_admin_proto_init() {
	if File_admin_sms_admin_proto != nil {
		return
	}
	file_admin_admin_proto_init()
	file_model_sms_home_advertise_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_admin_sms_admin_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddOrUpdateHomeAdvertiseParam); i {
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
		file_admin_sms_admin_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetHomeAdvertisesParam); i {
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
		file_admin_sms_admin_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*HomeAdvertisesData); i {
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
		file_admin_sms_admin_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetHomeAdvertisesRsp); i {
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
		file_admin_sms_admin_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetHomeAdvertiseReq); i {
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
		file_admin_sms_admin_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetHomeAdvertiseRsp); i {
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
		file_admin_sms_admin_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteHomeAdvertiseReq); i {
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
			RawDescriptor: file_admin_sms_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_sms_admin_proto_goTypes,
		DependencyIndexes: file_admin_sms_admin_proto_depIdxs,
		MessageInfos:      file_admin_sms_admin_proto_msgTypes,
	}.Build()
	File_admin_sms_admin_proto = out.File
	file_admin_sms_admin_proto_rawDesc = nil
	file_admin_sms_admin_proto_goTypes = nil
	file_admin_sms_admin_proto_depIdxs = nil
}
