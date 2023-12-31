package command_dao_if

import (
	"github.com/takuya-okada-01/badminist/api/domain/user"
	"github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/entity"
	"gorm.io/gorm"
)

type UserDao interface {
	FindUserByEmail(db *gorm.DB, email user.UserEmail) (entity.User, error)
	FindUserById(db *gorm.DB, id user.UserId) (entity.User, error)
	InsertUser(db *gorm.DB, user entity.User) error
	UpdateUserName(db *gorm.DB, id user.UserId, name user.UserName) error
	UpdateUserStatus(db *gorm.DB, id user.UserId, status user.UserStatus) error
	UpdateUserConfirmPass(db *gorm.DB, id user.UserId, confirmPass user.UserConfirmPass) error
}
