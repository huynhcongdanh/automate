// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/common/rules.proto

package common

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

type RuleType int32

const (
	RuleType_RULE_TYPE_UNSET RuleType = 0
	RuleType_NODE            RuleType = 1
	RuleType_EVENT           RuleType = 2
)

var RuleType_name = map[int32]string{
	0: "RULE_TYPE_UNSET",
	1: "NODE",
	2: "EVENT",
}

var RuleType_value = map[string]int32{
	"RULE_TYPE_UNSET": 0,
	"NODE":            1,
	"EVENT":           2,
}

func (x RuleType) String() string {
	return proto.EnumName(RuleType_name, int32(x))
}

func (RuleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_313d0dfa57edd469, []int{0}
}

type ConditionAttribute int32

const (
	ConditionAttribute_CONDITION_ATTRIBUTE_UNSET ConditionAttribute = 0
	ConditionAttribute_CHEF_SERVER               ConditionAttribute = 1
	ConditionAttribute_CHEF_ORGANIZATION         ConditionAttribute = 2
	ConditionAttribute_ENVIRONMENT               ConditionAttribute = 3
	ConditionAttribute_CHEF_ROLE                 ConditionAttribute = 4
	ConditionAttribute_CHEF_TAG                  ConditionAttribute = 5
	ConditionAttribute_CHEF_POLICY_GROUP         ConditionAttribute = 6
	ConditionAttribute_CHEF_POLICY_NAME          ConditionAttribute = 7
)

var ConditionAttribute_name = map[int32]string{
	0: "CONDITION_ATTRIBUTE_UNSET",
	1: "CHEF_SERVER",
	2: "CHEF_ORGANIZATION",
	3: "ENVIRONMENT",
	4: "CHEF_ROLE",
	5: "CHEF_TAG",
	6: "CHEF_POLICY_GROUP",
	7: "CHEF_POLICY_NAME",
}

var ConditionAttribute_value = map[string]int32{
	"CONDITION_ATTRIBUTE_UNSET": 0,
	"CHEF_SERVER":               1,
	"CHEF_ORGANIZATION":         2,
	"ENVIRONMENT":               3,
	"CHEF_ROLE":                 4,
	"CHEF_TAG":                  5,
	"CHEF_POLICY_GROUP":         6,
	"CHEF_POLICY_NAME":          7,
}

func (x ConditionAttribute) String() string {
	return proto.EnumName(ConditionAttribute_name, int32(x))
}

func (ConditionAttribute) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_313d0dfa57edd469, []int{1}
}

type ConditionOperator int32

const (
	ConditionOperator_CONDITION_OPERATOR_UNSET ConditionOperator = 0
	ConditionOperator_MEMBER_OF                ConditionOperator = 1
	ConditionOperator_EQUALS                   ConditionOperator = 2
)

var ConditionOperator_name = map[int32]string{
	0: "CONDITION_OPERATOR_UNSET",
	1: "MEMBER_OF",
	2: "EQUALS",
}

var ConditionOperator_value = map[string]int32{
	"CONDITION_OPERATOR_UNSET": 0,
	"MEMBER_OF":                1,
	"EQUALS":                   2,
}

func (x ConditionOperator) String() string {
	return proto.EnumName(ConditionOperator_name, int32(x))
}

func (ConditionOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_313d0dfa57edd469, []int{2}
}

type RuleStatus int32

const (
	RuleStatus_RULE_STATUS_UNSET RuleStatus = 0
	RuleStatus_STAGED            RuleStatus = 1
	RuleStatus_APPLIED           RuleStatus = 2
)

var RuleStatus_name = map[int32]string{
	0: "RULE_STATUS_UNSET",
	1: "STAGED",
	2: "APPLIED",
}

var RuleStatus_value = map[string]int32{
	"RULE_STATUS_UNSET": 0,
	"STAGED":            1,
	"APPLIED":           2,
}

func (x RuleStatus) String() string {
	return proto.EnumName(RuleStatus_name, int32(x))
}

func (RuleStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_313d0dfa57edd469, []int{3}
}

type ProjectRulesStatus int32

const (
	ProjectRulesStatus_PROJECT_RULES_STATUS_UNSET ProjectRulesStatus = 0
	ProjectRulesStatus_RULES_APPLIED              ProjectRulesStatus = 1
	ProjectRulesStatus_EDITS_PENDING              ProjectRulesStatus = 2
	ProjectRulesStatus_NO_RULES                   ProjectRulesStatus = 3
)

