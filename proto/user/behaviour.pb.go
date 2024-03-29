// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/behaviour.proto

package user

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

type BehaviourInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User                 string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Target               string   `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	Created              uint64   `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Creator              string   `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Type                 uint32   `protobuf:"varint,7,opt,name=type,proto3" json:"type,omitempty"`
	Action               uint32   `protobuf:"varint,8,opt,name=action,proto3" json:"action,omitempty"`
	Updated              uint64   `protobuf:"varint,9,opt,name=updated,proto3" json:"updated,omitempty"`
	Scene                string   `protobuf:"bytes,10,opt,name=scene,proto3" json:"scene,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BehaviourInfo) Reset()         { *m = BehaviourInfo{} }
func (m *BehaviourInfo) String() string { return proto.CompactTextString(m) }
func (*BehaviourInfo) ProtoMessage()    {}
func (*BehaviourInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_86b383772d9e07c3, []int{0}
}

func (m *BehaviourInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BehaviourInfo.Unmarshal(m, b)
}
func (m *BehaviourInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BehaviourInfo.Marshal(b, m, deterministic)
}
func (m *BehaviourInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BehaviourInfo.Merge(m, src)
}
func (m *BehaviourInfo) XXX_Size() int {
	return xxx_messageInfo_BehaviourInfo.Size(m)
}
func (m *BehaviourInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BehaviourInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BehaviourInfo proto.InternalMessageInfo

func (m *BehaviourInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *BehaviourInfo) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *BehaviourInfo) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *BehaviourInfo) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *BehaviourInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BehaviourInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *BehaviourInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *BehaviourInfo) GetAction() uint32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *BehaviourInfo) GetUpdated() uint64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *BehaviourInfo) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

type ReqBehaviourAdd struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Target               string   `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Type                 uint32   `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Action               uint32   `protobuf:"varint,5,opt,name=action,proto3" json:"action,omitempty"`
	Scene                string   `protobuf:"bytes,6,opt,name=scene,proto3" json:"scene,omitempty"`
	Entity               string   `protobuf:"bytes,7,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqBehaviourAdd) Reset()         { *m = ReqBehaviourAdd{} }
func (m *ReqBehaviourAdd) String() string { return proto.CompactTextString(m) }
func (*ReqBehaviourAdd) ProtoMessage()    {}
func (*ReqBehaviourAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_86b383772d9e07c3, []int{1}
}

func (m *ReqBehaviourAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqBehaviourAdd.Unmarshal(m, b)
}
func (m *ReqBehaviourAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqBehaviourAdd.Marshal(b, m, deterministic)
}
func (m *ReqBehaviourAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqBehaviourAdd.Merge(m, src)
}
func (m *ReqBehaviourAdd) XXX_Size() int {
	return xxx_messageInfo_ReqBehaviourAdd.Size(m)
}
func (m *ReqBehaviourAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqBehaviourAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqBehaviourAdd proto.InternalMessageInfo

func (m *ReqBehaviourAdd) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ReqBehaviourAdd) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ReqBehaviourAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqBehaviourAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqBehaviourAdd) GetAction() uint32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *ReqBehaviourAdd) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReqBehaviourAdd) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

type ReqBehaviourCheck struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Target               string   `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Action               uint32   `protobuf:"varint,3,opt,name=action,proto3" json:"action,omitempty"`
	Type                 uint32   `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqBehaviourCheck) Reset()         { *m = ReqBehaviourCheck{} }
func (m *ReqBehaviourCheck) String() string { return proto.CompactTextString(m) }
func (*ReqBehaviourCheck) ProtoMessage()    {}
func (*ReqBehaviourCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_86b383772d9e07c3, []int{2}
}

func (m *ReqBehaviourCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqBehaviourCheck.Unmarshal(m, b)
}
func (m *ReqBehaviourCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqBehaviourCheck.Marshal(b, m, deterministic)
}
func (m *ReqBehaviourCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqBehaviourCheck.Merge(m, src)
}
func (m *ReqBehaviourCheck) XXX_Size() int {
	return xxx_messageInfo_ReqBehaviourCheck.Size(m)
}
func (m *ReqBehaviourCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqBehaviourCheck.DiscardUnknown(m)
}

var xxx_messageInfo_ReqBehaviourCheck proto.InternalMessageInfo

func (m *ReqBehaviourCheck) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ReqBehaviourCheck) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ReqBehaviourCheck) GetAction() uint32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *ReqBehaviourCheck) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ReqBehaviourUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User                 string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Target               string   `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	Action               uint32   `protobuf:"varint,4,opt,name=action,proto3" json:"action,omitempty"`
	Entity               string   `protobuf:"bytes,5,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqBehaviourUpdate) Reset()         { *m = ReqBehaviourUpdate{} }
func (m *ReqBehaviourUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqBehaviourUpdate) ProtoMessage()    {}
func (*ReqBehaviourUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_86b383772d9e07c3, []int{3}
}

func (m *ReqBehaviourUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqBehaviourUpdate.Unmarshal(m, b)
}
func (m *ReqBehaviourUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqBehaviourUpdate.Marshal(b, m, deterministic)
}
func (m *ReqBehaviourUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqBehaviourUpdate.Merge(m, src)
}
func (m *ReqBehaviourUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqBehaviourUpdate.Size(m)
}
func (m *ReqBehaviourUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqBehaviourUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqBehaviourUpdate proto.InternalMessageInfo

func (m *ReqBehaviourUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqBehaviourUpdate) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ReqBehaviourUpdate) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ReqBehaviourUpdate) GetAction() uint32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *ReqBehaviourUpdate) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

type ReqBehaviourList struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Target               string   `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Type                 uint32   `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Scene                string   `protobuf:"bytes,7,opt,name=scene,proto3" json:"scene,omitempty"`
	Entity               string   `protobuf:"bytes,8,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqBehaviourList) Reset()         { *m = ReqBehaviourList{} }
