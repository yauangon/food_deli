package ginuser

import (
	"net/http"

	"food_deli/common"

	"food_deli/module/user/userbiz"

	"food_deli/module/user/userstorage"

	"github.com/gin-gonic/gin"
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
			c.JSON(err.StatusCode, err)
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}
