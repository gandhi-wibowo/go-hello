package controllers

import (
	"os"
	"time"

	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type jwtData struct {
	UserName  string `json:"user_name" validate:"required" binding:"required"`
	UserEmail string `json:"user_email" validate:"required" binding:"required"`
	UserId    string `json:"user_id" validate:"required" binding:"required"`
	*jwt.StandardClaims
}

func createJwtData(username string, useremail string, userid string) *jwtData {
	return &jwtData{
		UserName:  username,
		UserEmail: useremail,
		UserId:    userid,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1000).Unix(),
		},
	}
}

func GenerateJwt(ctx *gin.Context) {
	var data jwtData

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(os.Getenv("JWT_SIGNING_METHOD")), createJwtData(data.UserEmail, data.UserName, data.UserId))
	tokenString, er := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if er != nil {
		ctx.JSON(http.StatusBadRequest, er.Error())
		return
	} else {
		ctx.JSON(http.StatusOK, tokenString)
		return
	}
}

type RequestToken struct {
	Token string `json:"token" validate:"required" binding:"required"`
}

func ValidateJwt(ctx *gin.Context) {
	var requestToken RequestToken
	if err := ctx.ShouldBindJSON(&requestToken); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

}
