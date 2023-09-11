package user

import (
	"reflect"
	"testing"
)

func TestNewUserEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    UserEmail
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				email: "hoge@hoge.com",
			},
			want: UserEmail{
				value: "hoge@hoge.com",
			},
			wantErr: false,
		},
		{
			name: "異常系",
			args: args{
				email: "hogehoge.com",
			},
			want:    UserEmail{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
