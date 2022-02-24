// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/init/backend/interface.go

// Package backend is a generated GoMock package.
package backend

import (
	v1alpha2 "github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	parser "github.com/devfile/library/pkg/devfile/parser"
	gomock "github.com/golang/mock/gomock"
	filesystem "github.com/redhat-developer/odo/pkg/testingutil/filesystem"
	reflect "reflect"
)

// MockInitBackend is a mock of InitBackend interface
type MockInitBackend struct {
	ctrl     *gomock.Controller
	recorder *MockInitBackendMockRecorder
}

// MockInitBackendMockRecorder is the mock recorder for MockInitBackend
type MockInitBackendMockRecorder struct {
	mock *MockInitBackend
}

// NewMockInitBackend creates a new mock instance
func NewMockInitBackend(ctrl *gomock.Controller) *MockInitBackend {
	mock := &MockInitBackend{ctrl: ctrl}
	mock.recorder = &MockInitBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInitBackend) EXPECT() *MockInitBackendMockRecorder {
	return m.recorder
}

// Validate mocks base method
func (m *MockInitBackend) Validate(flags map[string]string, fs filesystem.Filesystem, dir string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", flags, fs, dir)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate
func (mr *MockInitBackendMockRecorder) Validate(flags, fs, dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockInitBackend)(nil).Validate), flags, fs, dir)
}

// SelectDevfile mocks base method
func (m *MockInitBackend) SelectDevfile(flags map[string]string, fs filesystem.Filesystem, dir string) (*DevfileLocation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectDevfile", flags, fs, dir)
	ret0, _ := ret[0].(*DevfileLocation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectDevfile indicates an expected call of SelectDevfile
func (mr *MockInitBackendMockRecorder) SelectDevfile(flags, fs, dir interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectDevfile", reflect.TypeOf((*MockInitBackend)(nil).SelectDevfile), flags, fs, dir)
}

// SelectStarterProject mocks base method
func (m *MockInitBackend) SelectStarterProject(devfile parser.DevfileObj, flags map[string]string) (*v1alpha2.StarterProject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectStarterProject", devfile, flags)
	ret0, _ := ret[0].(*v1alpha2.StarterProject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectStarterProject indicates an expected call of SelectStarterProject
func (mr *MockInitBackendMockRecorder) SelectStarterProject(devfile, flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectStarterProject", reflect.TypeOf((*MockInitBackend)(nil).SelectStarterProject), devfile, flags)
}

// PersonalizeName mocks base method
func (m *MockInitBackend) PersonalizeName(devfile parser.DevfileObj, flags map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PersonalizeName", devfile, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

// PersonalizeName indicates an expected call of PersonalizeName
func (mr *MockInitBackendMockRecorder) PersonalizeName(devfile, flags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersonalizeName", reflect.TypeOf((*MockInitBackend)(nil).PersonalizeName), devfile, flags)
}

// PersonalizeDevfileconfig mocks base method
func (m *MockInitBackend) PersonalizeDevfileconfig(devfileobj parser.DevfileObj) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PersonalizeDevfileconfig", devfileobj)
	ret0, _ := ret[0].(error)
	return ret0
}

// PersonalizeDevfileconfig indicates an expected call of PersonalizeDevfileconfig
func (mr *MockInitBackendMockRecorder) PersonalizeDevfileconfig(devfileobj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PersonalizeDevfileconfig", reflect.TypeOf((*MockInitBackend)(nil).PersonalizeDevfileconfig), devfileobj)
}
