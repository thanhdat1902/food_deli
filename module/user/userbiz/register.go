package userbiz

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
)

type CreateUserStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	Create(ctx context.Context, user *usermodel.UserCreate) error
}
type Hasher interface {
	Hash(data string) string
}
type createUserBiz struct {
	store  CreateUserStore
	hasher Hasher
}

func NewCreateUserBiz(store CreateUserStore, hasher Hasher) *createUserBiz {
	return &createUserBiz{store: store, hasher: hasher}
}

func (biz *createUserBiz) Create(ctx context.Context, user *usermodel.UserCreate) error {
	data, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"email": user.Email})
	if data != nil {
		return common.ErrEntityExisted(usermodel.EntityName, err)
	}
	salt := common.GenSalt(50)
	user.Salt = salt
	user.Role = "USER"
	hashed := biz.hasher.Hash(user.Password + user.Salt)
	user.Password = hashed
	if err := biz.store.Create(ctx, user); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
