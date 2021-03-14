package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userbiz"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userstorage"
)

func ListUsers(provider common.DBProvider) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusOK, common.ErrInvalidRequest(err))
			return
		}
		db := provider.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewListUserBiz(store)

		data, err := biz.ListUsers(c.Request.Context(), &paging)
		if err != nil {

			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}
