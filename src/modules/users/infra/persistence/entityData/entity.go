package entityData

import (
	"gorm.io/gorm"
	"talentpitch/src/modules/users/domain"
)

type User struct {
	gorm.Model
	ID    string
	Name  string
	Email string
}

func (u *User) ToEntity() *domain.User {
	return &domain.User{
		Name:  u.Name,
		Email: u.Email,
		ID:    u.ID,
	}
}
