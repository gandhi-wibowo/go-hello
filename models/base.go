package models

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at,string,omitempty"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at,string,omitempty"`
	DeletedAt time.Time `gorm:"type:datetime" json:"deleted_at,string,omitempty"`
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
