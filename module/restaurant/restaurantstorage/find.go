package restaurantstorage

import "github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"

func (s *store) FindDataWithCondition(condition map[string]interface{}) (*restaurantmodel.Restaurant, error) {
	db := s.db
	var data restaurantmodel.Restaurant
	if err := db.Table(data.TableName()).Where(condition).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
