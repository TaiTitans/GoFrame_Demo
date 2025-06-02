package user

import (
	"GoFrame_Demo/internal/dao"
	"GoFrame_Demo/internal/model/entity"
	"context"
	"fmt"

	"github.com/bluele/gcache"
	"github.com/gogf/gf/v2/errors/gerror"
)

var UserCache = gcache.New(100).
	LRU().
	Build()

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
	exist, err := dao.User.Ctx(ctx).WherePri(user.Id).One()
	if err != nil {
		return err
	}
	if exist.IsEmpty() {
		return gerror.Newf("user with id %d does not exist", user.Id)
	}
	_, err = dao.User.Ctx(ctx).WherePri(user.Id).Data(user).Update()
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
	ageVal, ageExists := filter["age"]
	if !ageExists {
		return nil, fmt.Errorf("missing 'age' in filter")
	}

	var ageInt int
	switch v := ageVal.(type) {
	case float64:
		ageInt = int(v)
	case int:
		ageInt = v
	default:
		return nil, fmt.Errorf("invalid age type: %T", v)
	}

	cacheKey := fmt.Sprintf("user_list_age_%d", ageInt)

	// Kiểm tra cache
	cachedList, err := UserCache.Get(cacheKey)
	if err == nil && cachedList != nil {
		fmt.Println("Cache hit for", cacheKey)
		return cachedList.([]*entity.User), nil
	}

	// Query DB với filter đầy đủ
	query := dao.User.Ctx(ctx)
	for key, value := range filter {
		query = query.Where(key, value)
	}

	err = query.Scan(&list)
	if err != nil {
		return nil, err
	}

	// Set cache theo key đã sinh từ age
	_ = UserCache.Set(cacheKey, list)

	return list, nil
}

func (l *UserLogic) GetOne(ctx context.Context, id int64) (user *entity.User, err error) {
	user = &entity.User{}
	err = dao.User.Ctx(ctx).WherePri(id).Scan(user) // Sử dụng Scan để lấy dữ liệu vào struct nếu scan ko  tìm thấy sẽ trả về rỗng
	if err != nil {
		return nil, err
	}
	if user.Id == 0 {
		return nil, nil // Không tìm thấy người dùng
	}
	return user, nil
}
