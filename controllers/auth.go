package controllers

import (
	"hello/models"
	"hello/repository"
	"hello/utils"

	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Conn *gorm.DB
}

func (ctrl *AuthController) Register(ctx *gin.Context) {
	data := models.User{}
	if err := utils.ExtractRawData(ctx, &data); err != nil {
		return
	}
	repo := repository.UserRepo(ctrl.Conn)
	err := repo.Create(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response(http.StatusBadRequest, err.Error(), nil))
		return
	} else {
		ctx.JSON(http.StatusOK, models.Response(http.StatusOK, "success", nil))
		return
	}
}

func (ctrl *AuthController) Login(ctx *gin.Context) {
	data := models.RequestLogin{}
	if err := utils.ExtractRawData(ctx, &data); err != nil {
		return
	}
	repo := repository.UserRepo(ctrl.Conn)
	result := repo.Read(data.CredentialId)
	ctx.JSON(http.StatusOK, models.Response(http.StatusOK, "success", result))
	return

	// ambil credential id nya.
	// trus ambil 1 data dari database berdasarkan credential id nya.

	// ambil data
	// check ke database
	// Generate JWT.
	//	GenerateJwt(ctx)
}
func ResetPassword(ctx *gin.Context) {
	// ambil data
	// generate id untuk reset.
	// kirimkan id ke email yang di inputkan.
}
