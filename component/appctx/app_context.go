package appctx

import "gorm.io/gorm"

type ctx struct {
	mainDB *gorm.DB
	secret string
}

func New(mainDB *gorm.DB, secret string) *ctx {
	return &ctx{mainDB: mainDB, secret: secret}
}

// Get method
func (c ctx) GetMainDBConnection() *gorm.DB {
	return c.mainDB
}

func (c ctx) GetSecretKey() string {
	return c.secret
}

// token expire struct
type tokenExpiry struct {
	atExp int
	rtExp int
}

func NewTokenConfig() tokenExpiry {
	return tokenExpiry{
		atExp: 60 * 60 * 24 * 7,
		rtExp: 60 * 60 * 25 * 7 * 2,
	}
}

func (t tokenExpiry) GetAtExp() int {
	return t.atExp
}

func (t tokenExpiry) GetRtExp() int {
	return t.rtExp
}
