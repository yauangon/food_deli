package main

import (
	"food_deli/common"

	"food_deli/module/restaurant/restauranttransport/ginrestaurant"

	"food_deli/module/user/usertransport/ginuser"

	"food_deli/middleware"

	"github.com/gin-gonic/gin"
)

// SetupHomeRoute : Home router
func SetUpHomeRoute(r *gin.Engine, appCtx common.DBProvider) {
	// Apply recover middleware
	r.Use(middleware.Recover(appCtx))
	// API list
	routerV1 := r.Group("/v1")
	resRoute := routerV1.Group("/restaurants")
	{
		resRoute.GET("", ginrestaurant.ListRestaurant(appCtx))
		resRoute.POST("", ginrestaurant.CreateRestaurant(appCtx))
		resRoute.DELETE("/:restaurant-id", ginrestaurant.DeleteRestaurant(appCtx))
		resRoute.GET("/:restaurant-id", ginrestaurant.GetRestaurantByID(appCtx))
	}
	userRoute := routerV1.Group("/users")
	{
		userRoute.POST("", ginuser.CreateUser(appCtx))
		userRoute.GET("/:user-id", ginuser.GetUserByID(appCtx))
		userRoute.DELETE("/:user-id", ginuser.DeleteUser(appCtx))
		userRoute.GET("", ginuser.ListUsers(appCtx))
	}
}

func SetUpAdminRoute(r *gin.Engine, appCtx common.DBProvider) {

}
