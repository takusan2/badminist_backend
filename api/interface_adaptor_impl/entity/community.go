package entity

import (
	"time"

	"github.com/takuya-okada-01/badminist/api/processor/query/read_model"
)

type Community struct {
	Id          string    `gorm:"type:varchar(36);primaryKey;"`
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

func (c *Community) ToReadModel() read_model.Community {
	return read_model.Community{
		ID:          c.Id,
		Name:        c.Name,
		Description: c.Description,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}
