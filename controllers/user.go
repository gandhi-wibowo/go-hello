package controllers

import (
	"fmt"
	"hello/utils"

	"github.com/gin-gonic/gin"
)

func GetProfile(ctx *gin.Context) {
	userName := utils.ExtractJwt(ctx, "user_id")
	fmt.Println(*userName)
	// ambil data.
	// ada userId
}

func PatchProfile(ctx *gin.Context) {
	// ambil data.
	// update data sesuai dengan yang di inputkan.
}

func DeleteProfile(ctx *gin.Context) {
	// ambil data usernya.
	// soft delete
}

func PutChangePassword(ctx *gin.Context) {
	// ambil data
	// ubah passwordnya.
}
