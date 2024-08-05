// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: model/sms_coupon.proto

package mall

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

// 使用类型
type CouponUseType int32

const (
	CouponUseType_COUPON_USE_TYPE_GENERAL           CouponUseType = 0 // 全场通用
	CouponUseType_COUPON_USE_TYPE_SPECIFIC_CATEGORY CouponUseType = 1 // 指定分类
	CouponUseType_COUPON_USE_TYPE_SPECIFIC_PRODUCT  CouponUseType = 2 // 指定商品
)

// Enum value maps for CouponUseType.
var (
	CouponUseType_name = map[int32]string{
		0: "COUPON_USE_TYPE_GENERAL",
		1: "COUPON_USE_TYPE_SPECIFIC_CATEGORY",
		2: "COUPON_USE_TYPE_SPECIFIC_PRODUCT",
	}
	CouponUseType_value = map[string]int32{
		"COUPON_USE_TYPE_GENERAL":           0,
		"COUPON_USE_TYPE_SPECIFIC_CATEGORY": 1,
		"COUPON_USE_TYPE_SPECIFIC_PRODUCT":  2,
	}
)

func (x CouponUseType) Enum() *CouponUseType {
	p := new(CouponUseType)
	*p = x
	return p
}

func (x CouponUseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CouponUseType) Descriptor() protoreflect.EnumDescriptor {
	return file_model_sms_coupon_proto_enumTypes[0].Descriptor()
}

func (CouponUseType) Type() protoreflect.EnumType {
	return &file_model_sms_coupon_proto_enumTypes[0]
}

func (x CouponUseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CouponUseType.Descriptor instead.
func (CouponUseType) EnumDescriptor() ([]byte, []int) {
	return file_model_sms_coupon_proto_rawDescGZIP(), []int{0}
}

// Coupon 优惠券表
// 用于存储优惠券信息，需要注意的是优惠券的使用类型：0->全场通用；1->指定分类；2->指定商品，不同使用类型的优惠券使用范围不一样。
type Coupon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基本信息
	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`        //
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`     // 名称
	Amount string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"` // 金额，保留两位小数
	Note   string `protobuf:"bytes,4,opt,name=note,proto3" json:"note,omitempty"`     // 备注
	Code   string `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`     // 优惠码
	// 数量
	Count        uint32 `protobuf:"varint,6,opt,name=count,proto3" json:"count,omitempty"`                                   // 数量
	PublishCount uint32 `protobuf:"varint,7,opt,name=publish_count,json=publishCount,proto3" json:"publish_count,omitempty"` // 发行数量
	UseCount     uint32 `protobuf:"varint,8,opt,name=use_count,json=useCount,proto3" json:"use_count,omitempty"`             // 已使用数量
	ReceiveCount uint32 `protobuf:"varint,9,opt,name=receive_count,json=receiveCount,proto3" json:"receive_count,omitempty"` // 领取数量
	// 类型
	Type    uint32        `protobuf:"varint,10,opt,name=type,proto3" json:"type,omitempty"`                                               // 优惠卷类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券
	UseType CouponUseType `protobuf:"varint,11,opt,name=use_type,json=useType,proto3,enum=model.CouponUseType" json:"use_type,omitempty"` // 使用类型：0->全场通用；1->指定分类；2->指定商品
	// 领取限制
	PerLimit    uint32 `protobuf:"varint,12,opt,name=per_limit,json=perLimit,proto3" json:"per_limit,omitempty"`          // 每人限领张数
	EnableTime  uint32 `protobuf:"varint,13,opt,name=enable_time,json=enableTime,proto3" json:"enable_time,omitempty"`    // 可以领取的日期
	MemberLevel uint32 `protobuf:"varint,14,opt,name=member_level,json=memberLevel,proto3" json:"member_level,omitempty"` // 可领取的会员类型：0->无限制
	// 使用限制
	MinPoint  string `protobuf:"bytes,15,opt,name=min_point,json=minPoint,proto3" json:"min_point,omitempty"`     // 使用门槛；0表示无门槛
	Platform  uint32 `protobuf:"varint,16,opt,name=platform,proto3" json:"platform,omitempty"`                    // 使用平台：0->全部；1->移动；2->PC
	StartTime uint32 `protobuf:"varint,17,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"` // 开始使用时间
	EndTime   uint32 `protobuf:"varint,18,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`       // 结束使用时间
}

