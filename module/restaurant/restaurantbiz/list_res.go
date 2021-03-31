package restaurantbiz

import (
	"context"

	"food_deli/common"

	"food_deli/module/restaurant/restaurantmodel"
)

type ListResStore interface {
	ListDataWithCondition(ctx context.Context, condition map[string]interface{}, paging *common.Paging) ([]restaurantmodel.Restaurant, *common.AppError)
}
type listResBiz struct {
	store ListResStore
}

func NewListResBiz(store ListResStore) *listResBiz {
	return &listResBiz{store: store}
}

func (biz *listResBiz) ListRestaurant(ctx context.Context, paging *common.Paging) ([]restaurantmodel.Restaurant, *common.AppError) {
	result, err := biz.store.ListDataWithCondition(ctx, nil, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.Entity, err)
	}
	return result, nil
}
