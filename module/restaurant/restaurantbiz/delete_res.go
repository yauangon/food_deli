package restaurantbiz

import (
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(condition map[string]interface{}) (*restaurantmodel.Restaurant, error)
	Delete(resID int) error
}
type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(resID int) error {
	res, err := biz.store.FindDataWithCondition(map[string]interface{}{"id": resID})
	if err != nil {
		return common.ErrEntityNotFound(restaurantmodel.Entity, err)
	}

	if res.Status == 0 {
		return common.ErrDeletedBefore(restaurantmodel.Entity, err)
	}
	if err := biz.store.Delete(res.ID); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.Entity, err)
	}
	return nil
}
