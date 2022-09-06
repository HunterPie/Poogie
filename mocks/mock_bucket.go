// Code generated by MockGen. DO NOT EDIT.
// Source: ./core/persistence/bucket/bucket.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIBucket is a mock of IBucket interface.
type MockIBucket struct {
	ctrl     *gomock.Controller
	recorder *MockIBucketMockRecorder
}

// MockIBucketMockRecorder is the mock recorder for MockIBucket.
type MockIBucketMockRecorder struct {
	mock *MockIBucket
}

// NewMockIBucket creates a new mock instance.
func NewMockIBucket(ctrl *gomock.Controller) *MockIBucket {
	mock := &MockIBucket{ctrl: ctrl}
	mock.recorder = &MockIBucketMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBucket) EXPECT() *MockIBucketMockRecorder {
	return m.recorder
}

// FindBy mocks base method.
func (m *MockIBucket) FindBy(name string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBy", name)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindBy indicates an expected call of FindBy.
func (mr *MockIBucketMockRecorder) FindBy(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBy", reflect.TypeOf((*MockIBucket)(nil).FindBy), name)
}

// FindMostRecent mocks base method.
func (m *MockIBucket) FindMostRecent() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMostRecent")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMostRecent indicates an expected call of FindMostRecent.
func (mr *MockIBucketMockRecorder) FindMostRecent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMostRecent", reflect.TypeOf((*MockIBucket)(nil).FindMostRecent))
}
