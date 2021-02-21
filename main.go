package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/component/appctx"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restauranttransport/ginrestaurant"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usertransport/ginuser"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Location struct
// type Location struct {
// 	LocationID int `json:"id,omitempty" form:"id" gorm:"id"`

// 	City   string `json:"city" form:"city" gorm:"city"`
// 	Ward   string `json:"ward" form:"ward" gorm:"ward"`
// 	Street string `json:"street" form:"street" gorm:"street"`

// 	CreatedAt time.Time `json:"created_at" form:"created_at" gorm:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at" form:"updated_at" gorm:"updated_at"`
// 	Status    int       `json:"status" form:"status" gorm:"status"`
// }

func main() {

	// Connect to db
	dsn := "admin:admin@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug()
	appCtx := appctx.New(db)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connected to db", db)

	// Router
	router := gin.Default()
	// Test new Res
	routerV1 := router.Group("/v1")
	resRoute := routerV1.Group("/restaurants")
	{
		resRoute.GET("", ginrestaurant.ListRestaurant(appCtx))
		resRoute.POST("", ginrestaurant.CreateRestaurant(appCtx))
	}
	userRoute := routerV1.Group("/users")
	{
		userRoute.DELETE("/:user-id", ginuser.DeleteUser(appCtx))
	}

	router.Run(":8080")
}
