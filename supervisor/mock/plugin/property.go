// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/baetyl/baetyl-cloud/v2/plugin (interfaces: Property)

// Package plugin is a generated GoMock package.
package plugin

import (
	models "github.com/baetyl/baetyl-cloud/v2/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProperty is a mock of Property interface
type MockProperty struct {
	ctrl     *gomock.Controller
	recorder *MockPropertyMockRecorder
}

// MockPropertyMockRecorder is the mock recorder for MockProperty
type MockPropertyMockRecorder struct {
	mock *MockProperty
}

// NewMockProperty creates a new mock instance
func NewMockProperty(ctrl *gomock.Controller) *MockProperty {
	mock := &MockProperty{ctrl: ctrl}
	mock.recorder = &MockPropertyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProperty) EXPECT() *MockPropertyMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockProperty) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockPropertyMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockProperty)(nil).Close))
}

// CountProperty mocks base method
func (m *MockProperty) CountProperty(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountProperty", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountProperty indicates an expected call of CountProperty
func (mr *MockPropertyMockRecorder) CountProperty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountProperty", reflect.TypeOf((*MockProperty)(nil).CountProperty), arg0)
}

// CreateProperty mocks base method
func (m *MockProperty) CreateProperty(arg0 *models.Property) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProperty", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProperty indicates an expected call of CreateProperty
func (mr *MockPropertyMockRecorder) CreateProperty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProperty", reflect.TypeOf((*MockProperty)(nil).CreateProperty), arg0)
}

// DeleteProperty mocks base method
func (m *MockProperty) DeleteProperty(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProperty", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProperty indicates an expected call of DeleteProperty
func (mr *MockPropertyMockRecorder) DeleteProperty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProperty", reflect.TypeOf((*MockProperty)(nil).DeleteProperty), arg0)
}

// GetProperty mocks base method
func (m *MockProperty) GetProperty(arg0 string) (*models.Property, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProperty", arg0)
	ret0, _ := ret[0].(*models.Property)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProperty indicates an expected call of GetProperty
func (mr *MockPropertyMockRecorder) GetProperty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProperty", reflect.TypeOf((*MockProperty)(nil).GetProperty), arg0)
}

// GetPropertyValue mocks base method
func (m *MockProperty) GetPropertyValue(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPropertyValue", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPropertyValue indicates an expected call of GetPropertyValue
func (mr *MockPropertyMockRecorder) GetPropertyValue(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPropertyValue", reflect.TypeOf((*MockProperty)(nil).GetPropertyValue), arg0)
}

// ListProperty mocks base method
func (m *MockProperty) ListProperty(arg0 *models.Filter) ([]models.Property, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProperty", arg0)
	ret0, _ := ret[0].([]models.Property)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProperty indicates an expected call of ListProperty
func (mr *MockPropertyMockRecorder) ListProperty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProperty", reflect.TypeOf((*MockProperty)(nil).ListProperty), arg0)
}

// UpdateProperty mocks base method
func (m *MockProperty) UpdateProperty(arg0 *models.Property) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProperty", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProperty indicates an expected call of UpdateProperty
func (mr *MockPropertyMockRecorder) UpdateProperty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProperty", reflect.TypeOf((*MockProperty)(nil).UpdateProperty), arg0)
}