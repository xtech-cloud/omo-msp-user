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
	api "github.com/micro/go-micro/v2/api"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	AddOne(ctx context.Context, in *ReqUserAdd, opts ...client.CallOption) (*ReplyUserOne, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyUserOne, error)
	GetByAccount(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyUserOne, error)
	UpdateBase(ctx context.Context, in *ReqUserUpdate, opts ...client.CallOption) (*ReplyInfo, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetList(ctx context.Context, in *ReqUserList, opts ...client.CallOption) (*ReplyUserList, error)
	GetByPage(ctx context.Context, in *RequestPage, opts ...client.CallOption) (*ReplyUserList, error)
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

func (c *userService) GetByAccount(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyUserOne, error) {
	req := c.c.NewRequest(c.name, "UserService.GetByAccount", in)
	out := new(ReplyUserOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UpdateBase(ctx context.Context, in *ReqUserUpdate, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "UserService.UpdateBase", in)
	out := new(ReplyInfo)
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

// Server API for UserService service

type UserServiceHandler interface {
	AddOne(context.Context, *ReqUserAdd, *ReplyUserOne) error
	GetOne(context.Context, *RequestInfo, *ReplyUserOne) error
	GetByAccount(context.Context, *RequestInfo, *ReplyUserOne) error
	UpdateBase(context.Context, *ReqUserUpdate, *ReplyInfo) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetList(context.Context, *ReqUserList, *ReplyUserList) error
	GetByPage(context.Context, *RequestPage, *ReplyUserList) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		AddOne(ctx context.Context, in *ReqUserAdd, out *ReplyUserOne) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyUserOne) error
		GetByAccount(ctx context.Context, in *RequestInfo, out *ReplyUserOne) error
		UpdateBase(ctx context.Context, in *ReqUserUpdate, out *ReplyInfo) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetList(ctx context.Context, in *ReqUserList, out *ReplyUserList) error
		GetByPage(ctx context.Context, in *RequestPage, out *ReplyUserList) error
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

func (h *userServiceHandler) GetByAccount(ctx context.Context, in *RequestInfo, out *ReplyUserOne) error {
	return h.UserServiceHandler.GetByAccount(ctx, in, out)
}

func (h *userServiceHandler) UpdateBase(ctx context.Context, in *ReqUserUpdate, out *ReplyInfo) error {
	return h.UserServiceHandler.UpdateBase(ctx, in, out)
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
