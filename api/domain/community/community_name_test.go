package community

import (
	"reflect"
	"testing"
)

func TestNewCommunityName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    CommunityName
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				name: "test",
			},
			want: CommunityName{
				value: "test",
			},
			wantErr: false,
		},
		{
			name: "異常系",
			args: args{
				name: "",
			},
			want:    CommunityName{},
			wantErr: true,
		},
		{
			name: "異常系",
			args: args{
				name: "     　　　",
			},
			want:    CommunityName{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCommunityName(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCommunityName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCommunityName() = %v, want %v", got, tt.want)
			}
		})
	}
}
