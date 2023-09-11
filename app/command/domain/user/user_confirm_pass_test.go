package user

import (
	"reflect"
	"testing"
)

var (
	confirmPass = NewUserConfirmPass()
)

func TestUserConfirmPassFromStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    UserConfirmPass
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				str: confirmPass.value,
			},
			want: UserConfirmPass{
				value: confirmPass.value,
			},
			wantErr: false,
		},
		{
			name: "異常系",
			args: args{
				str: "12345",
			},
			want:    UserConfirmPass{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UserConfirmPassFromStr(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserConfirmPassFromStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserConfirmPassFromStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
