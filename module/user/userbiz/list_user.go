package userbiz

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
)

type ListUserStore interface {
	ListUsers(ctx context.Context, paging *common.Paging) ([]usermodel.User, error)
}

type listUserBiz struct {
	store ListUserStore
}

func NewListUserBiz(store ListUserStore) *listUserBiz {
	return &listUserBiz{store: store}
}

func (biz *listUserBiz) ListUsers(ctx context.Context, paging *common.Paging) ([]usermodel.User, error) {
	data, err := biz.store.ListUsers(ctx, paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
