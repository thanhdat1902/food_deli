package userstorage

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
)

func (s *store) FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())
	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}
	var data usermodel.User

	if err := db.Where(condition).First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
