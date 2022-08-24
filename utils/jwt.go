package utils

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func MidleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.Request.Header.Get("Authorization")
		tokenString := strings.ReplaceAll(bearerToken, "Bearer ", "")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod(os.Getenv("JWT_SIGNING_METHOD")) != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if token == nil || err != nil {
			ctx.JSON(http.StatusUnauthorized, "not authorize")
			ctx.Abort()
			return
		}
		_, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			ctx.JSON(http.StatusUnauthorized, "not authorize")
			ctx.Abort()
			return
		}
	}
}

func ExtractJwt(ctx *gin.Context, key string) *string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	tokenString := strings.ReplaceAll(bearerToken, "Bearer ", "")
	tkn, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		fmt.Println(tkn)
		fmt.Println(err.Error())
		return nil
	}

	tokenData, ok := tkn.Claims.(jwt.MapClaims)
	if ok && tkn.Valid {
		ret := tokenData[key].(string)
		if ret == "" {
			return nil
		} else {
			return &ret
		}
	} else {
		return nil
	}
}
