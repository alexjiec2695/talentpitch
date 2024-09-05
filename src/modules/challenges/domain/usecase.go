package domainchallenge

type useCase struct {
	repo ChallengesRepository
}

type UseCase interface {
	CreateChallenges(challenges Challenges) error
	GetChallengesByID(ID string) (*Challenges, error)
	Update(challenges Challenges) error
	DeleteByID(ID string) error
	GetChallenges(pageSize, offset int) ([]*Challenges, error)
}

func NewUseCase(repo ChallengesRepository) UseCase {
	return &useCase{
		repo: repo,
	}
}

func (u *useCase) CreateChallenges(challenge Challenges) error {
	return u.repo.Create(challenge)
}

func (u *useCase) GetChallengesByID(ID string) (*Challenges, error) {
	return u.repo.GetChallengesByID(ID)
}

func (u *useCase) Update(challenge Challenges) error {
	return u.repo.Update(challenge)
}

func (u *useCase) DeleteByID(ID string) error {
	return u.repo.DeleteByID(ID)
}

func (u *useCase) GetChallenges(pageSize, offset int) ([]*Challenges, error) {
	return u.repo.GetChallenges(pageSize, offset)
}
