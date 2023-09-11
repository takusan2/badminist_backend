package player

import (
	"reflect"
	"testing"
)

func TestNewPlayerStatus(t *testing.T) {
	type args struct {
		status Status
	}
	tests := []struct {
		name    string
		args    args
		want    PlayerStatus
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				status: Attend,
			},
			want: PlayerStatus{
				value: "attend",
			},
			wantErr: false,
		},
		{
			name: "正常系",
			args: args{
				status: Break,
			},
			want: PlayerStatus{
				value: "break",
			},
			wantErr: false,
		},
		{
			name: "正常系",
			args: args{
				status: Absence,
			},
			want: PlayerStatus{
				value: "absence",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPlayerStatus(tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPlayerStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlayerStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayerStatusFromStr(t *testing.T) {
	type args struct {
		status string
	}
	tests := []struct {
		name    string
		args    args
		want    PlayerStatus
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				status: "attend",
			},
			want: PlayerStatus{
				value: "attend",
			},
			wantErr: false,
		},
		{
			name: "正常系",
			args: args{
				status: "break",
			},
			want: PlayerStatus{
				value: "break",
			},
			wantErr: false,
		},
		{
			name: "正常系",
			args: args{
				status: "absence",
			},
			want: PlayerStatus{
				value: "absence",
			},
			wantErr: false,
		},
		{
			name: "正常系",
			args: args{
				status: "attend",
			},
			want: PlayerStatus{
				value: "attend",
			},
			wantErr: false,
		},
		{
			name: "異常系",
			args: args{
				status: "invalid",
			},
			want:    PlayerStatus{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PlayerStatusFromStr(tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlayerStatusFromStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlayerStatusFromStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
