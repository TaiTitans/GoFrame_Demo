package user

import (
	"GoFrame_Demo/internal/dao"
	"GoFrame_Demo/internal/model/entity"
	"context"
)

type UserLogic struct{}

func (l *UserLogic) Create(ctx context.Context, user *entity.User) (id int64, err error) {
	if user.Status == 0 {
		user.Status = 1 // Default status
	}
	result, err := dao.User.Ctx(ctx).Data(user).Insert()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (l *UserLogic) Update(ctx context.Context, user *entity.User) error {
	_, err := dao.User.Ctx(ctx).Data(user).WherePri(user.Id).Update()
	if err != nil {
		return err
	}
	return nil
}
func (l *UserLogic) Delete(ctx context.Context, id int64) error {
	_, err := dao.User.Ctx(ctx).WherePri(id).Delete()
	if err != nil {
		return err
	}
	return nil
}
func (l *UserLogic) GetList(ctx context.Context, filter map[string]interface{}) (list []*entity.User, err error) {
	query := dao.User.Ctx(ctx)
	for key, value := range filter {
		query = query.Where(key, value)
	}
	err = query.Scan(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
func (l *UserLogic) GetOne(ctx context.Context, id int64) (user *entity.User, err error) {
	user = &entity.User{}
	err = dao.User.Ctx(ctx).WherePri(id).Scan(user)
	if err != nil {
		return nil, err
	}
	if user.Id == 0 {
		return nil, nil // Không tìm thấy người dùng
	}
	return user, nil
}
