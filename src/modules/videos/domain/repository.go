package domainvideos

type VideosRepository interface {
	Create(videos Videos) error
	GetVideosByID(Id string) (*Videos, error)
	Update(videosEntity Videos) error
	DeleteByID(Id string) error
	GetVideos() ([]*Videos, error)
	MassiveCreate()
}
