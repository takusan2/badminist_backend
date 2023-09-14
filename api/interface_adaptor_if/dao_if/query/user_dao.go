package query_dao_if

import (
	"github.com/takuya-okada-01/badminist/api/domain/user"
	"github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/entity"
	"gorm.io/gorm"
)

type UserDao interface {
	FindUserById(db *gorm.DB, id user.UserId) (entity.User, error)
}
