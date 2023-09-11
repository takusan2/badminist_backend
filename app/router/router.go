package router

import (
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/takuya-okada-01/badminist/app/command/interface_adoptor_impl/controller"
)

type Router struct {
	controller controller.Controller
}

func NewRouter(controller controller.Controller) *echo.Echo {
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
		ec.POST("/create-community", controller.CreateCommunity)
		ec.PUT("/rename-community", controller.RenameCommunity)
		ec.PUT("/edit-community-description", controller.EditCommunityDescription)
		ec.DELETE("/delete-community", controller.DeleteCommunity)

		// Player
		ec.POST("/add-player", controller.AddPlayer)
		ec.DELETE("/remove-player", controller.RemovePlayer)
		ec.PUT("/change-player-property", controller.ChangePlayerProperty)
		ec.PUT("/reset-player-num-games", controller.ResetPlayerNumGames)

		// Member
		ec.POST("/add-member", controller.AddMember)
		ec.DELETE("/remove-member", controller.RemoveMember)
		ec.PUT("/change-member-role", controller.ChangeMemberRole)
	}

	eu := e.Group("/users")
	{
		eu.POST("/temporary-registration", controller.TemporaryRegistration)
		eu.POST("/activate-user", controller.ActivateUser)
		eu.POST("/login", controller.Login)
	}

	em := e.Group("/mathces")
	{
		em.GET("/generate-match-combination/:community-id/:num-court/:rule", controller.GenerateMatchCombination)
	}
	return e
}
