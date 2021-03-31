package ginrestaurant

import (
	"net/http"
	"strconv"

	"food_deli/common"

	"food_deli/module/restaurant/restaurantbiz"
	"food_deli/module/restaurant/restaurantmodel"
	"food_deli/module/restaurant/restaurantstorage"

	"github.com/gin-gonic/gin"
)

func UpdateRestaurant(provider common.DBProvider) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data restaurantmodel.Restaurant
		id, _ := strconv.Atoi(c.Param("restaurant-id"))
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		data.ID = id

		db := provider.GetMainDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)
		if err := biz.UpdateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(err.StatusCode, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(1))
	}
}
