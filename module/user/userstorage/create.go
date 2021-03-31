package userstorage

import (
	"context"

	"food_deli/common"

	"food_deli/module/user/usermodel"
)

func (s *store) Create(ctx context.Context, user *usermodel.User) *common.AppError {
	db := s.db
	if err := db.Table(usermodel.User{}.TableName()).Create(user).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
