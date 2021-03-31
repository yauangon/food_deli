package restaurantstorage

import (
	"context"

	"food_deli/common"

	"food_deli/module/restaurant/restaurantmodel"
)

func (s *store) ListDataWithCondition(ctx context.Context, condition map[string]interface{}, paging *common.Paging) ([]restaurantmodel.Restaurant, *common.AppError) {
	db := s.db
	db = db.Table(restaurantmodel.Restaurant{}.TableName())
	db = db.Where("status=?", 1)
	if err := db.Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	var data []restaurantmodel.Restaurant
	paging.Fulfill()
	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return data, nil
}
