package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/component/appctx"
	"github.com/thanhdat1902/restapi/food_deli/component/hasher"
	"github.com/thanhdat1902/restapi/food_deli/component/tokenprovider/jwt"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userbiz"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userstorage"
)

func Login(provider common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var account usermodel.UserLogin
		if err := c.ShouldBind(&account); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := provider.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)
		tokenProvider := jwt.NewTokenJWTProvider(provider.GetSecretKey())
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewLoginBusiness(store, tokenProvider, md5, appctx.NewTokenConfig())
		data, err := biz.Login(c.Request.Context(), &account)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
