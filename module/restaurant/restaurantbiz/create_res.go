package restaurantbiz

import (
	"context"

	"food_deli/common"

	"food_deli/module/restaurant/restaurantmodel"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, res *restaurantmodel.CreateRestaurant) *common.AppError
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, res *restaurantmodel.CreateRestaurant) *common.AppError {
	if err := biz.store.Create(ctx, res); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.Entity, err)
	}
	return nil
}
