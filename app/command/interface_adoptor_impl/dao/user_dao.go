package dao

import (
	"github.com/takuya-okada-01/badminist/app/command/domain/user"
	"github.com/takuya-okada-01/badminist/app/command/interface_adaptor_if/dao_if"
	"github.com/takuya-okada-01/badminist/app/infrastructure/entity"
	"gorm.io/gorm"
)

type UserDaoImpl struct{}

func (*UserDaoImpl) UpdateUserStatus(
	db *gorm.DB,
	id user.UserId,
	status user.UserStatus,
) error {
	if err := db.Model(&entity.User{}).
		Where("id = ?", id.Value()).
		Update("status", status.Value()).Error; err != nil {
		return err
	}
	return nil
}

func (*UserDaoImpl) FindUserByEmail(db *gorm.DB, email user.UserEmail) (entity.User, error) {
	var user entity.User
	if err := db.Find(&user, map[string]any{
		"email": email.Value(),
	}).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (*UserDaoImpl) InsertUser(db *gorm.DB, user entity.User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (*UserDaoImpl) UpdateUserName(db *gorm.DB, id user.UserId, name user.UserName) error {
	if err := db.Model(&entity.User{}).Where("id = ?", id).Update("name", name.Value()).Error; err != nil {
		return err
	}
	return nil
}

func NewUserDaoImpl() dao_if.UserDao {
	return &UserDaoImpl{}
}
