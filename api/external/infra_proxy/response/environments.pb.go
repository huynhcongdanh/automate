// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/infra_proxy/response/environments.proto

package response

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Environments struct {
	// Environments list.
	Environments         []*EnvironmentListItem `protobuf:"bytes,1,rep,name=environments,proto3" json:"environments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Environments) Reset()         { *m = Environments{} }
func (m *Environments) String() string { return proto.CompactTextString(m) }
func (*Environments) ProtoMessage()    {}
func (*Environments) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6cb2a4fff0e747c, []int{0}
}

func (m *Environments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Environments.Unmarshal(m, b)
}
func (m *Environments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Environments.Marshal(b, m, deterministic)
}
func (m *Environments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Environments.Merge(m, src)
}
func (m *Environments) XXX_Size() int {
	return xxx_messageInfo_Environments.Size(m)
}
func (m *Environments) XXX_DiscardUnknown() {
	xxx_messageInfo_Environments.DiscardUnknown(m)
}

var xxx_messageInfo_Environments proto.InternalMessageInfo

func (m *Environments) GetEnvironments() []*EnvironmentListItem {
	if m != nil {
		return m.Environments
	}
	return nil
}

type EnvironmentListItem struct {
	// Environment name.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvironmentListItem) Reset()         { *m = EnvironmentListItem{} }
func (m *EnvironmentListItem) String() string { return proto.CompactTextString(m) }
func (*EnvironmentListItem) ProtoMessage()    {}
func (*EnvironmentListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6cb2a4fff0e747c, []int{1}
}

func (m *EnvironmentListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvironmentListItem.Unmarshal(m, b)
}
func (m *EnvironmentListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvironmentListItem.Marshal(b, m, deterministic)
}
func (m *EnvironmentListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvironmentListItem.Merge(m, src)
}
func (m *EnvironmentListItem) XXX_Size() int {
	return xxx_messageInfo_EnvironmentListItem.Size(m)
}
func (m *EnvironmentListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvironmentListItem.DiscardUnknown(m)
}

var xxx_messageInfo_EnvironmentListItem proto.InternalMessageInfo

func (m *EnvironmentListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Environment struct {
	// Environment name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Chef object type.
	ChefType string `protobuf:"bytes,2,opt,name=chef_type,json=chefType,proto3" json:"chef_type,omitempty"`
	// Environment description.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Environment JSON class.
	JsonClass string `protobuf:"bytes,4,opt,name=json_class,json=jsonClass,proto3" json:"json_class,omitempty"`
	// Environment versined cookbooks constraints.
	CookbookVersions map[string]string `protobuf:"bytes,5,rep,name=cookbook_versions,json=cookbookVersions,proto3" json:"cookbook_versions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Stringified json of the default attributes.
	DefaultAttributes string `protobuf:"bytes,6,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty"`
	// Stringified json of the override attributes.
	OverrideAttributes   string   `protobuf:"bytes,7,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Environment) Reset()         { *m = Environment{} }
func (m *Environment) String() string { return proto.CompactTextString(m) }
func (*Environment) ProtoMessage()    {}
func (*Environment) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6cb2a4fff0e747c, []int{2}
}

func (m *Environment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Environment.Unmarshal(m, b)
}
func (m *Environment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Environment.Marshal(b, m, deterministic)
}
func (m *Environment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Environment.Merge(m, src)
}
func (m *Environment) XXX_Size() int {
	return xxx_messageInfo_Environment.Size(m)
}
func (m *Environment) XXX_DiscardUnknown() {
	xxx_messageInfo_Environment.DiscardUnknown(m)
}

var xxx_messageInfo_Environment proto.InternalMessageInfo

func (m *Environment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Environment) GetChefType() string {
	if m != nil {
		return m.ChefType
	}
	return ""
}

func (m *Environment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Environment) GetJsonClass() string {
	if m != nil {
		return m.JsonClass
	}
	return ""
}

func (m *Environment) GetCookbookVersions() map[string]string {
	if m != nil {
		return m.CookbookVersions
	}
	return nil
}

func (m *Environment) GetDefaultAttributes() string {
	if m != nil {
		return m.DefaultAttributes
	}
	return ""
}

func (m *Environment) GetOverrideAttributes() string {
	if m != nil {
		return m.OverrideAttributes
	}
	return ""
}

func init() {
	proto.RegisterType((*Environments)(nil), "chef.automate.api.infra_proxy.response.Environments")
	proto.RegisterType((*EnvironmentListItem)(nil), "chef.automate.api.infra_proxy.response.EnvironmentListItem")
	proto.RegisterType((*Environment)(nil), "chef.automate.api.infra_proxy.response.Environment")
	proto.RegisterMapType((map[string]string)(nil), "chef.automate.api.infra_proxy.response.Environment.CookbookVersionsEntry")
}

func init() {
	proto.RegisterFile("api/external/infra_proxy/response/environments.proto", fileDescriptor_b6cb2a4fff0e747c)
}

var fileDescriptor_b6cb2a4fff0e747c = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x25, 0x2f, 0xef, 0x3d, 0xed, 0xed, 0x5b, 0xb4, 0x53, 0x85, 0xa0, 0x08, 0xa5, 0x0b, 0xa9,
	0x0b, 0x67, 0x40, 0x5d, 0x48, 0x75, 0xa3, 0xa5, 0x8b, 0x82, 0xab, 0x22, 0x2e, 0xdc, 0x84, 0x49,
	0x7a, 0x6b, 0xc7, 0x26, 0x73, 0x87, 0x99, 0x49, 0x68, 0xfe, 0x8b, 0x3f, 0x56, 0x32, 0x26, 0x1a,
	0xa5, 0xa0, 0xb8, 0x9b, 0x39, 0x5f, 0x1c, 0x0e, 0x17, 0x5e, 0x49, 0xa3, 0x04, 0x9e, 0x3d, 0x5a,
	0x2d, 0x0b, 0xa1, 0xf4, 0xc1, 0xca, 0xd4, 0x58, 0x3a, 0x37, 0xc2, 0xa2, 0x33, 0xa4, 0x1d, 0x0a,
	0xd4, 0xb5, 0xb2, 0xa4, 0x4b, 0xd4, 0xde, 0x71, 0x63, 0xc9, 0x13, 0x7b, 0x9a, 0x1f, 0xf1, 0xc0,
	0x65, 0xe5, 0xa9, 0x94, 0x1e, 0xb9, 0x34, 0x8a, 0x0f, 0xac, 0xbc, 0xb7, 0x2e, 0x08, 0xee, 0x36,
	0x03, 0x37, 0x4b, 0xe1, 0x6e, 0x98, 0x96, 0x44, 0xf3, 0x78, 0x39, 0x7e, 0xf1, 0x86, 0xff, 0x5b,
	0x1c, 0x1f, 0x64, 0x7d, 0x50, 0xce, 0x6f, 0x3d, 0x96, 0xbb, 0xdf, 0x02, 0x17, 0xcf, 0x60, 0x76,
	0x41, 0xc4, 0x18, 0x5c, 0x6b, 0x59, 0x62, 0x12, 0xcd, 0xa3, 0xe5, 0x68, 0x17, 0xde, 0x8b, 0x6f,
	0x31, 0x8c, 0x07, 0xda, 0x4b, 0x1a, 0xf6, 0x18, 0x46, 0x6d, 0xb5, 0xd4, 0x37, 0x06, 0x93, 0xab,
	0x40, 0xdc, 0x6f, 0x81, 0x8f, 0x8d, 0x41, 0x36, 0x87, 0xf1, 0x1e, 0x5d, 0x6e, 0x95, 0xf1, 0x8a,
	0x74, 0x12, 0x07, 0x7a, 0x08, 0xb1, 0x27, 0x00, 0x5f, 0x1d, 0xe9, 0x34, 0x2f, 0xa4, 0x73, 0xc9,
	0x75, 0x10, 0x8c, 0x5a, 0x64, 0xdd, 0x02, 0xac, 0x86, 0x69, 0x4e, 0x74, 0xca, 0x88, 0x4e, 0x69,
	0x8d, 0xd6, 0x29, 0xd2, 0x2e, 0xb9, 0x09, 0x93, 0x6c, 0xff, 0x63, 0x12, 0xbe, 0xee, 0xc2, 0x3e,
	0x75, 0x59, 0x1b, 0xed, 0x6d, 0xb3, 0x9b, 0xe4, 0x7f, 0xc0, 0xec, 0x39, 0xb0, 0x3d, 0x1e, 0x64,
	0x55, 0xf8, 0x54, 0x7a, 0x6f, 0x55, 0x56, 0x79, 0x74, 0xc9, 0x6d, 0xa8, 0x37, 0xed, 0x98, 0x77,
	0x3f, 0x09, 0x26, 0x60, 0x46, 0x35, 0x5a, 0xab, 0xf6, 0x38, 0xd4, 0xdf, 0x0b, 0x7a, 0xd6, 0x53,
	0xbf, 0x0c, 0x8f, 0xd6, 0xf0, 0xf0, 0x62, 0x15, 0x36, 0x81, 0xf8, 0x84, 0x4d, 0xb7, 0x70, 0xfb,
	0x64, 0x0f, 0xe0, 0xa6, 0x96, 0x45, 0xd5, 0x8f, 0xfb, 0xe3, 0xb3, 0xba, 0x7a, 0x1d, 0xbd, 0x7f,
	0xfb, 0x79, 0xf5, 0x45, 0xf9, 0x63, 0x95, 0xf1, 0x9c, 0x4a, 0xd1, 0xae, 0x21, 0xfa, 0x35, 0xc4,
	0x5f, 0x6f, 0x36, 0xbb, 0x0d, 0x77, 0xfa, 0xf2, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xec,
	0x43, 0xdc, 0xdf, 0x02, 0x00, 0x00,
}
