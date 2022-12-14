// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/husamettinarabaci/tinyurl/db/sqlc (interfaces: Store)

// Package mock_sqlc is a generated GoMock package.
package mock_sqlc

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/husamettinarabaci/tinyurl/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateTransform mocks base method.
func (m *MockStore) CreateTransform(arg0 context.Context, arg1 db.CreateTransformParams) (db.Transform, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransform", arg0, arg1)
	ret0, _ := ret[0].(db.Transform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransform indicates an expected call of CreateTransform.
func (mr *MockStoreMockRecorder) CreateTransform(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransform", reflect.TypeOf((*MockStore)(nil).CreateTransform), arg0, arg1)
}

// CreateTransformTx mocks base method.
func (m *MockStore) CreateTransformTx(arg0 context.Context, arg1 db.CreateTransformParams) (db.Transform, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransformTx", arg0, arg1)
	ret0, _ := ret[0].(db.Transform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransformTx indicates an expected call of CreateTransformTx.
func (mr *MockStoreMockRecorder) CreateTransformTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransformTx", reflect.TypeOf((*MockStore)(nil).CreateTransformTx), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DecUrlCount mocks base method.
func (m *MockStore) DecUrlCount(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecUrlCount", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecUrlCount indicates an expected call of DecUrlCount.
func (mr *MockStoreMockRecorder) DecUrlCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecUrlCount", reflect.TypeOf((*MockStore)(nil).DecUrlCount), arg0, arg1)
}

// DeleteTransform mocks base method.
func (m *MockStore) DeleteTransform(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransform", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTransform indicates an expected call of DeleteTransform.
func (mr *MockStoreMockRecorder) DeleteTransform(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransform", reflect.TypeOf((*MockStore)(nil).DeleteTransform), arg0, arg1)
}

// DeleteTransformTx mocks base method.
func (m *MockStore) DeleteTransformTx(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransformTx", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTransformTx indicates an expected call of DeleteTransformTx.
func (mr *MockStoreMockRecorder) DeleteTransformTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransformTx", reflect.TypeOf((*MockStore)(nil).DeleteTransformTx), arg0, arg1)
}

// GetTransform mocks base method.
func (m *MockStore) GetTransform(arg0 context.Context, arg1 int64) (db.Transform, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransform", arg0, arg1)
	ret0, _ := ret[0].(db.Transform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransform indicates an expected call of GetTransform.
func (mr *MockStoreMockRecorder) GetTransform(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransform", reflect.TypeOf((*MockStore)(nil).GetTransform), arg0, arg1)
}

// GetTransformForUpdate mocks base method.
func (m *MockStore) GetTransformForUpdate(arg0 context.Context, arg1 int64) (db.Transform, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransformForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Transform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransformForUpdate indicates an expected call of GetTransformForUpdate.
func (mr *MockStoreMockRecorder) GetTransformForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransformForUpdate", reflect.TypeOf((*MockStore)(nil).GetTransformForUpdate), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// IncUrlCount mocks base method.
func (m *MockStore) IncUrlCount(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncUrlCount", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IncUrlCount indicates an expected call of IncUrlCount.
func (mr *MockStoreMockRecorder) IncUrlCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncUrlCount", reflect.TypeOf((*MockStore)(nil).IncUrlCount), arg0, arg1)
}

// ListTransforms mocks base method.
func (m *MockStore) ListTransforms(arg0 context.Context, arg1 db.ListTransformsParams) ([]db.Transform, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTransforms", arg0, arg1)
	ret0, _ := ret[0].([]db.Transform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTransforms indicates an expected call of ListTransforms.
func (mr *MockStoreMockRecorder) ListTransforms(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTransforms", reflect.TypeOf((*MockStore)(nil).ListTransforms), arg0, arg1)
}

// UpdateTransform mocks base method.
func (m *MockStore) UpdateTransform(arg0 context.Context, arg1 db.UpdateTransformParams) (db.Transform, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransform", arg0, arg1)
	ret0, _ := ret[0].(db.Transform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTransform indicates an expected call of UpdateTransform.
func (mr *MockStoreMockRecorder) UpdateTransform(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransform", reflect.TypeOf((*MockStore)(nil).UpdateTransform), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}
