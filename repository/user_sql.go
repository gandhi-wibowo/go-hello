package repository

import (
	"errors"
	"hello/models"
	"hello/utils"
	"net/http"

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

func (db *userRepo) Read(credentialId string) (*models.User, int, error) {
	user := models.User{}
	result := db.Conn.Model(&user).Where("id = ?", credentialId).Or("email = ?", credentialId).Or("phone_number = ?", credentialId).Find(&user)
	if user.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return nil, http.StatusNotFound, errors.New("User account not found")
	}
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, errors.New("User account not found")
		}
		return nil, http.StatusInternalServerError, errors.New(result.Error.Error())
	}
	return &user, http.StatusOK, nil
}

func (db *userRepo) Update(user models.User, data models.User) error {
	err := db.Conn.Model(&user).Updates(data).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (db *userRepo) Delete() {

}
