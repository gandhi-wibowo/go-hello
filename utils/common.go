package utils

import (
	"hello/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExtractRawData(ctx *gin.Context, data interface{}) error {
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response(http.StatusBadRequest, err.Error(), nil))
	}
	return err
}
