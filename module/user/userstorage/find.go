package userstorage

import (
	"context"

	"food_deli/common"

	"food_deli/module/user/usermodel"
)

func (s *store) FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*usermodel.User, *common.AppError) {
	db := s.db

	var data usermodel.User

	if err := db.Table(data.TableName()).Where(condition).First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
