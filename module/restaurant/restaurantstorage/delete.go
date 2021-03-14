package restaurantstorage

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
)

func (s *store) Delete(ctx context.Context, id int) *common.AppError {
	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id=?", id).Updates(map[string]interface{}{"Status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
