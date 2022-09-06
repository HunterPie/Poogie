// Code generated by MockGen. DO NOT EDIT.
// Source: ./core/persistence/database/database.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	account "github.com/Haato3o/poogie/core/persistence/account"
	supporter "github.com/Haato3o/poogie/core/persistence/supporter"
	gomock "github.com/golang/mock/gomock"
)

// MockIDatabase is a mock of IDatabase interface.
type MockIDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockIDatabaseMockRecorder
}

// MockIDatabaseMockRecorder is the mock recorder for MockIDatabase.
type MockIDatabaseMockRecorder struct {
	mock *MockIDatabase
}

// NewMockIDatabase creates a new mock instance.
func NewMockIDatabase(ctrl *gomock.Controller) *MockIDatabase {
	mock := &MockIDatabase{ctrl: ctrl}
	mock.recorder = &MockIDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDatabase) EXPECT() *MockIDatabaseMockRecorder {
	return m.recorder
}

// GetAccountRepository mocks base method.
func (m *MockIDatabase) GetAccountRepository() account.IAccountRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountRepository")
	ret0, _ := ret[0].(account.IAccountRepository)
	return ret0
}

// GetAccountRepository indicates an expected call of GetAccountRepository.
func (mr *MockIDatabaseMockRecorder) GetAccountRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountRepository", reflect.TypeOf((*MockIDatabase)(nil).GetAccountRepository))
}

// GetSessionRepository mocks base method.
func (m *MockIDatabase) GetSessionRepository() account.IAccountSessionRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionRepository")
	ret0, _ := ret[0].(account.IAccountSessionRepository)
	return ret0
}

// GetSessionRepository indicates an expected call of GetSessionRepository.
func (mr *MockIDatabaseMockRecorder) GetSessionRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionRepository", reflect.TypeOf((*MockIDatabase)(nil).GetSessionRepository))
}

// GetSupporterRepository mocks base method.
func (m *MockIDatabase) GetSupporterRepository() supporter.ISupporterRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupporterRepository")
	ret0, _ := ret[0].(supporter.ISupporterRepository)
	return ret0
}

// GetSupporterRepository indicates an expected call of GetSupporterRepository.
func (mr *MockIDatabaseMockRecorder) GetSupporterRepository() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupporterRepository", reflect.TypeOf((*MockIDatabase)(nil).GetSupporterRepository))
}

// IsHealthy mocks base method.
func (m *MockIDatabase) IsHealthy(ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsHealthy", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsHealthy indicates an expected call of IsHealthy.
func (mr *MockIDatabaseMockRecorder) IsHealthy(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsHealthy", reflect.TypeOf((*MockIDatabase)(nil).IsHealthy), ctx)
}
