// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/deploy/cloudformation/stack/workload.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	cloudformation "github.com/aws/aws-sdk-go/service/cloudformation"
	gomock "github.com/golang/mock/gomock"
)

// MockNestedStackConfigurer is a mock of NestedStackConfigurer interface.
type MockNestedStackConfigurer struct {
	ctrl     *gomock.Controller
	recorder *MockNestedStackConfigurerMockRecorder
}

// MockNestedStackConfigurerMockRecorder is the mock recorder for MockNestedStackConfigurer.
type MockNestedStackConfigurerMockRecorder struct {
	mock *MockNestedStackConfigurer
}

// NewMockNestedStackConfigurer creates a new mock instance.
func NewMockNestedStackConfigurer(ctrl *gomock.Controller) *MockNestedStackConfigurer {
	mock := &MockNestedStackConfigurer{ctrl: ctrl}
	mock.recorder = &MockNestedStackConfigurerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNestedStackConfigurer) EXPECT() *MockNestedStackConfigurerMockRecorder {
	return m.recorder
}

// Parameters mocks base method.
func (m *MockNestedStackConfigurer) Parameters() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parameters indicates an expected call of Parameters.
func (mr *MockNestedStackConfigurerMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*MockNestedStackConfigurer)(nil).Parameters))
}

// Template mocks base method.
func (m *MockNestedStackConfigurer) Template() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Template")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Template indicates an expected call of Template.
func (mr *MockNestedStackConfigurerMockRecorder) Template() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Template", reflect.TypeOf((*MockNestedStackConfigurer)(nil).Template))
}

// Mocklocation is a mock of location interface.
type Mocklocation struct {
	ctrl     *gomock.Controller
	recorder *MocklocationMockRecorder
}

// MocklocationMockRecorder is the mock recorder for Mocklocation.
type MocklocationMockRecorder struct {
	mock *Mocklocation
}

// NewMocklocation creates a new mock instance.
func NewMocklocation(ctrl *gomock.Controller) *Mocklocation {
	mock := &Mocklocation{ctrl: ctrl}
	mock.recorder = &MocklocationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocklocation) EXPECT() *MocklocationMockRecorder {
	return m.recorder
}

// GetLocation mocks base method.
func (m *Mocklocation) GetLocation() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocation")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLocation indicates an expected call of GetLocation.
func (mr *MocklocationMockRecorder) GetLocation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocation", reflect.TypeOf((*Mocklocation)(nil).GetLocation))
}

// Mockuploadable is a mock of uploadable interface.
type Mockuploadable struct {
	ctrl     *gomock.Controller
	recorder *MockuploadableMockRecorder
}

// MockuploadableMockRecorder is the mock recorder for Mockuploadable.
type MockuploadableMockRecorder struct {
	mock *Mockuploadable
}

// NewMockuploadable creates a new mock instance.
func NewMockuploadable(ctrl *gomock.Controller) *Mockuploadable {
	mock := &Mockuploadable{ctrl: ctrl}
	mock.recorder = &MockuploadableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockuploadable) EXPECT() *MockuploadableMockRecorder {
	return m.recorder
}

// ArtifactPath mocks base method.
func (m *Mockuploadable) ArtifactPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ArtifactPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// ArtifactPath indicates an expected call of ArtifactPath.
func (mr *MockuploadableMockRecorder) ArtifactPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArtifactPath", reflect.TypeOf((*Mockuploadable)(nil).ArtifactPath))
}

// Name mocks base method.
func (m *Mockuploadable) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockuploadableMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*Mockuploadable)(nil).Name))
}

// MocktemplateConfigurer is a mock of templateConfigurer interface.
type MocktemplateConfigurer struct {
	ctrl     *gomock.Controller
	recorder *MocktemplateConfigurerMockRecorder
}

// MocktemplateConfigurerMockRecorder is the mock recorder for MocktemplateConfigurer.
type MocktemplateConfigurerMockRecorder struct {
	mock *MocktemplateConfigurer
}

// NewMocktemplateConfigurer creates a new mock instance.
func NewMocktemplateConfigurer(ctrl *gomock.Controller) *MocktemplateConfigurer {
	mock := &MocktemplateConfigurer{ctrl: ctrl}
	mock.recorder = &MocktemplateConfigurerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktemplateConfigurer) EXPECT() *MocktemplateConfigurerMockRecorder {
	return m.recorder
}

// Parameters mocks base method.
func (m *MocktemplateConfigurer) Parameters() ([]*cloudformation.Parameter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parameters")
	ret0, _ := ret[0].([]*cloudformation.Parameter)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parameters indicates an expected call of Parameters.
func (mr *MocktemplateConfigurerMockRecorder) Parameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parameters", reflect.TypeOf((*MocktemplateConfigurer)(nil).Parameters))
}

// Tags mocks base method.
func (m *MocktemplateConfigurer) Tags() []*cloudformation.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tags")
	ret0, _ := ret[0].([]*cloudformation.Tag)
	return ret0
}

// Tags indicates an expected call of Tags.
func (mr *MocktemplateConfigurerMockRecorder) Tags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tags", reflect.TypeOf((*MocktemplateConfigurer)(nil).Tags))
}
