package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/middleware"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restauranttransport/ginrestaurant"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usertransport/ginuser"
)

// SetupHomeRoute : Home router
func SetUpHomeRoute(r *gin.Engine, appCtx common.AppContext) {
	// Apply recover middleware
	r.Use(middleware.Recover(appCtx))
	// API list
	routerV1 := r.Group("/v1")
	routerV1.POST("/register", ginuser.CreateUser(appCtx))
	routerV1.POST("/login", ginuser.Login(appCtx))
	routerV1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.ProfileUser())
	resRoute := routerV1.Group("/restaurants")
	{
		resRoute.GET("", ginrestaurant.ListRestaurant(appCtx))
		resRoute.POST("", middleware.RequiredAuth(appCtx), ginrestaurant.CreateRestaurant(appCtx))
		resRoute.DELETE("/:restaurant-id", middleware.RequiredAuth(appCtx), ginrestaurant.DeleteRestaurant(appCtx))
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

func SetUpAdminRoute(r *gin.Engine, appCtx common.AppContext) {

}
