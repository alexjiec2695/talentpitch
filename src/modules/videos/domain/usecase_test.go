package domainvideos_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	domainvideos "talentpitch/src/modules/videos/domain"
	"testing"
)

type videosRepositoryMock struct {
	mock.Mock
}

func (m *videosRepositoryMock) Create(videos domainvideos.Videos) error {
	args := m.Called(videos)
	return args.Error(0)
}

func (m *videosRepositoryMock) GetVideosByID(Id string) (*domainvideos.Videos, error) {
	args := m.Called(Id)
	return args.Get(0).(*domainvideos.Videos), args.Error(1)
}

func (m *videosRepositoryMock) Update(videosEntity domainvideos.Videos) error {
	args := m.Called(videosEntity)
	return args.Error(0)
}

func (m *videosRepositoryMock) DeleteByID(Id string) error {
	args := m.Called(Id)
	return args.Error(0)
}

func (m *videosRepositoryMock) GetVideos(pageSize, offset int) ([]*domainvideos.Videos, error) {
	args := m.Called(pageSize, offset)
	return args.Get(0).([]*domainvideos.Videos), args.Error(1)
}

func (m *videosRepositoryMock) MassiveCreate() {
	return
}

func TestUseCase_CreateVideo(t *testing.T) {
	repository := new(videosRepositoryMock)
	useCaseInstance := domainvideos.NewUseCase(repository)
	req := domainvideos.Videos{}

	repository.On("Create", req).Return(nil)

	err := useCaseInstance.CreateVideo(req)

	assert.NoError(t, err)
}

func TestUseCase_GetVideoByID(t *testing.T) {
	repository := new(videosRepositoryMock)
	useCaseInstance := domainvideos.NewUseCase(repository)
	id := "12346789"
	expect := &domainvideos.Videos{}

	repository.On("GetVideosByID", id).Return(expect, nil)

	videos, err := useCaseInstance.GetVideoByID(id)

	assert.NoError(t, err)
	assert.Equal(t, videos, videos)
}

func TestUseCase_UpdateVideo(t *testing.T) {
	repository := new(videosRepositoryMock)
	useCaseInstance := domainvideos.NewUseCase(repository)
	req := domainvideos.Videos{}

	repository.On("Update", req).Return(nil)

	err := useCaseInstance.Update(req)

	assert.NoError(t, err)
}

func TestUseCase_DeleteByIDVideo(t *testing.T) {
	repository := new(videosRepositoryMock)
	useCaseInstance := domainvideos.NewUseCase(repository)
	id := "1234565677"

	repository.On("DeleteByID", id).Return(nil)

	err := useCaseInstance.DeleteByID(id)

	assert.NoError(t, err)
}

func TestUseCase_GetVideos(t *testing.T) {
	repository := new(videosRepositoryMock)
	useCaseInstance := domainvideos.NewUseCase(repository)
	expect := []*domainvideos.Videos{}

	repository.On("GetVideos", 10, 0).Return(expect, nil)

	videos, err := useCaseInstance.GetVideos(10, 0)

	assert.NoError(t, err)
	assert.Equal(t, expect, videos)
}
