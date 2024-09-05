package restchallenges

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
	domainchallenge "talentpitch/src/modules/challenges/domain"
	"testing"
)

type useCaseMock struct {
	mock.Mock
}

func (m *useCaseMock) CreateChallenges(challenges domainchallenge.Challenges) error {
	args := m.Called(challenges)
	return args.Error(0)
}

func (m *useCaseMock) GetChallengesByID(ID string) (*domainchallenge.Challenges, error) {
	args := m.Called(ID)
	return args.Get(0).(*domainchallenge.Challenges), args.Error(1)
}

func (m *useCaseMock) Update(challenges domainchallenge.Challenges) error {
	args := m.Called(challenges)
	return args.Error(0)
}

func (m *useCaseMock) DeleteByID(ID string) error {
	args := m.Called(ID)
	return args.Error(0)
}

func (m *useCaseMock) GetChallenges(pageSize, offset int) ([]*domainchallenge.Challenges, error) {
	args := m.Called(pageSize, offset)
	return args.Get(0).([]*domainchallenge.Challenges), args.Error(1)
}

func TestGetChallengessSuccessful(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	usecases.On("GetChallenges", 10, 0).Return([]*domainchallenge.Challenges{}, nil)

	instance.getChallenges(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestGetChallengessWithErrorWhenExecutedUseCase(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	usecases.On("GetChallenges", 10, 0).Return([]*domainchallenge.Challenges{}, errors.New("error service"))

	instance.getChallenges(ginContext)
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestCreateChallengesSuccessful(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	req := domainchallenge.Challenges{
		Title:       "titulo 1",
		Description: "descripcion test",
	}
	usecases.On("CreateChallenges", req).Return(nil)

	bytes, _ := json.Marshal(req)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	httpMock := &http.Request{
		Body: r,
	}
	ginContext.Request = httpMock

	instance.createChallenges(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestCreateChallengesWithErrorsWhenExecutedUseCase(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	req := domainchallenge.Challenges{
		Title:       "titulo 1",
		Description: "descripcion test",
	}
	usecases.On("CreateChallenges", req).Return(errors.New("services error"))

	bytes, _ := json.Marshal(req)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	httpMock := &http.Request{
		Body: r,
	}
	ginContext.Request = httpMock

	instance.createChallenges(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestCreateChallengesWithErrorsWhenShouldBindJSONRequest(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	instance.createChallenges(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestGetChallengesByIDSuccessful(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	usecases.On("GetChallengesByID", id).Return(&domainchallenge.Challenges{}, nil)

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}

	instance.getChallengesByID(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestGetChallengesByIDWithErrorWhenExecutedUseCase(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	usecases.On("GetChallengesByID", id).Return(&domainchallenge.Challenges{}, errors.New("services error"))

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}

	instance.getChallengesByID(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestUpdateChallengeSuccessful(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}
	req := domainchallenge.Challenges{
		ID:          id,
		Title:       "titulo 1",
		Description: "descripcion test",
	}
	bytes, _ := json.Marshal(req)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	httpMock := &http.Request{
		Body: r,
	}
	ginContext.Request = httpMock

	usecases.On("Update", req).Return(nil)

	instance.updateChallenges(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestUpdateChallengeWithErrorWhenCallingUseCase(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}
	req := domainchallenge.Challenges{
		ID:          id,
		Title:       "titulo 1",
		Description: "descripcion test",
	}
	bytes, _ := json.Marshal(req)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	httpMock := &http.Request{
		Body: r,
	}
	ginContext.Request = httpMock

	usecases.On("Update", req).Return(errors.New("error services"))

	instance.updateChallenges(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestUpdateChallengeWithErrorWhenShouldBindJSONRequest(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	instance.updateChallenges(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestDeleteChallengeByIDSuccessful(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	usecases.On("DeleteByID", id).Return(nil)

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}

	instance.deleteChallenges(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestDeleteChallengeByIDWithErrorWhenExecutedUseCase(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	usecases.On("DeleteByID", id).Return(errors.New("error services"))

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}

	instance.deleteChallenges(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}
