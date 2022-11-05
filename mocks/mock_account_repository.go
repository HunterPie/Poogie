// Code generated by MockGen. DO NOT EDIT.
// Source: ./core/persistence/account/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	account "github.com/Haato3o/poogie/core/persistence/account"
	gomock "github.com/golang/mock/gomock"
)

// MockIAccountRepository is a mock of IAccountRepository interface.
type MockIAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountRepositoryMockRecorder
}

// MockIAccountRepositoryMockRecorder is the mock recorder for MockIAccountRepository.
type MockIAccountRepositoryMockRecorder struct {
	mock *MockIAccountRepository
}

// NewMockIAccountRepository creates a new mock instance.
func NewMockIAccountRepository(ctrl *gomock.Controller) *MockIAccountRepository {
	mock := &MockIAccountRepository{ctrl: ctrl}
	mock.recorder = &MockIAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAccountRepository) EXPECT() *MockIAccountRepositoryMockRecorder {
	return m.recorder
}

// AreCredentialsValid mocks base method.
func (m *MockIAccountRepository) AreCredentialsValid(ctx context.Context, email, password string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AreCredentialsValid", ctx, email, password)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AreCredentialsValid indicates an expected call of AreCredentialsValid.
func (mr *MockIAccountRepositoryMockRecorder) AreCredentialsValid(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AreCredentialsValid", reflect.TypeOf((*MockIAccountRepository)(nil).AreCredentialsValid), ctx, email, password)
}

// Create mocks base method.
func (m *MockIAccountRepository) Create(ctx context.Context, model account.AccountModel) (account.AccountModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, model)
	ret0, _ := ret[0].(account.AccountModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIAccountRepositoryMockRecorder) Create(ctx, model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIAccountRepository)(nil).Create), ctx, model)
}

// DeleteBy mocks base method.
func (m *MockIAccountRepository) DeleteBy(ctx context.Context, userId string) account.AccountModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBy", ctx, userId)
	ret0, _ := ret[0].(account.AccountModel)
	return ret0
}

// DeleteBy indicates an expected call of DeleteBy.
func (mr *MockIAccountRepositoryMockRecorder) DeleteBy(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBy", reflect.TypeOf((*MockIAccountRepository)(nil).DeleteBy), ctx, userId)
}

// GetByEmail mocks base method.
func (m *MockIAccountRepository) GetByEmail(ctx context.Context, email string) (account.AccountModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(account.AccountModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockIAccountRepositoryMockRecorder) GetByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockIAccountRepository)(nil).GetByEmail), ctx, email)
}

// GetById mocks base method.
func (m *MockIAccountRepository) GetById(ctx context.Context, userId string) (account.AccountModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, userId)
	ret0, _ := ret[0].(account.AccountModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIAccountRepositoryMockRecorder) GetById(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIAccountRepository)(nil).GetById), ctx, userId)
}

// IsEmailTaken mocks base method.
func (m *MockIAccountRepository) IsEmailTaken(ctx context.Context, email string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEmailTaken", ctx, email)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEmailTaken indicates an expected call of IsEmailTaken.
func (mr *MockIAccountRepositoryMockRecorder) IsEmailTaken(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEmailTaken", reflect.TypeOf((*MockIAccountRepository)(nil).IsEmailTaken), ctx, email)
}

// IsUsernameTaken mocks base method.
func (m *MockIAccountRepository) IsUsernameTaken(ctx context.Context, username string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUsernameTaken", ctx, username)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUsernameTaken indicates an expected call of IsUsernameTaken.
func (mr *MockIAccountRepositoryMockRecorder) IsUsernameTaken(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUsernameTaken", reflect.TypeOf((*MockIAccountRepository)(nil).IsUsernameTaken), ctx, username)
}

// UpdateAvatar mocks base method.
func (m *MockIAccountRepository) UpdateAvatar(ctx context.Context, userId, avatar string) account.AccountModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", ctx, userId, avatar)
	ret0, _ := ret[0].(account.AccountModel)
	return ret0
}

// UpdateAvatar indicates an expected call of UpdateAvatar.
func (mr *MockIAccountRepositoryMockRecorder) UpdateAvatar(ctx, userId, avatar interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*MockIAccountRepository)(nil).UpdateAvatar), ctx, userId, avatar)
}

// UpdatePassword mocks base method.
func (m *MockIAccountRepository) UpdatePassword(ctx context.Context, userId, password string) account.AccountModel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", ctx, userId, password)
	ret0, _ := ret[0].(account.AccountModel)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockIAccountRepositoryMockRecorder) UpdatePassword(ctx, userId, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockIAccountRepository)(nil).UpdatePassword), ctx, userId, password)
}

