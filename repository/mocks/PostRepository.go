// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entity "github.com/vipindasvg/golang-rest-api/entity"
)

// PostRepository is an autogenerated mock type for the PostRepository type
type PostRepository struct {
	mock.Mock
}

// FindAll provides a mock function with given fields:
func (_m *PostRepository) FindAll() ([]entity.Post, error) {
	ret := _m.Called()

	var r0 []entity.Post
	if rf, ok := ret.Get(0).(func() []entity.Post); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: post
func (_m *PostRepository) Save(post *entity.Post) (*entity.Post, error) {
	ret := _m.Called(post)

	var r0 *entity.Post
	if rf, ok := ret.Get(0).(func(*entity.Post) *entity.Post); ok {
		r0 = rf(post)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Post) error); ok {
		r1 = rf(post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
