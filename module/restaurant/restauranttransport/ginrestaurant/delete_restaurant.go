package ginrestaurant

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantbiz"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantstorage"
)

func DeleteRestaurant(provider common.DBProvider) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("restaurant-id"))
		log.Println(id)
		db := provider.GetMainDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)
		if err := biz.DeleteRestaurant(id); err != nil {
			c.JSON(http.StatusGone, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusAccepted, common.SimpleSuccessResponse(1))
	}
}
