package user

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

var (
	userId = uuid.New().String()
)

func TestUserIdFromStr(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    UserId
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				id: userId,
			},
			want:    UserId{userId},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UserIdFromStr(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserIdFromStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserIdFromStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
