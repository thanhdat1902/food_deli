package userbiz

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
)

type GetUserStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type getUserBiz struct {
	store GetUserStore
}

func NewGetUserBiz(store GetUserStore) *getUserBiz {
	return &getUserBiz{store: store}
}

func (biz getUserBiz) GetUserByID(ctx context.Context, id int) (*usermodel.User, *common.AppError) {
	user, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}
	return user, nil
}
