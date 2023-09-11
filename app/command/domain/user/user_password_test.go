package user

import (
	"reflect"
	"testing"
)

func TestNewUserPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    UserPassword
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				password: "testpass",
			},
			want: UserPassword{
				value: "testpass",
			},
			wantErr: false,
		},
		{
			name: "異常系",
			args: args{
				password: "",
			},
			want:    UserPassword{},
			wantErr: true,
		},
		{
			name: "異常系",
			args: args{
				password: "     　　　",
			},
			want:    UserPassword{},
			wantErr: true,
		},
		{
			name: "異常系",
			args: args{
				password: "test",
			},
			want:    UserPassword{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
