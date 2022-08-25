package utils

import (
	"fmt"

	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB
var dbOnce sync.Once

func OpenDbConnection() *gorm.DB {
	dbOnce.Do(func() {
		DATABASE_HOST := os.Getenv("DATABASE_HOST")
		DATABASE_PORT := os.Getenv("DATABASE_PORT")
		DATABASE_NAME := os.Getenv("DATABASE_NAME")
		DATABASE_LOGIN_USERNAME := os.Getenv("DATABASE_LOGIN_USERNAME")
		DATABASE_LOGIN_PASSWORD := os.Getenv("DATABASE_LOGIN_PASSWORD")

		DSN := DATABASE_LOGIN_USERNAME + ":" + DATABASE_LOGIN_PASSWORD + "@tcp(" + DATABASE_HOST + ":" + DATABASE_PORT + ")/" + DATABASE_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		sqlDb, err := db.DB()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		sqlDb.SetConnMaxLifetime(time.Second * 60)
		dbInstance = db
	})
	return dbInstance
}
