// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Xanvial/todo-app-go/backend/datastore (interfaces: DataStore)

// Package mock_datastore is a generated GoMock package.
package mock_datastore

import (
	context "context"
	entity "github.com/Xanvial/todo-app-go/backend/entity"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDataStore is a mock of DataStore interface
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// CreateTodo mocks base method
func (m *MockDataStore) CreateTodo(arg0 context.Context, arg1 string) (*entity.TodoData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo", arg0, arg1)
	ret0, _ := ret[0].(*entity.TodoData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTodo indicates an expected call of CreateTodo
func (mr *MockDataStoreMockRecorder) CreateTodo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockDataStore)(nil).CreateTodo), arg0, arg1)
}

// DeleteTodo mocks base method
func (m *MockDataStore) DeleteTodo(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTodo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTodo indicates an expected call of DeleteTodo
func (mr *MockDataStoreMockRecorder) DeleteTodo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTodo", reflect.TypeOf((*MockDataStore)(nil).DeleteTodo), arg0, arg1)
}

// GetCompleted mocks base method
func (m *MockDataStore) GetCompleted(arg0 context.Context) ([]*entity.TodoData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompleted", arg0)
	ret0, _ := ret[0].([]*entity.TodoData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompleted indicates an expected call of GetCompleted
func (mr *MockDataStoreMockRecorder) GetCompleted(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompleted", reflect.TypeOf((*MockDataStore)(nil).GetCompleted), arg0)
}

// GetIncomplete mocks base method
func (m *MockDataStore) GetIncomplete(arg0 context.Context) ([]*entity.TodoData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncomplete", arg0)
	ret0, _ := ret[0].([]*entity.TodoData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncomplete indicates an expected call of GetIncomplete
func (mr *MockDataStoreMockRecorder) GetIncomplete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncomplete", reflect.TypeOf((*MockDataStore)(nil).GetIncomplete), arg0)
}

// UpdateTodo mocks base method
func (m *MockDataStore) UpdateTodo(arg0 context.Context, arg1 int, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTodo", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTodo indicates an expected call of UpdateTodo
func (mr *MockDataStoreMockRecorder) UpdateTodo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTodo", reflect.TypeOf((*MockDataStore)(nil).UpdateTodo), arg0, arg1, arg2)
}
