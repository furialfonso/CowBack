package services

import (
	"context"
	"docker-go-project/api/dto/request"
	"docker-go-project/api/dto/response"
	"docker-go-project/mocks"
	"docker-go-project/pkg/repository/model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockGroupService struct {
	groupRepository *mocks.IGroupRepository
}

type groupMocks struct {
	groupService func(f *mockGroupService)
}

func Test_GetGroups(t *testing.T) {
	tests := []struct {
		name   string
		mocks  groupMocks
		outPut []response.GroupResponse
		expErr error
	}{
		{
			name: "error get groups",
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("GetGroups", mock.Anything).Return([]model.Group{}, errors.New("error x"))
				},
			},
			expErr: errors.New("error x"),
		},
		{
			name: "full flow",
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("GetGroups", mock.Anything).Return([]model.Group{
						{
							ID:        1,
							Code:      "YOU&I",
							Debt:      200000,
							CreatedAt: "2023-05-01T08:00:00",
						},
						{
							ID:        2,
							Code:      "test1",
							Debt:      300000,
							CreatedAt: "2023-05-01T08:00:00",
						},
					}, nil)
				},
			},
			outPut: []response.GroupResponse{
				{
					Code: "YOU&I",
				},
				{
					Code: "test1",
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockGroupService{
				groupRepository: &mocks.IGroupRepository{},
			}
			tc.mocks.groupService(m)
			service := NewGroupService(m.groupRepository)
			groups, err := service.GetGroups(context.Background())
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
			assert.Equal(t, tc.outPut, groups)
		})
	}
}

func Test_GetGroupByCode(t *testing.T) {
	tests := []struct {
		name   string
		code   string
		mocks  groupMocks
		outPut response.GroupResponse
		expErr error
	}{
		{
			name: "error get groups",
			code: "YOU&I",
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("GetGroupByCode", mock.Anything, "YOU&I").Return(model.Group{}, errors.New("error x"))
				},
			},
			expErr: errors.New("error x"),
		},
		{
			name: "full flow",
			code: "YOU&I",
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("GetGroupByCode", mock.Anything, "YOU&I").Return(model.Group{
						ID:        1,
						Code:      "YOU&I",
						Debt:      200000,
						CreatedAt: "2023-05-01T08:00:00",
					}, nil)
				},
			},
			outPut: response.GroupResponse{
				Code: "YOU&I",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockGroupService{
				groupRepository: &mocks.IGroupRepository{},
			}
			tc.mocks.groupService(m)
			service := NewGroupService(m.groupRepository)
			groups, err := service.GetGroupByCode(context.Background(), tc.code)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
			assert.Equal(t, tc.outPut, groups)
		})
	}
}

func Test_CreateGroup(t *testing.T) {
	tests := []struct {
		name   string
		input  request.GroupDTO
		mocks  groupMocks
		expErr error
	}{
		{
			name: "error",
			input: request.GroupDTO{
				Code: "Test1",
			},
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("Create", mock.Anything, "Test1").Return(int64(0), errors.New("error x"))
				},
			},
			expErr: errors.New("error x"),
		},
		{
			name: "full flow",
			input: request.GroupDTO{
				Code: "Test1",
			},
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("Create", mock.Anything, "Test1").Return(int64(1), nil)
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockGroupService{
				groupRepository: &mocks.IGroupRepository{},
			}
			tc.mocks.groupService(m)
			service := NewGroupService(m.groupRepository)
			err := service.CreateGroup(context.Background(), tc.input)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
		})
	}
}

func Test_UpdateDebtByCode(t *testing.T) {
	tests := []struct {
		name   string
		input  request.GroupDTO
		mocks  groupMocks
		expErr error
	}{
		{
			name: "error",
			input: request.GroupDTO{
				Code: "Test1",
				Debt: 1000,
			},
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("UpdateGroupDebtByCode", mock.Anything, model.Group{
						Code: "Test1",
						Debt: 1000,
					}).Return(errors.New("error x"))
				},
			},
			expErr: errors.New("error x"),
		},
		{
			name: "full flow",
			input: request.GroupDTO{
				Code: "Test1",
				Debt: 1000,
			},
			mocks: groupMocks{
				groupService: func(f *mockGroupService) {
					f.groupRepository.Mock.On("UpdateGroupDebtByCode", mock.Anything, model.Group{
						Code: "Test1",
						Debt: 1000,
					}).Return(nil)
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &mockGroupService{
				groupRepository: &mocks.IGroupRepository{},
			}
			tc.mocks.groupService(m)
			service := NewGroupService(m.groupRepository)
			err := service.UpdateDebtByCode(context.Background(), tc.input)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
		})
	}
}
