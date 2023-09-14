package query_controller

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/takuya-okada-01/badminist/api/domain/community"
	"github.com/takuya-okada-01/badminist/api/domain/user"
	"github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/controller/auth"
	query_processor "github.com/takuya-okada-01/badminist/api/processor/query"
)

type Controller interface {
	GenerateMatchCombination(ctx echo.Context) error
	GetCommunityList(ctx echo.Context) error
	GetPlayerList(ctx echo.Context) error
	GetMemberList(ctx echo.Context) error
}

type controller struct {
	processor query_processor.QueryProcessor
}

func NewController(
	processor query_processor.QueryProcessor,
) Controller {
	return &controller{
		processor: processor,
	}
}

func (c *controller) GenerateMatchCombination(ctx echo.Context) error {
	params := ctx.QueryParams()
	communityId, err := community.CommunityIdFromStr(ctx.Param("community-id"))
	if err != nil {
		return ctx.JSON(400, err)
	}
	numCourt, err := strconv.Atoi(params.Get("num-court"))
	if err != nil {
		return ctx.JSON(400, err)
	}
	rule, err := query_processor.RuleFromStr(params.Get("rule"))
	if err != nil {
		return ctx.JSON(400, err)
	}
	response, err := c.processor.GenerateMatchCombination(
		communityId,
		numCourt,
		rule,
	)
	if err != nil {
		return ctx.JSON(500, err)
	}
	return ctx.JSON(200, response)
}

func (c *controller) GetCommunityList(ctx echo.Context) error {
	id := auth.GetCurrentUser(ctx)
	userId, err := user.UserIdFromStr(id)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	communityList, err := c.processor.GetCommunityList(
		userId,
	)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, communityList)
}

func (c *controller) GetPlayerList(ctx echo.Context) error {
	paramCommunityId := ctx.Param("community-id")
	communityId, err := community.CommunityIdFromStr(paramCommunityId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	playerList, err := c.processor.GetPlayerList(
		communityId,
	)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, playerList)
}

func (c *controller) GetMemberList(ctx echo.Context) error {
	paramCommunityId := ctx.Param("community-id")
	communityId, err := community.CommunityIdFromStr(paramCommunityId)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	memberList, err := c.processor.GetMemberList(
		communityId,
	)
	if err != nil {
		return ctx.JSON(400, err.Error())
	}
	return ctx.JSON(200, memberList)
}
