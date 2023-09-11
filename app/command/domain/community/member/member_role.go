package member

import "errors"

type MemberRole struct {
	value string
}

type Role int

const (
	Admin Role = iota
	Staff
	General
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "admin"
	case Staff:
		return "staff"
	default:
		return "general"
	}
}

func NewMemberRole(role Role) (MemberRole, error) {
	return MemberRole{role.String()}, nil
}

func MemberRoleFromStr(role string) (MemberRole, error) {
	switch role {
	case Admin.String():
		return NewMemberRole(Admin)
	case Staff.String():
		return NewMemberRole(Staff)
	case General.String():
		return NewMemberRole(General)
	default:
		role, _ := NewMemberRole(General)
		return role, errors.New("invalid role")
	}
}

func (m MemberRole) Value() string {
	return m.value
}
