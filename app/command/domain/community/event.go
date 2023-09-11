package community

import (
	"time"

	"github.com/takuya-okada-01/badminist/app/command/domain/community/member"
	"github.com/takuya-okada-01/badminist/app/command/domain/community/player"
	"github.com/takuya-okada-01/badminist/app/command/domain/user"
)

// CommunityEventBody
type CommunityEventCreatedBody struct {
	CommunityId CommunityId
	Name        CommunityName
	Description CommunityDescription
	Players     player.Players
	Members     member.Members
	OccurredAt  time.Time
}

// CommunityEventDeletedBody
type CommunityEventDeletedBody struct {
	CommunityId CommunityId
	OccurredAt  time.Time
}

// CommunityEventRenamedBody
type CommunityEventRenamedBody struct {
	CommunityId CommunityId
	Name        CommunityName
	OccurredAt  time.Time
}

// CommunityEventEditDescriptionBody
type CommunityEventEditDescriptionBody struct {
	CommunityId CommunityId
	Description CommunityDescription
	OccurredAt  time.Time
}

// CommunityEventAddMemberBody
type CommunityEventAddMemberBody struct {
	CommunityId CommunityId
	Member      member.Member
	OccurredAt  time.Time
}

// CommunityEventRemoveMemberBody
type CommunityEventRemoveMemberBody struct {
	CommunityId CommunityId
	UserId      user.UserId
	OccurredAt  time.Time
}

// CommunityEventChangeMemberRoleBody
type CommunityEventChangeMemberRoleBody struct {
	CommunityId CommunityId
	Member      member.Member
	OccurredAt  time.Time
}

// CommunityEventRemovePlayerBody
type CommunityEventAddPlayerBody struct {
	CommunityId CommunityId
	Player      player.Player
	OccurredAt  time.Time
}

// CommunityEventRemovePlayerBody
type CommunityEventRemovePlayerBody struct {
	CommunityId CommunityId
	PlayerId    player.PlayerId
	OccurredAt  time.Time
}

// CommunityEventChangePlayerPropertyBody
type CommunityEventChangePlayerPropertyBody struct {
	CommunityId CommunityId
	Player      player.Player
	OccurredAt  time.Time
}

// CommunityEventResetNumGamseBody
type CommunityEventResetNumGamseBody struct {
	CommunityId CommunityId
	PlayerId    player.PlayerId
	OccurredAt  time.Time
}
