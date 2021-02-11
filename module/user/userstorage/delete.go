package userstorage

import "github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"

func (s *store) Delete(id int) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).Where("id=?", id).Updates(map[string]interface{}{"Status": 0}).Error; err != nil {
		return err
	}
	return nil
}
