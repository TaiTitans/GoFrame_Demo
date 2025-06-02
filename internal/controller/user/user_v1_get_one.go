package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "GoFrame_Demo/api/user/v1"
	"GoFrame_Demo/internal/service"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	user, err := service.UserService.GetOne(ctx, req.Id)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to get user")
	}
	return &v1.GetOneRes{
		User: user,
	}, nil
}
