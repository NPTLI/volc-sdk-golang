// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: live/business/deny_config.proto

package business

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DenyConfigDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 协议类型，比如tcp,kcp,quic
	ProType []string `protobuf:"bytes,1,rep,name=ProType,proto3" json:"ProType,omitempty"`
	// 格式类型，比如http，rtmp
	FmtType []string `protobuf:"bytes,2,rep,name=FmtType,proto3" json:"FmtType,omitempty"`
	// 大洲
	Continent string `protobuf:"bytes,3,opt,name=Continent,proto3" json:"Continent,omitempty"`
	// 国家码
	Country string `protobuf:"bytes,4,opt,name=Country,proto3" json:"Country,omitempty"`
	// 区域
	Region string `protobuf:"bytes,5,opt,name=Region,proto3" json:"Region,omitempty"`
	// 城市
	City string `protobuf:"bytes,6,opt,name=City,proto3" json:"City,omitempty"`
	// 运营商
	ISP string `protobuf:"bytes,7,opt,name=ISP,proto3" json:"ISP,omitempty"`
	// 黑名单
	DenyList []string `protobuf:"bytes,8,rep,name=DenyList,proto3" json:"DenyList,omitempty"`
	// 白名单
	AllowList []string `protobuf:"bytes,9,rep,name=AllowList,proto3" json:"AllowList,omitempty"`
}

func (x *DenyConfigDetail) Reset() {
	*x = DenyConfigDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_business_deny_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DenyConfigDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenyConfigDetail) ProtoMessage() {}

func (x *DenyConfigDetail) ProtoReflect() protoreflect.Message {
	mi := &file_live_business_deny_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenyConfigDetail.ProtoReflect.Descriptor instead.
func (*DenyConfigDetail) Descriptor() ([]byte, []int) {
	return file_live_business_deny_config_proto_rawDescGZIP(), []int{0}
}

func (x *DenyConfigDetail) GetProType() []string {
	if x != nil {
		return x.ProType
	}
	return nil
}

func (x *DenyConfigDetail) GetFmtType() []string {
	if x != nil {
		return x.FmtType
	}
	return nil
}

func (x *DenyConfigDetail) GetContinent() string {
	if x != nil {
		return x.Continent
	}
	return ""
}

func (x *DenyConfigDetail) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *DenyConfigDetail) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *DenyConfigDetail) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *DenyConfigDetail) GetISP() string {
	if x != nil {
		return x.ISP
	}
	return ""
}

func (x *DenyConfigDetail) GetDenyList() []string {
	if x != nil {
		return x.DenyList
	}
	return nil
}

func (x *DenyConfigDetail) GetAllowList() []string {
	if x != nil {
		return x.AllowList
	}
	return nil
}

type DescribeDenyConfigResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 配置列表
	DenyList []*VhostWithDenyConfig `protobuf:"bytes,1,rep,name=DenyList,proto3" json:"DenyList,omitempty"`
}

func (x *DescribeDenyConfigResult) Reset() {
	*x = DescribeDenyConfigResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_business_deny_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeDenyConfigResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeDenyConfigResult) ProtoMessage() {}

func (x *DescribeDenyConfigResult) ProtoReflect() protoreflect.Message {
	mi := &file_live_business_deny_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeDenyConfigResult.ProtoReflect.Descriptor instead.
func (*DescribeDenyConfigResult) Descriptor() ([]byte, []int) {
	return file_live_business_deny_config_proto_rawDescGZIP(), []int{1}
}

func (x *DescribeDenyConfigResult) GetDenyList() []*VhostWithDenyConfig {
	if x != nil {
		return x.DenyList
	}
	return nil
}

type VhostWithDenyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 域名空间名称
	Vhost string `protobuf:"bytes,1,opt,name=Vhost,proto3" json:"Vhost,omitempty"`
	// 推拉流域名
	Domain string `protobuf:"bytes,2,opt,name=Domain,proto3" json:"Domain,omitempty"`
	// App的名称，由 1 到 30 位数字、字母、下划线及"-"和"."组成。
	App string `protobuf:"bytes,3,opt,name=App,proto3" json:"App,omitempty"`
	// 配置详情列表
	DenyConfig []*DenyConfigDetail `protobuf:"bytes,4,rep,name=DenyConfig,proto3" json:"DenyConfig,omitempty"`
}

func (x *VhostWithDenyConfig) Reset() {
	*x = VhostWithDenyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_live_business_deny_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VhostWithDenyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VhostWithDenyConfig) ProtoMessage() {}

