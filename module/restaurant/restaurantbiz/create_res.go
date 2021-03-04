package restaurantbiz

import (
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
)

type CreateRestaurantStore interface {
	CreateNewRestaurant(res *restaurantmodel.Restaurant) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(res *restaurantmodel.Restaurant) error {
	if err := biz.store.CreateNewRestaurant(res); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.Entity, err)
	}
	return nil
}
