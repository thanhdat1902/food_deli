package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userbiz"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
	"github.com/thanhdat1902/restapi/food_deli/module/user/userstorage"
)

func ListUsers(provider common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		db := provider.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)
		biz := userbiz.NewListUserBiz(store)
		paging.Fulfill()

		data, err := biz.ListUsers(c.Request.Context(), &paging)
		if err != nil {
			panic(common.ErrCannotListEntity(usermodel.EntityName, err))
		}
		for i := 0; i < len(data); i++ {
			data[i].Mask()
			if i == 0 {
				paging.PreviousCursor = data[0].FakeID.String()
			}
			if i == paging.Limit-1 {
				paging.NextCursor = data[i].FakeID.String()
			}
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}
