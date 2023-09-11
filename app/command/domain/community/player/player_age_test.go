package player

import (
	"reflect"
	"testing"
)

func TestNewPlayerAge(t *testing.T) {
	type args struct {
		age int
	}
	tests := []struct {
		name    string
		args    args
		want    PlayerAge
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				age: 0,
			},
			want: PlayerAge{
				value: 0,
			},
			wantErr: false,
		},
		{
			name: "正常系",
			args: args{
				age: 150,
			},
			want: PlayerAge{
				value: 150,
			},
			wantErr: false,
		},
		{
			name: "異常系",
			args: args{
				age: -1,
			},
			want:    PlayerAge{},
			wantErr: true,
		},
		{
			name: "異常系",
			args: args{
				age: 151,
			},
			want:    PlayerAge{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlayerAge(tt.args.age)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlayerAge() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlayerAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
