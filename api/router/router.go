package router

import (
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	command_controller "github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/controller/command"
	query_controller "github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/controller/query"
)

type Router struct {
	commandController command_controller.Controller
	query_controller  query_controller.Controller
}

func NewRouter(
	commandController command_controller.Controller,
	query_controller query_controller.Controller,
) *echo.Echo {
	e := echo.New()

	ec := e.Group("/communities")

	ec.Use(
		echojwt.WithConfig(echojwt.Config{
			SigningKey:  []byte(os.Getenv("SECRET_KEY")),
			TokenLookup: "header:Authorization:Bearer ",
		}),
	)
	{
		// Community
		ec.GET("", query_controller.GetCommunityList)
		ec.POST("/create-community", commandController.CreateCommunity)
		ec.PUT("/rename-community", commandController.RenameCommunity)
		ec.PUT("/edit-community-description", commandController.EditCommunityDescription)
		ec.DELETE("/delete-community", commandController.DeleteCommunity)

		// Player
		ec.GET("/:community-id/players", query_controller.GetPlayerList)
		ec.POST("/add-player", commandController.AddPlayer)
		ec.PUT("/change-player-property", commandController.ChangePlayerProperty)
		ec.PUT("/reset-player-num-games", commandController.ResetPlayerNumGames)
		ec.DELETE("/remove-player", commandController.RemovePlayer)

		// Member
		ec.GET("/:community-id/members", query_controller.GetMemberList)
		ec.POST("/add-member", commandController.AddMember)
		ec.PUT("/change-member-role", commandController.ChangeMemberRole)
		ec.DELETE("/remove-member", commandController.RemoveMember)

		// Match
		ec.GET("/:community-id/generate-matches", query_controller.GenerateMatchCombination)
	}

	eu := e.Group("/auth")
	{
		eu.POST("/temporary-registration", commandController.TemporaryRegistration)
		eu.POST("/activate-user", commandController.ActivateUser)
		eu.POST("/login", commandController.Login)
	}

	return e
}
