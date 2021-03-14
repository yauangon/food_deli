package restaurantbiz

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*restaurantmodel.Restaurant, *common.AppError)
	Delete(ctx context.Context, resID int) *common.AppError
}
type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, resID int) *common.AppError {
	res, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": resID})
	if err != nil {
		return common.ErrEntityNotFound(restaurantmodel.Entity, err)
	}

	if res.Status == 0 {
		return common.ErrDeletedBefore(restaurantmodel.Entity)
	}
	if err := biz.store.Delete(ctx, res.ID); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.Entity, err)
	}
	return nil
}
