package command_repository_if

import "github.com/takuya-okada-01/badminist/api/domain/user"

type UserRepository interface {
	FindUserByEmail(email user.UserEmail) (user.User, error)
	FindUserById(id user.UserId) (user.User, error)
	CreateUser(
		id user.UserId,
		name user.UserName,
		email user.UserEmail,
		password user.UserPassword,
		confirm_pass user.UserConfirmPass,
		status user.UserStatus,
	) error

	UpdateUserName(
		id user.UserId,
		name user.UserName,
	) error

	ActivateUser(id user.UserId) error
	ReissueConfirmPass(id user.UserId, confirmPass user.UserConfirmPass) error
}
