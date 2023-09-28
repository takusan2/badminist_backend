package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/resendlabs/resend-go"
	log "github.com/sirupsen/logrus"
	"github.com/takuya-okada-01/badminist/api/infrastructure/database"
	"github.com/takuya-okada-01/badminist/api/utils"

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
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	godotenv.Load(".env")
	db := database.Connect()
	defer database.CloseDB(db)

	resendApiKey := os.Getenv("RESEND_API_KEY")
	resendClient := resend.NewClient(resendApiKey)
	emailServer := utils.NewEmailServer(resendClient)

	communityDao := command_dao.NewCommunityDaoImpl()
	userDao := command_dao.NewUserDaoImpl()

	queryCommunityDao := query_dao.NewCommunityDaoImpl()
	queryUserDao := query_dao.NewUserDaoImpl()

	commandCommunityRepo := command_repository.NewCommunityRepositoryImpl(db, communityDao)
	commandUserRepo := command_repository.NewUserRepositoryImpl(db, userDao)

	commandProcessor := command_processor.NewCommandProcessor(
		commandCommunityRepo,
		commandUserRepo,
		emailServer,
	)
	queryProcessor := query_processor.NewQueryProcessor(
		db,
		queryCommunityDao,
		queryUserDao,
	)

	commandController := command_controller.NewController(commandProcessor, queryProcessor)
	queryController := query_controller.NewController(queryProcessor)

	router := router.NewRouter(commandController, queryController)
	router.Logger.Fatal(router.Start(":8080"))
}
