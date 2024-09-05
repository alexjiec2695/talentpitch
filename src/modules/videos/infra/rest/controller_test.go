package restvideo

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
	domainvideos "talentpitch/src/modules/videos/domain"
	"testing"
)

type useCaseMock struct {
	mock.Mock
}

func (m *useCaseMock) CreateVideo(video domainvideos.Videos) error {
	args := m.Called(video)
	return args.Error(0)
}

func (m *useCaseMock) GetVideoByID(ID string) (*domainvideos.Videos, error) {
	args := m.Called(ID)
	return args.Get(0).(*domainvideos.Videos), args.Error(1)
}

func (m *useCaseMock) Update(video domainvideos.Videos) error {
	args := m.Called(video)
	return args.Error(0)
}

func (m *useCaseMock) DeleteByID(ID string) error {
	args := m.Called(ID)
	return args.Error(0)
}

func (m *useCaseMock) GetVideos(pageSize, offset int) ([]*domainvideos.Videos, error) {
	args := m.Called(pageSize, offset)
	return args.Get(0).([]*domainvideos.Videos), args.Error(1)
}

func TestGetVideoSuccessful(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	usecases.On("GetVideos", 10, 0).Return([]*domainvideos.Videos{}, nil)

	instance.getVideo(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestGetVideoWithErrorWhenExecutedUseCase(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	usecases.On("GetVideos", 10, 0).Return([]*domainvideos.Videos{}, errors.New("error service"))

	instance.getVideo(ginContext)
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestCreateVideoSuccessful(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	req := domainvideos.Videos{
		Name: "TEST ",
		Url:  "www.test.com",
	}
	usecases.On("CreateVideo", req).Return(nil)

	bytes, _ := json.Marshal(req)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	httpMock := &http.Request{
		Body: r,
	}
	ginContext.Request = httpMock

	instance.createVideo(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestCreateVideoWithErrorsWhenExecutedUseCase(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	req := domainvideos.Videos{
		Name: "TEST ",
		Url:  "www.test.com",
	}
	usecases.On("CreateVideo", req).Return(errors.New("services error"))

	bytes, _ := json.Marshal(req)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	httpMock := &http.Request{
		Body: r,
	}
	ginContext.Request = httpMock

	instance.createVideo(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestCreateVideoWithErrorsWhenShouldBindJSONRequest(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	instance.createVideo(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestGetVideoByIDSuccessful(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	usecases.On("GetVideoByID", id).Return(&domainvideos.Videos{}, nil)

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}

	instance.getVideoByID(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestGetVideoByIDWithErrorWhenExecutedUseCase(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	usecases.On("GetVideoByID", id).Return(&domainvideos.Videos{}, errors.New("services error"))

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}

	instance.getVideoByID(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestUpdateVideoSuccessful(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}
	req := domainvideos.Videos{
		ID:   id,
		Name: "TEST ",
		Url:  "www.test.com",
	}
	bytes, _ := json.Marshal(req)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	httpMock := &http.Request{
		Body: r,
	}
	ginContext.Request = httpMock

	usecases.On("Update", req).Return(nil)

	instance.updateVideo(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestUpdateVideoWithErrorWhenCallingUseCase(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}
	req := domainvideos.Videos{
		ID:   id,
		Name: "TEST ",
		Url:  "www.test.com",
	}
	bytes, _ := json.Marshal(req)
	r := io.NopCloser(strings.NewReader(string(bytes)))
	httpMock := &http.Request{
		Body: r,
	}
	ginContext.Request = httpMock

	usecases.On("Update", req).Return(errors.New("error services"))

	instance.updateVideo(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestUpdateVideoWithErrorWhenShouldBindJSONRequest(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)

	instance.updateVideo(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestDeleteVideoByIDSuccessful(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	usecases.On("DeleteByID", id).Return(nil)

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}

	instance.deleteVideo(ginContext)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestDeleteVideoByIDWithErrorWhenExecutedUseCase(t *testing.T) {
	usecases := new(useCaseMock)
	instance := NewController(usecases)
	responseRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(responseRecorder)
	id := "123456789"

	usecases.On("DeleteByID", id).Return(errors.New("error services"))

	ginContext.Params = []gin.Param{
		{Key: "id", Value: id},
	}

	instance.deleteVideo(ginContext)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}
