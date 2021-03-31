package ginrestaurant

import (
	"net/http"

	"food_deli/common"

	"food_deli/module/restaurant/restaurantmodel"

	"food_deli/module/restaurant/restaurantbiz"

	"food_deli/module/restaurant/restaurantstorage"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(provider common.DBProvider) func(c *gin.Context) {
	return func(c *gin.Context) {

		var res restaurantmodel.CreateRestaurant
		if err := c.ShouldBind(&res); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := provider.GetMainDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &res); err != nil {
			c.JSON(err.StatusCode, err)
			return
		}
		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(1))
	}
}
