package ginrestaurant

import (
	"net/http"

	"food_deli/common"

	"food_deli/module/restaurant/restaurantbiz"
	"food_deli/module/restaurant/restaurantstorage"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(provider common.DBProvider) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		db := provider.GetMainDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewListResBiz(store)
		paging.Fulfill()

		data, err := biz.ListRestaurant(c.Request.Context(), &paging)
		if err != nil {
			c.JSON(err.StatusCode, err)
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}
