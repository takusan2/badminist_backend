package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/takuya-okada-01/badminist/api/infrastructure/database"
	"github.com/takuya-okada-01/badminist/api/infrastructure/entity"
)

func main() {
	godotenv.Load(".env")
	dbConn := database.Connect()
	defer fmt.Println("Successfully migrated")
	defer database.CloseDB(dbConn)
	dbConn.AutoMigrate(
		&entity.User{},
		&entity.Member{},
		&entity.Community{},
		&entity.Player{},
	)
}
