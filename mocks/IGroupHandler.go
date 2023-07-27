// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// IGroupHandler is an autogenerated mock type for the IGroupHandler type
type IGroupHandler struct {
	mock.Mock
}

// Create provides a mock function with given fields: c
func (_m *IGroupHandler) Create(c *gin.Context) {
	_m.Called(c)
}

// GetGroupByCode provides a mock function with given fields: c
func (_m *IGroupHandler) GetGroupByCode(c *gin.Context) {
	_m.Called(c)
}

// GetGroups provides a mock function with given fields: c
func (_m *IGroupHandler) GetGroups(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewIGroupHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewIGroupHandler creates a new instance of IGroupHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIGroupHandler(t mockConstructorTestingTNewIGroupHandler) *IGroupHandler {
	mock := &IGroupHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}