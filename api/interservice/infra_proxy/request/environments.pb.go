// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/infra_proxy/request/environments.proto

package request

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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
	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId             string   `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Environments) Reset()         { *m = Environments{} }
func (m *Environments) String() string { return proto.CompactTextString(m) }
func (*Environments) ProtoMessage()    {}
func (*Environments) Descriptor() ([]byte, []int) {
	return fileDescriptor_e190c707558fe428, []int{0}
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

func (m *Environments) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *Environments) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

type Environment struct {
	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Environment name.
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Environment) Reset()         { *m = Environment{} }
func (m *Environment) String() string { return proto.CompactTextString(m) }
func (*Environment) ProtoMessage()    {}
func (*Environment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e190c707558fe428, []int{1}
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

func (m *Environment) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *Environment) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *Environment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateEnvironment struct {
	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Environment name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Environment description.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty" toml:"description,omitempty" mapstructure:"description,omitempty"`
	// Class name.
	JsonClass string `protobuf:"bytes,5,opt,name=json_class,json=jsonClass,proto3" json:"json_class,omitempty" toml:"json_class,omitempty" mapstructure:"json_class,omitempty"`
	// Environment versioned cookbooks constraints.
	CookbookVersions map[string]string `protobuf:"bytes,6,rep,name=cookbook_versions,json=cookbookVersions,proto3" json:"cookbook_versions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" toml:"cookbook_versions,omitempty" mapstructure:"cookbook_versions,omitempty"`
	// Environment default attributes JSON.
	DefaultAttributes *_struct.Struct `protobuf:"bytes,7,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty" toml:"default_attributes,omitempty" mapstructure:"default_attributes,omitempty"`
	// Environment override attributes JSON.
	OverrideAttributes   *_struct.Struct `protobuf:"bytes,8,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty" toml:"override_attributes,omitempty" mapstructure:"override_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32           `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *CreateEnvironment) Reset()         { *m = CreateEnvironment{} }
func (m *CreateEnvironment) String() string { return proto.CompactTextString(m) }
func (*CreateEnvironment) ProtoMessage()    {}
func (*CreateEnvironment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e190c707558fe428, []int{2}
}

func (m *CreateEnvironment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEnvironment.Unmarshal(m, b)
}
func (m *CreateEnvironment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEnvironment.Marshal(b, m, deterministic)
}
func (m *CreateEnvironment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEnvironment.Merge(m, src)
}
func (m *CreateEnvironment) XXX_Size() int {
	return xxx_messageInfo_CreateEnvironment.Size(m)
}
func (m *CreateEnvironment) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEnvironment.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEnvironment proto.InternalMessageInfo

func (m *CreateEnvironment) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *CreateEnvironment) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *CreateEnvironment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateEnvironment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateEnvironment) GetJsonClass() string {
	if m != nil {
		return m.JsonClass
	}
	return ""
}

func (m *CreateEnvironment) GetCookbookVersions() map[string]string {
	if m != nil {
		return m.CookbookVersions
	}
	return nil
}

func (m *CreateEnvironment) GetDefaultAttributes() *_struct.Struct {
	if m != nil {
		return m.DefaultAttributes
	}
	return nil
}

func (m *CreateEnvironment) GetOverrideAttributes() *_struct.Struct {
	if m != nil {
		return m.OverrideAttributes
	}
	return nil
}

type UpdateEnvironment struct {
	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Environment name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Environment description.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty" toml:"description,omitempty" mapstructure:"description,omitempty"`
	// Class name.
	JsonClass string `protobuf:"bytes,5,opt,name=json_class,json=jsonClass,proto3" json:"json_class,omitempty" toml:"json_class,omitempty" mapstructure:"json_class,omitempty"`
	// Environment versioned cookbooks constraints.
	CookbookVersions map[string]string `protobuf:"bytes,6,rep,name=cookbook_versions,json=cookbookVersions,proto3" json:"cookbook_versions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" toml:"cookbook_versions,omitempty" mapstructure:"cookbook_versions,omitempty"`
	// Environment default attributes JSON.
	DefaultAttributes *_struct.Struct `protobuf:"bytes,7,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty" toml:"default_attributes,omitempty" mapstructure:"default_attributes,omitempty"`
	// Environment override attributes JSON.
	OverrideAttributes   *_struct.Struct `protobuf:"bytes,8,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty" toml:"override_attributes,omitempty" mapstructure:"override_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32           `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *UpdateEnvironment) Reset()         { *m = UpdateEnvironment{} }
