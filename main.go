package main

import (
	"hello/config"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config.Setup("8080")
}
