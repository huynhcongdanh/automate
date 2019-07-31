// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lib/license/license.proto

package license

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type License struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version              string               `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Type                 string               `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Generator            string               `protobuf:"bytes,20,opt,name=generator,proto3" json:"generator,omitempty"`
	KeySha256            string               `protobuf:"bytes,6,opt,name=key_sha256,json=keySha256,proto3" json:"key_sha256,omitempty"`
	GenerationDate       *timestamp.Timestamp `protobuf:"bytes,21,opt,name=generation_date,json=generationDate,proto3" json:"generation_date,omitempty"`
	Customer             string               `protobuf:"bytes,24,opt,name=customer,proto3" json:"customer,omitempty"`
	CustomerId           string               `protobuf:"bytes,4,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	CustomerIdVersion    string               `protobuf:"bytes,25,opt,name=customer_id_version,json=customerIdVersion,proto3" json:"customer_id_version,omitempty"`
	Entitlements         []*Entitlement       `protobuf:"bytes,5,rep,name=entitlements,proto3" json:"entitlements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *License) Reset()         { *m = License{} }
func (m *License) String() string { return proto.CompactTextString(m) }
func (*License) ProtoMessage()    {}
func (*License) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d88e617a98a96bd, []int{0}
}

func (m *License) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_License.Unmarshal(m, b)
}
func (m *License) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_License.Marshal(b, m, deterministic)
}
func (m *License) XXX_Merge(src proto.Message) {
	xxx_messageInfo_License.Merge(m, src)
}
func (m *License) XXX_Size() int {
	return xxx_messageInfo_License.Size(m)
}
func (m *License) XXX_DiscardUnknown() {
	xxx_messageInfo_License.DiscardUnknown(m)
}

var xxx_messageInfo_License proto.InternalMessageInfo

func (m *License) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *License) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *License) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *License) GetGenerator() string {
	if m != nil {
		return m.Generator
	}
	return ""
}

func (m *License) GetKeySha256() string {
	if m != nil {
		return m.KeySha256
	}
	return ""
}

func (m *License) GetGenerationDate() *timestamp.Timestamp {
	if m != nil {
		return m.GenerationDate
	}
	return nil
}

func (m *License) GetCustomer() string {
	if m != nil {
		return m.Customer
	}
	return ""
}

func (m *License) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *License) GetCustomerIdVersion() string {
	if m != nil {
		return m.CustomerIdVersion
	}
	return ""
}

func (m *License) GetEntitlements() []*Entitlement {
	if m != nil {
		return m.Entitlements
	}
	return nil
}

type Entitlement struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Measure              string               `protobuf:"bytes,2,opt,name=measure,proto3" json:"measure,omitempty"`
	Limit                int64                `protobuf:"zigzag64,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Start                *timestamp.Timestamp `protobuf:"bytes,22,opt,name=start,proto3" json:"start,omitempty"`
	End                  *timestamp.Timestamp `protobuf:"bytes,23,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Entitlement) Reset()         { *m = Entitlement{} }
func (m *Entitlement) String() string { return proto.CompactTextString(m) }
func (*Entitlement) ProtoMessage()    {}
func (*Entitlement) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d88e617a98a96bd, []int{1}
}

func (m *Entitlement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entitlement.Unmarshal(m, b)
}
func (m *Entitlement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entitlement.Marshal(b, m, deterministic)
}
func (m *Entitlement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entitlement.Merge(m, src)
}
func (m *Entitlement) XXX_Size() int {
	return xxx_messageInfo_Entitlement.Size(m)
}
func (m *Entitlement) XXX_DiscardUnknown() {
	xxx_messageInfo_Entitlement.DiscardUnknown(m)
}

var xxx_messageInfo_Entitlement proto.InternalMessageInfo

func (m *Entitlement) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entitlement) GetMeasure() string {
	if m != nil {
		return m.Measure
	}
	return ""
}

func (m *Entitlement) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Entitlement) GetStart() *timestamp.Timestamp {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Entitlement) GetEnd() *timestamp.Timestamp {
	if m != nil {
		return m.End
	}
	return nil
}

func init() {
	proto.RegisterType((*License)(nil), "license.License")
	proto.RegisterType((*Entitlement)(nil), "license.Entitlement")
}

func init() { proto.RegisterFile("lib/license/license.proto", fileDescriptor_1d88e617a98a96bd) }

