// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/common/version/version.proto

package version

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

type VersionInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Sha                  string   `protobuf:"bytes,3,opt,name=sha,proto3" json:"sha,omitempty"`
	Built                string   `protobuf:"bytes,4,opt,name=built,proto3" json:"built,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionInfo) Reset()         { *m = VersionInfo{} }
func (m *VersionInfo) String() string { return proto.CompactTextString(m) }
func (*VersionInfo) ProtoMessage()    {}
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d8071f8bbbe2d80, []int{0}
}

func (m *VersionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionInfo.Unmarshal(m, b)
}
func (m *VersionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionInfo.Marshal(b, m, deterministic)
}
func (m *VersionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionInfo.Merge(m, src)
}
func (m *VersionInfo) XXX_Size() int {
	return xxx_messageInfo_VersionInfo.Size(m)
}
func (m *VersionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VersionInfo proto.InternalMessageInfo

func (m *VersionInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VersionInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionInfo) GetSha() string {
	if m != nil {
		return m.Sha
	}
	return ""
}

func (m *VersionInfo) GetBuilt() string {
	if m != nil {
		return m.Built
	}
	return ""
}

type VersionInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionInfoRequest) Reset()         { *m = VersionInfoRequest{} }
func (m *VersionInfoRequest) String() string { return proto.CompactTextString(m) }
func (*VersionInfoRequest) ProtoMessage()    {}
func (*VersionInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d8071f8bbbe2d80, []int{1}
}

func (m *VersionInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionInfoRequest.Unmarshal(m, b)
}
func (m *VersionInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionInfoRequest.Marshal(b, m, deterministic)
}
func (m *VersionInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionInfoRequest.Merge(m, src)
}
func (m *VersionInfoRequest) XXX_Size() int {
	return xxx_messageInfo_VersionInfoRequest.Size(m)
}
func (m *VersionInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionInfoRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VersionInfo)(nil), "chef.automate.api.common.version.VersionInfo")
	proto.RegisterType((*VersionInfoRequest)(nil), "chef.automate.api.common.version.VersionInfoRequest")
}

func init() {
	proto.RegisterFile("api/external/common/version/version.proto", fileDescriptor_8d8071f8bbbe2d80)
}

var fileDescriptor_8d8071f8bbbe2d80 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x4f, 0xad, 0x28, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3,
	0x2f, 0x4b, 0x2d, 0x2a, 0xce, 0x44, 0xd0, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x0a, 0xc9,
	0x19, 0xa9, 0x69, 0x7a, 0x89, 0xa5, 0x25, 0xf9, 0xb9, 0x89, 0x25, 0xa9, 0x7a, 0x89, 0x05, 0x99,
	0x7a, 0x10, 0xf5, 0x7a, 0x50, 0x75, 0x4a, 0xc9, 0x5c, 0xdc, 0x61, 0x10, 0xa6, 0x67, 0x5e, 0x5a,
	0xbe, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x98, 0x2d, 0x24, 0xc1, 0xc5, 0x0e, 0x55, 0x2d, 0xc1, 0x04, 0x16, 0x86, 0x71, 0x85, 0x04, 0xb8,
	0x98, 0x8b, 0x33, 0x12, 0x25, 0x98, 0xc1, 0xa2, 0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b, 0x52, 0x69,
	0x66, 0x4e, 0x89, 0x04, 0x0b, 0x58, 0x0c, 0xc2, 0x51, 0x12, 0xe1, 0x12, 0x42, 0xb2, 0x24, 0x28,
	0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0xc4, 0xc9, 0x2c, 0xca, 0x24, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x09,
	0xe4, 0x26, 0x7d, 0x90, 0x4b, 0xf5, 0x61, 0x2e, 0xd5, 0xc7, 0xe3, 0xc5, 0x24, 0x36, 0xb0, 0xdf,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x5e, 0x85, 0x09, 0x08, 0x01, 0x00, 0x00,
}
