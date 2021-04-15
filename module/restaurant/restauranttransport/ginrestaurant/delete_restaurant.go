package ginrestaurant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantbiz"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantstorage"
)

func DeleteRestaurant(provider common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		id, _ := strconv.Atoi(c.Param("restaurant-id"))
		db := provider.GetMainDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store, requester)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(err.StatusCode, err)
			return
		}
		c.JSON(http.StatusNoContent, common.SimpleSuccessResponse(1))
	}
}
