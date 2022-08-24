package models

type User struct {
	BaseModel
	Name        string
	Email       string
	PhoneNumber string
	Password    string
	FcmId       string
	Token       string
}
