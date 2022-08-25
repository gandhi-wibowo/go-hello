package models

type RequestLogin struct {
	CredentialId string `json:"credential_id" validate:"required" binding:"required"`
	Password     string `json:"password" validate:"required" binding:"required"`
}
