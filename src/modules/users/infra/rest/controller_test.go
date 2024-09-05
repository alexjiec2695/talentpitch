package restuser

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	domainuser "talentpitch/src/modules/users/domain"
	"testing"
)

type useCaseMock struct {
	mock.Mock
}

func (m *useCaseMock) CreateUser(user domainuser.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *useCaseMock) GetUserByID(ID string) (*domainuser.User, error) {
	args := m.Called(ID)
	return args.Get(0).(*domainuser.User), args.Error(1)
}

func (m *useCaseMock) Update(user domainuser.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *useCaseMock) DeleteByID(ID string) error {
	args := m.Called(ID)
	return args.Error(0)
}

func (m *useCaseMock) GetUsers(pageSize, offset int) ([]*domainuser.User, error) {
	args := m.Called(pageSize, offset)
	return args.Get(0).([]*domainuser.User), args.Error(1)
}

func TestGetUsersSuccessful(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	usecases.On("GetUsers", 10, 0).Return([]*domainuser.User{}, nil)

	instance.getUsers(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestGetUsersWithErrorWhenExecutedUseCase(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	usecases.On("GetUsers", 10, 0).Return([]*domainuser.User{}, errors.New("error service"))

	instance.getUsers(ginContext)
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestCreateUserSuccessful(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	req := domainuser.User{
		Name:  "TEST ",
		Email: "test@gmail.com",
	}
	usecases.On("CreateUser", req).Return(nil)

	bytes, _ := json.Marshal(req)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	httpMock := &http.Request{
		Body: r,
	}
	ginContext.Request = httpMock

	instance.createUser(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestCreateUserWithErrorsWhenExecutedUseCase(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	req := domainuser.User{
		Name:  "TEST ",
		Email: "test@gmail.com",
	}
	usecases.On("CreateUser", req).Return(errors.New("services error"))

	bytes, _ := json.Marshal(req)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	httpMock := &http.Request{
		Body: r,
	}
	ginContext.Request = httpMock

	instance.createUser(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestCreateUserWithErrorsWhenShouldBindJSONRequest(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	instance.createUser(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestGetUsersByIDSuccessful(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	usecases.On("GetUserByID", id).Return(&domainuser.User{}, nil)

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}

	instance.getUserByID(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestGetUsersByIDWithErrorWhenExecutedUseCase(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	usecases.On("GetUserByID", id).Return(&domainuser.User{}, errors.New("services error"))

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}

	instance.getUserByID(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestUpdateUserSuccessful(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}
	req := domainuser.User{
		ID:    id,
		Name:  "TEST ",
		Email: "test@gmail.com",
	}
	bytes, _ := json.Marshal(req)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	httpMock := &http.Request{
		Body: r,
	}
	ginContext.Request = httpMock

	usecases.On("Update", req).Return(nil)

	instance.updateUser(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestUpdateUserWithErrorWhenCallingUseCase(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}
	req := domainuser.User{
		ID:    id,
		Name:  "TEST ",
		Email: "test@gmail.com",
	}
	bytes, _ := json.Marshal(req)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	httpMock := &http.Request{
		Body: r,
	}
	ginContext.Request = httpMock

	usecases.On("Update", req).Return(errors.New("error services"))

	instance.updateUser(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestUpdateUserWithErrorWhenShouldBindJSONRequest(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	instance.updateUser(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestDeleteUserByIDSuccessful(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	usecases.On("DeleteByID", id).Return(nil)

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}

	instance.deleteUser(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestDeleteUserByIDWithErrorWhenExecutedUseCase(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	usecases.On("DeleteByID", id).Return(errors.New("error services"))

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}

	instance.deleteUser(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}