// UpdateSupporterStatus mocks base method.
func (m *MockIAccountRepository) UpdateSupporterStatus(ctx context.Context, userId string, isSupporter bool) (account.AccountModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSupporterStatus", ctx, userId, isSupporter)
	ret0, _ := ret[0].(account.AccountModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSupporterStatus indicates an expected call of UpdateSupporterStatus.
func (mr *MockIAccountRepositoryMockRecorder) UpdateSupporterStatus(ctx, userId, isSupporter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSupporterStatus", reflect.TypeOf((*MockIAccountRepository)(nil).UpdateSupporterStatus), ctx, userId, isSupporter)
}

// VerifyAccount mocks base method.
func (m *MockIAccountRepository) VerifyAccount(ctx context.Context, userId string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "VerifyAccount", ctx, userId)
}

// VerifyAccount indicates an expected call of VerifyAccount.
func (mr *MockIAccountRepositoryMockRecorder) VerifyAccount(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyAccount", reflect.TypeOf((*MockIAccountRepository)(nil).VerifyAccount), ctx, userId)
}

// MockIAccountSessionRepository is a mock of IAccountSessionRepository interface.
type MockIAccountSessionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountSessionRepositoryMockRecorder
}

// MockIAccountSessionRepositoryMockRecorder is the mock recorder for MockIAccountSessionRepository.
type MockIAccountSessionRepositoryMockRecorder struct {
	mock *MockIAccountSessionRepository
}

// NewMockIAccountSessionRepository creates a new mock instance.
func NewMockIAccountSessionRepository(ctrl *gomock.Controller) *MockIAccountSessionRepository {
	mock := &MockIAccountSessionRepository{ctrl: ctrl}
	mock.recorder = &MockIAccountSessionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAccountSessionRepository) EXPECT() *MockIAccountSessionRepositoryMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockIAccountSessionRepository) CreateSession(ctx context.Context, token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockIAccountSessionRepositoryMockRecorder) CreateSession(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockIAccountSessionRepository)(nil).CreateSession), ctx, token)
}

// IsSessionValid mocks base method.
func (m *MockIAccountSessionRepository) IsSessionValid(ctx context.Context, token string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSessionValid", ctx, token)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSessionValid indicates an expected call of IsSessionValid.
func (mr *MockIAccountSessionRepositoryMockRecorder) IsSessionValid(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSessionValid", reflect.TypeOf((*MockIAccountSessionRepository)(nil).IsSessionValid), ctx, token)
}

// RevokeSession mocks base method.
func (m *MockIAccountSessionRepository) RevokeSession(ctx context.Context, token string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeSession", ctx, token)
	ret0, _ := ret[0].(string)
	return ret0
}

// RevokeSession indicates an expected call of RevokeSession.
func (mr *MockIAccountSessionRepositoryMockRecorder) RevokeSession(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeSession", reflect.TypeOf((*MockIAccountSessionRepository)(nil).RevokeSession), ctx, token)
}

// MockIAccountBadgesRepository is a mock of IAccountBadgesRepository interface.
type MockIAccountBadgesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountBadgesRepositoryMockRecorder
}

// MockIAccountBadgesRepositoryMockRecorder is the mock recorder for MockIAccountBadgesRepository.
type MockIAccountBadgesRepositoryMockRecorder struct {
	mock *MockIAccountBadgesRepository
}

// NewMockIAccountBadgesRepository creates a new mock instance.
func NewMockIAccountBadgesRepository(ctrl *gomock.Controller) *MockIAccountBadgesRepository {
	mock := &MockIAccountBadgesRepository{ctrl: ctrl}
	mock.recorder = &MockIAccountBadgesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAccountBadgesRepository) EXPECT() *MockIAccountBadgesRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIAccountBadgesRepository) Create(ctx context.Context, userId, badgeId string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", ctx, userId, badgeId)
}

// Create indicates an expected call of Create.
func (mr *MockIAccountBadgesRepositoryMockRecorder) Create(ctx, userId, badgeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIAccountBadgesRepository)(nil).Create), ctx, userId, badgeId)
}

// Delete mocks base method.
func (m *MockIAccountBadgesRepository) Delete(ctx context.Context, userId, badgeId string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", ctx, userId, badgeId)
}

// Delete indicates an expected call of Delete.
func (mr *MockIAccountBadgesRepositoryMockRecorder) Delete(ctx, userId, badgeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIAccountBadgesRepository)(nil).Delete), ctx, userId, badgeId)
}

// MockIAccountHuntStatisticSummaryRepository is a mock of IAccountHuntStatisticSummaryRepository interface.
type MockIAccountHuntStatisticSummaryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountHuntStatisticSummaryRepositoryMockRecorder
}

