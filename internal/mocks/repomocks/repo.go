// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repo/repo.go

// Package repomocks is a generated GoMock package.
package repomocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/amiosamu/todo-list/internal/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockTask is a mock of Task interface.
type MockTask struct {
	ctrl     *gomock.Controller
	recorder *MockTaskMockRecorder
}

// MockTaskMockRecorder is the mock recorder for MockTask.
type MockTaskMockRecorder struct {
	mock *MockTask
}

// NewMockTask creates a new mock instance.
func NewMockTask(ctrl *gomock.Controller) *MockTask {
	mock := &MockTask{ctrl: ctrl}
	mock.recorder = &MockTaskMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTask) EXPECT() *MockTaskMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockTask) CreateTask(ctx context.Context, t entity.Task) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", ctx, t)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockTaskMockRecorder) CreateTask(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockTask)(nil).CreateTask), ctx, t)
}

// DeleteTask mocks base method.
func (m *MockTask) DeleteTask(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockTaskMockRecorder) DeleteTask(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockTask)(nil).DeleteTask), ctx, id)
}

// GetTaskByID mocks base method.
func (m *MockTask) GetTaskByID(ctx context.Context, id string) (entity.UpdateTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskByID", ctx, id)
	ret0, _ := ret[0].(entity.UpdateTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskByID indicates an expected call of GetTaskByID.
func (mr *MockTaskMockRecorder) GetTaskByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskByID", reflect.TypeOf((*MockTask)(nil).GetTaskByID), ctx, id)
}

// GetTasksByStatus mocks base method.
func (m *MockTask) GetTasksByStatus(ctx context.Context, status string) ([]entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasksByStatus", ctx, status)
	ret0, _ := ret[0].([]entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasksByStatus indicates an expected call of GetTasksByStatus.
func (mr *MockTaskMockRecorder) GetTasksByStatus(ctx, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksByStatus", reflect.TypeOf((*MockTask)(nil).GetTasksByStatus), ctx, status)
}

// TaskDone mocks base method.
func (m *MockTask) TaskDone(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TaskDone", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// TaskDone indicates an expected call of TaskDone.
func (mr *MockTaskMockRecorder) TaskDone(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TaskDone", reflect.TypeOf((*MockTask)(nil).TaskDone), ctx, id)
}

// UpdateTask mocks base method.
func (m *MockTask) UpdateTask(ctx context.Context, t entity.UpdateTask, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", ctx, t, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockTaskMockRecorder) UpdateTask(ctx, t, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockTask)(nil).UpdateTask), ctx, t, id)
}
