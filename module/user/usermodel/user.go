package usermodel

import "github.com/thanhdat1902/restapi/food_deli/common"

// EntityName variable
const EntityName = "User"

// JSON type
type JSON []byte

//User struct
type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" form:"email" gorm:"column:email"`
	FacebookID      *string       `json:"fb_id" form:"fb_id" gorm:"column:fb_id"`
	GoogleID        *string       `json:"gg_id" form:"gg_id" gorm:"column:gg_id"`
	Password        string        `json:"password" form:"password" gorm:"column:password"`
	Salt            string        `json:"-" gorm:"column:salt"`
	Lastname        string        `json:"last_name" form:"last_name" gorm:"column:last_name"`
	Firstname       string        `json:"first_name" form:"first_name" gorm:"column:first_name;default:'User'"`
	Phone           string        `json:"phone" form:"phone" gorm:"column:phone"`
	Role            string        `json:"role" form:"role" gorm:"column:role;default:'USER'"`
	Avatar          *common.Image `json:"avatar" form:"avatar" gorm:"column:avatar"`
}

// TableName of user
func (User) TableName() string {
	return "users"
}