func (m *UpdateEnvironment) String() string { return proto.CompactTextString(m) }
func (*UpdateEnvironment) ProtoMessage()    {}
func (*UpdateEnvironment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e190c707558fe428, []int{3}
}

func (m *UpdateEnvironment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEnvironment.Unmarshal(m, b)
}
func (m *UpdateEnvironment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEnvironment.Marshal(b, m, deterministic)
}
func (m *UpdateEnvironment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEnvironment.Merge(m, src)
}
func (m *UpdateEnvironment) XXX_Size() int {
	return xxx_messageInfo_UpdateEnvironment.Size(m)
}
func (m *UpdateEnvironment) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEnvironment.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEnvironment proto.InternalMessageInfo

func (m *UpdateEnvironment) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *UpdateEnvironment) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *UpdateEnvironment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateEnvironment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateEnvironment) GetJsonClass() string {
	if m != nil {
		return m.JsonClass
	}
	return ""
}

func (m *UpdateEnvironment) GetCookbookVersions() map[string]string {
	if m != nil {
		return m.CookbookVersions
	}
	return nil
}

func (m *UpdateEnvironment) GetDefaultAttributes() *_struct.Struct {
	if m != nil {
		return m.DefaultAttributes
	}
	return nil
}

func (m *UpdateEnvironment) GetOverrideAttributes() *_struct.Struct {
	if m != nil {
		return m.OverrideAttributes
	}
	return nil
}

func init() {
	proto.RegisterType((*Environments)(nil), "chef.automate.domain.infra_proxy.request.Environments")
	proto.RegisterType((*Environment)(nil), "chef.automate.domain.infra_proxy.request.Environment")
	proto.RegisterType((*CreateEnvironment)(nil), "chef.automate.domain.infra_proxy.request.CreateEnvironment")
	proto.RegisterMapType((map[string]string)(nil), "chef.automate.domain.infra_proxy.request.CreateEnvironment.CookbookVersionsEntry")
	proto.RegisterType((*UpdateEnvironment)(nil), "chef.automate.domain.infra_proxy.request.UpdateEnvironment")
	proto.RegisterMapType((map[string]string)(nil), "chef.automate.domain.infra_proxy.request.UpdateEnvironment.CookbookVersionsEntry")
}

func init() {
	proto.RegisterFile("api/interservice/infra_proxy/request/environments.proto", fileDescriptor_e190c707558fe428)
}

