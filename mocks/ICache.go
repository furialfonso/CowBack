// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	cache "cow_back/pkg/platform/cache"

	mock "github.com/stretchr/testify/mock"
)

// ICache is an autogenerated mock type for the ICache type
type ICache struct {
	mock.Mock
}

// Get provides a mock function with given fields: key
func (_m *ICache) Get(key string) (cache.User, bool) {
	ret := _m.Called(key)

	var r0 cache.User
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (cache.User, bool)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) cache.User); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(cache.User)
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetByNickName provides a mock function with given fields: nickName
func (_m *ICache) GetByNickName(nickName string) (cache.User, bool) {
	ret := _m.Called(nickName)

	var r0 cache.User
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (cache.User, bool)); ok {
		return rf(nickName)
	}
	if rf, ok := ret.Get(0).(func(string) cache.User); ok {
		r0 = rf(nickName)
	} else {
		r0 = ret.Get(0).(cache.User)
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(nickName)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Set provides a mock function with given fields: key, user
func (_m *ICache) Set(key string, user cache.User) {
	_m.Called(key, user)
}

// NewICache creates a new instance of ICache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewICache(t interface {
	mock.TestingT
	Cleanup(func())
}) *ICache {
	mock := &ICache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}