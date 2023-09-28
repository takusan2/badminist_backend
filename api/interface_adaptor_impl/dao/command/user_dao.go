package command_dao

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/takuya-okada-01/badminist/api/domain/user"
	command_dao_if "github.com/takuya-okada-01/badminist/api/interface_adaptor_if/dao_if/command"
	"github.com/takuya-okada-01/badminist/api/interface_adaptor_impl/entity"
	"gorm.io/gorm"
)

type UserDaoImpl struct{}

func NewUserDaoImpl() command_dao_if.UserDao {
	return &UserDaoImpl{}
}

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
	if err := db.Where("email = ?", email.Value()).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, errors.New("そのメールアドレスは登録されていません")
		}
		return entity.User{}, err
	}
	return user, nil
}

func (*UserDaoImpl) FindUserById(db *gorm.DB, id user.UserId) (entity.User, error) {
	var user entity.User
	if err := db.Where("id = ?", id.Value()).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, errors.New("そのユーザーは登録されていません")
		}
		return entity.User{}, err
	}
	return user, nil
}

func (*UserDaoImpl) InsertUser(db *gorm.DB, user entity.User) error {
	if err := db.Create(&user).Error; err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return errors.New("そのメールアドレスはすでに登録されています")
			}
		}
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

func (*UserDaoImpl) UpdateUserConfirmPass(db *gorm.DB, id user.UserId, confirmPass user.UserConfirmPass) error {
	if err := db.Model(&entity.User{}).Where("id = ?", id.Value()).Update("confirm_pass", confirmPass.Value()).Error; err != nil {
		return err
	}
	return nil
}