var ProjectRulesStatus_name = map[int32]string{
	0: "PROJECT_RULES_STATUS_UNSET",
	1: "RULES_APPLIED",
	2: "EDITS_PENDING",
	3: "NO_RULES",
}

var ProjectRulesStatus_value = map[string]int32{
	"PROJECT_RULES_STATUS_UNSET": 0,
	"RULES_APPLIED":              1,
	"EDITS_PENDING":              2,
	"NO_RULES":                   3,
}

func (x ProjectRulesStatus) String() string {
	return proto.EnumName(ProjectRulesStatus_name, int32(x))
}

func (ProjectRulesStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_313d0dfa57edd469, []int{4}
}

type Rule struct {
	// Unique ID. Cannot be changed.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique ID of the project this rule belongs to. Cannot be changed.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Name for the project rule.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Whether the rule applies to ingested `NODE` or `EVENT resources.
	// Cannot be changed.
	Type RuleType `protobuf:"varint,4,opt,name=type,proto3,enum=chef.automate.api.iam.v2.RuleType" json:"type,omitempty"`
	// Conditions that ingested resources must match to belong to the project.
	// Will contain one or more.
	Conditions []*Condition `protobuf:"bytes,5,rep,name=conditions,proto3" json:"conditions,omitempty"`
	// Whether the rule is `STAGED` (not in effect) or `APPLIED` (in effect).
	Status               RuleStatus `protobuf:"varint,6,opt,name=status,proto3,enum=chef.automate.api.iam.v2.RuleStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_313d0dfa57edd469, []int{0}
}

func (m *Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rule.Unmarshal(m, b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
}
func (m *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(m, src)
}
func (m *Rule) XXX_Size() int {
	return xxx_messageInfo_Rule.Size(m)
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Rule) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Rule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rule) GetType() RuleType {
	if m != nil {
		return m.Type
	}
	return RuleType_RULE_TYPE_UNSET
}

func (m *Rule) GetConditions() []*Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *Rule) GetStatus() RuleStatus {
	if m != nil {
		return m.Status
	}
	return RuleStatus_RULE_STATUS_UNSET
}

type Condition struct {
	// Represents a property of an ingested resource. Depends on the rule type.
	Attribute ConditionAttribute `protobuf:"varint,1,opt,name=attribute,proto3,enum=chef.automate.api.iam.v2.ConditionAttribute" json:"attribute,omitempty"`
	// The value(s) of the attribute that an ingested resource must match.
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	// Whether the attribute matches a single value (`EQUALS`) or
	// matches at least one of a set of values (`MEMBER_OF`).
	Operator             ConditionOperator `protobuf:"varint,3,opt,name=operator,proto3,enum=chef.automate.api.iam.v2.ConditionOperator" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Condition) Reset()         { *m = Condition{} }
func (m *Condition) String() string { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()    {}
func (*Condition) Descriptor() ([]byte, []int) {
	return fileDescriptor_313d0dfa57edd469, []int{1}
}

func (m *Condition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Condition.Unmarshal(m, b)
}
func (m *Condition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Condition.Marshal(b, m, deterministic)
}
func (m *Condition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Condition.Merge(m, src)
}
func (m *Condition) XXX_Size() int {
	return xxx_messageInfo_Condition.Size(m)
}
func (m *Condition) XXX_DiscardUnknown() {
	xxx_messageInfo_Condition.DiscardUnknown(m)
}

var xxx_messageInfo_Condition proto.InternalMessageInfo

func (m *Condition) GetAttribute() ConditionAttribute {
	if m != nil {
		return m.Attribute
	}
	return ConditionAttribute_CONDITION_ATTRIBUTE_UNSET
}

func (m *Condition) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Condition) GetOperator() ConditionOperator {
	if m != nil {
		return m.Operator
	}
	return ConditionOperator_CONDITION_OPERATOR_UNSET
}

func init() {
	proto.RegisterEnum("chef.automate.api.iam.v2.RuleType", RuleType_name, RuleType_value)
	proto.RegisterEnum("chef.automate.api.iam.v2.ConditionAttribute", ConditionAttribute_name, ConditionAttribute_value)
	proto.RegisterEnum("chef.automate.api.iam.v2.ConditionOperator", ConditionOperator_name, ConditionOperator_value)
	proto.RegisterEnum("chef.automate.api.iam.v2.RuleStatus", RuleStatus_name, RuleStatus_value)
	proto.RegisterEnum("chef.automate.api.iam.v2.ProjectRulesStatus", ProjectRulesStatus_name, ProjectRulesStatus_value)
	proto.RegisterType((*Rule)(nil), "chef.automate.api.iam.v2.Rule")
	proto.RegisterType((*Condition)(nil), "chef.automate.api.iam.v2.Condition")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/common/rules.proto", fileDescriptor_313d0dfa57edd469)
}