func (m *ReqBehaviourList) String() string { return proto.CompactTextString(m) }
func (*ReqBehaviourList) ProtoMessage()    {}
func (*ReqBehaviourList) Descriptor() ([]byte, []int) {
	return fileDescriptor_86b383772d9e07c3, []int{4}
}

func (m *ReqBehaviourList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqBehaviourList.Unmarshal(m, b)
}
func (m *ReqBehaviourList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqBehaviourList.Marshal(b, m, deterministic)
}
func (m *ReqBehaviourList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqBehaviourList.Merge(m, src)
}
func (m *ReqBehaviourList) XXX_Size() int {
	return xxx_messageInfo_ReqBehaviourList.Size(m)
}
func (m *ReqBehaviourList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqBehaviourList.DiscardUnknown(m)
}

var xxx_messageInfo_ReqBehaviourList proto.InternalMessageInfo

func (m *ReqBehaviourList) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ReqBehaviourList) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ReqBehaviourList) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqBehaviourList) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReqBehaviourList) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

type ReplyBehaviourCheck struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Had                  bool         `protobuf:"varint,2,opt,name=had,proto3" json:"had,omitempty"`
	Count                uint32       `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyBehaviourCheck) Reset()         { *m = ReplyBehaviourCheck{} }
func (m *ReplyBehaviourCheck) String() string { return proto.CompactTextString(m) }
func (*ReplyBehaviourCheck) ProtoMessage()    {}
func (*ReplyBehaviourCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_86b383772d9e07c3, []int{5}
}

func (m *ReplyBehaviourCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyBehaviourCheck.Unmarshal(m, b)
}
func (m *ReplyBehaviourCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyBehaviourCheck.Marshal(b, m, deterministic)
}
func (m *ReplyBehaviourCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyBehaviourCheck.Merge(m, src)
}
func (m *ReplyBehaviourCheck) XXX_Size() int {
	return xxx_messageInfo_ReplyBehaviourCheck.Size(m)
}
func (m *ReplyBehaviourCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyBehaviourCheck.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyBehaviourCheck proto.InternalMessageInfo

func (m *ReplyBehaviourCheck) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyBehaviourCheck) GetHad() bool {
	if m != nil {
		return m.Had
	}
	return false
}

func (m *ReplyBehaviourCheck) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ReplyBehaviourList struct {
	Status               *ReplyStatus     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*BehaviourInfo `protobuf:"bytes,11,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReplyBehaviourList) Reset()         { *m = ReplyBehaviourList{} }
func (m *ReplyBehaviourList) String() string { return proto.CompactTextString(m) }
func (*ReplyBehaviourList) ProtoMessage()    {}
func (*ReplyBehaviourList) Descriptor() ([]byte, []int) {
	return fileDescriptor_86b383772d9e07c3, []int{6}
}

