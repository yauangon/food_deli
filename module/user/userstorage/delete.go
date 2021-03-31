package userstorage

import (
	"context"

	"food_deli/common"

	"food_deli/module/user/usermodel"
)

func (s *store) Delete(ctx context.Context, id int) *common.AppError {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).Where("id=?", id).Updates(map[string]interface{}{"Status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
