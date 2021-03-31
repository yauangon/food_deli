package ginuser

import (
	"net/http"
	"strconv"

	"food_deli/common"

	"food_deli/module/user/userbiz"
	"food_deli/module/user/userstorage"

	"github.com/gin-gonic/gin"
)

func DeleteUser(provider common.DBProvider) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("user-id"))
		db := provider.GetMainDBConnection()

		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewDeleteUserBiz(store)
		if err := biz.DeleteUser(c.Request.Context(), id); err != nil {
			c.JSON(err.StatusCode, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(1))
	}
}
