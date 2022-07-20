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

type ReqBehaviourAdd struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Target               string   `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Type                 uint32   `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Action               uint32   `protobuf:"varint,5,opt,name=action,proto3" json:"action,omitempty"`
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

type ReqBehaviourList struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Target               string   `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Type                 uint32   `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
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
	List                 []*BehaviourInfo `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
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
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x6e, 0x96, 0x34, 0x6d, 0xcf, 0x36, 0x51, 0x0c, 0x62, 0x26, 0x08, 0x08, 0xbe, 0xea, 0x55,
	0x2a, 0xca, 0x13, 0xac, 0xbb, 0xd8, 0x90, 0xf8, 0x93, 0x27, 0x6e, 0xb8, 0xcb, 0x12, 0x43, 0xc3,
	0x96, 0x38, 0x24, 0xce, 0xc4, 0x1e, 0x80, 0x07, 0xe0, 0xf9, 0x78, 0x19, 0xe4, 0x93, 0xa4, 0x71,
	0xa6, 0x76, 0x12, 0x74, 0x77, 0xe7, 0x9c, 0x7c, 0xfe, 0x7e, 0x8e, 0xdd, 0x82, 0x97, 0x17, 0x52,
	0xc9, 0x79, 0x55, 0x8a, 0x62, 0x7e, 0x21, 0x56, 0xe1, 0x75, 0x22, 0xab, 0x22, 0xc0, 0x21, 0x39,
	0x90, 0xa9, 0x0c, 0xd2, 0x32, 0x0f, 0xf4, 0x57, 0xef, 0xc8, 0x40, 0x46, 0x32, 0x4d, 0x65, 0x56,
	0xc3, 0xd8, 0x1f, 0x0b, 0x0e, 0x97, 0xed, 0xd1, 0xb7, 0xd9, 0x57, 0x49, 0xa6, 0x60, 0x57, 0x49,
	0x4c, 0x2d, 0xdf, 0x9a, 0x4d, 0xb8, 0x2e, 0x09, 0x01, 0x47, 0x1f, 0xa4, 0x7b, 0x38, 0xc2, 0x9a,
	0x3c, 0x01, 0x57, 0x85, 0xc5, 0x37, 0xa1, 0xa8, 0x8d, 0xd3, 0xa6, 0x23, 0x14, 0x46, 0x51, 0x21,
	0x42, 0x25, 0x62, 0xea, 0xf8, 0xd6, 0xcc, 0xe1, 0x6d, 0xab, 0x59, 0xb2, 0x30, 0x15, 0x74, 0x58,
	0xb3, 0xe8, 0x7a, 0x8d, 0x96, 0x05, 0x75, 0x71, 0xdc, 0xb6, 0x1a, 0xad, 0x6e, 0x72, 0x41, 0x47,
	0xbe, 0x35, 0x3b, 0xe4, 0x58, 0x6b, 0xcd, 0x30, 0x52, 0x89, 0xcc, 0xe8, 0x18, 0xa7, 0x4d, 0xa7,
	0x59, 0xaa, 0x3c, 0x46, 0xcd, 0x49, 0xad, 0xd9, 0xb4, 0xec, 0x97, 0x05, 0x0f, 0xb8, 0xf8, 0xb1,
	0x0e, 0x78, 0x1c, 0x77, 0x69, 0xac, 0x8d, 0x69, 0xf6, 0x7a, 0x69, 0x5a, 0x17, 0xb6, 0xe1, 0xc2,
	0x83, 0xb1, 0xcc, 0x45, 0x81, 0xa6, 0x1d, 0x44, 0xaf, 0x7b, 0xc3, 0xe1, 0xd0, 0x74, 0xc8, 0x2e,
	0xe1, 0xa1, 0x69, 0xe3, 0x64, 0x25, 0xa2, 0xcb, 0x7f, 0x32, 0xd2, 0x11, 0xdb, 0xbd, 0xe8, 0xad,
	0x41, 0xa7, 0x33, 0xc8, 0xbe, 0x03, 0x31, 0xc5, 0x3e, 0xe3, 0x2e, 0x76, 0xbc, 0xd6, 0x4e, 0xdf,
	0xe9, 0x05, 0xe3, 0x30, 0x35, 0xb5, 0xde, 0x25, 0xa5, 0xda, 0x75, 0xc1, 0x2c, 0x87, 0x47, 0x5c,
	0xe4, 0x57, 0x37, 0xb7, 0xd6, 0xf5, 0x1a, 0xdc, 0x52, 0x85, 0xaa, 0x2a, 0x91, 0x78, 0x7f, 0xf1,
	0x34, 0x30, 0x5f, 0x78, 0x80, 0x47, 0xce, 0x11, 0xc0, 0x1b, 0xa0, 0xce, 0xbc, 0x0a, 0x63, 0x94,
	0x1c, 0x73, 0x5d, 0x92, 0xc7, 0x30, 0x8c, 0x64, 0x95, 0xa9, 0x46, 0xb0, 0x6e, 0xd8, 0x4f, 0xbd,
	0x31, 0x53, 0x11, 0x73, 0xfc, 0x87, 0xe0, 0x1c, 0x9c, 0xab, 0xa4, 0xd4, 0xec, 0xf6, 0x6c, 0x7f,
	0xf1, 0xac, 0x7f, 0xa0, 0xf7, 0x33, 0xe3, 0x08, 0x5c, 0xfc, 0xb6, 0x61, 0xba, 0x9e, 0x9f, 0x8b,
	0xe2, 0x3a, 0x89, 0x04, 0x59, 0x82, 0x7b, 0x1c, 0xc7, 0x1f, 0x33, 0x41, 0x9e, 0xdf, 0x96, 0xec,
	0x3d, 0x65, 0xef, 0x68, 0x83, 0x23, 0x4d, 0xce, 0x06, 0xe4, 0x03, 0xb8, 0x67, 0x21, 0x72, 0xbc,
	0xdc, 0xce, 0x81, 0x8b, 0xf5, 0x5e, 0x6d, 0x60, 0xe9, 0x43, 0xd8, 0x80, 0x9c, 0xc1, 0xa4, 0x7e,
	0x48, 0x9a, 0xd2, 0xdf, 0x4e, 0x59, 0x83, 0xee, 0x72, 0xf6, 0x09, 0xc6, 0xa7, 0x42, 0x9d, 0xe8,
	0xc5, 0xdf, 0x93, 0xb7, 0xf7, 0x30, 0x3a, 0x15, 0x0a, 0xef, 0xec, 0xc5, 0x76, 0x42, 0xfd, 0xdd,
	0xf3, 0xef, 0xe2, 0xd3, 0x08, 0x36, 0x58, 0x1e, 0x7c, 0x81, 0xee, 0xdf, 0xf2, 0xc2, 0xc5, 0xfa,
	0xcd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0xb8, 0x3c, 0xae, 0x6c, 0x05, 0x00, 0x00,
}
