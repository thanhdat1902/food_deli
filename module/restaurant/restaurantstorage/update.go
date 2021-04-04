package restaurantstorage

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
)

func (s *store) Update(ctx context.Context, res *restaurantmodel.RestaurantUpdate) *common.AppError {
	db := s.db
	db = db.Table(restaurantmodel.Restaurant{}.TableName())
	db = db.Where("id=?", res.ID)
	if err := db.Updates(res).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
