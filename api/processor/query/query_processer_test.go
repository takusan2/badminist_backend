package query_processor

import (
	"reflect"
	"testing"

	query_dao_if "github.com/takuya-okada-01/badminist/api/interface_adaptor_if/dao_if/query"
	"github.com/takuya-okada-01/badminist/api/processor/query/read_model"
	"gorm.io/gorm"
)

func Test_queryProcessor_generateSinglesMatchCombination(t *testing.T) {
	type fields struct {
		db           *gorm.DB
		communityDao query_dao_if.CommunityDao
		userDao      query_dao_if.UserDao
	}
	type args struct {
		players  read_model.PlayerList
		numCourt int
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantMatchLength int
		wantRestLength  int
		wantErr         bool
	}{
		{
			name: "success",
			fields: fields{
				db:           nil,
				communityDao: nil,
				userDao:      nil,
			},
			args: args{
				players: read_model.PlayerList{
					{
						ID:   "1",
						Name: "player1",
					},
					{
						ID:   "2",
						Name: "player2",
					},
					{
						ID:   "3",
						Name: "player3",
					},
					{
						ID:   "4",
						Name: "player4",
					},
				},
				numCourt: 1,
			},
			wantMatchLength: 1,
			wantRestLength:  2,
			wantErr:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queryProcessor{
				db:           tt.fields.db,
				communityDao: tt.fields.communityDao,
				userDao:      tt.fields.userDao,
			}
			got, got1, err := q.generateSinglesMatchCombination(tt.args.players, tt.args.numCourt)
			if (err != nil) != tt.wantErr {
				t.Errorf("queryProcessor.generateSinglesMatchCombination() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.wantMatchLength) {
				t.Errorf("queryProcessor.generateSinglesMatchCombination() got = %v, want %v", len(got), tt.wantMatchLength)
			}
			if !reflect.DeepEqual(len(got1), tt.wantRestLength) {
				t.Errorf("queryProcessor.generateSinglesMatchCombination() got1 = %v, want %v", len(got1), tt.wantRestLength)
			}
		})
	}
}
