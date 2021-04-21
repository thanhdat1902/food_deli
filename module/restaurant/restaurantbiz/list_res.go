package restaurantbiz

import (
	"context"

	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
)

type ListResStore interface {
	ListDataWithCondition(ctx context.Context, condition map[string]interface{}, paging *common.Paging, moreData ...string) ([]restaurantmodel.Restaurant, *common.AppError)
}
type listResBiz struct {
	store ListResStore
}

func NewListResBiz(store ListResStore) *listResBiz {
	return &listResBiz{store: store}
}

func (biz *listResBiz) ListRestaurant(ctx context.Context, paging *common.Paging) ([]restaurantmodel.Restaurant, *common.AppError) {
	result, err := biz.store.ListDataWithCondition(ctx, nil, paging, "Owner")

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.Entity, err)
	}
	return result, nil
}
