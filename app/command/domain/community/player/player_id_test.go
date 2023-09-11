package player

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

var (
	id = uuid.New().String()
)

func TestPlayerIdFrom(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    PlayerId
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				id: id,
			},
			want: PlayerId{
				value: id,
			},
			wantErr: false,
		},
		{
			name: "異常系",
			args: args{
				id: "invalid",
			},
			want:    PlayerId{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PlayerIdFromStr(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlayerIdFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlayerIdFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
