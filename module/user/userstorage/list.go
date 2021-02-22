package userstorage

import (
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
)

func (s *store) ListUsers(paging *common.Paging) ([]usermodel.User, error) {
	db := s.db
	db = db.Table(usermodel.User{}.TableName())
	db = db.Where("status=?", 1)
	var data []usermodel.User
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	paging.Fulfill()
	if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
