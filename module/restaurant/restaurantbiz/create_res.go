package restaurantbiz

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
)

type CreateRestaurantStore interface {
	CreateNewRestaurant(ctx context.Context, res *restaurantmodel.Restaurant) *common.AppError
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, res *restaurantmodel.Restaurant) *common.AppError {
	if err := biz.store.CreateNewRestaurant(ctx, res); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.Entity, err)
	}
	return nil
}
