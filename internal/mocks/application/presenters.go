// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/application/presenters.go

// Package application_mock is a generated GoMock package.
package application_mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/user-ranking/internal/domain"
)

// MockRankingPresenter is a mock of RankingPresenter interface.
type MockRankingPresenter struct {
	ctrl     *gomock.Controller
	recorder *MockRankingPresenterMockRecorder
}

// MockRankingPresenterMockRecorder is the mock recorder for MockRankingPresenter.
type MockRankingPresenterMockRecorder struct {
	mock *MockRankingPresenter
}

// NewMockRankingPresenter creates a new mock instance.
func NewMockRankingPresenter(ctrl *gomock.Controller) *MockRankingPresenter {
	mock := &MockRankingPresenter{ctrl: ctrl}
	mock.recorder = &MockRankingPresenterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRankingPresenter) EXPECT() *MockRankingPresenterMockRecorder {
	return m.recorder
}

// PresentMany mocks base method.
func (m *MockRankingPresenter) PresentMany(ranking []domain.UserRank) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PresentMany", ranking)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// PresentMany indicates an expected call of PresentMany.
func (mr *MockRankingPresenterMockRecorder) PresentMany(ranking interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresentMany", reflect.TypeOf((*MockRankingPresenter)(nil).PresentMany), ranking)
}
