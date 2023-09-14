package main

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/takuya-okada-01/badminist/api/infrastructure/database"

	command_controller "github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/controller/command"
	query_controller "github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/controller/query"
	command_dao "github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/dao/command"
	query_dao "github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/dao/query"

	command_repository "github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/repository/command"
	command_processor "github.com/takuya-okada-01/badminist/api/processor/command"
	query_processor "github.com/takuya-okada-01/badminist/api/processor/query"

	"github.com/takuya-okada-01/badminist/api/router"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	godotenv.Load(".env")
	db := database.Connect()
	defer database.CloseDB(db)

	communityDao := command_dao.NewCommunityDaoImpl()
	userDao := command_dao.NewUserDaoImpl()

	queryCommunityDao := query_dao.NewCommunityDaoImpl()
	queryUserDao := query_dao.NewUserDaoImpl()

	commandCommunityRepo := command_repository.NewCommunityRepositoryImpl(db, communityDao)
	commandUserRepo := command_repository.NewUserRepositoryImpl(db, userDao)

	commandProcessor := command_processor.NewCommandProcessor(commandCommunityRepo, commandUserRepo)
	queryProcessor := query_processor.NewQueryProcessor(db, queryCommunityDao, queryUserDao)

	commandController := command_controller.NewController(commandProcessor, queryProcessor)
	queryController := query_controller.NewController(queryProcessor)

	router := router.NewRouter(commandController, queryController)
	router.Logger.Fatal(router.Start(":8080"))
}
