// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"
	group "docker-go-project/pkg/repository/group"

	mock "github.com/stretchr/testify/mock"
)

// IGroupRepository is an autogenerated mock type for the IGroupRepository type
type IGroupRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, code
func (_m *IGroupRepository) Create(ctx context.Context, code string) (int64, error) {
	ret := _m.Called(ctx, code)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(ctx, code)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, code
func (_m *IGroupRepository) Delete(ctx context.Context, code string) error {
	ret := _m.Called(ctx, code)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *IGroupRepository) GetAll(ctx context.Context) ([]group.Group, error) {
	ret := _m.Called(ctx)

	var r0 []group.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]group.Group, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []group.Group); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]group.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByCode provides a mock function with given fields: ctx, code
func (_m *IGroupRepository) GetByCode(ctx context.Context, code string) (group.Group, error) {
	ret := _m.Called(ctx, code)

	var r0 group.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (group.Group, error)); ok {
		return rf(ctx, code)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) group.Group); ok {
		r0 = rf(ctx, code)
	} else {
		r0 = ret.Get(0).(group.Group)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDebtByCode provides a mock function with given fields: ctx, _a1
func (_m *IGroupRepository) UpdateDebtByCode(ctx context.Context, _a1 group.Group) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, group.Group) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIGroupRepository creates a new instance of IGroupRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIGroupRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IGroupRepository {
	mock := &IGroupRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