func (x *VhostWithDenyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_live_business_deny_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VhostWithDenyConfig.ProtoReflect.Descriptor instead.
func (*VhostWithDenyConfig) Descriptor() ([]byte, []int) {
	return file_live_business_deny_config_proto_rawDescGZIP(), []int{2}
}

func (x *VhostWithDenyConfig) GetVhost() string {
	if x != nil {
		return x.Vhost
	}
	return ""
}

func (x *VhostWithDenyConfig) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *VhostWithDenyConfig) GetApp() string {
	if x != nil {
		return x.App
	}
	return ""
}

func (x *VhostWithDenyConfig) GetDenyConfig() []*DenyConfigDetail {
	if x != nil {
		return x.DenyConfig
	}
	return nil
}

var File_live_business_deny_config_proto protoreflect.FileDescriptor

var file_live_business_deny_config_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f,
	0x64, 0x65, 0x6e, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1f, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69,
	0x76, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x22, 0xf6, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x6e, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x50, 0x72, 0x6f, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x6d, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x46, 0x6d, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x43,
	0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x49, 0x53, 0x50, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x49, 0x53,
	0x50, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x6c, 0x0a, 0x18, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x65, 0x6e, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x50, 0x0a, 0x08, 0x44, 0x65, 0x6e, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x56, 0x6f, 0x6c, 0x63,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x68, 0x6f, 0x73,
	0x74, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x6e, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x08, 0x44, 0x65, 0x6e, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xa8, 0x01, 0x0a, 0x13, 0x56, 0x68,
	0x6f, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x44, 0x65, 0x6e, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x56, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x41, 0x70,
	0x70, 0x12, 0x51, 0x0a, 0x0a, 0x44, 0x65, 0x6e, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x44, 0x65, 0x6e, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0a, 0x44, 0x65, 0x6e, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0xd2, 0x01, 0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x6f, 0x6c,
	0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x6c, 0x69, 0x76, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x42, 0x0a, 0x44, 0x65, 0x6e, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50,
	0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f,
	0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x2d, 0x73, 0x64,
	0x6b, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x6c, 0x69, 0x76, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0xa0, 0x01, 0x01, 0xd8, 0x01, 0x01, 0xc2, 0x02, 0x00, 0xca, 0x02,
	0x21, 0x56, 0x6f, 0x6c, 0x63, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x4c, 0x69,
	0x76, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0xe2, 0x02, 0x24, 0x56, 0x6f, 0x6c, 0x63, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5c, 0x4c, 0x69, 0x76, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_live_business_deny_config_proto_rawDescOnce sync.Once
	file_live_business_deny_config_proto_rawDescData = file_live_business_deny_config_proto_rawDesc
)

func file_live_business_deny_config_proto_rawDescGZIP() []byte {
	file_live_business_deny_config_proto_rawDescOnce.Do(func() {
		file_live_business_deny_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_live_business_deny_config_proto_rawDescData)
	})
	return file_live_business_deny_config_proto_rawDescData
}

var file_live_business_deny_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_live_business_deny_config_proto_goTypes = []interface{}{
	(*DenyConfigDetail)(nil),         // 0: Volcengine.Live.Models.Business.DenyConfigDetail
	(*DescribeDenyConfigResult)(nil), // 1: Volcengine.Live.Models.Business.DescribeDenyConfigResult
	(*VhostWithDenyConfig)(nil),      // 2: Volcengine.Live.Models.Business.VhostWithDenyConfig
}
var file_live_business_deny_config_proto_depIdxs = []int32{
	2, // 0: Volcengine.Live.Models.Business.DescribeDenyConfigResult.DenyList:type_name -> Volcengine.Live.Models.Business.VhostWithDenyConfig
	0, // 1: Volcengine.Live.Models.Business.VhostWithDenyConfig.DenyConfig:type_name -> Volcengine.Live.Models.Business.DenyConfigDetail
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_live_business_deny_config_proto_init() }
func file_live_business_deny_config_proto_init() {
	if File_live_business_deny_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_live_business_deny_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DenyConfigDetail); i {
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
		file_live_business_deny_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeDenyConfigResult); i {
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
		file_live_business_deny_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VhostWithDenyConfig); i {
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
			RawDescriptor: file_live_business_deny_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_live_business_deny_config_proto_goTypes,
		DependencyIndexes: file_live_business_deny_config_proto_depIdxs,
		MessageInfos:      file_live_business_deny_config_proto_msgTypes,
	}.Build()
	File_live_business_deny_config_proto = out.File
	file_live_business_deny_config_proto_rawDesc = nil
	file_live_business_deny_config_proto_goTypes = nil
	file_live_business_deny_config_proto_depIdxs = nil
}
