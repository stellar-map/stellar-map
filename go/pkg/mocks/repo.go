// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stellar-map/stellar-map/go/pkg/entities (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	entities "github.com/stellar-map/stellar-map/go/pkg/entities"
	reflect "reflect"
)

// MockRepo is a mock of Repo interface
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// BatchCreateLedgers mocks base method
func (m *MockRepo) BatchCreateLedgers(arg0 []*entities.Ledger) error {
	ret := m.ctrl.Call(m, "BatchCreateLedgers", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchCreateLedgers indicates an expected call of BatchCreateLedgers
func (mr *MockRepoMockRecorder) BatchCreateLedgers(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateLedgers", reflect.TypeOf((*MockRepo)(nil).BatchCreateLedgers), arg0)
}

// BatchCreateTransactions mocks base method
func (m *MockRepo) BatchCreateTransactions(arg0 []*entities.Transaction) error {
	ret := m.ctrl.Call(m, "BatchCreateTransactions", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchCreateTransactions indicates an expected call of BatchCreateTransactions
func (mr *MockRepoMockRecorder) BatchCreateTransactions(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchCreateTransactions", reflect.TypeOf((*MockRepo)(nil).BatchCreateTransactions), arg0)
}

// PagingToken mocks base method
func (m *MockRepo) PagingToken(arg0 string) (string, error) {
	ret := m.ctrl.Call(m, "PagingToken", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PagingToken indicates an expected call of PagingToken
func (mr *MockRepoMockRecorder) PagingToken(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PagingToken", reflect.TypeOf((*MockRepo)(nil).PagingToken), arg0)
}