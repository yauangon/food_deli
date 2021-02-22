package ginrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantbiz"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantstorage"
)

func CreateRestaurant(provider common.DBProvider) func(c *gin.Context) {
	return func(c *gin.Context) {

		var res restaurantmodel.Restaurant
		if err := c.ShouldBind(&res); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		db := provider.GetMainDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(&res); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(1))
	}
}
