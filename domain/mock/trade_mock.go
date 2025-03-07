// Code generated by MockGen. DO NOT EDIT.
// Source: ./ (interfaces: TradeUsecase)
//
// Generated by this command:
//
//	mockgen -destination mock/trade_mock.go ./ TradeUsecase
//

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTradeUsecase is a mock of TradeUsecase interface.
type MockTradeUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockTradeUsecaseMockRecorder
	isgomock struct{}
}

// MockTradeUsecaseMockRecorder is the mock recorder for MockTradeUsecase.
type MockTradeUsecaseMockRecorder struct {
	mock *MockTradeUsecase
}

// NewMockTradeUsecase creates a new mock instance.
func NewMockTradeUsecase(ctrl *gomock.Controller) *MockTradeUsecase {
	mock := &MockTradeUsecase{ctrl: ctrl}
	mock.recorder = &MockTradeUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTradeUsecase) EXPECT() *MockTradeUsecaseMockRecorder {
	return m.recorder
}

// GetLastTradePrice mocks base method.
func (m *MockTradeUsecase) GetLastTradePrice() (int, map[string]map[string]float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastTradePrice")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(map[string]map[string]float64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetLastTradePrice indicates an expected call of GetLastTradePrice.
func (mr *MockTradeUsecaseMockRecorder) GetLastTradePrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastTradePrice", reflect.TypeOf((*MockTradeUsecase)(nil).GetLastTradePrice))
}
