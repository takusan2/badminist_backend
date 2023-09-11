package member

import (
	"reflect"
	"testing"
)

func TestNewMemberRole(t *testing.T) {
	type args struct {
		role Role
	}
	tests := []struct {
		name    string
		args    args
		want    MemberRole
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				role: Admin,
			},
			want: MemberRole{
				value: Admin.String(),
			},
			wantErr: false,
		},
		{
			name: "正常系",
			args: args{
				role: Staff,
			},
			want: MemberRole{
				value: Staff.String(),
			},
			wantErr: false,
		},
		{
			name: "正常系",
			args: args{
				role: General,
			},
			want: MemberRole{
				value: General.String(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMemberRole(tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMemberRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemberRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemberRoleFromString(t *testing.T) {
	type args struct {
		role string
	}
	tests := []struct {
		name    string
		args    args
		want    MemberRole
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				role: Admin.String(),
			},
			want: MemberRole{
				value: Admin.String(),
			},
			wantErr: false,
		},
		{
			name: "正常系",
			args: args{
				role: Staff.String(),
			},
			want: MemberRole{
				value: Staff.String(),
			},
			wantErr: false,
		},
		{
			name: "正常系",
			args: args{
				role: General.String(),
			},
			want: MemberRole{
				value: General.String(),
			},
			wantErr: false,
		},
		{
			name: "異常系",
			args: args{
				role: "invalid",
			},
			want: MemberRole{
				value: General.String(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MemberRoleFromStr(tt.args.role)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemberRoleFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemberRoleFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
