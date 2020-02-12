// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/response/policy.proto

package response

import (
	fmt "fmt"
	common "github.com/chef/automate/components/automate-gateway/api/iam/v2/common"
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

type CreatePolicyResp struct {
	Policy               *common.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreatePolicyResp) Reset()         { *m = CreatePolicyResp{} }
func (m *CreatePolicyResp) String() string { return proto.CompactTextString(m) }
func (*CreatePolicyResp) ProtoMessage()    {}
func (*CreatePolicyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{0}
}

func (m *CreatePolicyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePolicyResp.Unmarshal(m, b)
}
func (m *CreatePolicyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePolicyResp.Marshal(b, m, deterministic)
}
func (m *CreatePolicyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePolicyResp.Merge(m, src)
}
func (m *CreatePolicyResp) XXX_Size() int {
	return xxx_messageInfo_CreatePolicyResp.Size(m)
}
func (m *CreatePolicyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePolicyResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePolicyResp proto.InternalMessageInfo

func (m *CreatePolicyResp) GetPolicy() *common.Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type GetPolicyResp struct {
	Policy               *common.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetPolicyResp) Reset()         { *m = GetPolicyResp{} }
func (m *GetPolicyResp) String() string { return proto.CompactTextString(m) }
func (*GetPolicyResp) ProtoMessage()    {}
func (*GetPolicyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{1}
}

func (m *GetPolicyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPolicyResp.Unmarshal(m, b)
}
func (m *GetPolicyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPolicyResp.Marshal(b, m, deterministic)
}
func (m *GetPolicyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPolicyResp.Merge(m, src)
}
func (m *GetPolicyResp) XXX_Size() int {
	return xxx_messageInfo_GetPolicyResp.Size(m)
}
func (m *GetPolicyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPolicyResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetPolicyResp proto.InternalMessageInfo

func (m *GetPolicyResp) GetPolicy() *common.Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type UpdatePolicyResp struct {
	Policy               *common.Policy `protobuf:"bytes,1,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdatePolicyResp) Reset()         { *m = UpdatePolicyResp{} }
func (m *UpdatePolicyResp) String() string { return proto.CompactTextString(m) }
func (*UpdatePolicyResp) ProtoMessage()    {}
func (*UpdatePolicyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{2}
}

func (m *UpdatePolicyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePolicyResp.Unmarshal(m, b)
}
func (m *UpdatePolicyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePolicyResp.Marshal(b, m, deterministic)
}
func (m *UpdatePolicyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePolicyResp.Merge(m, src)
}
func (m *UpdatePolicyResp) XXX_Size() int {
	return xxx_messageInfo_UpdatePolicyResp.Size(m)
}
func (m *UpdatePolicyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePolicyResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePolicyResp proto.InternalMessageInfo

func (m *UpdatePolicyResp) GetPolicy() *common.Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type ListPoliciesResp struct {
	Policies             []*common.Policy `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListPoliciesResp) Reset()         { *m = ListPoliciesResp{} }
func (m *ListPoliciesResp) String() string { return proto.CompactTextString(m) }
func (*ListPoliciesResp) ProtoMessage()    {}
func (*ListPoliciesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{3}
}

func (m *ListPoliciesResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPoliciesResp.Unmarshal(m, b)
}
func (m *ListPoliciesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPoliciesResp.Marshal(b, m, deterministic)
}
func (m *ListPoliciesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPoliciesResp.Merge(m, src)
}
func (m *ListPoliciesResp) XXX_Size() int {
	return xxx_messageInfo_ListPoliciesResp.Size(m)
}
func (m *ListPoliciesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPoliciesResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListPoliciesResp proto.InternalMessageInfo

func (m *ListPoliciesResp) GetPolicies() []*common.Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type AddPolicyMembersResp struct {
	Members              []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPolicyMembersResp) Reset()         { *m = AddPolicyMembersResp{} }
func (m *AddPolicyMembersResp) String() string { return proto.CompactTextString(m) }
func (*AddPolicyMembersResp) ProtoMessage()    {}
func (*AddPolicyMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{4}
}

func (m *AddPolicyMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPolicyMembersResp.Unmarshal(m, b)
}
func (m *AddPolicyMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPolicyMembersResp.Marshal(b, m, deterministic)
}
func (m *AddPolicyMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPolicyMembersResp.Merge(m, src)
}
func (m *AddPolicyMembersResp) XXX_Size() int {
	return xxx_messageInfo_AddPolicyMembersResp.Size(m)
}
func (m *AddPolicyMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPolicyMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddPolicyMembersResp proto.InternalMessageInfo

func (m *AddPolicyMembersResp) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type DeletePolicyResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePolicyResp) Reset()         { *m = DeletePolicyResp{} }
func (m *DeletePolicyResp) String() string { return proto.CompactTextString(m) }
func (*DeletePolicyResp) ProtoMessage()    {}
func (*DeletePolicyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{5}
}

func (m *DeletePolicyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePolicyResp.Unmarshal(m, b)
}
func (m *DeletePolicyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePolicyResp.Marshal(b, m, deterministic)
}
func (m *DeletePolicyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePolicyResp.Merge(m, src)
}
func (m *DeletePolicyResp) XXX_Size() int {
	return xxx_messageInfo_DeletePolicyResp.Size(m)
}
func (m *DeletePolicyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePolicyResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePolicyResp proto.InternalMessageInfo

type GetPolicyVersionResp struct {
	Version              *common.Version `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetPolicyVersionResp) Reset()         { *m = GetPolicyVersionResp{} }
func (m *GetPolicyVersionResp) String() string { return proto.CompactTextString(m) }
func (*GetPolicyVersionResp) ProtoMessage()    {}
func (*GetPolicyVersionResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{6}
}

func (m *GetPolicyVersionResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPolicyVersionResp.Unmarshal(m, b)
}
func (m *GetPolicyVersionResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPolicyVersionResp.Marshal(b, m, deterministic)
}
func (m *GetPolicyVersionResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPolicyVersionResp.Merge(m, src)
}
func (m *GetPolicyVersionResp) XXX_Size() int {
	return xxx_messageInfo_GetPolicyVersionResp.Size(m)
}
func (m *GetPolicyVersionResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPolicyVersionResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetPolicyVersionResp proto.InternalMessageInfo

func (m *GetPolicyVersionResp) GetVersion() *common.Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type ResetToV1Resp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetToV1Resp) Reset()         { *m = ResetToV1Resp{} }
func (m *ResetToV1Resp) String() string { return proto.CompactTextString(m) }
func (*ResetToV1Resp) ProtoMessage()    {}
func (*ResetToV1Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{7}
}

func (m *ResetToV1Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetToV1Resp.Unmarshal(m, b)
}
func (m *ResetToV1Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetToV1Resp.Marshal(b, m, deterministic)
}
func (m *ResetToV1Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetToV1Resp.Merge(m, src)
}
func (m *ResetToV1Resp) XXX_Size() int {
	return xxx_messageInfo_ResetToV1Resp.Size(m)
}
func (m *ResetToV1Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetToV1Resp.DiscardUnknown(m)
}

var xxx_messageInfo_ResetToV1Resp proto.InternalMessageInfo

type ListPolicyMembersResp struct {
	Members              []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPolicyMembersResp) Reset()         { *m = ListPolicyMembersResp{} }
func (m *ListPolicyMembersResp) String() string { return proto.CompactTextString(m) }
func (*ListPolicyMembersResp) ProtoMessage()    {}
func (*ListPolicyMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{8}
}

func (m *ListPolicyMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPolicyMembersResp.Unmarshal(m, b)
}
func (m *ListPolicyMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPolicyMembersResp.Marshal(b, m, deterministic)
}
func (m *ListPolicyMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPolicyMembersResp.Merge(m, src)
}
func (m *ListPolicyMembersResp) XXX_Size() int {
	return xxx_messageInfo_ListPolicyMembersResp.Size(m)
}
func (m *ListPolicyMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPolicyMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListPolicyMembersResp proto.InternalMessageInfo

func (m *ListPolicyMembersResp) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type ReplacePolicyMembersResp struct {
	Members              []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplacePolicyMembersResp) Reset()         { *m = ReplacePolicyMembersResp{} }
func (m *ReplacePolicyMembersResp) String() string { return proto.CompactTextString(m) }
func (*ReplacePolicyMembersResp) ProtoMessage()    {}
func (*ReplacePolicyMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{9}
}

func (m *ReplacePolicyMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplacePolicyMembersResp.Unmarshal(m, b)
}
func (m *ReplacePolicyMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplacePolicyMembersResp.Marshal(b, m, deterministic)
}
func (m *ReplacePolicyMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplacePolicyMembersResp.Merge(m, src)
}
func (m *ReplacePolicyMembersResp) XXX_Size() int {
	return xxx_messageInfo_ReplacePolicyMembersResp.Size(m)
}
func (m *ReplacePolicyMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplacePolicyMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_ReplacePolicyMembersResp proto.InternalMessageInfo

func (m *ReplacePolicyMembersResp) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type RemovePolicyMembersResp struct {
	Members              []string `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemovePolicyMembersResp) Reset()         { *m = RemovePolicyMembersResp{} }
func (m *RemovePolicyMembersResp) String() string { return proto.CompactTextString(m) }
func (*RemovePolicyMembersResp) ProtoMessage()    {}
func (*RemovePolicyMembersResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{10}
}

func (m *RemovePolicyMembersResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemovePolicyMembersResp.Unmarshal(m, b)
}
func (m *RemovePolicyMembersResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemovePolicyMembersResp.Marshal(b, m, deterministic)
}
func (m *RemovePolicyMembersResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemovePolicyMembersResp.Merge(m, src)
}
func (m *RemovePolicyMembersResp) XXX_Size() int {
	return xxx_messageInfo_RemovePolicyMembersResp.Size(m)
}
func (m *RemovePolicyMembersResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RemovePolicyMembersResp.DiscardUnknown(m)
}

var xxx_messageInfo_RemovePolicyMembersResp proto.InternalMessageInfo

func (m *RemovePolicyMembersResp) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type CreateRoleResp struct {
	Role                 *common.Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateRoleResp) Reset()         { *m = CreateRoleResp{} }
func (m *CreateRoleResp) String() string { return proto.CompactTextString(m) }
func (*CreateRoleResp) ProtoMessage()    {}
func (*CreateRoleResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{11}
}

func (m *CreateRoleResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoleResp.Unmarshal(m, b)
}
func (m *CreateRoleResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoleResp.Marshal(b, m, deterministic)
}
func (m *CreateRoleResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleResp.Merge(m, src)
}
func (m *CreateRoleResp) XXX_Size() int {
	return xxx_messageInfo_CreateRoleResp.Size(m)
}
func (m *CreateRoleResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleResp proto.InternalMessageInfo

func (m *CreateRoleResp) GetRole() *common.Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type GetRoleResp struct {
	Role                 *common.Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetRoleResp) Reset()         { *m = GetRoleResp{} }
func (m *GetRoleResp) String() string { return proto.CompactTextString(m) }
func (*GetRoleResp) ProtoMessage()    {}
func (*GetRoleResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{12}
}

func (m *GetRoleResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoleResp.Unmarshal(m, b)
}
func (m *GetRoleResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoleResp.Marshal(b, m, deterministic)
}
func (m *GetRoleResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoleResp.Merge(m, src)
}
func (m *GetRoleResp) XXX_Size() int {
	return xxx_messageInfo_GetRoleResp.Size(m)
}
func (m *GetRoleResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoleResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoleResp proto.InternalMessageInfo

func (m *GetRoleResp) GetRole() *common.Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type DeleteRoleResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRoleResp) Reset()         { *m = DeleteRoleResp{} }
func (m *DeleteRoleResp) String() string { return proto.CompactTextString(m) }
func (*DeleteRoleResp) ProtoMessage()    {}
func (*DeleteRoleResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{13}
}

func (m *DeleteRoleResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRoleResp.Unmarshal(m, b)
}
func (m *DeleteRoleResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRoleResp.Marshal(b, m, deterministic)
}
func (m *DeleteRoleResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRoleResp.Merge(m, src)
}
func (m *DeleteRoleResp) XXX_Size() int {
	return xxx_messageInfo_DeleteRoleResp.Size(m)
}
func (m *DeleteRoleResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRoleResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRoleResp proto.InternalMessageInfo

type UpdateRoleResp struct {
	Role                 *common.Role `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateRoleResp) Reset()         { *m = UpdateRoleResp{} }
func (m *UpdateRoleResp) String() string { return proto.CompactTextString(m) }
func (*UpdateRoleResp) ProtoMessage()    {}
func (*UpdateRoleResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{14}
}

func (m *UpdateRoleResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRoleResp.Unmarshal(m, b)
}
func (m *UpdateRoleResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRoleResp.Marshal(b, m, deterministic)
}
func (m *UpdateRoleResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRoleResp.Merge(m, src)
}
func (m *UpdateRoleResp) XXX_Size() int {
	return xxx_messageInfo_UpdateRoleResp.Size(m)
}
func (m *UpdateRoleResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRoleResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRoleResp proto.InternalMessageInfo

func (m *UpdateRoleResp) GetRole() *common.Role {
	if m != nil {
		return m.Role
	}
	return nil
}

type ListRolesResp struct {
	Roles                []*common.Role `protobuf:"bytes,1,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListRolesResp) Reset()         { *m = ListRolesResp{} }
func (m *ListRolesResp) String() string { return proto.CompactTextString(m) }
func (*ListRolesResp) ProtoMessage()    {}
func (*ListRolesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{15}
}

func (m *ListRolesResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRolesResp.Unmarshal(m, b)
}
func (m *ListRolesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRolesResp.Marshal(b, m, deterministic)
}
func (m *ListRolesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRolesResp.Merge(m, src)
}
func (m *ListRolesResp) XXX_Size() int {
	return xxx_messageInfo_ListRolesResp.Size(m)
}
func (m *ListRolesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRolesResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListRolesResp proto.InternalMessageInfo

func (m *ListRolesResp) GetRoles() []*common.Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type GetProjectResp struct {
	Project              *common.Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetProjectResp) Reset()         { *m = GetProjectResp{} }
func (m *GetProjectResp) String() string { return proto.CompactTextString(m) }
func (*GetProjectResp) ProtoMessage()    {}
func (*GetProjectResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{16}
}

func (m *GetProjectResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProjectResp.Unmarshal(m, b)
}
func (m *GetProjectResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProjectResp.Marshal(b, m, deterministic)
}
func (m *GetProjectResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProjectResp.Merge(m, src)
}
func (m *GetProjectResp) XXX_Size() int {
	return xxx_messageInfo_GetProjectResp.Size(m)
}
func (m *GetProjectResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProjectResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetProjectResp proto.InternalMessageInfo

func (m *GetProjectResp) GetProject() *common.Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type ListProjectsResp struct {
	Projects             []*common.Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListProjectsResp) Reset()         { *m = ListProjectsResp{} }
func (m *ListProjectsResp) String() string { return proto.CompactTextString(m) }
func (*ListProjectsResp) ProtoMessage()    {}
func (*ListProjectsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{17}
}

func (m *ListProjectsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListProjectsResp.Unmarshal(m, b)
}
func (m *ListProjectsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListProjectsResp.Marshal(b, m, deterministic)
}
func (m *ListProjectsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListProjectsResp.Merge(m, src)
}
func (m *ListProjectsResp) XXX_Size() int {
	return xxx_messageInfo_ListProjectsResp.Size(m)
}
func (m *ListProjectsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ListProjectsResp.DiscardUnknown(m)
}

var xxx_messageInfo_ListProjectsResp proto.InternalMessageInfo

func (m *ListProjectsResp) GetProjects() []*common.Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

type CreateProjectResp struct {
	Project              *common.Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateProjectResp) Reset()         { *m = CreateProjectResp{} }
func (m *CreateProjectResp) String() string { return proto.CompactTextString(m) }
func (*CreateProjectResp) ProtoMessage()    {}
func (*CreateProjectResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{18}
}

func (m *CreateProjectResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProjectResp.Unmarshal(m, b)
}
func (m *CreateProjectResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProjectResp.Marshal(b, m, deterministic)
}
func (m *CreateProjectResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProjectResp.Merge(m, src)
}
func (m *CreateProjectResp) XXX_Size() int {
	return xxx_messageInfo_CreateProjectResp.Size(m)
}
func (m *CreateProjectResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProjectResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProjectResp proto.InternalMessageInfo

func (m *CreateProjectResp) GetProject() *common.Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type UpdateProjectResp struct {
	Project              *common.Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateProjectResp) Reset()         { *m = UpdateProjectResp{} }
func (m *UpdateProjectResp) String() string { return proto.CompactTextString(m) }
func (*UpdateProjectResp) ProtoMessage()    {}
func (*UpdateProjectResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{19}
}

func (m *UpdateProjectResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProjectResp.Unmarshal(m, b)
}
func (m *UpdateProjectResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProjectResp.Marshal(b, m, deterministic)
}
func (m *UpdateProjectResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProjectResp.Merge(m, src)
}
func (m *UpdateProjectResp) XXX_Size() int {
	return xxx_messageInfo_UpdateProjectResp.Size(m)
}
func (m *UpdateProjectResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProjectResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProjectResp proto.InternalMessageInfo

func (m *UpdateProjectResp) GetProject() *common.Project {
	if m != nil {
		return m.Project
	}
	return nil
}

type DeleteProjectResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProjectResp) Reset()         { *m = DeleteProjectResp{} }
func (m *DeleteProjectResp) String() string { return proto.CompactTextString(m) }
func (*DeleteProjectResp) ProtoMessage()    {}
func (*DeleteProjectResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94e203b4b99ed5f, []int{20}
}

func (m *DeleteProjectResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProjectResp.Unmarshal(m, b)
}
func (m *DeleteProjectResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProjectResp.Marshal(b, m, deterministic)
}
func (m *DeleteProjectResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProjectResp.Merge(m, src)
}
func (m *DeleteProjectResp) XXX_Size() int {
	return xxx_messageInfo_DeleteProjectResp.Size(m)
}
func (m *DeleteProjectResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProjectResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProjectResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreatePolicyResp)(nil), "chef.automate.api.iam.v2.CreatePolicyResp")
	proto.RegisterType((*GetPolicyResp)(nil), "chef.automate.api.iam.v2.GetPolicyResp")
	proto.RegisterType((*UpdatePolicyResp)(nil), "chef.automate.api.iam.v2.UpdatePolicyResp")
	proto.RegisterType((*ListPoliciesResp)(nil), "chef.automate.api.iam.v2.ListPoliciesResp")
	proto.RegisterType((*AddPolicyMembersResp)(nil), "chef.automate.api.iam.v2.AddPolicyMembersResp")
	proto.RegisterType((*DeletePolicyResp)(nil), "chef.automate.api.iam.v2.DeletePolicyResp")
	proto.RegisterType((*GetPolicyVersionResp)(nil), "chef.automate.api.iam.v2.GetPolicyVersionResp")
	proto.RegisterType((*ResetToV1Resp)(nil), "chef.automate.api.iam.v2.ResetToV1Resp")
	proto.RegisterType((*ListPolicyMembersResp)(nil), "chef.automate.api.iam.v2.ListPolicyMembersResp")
	proto.RegisterType((*ReplacePolicyMembersResp)(nil), "chef.automate.api.iam.v2.ReplacePolicyMembersResp")
	proto.RegisterType((*RemovePolicyMembersResp)(nil), "chef.automate.api.iam.v2.RemovePolicyMembersResp")
	proto.RegisterType((*CreateRoleResp)(nil), "chef.automate.api.iam.v2.CreateRoleResp")
	proto.RegisterType((*GetRoleResp)(nil), "chef.automate.api.iam.v2.GetRoleResp")
	proto.RegisterType((*DeleteRoleResp)(nil), "chef.automate.api.iam.v2.DeleteRoleResp")
	proto.RegisterType((*UpdateRoleResp)(nil), "chef.automate.api.iam.v2.UpdateRoleResp")
	proto.RegisterType((*ListRolesResp)(nil), "chef.automate.api.iam.v2.ListRolesResp")
	proto.RegisterType((*GetProjectResp)(nil), "chef.automate.api.iam.v2.GetProjectResp")
	proto.RegisterType((*ListProjectsResp)(nil), "chef.automate.api.iam.v2.ListProjectsResp")
	proto.RegisterType((*CreateProjectResp)(nil), "chef.automate.api.iam.v2.CreateProjectResp")
	proto.RegisterType((*UpdateProjectResp)(nil), "chef.automate.api.iam.v2.UpdateProjectResp")
	proto.RegisterType((*DeleteProjectResp)(nil), "chef.automate.api.iam.v2.DeleteProjectResp")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/response/policy.proto", fileDescriptor_b94e203b4b99ed5f)
}

var fileDescriptor_b94e203b4b99ed5f = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x15, 0x01, 0x6d, 0x99, 0x2a, 0xc1, 0x09, 0x45, 0x58, 0x1c, 0x50, 0xd8, 0x13, 0x17,
	0xbc, 0x34, 0xed, 0x01, 0x41, 0x39, 0x14, 0x8a, 0x0a, 0x28, 0x95, 0x82, 0x81, 0x1e, 0xb8, 0x6d,
	0x9c, 0xa1, 0x5d, 0xe4, 0xf5, 0xac, 0xbc, 0xdb, 0xa0, 0x7e, 0x7b, 0xe4, 0xfd, 0x63, 0xe0, 0x10,
	0x35, 0x51, 0x7d, 0x9c, 0xf5, 0xbc, 0xdf, 0xbe, 0x19, 0xcf, 0x0e, 0x1c, 0x15, 0xa4, 0x34, 0x55,
	0x58, 0x59, 0xc3, 0xc5, 0x95, 0x25, 0x25, 0x2c, 0xbe, 0xb8, 0x10, 0x16, 0x7f, 0x8b, 0x6b, 0x2e,
	0xb4, 0xe4, 0x52, 0x28, 0xbe, 0x9c, 0xf0, 0x1a, 0x8d, 0xa6, 0xca, 0x20, 0xd7, 0x54, 0xca, 0xe2,
	0x3a, 0xd3, 0x35, 0x59, 0x1a, 0xa5, 0xc5, 0x25, 0xfe, 0xcc, 0xa2, 0x2e, 0x13, 0x5a, 0x66, 0x52,
	0xa8, 0x6c, 0x39, 0x79, 0xf2, 0x7a, 0x4d, 0x6e, 0x41, 0x4a, 0x51, 0xf5, 0x1f, 0x95, 0x4d, 0x21,
	0x79, 0x5f, 0xa3, 0xb0, 0x38, 0x73, 0xa7, 0x39, 0x1a, 0x3d, 0x7a, 0x05, 0x5b, 0x3e, 0x27, 0xed,
	0x8d, 0x7b, 0xcf, 0x77, 0x27, 0xe3, 0x6c, 0xd5, 0xd5, 0x59, 0x50, 0x85, 0x7c, 0xf6, 0x09, 0xfa,
	0xa7, 0x68, 0x3b, 0x41, 0x4d, 0x21, 0xf9, 0xae, 0x17, 0x5d, 0x19, 0x9b, 0x41, 0x32, 0x95, 0xc6,
	0x3b, 0x93, 0x68, 0x1c, 0xed, 0x08, 0x76, 0x74, 0x88, 0xd3, 0xde, 0xf8, 0xce, 0x5a, 0xbc, 0x56,
	0xc1, 0x5e, 0xc2, 0xde, 0xf1, 0x62, 0xe1, 0x8f, 0xcf, 0x50, 0xcd, 0xb1, 0xf6, 0xd4, 0x14, 0xb6,
	0x95, 0x0f, 0x1d, 0xf4, 0x7e, 0x1e, 0x43, 0x36, 0x82, 0xe4, 0x04, 0x4b, 0xfc, 0xb7, 0x22, 0xf6,
	0x15, 0xf6, 0xda, 0x86, 0x9d, 0x63, 0x6d, 0x24, 0x55, 0x8e, 0xf2, 0x06, 0xb6, 0x97, 0x3e, 0x0c,
	0xa5, 0x3e, 0x5b, 0x6d, 0x2d, 0xea, 0xa2, 0x82, 0x3d, 0x80, 0x7e, 0x8e, 0x06, 0xed, 0x37, 0x3a,
	0xdf, 0x77, 0xb7, 0xec, 0xc3, 0xa3, 0xb6, 0xfa, 0x35, 0xcd, 0x1e, 0x42, 0x9a, 0xa3, 0x2e, 0x45,
	0x81, 0x9b, 0xa8, 0x0e, 0xe0, 0x71, 0x8e, 0x8a, 0x96, 0x1b, 0x89, 0x4e, 0x60, 0xe0, 0x47, 0x30,
	0xa7, 0x12, 0x5d, 0xee, 0x04, 0xee, 0xd6, 0x54, 0x62, 0x28, 0xfd, 0xe9, 0xea, 0xd2, 0x9d, 0xc2,
	0xe5, 0xb2, 0x63, 0xd8, 0x3d, 0x45, 0x7b, 0x2b, 0x44, 0x02, 0x03, 0xff, 0x83, 0x22, 0xa5, 0xb1,
	0xe6, 0x87, 0xf0, 0x56, 0xdc, 0x0f, 0xd0, 0x6f, 0xda, 0xdf, 0x9c, 0xf8, 0x5e, 0x1c, 0xc2, 0xbd,
	0xe6, 0x43, 0x1c, 0xbb, 0x9b, 0x28, 0x3e, 0x99, 0x9d, 0xc1, 0xa0, 0x99, 0x95, 0x9a, 0x7e, 0x61,
	0x61, 0xe3, 0x94, 0x68, 0x1f, 0xde, 0x3c, 0x25, 0x51, 0x17, 0x15, 0xec, 0x4b, 0x78, 0x12, 0x3e,
	0xf4, 0xc6, 0xde, 0xc2, 0x4e, 0xf8, 0x1c, 0xbd, 0xad, 0x41, 0x6c, 0x25, 0x6c, 0x06, 0xc3, 0xb0,
	0x4c, 0xba, 0x32, 0x39, 0x83, 0x61, 0xd8, 0x02, 0x5d, 0x11, 0x1f, 0xc2, 0x30, 0xbc, 0xc2, 0xbf,
	0xc4, 0x77, 0x9f, 0x7f, 0x7c, 0xbc, 0x90, 0xf6, 0xf2, 0x6a, 0x9e, 0x15, 0xa4, 0x78, 0x03, 0x6b,
	0x17, 0x29, 0xdf, 0x70, 0x69, 0xcf, 0xb7, 0xdc, 0x62, 0x3d, 0xf8, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x42, 0x4e, 0x5f, 0xdb, 0xee, 0x05, 0x00, 0x00,
}