func (m *ReplyBehaviourList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyBehaviourList.Unmarshal(m, b)
}
func (m *ReplyBehaviourList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyBehaviourList.Marshal(b, m, deterministic)
}
func (m *ReplyBehaviourList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyBehaviourList.Merge(m, src)
}
func (m *ReplyBehaviourList) XXX_Size() int {
	return xxx_messageInfo_ReplyBehaviourList.Size(m)
}
func (m *ReplyBehaviourList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyBehaviourList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyBehaviourList proto.InternalMessageInfo

func (m *ReplyBehaviourList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyBehaviourList) GetList() []*BehaviourInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*BehaviourInfo)(nil), "omo.msp.user.BehaviourInfo")
	proto.RegisterType((*ReqBehaviourAdd)(nil), "omo.msp.user.ReqBehaviourAdd")
	proto.RegisterType((*ReqBehaviourCheck)(nil), "omo.msp.user.ReqBehaviourCheck")
	proto.RegisterType((*ReqBehaviourUpdate)(nil), "omo.msp.user.ReqBehaviourUpdate")
	proto.RegisterType((*ReqBehaviourList)(nil), "omo.msp.user.ReqBehaviourList")
	proto.RegisterType((*ReplyBehaviourCheck)(nil), "omo.msp.user.ReplyBehaviourCheck")
	proto.RegisterType((*ReplyBehaviourList)(nil), "omo.msp.user.ReplyBehaviourList")
}

func init() {
	proto.RegisterFile("proto/user/behaviour.proto", fileDescriptor_86b383772d9e07c3)
}

