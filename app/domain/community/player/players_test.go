package player

import (
	"reflect"
	"testing"
)

func TestPlayers_ChangePlayerAllProperty(t *testing.T) {
	playerId := NewPlayerId()
	playerName, _ := NewPlayerName("test")
	playerAge, _ := NewPlayerAge(0)
	playerGender, _ := NewPlayerGender(Unknown)
	playerLevel, _ := NewPlayerLevel(Beginner)
	playerNumGames, _ := NewPlayerNumGames(0)
	playerStatus, _ := NewPlayerStatus(Attend)
	changedPlayerName, _ := NewPlayerName("changed")
	type fields struct {
		value []Player
	}
	type args struct {
		changePlayer Player
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "正常系",
			fields: fields{
				value: []Player{
					NewPlayer(
						playerId,
						playerName,
						playerGender,
						playerAge,
						playerLevel,
						playerNumGames,
						playerStatus,
					),
				},
			},
			args: args{
				changePlayer: NewPlayer(
					playerId,
					changedPlayerName,
					playerGender,
					playerAge,
					playerLevel,
					playerNumGames,
					playerStatus,
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Players{
				value: tt.fields.value,
			}
			p.ChangePlayerProperty(tt.args.changePlayer)
			if !reflect.DeepEqual(p.value[0], tt.args.changePlayer) {
				t.Errorf("Players.ChangePlayer() = %v, want %v", p.value[0], tt.args.changePlayer)
			}
		})
	}
}

func TestPlayers_AddPlayer(t *testing.T) {
	playerId := NewPlayerId()
	playerName, _ := NewPlayerName("test")
	playerAge, _ := NewPlayerAge(0)
	playerLevel, _ := NewPlayerLevel(Beginner)
	playerNumGames, _ := NewPlayerNumGames(0)
	playerStatus, _ := NewPlayerStatus(Attend)
	playerGender, _ := NewPlayerGender(Unknown)
	type fields struct {
		value []Player
	}
	type args struct {
		player Player
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "正常系",
			fields: fields{
				value: []Player{},
			},
			args: args{
				player: NewPlayer(
					playerId,
					playerName,
					playerGender,
					playerAge,
					playerLevel,
					playerNumGames,
					playerStatus,
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Players{
				value: tt.fields.value,
			}
			p.AddPlayer(tt.args.player)
			if !reflect.DeepEqual(p.value[0], tt.args.player) {
				t.Errorf("Players.AddPlayer() = %v, want %v", p.value[0], tt.args.player)
			}
		})
	}
}

func TestPlayers_RemovePlayer(t *testing.T) {
	playerId := NewPlayerId()
	playerName, _ := NewPlayerName("test")
	playerAge, _ := NewPlayerAge(0)
	playerLevel, _ := NewPlayerLevel(Beginner)
	playerNumGames, _ := NewPlayerNumGames(0)
	playerStatus, _ := NewPlayerStatus(Attend)
	playerGender, _ := NewPlayerGender(Unknown)
	type fields struct {
		value []Player
	}
	type args struct {
		remoevPlayerId PlayerId
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "正常系",
			fields: fields{
				value: []Player{
					NewPlayer(
						playerId,
						playerName,
						playerGender,
						playerAge,
						playerLevel,
						playerNumGames,
						playerStatus,
					),
				},
			},
			args: args{
				remoevPlayerId: playerId,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Players{
				value: tt.fields.value,
			}
			p.RemovePlayer(tt.args.remoevPlayerId)
			if len(p.value) != 0 {
				t.Errorf("Players.RemovePlayer() = %v, want %v", p.value, []Player{})
			}
		})
	}
}
