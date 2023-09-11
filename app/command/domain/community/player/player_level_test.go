package player

import (
	"reflect"
	"testing"
)

func TestNewPlayerLevel(t *testing.T) {
	type args struct {
		level Level
	}
	tests := []struct {
		name    string
		args    args
		want    PlayerLevel
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				level: Beginner,
			},
			want: PlayerLevel{
				value: "beginner",
			},
			wantErr: false,
		},
		{

			name: "正常系",
			args: args{
				level: Intermediate,
			},
			want: PlayerLevel{
				value: "intermediate",
			},
			wantErr: false,
		},
		{
			name: "正常系",
			args: args{
				level: Advanced,
			},
			want: PlayerLevel{
				value: "advanced",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlayerLevel(tt.args.level)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlayerLevel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlayerLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayerLevelFromStr(t *testing.T) {
	type args struct {
		level string
	}
	tests := []struct {
		name    string
		args    args
		want    PlayerLevel
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				level: "beginner",
			},
			want: PlayerLevel{
				value: "beginner",
			},
			wantErr: false,
		},
		{
			name: "正常系",
			args: args{
				level: "intermediate",
			},
			want: PlayerLevel{
				value: "intermediate",
			},
			wantErr: false,
		},
		{
			name: "正常系",
			args: args{
				level: "advanced",
			},
			want: PlayerLevel{
				value: "advanced",
			},
			wantErr: false,
		},
		{
			name: "異常系",
			args: args{
				level: "test",
			},
			want:    PlayerLevel{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PlayerLevelFromStr(tt.args.level)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlayerLevelFromStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlayerLevelFromStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
