package user

import "errors"

type UserStatus struct {
	value string
}

type Status int

const (
	Active Status = iota
	Inactive
)

func (s Status) String() string {
	switch s {
	case Active:
		return "active"
	default:
		return "inactive"
	}
}

func NewUserStatus(status Status) (UserStatus, error) {
	return UserStatus{status.String()}, nil
}

func UserStatusFromStr(status string) (UserStatus, error) {
	switch status {
	case Active.String():
		return NewUserStatus(Active)
	case Inactive.String():
		return NewUserStatus(Inactive)
	default:
		status, _ := NewUserStatus(Inactive)
		return status, errors.New("invalid status")
	}
}

func (u UserStatus) Value() string {
	return u.value
}

func (u *UserStatus) Activate() {
	u.value = Active.String()
}
