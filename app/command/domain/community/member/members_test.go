package member

import (
	"testing"

	"github.com/takuya-okada-01/badminist/app/command/domain/user"
)

func TestMembers_AddMember(t *testing.T) {
	memberId := NewMemberId()
	userId, _ := user.UserIdFromStr("test")
	memberRole, _ := NewMemberRole(Admin)

	type fields struct {
		value []Member
	}
	type args struct {
		member Member
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "正常系",
			fields: fields{
				value: []Member{},
			},
			args: args{
				member: NewMember(
					memberId,
					memberRole,
					userId,
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Members{
				value: tt.fields.value,
			}
			m.AddMember(tt.args.member)
			if len(m.value) != 1 {
				t.Errorf("Members.AddMember() = %v, want %v", len(m.value), 1)
			}
		})
	}
}

func TestMembers_RemoveMember(t *testing.T) {
	memberId := NewMemberId()
	userId, _ := user.UserIdFromStr("test")
	memberRole, _ := NewMemberRole(Admin)

	type fields struct {
		value []Member
	}
	type args struct {
		userId user.UserId
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "正常系",
			fields: fields{
				value: []Member{
					NewMember(
						memberId,
						memberRole,
						userId,
					),
				},
			},
			args: args{
				userId: userId,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Members{
				value: tt.fields.value,
			}
			m.RemoveMember(tt.args.userId)
			if len(m.value) != 0 {
				t.Errorf("Members.RemoveMember() = %v, want %v", len(m.value), 0)
			}
		})
	}
}

func TestMembers_ChangeRole(t *testing.T) {
	memberId := NewMemberId()
	userId, _ := user.UserIdFromStr("test")
	memberRole, _ := NewMemberRole(Admin)
	staff, _ := NewMemberRole(Staff)
	type fields struct {
		value []Member
	}
	type args struct {
		userId user.UserId
		role   MemberRole
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "正常系",
			fields: fields{
				value: []Member{
					NewMember(
						memberId,
						memberRole,
						userId,
					),
				},
			},
			args: args{
				userId: userId,
				role:   staff,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Members{
				value: tt.fields.value,
			}
			m.ChangeRole(tt.args.userId, tt.args.role)
			if m.value[0].role != staff {
				t.Errorf("Members.ChangeRole() = %v, want %v", m.value[0].role, staff)
			}
		})
	}
}

func TestMembers_IsMember(t *testing.T) {
	memberId := NewMemberId()
	userId, _ := user.UserIdFromStr("test")
	admin, _ := NewMemberRole(Admin)

	type fields struct {
		value []Member
	}
	type args struct {
		userId user.UserId
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "正常系",
			fields: fields{
				value: []Member{
					NewMember(
						memberId,
						admin,
						userId,
					),
				},
			},
			args: args{
				userId: userId,
			},
			want: true,
		},
		{
			name: "正常系",
			fields: fields{
				value: []Member{
					NewMember(
						memberId,
						admin,
						userId,
					),
				},
			},
			args: args{
				userId: user.UserId{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Members{
				value: tt.fields.value,
			}
			if got := m.IsMember(tt.args.userId); got != tt.want {
				t.Errorf("Members.IsMember() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMembers_IsAdmin(t *testing.T) {
	memberId := NewMemberId()
	userId, _ := user.UserIdFromStr("test")
	admin, _ := NewMemberRole(Admin)
	general, _ := NewMemberRole(General)
	type fields struct {
		value []Member
	}
	type args struct {
		userId user.UserId
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "正常系",
			fields: fields{
				value: []Member{
					NewMember(
						memberId,
						admin,
						userId,
					),
				},
			},
			args: args{
				userId: userId,
			},
			want: true,
		},
		{
			name: "正常系",
			fields: fields{
				value: []Member{
					NewMember(
						memberId,
						general,
						userId,
					),
				},
			},
			args: args{
				userId: userId,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Members{
				value: tt.fields.value,
			}
			if got := m.IsAdmin(tt.args.userId); got != tt.want {
				t.Errorf("Members.IsAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMembers_IsStaff(t *testing.T) {
	memberId := NewMemberId()
	userId, _ := user.UserIdFromStr("test")
	staff, _ := NewMemberRole(Admin)
	general, _ := NewMemberRole(General)

	type fields struct {
		value []Member
	}
	type args struct {
		userId user.UserId
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "正常系",
			fields: fields{
				value: []Member{
					NewMember(
						memberId,
						staff,
						userId,
					),
				},
			},
			args: args{
				userId: userId,
			},
			want: true,
		},
		{
			name: "正常系",
			fields: fields{
				value: []Member{
					NewMember(
						memberId,
						general,
						userId,
					),
				},
			},
			args: args{
				userId: userId,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Members{
				value: tt.fields.value,
			}
			if got := m.IsStaff(tt.args.userId); got != tt.want {
				t.Errorf("Members.IsStaff() = %v, want %v", got, tt.want)
			}
		})
	}
}
