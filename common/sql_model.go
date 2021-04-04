package common

import "time"

// SQLModel type
type SQLModel struct {
	ID        int        `json:"-" gorm:"column:id;"`
	FakeID    *UID       `json:"id" gorm:"-"`
	CreatedAt *time.Time `json:"created_at" form:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" form:"updated_at" gorm:"column:updated_at"`
	Status    int        `json:"status" form:"status" gorm:"column:status;default:1"`
}

func (sqlmodel *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(sqlmodel.ID), dbType, 1)
	sqlmodel.FakeID = &uid
}
