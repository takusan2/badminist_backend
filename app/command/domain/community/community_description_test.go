package community

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewCommunityDescription(t *testing.T) {
	type args struct {
		description string
	}
	tests := []struct {
		name    string
		args    args
		want    CommunityDescription
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				description: "test",
			},
			want: CommunityDescription{
				value: "test",
			},
			wantErr: false,
		},
		{
			name: "異常系",
			args: args{
				description: strings.Repeat("a", 256),
			},
			want:    CommunityDescription{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCommunityDescription(tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCommunityDescription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommunityDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}
