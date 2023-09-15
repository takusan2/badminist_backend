package entity

import (
	"time"

	"github.com/takuya-okada-01/badminist/api/domain/user"
	"github.com/takuya-okada-01/badminist/api/processor/query/read_model"
)

type User struct {
	Id          string    `json:"id" gorm:"type:varchar(36);primaryKey;"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null;"`
	Email       string    `json:"email" gorm:"type:varchar(255);not null;"`
	Password    string    `json:"password" gorm:"type:varchar(255);not null;"`
	ConfirmPass string    `json:"confirm_pass" gorm:"type:varchar(255);not null;"`
	Status      string    `json:"status" gorm:"type:varchar(255);not null;"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}

func NewUser(
	id string,
	name string,
	email string,
	password string,
	confirmPass string,
	status string,
) User {
	return User{
		Id:          id,
		Name:        name,
		Email:       email,
		Password:    password,
		ConfirmPass: confirmPass,
		Status:      status,
	}
}

func (u *User) ToDomainObject() user.User {
	userId, err := user.UserIdFromStr(u.Id)
	if err != nil {
		panic(err)
	}

	name, err := user.NewUserName(u.Name)
	if err != nil {
		panic(err)
	}

	email, err := user.NewUserEmail(u.Email)
	if err != nil {
		panic(err)
	}

	password, err := user.NewUserPassword(u.Password)
	if err != nil {
		panic(err)
	}

	confirmPass, err := user.UserConfirmPassFromStr(u.ConfirmPass)
	if err != nil {
		panic(err)
	}

	status, err := user.UserStatusFromStr(u.Status)
	if err != nil {
		panic(err)
	}

	return user.NewUser(
		userId,
		name,
		email,
		password,
		confirmPass,
		status,
	)
}

func (u *User) ToReadModel() read_model.User {
	return read_model.User{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
	}
}
