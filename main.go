package main

import (
	"hello/config"
	"hello/repository"
	"hello/utils"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := utils.OpenDbConnection()
	repository.Migration(*db)
	conf := config.ConfigRoute{Conn: db}
	conf.Setup("8080")
}
