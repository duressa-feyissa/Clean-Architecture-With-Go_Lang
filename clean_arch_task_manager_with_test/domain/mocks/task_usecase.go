// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "cleantaskmanager/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskUsecase is an autogenerated mock type for the TaskUsecase type
type TaskUsecase struct {
	mock.Mock
}

// AddTask provides a mock function with given fields: c, claims, task
func (_m *TaskUsecase) AddTask(c context.Context, claims *domain.Claims, task *domain.Task) (primitive.ObjectID, error) {
	ret := _m.Called(c, claims, task)

	var r0 primitive.ObjectID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Claims, *domain.Task) (primitive.ObjectID, error)); ok {
		r0, r1 = rf(c, claims, task)
	} else {
		r1 = ret.Error(0)
	}

	return r0, r1
}

// DeleteTask provides a mock function with given fields: c, claims, id
func (_m *TaskUsecase) DeleteTask(c context.Context, claims *domain.Claims, id primitive.ObjectID) error {
	ret := _m.Called(c, claims, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Claims, primitive.ObjectID) error); ok {
		r0 = rf(c, claims, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTask provides a mock function with given fields: c, claims, id
func (_m *TaskUsecase) GetTask(c context.Context, claims *domain.Claims, id primitive.ObjectID) (*domain.Task, error) {
	ret := _m.Called(c, claims, id)

	var r0 *domain.Task
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Claims, primitive.ObjectID) *domain.Task); ok {
		r0 = rf(c, claims, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Claims, primitive.ObjectID) error); ok {
		r1 = rf(c, claims, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTasks provides a mock function with given fields: c, claims
func (_m *TaskUsecase) GetTasks(c context.Context, claims *domain.Claims) ([]domain.Task, error) {
	ret := _m.Called(c, claims)

	var r0 []domain.Task
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Claims) []domain.Task); ok {
		r0 = rf(c, claims)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Claims) error); ok {
		r1 = rf(c, claims)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTask provides a mock function with given fields: c, claims, id, task
func (_m *TaskUsecase) UpdateTask(c context.Context, claims *domain.Claims, id primitive.ObjectID, task *domain.UpdateTask) error {
	ret := _m.Called(c, claims, id, task)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Claims, primitive.ObjectID, *domain.UpdateTask) error); ok {
		r0 = rf(c, claims, id, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTaskUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewTaskUsecase creates a new instance of TaskUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTaskUsecase(t mockConstructorTestingTNewTaskUsecase) *TaskUsecase {
	mock := &TaskUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}