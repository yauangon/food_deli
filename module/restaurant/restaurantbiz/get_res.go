package restaurantbiz

import (
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
)

type GetRestaurantStore interface {
	FindDataWithCondition(condition map[string]interface{}) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetReataurant(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurantById(id int) (*restaurantmodel.Restaurant, error) {
	res, err := biz.store.FindDataWithCondition(map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(restaurantmodel.Entity, err)
	}
	return res, nil
}
