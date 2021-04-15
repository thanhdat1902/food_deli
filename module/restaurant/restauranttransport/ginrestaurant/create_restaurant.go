package ginrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantbiz"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantstorage"
)

func CreateRestaurant(provider common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		var res restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&res); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		resquester := c.MustGet(common.CurrentUser).(common.Requester)
		res.OwnerID = resquester.GetID()

		db := provider.GetMainDBConnection()
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &res); err != nil {
			c.JSON(err.StatusCode, err)
			return
		}
		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(1))
	}
}
