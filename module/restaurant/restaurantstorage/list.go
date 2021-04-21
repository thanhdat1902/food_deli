package restaurantstorage

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
)

func (s *store) ListDataWithCondition(ctx context.Context, condition map[string]interface{}, paging *common.Paging, moreData ...string) ([]restaurantmodel.Restaurant, *common.AppError) {
	db := s.db
	db = db.Table(restaurantmodel.Restaurant{}.TableName())
	db = db.Where("status=?", 1)
	if err := db.Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	var data []restaurantmodel.Restaurant
	paging.Fulfill()

	for i := range moreData {
		db.Preload(moreData[i])
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return data, nil
}
