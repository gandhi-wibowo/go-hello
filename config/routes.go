package config

import (
	"fmt"

	"hello/controllers"
	"hello/utils"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// CORS is function for enable cors on backend
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, x-api-key, X-BCA-Key, X-BCA-Timestamp, X-BCA-Signature")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

type ConfigRoute struct {
	Conn *gorm.DB
}

func (conf *ConfigRoute) Setup(port string) {

	app := gin.Default()
	app.Use(cors())

	authController := controllers.AuthController{Conn: conf.Conn}

	auth := app.Group("/auth")
	auth.POST("/register", authController.Register)                // Register
	auth.POST("/login", authController.Login)                      // Login
	auth.POST("/reset-password", controllers.ResetPassword)        // Reset Password
	auth.POST("/logout", utils.MidleWare(), authController.Logout) // Reset Password

	user := app.Group("/user")
	user.GET("", utils.MidleWare(), controllers.GetProfile)                        // Read
	user.PATCH("", utils.MidleWare(), controllers.PatchProfile)                    // Update
	user.DELETE("", utils.MidleWare(), controllers.DeleteProfile)                  // DELETE
	user.PUT("/change-password", utils.MidleWare(), controllers.PutChangePassword) // Change password

	runningOnPort := fmt.Sprintf(":%s", port)
	app.Run(runningOnPort)
}
