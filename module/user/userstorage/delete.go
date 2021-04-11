package userstorage

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/user/usermodel"
)

func (s *store) Delete(ctx context.Context, id int) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).Where("id=?", id).Updates(map[string]interface{}{"Status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
