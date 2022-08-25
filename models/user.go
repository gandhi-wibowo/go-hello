package models

import (
	"errors"

	"gorm.io/gorm"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	BaseModel
	Name        string `json:"name" validate:"required" binding:"required"`
	Email       string `json:"email" validate:"required" binding:"required" gorm:"unique"`
	PhoneNumber string `json:"phone_number" validate:"required" binding:"required" gorm:"unique"`
	Password    string `json:"password" validate:"required" binding:"required"`
	FcmId       string `json:"fcm_id"`
	Token       string `json:"token"`
}

func (User) TableName() string {
	return "users"
}

func (user *User) BeforeCreate(db *gorm.DB) (err error) {
	uuid := uuid.NewV4()
	if &uuid == nil {
		err = errors.New("can't save invalid data")
		return
	}
	db.Statement.SetColumn("ID", uuid)
	return
}
