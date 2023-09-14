package query_dao

import (
	"github.com/takuya-okada-01/badminist/api/domain/user"
	query_dao_if "github.com/takuya-okada-01/badminist/api/interface_adaptor_if/dao_if/query"
	"github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/entity"
	"gorm.io/gorm"
)

type UserDaoImpl struct{}

func NewUserDaoImpl() query_dao_if.UserDao {
	return &UserDaoImpl{}
}

func (u *UserDaoImpl) FindUserById(
	db *gorm.DB,
	id user.UserId,
) (entity.User, error) {
	var user entity.User
	if err := db.Where("id = ?", id.Value()).First(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}
