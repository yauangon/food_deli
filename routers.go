package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/middleware"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restauranttransport/ginrestaurant"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usertransport/ginuser"
)

// SetupHomeRoute : Home router
func SetUpHomeRoute(r *gin.Engine, appCtx common.DBProvider) {
	r.Use(middleware.Recover(appCtx))
	// API list
	routerV1 := r.Group("/v1")
	resRoute := routerV1.Group("/restaurants")
	{
		resRoute.GET("", ginrestaurant.ListRestaurant(appCtx))
		resRoute.POST("", ginrestaurant.CreateRestaurant(appCtx))
		resRoute.DELETE("/:restaurant-id", ginrestaurant.DeleteRestaurant(appCtx))
		resRoute.GET("/:restaurant-id")
	}
	userRoute := routerV1.Group("/users")
	{
		userRoute.POST("", ginuser.CreateUser(appCtx))
		userRoute.DELETE("/:user-id", ginuser.DeleteUser(appCtx))
	}
}

func SetUpAdminRoute(r *gin.Engine, appCtx common.DBProvider) {

}
