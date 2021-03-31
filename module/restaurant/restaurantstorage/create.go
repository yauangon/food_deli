package restaurantstorage

import (
	"context"
	"log"

	"food_deli/common"

	"food_deli/module/restaurant/restaurantmodel"
)

func (s *store) Create(ctx context.Context, res *restaurantmodel.CreateRestaurant) *common.AppError {

	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Create(&res).Error; err != nil {
		log.Println("Err at DB")
		return common.ErrDB(err)
	}
	return nil
}
