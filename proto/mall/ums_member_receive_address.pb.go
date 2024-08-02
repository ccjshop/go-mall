// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: model/ums_member_receive_address.proto

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

// 会员收货地址
type MemberReceiveAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                            // 地址id
	MemberId      uint64 `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`                // 会员id
	Name          string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`                                         // 收货人名称
	PhoneNumber   string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`        // 电话号码
	DefaultStatus uint32 `protobuf:"varint,5,opt,name=default_status,json=defaultStatus,proto3" json:"default_status,omitempty"` // 是否为默认
	// 地址
	PostCode      string `protobuf:"bytes,6,opt,name=post_code,json=postCode,proto3" json:"post_code,omitempty"`                 // 邮政编码
	Province      string `protobuf:"bytes,7,opt,name=province,proto3" json:"province,omitempty"`                                 // 省份/直辖市
	City          string `protobuf:"bytes,8,opt,name=city,proto3" json:"city,omitempty"`                                         // 城市
	Region        string `protobuf:"bytes,9,opt,name=region,proto3" json:"region,omitempty"`                                     // 区
	DetailAddress string `protobuf:"bytes,10,opt,name=detail_address,json=detailAddress,proto3" json:"detail_address,omitempty"` // 详细地址(街道)
}

func (x *MemberReceiveAddress) Reset() {
	*x = MemberReceiveAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_ums_member_receive_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberReceiveAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberReceiveAddress) ProtoMessage() {}

func (x *MemberReceiveAddress) ProtoReflect() protoreflect.Message {
	mi := &file_model_ums_member_receive_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberReceiveAddress.ProtoReflect.Descriptor instead.
func (*MemberReceiveAddress) Descriptor() ([]byte, []int) {
	return file_model_ums_member_receive_address_proto_rawDescGZIP(), []int{0}
}

func (x *MemberReceiveAddress) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberReceiveAddress) GetMemberId() uint64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *MemberReceiveAddress) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MemberReceiveAddress) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *MemberReceiveAddress) GetDefaultStatus() uint32 {
	if x != nil {
		return x.DefaultStatus
	}
	return 0
}

func (x *MemberReceiveAddress) GetPostCode() string {
	if x != nil {
		return x.PostCode
	}
	return ""
}

func (x *MemberReceiveAddress) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *MemberReceiveAddress) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *MemberReceiveAddress) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *MemberReceiveAddress) GetDetailAddress() string {
	if x != nil {
		return x.DetailAddress
	}
	return ""
}

var File_model_ums_member_receive_address_proto protoreflect.FileDescriptor

var file_model_ums_member_receive_address_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x6d, 0x73, 0x5f, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22,
	0xad, 0x02, 0x0a, 0x14, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_model_ums_member_receive_address_proto_rawDescOnce sync.Once
	file_model_ums_member_receive_address_proto_rawDescData = file_model_ums_member_receive_address_proto_rawDesc
)

func file_model_ums_member_receive_address_proto_rawDescGZIP() []byte {
	file_model_ums_member_receive_address_proto_rawDescOnce.Do(func() {
		file_model_ums_member_receive_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_ums_member_receive_address_proto_rawDescData)
	})
	return file_model_ums_member_receive_address_proto_rawDescData
}

var file_model_ums_member_receive_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_ums_member_receive_address_proto_goTypes = []interface{}{
	(*MemberReceiveAddress)(nil), // 0: model.MemberReceiveAddress
}
var file_model_ums_member_receive_address_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_model_ums_member_receive_address_proto_init() }
func file_model_ums_member_receive_address_proto_init() {
	if File_model_ums_member_receive_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_ums_member_receive_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberReceiveAddress); i {
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
			RawDescriptor: file_model_ums_member_receive_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_ums_member_receive_address_proto_goTypes,
		DependencyIndexes: file_model_ums_member_receive_address_proto_depIdxs,
		MessageInfos:      file_model_ums_member_receive_address_proto_msgTypes,
	}.Build()
	File_model_ums_member_receive_address_proto = out.File
	file_model_ums_member_receive_address_proto_rawDesc = nil
	file_model_ums_member_receive_address_proto_goTypes = nil
	file_model_ums_member_receive_address_proto_depIdxs = nil
}