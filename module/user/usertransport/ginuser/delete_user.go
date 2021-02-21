package ginuser

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userbiz"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userstorage"
)

func DeleteUser(provider common.DBProvider) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("user-id"))
		db := provider.GetMainDBConnection()

		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewDeleteUserBiz(store)
		if err := biz.DeleteUser(id); err != nil {
			c.JSON(http.StatusGone, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": 1,
		})
	}
}