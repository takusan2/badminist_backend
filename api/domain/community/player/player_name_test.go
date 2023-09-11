package player

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewPlayerName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    PlayerName
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				name: "test",
			},
			want: PlayerName{
				value: "test",
			},
			wantErr: false,
		},
		{
			name: "異常系",
			args: args{
				name: "",
			},
			want:    PlayerName{},
			wantErr: true,
		},
		{
			name: "異常系",
			args: args{
				name: "     　　　",
			},
			want:    PlayerName{},
			wantErr: true,
		},
		{
			name: "異常系",
			args: args{
				name: strings.Repeat("a", 256),
			},
			want:    PlayerName{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlayerName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlayerName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlayerName() = %v, want %v", got, tt.want)
			}
		})
	}
}
