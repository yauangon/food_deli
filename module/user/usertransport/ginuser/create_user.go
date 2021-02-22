package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userbiz"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userstorage"
)

func CreateUser(provider common.DBProvider) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user usermodel.User

		db := provider.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewCreateUserBiz(store)

		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusGone, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := biz.Create(&user); err != nil {
			c.JSON(http.StatusGone, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(1))
	}
}
