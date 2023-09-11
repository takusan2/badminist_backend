package player

import (
	"reflect"
	"testing"
)

func TestNewPlayerGender(t *testing.T) {
	type args struct {
		gender Gender
	}
	tests := []struct {
		name    string
		args    args
		want    PlayerGender
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				gender: Unknown,
			},
			want: PlayerGender{
				value: "unknown",
			},
			wantErr: false,
		},
		{
			name: "正常系",
			args: args{
				gender: Male,
			},
			want: PlayerGender{
				value: "male",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlayerGender(tt.args.gender)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlayerGender() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlayerGender() = %v, want %v", got, tt.want)
			}
		})
	}
}