var fileDescriptor_313d0dfa57edd469 = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5b, 0x6f, 0xd3, 0x30,
	0x14, 0x26, 0x69, 0xd7, 0x35, 0x67, 0xac, 0xf3, 0xcc, 0x45, 0x01, 0x31, 0x54, 0x0d, 0x1e, 0xaa,
	0x02, 0x89, 0x54, 0x10, 0x12, 0xd2, 0x5e, 0xb2, 0xd6, 0xeb, 0x32, 0xb5, 0x71, 0x70, 0xdc, 0x49,
	0xdb, 0x4b, 0x94, 0xb5, 0x61, 0x0b, 0x5a, 0xea, 0xa8, 0x71, 0x86, 0xf6, 0xce, 0x6f, 0x42, 0xfc,
	0x3c, 0x14, 0xaf, 0x97, 0x89, 0x69, 0x68, 0xbc, 0xc5, 0x5f, 0xce, 0x77, 0xd1, 0x27, 0xfb, 0xc0,
	0x97, 0xb1, 0x48, 0x33, 0x31, 0x8d, 0xa7, 0x32, 0xb7, 0xa3, 0x42, 0x8a, 0x34, 0x92, 0xf1, 0x87,
	0xf3, 0x48, 0xc6, 0x3f, 0xa2, 0x6b, 0x3b, 0xca, 0x12, 0x3b, 0x89, 0x52, 0xfb, 0xaa, 0x63, 0x8f,
	0x45, 0x9a, 0x8a, 0xa9, 0x3d, 0x2b, 0x2e, 0xe3, 0xdc, 0xca, 0x66, 0x42, 0x0a, 0x6c, 0x8e, 0x2f,
	0xe2, 0x6f, 0xd6, 0x82, 0x64, 0x45, 0x59, 0x62, 0x25, 0x51, 0x6a, 0x5d, 0x75, 0x76, 0x7f, 0xea,
	0x50, 0x65, 0xc5, 0x65, 0x8c, 0x1b, 0xa0, 0x27, 0x13, 0x53, 0x6b, 0x6a, 0x2d, 0x83, 0xe9, 0xc9,
	0x04, 0xef, 0x00, 0x64, 0x33, 0xf1, 0x3d, 0x1e, 0xcb, 0x30, 0x99, 0x98, 0xba, 0xc2, 0x8d, 0x39,
	0xe2, 0x4e, 0x30, 0x86, 0xea, 0x34, 0x4a, 0x63, 0xb3, 0xa2, 0x7e, 0xa8, 0x6f, 0xfc, 0x19, 0xaa,
	0xf2, 0x3a, 0x8b, 0xcd, 0x6a, 0x53, 0x6b, 0x35, 0x3a, 0xbb, 0xd6, 0x7d, 0xa6, 0x56, 0x69, 0xc8,
	0xaf, 0xb3, 0x98, 0xa9, 0x79, 0xdc, 0x05, 0x18, 0x8b, 0xe9, 0x24, 0x91, 0x89, 0x98, 0xe6, 0xe6,
	0x5a, 0xb3, 0xd2, 0xda, 0xe8, 0xbc, 0xb9, 0x9f, 0xdd, 0x5d, 0xcc, 0xb2, 0x5b, 0x34, 0xbc, 0x07,
	0xb5, 0x5c, 0x46, 0xb2, 0xc8, 0xcd, 0x9a, 0xb2, 0x7f, 0xfb, 0x6f, 0xfb, 0x40, 0xcd, 0xb2, 0x39,
	0x67, 0xf7, 0xb7, 0x06, 0xc6, 0x52, 0x17, 0x1f, 0x81, 0x11, 0x49, 0x39, 0x4b, 0xce, 0x0a, 0x19,
	0xab, 0x4a, 0x1a, 0x9d, 0xf7, 0x0f, 0xc8, 0xe3, 0x2c, 0x38, 0x6c, 0x45, 0xc7, 0xcf, 0xa1, 0x76,
	0x15, 0x5d, 0x16, 0x71, 0x6e, 0xea, 0xcd, 0x4a, 0xcb, 0x60, 0xf3, 0x13, 0xee, 0x43, 0x5d, 0x64,
	0xf1, 0x2c, 0x92, 0x62, 0xa6, 0x4a, 0x6c, 0x74, 0xde, 0x3d, 0xc0, 0x82, 0xce, 0x29, 0x6c, 0x49,
	0x6e, 0x7f, 0x82, 0xfa, 0xa2, 0x4f, 0xfc, 0x04, 0xb6, 0xd8, 0x68, 0x40, 0x42, 0x7e, 0xe2, 0x93,
	0x70, 0xe4, 0x05, 0x84, 0xa3, 0x47, 0xb8, 0x0e, 0x55, 0x8f, 0xf6, 0x08, 0xd2, 0xb0, 0x01, 0x6b,
	0xe4, 0x98, 0x78, 0x1c, 0xe9, 0xed, 0x5f, 0x1a, 0xe0, 0xbb, 0xc1, 0xf1, 0x0e, 0xbc, 0xe8, 0x52,
	0xaf, 0xe7, 0x72, 0x97, 0x7a, 0xa1, 0xc3, 0x39, 0x73, 0xf7, 0x47, 0x7c, 0x25, 0xb5, 0x05, 0x1b,
	0xdd, 0x43, 0x72, 0x10, 0x06, 0x84, 0x1d, 0x13, 0x86, 0x34, 0xfc, 0x0c, 0xb6, 0x15, 0x40, 0x59,
	0xdf, 0xf1, 0xdc, 0x53, 0xa7, 0xe4, 0x21, 0xbd, 0x9c, 0x23, 0xde, 0xb1, 0xcb, 0xa8, 0x37, 0x2c,
	0xed, 0x2a, 0x78, 0x13, 0x0c, 0x35, 0xc7, 0xe8, 0x80, 0xa0, 0x2a, 0x7e, 0x0c, 0x75, 0x75, 0xe4,
	0x4e, 0x1f, 0xad, 0x2d, 0x45, 0x7c, 0x3a, 0x70, 0xbb, 0x27, 0x61, 0x9f, 0xd1, 0x91, 0x8f, 0x6a,
	0xf8, 0x29, 0xa0, 0xdb, 0xb0, 0xe7, 0x0c, 0x09, 0x5a, 0x6f, 0x0f, 0x60, 0xfb, 0x4e, 0x1b, 0xf8,
	0x15, 0x98, 0xab, 0xd8, 0xd4, 0x27, 0xcc, 0xe1, 0x94, 0x2d, 0x53, 0x6f, 0x82, 0x31, 0x24, 0xc3,
	0x7d, 0xc2, 0x42, 0x7a, 0x80, 0x34, 0x0c, 0x50, 0x23, 0x5f, 0x47, 0xce, 0x20, 0x40, 0x7a, 0x7b,
	0x0f, 0x60, 0x75, 0x1b, 0xca, 0x20, 0xaa, 0xbe, 0x80, 0x3b, 0x7c, 0x14, 0x2c, 0xf9, 0x00, 0xb5,
	0x80, 0x3b, 0x7d, 0xd2, 0x43, 0x1a, 0xde, 0x80, 0x75, 0xc7, 0xf7, 0x07, 0x2e, 0xe9, 0x21, 0xbd,
	0x7d, 0x01, 0xd8, 0xbf, 0x79, 0x11, 0xa5, 0x48, 0x3e, 0x57, 0x79, 0x0d, 0x2f, 0x7d, 0x46, 0x8f,
	0x48, 0x97, 0x87, 0xa5, 0x5a, 0xf0, 0xb7, 0xdc, 0x36, 0x6c, 0xde, 0xe0, 0x0b, 0x21, 0xad, 0x84,
	0x48, 0xcf, 0xe5, 0x41, 0xe8, 0x13, 0xaf, 0xe7, 0x7a, 0x7d, 0xa4, 0x97, 0x15, 0x79, 0xf4, 0x46,
	0x00, 0x55, 0xf6, 0x0f, 0x4f, 0x0f, 0xce, 0x13, 0x79, 0x51, 0x9c, 0x59, 0x63, 0x91, 0xda, 0xe5,
	0x3d, 0x59, 0xae, 0x00, 0xfb, 0xbf, 0xd6, 0xc2, 0x59, 0x4d, 0x6d, 0x84, 0x8f, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xef, 0x87, 0x69, 0xf6, 0x4e, 0x04, 0x00, 0x00,
}
