package restaurantstorage

import (
	"context"

	"food_deli/common"

	"food_deli/module/restaurant/restaurantmodel"
)

func (s *store) FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*restaurantmodel.Restaurant, *common.AppError) {
	db := s.db
	var data restaurantmodel.Restaurant
	if err := db.Table(data.TableName()).Where(condition).First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
