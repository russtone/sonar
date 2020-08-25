// Code generated by mockery v1.0.0. DO NOT EDIT.

package actions_mock

import (
	context "context"

	actions "github.com/bi-zone/sonar/internal/actions"

	mock "github.com/stretchr/testify/mock"
)

// ResultHandler is an autogenerated mock type for the ResultHandler type
type ResultHandler struct {
	mock.Mock
}

// DNSRecordsCreate provides a mock function with given fields: _a0, _a1
func (_m *ResultHandler) DNSRecordsCreate(_a0 context.Context, _a1 *actions.CreateDNSRecordResultData) {
	_m.Called(_a0, _a1)
}

// DNSRecordsDelete provides a mock function with given fields: _a0, _a1
func (_m *ResultHandler) DNSRecordsDelete(_a0 context.Context, _a1 actions.DNSRecordsDeleteResult) {
	_m.Called(_a0, _a1)
}

// DNSRecordsList provides a mock function with given fields: _a0, _a1
func (_m *ResultHandler) DNSRecordsList(_a0 context.Context, _a1 *actions.ListDNSRecordsResultData) {
	_m.Called(_a0, _a1)
}

// PayloadsCreate provides a mock function with given fields: _a0, _a1
func (_m *ResultHandler) PayloadsCreate(_a0 context.Context, _a1 actions.PayloadsCreateResult) {
	_m.Called(_a0, _a1)
}

// PayloadsDelete provides a mock function with given fields: _a0, _a1
func (_m *ResultHandler) PayloadsDelete(_a0 context.Context, _a1 actions.PayloadsDeleteResult) {
	_m.Called(_a0, _a1)
}

// PayloadsList provides a mock function with given fields: _a0, _a1
func (_m *ResultHandler) PayloadsList(_a0 context.Context, _a1 actions.PayloadsListResult) {
	_m.Called(_a0, _a1)
}

// PayloadsUpdate provides a mock function with given fields: _a0, _a1
func (_m *ResultHandler) PayloadsUpdate(_a0 context.Context, _a1 actions.PayloadsUpdateResult) {
	_m.Called(_a0, _a1)
}

// UserCurrent provides a mock function with given fields: _a0, _a1
func (_m *ResultHandler) UserCurrent(_a0 context.Context, _a1 actions.UserCurrentResult) {
	_m.Called(_a0, _a1)
}

// UsersCreate provides a mock function with given fields: _a0, _a1
func (_m *ResultHandler) UsersCreate(_a0 context.Context, _a1 actions.UsersCreateResult) {
	_m.Called(_a0, _a1)
}

// UsersDelete provides a mock function with given fields: _a0, _a1
func (_m *ResultHandler) UsersDelete(_a0 context.Context, _a1 actions.UsersDeleteResult) {
	_m.Called(_a0, _a1)
}
