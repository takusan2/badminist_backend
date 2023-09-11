package dao_if

import (
	"github.com/takuya-okada-01/badminist/app/domain/user"
	"github.com/takuya-okada-01/badminist/app/infrastructure/entity"
	"gorm.io/gorm"
)

type UserDao interface {
	FindUserByEmail(db *gorm.DB, email user.UserEmail) (entity.User, error)
	InsertUser(db *gorm.DB, user entity.User) error
	UpdateUserName(db *gorm.DB, id user.UserId, name user.UserName) error
	UpdateUserStatus(db *gorm.DB, id user.UserId, status user.UserStatus) error
}
