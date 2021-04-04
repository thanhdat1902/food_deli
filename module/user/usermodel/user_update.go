package usermodel

import "github.com/thanhdat1902/restapi/food_deli/common"

type UserUpdate struct {
	common.SQLModel `json:",inline"`
	Fullname        string `json:"fullname" form:"fullname" gorm:"default:'User'"`
	Role            string `json:"role" form:"role" gorm:"default:'USER'"`
	Avatar          JSON   `json:"avatar" form:"avatar"`
	Password        string `json:"password" form:"password"`
	Email           string `json:"email" form:"email"`
	PhoneNumber     string `json:"phone_number" form:"phone_number"`
}
func (UserUpdate) TableName() string {
	return "users"
}