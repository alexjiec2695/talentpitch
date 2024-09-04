package domainvideos

type useCase struct {
	repo VideosRepository
}

type UseCase interface {
	CreateVideo(video Videos) error
	GetVideoByID(ID string) (*Videos, error)
	Update(video Videos) error
	DeleteByID(ID string) error
	GetVideos() ([]*Videos, error)
}

func NewUseCase(repo VideosRepository) UseCase {
	return &useCase{
		repo: repo,
	}
}

func (u *useCase) CreateVideo(video Videos) error {
	return u.repo.Create(video)
}

func (u *useCase) GetVideoByID(ID string) (*Videos, error) {
	return u.repo.GetVideosByID(ID)
}

func (u *useCase) Update(video Videos) error {
	return u.repo.Update(video)
}

func (u *useCase) DeleteByID(ID string) error {
	return u.repo.DeleteByID(ID)
}

func (u *useCase) GetVideos() ([]*Videos, error) {
	return u.repo.GetVideos()
}
