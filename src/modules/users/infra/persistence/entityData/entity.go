package entityData

import (
	"gorm.io/gorm"
	domainuser "talentpitch/src/modules/users/domain"
)

type User struct {
	gorm.Model
	ID    string
	Name  string
	Email string
}

func (u *User) ToEntity() *domainuser.User {
	return &domainuser.User{
		Name:  u.Name,
		Email: u.Email,
		ID:    u.ID,
	}
}
