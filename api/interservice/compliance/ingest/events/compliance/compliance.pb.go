// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/interservice/compliance/ingest/events/compliance/compliance.proto

package compliance

import (
	fmt "fmt"
	common "github.com/chef/automate/api/interservice/compliance/common"
	inspec "github.com/chef/automate/api/interservice/compliance/ingest/events/inspec"
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

type Report struct {
	// inspec full json report fields
	Version     string             `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty" toml:"version,omitempty" mapstructure:"version,omitempty"`
	Platform    *inspec.Platform   `protobuf:"bytes,16,opt,name=platform,proto3" json:"platform,omitempty" toml:"platform,omitempty" mapstructure:"platform,omitempty"`
	Statistics  *inspec.Statistics `protobuf:"bytes,17,opt,name=statistics,proto3" json:"statistics,omitempty" toml:"statistics,omitempty" mapstructure:"statistics,omitempty"`
	Profiles    []*inspec.Profile  `protobuf:"bytes,18,rep,name=profiles,proto3" json:"profiles,omitempty" toml:"profiles,omitempty" mapstructure:"profiles,omitempty"`
	OtherChecks []string           `protobuf:"bytes,19,rep,name=other_checks,json=otherChecks,proto3" json:"other_checks,omitempty" toml:"other_checks,omitempty" mapstructure:"other_checks,omitempty"`
	// extra report fields added by the audit cookbook
	ReportUuid           string       `protobuf:"bytes,20,opt,name=report_uuid,json=reportUuid,proto3" json:"report_uuid,omitempty" toml:"report_uuid,omitempty" mapstructure:"report_uuid,omitempty"`
	NodeUuid             string       `protobuf:"bytes,21,opt,name=node_uuid,json=nodeUuid,proto3" json:"node_uuid,omitempty" toml:"node_uuid,omitempty" mapstructure:"node_uuid,omitempty"`
	JobUuid              string       `protobuf:"bytes,22,opt,name=job_uuid,json=jobUuid,proto3" json:"job_uuid,omitempty" toml:"job_uuid,omitempty" mapstructure:"job_uuid,omitempty"`
	NodeName             string       `protobuf:"bytes,23,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty" toml:"node_name,omitempty" mapstructure:"node_name,omitempty"`
	Environment          string       `protobuf:"bytes,24,opt,name=environment,proto3" json:"environment,omitempty" toml:"environment,omitempty" mapstructure:"environment,omitempty"`
	Roles                []string     `protobuf:"bytes,25,rep,name=roles,proto3" json:"roles,omitempty" toml:"roles,omitempty" mapstructure:"roles,omitempty"`
	Recipes              []string     `protobuf:"bytes,26,rep,name=recipes,proto3" json:"recipes,omitempty" toml:"recipes,omitempty" mapstructure:"recipes,omitempty"`
	EndTime              string       `protobuf:"bytes,27,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty" toml:"end_time,omitempty" mapstructure:"end_time,omitempty"`
	Type                 string       `protobuf:"bytes,28,opt,name=type,proto3" json:"type,omitempty" toml:"type,omitempty" mapstructure:"type,omitempty"`
	SourceId             string       `protobuf:"bytes,29,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty" toml:"source_id,omitempty" mapstructure:"source_id,omitempty"`
	SourceRegion         string       `protobuf:"bytes,30,opt,name=source_region,json=sourceRegion,proto3" json:"source_region,omitempty" toml:"source_region,omitempty" mapstructure:"source_region,omitempty"`
	SourceAccountId      string       `protobuf:"bytes,31,opt,name=source_account_id,json=sourceAccountId,proto3" json:"source_account_id,omitempty" toml:"source_account_id,omitempty" mapstructure:"source_account_id,omitempty"`
	PolicyName           string       `protobuf:"bytes,32,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty" toml:"policy_name,omitempty" mapstructure:"policy_name,omitempty"`
	PolicyGroup          string       `protobuf:"bytes,33,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty" toml:"policy_group,omitempty" mapstructure:"policy_group,omitempty"`
	OrganizationName     string       `protobuf:"bytes,34,opt,name=organization_name,json=organizationName,proto3" json:"organization_name,omitempty" toml:"organization_name,omitempty" mapstructure:"organization_name,omitempty"`
	SourceFqdn           string       `protobuf:"bytes,35,opt,name=source_fqdn,json=sourceFqdn,proto3" json:"source_fqdn,omitempty" toml:"source_fqdn,omitempty" mapstructure:"source_fqdn,omitempty"`
	ChefTags             []string     `protobuf:"bytes,36,rep,name=chef_tags,json=chefTags,proto3" json:"chef_tags,omitempty" toml:"chef_tags,omitempty" mapstructure:"chef_tags,omitempty"`
	Ipaddress            string       `protobuf:"bytes,37,opt,name=ipaddress,proto3" json:"ipaddress,omitempty" toml:"ipaddress,omitempty" mapstructure:"ipaddress,omitempty"`
	Fqdn                 string       `protobuf:"bytes,38,opt,name=fqdn,proto3" json:"fqdn,omitempty" toml:"fqdn,omitempty" mapstructure:"fqdn,omitempty"`
	Tags                 []*common.Kv `protobuf:"bytes,39,rep,name=tags,proto3" json:"tags,omitempty" toml:"tags,omitempty" mapstructure:"tags,omitempty"`
	AutomateManagerId    string       `protobuf:"bytes,40,opt,name=automate_manager_id,json=automateManagerId,proto3" json:"automate_manager_id,omitempty" toml:"automate_manager_id,omitempty" mapstructure:"automate_manager_id,omitempty"`
	RunTimeLimit         float32      `protobuf:"fixed32,41,opt,name=run_time_limit,json=runTimeLimit,proto3" json:"run_time_limit,omitempty" toml:"run_time_limit,omitempty" mapstructure:"run_time_limit,omitempty"`
	AutomateManagerType  string       `protobuf:"bytes,42,opt,name=automate_manager_type,json=automateManagerType,proto3" json:"automate_manager_type,omitempty" toml:"automate_manager_type,omitempty" mapstructure:"automate_manager_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte       `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32        `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d8e0277b12342b3, []int{0}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Report) GetPlatform() *inspec.Platform {
	if m != nil {
		return m.Platform
	}
	return nil
}

func (m *Report) GetStatistics() *inspec.Statistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (m *Report) GetProfiles() []*inspec.Profile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func (m *Report) GetOtherChecks() []string {
	if m != nil {
		return m.OtherChecks
	}
	return nil
}

func (m *Report) GetReportUuid() string {
	if m != nil {
		return m.ReportUuid
	}
	return ""
}

func (m *Report) GetNodeUuid() string {
	if m != nil {
		return m.NodeUuid
	}
	return ""
}

func (m *Report) GetJobUuid() string {
	if m != nil {
		return m.JobUuid
	}
	return ""
}

func (m *Report) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Report) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *Report) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Report) GetRecipes() []string {
	if m != nil {
		return m.Recipes
	}
	return nil
}

func (m *Report) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *Report) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Report) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *Report) GetSourceRegion() string {
	if m != nil {
		return m.SourceRegion
	}
	return ""
}

func (m *Report) GetSourceAccountId() string {
	if m != nil {
		return m.SourceAccountId
	}
	return ""
}

func (m *Report) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *Report) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

func (m *Report) GetOrganizationName() string {
	if m != nil {
		return m.OrganizationName
	}
	return ""
}

func (m *Report) GetSourceFqdn() string {
	if m != nil {
		return m.SourceFqdn
	}
	return ""
}

func (m *Report) GetChefTags() []string {
	if m != nil {
		return m.ChefTags
	}
	return nil
}

func (m *Report) GetIpaddress() string {
	if m != nil {
		return m.Ipaddress
	}
	return ""
}

func (m *Report) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *Report) GetTags() []*common.Kv {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Report) GetAutomateManagerId() string {
	if m != nil {
		return m.AutomateManagerId
	}
	return ""
}

func (m *Report) GetRunTimeLimit() float32 {
	if m != nil {
		return m.RunTimeLimit
	}
	return 0
}

func (m *Report) GetAutomateManagerType() string {
	if m != nil {
		return m.AutomateManagerType
	}
	return ""
}

func init() {
	proto.RegisterType((*Report)(nil), "chef.automate.domain.compliance.ingest.events.compliance.Report")
}

func init() {
	proto.RegisterFile("api/interservice/compliance/ingest/events/compliance/compliance.proto", fileDescriptor_2d8e0277b12342b3)
}

var fileDescriptor_2d8e0277b12342b3 = []byte{
	// 667 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6e, 0x13, 0x31,
	0x14, 0x55, 0x68, 0x29, 0x89, 0x13, 0xa0, 0x71, 0x5b, 0x70, 0x1f, 0xd0, 0xf4, 0x01, 0x84, 0x22,
	0x65, 0xa4, 0xb2, 0x61, 0xd3, 0x8a, 0x87, 0x00, 0x55, 0x50, 0x40, 0x43, 0x59, 0xd0, 0xcd, 0xe0,
	0x8c, 0x6f, 0x26, 0x2e, 0x19, 0x7b, 0x6a, 0x7b, 0x22, 0x95, 0x8f, 0xe6, 0x1b, 0x90, 0xaf, 0x27,
	0x6d, 0x00, 0x89, 0x47, 0x57, 0xb1, 0xcf, 0xb9, 0xf7, 0xf8, 0xf8, 0xf8, 0x66, 0xc8, 0x4b, 0x5e,
	0xc8, 0x48, 0x2a, 0x07, 0xc6, 0x82, 0x19, 0xcb, 0x14, 0xa2, 0x54, 0xe7, 0xc5, 0x48, 0x72, 0x95,
	0x42, 0x24, 0x55, 0x06, 0xd6, 0x45, 0x30, 0x06, 0xe5, 0xec, 0x34, 0x71, 0xb1, 0xec, 0x15, 0x46,
	0x3b, 0x4d, 0x9f, 0xa4, 0x43, 0x18, 0xf4, 0x78, 0xe9, 0x74, 0xce, 0x1d, 0xf4, 0x84, 0xce, 0xb9,
	0x54, 0xbd, 0xa9, 0xb2, 0x20, 0xd5, 0x0b, 0x52, 0x53, 0xc4, 0xca, 0xde, 0xbf, 0x1b, 0x90, 0xca,
	0x16, 0x90, 0x56, 0x3f, 0xe1, 0xe0, 0x95, 0xe8, 0x4f, 0xed, 0xa9, 0xce, 0x73, 0xad, 0xaa, 0x9f,
	0xd0, 0xb0, 0xf9, 0xbd, 0x4e, 0xe6, 0x62, 0x28, 0xb4, 0x71, 0x94, 0x91, 0x6b, 0x63, 0x30, 0x56,
	0x6a, 0xc5, 0x6a, 0x9d, 0x5a, 0xb7, 0x11, 0x4f, 0xb6, 0xf4, 0x98, 0xd4, 0x8b, 0x11, 0x77, 0x03,
	0x6d, 0x72, 0x36, 0xdf, 0xa9, 0x75, 0x9b, 0xbb, 0xfb, 0xbd, 0xff, 0xbb, 0x61, 0x65, 0xf2, 0x43,
	0xa5, 0x12, 0x9f, 0xeb, 0xd1, 0x2f, 0x84, 0x58, 0xc7, 0x9d, 0xb4, 0x4e, 0xa6, 0x96, 0xb5, 0x51,
	0xfd, 0xe9, 0xe5, 0xd4, 0x3f, 0x9e, 0xeb, 0xc4, 0x53, 0x9a, 0xf4, 0x33, 0xa9, 0x17, 0x46, 0x0f,
	0xe4, 0x08, 0x2c, 0xa3, 0x9d, 0x99, 0x6e, 0x73, 0x77, 0xef, 0x92, 0xee, 0x83, 0x4a, 0x7c, 0x2e,
	0x47, 0x37, 0x48, 0x4b, 0xbb, 0x21, 0x98, 0x24, 0x1d, 0x42, 0xfa, 0xd5, 0xb2, 0x85, 0xce, 0x4c,
	0xb7, 0x11, 0x37, 0x11, 0x7b, 0x81, 0x10, 0x5d, 0x27, 0x4d, 0x83, 0xf9, 0x26, 0x65, 0x29, 0x05,
	0x5b, 0xc4, 0x64, 0x49, 0x80, 0x3e, 0x95, 0x52, 0xd0, 0x55, 0xd2, 0x50, 0x5a, 0x40, 0xa0, 0x97,
	0x90, 0xae, 0x7b, 0x00, 0xc9, 0x65, 0x52, 0x3f, 0xd1, 0xfd, 0xc0, 0xdd, 0x0a, 0x8f, 0x72, 0xa2,
	0xfb, 0x3f, 0xf5, 0x29, 0x9e, 0x03, 0xbb, 0x7d, 0xd1, 0xf7, 0x8e, 0xe7, 0x40, 0x3b, 0xa4, 0x09,
	0x6a, 0x2c, 0x8d, 0x56, 0x39, 0x28, 0xc7, 0x18, 0xd2, 0xd3, 0x10, 0x5d, 0x24, 0x57, 0x8d, 0xf6,
	0x91, 0x2c, 0xa3, 0xe7, 0xb0, 0xf1, 0x33, 0x60, 0x20, 0x95, 0x05, 0x58, 0xb6, 0x82, 0xf8, 0x64,
	0xeb, 0x9d, 0x80, 0x12, 0x89, 0x93, 0x39, 0xb0, 0xd5, 0xe0, 0x04, 0x94, 0x38, 0x92, 0x39, 0x50,
	0x4a, 0x66, 0xdd, 0x59, 0x01, 0x6c, 0x0d, 0x61, 0x5c, 0x7b, 0x77, 0x56, 0x97, 0x26, 0x85, 0x44,
	0x0a, 0x76, 0x27, 0xb8, 0x0b, 0xc0, 0x81, 0xa0, 0x5b, 0xe4, 0x7a, 0x45, 0x1a, 0xc8, 0xfc, 0xbc,
	0xdd, 0xc5, 0x82, 0x56, 0x00, 0x63, 0xc4, 0xe8, 0x0e, 0x69, 0x57, 0x45, 0x3c, 0x4d, 0x75, 0xa9,
	0x9c, 0x57, 0x5a, 0xc7, 0xc2, 0x9b, 0x81, 0x78, 0x16, 0xf0, 0x03, 0xe1, 0x43, 0x2e, 0xf4, 0x48,
	0xa6, 0x67, 0x21, 0x8d, 0x4e, 0x08, 0x39, 0x40, 0x98, 0xc7, 0x06, 0x69, 0x55, 0x05, 0x99, 0xd1,
	0x65, 0xc1, 0x36, 0x42, 0x20, 0x01, 0x7b, 0xed, 0x21, 0xfa, 0x88, 0xb4, 0xb5, 0xc9, 0xb8, 0x92,
	0xdf, 0xb8, 0x93, 0x5a, 0x05, 0xa5, 0x4d, 0xac, 0x9b, 0x9f, 0x26, 0x50, 0x6f, 0x9d, 0x34, 0x2b,
	0x73, 0x83, 0x53, 0xa1, 0xd8, 0x56, 0x38, 0x30, 0x40, 0xaf, 0x4e, 0x85, 0xf2, 0xf7, 0xf7, 0x33,
	0x96, 0x38, 0x9e, 0x59, 0xb6, 0x8d, 0x51, 0xd6, 0x3d, 0x70, 0xc4, 0x33, 0x4b, 0xd7, 0x48, 0x43,
	0x16, 0x5c, 0x08, 0x03, 0xd6, 0xb2, 0x7b, 0xd8, 0x7b, 0x01, 0xf8, 0x38, 0x51, 0xf4, 0x7e, 0x88,
	0xd3, 0xaf, 0xe9, 0x3e, 0x99, 0x45, 0xa5, 0x07, 0x38, 0xbf, 0x3b, 0x7f, 0x9d, 0xdf, 0xea, 0x3f,
	0xfe, 0x66, 0x1c, 0x63, 0x1f, 0xed, 0x91, 0x85, 0x49, 0x75, 0x92, 0x73, 0xc5, 0x33, 0x30, 0x3e,
	0xce, 0x2e, 0x1e, 0xd1, 0x9e, 0x50, 0x87, 0x81, 0x39, 0x10, 0x74, 0x9b, 0xdc, 0x30, 0xa5, 0xc2,
	0xd7, 0x4e, 0x46, 0x32, 0x97, 0x8e, 0x3d, 0xec, 0xd4, 0xba, 0x57, 0xe2, 0x96, 0x29, 0x95, 0x7f,
	0xf3, 0xb7, 0x1e, 0xa3, 0xbb, 0x64, 0xe9, 0x37, 0x55, 0x9c, 0x84, 0x1d, 0xd4, 0x5d, 0xf8, 0x45,
	0xf7, 0xe8, 0xac, 0x80, 0xe7, 0xef, 0x8f, 0x0f, 0x33, 0xe9, 0x86, 0x65, 0xdf, 0x7b, 0x8c, 0xfc,
	0x3d, 0xa2, 0x49, 0x59, 0x74, 0x99, 0x8f, 0x6f, 0x7f, 0x0e, 0x3f, 0x64, 0x8f, 0x7f, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x4f, 0xba, 0x2a, 0x3a, 0xbb, 0x05, 0x00, 0x00,
}
