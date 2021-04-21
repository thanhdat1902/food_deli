package common

type SimpleUser struct {
	SQLModel  `json:",inline"`
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	FirstName string `json:"first_name" gorm:"column:first_name;"`
	Role      string `json:"-" gorm:"column:role;"`
	Avatar    *Image `json:"avatar" gorm:"column:avatar;type:json"`
}

func (u SimpleUser) TableName() string {
	return "users"
}
