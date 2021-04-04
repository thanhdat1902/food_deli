package userbiz

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
)

// DeleteUserStore delete action biz
type DeleteUserStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, *common.AppError)
	Delete(ctx context.Context, id int) *common.AppError
}

type deleteUserBiz struct {
	store DeleteUserStore
}

// NewDeleteUserBiz biz main action
func NewDeleteUserBiz(store DeleteUserStore) *deleteUserBiz {
	return &deleteUserBiz{store: store}
}

func (biz *deleteUserBiz) DeleteUser(ctx context.Context, userID int) *common.AppError {
	user, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": userID})
	if err != nil {
		return common.ErrEntityNotFound(usermodel.EntityName, err)
	}
	if user.Status == 0 {
		return common.ErrDeletedBefore(usermodel.EntityName)
	}
	if err := biz.store.Delete(ctx, userID); err != nil {
		return common.ErrCannotDeleteEntity(usermodel.EntityName, err)
	}
	return nil
}
