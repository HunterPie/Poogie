// Code generated by MockGen. DO NOT EDIT.
// Source: ./core/persistence/supporter/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	supporter "github.com/Haato3o/poogie/core/persistence/supporter"
	gomock "github.com/golang/mock/gomock"
)

// MockISupporterRepository is a mock of ISupporterRepository interface.
type MockISupporterRepository struct {
	ctrl     *gomock.Controller
	recorder *MockISupporterRepositoryMockRecorder
}

// MockISupporterRepositoryMockRecorder is the mock recorder for MockISupporterRepository.
type MockISupporterRepositoryMockRecorder struct {
	mock *MockISupporterRepository
}

// NewMockISupporterRepository creates a new mock instance.
func NewMockISupporterRepository(ctrl *gomock.Controller) *MockISupporterRepository {
	mock := &MockISupporterRepository{ctrl: ctrl}
	mock.recorder = &MockISupporterRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISupporterRepository) EXPECT() *MockISupporterRepositoryMockRecorder {
	return m.recorder
}

// AssociateToUser mocks base method.
func (m *MockISupporterRepository) AssociateToUser(ctx context.Context, email, userId string) supporter.SupporterModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssociateToUser", ctx, email, userId)
	ret0, _ := ret[0].(supporter.SupporterModel)
	return ret0
}

// AssociateToUser indicates an expected call of AssociateToUser.
func (mr *MockISupporterRepositoryMockRecorder) AssociateToUser(ctx, email, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssociateToUser", reflect.TypeOf((*MockISupporterRepository)(nil).AssociateToUser), ctx, email, userId)
}

// ExistsSupporter mocks base method.
func (m *MockISupporterRepository) ExistsSupporter(ctx context.Context, email string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsSupporter", ctx, email)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExistsSupporter indicates an expected call of ExistsSupporter.
func (mr *MockISupporterRepositoryMockRecorder) ExistsSupporter(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsSupporter", reflect.TypeOf((*MockISupporterRepository)(nil).ExistsSupporter), ctx, email)
}

// ExistsToken mocks base method.
func (m *MockISupporterRepository) ExistsToken(ctx context.Context, token string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsToken", ctx, token)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExistsToken indicates an expected call of ExistsToken.
func (mr *MockISupporterRepositoryMockRecorder) ExistsToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsToken", reflect.TypeOf((*MockISupporterRepository)(nil).ExistsToken), ctx, token)
}

// FindBy mocks base method.
func (m *MockISupporterRepository) FindBy(ctx context.Context, email string) supporter.SupporterModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindBy", ctx, email)
	ret0, _ := ret[0].(supporter.SupporterModel)
	return ret0
}

// FindBy indicates an expected call of FindBy.
func (mr *MockISupporterRepositoryMockRecorder) FindBy(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindBy", reflect.TypeOf((*MockISupporterRepository)(nil).FindBy), ctx, email)
}

// FindByAssociation mocks base method.
func (m *MockISupporterRepository) FindByAssociation(ctx context.Context, userId string) supporter.SupporterModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByAssociation", ctx, userId)
	ret0, _ := ret[0].(supporter.SupporterModel)
	return ret0
}

// FindByAssociation indicates an expected call of FindByAssociation.
func (mr *MockISupporterRepositoryMockRecorder) FindByAssociation(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByAssociation", reflect.TypeOf((*MockISupporterRepository)(nil).FindByAssociation), ctx, userId)
}

// Insert mocks base method.
func (m *MockISupporterRepository) Insert(ctx context.Context, model supporter.SupporterModel) supporter.SupporterModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, model)
	ret0, _ := ret[0].(supporter.SupporterModel)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockISupporterRepositoryMockRecorder) Insert(ctx, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockISupporterRepository)(nil).Insert), ctx, model)
}

// RenewBy mocks base method.
func (m *MockISupporterRepository) RenewBy(ctx context.Context, email string) supporter.SupporterModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenewBy", ctx, email)
	ret0, _ := ret[0].(supporter.SupporterModel)
	return ret0
}

// RenewBy indicates an expected call of RenewBy.
func (mr *MockISupporterRepositoryMockRecorder) RenewBy(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewBy", reflect.TypeOf((*MockISupporterRepository)(nil).RenewBy), ctx, email)
}

// RevokeBy mocks base method.
func (m *MockISupporterRepository) RevokeBy(ctx context.Context, email string) supporter.SupporterModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeBy", ctx, email)
	ret0, _ := ret[0].(supporter.SupporterModel)
	return ret0
}

// RevokeBy indicates an expected call of RevokeBy.
func (mr *MockISupporterRepositoryMockRecorder) RevokeBy(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeBy", reflect.TypeOf((*MockISupporterRepository)(nil).RevokeBy), ctx, email)
}
