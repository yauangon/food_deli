package ginrestaurant

import (
	"net/http"
	"strconv"

	"food_deli/common"

	"food_deli/module/restaurant/restaurantbiz"

	"food_deli/module/restaurant/restaurantstorage"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(provider common.DBProvider) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("restaurant-id"))
		db := provider.GetMainDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)
		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(err.StatusCode, err)
			return
		}
		c.JSON(http.StatusNoContent, common.SimpleSuccessResponse(1))
	}
}
