package member

import "github.com/takuya-okada-01/badminist/api/domain/user"

type Members struct {
	value []Member
}

func NewMembers(userId user.UserId) Members {
	id := NewMemberId()
	role, _ := NewMemberRole(Admin)
	return Members{
		value: []Member{
			NewMember(
				id,
				role,
				userId,
			),
		},
	}
}

func MembersFromList(members []Member) Members {
	return Members{
		value: members,
	}
}

func (m *Members) AddMember(member Member) {
	m.value = append(m.value, member)
}

func (m *Members) RemoveMember(userId user.UserId) {
	for i, member := range m.value {
		if member.userId == userId {
			m.value = append(m.value[:i], m.value[i+1:]...)
		}
	}
}

func (m *Members) ChangeRole(userId user.UserId, role MemberRole) {
	for i, member := range m.value {
		if member.userId == userId {
			m.value[i].ChangeRole(role)
		}
	}
}

func (m *Members) IsMember(userId user.UserId) bool {
	member := m.GetMemberByUserId(userId)
	if member.id.value != "" {
		return true
	}
	return false
}

func (m *Members) IsAdmin(userId user.UserId) bool {
	admin, _ := NewMemberRole(Admin)
	return m.IsRole(userId, []MemberRole{admin})
}

func (m *Members) IsStaff(userId user.UserId) bool {
	admin, _ := NewMemberRole(Admin)
	staff, _ := NewMemberRole(Staff)
	return m.IsRole(userId, []MemberRole{admin, staff})
}

func (m *Members) IsRole(userId user.UserId, roles []MemberRole) bool {
	member := m.GetMemberByUserId(userId)
	for _, role := range roles {
		if member.role == role {
			return true
		}
	}
	return false
}

func (m *Members) GetMemberByUserId(userId user.UserId) Member {
	for _, member := range m.value {
		if member.userId == userId {
			return member
		}
	}
	return Member{}
}

func (m *Members) Value() []Member {
	return m.value
}
