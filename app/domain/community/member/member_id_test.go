package member

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

var (
	id = uuid.New().String()
)

func TestMemberIdFrom(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    MemberId
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				id: id,
			},
			want: MemberId{
				value: id,
			},
			wantErr: false,
		},
		{
			name: "異常系",
			args: args{
				id: "invalid",
			},
			want:    MemberId{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MemberIdFromStr(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("MemberIdFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MemberIdFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
