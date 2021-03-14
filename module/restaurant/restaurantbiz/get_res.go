package restaurantbiz

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
)

type GetRestaurantStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*restaurantmodel.Restaurant, *common.AppError)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetReataurant(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurantById(ctx context.Context, id int) (*restaurantmodel.Restaurant, *common.AppError) {
	res, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantmodel.Entity, err)
	}
	return res, nil
}
