package service

import (
	"GoFrame_Demo/internal/model/entity"
	"context"
)

type IUser interface {
	Create(ctx context.Context, user *entity.User) (id int64, err error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int64) error
	GetOne(ctx context.Context, id int64) (user *entity.User, err error)
	GetList(ctx context.Context, filter map[string]interface{}) (list []*entity.User, err error)
}
