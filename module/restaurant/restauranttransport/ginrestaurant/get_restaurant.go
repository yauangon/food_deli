package ginrestaurant

import (
	"net/http"
	"strconv"

	"food_deli/common"

	"food_deli/module/restaurant/restaurantbiz"
	"food_deli/module/restaurant/restaurantstorage"

	"github.com/gin-gonic/gin"
)

func GetRestaurantByID(provider common.DBProvider) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("restaurant-id"))

		db := provider.GetMainDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewGetReataurant(store)
		res, err := biz.GetRestaurantById(c.Request.Context(), id)
		if err != nil {
			c.JSON(err.StatusCode, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(res))
	}
}
