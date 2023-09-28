package main

import (
	"github.com/joho/godotenv"
	"github.com/takuya-okada-01/badminist/api/infrastructure/database"
	"github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/entity"
)

func main() {
	godotenv.Load(".env")
	dbConn := database.Connect()
	defer database.CloseDB(dbConn)
	dbConn.AutoMigrate(
		&entity.User{},
		&entity.Member{},
		&entity.Community{},
		&entity.Player{},
	)
}
