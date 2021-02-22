package restaurantstorage

import (
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
)

func (s *store) CreateNewRestaurant(res *restaurantmodel.Restaurant) error {

	db := s.db
	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Create(res).Error; err != nil {
		return err
	}
	return nil
}
