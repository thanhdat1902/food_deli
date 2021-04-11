package ginuser

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userbiz"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userstorage"
)

func DeleteUser(provider common.DBProvider) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("user-id"))
		db := provider.GetMainDBConnection()

		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewDeleteUserBiz(store)
		if err := biz.DeleteUser(c.Request.Context(), id); err != nil {
			panic(common.ErrCannotDeleteEntity(usermodel.EntityName, err))
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(1))
	}
}
