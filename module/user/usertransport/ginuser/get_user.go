package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userbiz"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userstorage"
)

func GetUserByID(provider common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		uid, _ := common.FromBase58(c.Param("user-id"))

		db := provider.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewGetUserBiz(store)
		user, err := biz.GetUserByID(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			c.JSON(err.StatusCode, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(user))
	}
}
