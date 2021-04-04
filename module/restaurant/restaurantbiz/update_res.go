package restaurantbiz

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*restaurantmodel.Restaurant, *common.AppError)
	Update(ctx context.Context, res *restaurantmodel.RestaurantUpdate) *common.AppError
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, res *restaurantmodel.RestaurantUpdate) *common.AppError {
	data, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": res.ID})
	if err != nil {
		return common.ErrCannotGetEntity(restaurantmodel.Entity, err)
	}
	if data.Status == 0 {
		return common.ErrDeletedBefore(restaurantmodel.Entity)
	}
	if err := biz.store.Update(ctx, res); err != nil {
		return common.ErrCannotUpdateEntity(restaurantmodel.Entity, err)
	}
	return nil
}
