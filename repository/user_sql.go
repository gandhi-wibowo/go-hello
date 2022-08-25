package repository

import (
	"errors"
	"hello/models"
	"hello/utils"

	"gorm.io/gorm"
)

type userRepo struct {
	Conn *gorm.DB
}

func UserRepo(conn *gorm.DB) userRepo {
	return userRepo{
		Conn: conn,
	}
}
func (db *userRepo) Create(data models.User) error {
	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return errors.New("failed to Create User :" + err.Error())
	}
	data.Password = hashedPassword
	if err := db.Conn.Create(&data).Error; err != nil {
		return errors.New("failed to Create User :" + err.Error())
	}
	return nil
}

func (db *userRepo) Read(credentialId string) interface{} {
	// bisa dari id, email, phone number
	result := db.Conn.Model(models.User{}).Where("id = ?", credentialId).Or("email = ?", credentialId).Or("phone_number = ?", credentialId) //.Find(models.User{})
	return result
}

func (db *userRepo) Update() {

}

func (db *userRepo) Delete() {

}
