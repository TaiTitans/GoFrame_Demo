package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"GoFrame_Demo/api/user/v1"
	"GoFrame_Demo/internal/model/entity"
	"GoFrame_Demo/internal/service"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	user := &entity.User{
		Id:   uint(req.Id),
		Name: req.Name,
		Age:  req.Age,
	}
	err = service.UserService.Update(ctx, user)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to update user")
	}
	return &v1.UpdateRes{}, nil
}
