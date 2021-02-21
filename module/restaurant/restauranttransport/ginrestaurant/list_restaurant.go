package ginrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantbiz"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantstorage"
)

func ListRestaurant(provider common.DBProvider) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		db := provider.GetMainDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.ListResBiz(store)

		data, err := biz.ListRestaurant(&paging)
		if err != nil {
			c.JSON(http.StatusGone, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}
