// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserService interface {
	AddOne(ctx context.Context, in *ReqUserAdd, opts ...client.CallOption) (*ReplyUserOne, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyUserOne, error)
	GetByID(ctx context.Context, in *RequestIDInfo, opts ...client.CallOption) (*ReplyUserOne, error)
	GetByPhone(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyUserOne, error)
	GetByEntity(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyUserOne, error)
	GetBySNS(ctx context.Context, in *ReqUserBy, opts ...client.CallOption) (*ReplyUserOne, error)
	UpdateBase(ctx context.Context, in *ReqUserUpdate, opts ...client.CallOption) (*ReplyUserOne, error)
	UpdateEntity(ctx context.Context, in *ReqUserEntity, opts ...client.CallOption) (*ReplyUserOne, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetList(ctx context.Context, in *ReqUserList, opts ...client.CallOption) (*ReplyUserList, error)
	GetByPage(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyUserList, error)
	GetByKey(ctx context.Context, in *ReqUserSearch, opts ...client.CallOption) (*ReplyUserList, error)
	UpdateSNS(ctx context.Context, in *ReqUserSNS, opts ...client.CallOption) (*ReplyUserOne, error)
	UpdateTags(ctx context.Context, in *ReqUserTags, opts ...client.CallOption) (*ReplyUserOne, error)
	UpdatePhone(ctx context.Context, in *ReqUserPhone, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateFollows(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, opts ...client.CallOption) (*ReplyInfo, error)
	GetStatistic(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyStatistic, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) AddOne(ctx context.Context, in *ReqUserAdd, opts ...client.CallOption) (*ReplyUserOne, error) {
	req := c.c.NewRequest(c.name, "UserService.AddOne", in)
	out := new(ReplyUserOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyUserOne, error) {
	req := c.c.NewRequest(c.name, "UserService.GetOne", in)
	out := new(ReplyUserOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetByID(ctx context.Context, in *RequestIDInfo, opts ...client.CallOption) (*ReplyUserOne, error) {
	req := c.c.NewRequest(c.name, "UserService.GetByID", in)
	out := new(ReplyUserOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetByPhone(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyUserOne, error) {
	req := c.c.NewRequest(c.name, "UserService.GetByPhone", in)
	out := new(ReplyUserOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetByEntity(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyUserOne, error) {
	req := c.c.NewRequest(c.name, "UserService.GetByEntity", in)
	out := new(ReplyUserOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetBySNS(ctx context.Context, in *ReqUserBy, opts ...client.CallOption) (*ReplyUserOne, error) {
	req := c.c.NewRequest(c.name, "UserService.GetBySNS", in)
	out := new(ReplyUserOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateBase(ctx context.Context, in *ReqUserUpdate, opts ...client.CallOption) (*ReplyUserOne, error) {
	req := c.c.NewRequest(c.name, "UserService.UpdateBase", in)
	out := new(ReplyUserOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateEntity(ctx context.Context, in *ReqUserEntity, opts ...client.CallOption) (*ReplyUserOne, error) {
	req := c.c.NewRequest(c.name, "UserService.UpdateEntity", in)
	out := new(ReplyUserOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "UserService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetList(ctx context.Context, in *ReqUserList, opts ...client.CallOption) (*ReplyUserList, error) {
	req := c.c.NewRequest(c.name, "UserService.GetList", in)
	out := new(ReplyUserList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetByPage(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyUserList, error) {
	req := c.c.NewRequest(c.name, "UserService.GetByPage", in)
	out := new(ReplyUserList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetByKey(ctx context.Context, in *ReqUserSearch, opts ...client.CallOption) (*ReplyUserList, error) {
	req := c.c.NewRequest(c.name, "UserService.GetByKey", in)
	out := new(ReplyUserList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateSNS(ctx context.Context, in *ReqUserSNS, opts ...client.CallOption) (*ReplyUserOne, error) {
	req := c.c.NewRequest(c.name, "UserService.UpdateSNS", in)
	out := new(ReplyUserOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateTags(ctx context.Context, in *ReqUserTags, opts ...client.CallOption) (*ReplyUserOne, error) {
	req := c.c.NewRequest(c.name, "UserService.UpdateTags", in)
	out := new(ReplyUserOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdatePhone(ctx context.Context, in *ReqUserPhone, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "UserService.UpdatePhone", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateFollows(ctx context.Context, in *RequestList, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "UserService.UpdateFollows", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "UserService.UpdateByFilter", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetStatistic(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyStatistic, error) {
	req := c.c.NewRequest(c.name, "UserService.GetStatistic", in)
	out := new(ReplyStatistic)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	AddOne(context.Context, *ReqUserAdd, *ReplyUserOne) error
	GetOne(context.Context, *RequestInfo, *ReplyUserOne) error
	GetByID(context.Context, *RequestIDInfo, *ReplyUserOne) error
	GetByPhone(context.Context, *RequestInfo, *ReplyUserOne) error
	GetByEntity(context.Context, *RequestInfo, *ReplyUserOne) error
	GetBySNS(context.Context, *ReqUserBy, *ReplyUserOne) error
	UpdateBase(context.Context, *ReqUserUpdate, *ReplyUserOne) error
	UpdateEntity(context.Context, *ReqUserEntity, *ReplyUserOne) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetList(context.Context, *ReqUserList, *ReplyUserList) error
	GetByPage(context.Context, *RequestPage, *ReplyUserList) error
	GetByKey(context.Context, *ReqUserSearch, *ReplyUserList) error
	UpdateSNS(context.Context, *ReqUserSNS, *ReplyUserOne) error
	UpdateTags(context.Context, *ReqUserTags, *ReplyUserOne) error
	UpdatePhone(context.Context, *ReqUserPhone, *ReplyInfo) error
	UpdateFollows(context.Context, *RequestList, *ReplyInfo) error
	UpdateByFilter(context.Context, *ReqUpdateFilter, *ReplyInfo) error
	GetStatistic(context.Context, *RequestPage, *ReplyStatistic) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		AddOne(ctx context.Context, in *ReqUserAdd, out *ReplyUserOne) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyUserOne) error
		GetByID(ctx context.Context, in *RequestIDInfo, out *ReplyUserOne) error
		GetByPhone(ctx context.Context, in *RequestInfo, out *ReplyUserOne) error
		GetByEntity(ctx context.Context, in *RequestInfo, out *ReplyUserOne) error
		GetBySNS(ctx context.Context, in *ReqUserBy, out *ReplyUserOne) error
		UpdateBase(ctx context.Context, in *ReqUserUpdate, out *ReplyUserOne) error
		UpdateEntity(ctx context.Context, in *ReqUserEntity, out *ReplyUserOne) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetList(ctx context.Context, in *ReqUserList, out *ReplyUserList) error
		GetByPage(ctx context.Context, in *RequestPage, out *ReplyUserList) error
		GetByKey(ctx context.Context, in *ReqUserSearch, out *ReplyUserList) error
		UpdateSNS(ctx context.Context, in *ReqUserSNS, out *ReplyUserOne) error
		UpdateTags(ctx context.Context, in *ReqUserTags, out *ReplyUserOne) error
		UpdatePhone(ctx context.Context, in *ReqUserPhone, out *ReplyInfo) error
		UpdateFollows(ctx context.Context, in *RequestList, out *ReplyInfo) error
		UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, out *ReplyInfo) error
		GetStatistic(ctx context.Context, in *RequestPage, out *ReplyStatistic) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) AddOne(ctx context.Context, in *ReqUserAdd, out *ReplyUserOne) error {
	return h.UserServiceHandler.AddOne(ctx, in, out)
}

func (h *userServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyUserOne) error {
	return h.UserServiceHandler.GetOne(ctx, in, out)
}

func (h *userServiceHandler) GetByID(ctx context.Context, in *RequestIDInfo, out *ReplyUserOne) error {
	return h.UserServiceHandler.GetByID(ctx, in, out)
}

func (h *userServiceHandler) GetByPhone(ctx context.Context, in *RequestInfo, out *ReplyUserOne) error {
	return h.UserServiceHandler.GetByPhone(ctx, in, out)
}

func (h *userServiceHandler) GetByEntity(ctx context.Context, in *RequestInfo, out *ReplyUserOne) error {
	return h.UserServiceHandler.GetByEntity(ctx, in, out)
}

func (h *userServiceHandler) GetBySNS(ctx context.Context, in *ReqUserBy, out *ReplyUserOne) error {
	return h.UserServiceHandler.GetBySNS(ctx, in, out)
}

func (h *userServiceHandler) UpdateBase(ctx context.Context, in *ReqUserUpdate, out *ReplyUserOne) error {
	return h.UserServiceHandler.UpdateBase(ctx, in, out)
}

func (h *userServiceHandler) UpdateEntity(ctx context.Context, in *ReqUserEntity, out *ReplyUserOne) error {
	return h.UserServiceHandler.UpdateEntity(ctx, in, out)
}

func (h *userServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.UserServiceHandler.RemoveOne(ctx, in, out)
}

func (h *userServiceHandler) GetList(ctx context.Context, in *ReqUserList, out *ReplyUserList) error {
	return h.UserServiceHandler.GetList(ctx, in, out)
}

func (h *userServiceHandler) GetByPage(ctx context.Context, in *RequestPage, out *ReplyUserList) error {
	return h.UserServiceHandler.GetByPage(ctx, in, out)
}

func (h *userServiceHandler) GetByKey(ctx context.Context, in *ReqUserSearch, out *ReplyUserList) error {
	return h.UserServiceHandler.GetByKey(ctx, in, out)
}

func (h *userServiceHandler) UpdateSNS(ctx context.Context, in *ReqUserSNS, out *ReplyUserOne) error {
	return h.UserServiceHandler.UpdateSNS(ctx, in, out)
}

func (h *userServiceHandler) UpdateTags(ctx context.Context, in *ReqUserTags, out *ReplyUserOne) error {
	return h.UserServiceHandler.UpdateTags(ctx, in, out)
}

func (h *userServiceHandler) UpdatePhone(ctx context.Context, in *ReqUserPhone, out *ReplyInfo) error {
	return h.UserServiceHandler.UpdatePhone(ctx, in, out)
}

func (h *userServiceHandler) UpdateFollows(ctx context.Context, in *RequestList, out *ReplyInfo) error {
	return h.UserServiceHandler.UpdateFollows(ctx, in, out)
}

func (h *userServiceHandler) UpdateByFilter(ctx context.Context, in *ReqUpdateFilter, out *ReplyInfo) error {
	return h.UserServiceHandler.UpdateByFilter(ctx, in, out)
}

func (h *userServiceHandler) GetStatistic(ctx context.Context, in *RequestPage, out *ReplyStatistic) error {
	return h.UserServiceHandler.GetStatistic(ctx, in, out)
}
