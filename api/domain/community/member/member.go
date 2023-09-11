package member

import "github.com/takuya-okada-01/badminist/api/domain/user"

type Member struct {
	id     MemberId
	role   MemberRole
	userId user.UserId
}

func NewMember(
	id MemberId,
	role MemberRole,
	userId user.UserId,
) Member {
	return Member{
		id:     id,
		role:   role,
		userId: userId,
	}
}

func (m *Member) BreachEncapsulationOfId() MemberId {
	return m.id
}
func (m *Member) BreachEncapsulationOfRole() MemberRole {
	return m.role
}
func (m *Member) BreachEncapsulationOfUserId() user.UserId {
	return m.userId
}

func (m *Member) ChangeRole(role MemberRole) {
	m.role = role
}
