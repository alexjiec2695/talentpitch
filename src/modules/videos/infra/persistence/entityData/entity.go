package entityData

import (
	"gorm.io/gorm"
	"talentpitch/src/modules/videos/domain"
)

type Videos struct {
	gorm.Model
	ID   string
	Name string
	Url  string
}

func (u *Videos) ToEntity() *domainvideos.Videos {
	return &domainvideos.Videos{
		Name: u.Name,
		Url:  u.Url,
		ID:   u.ID,
	}
}
