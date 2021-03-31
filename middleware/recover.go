package middleware

import (
	"food_deli/common"

	"github.com/gin-gonic/gin"
)

func Recover(sc common.DBProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
				}

				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
			}
		}()

		c.Next()
	}
}
