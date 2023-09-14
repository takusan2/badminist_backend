package main

import (
	"github.com/joho/godotenv"
	"github.com/takuya-okada-01/badminist/api/infrastructure/database"
	"github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/controller"
	"github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/dao"
	"github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/repository"
	"github.com/takuya-okada-01/badminist/api/processor"
	"github.com/takuya-okada-01/badminist/api/router"
)

func main() {
	godotenv.Load(".env")
	db := database.Connect()
	defer database.CloseDB(db)

	communityDao := dao.NewCommunityDaoImpl()
	userDao := dao.NewUserDaoImpl()

	communityRepo := repository.NewCommunityRepositoryImpl(db, communityDao)
	userRepo := repository.NewUserRepositoryImpl(db, userDao)

	commandProcessor := processor.NewCommandProcessor(communityRepo, userRepo)
	queryProcessor := processor.NewQueryProcessor(db, communityDao)

	controller := controller.NewController(commandProcessor, queryProcessor)

	router := router.NewRouter(controller)
	router.Logger.Fatal(router.Start(":8080"))
}
