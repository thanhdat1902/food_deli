package restaurantbiz

import (
	"github.com/thanhdat1902/restapi/food_deli/common"
	"github.com/thanhdat1902/restapi/food_deli/module/restaurant/restaurantmodel"
)

type ListResStore interface {
	ListDataWithCondition(condition map[string]interface{}, paging *common.Paging) ([]restaurantmodel.Restaurant, error)
}
type listResBiz struct {
	store ListResStore
}

func ListResBiz(store ListResStore) *listResBiz {
	return &listResBiz{store: store}
}

func (biz *listResBiz) ListRestaurant(paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(nil, paging)

	if err != nil {
		return nil, err
	}
	return result, nil
}
