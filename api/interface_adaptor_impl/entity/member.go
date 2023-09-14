package entity

import (
	"time"

	"github.com/takuya-okada-01/badminist/api/domain/community/member"
	"github.com/takuya-okada-01/badminist/api/domain/user"
)

type Member struct {
	Id          string    `gorm:"primary_key auto_increment;"`
	UserId      string    `gorm:"type:varchar(36);not null;"`
	User        User      `gorm:"foreignKey:UserId"`
	CommunityId string    `gorm:"type:varchar(36);not null;"`
	Community   Community `gorm:"foreignKey:CommunityId"`
	Role        string    `gorm:"type:varchar(255);not null;"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
	UpdatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}

func NewMember(
	id string,
	userId string,
	communityId string,
	role string,
) Member {
	return Member{
		Id:          id,
		UserId:      userId,
		CommunityId: communityId,
		Role:        role,
	}
}

func (m *Member) ToDomainObject() member.Member {
	memberId, err := member.MemberIdFromStr(m.Id)
	if err != nil {
		panic(err)
	}

	userId, err := user.UserIdFromStr(m.UserId)
	if err != nil {
		panic(err)
	}

	role, err := member.MemberRoleFromStr(m.Role)
	if err != nil {
		panic(err)
	}

	return member.NewMember(memberId, role, userId)
}
