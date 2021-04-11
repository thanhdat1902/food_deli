package userbiz

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/component/tokenprovider"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
)

type LoginStorage interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type TokenConfig interface {
	GetAtExp() int
	GetRtExp() int
}

type loginBusiness struct {
	appCtx        common.DBProvider
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	tkCfg         TokenConfig
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider tokenprovider.Provider,
	hasher Hasher, tkCfg TokenConfig) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		tkCfg:         tkCfg,
	}
}

// 1. Find user, email
// 2. Hash pass from input and compare with pass in db
// 3. Provider: issue JWT token for client
// 3.1. Access token and refresh token
// 4. Return token(s)

func (biz *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	user, err := biz.storeUser.FindDataWithCondition(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	passHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.ID,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.tkCfg.GetAtExp())
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := biz.tokenProvider.Generate(payload, biz.tkCfg.GetRtExp())
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(accessToken, refreshToken)

	return account, nil
}
