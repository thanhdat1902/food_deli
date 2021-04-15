package userstorage

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
)

func (s *store) ListUsers(ctx context.Context, cond map[string]interface{}, paging *common.Paging) ([]usermodel.User, error) {
	db := s.db
	db = db.Table(usermodel.User{}.TableName())
	db = db.Where(cond)
	db = db.Where("status=?", 1)
	db = db.Order("id desc")
	var data []usermodel.User

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if paging.FakeCursor != "" {
		uid, err := common.DecomposeUID(paging.FakeCursor)
		if err != nil {
			return nil, err
		}
		db.Where("id <?", uid.GetLocalID())
	} else {
		if err := db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&data).Error; err != nil {
			return nil, common.ErrDB(err)

		}
	}
	return data, nil
}
