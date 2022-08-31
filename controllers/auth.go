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
	result, code, errMesg := repo.Read(data.CredentialId)
	if errMesg != nil {
		ctx.JSON(http.StatusOK, models.Response(code, errMesg.Error(), nil))
	}
	if utils.CheckPasswordHash(data.Password, result.Password) {
		token, err := utils.GenerateJwt(*result)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.Response(http.StatusInternalServerError, "failed", nil))
		} else {
			ctx.JSON(http.StatusOK, models.Response(http.StatusOK, "success", token))
			// simpan token nya ke data si user.
			// panggil model untuk update
			data := models.User{
				Token: *token,
			}
			repo.Update(*result, data)
		}
	} else {
		// password salah
		ctx.JSON(http.StatusOK, models.Response(http.StatusBadRequest, "wrong password", nil))
	}
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
