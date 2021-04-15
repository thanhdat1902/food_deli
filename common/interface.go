package common

import "gorm.io/gorm"

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetSecretKey() string
}

type Requester interface {
	GetID() int
	GetEmail() string
	GetRole() string
}
