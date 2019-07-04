// Code generated by protoc-gen-go. DO NOT EDIT.
// source: branch.proto

package pb

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

type BranchStatus int32

const (
	BranchStatus_BS_NORMAL BranchStatus = 0
	BranchStatus_BS_FROZEN BranchStatus = 1
)

var BranchStatus_name = map[int32]string{
	0: "BS_NORMAL",
	1: "BS_FROZEN",
}

var BranchStatus_value = map[string]int32{
	"BS_NORMAL": 0,
	"BS_FROZEN": 1,
}

func (x BranchStatus) String() string {
	return proto.EnumName(BranchStatus_name, int32(x))
}

func (BranchStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_20ce1f5884b7e047, []int{0}
}

// 单纯树形结构体
type BranchTreeNode struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Children             []*BranchTreeNode `protobuf:"bytes,2,rep,name=children,proto3" json:"children"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BranchTreeNode) Reset()         { *m = BranchTreeNode{} }
func (m *BranchTreeNode) String() string { return proto.CompactTextString(m) }
func (*BranchTreeNode) ProtoMessage()    {}
func (*BranchTreeNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_20ce1f5884b7e047, []int{0}
}

func (m *BranchTreeNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BranchTreeNode.Unmarshal(m, b)
}
func (m *BranchTreeNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BranchTreeNode.Marshal(b, m, deterministic)
}
func (m *BranchTreeNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BranchTreeNode.Merge(m, src)
}
func (m *BranchTreeNode) XXX_Size() int {
	return xxx_messageInfo_BranchTreeNode.Size(m)
}
func (m *BranchTreeNode) XXX_DiscardUnknown() {
	xxx_messageInfo_BranchTreeNode.DiscardUnknown(m)
}

var xxx_messageInfo_BranchTreeNode proto.InternalMessageInfo

func (m *BranchTreeNode) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BranchTreeNode) GetChildren() []*BranchTreeNode {
	if m != nil {
		return m.Children
	}
	return nil
}

// 部门
type Branch struct {
	// 编号
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	// 名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// 备注
	Comment string `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment"`
	// 邀请码
	InviteCode []string `protobuf:"bytes,4,rep,name=invite_code,json=inviteCode,proto3" json:"inviteCode"`
	// 上级
	Parent string `protobuf:"bytes,5,opt,name=parent,proto3" json:"parent"`
	// 创建时间
	CreateTime int64 `protobuf:"varint,6,opt,name=create_time,json=createTime,proto3" json:"createTime"`
	// 部门状态
	Status               BranchStatus `protobuf:"varint,7,opt,name=status,proto3,enum=pb.BranchStatus" json:"status"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Branch) Reset()         { *m = Branch{} }
func (m *Branch) String() string { return proto.CompactTextString(m) }
func (*Branch) ProtoMessage()    {}
func (*Branch) Descriptor() ([]byte, []int) {
	return fileDescriptor_20ce1f5884b7e047, []int{1}
}

func (m *Branch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Branch.Unmarshal(m, b)
}
func (m *Branch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Branch.Marshal(b, m, deterministic)
}
func (m *Branch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Branch.Merge(m, src)
}
func (m *Branch) XXX_Size() int {
	return xxx_messageInfo_Branch.Size(m)
}
func (m *Branch) XXX_DiscardUnknown() {
	xxx_messageInfo_Branch.DiscardUnknown(m)
}

var xxx_messageInfo_Branch proto.InternalMessageInfo

func (m *Branch) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Branch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Branch) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Branch) GetInviteCode() []string {
	if m != nil {
		return m.InviteCode
	}
	return nil
}

func (m *Branch) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Branch) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Branch) GetStatus() BranchStatus {
	if m != nil {
		return m.Status
	}
	return BranchStatus_BS_NORMAL
}

func init() {
	proto.RegisterEnum("pb.BranchStatus", BranchStatus_name, BranchStatus_value)
	proto.RegisterType((*BranchTreeNode)(nil), "pb.BranchTreeNode")
	proto.RegisterType((*Branch)(nil), "pb.Branch")
}

func init() { proto.RegisterFile("branch.proto", fileDescriptor_20ce1f5884b7e047) }

var fileDescriptor_20ce1f5884b7e047 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x71, 0x52, 0x52, 0x72, 0x5b, 0xa2, 0xc8, 0x03, 0xf2, 0x46, 0xd4, 0xc9, 0x42, 0x28,
	0x43, 0x79, 0x02, 0x8a, 0x60, 0x82, 0x14, 0xb9, 0x9d, 0x58, 0xa2, 0x24, 0xbe, 0x52, 0x2d, 0x61,
	0x3b, 0x72, 0x0d, 0x0f, 0xc9, 0x53, 0xa1, 0x38, 0xe1, 0x4f, 0x6c, 0x3e, 0xdf, 0xf1, 0x67, 0xf9,
	0x5e, 0x58, 0xb6, 0xae, 0x31, 0xdd, 0xa1, 0xec, 0x9d, 0xf5, 0x96, 0x46, 0x7d, 0xbb, 0x7a, 0x86,
	0x6c, 0x13, 0xd8, 0xde, 0x21, 0x56, 0x56, 0x22, 0xcd, 0x20, 0x52, 0x92, 0x91, 0x82, 0xf0, 0x54,
	0x44, 0x4a, 0xd2, 0x12, 0xce, 0xba, 0x83, 0x7a, 0x95, 0x0e, 0x0d, 0x8b, 0x8a, 0x98, 0x2f, 0xd6,
	0xb4, 0xec, 0xdb, 0xf2, 0xaf, 0x25, 0xbe, 0xef, 0xac, 0x3e, 0x08, 0x24, 0x63, 0xf9, 0xef, 0x29,
	0x0a, 0x33, 0xd3, 0x68, 0x64, 0x51, 0x20, 0xe1, 0x4c, 0x19, 0xcc, 0x3b, 0xab, 0x35, 0x1a, 0xcf,
	0xe2, 0x80, 0xbf, 0x22, 0xbd, 0x84, 0x85, 0x32, 0xef, 0xca, 0x63, 0xdd, 0x59, 0x89, 0x6c, 0x56,
	0xc4, 0x3c, 0x15, 0x30, 0xa2, 0xbb, 0xe1, 0xa7, 0x17, 0x90, 0xf4, 0x8d, 0x1b, 0xcc, 0xd3, 0x60,
	0x4e, 0x69, 0x10, 0x3b, 0x87, 0x8d, 0xc7, 0xda, 0x2b, 0x8d, 0x2c, 0x29, 0x08, 0x8f, 0x05, 0x8c,
	0x68, 0xaf, 0x34, 0x52, 0x0e, 0xc9, 0xd1, 0x37, 0xfe, 0xed, 0xc8, 0xe6, 0x05, 0xe1, 0xd9, 0x3a,
	0xff, 0x19, 0x68, 0x17, 0xb8, 0x98, 0xfa, 0xab, 0x6b, 0x58, 0xfe, 0xe6, 0xf4, 0x1c, 0xd2, 0xcd,
	0xae, 0xae, 0xb6, 0xe2, 0xe9, 0xf6, 0x31, 0x3f, 0x99, 0xe2, 0x83, 0xd8, 0xbe, 0xdc, 0x57, 0x39,
	0x69, 0x93, 0xb0, 0xd7, 0x9b, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x95, 0x84, 0xe4, 0x67,
	0x01, 0x00, 0x00,
}