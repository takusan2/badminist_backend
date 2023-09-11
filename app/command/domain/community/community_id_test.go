package community

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

var (
	id = uuid.New().String()
)

func TestCommunityIdFrom(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    CommunityId
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				id: id,
			},
			want: CommunityId{
				value: id,
			},
			wantErr: false,
		},
		{
			name: "異常系",
			args: args{
				id: "invalid",
			},
			want:    CommunityId{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CommunityIdFromStr(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommunityIdFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommunityIdFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
