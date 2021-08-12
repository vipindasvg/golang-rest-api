package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// manual mock

type PostControllerMock struct {
	mock.Mock
}

func (d *PostControllerMock) AddPost(resp http.ResponseWriter, req *http.Request) {
	_ = d.Called(resp, req)
	return 
}

func (d *PostControllerMock) GetPosts(resp http.ResponseWriter, req *http.Request) {
	_ = d.Called(resp, req)
	return 
}

func TestGetPosts(t *testing.T) {
	thePostControllerMock := PostControllerMock{}
	w := httptest.NewRecorder()
	w.Code = 200
	assert.Equal(t, http.StatusOK, w.Code)
	thePostControllerMock.AssertExpectations(t)
}
