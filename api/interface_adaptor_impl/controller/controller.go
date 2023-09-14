package controller

import (
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/takuya-okada-01/badminist/api/domain/community"
	"github.com/takuya-okada-01/badminist/api/domain/community/member"
	"github.com/takuya-okada-01/badminist/api/domain/community/player"
	"github.com/takuya-okada-01/badminist/api/domain/user"
	"github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/dto"
	"github.com/takuya-okada-01/badminist/api/processor"
)

type Controller struct {
	commandProcessor processor.CommandProcessor
	queryProcessor   processor.QueryProcessor
}

func GetCurrentUser(ctx echo.Context) string {
	user := ctx.Get("user").(*jwt.Token)
	userClaims := user.Claims.(jwt.MapClaims)
	userID := userClaims["user_id"].(string)
	return userID
}

func NewController(
	commandProcessor processor.CommandProcessor,
	queryProcessor processor.QueryProcessor,
) Controller {
	return Controller{
		commandProcessor: commandProcessor,
		queryProcessor:   queryProcessor,
	}
}

func (c *Controller) CreateCommunity(
	ctx echo.Context,
) error {
	var request dto.CreateCommunityRequestBody
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

	executor := GetCurrentUser(ctx)
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

func (c *Controller) RenameCommunity(
	ctx echo.Context,
) error {
	var request dto.RenameCommunityRequestBody
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

	executor := GetCurrentUser(ctx)
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

func (c *Controller) EditCommunityDescription(
	ctx echo.Context,
) error {
	var request dto.EditCommunityDescriptionRequestBody
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

	executor := GetCurrentUser(ctx)
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

func (c *Controller) DeleteCommunity(
	ctx echo.Context,
) error {
	var request dto.DeleteCommunityRequestBody
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(400, err.Error())
	}

	communityId, err := community.CommunityIdFromStr(request.CommunityId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	executor := GetCurrentUser(ctx)
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

func (c *Controller) AddPlayer(
	ctx echo.Context,
) error {
	var request dto.AddPlayerRequestBody
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

	executor := GetCurrentUser(ctx)
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

func (c *Controller) RemovePlayer(
	ctx echo.Context,
) error {
	var request dto.RemovePlayerRequestBody
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

	executor := GetCurrentUser(ctx)
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

func (c *Controller) ChangePlayerProperty(ctx echo.Context) error {
	var request dto.ChangePlayerPropertyRequestBody
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

	executor := GetCurrentUser(ctx)
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

func (c *Controller) ResetPlayerNumGames(
	ctx echo.Context,
) error {
	var request dto.ResetPlayerNumGamesRequestBody
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
	executor := GetCurrentUser(ctx)
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

func (c *Controller) AddMember(
	ctx echo.Context,
) error {
	var request dto.AddMemberRequestBody
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
	executor := GetCurrentUser(ctx)
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

func (c *Controller) RemoveMember(
	ctx echo.Context,
) error {
	var request dto.RemoveMemberRequestBody
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
	executor := GetCurrentUser(ctx)
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

func (c *Controller) ChangeMemberRole(
	ctx echo.Context,
) error {
	var request dto.ChangeMemberRoleRequestBody
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

	executor := GetCurrentUser(ctx)
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

func (c *Controller) TemporaryRegistration(ctx echo.Context) error {
	var request dto.TemporaryRegistrationRequestBody
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

	if err := c.commandProcessor.TemporaryRegistration(
		name,
		email,
		password,
	); err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, map[string]any{"message": "success"})
}

func (c *Controller) ActivateUser(ctx echo.Context) error {
	var request dto.ActivateUserRequestBody
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

func (c *Controller) GenerateMatchCombination(ctx echo.Context) error {
	fmt.Print("a")
	paramCommunityId := ctx.Param("community-id")
	fmt.Print(ctx.Param("num-court"))
	paramNumCourt, err := strconv.Atoi(ctx.Param("num-court"))
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	paramRule := ctx.Param("rule")

	communityId, err := community.CommunityIdFromStr(paramCommunityId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}

	rule := processor.RuleFromStr(paramRule)

	mathcCombination, err := c.queryProcessor.GenerateMatchCombination(
		communityId,
		paramNumCourt,
		rule,
	)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, mathcCombination)
}

func (c *Controller) Login(ctx echo.Context) error {
	var request dto.LoginRequestBody
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
	return ctx.JSON(200, token)
}

func (c *Controller) GetCommunityList(ctx echo.Context) error {
	id := GetCurrentUser(ctx)
	userId, err := user.UserIdFromStr(id)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	communityList, err := c.queryProcessor.GetCommunityList(
		userId,
	)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, communityList)
}
