// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// PostController is an autogenerated mock type for the PostController type
type PostController struct {
	mock.Mock
}

// AddPost provides a mock function with given fields: resp, req
func (_m *PostController) AddPost(resp http.ResponseWriter, req *http.Request) {
	_m.Called(resp, req)
}

// GetPosts provides a mock function with given fields: resp, req
func (_m *PostController) GetPosts(resp http.ResponseWriter, req *http.Request) {
	_m.Called(resp, req)
}