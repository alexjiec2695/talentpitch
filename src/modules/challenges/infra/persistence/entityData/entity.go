package entityData

import (
	"gorm.io/gorm"
	"talentpitch/src/modules/challenges/domain"
)

type Challenges struct {
	gorm.Model
	ID          string
	Title       string
	Description string
}

func (c *Challenges) ToEntity() *domainchallenge.Challenges {
	return &domainchallenge.Challenges{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
	}
}