var fileDescriptor_1d88e617a98a96bd = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x95, 0x74, 0xbb, 0x65, 0x27, 0x68, 0x11, 0xa6, 0x80, 0xb7, 0x02, 0x6d, 0xb5, 0x42,
	0xa8, 0x07, 0x94, 0xa0, 0x22, 0x10, 0x67, 0xfe, 0x1c, 0x90, 0x38, 0x05, 0xc4, 0x81, 0x4b, 0xe4,
	0x24, 0xd3, 0xd4, 0x6a, 0x6c, 0x57, 0xf1, 0x04, 0xa9, 0x1f, 0x8b, 0x6f, 0xc6, 0x47, 0x40, 0x71,
	0xe2, 0xa6, 0x7b, 0xea, 0x29, 0xf3, 0xde, 0xfb, 0x25, 0xd1, 0x3c, 0x1b, 0x6e, 0x6a, 0x99, 0x27,
	0xb5, 0x2c, 0x50, 0x5b, 0xf4, 0xcf, 0x78, 0xdf, 0x18, 0x32, 0x6c, 0x36, 0xc8, 0xc5, 0x6d, 0x65,
	0x4c, 0x55, 0x63, 0xe2, 0xec, 0xbc, 0xdd, 0x24, 0x24, 0x15, 0x5a, 0x12, 0x6a, 0xdf, 0x93, 0x77,
	0xff, 0x42, 0x98, 0x7d, 0xef, 0x61, 0x76, 0x0d, 0xa1, 0x2c, 0x79, 0xb0, 0x0c, 0x56, 0x57, 0x69,
	0x28, 0x4b, 0xc6, 0x61, 0xf6, 0x07, 0x1b, 0x2b, 0x8d, 0xe6, 0xa1, 0x33, 0xbd, 0x64, 0x0c, 0x2e,
	0xe8, 0xb0, 0x47, 0x3e, 0x71, 0xb6, 0x9b, 0xd9, 0x0b, 0xb8, 0xaa, 0x50, 0x63, 0x23, 0xc8, 0x34,
	0x7c, 0xee, 0x82, 0xd1, 0x60, 0x2f, 0x01, 0x76, 0x78, 0xc8, 0xec, 0x56, 0xac, 0xdf, 0x7f, 0xe0,
	0x97, 0x7d, 0xbc, 0xc3, 0xc3, 0x0f, 0x67, 0xb0, 0xcf, 0xf0, 0x68, 0x60, 0xa5, 0xd1, 0x59, 0x29,
	0x08, 0xf9, 0xd3, 0x65, 0xb0, 0x8a, 0xd6, 0x8b, 0xb8, 0xdf, 0x20, 0xf6, 0x1b, 0xc4, 0x3f, 0xfd,
	0x06, 0xe9, 0xf5, 0xf8, 0xca, 0x17, 0x41, 0xc8, 0x16, 0xf0, 0xa0, 0x68, 0x2d, 0x19, 0x85, 0x0d,
	0xe7, 0xee, 0x0f, 0x47, 0xcd, 0x6e, 0x21, 0xf2, 0x73, 0x26, 0x4b, 0x7e, 0xe1, 0x62, 0xf0, 0xd6,
	0xb7, 0x92, 0xc5, 0xf0, 0xe4, 0x04, 0xc8, 0xfc, 0xe2, 0x37, 0x0e, 0x7c, 0x3c, 0x82, 0xbf, 0x86,
	0x0a, 0x3e, 0xc2, 0x43, 0xd4, 0x24, 0xa9, 0x46, 0x85, 0x9a, 0x2c, 0x9f, 0x2e, 0x27, 0xab, 0x68,
	0x3d, 0x8f, 0xfd, 0x41, 0x7c, 0x1d, 0xc3, 0xf4, 0x1e, 0x79, 0xf7, 0x37, 0x80, 0xe8, 0x24, 0xed,
	0xca, 0xd4, 0x42, 0xe1, 0x50, 0xbc, 0x9b, 0xbb, 0xea, 0x15, 0x0a, 0xdb, 0x36, 0xe8, 0xab, 0x1f,
	0x24, 0x9b, 0xc3, 0xb4, 0x96, 0x4a, 0x92, 0xeb, 0x9e, 0xa5, 0xbd, 0x60, 0x6f, 0x61, 0x6a, 0x49,
	0x34, 0xc4, 0x9f, 0x9d, 0x6d, 0xad, 0x07, 0xd9, 0x1b, 0x98, 0xa0, 0x2e, 0xf9, 0xf3, 0xb3, 0x7c,
	0x87, 0x7d, 0x7a, 0xfd, 0xfb, 0x55, 0x25, 0x69, 0xdb, 0xe6, 0x71, 0x61, 0x54, 0x52, 0x6c, 0x71,
	0x93, 0x88, 0x96, 0x8c, 0x12, 0xd4, 0x5d, 0xbd, 0xe3, 0x35, 0xcc, 0x2f, 0xdd, 0x07, 0xde, 0xfd,
	0x0f, 0x00, 0x00, 0xff, 0xff, 0x64, 0x06, 0x47, 0xd2, 0x9c, 0x02, 0x00, 0x00,
}
