package member

import (
	"testing"

	"github.com/takuya-okada-01/badminist/app/command/domain/user"
)

func TestMember_ChangeRole(t *testing.T) {
	memberId := NewMemberId()
	userId, _ := user.UserIdFromStr(id)
	memberRole, _ := NewMemberRole(Admin)
	staffRole, _ := NewMemberRole(Staff)
	type fields struct {
		id     MemberId
		role   MemberRole
		userId user.UserId
	}
	type args struct {
		role MemberRole
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "正常系",
			fields: fields{
				id:     memberId,
				role:   memberRole,
				userId: userId,
			},
			args: args{
				role: staffRole,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Member{
				id:     tt.fields.id,
				role:   tt.fields.role,
				userId: tt.fields.userId,
			}
			m.ChangeRole(tt.args.role)
			if m.role != staffRole {
				t.Errorf("Member.ChangeRole() = %v, want %v", m.role, staffRole)
			}
		})
	}
}
