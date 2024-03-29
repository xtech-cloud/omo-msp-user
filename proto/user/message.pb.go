// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/message.proto

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

type MessageInfo struct {
	Type                 uint32   `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Status               uint32   `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Created              uint64   `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              uint64   `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Stamp                uint64   `protobuf:"varint,5,opt,name=stamp,proto3" json:"stamp,omitempty"`
	Uid                  string   `protobuf:"bytes,6,opt,name=uid,proto3" json:"uid,omitempty"`
	Creator              string   `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
	Name                 string   `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	User                 string   `protobuf:"bytes,9,opt,name=user,proto3" json:"user,omitempty"`
	Quote                string   `protobuf:"bytes,10,opt,name=quote,proto3" json:"quote,omitempty"`
	Owner                string   `protobuf:"bytes,11,opt,name=owner,proto3" json:"owner,omitempty"`
	Targets              []string `protobuf:"bytes,12,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageInfo) Reset()         { *m = MessageInfo{} }
func (m *MessageInfo) String() string { return proto.CompactTextString(m) }
func (*MessageInfo) ProtoMessage()    {}
func (*MessageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_600b1562ab433ea1, []int{0}
}

func (m *MessageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageInfo.Unmarshal(m, b)
}
func (m *MessageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageInfo.Marshal(b, m, deterministic)
}
func (m *MessageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageInfo.Merge(m, src)
}
func (m *MessageInfo) XXX_Size() int {
	return xxx_messageInfo_MessageInfo.Size(m)
}
func (m *MessageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MessageInfo proto.InternalMessageInfo

func (m *MessageInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MessageInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *MessageInfo) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *MessageInfo) GetUpdated() uint64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *MessageInfo) GetStamp() uint64 {
	if m != nil {
		return m.Stamp
	}
	return 0
}

func (m *MessageInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *MessageInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MessageInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MessageInfo) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *MessageInfo) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *MessageInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MessageInfo) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReqMessageAdd struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Quote                string   `protobuf:"bytes,2,opt,name=quote,proto3" json:"quote,omitempty"`
	Type                 uint32   `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Stamp                uint64   `protobuf:"varint,5,opt,name=stamp,proto3" json:"stamp,omitempty"`
	Owner                string   `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	Targets              []string `protobuf:"bytes,11,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqMessageAdd) Reset()         { *m = ReqMessageAdd{} }
func (m *ReqMessageAdd) String() string { return proto.CompactTextString(m) }
func (*ReqMessageAdd) ProtoMessage()    {}
func (*ReqMessageAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_600b1562ab433ea1, []int{1}
}

func (m *ReqMessageAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqMessageAdd.Unmarshal(m, b)
}
func (m *ReqMessageAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqMessageAdd.Marshal(b, m, deterministic)
}
func (m *ReqMessageAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqMessageAdd.Merge(m, src)
}
func (m *ReqMessageAdd) XXX_Size() int {
	return xxx_messageInfo_ReqMessageAdd.Size(m)
}
func (m *ReqMessageAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqMessageAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqMessageAdd proto.InternalMessageInfo

func (m *ReqMessageAdd) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ReqMessageAdd) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *ReqMessageAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqMessageAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqMessageAdd) GetStamp() uint64 {
	if m != nil {
		return m.Stamp
	}
	return 0
}

func (m *ReqMessageAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqMessageAdd) GetTargets() []string {
	if m != nil {
		return m.Targets
	}
	return nil
}

type ReplyMessageList struct {
	Status               *ReplyStatus   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint32         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages                uint32         `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	List                 []*MessageInfo `protobuf:"bytes,11,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReplyMessageList) Reset()         { *m = ReplyMessageList{} }
func (m *ReplyMessageList) String() string { return proto.CompactTextString(m) }
func (*ReplyMessageList) ProtoMessage()    {}
func (*ReplyMessageList) Descriptor() ([]byte, []int) {
	return fileDescriptor_600b1562ab433ea1, []int{2}
}

func (m *ReplyMessageList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyMessageList.Unmarshal(m, b)
}
func (m *ReplyMessageList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyMessageList.Marshal(b, m, deterministic)
}
func (m *ReplyMessageList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyMessageList.Merge(m, src)
}
func (m *ReplyMessageList) XXX_Size() int {
	return xxx_messageInfo_ReplyMessageList.Size(m)
}
func (m *ReplyMessageList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyMessageList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyMessageList proto.InternalMessageInfo

func (m *ReplyMessageList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyMessageList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyMessageList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyMessageList) GetList() []*MessageInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageInfo)(nil), "omo.msp.user.MessageInfo")
	proto.RegisterType((*ReqMessageAdd)(nil), "omo.msp.user.ReqMessageAdd")
	proto.RegisterType((*ReplyMessageList)(nil), "omo.msp.user.ReplyMessageList")
}

func init() {
	proto.RegisterFile("proto/user/message.proto", fileDescriptor_600b1562ab433ea1)
}

