package v1

import (
	"GoFrame_Demo/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type Status int

const (
	StatusOK       Status = 0 // User is OK
	StatusDisabled Status = 1 // User is Disabled
)

type CreateReq struct {
	g.Meta `path:"/user" method:"post" tags:"User" summary:"Create user"`
	Name   string `v:"required|length:3,10" dc:"user name"`
	Age    uint   `v:"required|between:18,200" dc:"user age"`
}

type CreateRes struct {
	Id   int64  `json:"id" dc:"user id"`
	Name string `json:"name" dc:"user name"`
	Age  uint   `json:"age" dc:"user age"`
}

type UpdateReq struct {
	g.Meta `path:"/user/{id}" method:"put" tags:"User" summary:"Update user"`
	Id     int64  `v:"required|min:1" dc:"user id"`
	Name   string `v:"required|length:3,10" dc:"user name"`
	Age    uint   `v:"required|between:18,200" dc:"user age"`
}
type UpdateRes struct{}

type DeleteReq struct {
	g.Meta `path:"/user/{id}" method:"delete" tags:"User" summary:"Delete user"`
	Id     int64 `v:"required" dc:"user id"`
}

type DeleteRes struct {
	Success bool `json:"success" dc:"delete success"`
}
type GetOneReq struct {
	g.Meta `path:"/user/{id}" method:"get" tags:"User" summary:"Get user by id"`
	Id     int64 `v:"required:min:1" dc:"user id"`
}

type GetOneRes struct {
	*entity.User `dc:"user info"`
}

type GetListReq struct {
	g.Meta `path:"/user" method:"get" tags:"User" summary:"Get user list"`
	Age    *uint `v:"between:18,200" dc:"user age"`
	Status int64 `v:"in:0,1" dc:"user status"`
}

type GetListRes struct {
	List []*entity.User `json:"list" dc:"user list"`
}
