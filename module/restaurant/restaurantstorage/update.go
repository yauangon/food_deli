package restaurantstorage

import (
	"context"

	"food_deli/common"

	"food_deli/module/restaurant/restaurantmodel"
)

func (s *store) Update(ctx context.Context, res *restaurantmodel.Restaurant) *common.AppError {
	db := s.db
	db = db.Table(restaurantmodel.Restaurant{}.TableName())
	db = db.Where("id=?", res.ID)
	if err := db.Updates(res).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
