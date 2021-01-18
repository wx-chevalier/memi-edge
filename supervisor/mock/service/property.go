// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/baetyl/baetyl-cloud/v2/service (interfaces: PropertyService)

// Package service is a generated GoMock package.
package service

import (
	models "github.com/baetyl/baetyl-cloud/v2/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPropertyService is a mock of PropertyService interface
type MockPropertyService struct {
	ctrl     *gomock.Controller
	recorder *MockPropertyServiceMockRecorder
}

// MockPropertyServiceMockRecorder is the mock recorder for MockPropertyService
type MockPropertyServiceMockRecorder struct {
	mock *MockPropertyService
}

// NewMockPropertyService creates a new mock instance
func NewMockPropertyService(ctrl *gomock.Controller) *MockPropertyService {
	mock := &MockPropertyService{ctrl: ctrl}
	mock.recorder = &MockPropertyServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPropertyService) EXPECT() *MockPropertyServiceMockRecorder {
	return m.recorder
}

// CountProperty mocks base method
func (m *MockPropertyService) CountProperty(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountProperty", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountProperty indicates an expected call of CountProperty
func (mr *MockPropertyServiceMockRecorder) CountProperty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountProperty", reflect.TypeOf((*MockPropertyService)(nil).CountProperty), arg0)
}

// CreateProperty mocks base method
func (m *MockPropertyService) CreateProperty(arg0 *models.Property) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProperty", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProperty indicates an expected call of CreateProperty
func (mr *MockPropertyServiceMockRecorder) CreateProperty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProperty", reflect.TypeOf((*MockPropertyService)(nil).CreateProperty), arg0)
}

// CreateSysApp mocks base method
func (m *MockPropertyService) CreateSysApp(arg0 *models.NodeSysAppInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSysApp", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSysApp indicates an expected call of CreateSysApp
func (mr *MockPropertyServiceMockRecorder) CreateSysApp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSysApp", reflect.TypeOf((*MockPropertyService)(nil).CreateSysApp), arg0)
}

// DeleteProperty mocks base method
func (m *MockPropertyService) DeleteProperty(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProperty", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProperty indicates an expected call of DeleteProperty
func (mr *MockPropertyServiceMockRecorder) DeleteProperty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProperty", reflect.TypeOf((*MockPropertyService)(nil).DeleteProperty), arg0)
}

// DeleteSysApp mocks base method
func (m *MockPropertyService) DeleteSysApp(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSysApp", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSysApp indicates an expected call of DeleteSysApp
func (mr *MockPropertyServiceMockRecorder) DeleteSysApp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSysApp", reflect.TypeOf((*MockPropertyService)(nil).DeleteSysApp), arg0)
}

// GetProperty mocks base method
func (m *MockPropertyService) GetProperty(arg0 string) (*models.Property, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProperty", arg0)
	ret0, _ := ret[0].(*models.Property)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProperty indicates an expected call of GetProperty
func (mr *MockPropertyServiceMockRecorder) GetProperty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProperty", reflect.TypeOf((*MockPropertyService)(nil).GetProperty), arg0)
}

// GetPropertyValue mocks base method
func (m *MockPropertyService) GetPropertyValue(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPropertyValue", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPropertyValue indicates an expected call of GetPropertyValue
func (mr *MockPropertyServiceMockRecorder) GetPropertyValue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPropertyValue", reflect.TypeOf((*MockPropertyService)(nil).GetPropertyValue), arg0)
}

// GetSysApp mocks base method
func (m *MockPropertyService) GetSysApp(arg0 string) (*models.NodeSysAppInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSysApp", arg0)
	ret0, _ := ret[0].(*models.NodeSysAppInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSysApp indicates an expected call of GetSysApp
func (mr *MockPropertyServiceMockRecorder) GetSysApp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSysApp", reflect.TypeOf((*MockPropertyService)(nil).GetSysApp), arg0)
}

// ListOptionalSysApps mocks base method
func (m *MockPropertyService) ListOptionalSysApps() ([]models.NodeSysAppInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOptionalSysApps")
	ret0, _ := ret[0].([]models.NodeSysAppInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOptionalSysApps indicates an expected call of ListOptionalSysApps
func (mr *MockPropertyServiceMockRecorder) ListOptionalSysApps() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOptionalSysApps", reflect.TypeOf((*MockPropertyService)(nil).ListOptionalSysApps))
}

// ListProperty mocks base method
func (m *MockPropertyService) ListProperty(arg0 *models.Filter) ([]models.Property, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProperty", arg0)
	ret0, _ := ret[0].([]models.Property)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProperty indicates an expected call of ListProperty
func (mr *MockPropertyServiceMockRecorder) ListProperty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProperty", reflect.TypeOf((*MockPropertyService)(nil).ListProperty), arg0)
}

// ListSysApps mocks base method
func (m *MockPropertyService) ListSysApps() ([]models.NodeSysAppInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSysApps")
	ret0, _ := ret[0].([]models.NodeSysAppInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSysApps indicates an expected call of ListSysApps
func (mr *MockPropertyServiceMockRecorder) ListSysApps() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSysApps", reflect.TypeOf((*MockPropertyService)(nil).ListSysApps))
}

// UpdateProperty mocks base method
func (m *MockPropertyService) UpdateProperty(arg0 *models.Property) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProperty", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProperty indicates an expected call of UpdateProperty
func (mr *MockPropertyServiceMockRecorder) UpdateProperty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProperty", reflect.TypeOf((*MockPropertyService)(nil).UpdateProperty), arg0)
}

// UpdateSysApp mocks base method
func (m *MockPropertyService) UpdateSysApp(arg0 *models.NodeSysAppInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSysApp", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSysApp indicates an expected call of UpdateSysApp
func (mr *MockPropertyServiceMockRecorder) UpdateSysApp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSysApp", reflect.TypeOf((*MockPropertyService)(nil).UpdateSysApp), arg0)
}