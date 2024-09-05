package domainuser

type useCase struct {
	repo UserRepository
}

type UseCase interface {
	CreateUser(user User) error
	GetUserByID(ID string) (*User, error)
	Update(user User) error
	DeleteByID(ID string) error
	GetUsers(pageSize, offset int) ([]*User, error)
}

func NewUseCase(repo UserRepository) UseCase {
	return &useCase{
		repo: repo,
	}
}

func (u *useCase) CreateUser(user User) error {
	return u.repo.Create(user)
}

func (u *useCase) GetUserByID(ID string) (*User, error) {
	return u.repo.GetUserByID(ID)
}

func (u *useCase) Update(user User) error {
	return u.repo.Update(user)
}

func (u *useCase) DeleteByID(ID string) error {
	return u.repo.DeleteByID(ID)
}

func (u *useCase) GetUsers(pageSize, offset int) ([]*User, error) {
	return u.repo.GetUsers(pageSize, offset)
}
