package userstorage

import "github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"

func (s *store) Create(user *usermodel.User) error {
	db := s.db
	if err := db.Table(usermodel.User{}.TableName()).Create(user).Error; err != nil {
		return err
	}
	return nil
}
