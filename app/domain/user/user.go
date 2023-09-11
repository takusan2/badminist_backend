package user

import "time"

type User struct {
	id          UserId
	name        UserName
	email       UserEmail
	password    UserPassword
	confirmPass UserConfirmPass
	status      UserStatus
}

func NewUser(
	id UserId,
	name UserName,
	email UserEmail,
	password UserPassword,
	confirmPass UserConfirmPass,
	status UserStatus,
) User {
	return User{
		id:          id,
		name:        name,
		email:       email,
		password:    password,
		confirmPass: confirmPass,
		status:      status,
	}
}

func (u *User) BreachEncapsulationOfId() UserId {
	return u.id
}
func (u *User) BreachEncapsulationOfName() UserName {
	return u.name
}
func (u *User) BreachEncapsulationOfEmail() UserEmail {
	return u.email
}
func (u *User) BreachEncapsulationOfPassword() UserPassword {
	return u.password
}
func (u *User) BreachEncapsulationOfConfirmPass() UserConfirmPass {
	return u.confirmPass
}
func (u *User) BreachEncapsulationOfStatus() UserStatus {
	return u.status
}

func (u *User) ReName(name UserName) error {
	u.name = name
	return nil
}

func (u *User) Activate() (ActivateEventBody, error) {
	u.status.Activate()
	return ActivateEventBody{
		UserId:     u.id,
		OccurredAt: time.Now(),
	}, nil
}

func (u *User) Authenticate(password UserPassword) bool {
	return u.password.Authenticate(password)
}

func (u *User) CompareConfirmPass(confirmPass UserConfirmPass) bool {
	return u.confirmPass.CompareConfirmPass(confirmPass)
}
