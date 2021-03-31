package userstorage

import (
	"context"

	"food_deli/common"

	"food_deli/module/user/usermodel"
)

func (s *store) ListUsers(ctx context.Context, paging *common.Paging) ([]usermodel.User, *common.AppError) {
	db := s.db
	db = db.Table(usermodel.User{}.TableName())
	db = db.Where("status=?", 1)
	var data []usermodel.User
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	paging.Fulfill()
	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return data, nil
}