// MockIAccountHuntStatisticSummaryRepositoryMockRecorder is the mock recorder for MockIAccountHuntStatisticSummaryRepository.
type MockIAccountHuntStatisticSummaryRepositoryMockRecorder struct {
	mock *MockIAccountHuntStatisticSummaryRepository
}

// NewMockIAccountHuntStatisticSummaryRepository creates a new mock instance.
func NewMockIAccountHuntStatisticSummaryRepository(ctrl *gomock.Controller) *MockIAccountHuntStatisticSummaryRepository {
	mock := &MockIAccountHuntStatisticSummaryRepository{ctrl: ctrl}
	mock.recorder = &MockIAccountHuntStatisticSummaryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAccountHuntStatisticSummaryRepository) EXPECT() *MockIAccountHuntStatisticSummaryRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIAccountHuntStatisticSummaryRepository) Create(ctx context.Context, userId, badgeId string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", ctx, userId, badgeId)
}

// Create indicates an expected call of Create.
func (mr *MockIAccountHuntStatisticSummaryRepositoryMockRecorder) Create(ctx, userId, badgeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIAccountHuntStatisticSummaryRepository)(nil).Create), ctx, userId, badgeId)
}

// Delete mocks base method.
func (m *MockIAccountHuntStatisticSummaryRepository) Delete(ctx context.Context, userId, badgeId string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", ctx, userId, badgeId)
}

// Delete indicates an expected call of Delete.
func (mr *MockIAccountHuntStatisticSummaryRepositoryMockRecorder) Delete(ctx, userId, badgeId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIAccountHuntStatisticSummaryRepository)(nil).Delete), ctx, userId, badgeId)
}

// MockIAccountVerificationRepository is a mock of IAccountVerificationRepository interface.
type MockIAccountVerificationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountVerificationRepositoryMockRecorder
}

// MockIAccountVerificationRepositoryMockRecorder is the mock recorder for MockIAccountVerificationRepository.
type MockIAccountVerificationRepositoryMockRecorder struct {
	mock *MockIAccountVerificationRepository
}

// NewMockIAccountVerificationRepository creates a new mock instance.
func NewMockIAccountVerificationRepository(ctrl *gomock.Controller) *MockIAccountVerificationRepository {
	mock := &MockIAccountVerificationRepository{ctrl: ctrl}
	mock.recorder = &MockIAccountVerificationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAccountVerificationRepository) EXPECT() *MockIAccountVerificationRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIAccountVerificationRepository) Create(ctx context.Context, token, account string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", ctx, token, account)
}

// Create indicates an expected call of Create.
func (mr *MockIAccountVerificationRepositoryMockRecorder) Create(ctx, token, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIAccountVerificationRepository)(nil).Create), ctx, token, account)
}

// Find mocks base method.
func (m *MockIAccountVerificationRepository) Find(ctx context.Context, token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockIAccountVerificationRepositoryMockRecorder) Find(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockIAccountVerificationRepository)(nil).Find), ctx, token)
}

// MockIAccountResetRepository is a mock of IAccountResetRepository interface.
type MockIAccountResetRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAccountResetRepositoryMockRecorder
}

// MockIAccountResetRepositoryMockRecorder is the mock recorder for MockIAccountResetRepository.
type MockIAccountResetRepositoryMockRecorder struct {
	mock *MockIAccountResetRepository
}

// NewMockIAccountResetRepository creates a new mock instance.
func NewMockIAccountResetRepository(ctrl *gomock.Controller) *MockIAccountResetRepository {
	mock := &MockIAccountResetRepository{ctrl: ctrl}
	mock.recorder = &MockIAccountResetRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAccountResetRepository) EXPECT() *MockIAccountResetRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIAccountResetRepository) Create(ctx context.Context, code, email string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Create", ctx, code, email)
}

// Create indicates an expected call of Create.
func (mr *MockIAccountResetRepositoryMockRecorder) Create(ctx, code, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIAccountResetRepository)(nil).Create), ctx, code, email)
}

// IsTokenValid mocks base method.
func (m *MockIAccountResetRepository) IsTokenValid(ctx context.Context, code, email string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTokenValid", ctx, code, email)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTokenValid indicates an expected call of IsTokenValid.
func (mr *MockIAccountResetRepositoryMockRecorder) IsTokenValid(ctx, code, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTokenValid", reflect.TypeOf((*MockIAccountResetRepository)(nil).IsTokenValid), ctx, code, email)
}

// RevokeBy mocks base method.
func (m *MockIAccountResetRepository) RevokeBy(ctx context.Context, code, email string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RevokeBy", ctx, code, email)
}

// RevokeBy indicates an expected call of RevokeBy.
func (mr *MockIAccountResetRepositoryMockRecorder) RevokeBy(ctx, code, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeBy", reflect.TypeOf((*MockIAccountResetRepository)(nil).RevokeBy), ctx, code, email)
}
