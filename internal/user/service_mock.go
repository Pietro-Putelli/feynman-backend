// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=./service_mock.go -package=user
//

// Package user is a generated GoMock package.
package user

import (
	reflect "reflect"

	domain "github.com/pietro-putelli/feynman-backend/internal/domain"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateUserIfNotExists mocks base method.
func (m *MockService) CreateUserIfNotExists(user *domain.ThirdPartyUser) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserIfNotExists", user)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserIfNotExists indicates an expected call of CreateUserIfNotExists.
func (mr *MockServiceMockRecorder) CreateUserIfNotExists(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserIfNotExists", reflect.TypeOf((*MockService)(nil).CreateUserIfNotExists), user)
}

// GetUserByGuid mocks base method.
func (m *MockService) GetUserByGuid(guid uuid.UUID) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByGuid", guid)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByGuid indicates an expected call of GetUserByGuid.
func (mr *MockServiceMockRecorder) GetUserByGuid(guid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByGuid", reflect.TypeOf((*MockService)(nil).GetUserByGuid), guid)
}
