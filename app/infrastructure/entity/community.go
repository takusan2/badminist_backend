package entity

import (
	"time"
)

type Community struct {
	Id          string    `gorm:"type:varchar(36);primary_key;"`
	Name        string    `gorm:"type:varchar(255);not null;"`
	Description string    `gorm:"type:text;"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
	UpdatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP;"`
}

func NewCommunity(
	id string,
	name string,
	description string,
) Community {
	return Community{
		Id:          id,
		Name:        name,
		Description: description,
	}
}
