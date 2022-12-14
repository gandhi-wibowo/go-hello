package controllers

import (
	"hello/models"
	"hello/repository"
	"hello/utils"

	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
		ctx.JSON(code, models.Response(code, errMesg.Error(), nil))
		return
	}
	if utils.CheckPasswordHash(data.Password, result.Password) {
		token, err := utils.GenerateJwt(*result)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.Response(http.StatusInternalServerError, "failed", nil))
			return
		} else {
			ctx.JSON(http.StatusOK, models.Response(http.StatusOK, "success", token))
			data := models.User{
				Token: *token,
			}
			repo.Update(*result, data)
			return
		}
	} else {
		// password salah
		ctx.JSON(http.StatusBadRequest, models.Response(http.StatusBadRequest, "wrong password", nil))
		return
	}
}
func ResetPassword(ctx *gin.Context) {
	// ambil data
	// generate id untuk reset.
	// kirimkan id ke email yang di inputkan.
}

func (ctrl *AuthController) Logout(ctx *gin.Context) {
	userId, err := uuid.FromString(*utils.ExtractJwt(ctx, "id"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, models.Response(http.StatusUnauthorized, "unauthorize", nil))
		return
	}
	currentUser := models.User{
		BaseModel: models.BaseModel{
			ID: userId,
		},
	}

	updatedData := models.User{Token: "-"}
	repo := repository.UserRepo(ctrl.Conn)
	errMesg := repo.Update(currentUser, updatedData)
	if errMesg != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response(http.StatusInternalServerError, "error logout", nil))
		return
	} else {
		ctx.JSON(http.StatusOK, models.Response(http.StatusOK, "success", nil))
		return
	}
}
