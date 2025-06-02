package service

import (
	"GoFrame_Demo/internal/logic/user"
	"GoFrame_Demo/internal/model/entity"
	"context"
)

type userService struct {
	logic *user.UserLogic
}

func (s *userService) Create(ctx context.Context, user *entity.User) (int64, error) {
	return s.logic.Create(ctx, user)
}

func (s *userService) Update(ctx context.Context, user *entity.User) error {
	return s.logic.Update(ctx, user)
}

func (s *userService) Delete(ctx context.Context, id int64) error {
	return s.logic.Delete(ctx, id)
}

func (s *userService) GetList(ctx context.Context, filter map[string]interface{}) ([]*entity.User, error) {
	return s.logic.GetList(ctx, filter) // truyền filter xuống logic
}

func (s *userService) GetOne(ctx context.Context, id int64) (*entity.User, error) {
	return s.logic.GetOne(ctx, id)
}

var UserService IUser = &userService{
	logic: &user.UserLogic{},
}
