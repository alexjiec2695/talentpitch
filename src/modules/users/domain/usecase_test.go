package domainuser_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	domainuser "talentpitch/src/modules/users/domain"
	"testing"
)

type userRepositoryMock struct {
	mock.Mock
}

func (m *userRepositoryMock) Create(user domainuser.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *userRepositoryMock) GetUserByID(Id string) (*domainuser.User, error) {
	args := m.Called(Id)
	return args.Get(0).(*domainuser.User), args.Error(1)
}

func (m *userRepositoryMock) Update(user domainuser.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *userRepositoryMock) DeleteByID(Id string) error {
	args := m.Called(Id)
	return args.Error(0)
}

func (m *userRepositoryMock) GetUsers(pageSize, offset int) ([]*domainuser.User, error) {
	args := m.Called(pageSize, offset)
	return args.Get(0).([]*domainuser.User), args.Error(1)
}

func (m *userRepositoryMock) MassiveCreate() {
	return
}

func TestUseCase_CreateUser(t *testing.T) {
	repository := new(userRepositoryMock)
	useCaseInstance := domainuser.NewUseCase(repository)
	req := domainuser.User{}

	repository.On("Create", req).Return(nil)

	err := useCaseInstance.CreateUser(req)

	assert.NoError(t, err)
}

func TestUseCase_GetUserByID(t *testing.T) {
	repository := new(userRepositoryMock)
	useCaseInstance := domainuser.NewUseCase(repository)
	id := "12346789"
	expect := &domainuser.User{}

	repository.On("GetUserByID", id).Return(expect, nil)

	users, err := useCaseInstance.GetUserByID(id)

	assert.NoError(t, err)
	assert.Equal(t, users, users)
}

func TestUseCase_UpdateUser(t *testing.T) {
	repository := new(userRepositoryMock)
	useCaseInstance := domainuser.NewUseCase(repository)
	req := domainuser.User{}

	repository.On("Update", req).Return(nil)

	err := useCaseInstance.Update(req)

	assert.NoError(t, err)
}

func TestUseCase_DeleteByIDUser(t *testing.T) {
	repository := new(userRepositoryMock)
	useCaseInstance := domainuser.NewUseCase(repository)
	id := "1234565677"

	repository.On("DeleteByID", id).Return(nil)

	err := useCaseInstance.DeleteByID(id)

	assert.NoError(t, err)
}

func TestUseCase_Getusers(t *testing.T) {
	repository := new(userRepositoryMock)
	useCaseInstance := domainuser.NewUseCase(repository)
	expect := []*domainuser.User{}

	repository.On("GetUsers", 10, 0).Return(expect, nil)

	users, err := useCaseInstance.GetUsers(10, 0)

	assert.NoError(t, err)
	assert.Equal(t, expect, users)
}
