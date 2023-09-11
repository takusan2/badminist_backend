package repository

import (
	"github.com/takuya-okada-01/badminist/app/command/domain/user"
	"github.com/takuya-okada-01/badminist/app/command/interface_adaptor_if/dao_if"
	"github.com/takuya-okada-01/badminist/app/command/interface_adaptor_if/repository_if.go"
	"github.com/takuya-okada-01/badminist/app/infrastructure/entity"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db  *gorm.DB
	dao dao_if.UserDao
}

// FindUserByEmail implements repository_if.UserRepository.
func (u *UserRepositoryImpl) FindUserByEmail(email user.UserEmail) (user.User, error) {
	userEntity, err := u.dao.FindUserByEmail(u.db, email)
	if err != nil {
		return user.User{}, err
	}
	return userEntity.ToDomainObject(), nil
}

// CreateUser implements repository_if.UserRepository.
func (u *UserRepositoryImpl) CreateUser(
	id user.UserId,
	name user.UserName,
	email user.UserEmail,
	password user.UserPassword,
	confirmPass user.UserConfirmPass,
	status user.UserStatus,
) error {
	user := entity.NewUser(
		id.Value(),
		name.Value(),
		email.Value(),
		password.Value(),
		confirmPass.Value(),
		status.Value(),
	)
	err := u.dao.InsertUser(u.db, user)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserName implements repository_if.UserRepository.
func (u *UserRepositoryImpl) UpdateUserName(
	id user.UserId,
	name user.UserName,
) error {
	if err := u.dao.UpdateUserName(u.db, id, name); err != nil {
		return err
	}
	return nil
}

func (u *UserRepositoryImpl) ActivateUser(
	id user.UserId,
) error {
	status, _ := user.NewUserStatus(user.Active)
	if err := u.dao.UpdateUserStatus(u.db, id, status); err != nil {
		return err
	}
	return nil
}

func NewUserRepositoryImpl(
	db *gorm.DB,
	dao dao_if.UserDao,
) repository_if.UserRepository {
	return &UserRepositoryImpl{
		db:  db,
		dao: dao,
	}
}
