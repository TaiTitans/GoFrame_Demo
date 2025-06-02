package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "GoFrame_Demo/api/user/v1"
	"GoFrame_Demo/internal/service"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	filter := make(map[string]interface{})
	users, err := service.UserService.GetList(ctx, filter)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to get user list")
	}
	return &v1.GetListRes{
		List: users,
	}, nil
}
