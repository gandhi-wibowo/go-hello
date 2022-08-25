package models

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt time.Time `json:"deleted_at" gorm:"type:timestamp"`
}

func (base *BaseModel) BeforeCreate(db *gorm.DB) (err error) {
	uuid := uuid.NewV4()
	if &uuid == nil {
		err = errors.New("can't save invalid data")
		return
	}
	db.Statement.SetColumn("ID", uuid)
	return
}