var fileDescriptor_e190c707558fe428 = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x49, 0xd3, 0xc4, 0x66, 0xe2, 0xa1, 0x19, 0x2d, 0x2e, 0x55, 0x21, 0xe4, 0x94, 0xd3,
	0x0c, 0xd4, 0x83, 0x22, 0x88, 0xd8, 0x50, 0xb1, 0x47, 0x23, 0xf5, 0xe0, 0x65, 0x99, 0xdd, 0x7d,
	0xd9, 0x8e, 0xc9, 0xce, 0x5b, 0xdf, 0xbc, 0x5d, 0xcc, 0xc5, 0x2f, 0xed, 0x17, 0x90, 0x9d, 0x6c,
	0x74, 0xb1, 0x22, 0x45, 0xf0, 0xa2, 0xb7, 0xc9, 0xff, 0xe5, 0xff, 0xe3, 0xcd, 0xfc, 0x60, 0xc5,
	0x53, 0x53, 0x5a, 0x6d, 0x1d, 0x03, 0x79, 0xa0, 0xda, 0xa6, 0xa0, 0xad, 0x5b, 0x91, 0x89, 0x4b,
	0xc2, 0xcf, 0x5b, 0x4d, 0xf0, 0xa9, 0x02, 0xcf, 0x1a, 0x5c, 0x6d, 0x09, 0x5d, 0x01, 0x8e, 0xbd,
	0x2a, 0x09, 0x19, 0xe5, 0x3c, 0xbd, 0x86, 0x95, 0x32, 0x15, 0x63, 0x61, 0x18, 0x54, 0x86, 0x85,
	0xb1, 0x4e, 0x75, 0xca, 0xaa, 0x2d, 0x9f, 0x3e, 0xca, 0x11, 0xf3, 0x0d, 0xe8, 0xd0, 0x4b, 0xaa,
	0x95, 0xf6, 0x4c, 0x55, 0xca, 0x3b, 0xce, 0xec, 0x5c, 0xdc, 0xbd, 0xe8, 0xd0, 0xe5, 0x89, 0x18,
	0x22, 0xe5, 0xb1, 0xcd, 0xa2, 0xde, 0xb4, 0x37, 0x1f, 0x2d, 0x07, 0x48, 0xf9, 0x65, 0x26, 0x1f,
	0x8a, 0x51, 0xb3, 0x20, 0x50, 0x33, 0x39, 0x08, 0x93, 0xa3, 0x5d, 0x70, 0x99, 0xcd, 0xae, 0xc4,
	0xb8, 0xc3, 0xf8, 0x13, 0x84, 0x94, 0xe2, 0xd0, 0x99, 0x02, 0xa2, 0x7e, 0xc8, 0xc3, 0x79, 0xf6,
	0xb5, 0x2f, 0x26, 0x0b, 0x02, 0xc3, 0xf0, 0x17, 0xe8, 0x72, 0x2a, 0xc6, 0x19, 0xf8, 0x94, 0x6c,
	0xc9, 0x16, 0x5d, 0x74, 0x18, 0x46, 0xdd, 0x48, 0x3e, 0x16, 0xe2, 0xa3, 0x47, 0x17, 0xa7, 0x1b,
	0xe3, 0x7d, 0x34, 0x08, 0x7f, 0x18, 0x35, 0xc9, 0xa2, 0x09, 0xe4, 0x17, 0x31, 0x49, 0x11, 0xd7,
	0x09, 0xe2, 0x3a, 0xae, 0x81, 0xbc, 0x45, 0xe7, 0xa3, 0xe1, 0xb4, 0x3f, 0x1f, 0x9f, 0xbd, 0x55,
	0xb7, 0xb5, 0xa3, 0x6e, 0x5c, 0x50, 0x2d, 0x5a, 0xe8, 0xfb, 0x96, 0x79, 0xe1, 0x98, 0xb6, 0xcb,
	0xe3, 0xf4, 0xa7, 0x58, 0xbe, 0x16, 0x32, 0x83, 0x95, 0xa9, 0x36, 0x1c, 0x1b, 0x66, 0xb2, 0x49,
	0xc5, 0xe0, 0xa3, 0x3b, 0xd3, 0xde, 0x7c, 0x7c, 0xf6, 0x40, 0xed, 0xa4, 0xab, 0xbd, 0x74, 0xf5,
	0x2e, 0x48, 0x5f, 0x4e, 0xda, 0xca, 0xab, 0xef, 0x0d, 0xf9, 0x46, 0xdc, 0xc3, 0x1a, 0x88, 0x6c,
	0x06, 0x5d, 0xd0, 0xd1, 0xef, 0x41, 0x72, 0xdf, 0xf9, 0x41, 0x3a, 0x5d, 0x88, 0x93, 0x5f, 0x2e,
	0x2f, 0x8f, 0x45, 0x7f, 0x0d, 0xdb, 0x56, 0x58, 0x73, 0x94, 0xf7, 0xc5, 0xa0, 0x36, 0x9b, 0x0a,
	0x5a, 0x55, 0xbb, 0x1f, 0xcf, 0x0f, 0x9e, 0xf5, 0x82, 0xf5, 0xab, 0x32, 0xfb, 0xb7, 0xad, 0xdf,
	0xb8, 0xe0, 0x7f, 0x6e, 0xfd, 0xfc, 0xe5, 0x87, 0x17, 0xb9, 0xe5, 0xeb, 0x2a, 0x51, 0x29, 0x16,
	0xba, 0x79, 0x47, 0xbd, 0x7f, 0x47, 0x7d, 0x9b, 0x4f, 0x64, 0x32, 0x0c, 0xab, 0x3e, 0xf9, 0x16,
	0x00, 0x00, 0xff, 0xff, 0x57, 0x9a, 0x09, 0x67, 0x51, 0x05, 0x00, 0x00,
}
