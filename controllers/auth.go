package controllers

import (
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	// ambil data
	// simpan ke database
}
func Login(ctx *gin.Context) {
	// ambil data
	// check ke database
	// Generate JWT.
	GenerateJwt(ctx)
}
func ResetPassword(ctx *gin.Context) {
	// ambil data
	// generate id untuk reset.
	// kirimkan id ke email yang di inputkan.
}
