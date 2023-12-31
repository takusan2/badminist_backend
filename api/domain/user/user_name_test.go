package user

import (
	"reflect"
	"testing"
)

func TestNewUserName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    UserName
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				name: "test",
			},
			want: UserName{
				value: "test",
			},
			wantErr: false,
		},
		{
			name: "異常系",
			args: args{
				name: "",
			},
			want:    UserName{},
			wantErr: true,
		},
		{
			name: "異常系",
			args: args{
				name: "     　　　",
			},
			want:    UserName{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserName() = %v, want %v", got, tt.want)
			}
		})
	}
}