var fileDescriptor_86b383772d9e07c3 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xc7, 0xe3, 0xc6, 0x71, 0x92, 0x49, 0xaa, 0x27, 0xcf, 0x82, 0xe8, 0xe2, 0xf2, 0x62, 0x7c,
	0xca, 0x29, 0x11, 0xe1, 0x13, 0x34, 0x95, 0x48, 0x2b, 0xf1, 0x52, 0xb9, 0xe2, 0xc2, 0xcd, 0xb5,
	0x87, 0xc6, 0x6a, 0xe2, 0x75, 0xed, 0x75, 0x45, 0x6e, 0x88, 0xcf, 0xc3, 0x8d, 0x4f, 0xc6, 0x37,
	0x40, 0x3b, 0x76, 0xfc, 0x52, 0x9c, 0x0a, 0x28, 0xb7, 0x99, 0xf1, 0x7f, 0xff, 0xf3, 0x9b, 0x59,
	0x5b, 0x06, 0x33, 0x8a, 0x85, 0x14, 0xd3, 0x34, 0xc1, 0x78, 0x7a, 0x81, 0x4b, 0xf7, 0x26, 0x10,
	0x69, 0x3c, 0xa1, 0x22, 0x1b, 0x8a, 0xb5, 0x98, 0xac, 0x93, 0x68, 0xa2, 0x9e, 0x9a, 0x07, 0x15,
	0xa5, 0x27, 0xd6, 0x6b, 0x11, 0x66, 0x32, 0xfb, 0x87, 0x06, 0xfb, 0xf3, 0xed, 0xd1, 0xd3, 0xf0,
	0x93, 0x60, 0x23, 0x68, 0xa7, 0x81, 0xcf, 0x35, 0x4b, 0x1b, 0xf7, 0x1d, 0x15, 0x32, 0x06, 0xba,
	0x3a, 0xc8, 0xf7, 0xa8, 0x44, 0x31, 0x7b, 0x04, 0x86, 0x74, 0xe3, 0x4b, 0x94, 0xbc, 0x4d, 0xd5,
	0x3c, 0x63, 0x1c, 0xba, 0x5e, 0x8c, 0xae, 0x44, 0x9f, 0xeb, 0x96, 0x36, 0xd6, 0x9d, 0x6d, 0xaa,
	0x5c, 0x42, 0x77, 0x8d, 0xbc, 0x93, 0xb9, 0xa8, 0xb8, 0x50, 0x8b, 0x98, 0x1b, 0x54, 0xde, 0xa6,
	0x4a, 0x2d, 0x37, 0x11, 0xf2, 0xae, 0xa5, 0x8d, 0xf7, 0x1d, 0x8a, 0x55, 0x4f, 0xd7, 0x93, 0x81,
	0x08, 0x79, 0x8f, 0xaa, 0x79, 0xa6, 0x5c, 0xd2, 0xc8, 0xa7, 0x9e, 0xfd, 0xac, 0x67, 0x9e, 0xb2,
	0x87, 0xd0, 0x49, 0x3c, 0x0c, 0x91, 0x03, 0xb9, 0x67, 0x89, 0xfd, 0x5d, 0x83, 0xff, 0x1c, 0xbc,
	0x2e, 0xc6, 0x3e, 0xf2, 0xcb, 0x19, 0xb5, 0xc6, 0x19, 0xf7, 0x6a, 0x33, 0x6e, 0xd9, 0xda, 0x15,
	0x36, 0x13, 0x7a, 0x22, 0xc2, 0x98, 0x46, 0xd1, 0x49, 0x5d, 0xe4, 0x15, 0xee, 0x4e, 0x8d, 0xbb,
	0xa0, 0x33, 0x2a, 0x74, 0x4a, 0x8d, 0xa1, 0x0c, 0xe4, 0x86, 0x66, 0xef, 0x3b, 0x79, 0x66, 0x5f,
	0xc1, 0xff, 0x55, 0xe8, 0xe3, 0x25, 0x7a, 0x57, 0x7f, 0x84, 0x5d, 0x62, 0xb4, 0x6b, 0x18, 0xdb,
	0x71, 0xf4, 0x72, 0x1c, 0xfb, 0xab, 0x06, 0xac, 0xda, 0xed, 0x03, 0x2d, 0xf4, 0x9e, 0xef, 0x46,
	0x09, 0xa0, 0xd7, 0x00, 0xca, 0x89, 0x3b, 0xb5, 0x89, 0xbf, 0x68, 0x30, 0xaa, 0x42, 0xbc, 0x09,
	0x12, 0x79, 0xef, 0x8b, 0x2a, 0x96, 0xde, 0x6d, 0x5e, 0x7a, 0xaf, 0x86, 0x10, 0xc1, 0x03, 0x07,
	0xa3, 0xd5, 0xe6, 0xd6, 0xda, 0x5f, 0x82, 0x91, 0x48, 0x57, 0xa6, 0x09, 0x61, 0x0c, 0x66, 0x8f,
	0x27, 0xd5, 0xaf, 0x6d, 0x42, 0x47, 0xce, 0x49, 0xe0, 0xe4, 0x42, 0xb5, 0xba, 0xa5, 0xeb, 0x13,
	0x60, 0xcf, 0x51, 0xa1, 0x22, 0xf1, 0x44, 0x1a, 0xca, 0x1c, 0x2f, 0x4b, 0xec, 0xcf, 0x6a, 0xf1,
	0xd5, 0x8e, 0x34, 0xf5, 0x5f, 0x34, 0x9c, 0x82, 0xbe, 0x0a, 0x12, 0xc9, 0x07, 0x56, 0x7b, 0x3c,
	0x98, 0x1d, 0xd6, 0x0f, 0xd4, 0x3e, 0x79, 0x87, 0x84, 0xb3, 0x6f, 0x3a, 0x8c, 0x8a, 0xfa, 0x39,
	0xc6, 0x37, 0x81, 0x87, 0x6c, 0x0e, 0xc6, 0x91, 0xef, 0xbf, 0x0f, 0x91, 0x3d, 0xbd, 0xdd, 0xb2,
	0xf6, 0x01, 0x99, 0x07, 0x0d, 0x44, 0xca, 0xdc, 0x6e, 0xb1, 0x77, 0x60, 0x9c, 0xb8, 0xe4, 0xf1,
	0x7c, 0xb7, 0x07, 0x2d, 0xd6, 0x7c, 0xd1, 0xe0, 0x52, 0x97, 0xd8, 0x2d, 0x76, 0x02, 0xfd, 0xec,
	0x7d, 0x54, 0x96, 0xd6, 0x6e, 0xcb, 0x4c, 0x74, 0x17, 0xd9, 0x19, 0xf4, 0x16, 0x28, 0x8f, 0xd5,
	0xe2, 0xff, 0x11, 0xdb, 0x5b, 0xe8, 0x2e, 0x50, 0xd2, 0x9d, 0x3d, 0xdb, 0x6d, 0xa8, 0x9e, 0x9b,
	0xd6, 0x5d, 0x7e, 0x4a, 0x61, 0xb7, 0xd8, 0x29, 0x0c, 0x17, 0x28, 0xd5, 0xcd, 0x06, 0x89, 0x0c,
	0x3c, 0xf6, 0xcb, 0xbd, 0x5f, 0xa7, 0x98, 0xc8, 0x33, 0xf7, 0x12, 0xcd, 0x27, 0x3b, 0x5e, 0x09,
	0x3a, 0x48, 0xb7, 0x30, 0x58, 0xa0, 0x9c, 0x6f, 0x5e, 0x07, 0x2b, 0x89, 0x31, 0x3b, 0x6c, 0x74,
	0xca, 0x1e, 0xfe, 0x0e, 0xda, 0x7c, 0xf8, 0x11, 0xca, 0x9f, 0xca, 0x85, 0x41, 0xf1, 0xab, 0x9f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x49, 0x0d, 0x7e, 0x93, 0x06, 0x00, 0x00,
}
