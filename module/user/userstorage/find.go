package userstorage

import "github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"

func (s *store) FindDataWithCondition(condition map[string]interface{}) (*usermodel.User, error) {
	db := s.db

	var data usermodel.User

	if err := db.Table(data.TableName()).Where(condition).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
