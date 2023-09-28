package command_controller

import (
	"github.com/labstack/echo/v4"
	"github.com/takuya-okada-01/badminist/api/domain/community"
	"github.com/takuya-okada-01/badminist/api/domain/community/member"
	"github.com/takuya-okada-01/badminist/api/domain/community/player"
	"github.com/takuya-okada-01/badminist/api/domain/user"
	"github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/controller/auth"
	auth_dto "github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/dto/auth"
	command_dto "github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/dto/command"

	command_processor "github.com/takuya-okada-01/badminist/api/processor/command"
	query_processor "github.com/takuya-okada-01/badminist/api/processor/query"
)

type Controller interface {
	CreateCommunity(ctx echo.Context) error
	RenameCommunity(ctx echo.Context) error
	EditCommunityDescription(ctx echo.Context) error
	DeleteCommunity(ctx echo.Context) error
	AddPlayer(ctx echo.Context) error
	RemovePlayer(ctx echo.Context) error
	ChangePlayerProperty(ctx echo.Context) error
	ResetPlayerNumGames(ctx echo.Context) error
	ChangePlayerNumGames(ctx echo.Context) error
	AddMember(ctx echo.Context) error
	RemoveMember(ctx echo.Context) error
	ChangeMemberRole(ctx echo.Context) error
	TemporaryRegistration(ctx echo.Context) error
	ActivateUser(ctx echo.Context) error
	Login(ctx echo.Context) error
}

type controller struct {
	commandProcessor command_processor.CommandProcessor
	queryProcessor   query_processor.QueryProcessor
}

func NewController(
	commandProcessor command_processor.CommandProcessor,
	queryProcessor query_processor.QueryProcessor,
) Controller {
	return &controller{
		commandProcessor: commandProcessor,
		queryProcessor:   queryProcessor,
	}
}

func (c *controller) CreateCommunity(
	ctx echo.Context,
) error {
	var request command_dto.CreateCommunityRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(400, err.Error())
	}

	name, err := community.NewCommunityName(request.Name)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	description, err := community.NewCommunityDescription(request.Description)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	executor := auth.GetCurrentUser(ctx)
	executorId, err := user.UserIdFromStr(executor)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	if err := c.commandProcessor.CreateCommunity(
		name,
		description,
		executorId,
	); err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, map[string]any{"message": "success"})
}

func (c *controller) RenameCommunity(
	ctx echo.Context,
) error {
	var request command_dto.RenameCommunityRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(400, err.Error())

	}

	communityId, err := community.CommunityIdFromStr(request.CommunityId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	name, err := community.NewCommunityName(request.Name)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	executor := auth.GetCurrentUser(ctx)
	executorId, err := user.UserIdFromStr(executor)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	if err := c.commandProcessor.RenameCommunity(
		communityId,
		name,
		executorId,
	); err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, map[string]any{"message": "success"})
}

func (c *controller) EditCommunityDescription(
	ctx echo.Context,
) error {
	var request command_dto.EditCommunityDescriptionRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(400, err.Error())
	}

	communityId, err := community.CommunityIdFromStr(request.CommunityId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	description, err := community.NewCommunityDescription(request.Description)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	executor := auth.GetCurrentUser(ctx)
	executorId, err := user.UserIdFromStr(executor)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	if err := c.commandProcessor.EditCommunityDescription(
		communityId,
		description,
		executorId,
	); err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, map[string]any{"message": "success"})
}

func (c *controller) DeleteCommunity(
	ctx echo.Context,
) error {
	var request command_dto.DeleteCommunityRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(400, err.Error())
	}

	communityId, err := community.CommunityIdFromStr(request.CommunityId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	executor := auth.GetCurrentUser(ctx)
	executorId, err := user.UserIdFromStr(executor)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	if err := c.commandProcessor.DeleteCommunity(
		communityId,
		executorId,
	); err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, map[string]any{"message": "success"})
}

func (c *controller) AddPlayer(
	ctx echo.Context,
) error {
	var request command_dto.AddPlayerRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(400, err.Error())
	}

	communityId, err := community.CommunityIdFromStr(request.CommunityId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	playerName, err := player.NewPlayerName(request.PlayerName)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	playerGender, err := player.PlayerGenderFromStr(request.PlayerGender)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	playerAge, err := player.NewPlayerAge(request.PlayerAge)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	playerLevel, err := player.PlayerLevelFromStr(request.PlayerLevel)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	playerNumGames, err := player.NewPlayerNumGames(request.PlayerNumGames)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	playerStatus, err := player.PlayerStatusFromStr(request.PlayerStatus)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	executor := auth.GetCurrentUser(ctx)
	executorId, err := user.UserIdFromStr(executor)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	if err := c.commandProcessor.AddPlayer(
		communityId,
		playerName,
		playerGender,
		playerAge,
		playerLevel,
		playerNumGames,
		playerStatus,
		executorId,
	); err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, map[string]any{"message": "success"})
}

func (c *controller) RemovePlayer(
	ctx echo.Context,
) error {
	var request command_dto.RemovePlayerRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(400, err.Error())
	}

	communityId, err := community.CommunityIdFromStr(request.CommunityId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	playerId, err := player.PlayerIdFromStr(request.PlayerId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	executor := auth.GetCurrentUser(ctx)
	executorId, err := user.UserIdFromStr(executor)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	if err := c.commandProcessor.RemovePlayer(
		communityId,
		playerId,
		executorId,
	); err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, map[string]any{"message": "success"})
}

