package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/component/hasher"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userbiz"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userstorage"
)

func CreateUser(provider common.DBProvider) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user usermodel.UserCreate

		db := provider.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewCreateUserBiz(store, md5)

		if err := c.ShouldBind(&user); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := biz.Create(c.Request.Context(), &user); err != nil {
			panic(common.ErrCannotCreateEntity(usermodel.EntityName, err))
		}
		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(1))
	}
}
