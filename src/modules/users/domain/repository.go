package domainuser

type UserRepository interface {
	Create(user User) error
	GetUserByID(Id string) (*User, error)
	Update(userEntity User) error
	DeleteByID(Id string) error
	GetUsers() ([]*User, error)
}
