package restaurantstorage

import (
	"context"
	"log"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
)

func (s *store) Create(ctx context.Context, res *restaurantmodel.RestaurantCreate) *common.AppError {

	db := s.db

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Create(&res).Error; err != nil {
		log.Println("Err at DB")
		return common.ErrDB(err)
	}
	return nil
}
