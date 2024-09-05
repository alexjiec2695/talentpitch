package domainchallenge_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	domainchallenge "talentpitch/src/modules/challenges/domain"
	"testing"
)

type challengesRepositoryMock struct {
	mock.Mock
}

func (m *challengesRepositoryMock) Create(challenges domainchallenge.Challenges) error {
	args := m.Called(challenges)
	return args.Error(0)
}

func (m *challengesRepositoryMock) GetChallengesByID(Id string) (*domainchallenge.Challenges, error) {
	args := m.Called(Id)
	return args.Get(0).(*domainchallenge.Challenges), args.Error(1)
}

func (m *challengesRepositoryMock) Update(challenges domainchallenge.Challenges) error {
	args := m.Called(challenges)
	return args.Error(0)
}

func (m *challengesRepositoryMock) DeleteByID(Id string) error {
	args := m.Called(Id)
	return args.Error(0)
}

func (m *challengesRepositoryMock) GetChallenges(pageSize, offset int) ([]*domainchallenge.Challenges, error) {
	args := m.Called(pageSize, offset)
	return args.Get(0).([]*domainchallenge.Challenges), args.Error(1)
}

func (m *challengesRepositoryMock) MassiveCreate() {
	return
}

func TestUseCase_CreateChallenge(t *testing.T) {
	repository := new(challengesRepositoryMock)
	useCaseInstance := domainchallenge.NewUseCase(repository)
	req := domainchallenge.Challenges{}

	repository.On("Create", req).Return(nil)

	err := useCaseInstance.CreateChallenges(req)

	assert.NoError(t, err)
}

func TestUseCase_GetChallengeByID(t *testing.T) {
	repository := new(challengesRepositoryMock)
	useCaseInstance := domainchallenge.NewUseCase(repository)
	id := "12346789"
	expect := &domainchallenge.Challenges{}

	repository.On("GetChallengesByID", id).Return(expect, nil)

	challenge, err := useCaseInstance.GetChallengesByID(id)

	assert.NoError(t, err)
	assert.Equal(t, challenge, challenge)
}

func TestUseCase_UpdateChallenge(t *testing.T) {
	repository := new(challengesRepositoryMock)
	useCaseInstance := domainchallenge.NewUseCase(repository)
	req := domainchallenge.Challenges{}

	repository.On("Update", req).Return(nil)

	err := useCaseInstance.Update(req)

	assert.NoError(t, err)
}

func TestUseCase_DeleteByIDChallenge(t *testing.T) {
	repository := new(challengesRepositoryMock)
	useCaseInstance := domainchallenge.NewUseCase(repository)
	id := "1234565677"

	repository.On("DeleteByID", id).Return(nil)

	err := useCaseInstance.DeleteByID(id)

	assert.NoError(t, err)
}

func TestUseCase_Getchallenge(t *testing.T) {
	repository := new(challengesRepositoryMock)
	useCaseInstance := domainchallenge.NewUseCase(repository)
	expect := []*domainchallenge.Challenges{}

	repository.On("GetChallenges", 10, 0).Return(expect, nil)

	challenge, err := useCaseInstance.GetChallenges(10, 0)

	assert.NoError(t, err)
	assert.Equal(t, expect, challenge)
}
