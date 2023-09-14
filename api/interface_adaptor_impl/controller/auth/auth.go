package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetCurrentUser(ctx echo.Context) string {
	user := ctx.Get("user").(*jwt.Token)
	userClaims := user.Claims.(jwt.MapClaims)
	userID := userClaims["user_id"].(string)
	return userID
}
