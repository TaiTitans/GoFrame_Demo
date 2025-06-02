package user

import (
	"context"

	v1 "GoFrame_Demo/api/user/v1"
	"GoFrame_Demo/internal/model/entity"
	"GoFrame_Demo/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	user := &entity.User{
		Name: req.Name,
		Age:  req.Age,
	}
	id, err := service.UserService.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{Id: id, Name: user.Name, Age: user.Age}, nil
}
