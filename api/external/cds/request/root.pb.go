// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/cds/request/root.proto

package request

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

type ContentItems struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContentItems) Reset()         { *m = ContentItems{} }
func (m *ContentItems) String() string { return proto.CompactTextString(m) }
func (*ContentItems) ProtoMessage()    {}
func (*ContentItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4605f190c3b603e, []int{0}
}

func (m *ContentItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentItems.Unmarshal(m, b)
}
func (m *ContentItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentItems.Marshal(b, m, deterministic)
}
func (m *ContentItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentItems.Merge(m, src)
}
func (m *ContentItems) XXX_Size() int {
	return xxx_messageInfo_ContentItems.Size(m)
}
func (m *ContentItems) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentItems.DiscardUnknown(m)
}

var xxx_messageInfo_ContentItems proto.InternalMessageInfo

type InstallContentItem struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RequestUser          string   `protobuf:"bytes,2,opt,name=request_user,json=requestUser,proto3" json:"request_user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallContentItem) Reset()         { *m = InstallContentItem{} }
func (m *InstallContentItem) String() string { return proto.CompactTextString(m) }
func (*InstallContentItem) ProtoMessage()    {}
func (*InstallContentItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4605f190c3b603e, []int{1}
}

func (m *InstallContentItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallContentItem.Unmarshal(m, b)
}
func (m *InstallContentItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallContentItem.Marshal(b, m, deterministic)
}
func (m *InstallContentItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallContentItem.Merge(m, src)
}
func (m *InstallContentItem) XXX_Size() int {
	return xxx_messageInfo_InstallContentItem.Size(m)
}
func (m *InstallContentItem) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallContentItem.DiscardUnknown(m)
}

var xxx_messageInfo_InstallContentItem proto.InternalMessageInfo

func (m *InstallContentItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InstallContentItem) GetRequestUser() string {
	if m != nil {
		return m.RequestUser
	}
	return ""
}

type DownloadContentItem struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DownloadContentItem) Reset()         { *m = DownloadContentItem{} }
func (m *DownloadContentItem) String() string { return proto.CompactTextString(m) }
func (*DownloadContentItem) ProtoMessage()    {}
func (*DownloadContentItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4605f190c3b603e, []int{2}
}

func (m *DownloadContentItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownloadContentItem.Unmarshal(m, b)
}
func (m *DownloadContentItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownloadContentItem.Marshal(b, m, deterministic)
}
func (m *DownloadContentItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownloadContentItem.Merge(m, src)
}
func (m *DownloadContentItem) XXX_Size() int {
	return xxx_messageInfo_DownloadContentItem.Size(m)
}
func (m *DownloadContentItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DownloadContentItem.DiscardUnknown(m)
}

var xxx_messageInfo_DownloadContentItem proto.InternalMessageInfo

func (m *DownloadContentItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*ContentItems)(nil), "chef.automate.api.cds.request.ContentItems")
	proto.RegisterType((*InstallContentItem)(nil), "chef.automate.api.cds.request.InstallContentItem")
	proto.RegisterType((*DownloadContentItem)(nil), "chef.automate.api.cds.request.DownloadContentItem")
}

func init() {
	proto.RegisterFile("api/external/cds/request/root.proto", fileDescriptor_a4605f190c3b603e)
}

var fileDescriptor_a4605f190c3b603e = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x31, 0x6b, 0xc3, 0x30,
	0x10, 0x46, 0xb1, 0x87, 0x42, 0x55, 0xe3, 0x41, 0x5d, 0xbc, 0x14, 0x5a, 0x97, 0x42, 0x27, 0x89,
	0xe2, 0x7f, 0xd0, 0x04, 0x82, 0xd7, 0x40, 0x96, 0x2c, 0x41, 0x96, 0x2e, 0xb1, 0x40, 0xd6, 0x39,
	0xd2, 0x89, 0xe4, 0xe7, 0x87, 0x18, 0x07, 0xb2, 0x24, 0xeb, 0x77, 0xc7, 0xe3, 0x3d, 0xf6, 0xad,
	0x46, 0x2b, 0xe1, 0x4c, 0x10, 0xbc, 0x72, 0x52, 0x9b, 0x28, 0x03, 0x1c, 0x13, 0x44, 0x92, 0x01,
	0x91, 0xc4, 0x18, 0x90, 0x90, 0x7f, 0xe8, 0x1e, 0xf6, 0x42, 0x25, 0xc2, 0x41, 0x11, 0x08, 0x35,
	0x5a, 0xa1, 0x4d, 0x14, 0xf3, 0x67, 0x5d, 0xb2, 0x62, 0x81, 0x9e, 0xc0, 0x53, 0x4b, 0x30, 0xc4,
	0x7a, 0xc5, 0x78, 0xeb, 0x23, 0x29, 0xe7, 0xee, 0x66, 0x5e, 0xb2, 0xdc, 0x9a, 0x2a, 0xfb, 0xcc,
	0x7e, 0x5f, 0xd7, 0xb9, 0x35, 0xfc, 0x8b, 0x15, 0x33, 0x60, 0x97, 0x22, 0x84, 0x2a, 0x9f, 0x2e,
	0x6f, 0xf3, 0xb6, 0x89, 0x10, 0xea, 0x1f, 0xf6, 0xbe, 0xc4, 0x93, 0x77, 0xa8, 0xcc, 0x13, 0xd2,
	0x7f, 0xb3, 0xfd, 0x3b, 0x58, 0xea, 0x53, 0x27, 0x34, 0x0e, 0xf2, 0xea, 0x2a, 0x6f, 0xae, 0xf2,
	0x51, 0x5e, 0xf7, 0x32, 0xa5, 0x35, 0x97, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xc1, 0x2f, 0x15,
	0x01, 0x01, 0x00, 0x00,
}