func (x *Coupon) Reset() {
	*x = Coupon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_sms_coupon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coupon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coupon) ProtoMessage() {}

func (x *Coupon) ProtoReflect() protoreflect.Message {
	mi := &file_model_sms_coupon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coupon.ProtoReflect.Descriptor instead.
func (*Coupon) Descriptor() ([]byte, []int) {
	return file_model_sms_coupon_proto_rawDescGZIP(), []int{0}
}

func (x *Coupon) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Coupon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Coupon) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Coupon) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Coupon) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Coupon) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Coupon) GetPublishCount() uint32 {
	if x != nil {
		return x.PublishCount
	}
	return 0
}

func (x *Coupon) GetUseCount() uint32 {
	if x != nil {
		return x.UseCount
	}
	return 0
}

func (x *Coupon) GetReceiveCount() uint32 {
	if x != nil {
		return x.ReceiveCount
	}
	return 0
}

func (x *Coupon) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Coupon) GetUseType() CouponUseType {
	if x != nil {
		return x.UseType
	}
	return CouponUseType_COUPON_USE_TYPE_GENERAL
}

func (x *Coupon) GetPerLimit() uint32 {
	if x != nil {
		return x.PerLimit
	}
	return 0
}

func (x *Coupon) GetEnableTime() uint32 {
	if x != nil {
		return x.EnableTime
	}
	return 0
}

func (x *Coupon) GetMemberLevel() uint32 {
	if x != nil {
		return x.MemberLevel
	}
	return 0
}

func (x *Coupon) GetMinPoint() string {
	if x != nil {
		return x.MinPoint
	}
	return ""
}

func (x *Coupon) GetPlatform() uint32 {
	if x != nil {
		return x.Platform
	}
	return 0
}

func (x *Coupon) GetStartTime() uint32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Coupon) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

var File_model_sms_coupon_proto protoreflect.FileDescriptor

var file_model_sms_coupon_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x6d, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22,
	0x82, 0x04, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x2f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x55, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x75, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x2a, 0x79, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x55, 0x73,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x55, 0x50, 0x4f, 0x4e, 0x5f,
	0x55, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x45, 0x4e, 0x45, 0x52, 0x41, 0x4c,
	0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x4f, 0x55, 0x50, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x43, 0x5f, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x10, 0x01, 0x12, 0x24, 0x0a, 0x20, 0x43, 0x4f, 0x55,
	0x50, 0x4f, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x43, 0x5f, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x54, 0x10, 0x02, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_model_sms_coupon_proto_rawDescOnce sync.Once
	file_model_sms_coupon_proto_rawDescData = file_model_sms_coupon_proto_rawDesc
)

func file_model_sms_coupon_proto_rawDescGZIP() []byte {
	file_model_sms_coupon_proto_rawDescOnce.Do(func() {
		file_model_sms_coupon_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_sms_coupon_proto_rawDescData)
	})
	return file_model_sms_coupon_proto_rawDescData
}

var file_model_sms_coupon_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_model_sms_coupon_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_sms_coupon_proto_goTypes = []any{
	(CouponUseType)(0), // 0: model.CouponUseType
	(*Coupon)(nil),     // 1: model.Coupon
}
var file_model_sms_coupon_proto_depIdxs = []int32{
	0, // 0: model.Coupon.use_type:type_name -> model.CouponUseType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_model_sms_coupon_proto_init() }
func file_model_sms_coupon_proto_init() {
	if File_model_sms_coupon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_sms_coupon_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Coupon); i {
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
			RawDescriptor: file_model_sms_coupon_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_sms_coupon_proto_goTypes,
		DependencyIndexes: file_model_sms_coupon_proto_depIdxs,
		EnumInfos:         file_model_sms_coupon_proto_enumTypes,
		MessageInfos:      file_model_sms_coupon_proto_msgTypes,
	}.Build()
	File_model_sms_coupon_proto = out.File
	file_model_sms_coupon_proto_rawDesc = nil
	file_model_sms_coupon_proto_goTypes = nil
	file_model_sms_coupon_proto_depIdxs = nil
}
