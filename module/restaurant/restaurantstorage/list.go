package restaurantstorage

import (
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
)

func (s *store) ListDataWithCondition(condition map[string]interface{}, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	db := s.db
	db = db.Table(restaurantmodel.Restaurant{}.TableName())
	db = db.Where("status=?", 1)
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	var data []restaurantmodel.Restaurant
	paging.Fulfill()
	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
