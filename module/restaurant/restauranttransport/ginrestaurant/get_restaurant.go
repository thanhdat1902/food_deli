package ginrestaurant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantbiz"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantstorage"
)

func GetRestaurantByID(provider common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("restaurant-id"))

		db := provider.GetMainDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewGetReataurant(store)
		res, err := biz.GetRestaurantById(c.Request.Context(), id)
		if err != nil {
			c.JSON(err.StatusCode, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(res))
	}
}