func (c *controller) ChangePlayerProperty(ctx echo.Context) error {
	var request command_dto.ChangePlayerPropertyRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(400, err.Error())
	}

	communityId, err := community.CommunityIdFromStr(request.CommunityId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	playerId, err := player.PlayerIdFromStr(request.PlayerId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	playerName, err := player.NewPlayerName(request.PlayerName)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	playerGender, err := player.PlayerGenderFromStr(request.PlayerGender)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	playerAge, err := player.NewPlayerAge(request.PlayerAge)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	playerLevel, err := player.PlayerLevelFromStr(request.PlayerLevel)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	playerNumGames, err := player.NewPlayerNumGames(request.PlayerNumGames)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	playerStatus, err := player.PlayerStatusFromStr(request.PlayerStatus)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	executor := auth.GetCurrentUser(ctx)
	executorId, err := user.UserIdFromStr(executor)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	if err := c.commandProcessor.ChangePlayerProperty(
		communityId,
		playerId,
		playerName,
		playerGender,
		playerAge,
		playerLevel,
		playerNumGames,
		playerStatus,
		executorId,
	); err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, map[string]any{"message": "success"})
}

func (c *controller) ResetPlayerNumGames(
	ctx echo.Context,
) error {
	var request command_dto.ResetPlayerNumGamesRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(400, err.Error())
	}
	communityId, err := community.CommunityIdFromStr(request.CommunityId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	playerId, err := player.PlayerIdFromStr(request.PlayerId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	executor := auth.GetCurrentUser(ctx)
	executorId, err := user.UserIdFromStr(executor)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	if err := c.commandProcessor.ResetPlayerNumGames(
		communityId,
		playerId,
		executorId,
	); err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, map[string]any{"message": "success"})
}

func (c *controller) ChangePlayerNumGames(
	ctx echo.Context,
) error {
	var request command_dto.ChangePlayerNumGamesRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(400, err.Error())
	}
	communityId, err := community.CommunityIdFromStr(request.CommunityId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	playerId, err := player.PlayerIdFromStr(request.PlayerId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	playerNumGames, err := player.NewPlayerNumGames(request.NumGames)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	executor := auth.GetCurrentUser(ctx)
	executorId, err := user.UserIdFromStr(executor)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	if err := c.commandProcessor.ChangePlayerNumGames(
		communityId,
		playerId,
		playerNumGames,
		executorId,
	); err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, map[string]any{"message": "success"})

}

func (c *controller) AddMember(
	ctx echo.Context,
) error {
	var request command_dto.AddMemberRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(400, err.Error())
	}
	communityId, err := community.CommunityIdFromStr(request.CommunityId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	userId, err := user.UserIdFromStr(request.UserId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	memberRole, err := member.MemberRoleFromStr(request.MemberRole)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	executor := auth.GetCurrentUser(ctx)
	executorId, err := user.UserIdFromStr(executor)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	if err := c.commandProcessor.AddMember(
		communityId,
		userId,
		memberRole,
		executorId,
	); err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, map[string]any{"message": "success"})
}

func (c *controller) RemoveMember(
	ctx echo.Context,
) error {
	var request command_dto.RemoveMemberRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(400, err.Error())
	}
	communityId, err := community.CommunityIdFromStr(request.CommunityId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	userId, err := user.UserIdFromStr(request.UserId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	executor := auth.GetCurrentUser(ctx)
	executorId, err := user.UserIdFromStr(executor)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	if err := c.commandProcessor.RemoveMember(
		communityId,
		userId,
		executorId,
	); err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, map[string]any{"message": "success"})
}

func (c *controller) ChangeMemberRole(
	ctx echo.Context,
) error {
	var request command_dto.ChangeMemberRoleRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(400, err.Error())
	}
	communityId, err := community.CommunityIdFromStr(request.CommunityId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	userId, err := user.UserIdFromStr(request.UserId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	memberRole, err := member.MemberRoleFromStr(request.MemberRole)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	executor := auth.GetCurrentUser(ctx)
	executorId, err := user.UserIdFromStr(executor)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	if err := c.commandProcessor.ChangeMemberRole(
		communityId,
		userId,
		memberRole,
		executorId,
	); err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, map[string]any{"message": "success"})
}

func (c *controller) TemporaryRegistration(ctx echo.Context) error {
	var request auth_dto.TemporaryRegistrationRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(400, err.Error())
	}

	name, err := user.NewUserName(request.Name)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	email, err := user.NewUserEmail(request.Email)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	password, err := user.NewUserPassword(request.Password)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	token, err := c.commandProcessor.TemporaryRegistration(
		name,
		email,
		password,
	)

	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, map[string]any{"token": token})
}

func (c *controller) ActivateUser(ctx echo.Context) error {
	var request auth_dto.ActivateUserRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(400, err.Error())
	}

	email, err := user.NewUserEmail(request.Email)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	confirmPass, err := user.UserConfirmPassFromStr(request.ConfirmPass)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	if err := c.commandProcessor.ActivateUser(
		email,
		confirmPass,
	); err != nil {
		return ctx.JSON(400, err.Error())
	}

	return ctx.JSON(200, map[string]any{"message": "success"})

}

func (c *controller) Login(ctx echo.Context) error {
	var request auth_dto.LoginRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(400, err.Error())
	}

	email, err := user.NewUserEmail(request.Email)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	password, err := user.NewUserPassword(request.Password)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	token, err := c.commandProcessor.Login(
		email,
		password,
	)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, auth_dto.LoginResponseBody{Token: token})
}
