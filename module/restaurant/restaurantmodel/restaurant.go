package restaurantmodel

import (
	"github.com/thanhdat1902/restapi/food_deli/common"
)

// Restaurant type
type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" form:"name" gorm:"column:name"`
	CityID          *int               `json:"city_id" form:"city_id" gorm:"column:city_id"`
	Addr            *string            `json:"addr" form:"addr" gorm:"column:addr"`
	Lat             *float64           `json:"lat" form:"lat" gorm:"column:lat"`
	Long            *float64           `json:"lng" form:"lng" gorm:"column:lng"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo"`
	OpenHour        *string            `json:"open_hour" form:"open_hour" gorm:"column:open_hour"`
	CloseHour       *string            `json:"close_hour" form:"close_hour" gorm:"column:close_hour"`
	Description     string             `json:"description" form:"description" gorm:"column:description"`
	OwnerID         int                `json:"-" gorm:"column:owner_id"`
	FakeOwnerID     *common.UID        `json:"owner_id" gorm:"-"`
	Owner           *common.SimpleUser `json:"owner" gorm:"foreignKey:OwnerID;"`
}

// Entity name
var Entity = "Restaurant"

// Mask the UID
func (r *Restaurant) Mask() {
	r.GenUID(common.DbType)
	uid := common.NewUID(uint32(r.OwnerID), common.DbTypeUser, 1)
	r.FakeOwnerID = &uid
	if r.Owner != nil {
		r.Owner.GenUID(common.DbTypeUser)
	}
}

// TableName of restaurants
func (Restaurant) TableName() string {
	return "restaurants"
}