var fileDescriptor_600b1562ab433ea1 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x8d, 0x13, 0xc7, 0x69, 0x27, 0x69, 0x55, 0xad, 0xaa, 0xaf, 0xfb, 0x85, 0x1f, 0x45, 0xbe,
	0xca, 0x0d, 0xae, 0x28, 0x2f, 0x40, 0x7b, 0xd1, 0xaa, 0x08, 0x04, 0x72, 0xc5, 0x0d, 0x77, 0x4b,
	0x3c, 0x44, 0x96, 0x62, 0xaf, 0xb3, 0x3b, 0x06, 0xe5, 0x61, 0x90, 0x78, 0x03, 0xc4, 0x1b, 0xa2,
	0x9d, 0x75, 0x8c, 0x21, 0x49, 0xef, 0xe6, 0x9c, 0x19, 0xcf, 0x9c, 0x39, 0x3b, 0x06, 0x59, 0x19,
	0x4d, 0xfa, 0xb2, 0xb6, 0x68, 0x2e, 0x0b, 0xb4, 0x56, 0x2d, 0x31, 0x61, 0x4a, 0x4c, 0x74, 0xa1,
	0x93, 0xc2, 0x56, 0x89, 0xcb, 0x4d, 0x2f, 0x3a, 0x75, 0x0b, 0x5d, 0x14, 0xba, 0xf4, 0x65, 0xf1,
	0xf7, 0x3e, 0x8c, 0xdf, 0xf9, 0x0f, 0xef, 0xcb, 0x2f, 0x5a, 0x08, 0x08, 0x69, 0x53, 0xa1, 0x0c,
	0x66, 0xc1, 0xfc, 0x24, 0xe5, 0x58, 0xfc, 0x07, 0x91, 0x25, 0x45, 0xb5, 0x95, 0x7d, 0x66, 0x1b,
	0x24, 0x24, 0x8c, 0x16, 0x06, 0x15, 0x61, 0x26, 0x07, 0xb3, 0x60, 0x1e, 0xa6, 0x5b, 0xe8, 0x32,
	0x75, 0x95, 0x71, 0x26, 0xf4, 0x99, 0x06, 0x8a, 0x73, 0x18, 0x5a, 0x52, 0x45, 0x25, 0x87, 0xcc,
	0x7b, 0x20, 0xce, 0x60, 0x50, 0xe7, 0x99, 0x8c, 0x66, 0xc1, 0xfc, 0x38, 0x75, 0x61, 0xdb, 0x5b,
	0x1b, 0x39, 0x62, 0x76, 0x0b, 0x9d, 0xc2, 0x52, 0x15, 0x28, 0x8f, 0x98, 0xe6, 0xd8, 0x71, 0x6e,
	0x35, 0x79, 0xec, 0x39, 0x17, 0xbb, 0x49, 0xeb, 0x5a, 0x13, 0x4a, 0x60, 0xd2, 0x03, 0xc7, 0xea,
	0x6f, 0x25, 0x1a, 0x39, 0xf6, 0x2c, 0x03, 0x37, 0x8d, 0x94, 0x59, 0x22, 0x59, 0x39, 0x99, 0x0d,
	0xdc, 0xb4, 0x06, 0xc6, 0x3f, 0x03, 0x38, 0x49, 0x71, 0xdd, 0x58, 0x74, 0x9d, 0x65, 0xed, 0xac,
	0x60, 0xdf, 0xac, 0x7e, 0x77, 0xd6, 0xd6, 0xcb, 0x41, 0xc7, 0xcb, 0x29, 0x1c, 0xe9, 0x0a, 0x0d,
	0x2f, 0x16, 0x72, 0x71, 0x8b, 0x0f, 0x78, 0xd3, 0x2a, 0x8e, 0x0e, 0x28, 0x1e, 0xff, 0xad, 0xf8,
	0x47, 0x00, 0x67, 0x29, 0x56, 0xab, 0x4d, 0xa3, 0xf9, 0x6d, 0x6e, 0x49, 0xbc, 0x6c, 0x9f, 0xd0,
	0xc9, 0x1e, 0x5f, 0xfd, 0x9f, 0x74, 0xcf, 0x23, 0xe1, 0xfa, 0x07, 0x2e, 0x68, 0x5f, 0xf7, 0x1c,
	0x86, 0xa4, 0x49, 0xad, 0x9a, 0x47, 0xf7, 0xc0, 0xb1, 0x95, 0x5a, 0xa2, 0x6d, 0x96, 0xf2, 0x40,
	0xbc, 0x80, 0x70, 0x95, 0x5b, 0x62, 0x29, 0x3b, 0xcd, 0x3b, 0xe7, 0x95, 0x72, 0xd9, 0xd5, 0xaf,
	0x3e, 0x9c, 0x36, 0xec, 0x03, 0x9a, 0xaf, 0xf9, 0x02, 0xc5, 0x6b, 0x88, 0xae, 0xb3, 0xec, 0x7d,
	0x89, 0xe2, 0xc9, 0xbf, 0xd2, 0x3a, 0xe6, 0x4f, 0x2f, 0xf6, 0xe8, 0x76, 0x8d, 0xe3, 0x9e, 0xb8,
	0x87, 0xc9, 0x1d, 0x92, 0x5b, 0x22, 0xb7, 0x94, 0x2f, 0xc4, 0xce, 0x8a, 0xeb, 0x1a, 0x2d, 0x7d,
	0x50, 0x4b, 0x9c, 0x3e, 0x3d, 0xb0, 0x3d, 0x7f, 0x18, 0xf7, 0xc4, 0x2d, 0x8c, 0xee, 0x90, 0xd8,
	0xb8, 0x47, 0xba, 0x3c, 0xdf, 0xd3, 0xa5, 0xe3, 0x79, 0xdc, 0x13, 0x6f, 0xe0, 0xf4, 0x23, 0xdf,
	0xfd, 0xcd, 0xe6, 0x36, 0x5f, 0x11, 0x1a, 0xf1, 0x6c, 0xa7, 0x9d, 0x2f, 0xf0, 0xe9, 0x47, 0xd6,
	0xbb, 0x99, 0x7c, 0x82, 0x3f, 0xff, 0xf0, 0xe7, 0x88, 0xe3, 0x57, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xd3, 0xfb, 0xb6, 0x51, 0x00, 0x04, 0x00, 0x00,
}
