package ginrestaurant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantbiz"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantstorage"
)

func UpdateRestaurant(provider common.DBProvider) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantUpdate
		id, _ := strconv.Atoi(c.Param("restaurant-id"))
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		data.ID = id

		db := provider.GetMainDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)
		if err := biz.UpdateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(err.StatusCode, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(1))
	}
}
