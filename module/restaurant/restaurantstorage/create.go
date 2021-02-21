package restaurantstorage

import (
	"encoding/json"
	"log"

	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
)

func (s *store) CreateNewRestaurant(res *restaurantmodel.Restaurant) error {

	db := s.db
	tmp, err := json.Marshal(&res)
	if err != nil {
		return err
	}
	log.Println("FIle create")
	log.Println(string(tmp))

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Create(res).Error; err != nil {
		return err
	}
	return nil
}
