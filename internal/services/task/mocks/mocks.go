// Code generated by MockGen. DO NOT EDIT.
// Source: core.go
//
// Generated by this command:
//
//	mockgen -source=core.go -destination=mocks/mocks.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/vas-sh/todo/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// Mockrepoer is a mock of repoer interface.
type Mockrepoer struct {
	ctrl     *gomock.Controller
	recorder *MockrepoerMockRecorder
	isgomock struct{}
}

// MockrepoerMockRecorder is the mock recorder for Mockrepoer.
type MockrepoerMockRecorder struct {
	mock *Mockrepoer
}

// NewMockrepoer creates a new mock instance.
func NewMockrepoer(ctrl *gomock.Controller) *Mockrepoer {
	mock := &Mockrepoer{ctrl: ctrl}
	mock.recorder = &MockrepoerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockrepoer) EXPECT() *MockrepoerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *Mockrepoer) Create(ctx context.Context, title, description string) (models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, title, description)
	ret0, _ := ret[0].(models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockrepoerMockRecorder) Create(ctx, title, description any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*Mockrepoer)(nil).Create), ctx, title, description)
}

// List mocks base method.
func (m *Mockrepoer) List(ctx context.Context) ([]models.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx)
	ret0, _ := ret[0].([]models.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockrepoerMockRecorder) List(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*Mockrepoer)(nil).List), ctx)
}

// Remove mocks base method.
func (m *Mockrepoer) Remove(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockrepoerMockRecorder) Remove(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*Mockrepoer)(nil).Remove), ctx, id)
}
