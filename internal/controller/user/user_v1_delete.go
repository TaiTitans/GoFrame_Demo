package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "GoFrame_Demo/api/user/v1"
	"GoFrame_Demo/internal/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = service.UserService.Delete(ctx, req.Id)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "failed to delete user")
	}
	return &v1.DeleteRes{Success: true}, nil
}
