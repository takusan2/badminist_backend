package user

import "time"

type CreateUserEventBody struct {
	UserId      UserId
	Email       UserEmail
	Password    UserPassword
	ConfirmPass UserConfirmPass
	Status      UserStatus
	OccurredAt  time.Time
}

type RenameUserEventBody struct {
	UserId     UserId
	NewName    UserName
	OccurredAt time.Time
}

type ActivateEventBody struct {
	UserId     UserId
	OccurredAt time.Time
}

type ReissueConfirmPassEventBody struct {
	UserId      UserId
	ConfirmPass UserConfirmPass
	OccurredAt  time.Time
}
