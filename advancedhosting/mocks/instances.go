/*
Copyright 2021 Advanced Hosting

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mocks

import (
	context "context"
	reflect "reflect"

	ah "github.com/advancedhosting/advancedhosting-api-go/ah"
	gomock "github.com/golang/mock/gomock"
)

// MockInstancesAPI is a mock of InstancesAPI interface.
type MockInstancesAPI struct {
	ctrl     *gomock.Controller
	recorder *MockInstancesAPIMockRecorder
}

// MockInstancesAPIMockRecorder is the mock recorder for MockInstancesAPI.
type MockInstancesAPIMockRecorder struct {
	mock *MockInstancesAPI
}

// NewMockInstancesAPI creates a new mock instance.
func NewMockInstancesAPI(ctrl *gomock.Controller) *MockInstancesAPI {
	mock := &MockInstancesAPI{ctrl: ctrl}
	mock.recorder = &MockInstancesAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstancesAPI) EXPECT() *MockInstancesAPIMockRecorder {
	return m.recorder
}

// ActionInfo mocks base method.
func (m *MockInstancesAPI) ActionInfo(arg0 context.Context, arg1, arg2 string) (*ah.InstanceAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActionInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(*ah.InstanceAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActionInfo indicates an expected call of ActionInfo.
func (mr *MockInstancesAPIMockRecorder) ActionInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActionInfo", reflect.TypeOf((*MockInstancesAPI)(nil).ActionInfo), arg0, arg1, arg2)
}

// Actions mocks base method.
func (m *MockInstancesAPI) Actions(arg0 context.Context, arg1 string) ([]ah.InstanceAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Actions", arg0, arg1)
	ret0, _ := ret[0].([]ah.InstanceAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Actions indicates an expected call of Actions.
func (mr *MockInstancesAPIMockRecorder) Actions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Actions", reflect.TypeOf((*MockInstancesAPI)(nil).Actions), arg0, arg1)
}

// AttachVolume mocks base method.
func (m *MockInstancesAPI) AttachVolume(arg0 context.Context, arg1, arg2 string) (*ah.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(*ah.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachVolume indicates an expected call of AttachVolume.
func (mr *MockInstancesAPIMockRecorder) AttachVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachVolume", reflect.TypeOf((*MockInstancesAPI)(nil).AttachVolume), arg0, arg1, arg2)
}

// AvailableVolumes mocks base method.
func (m *MockInstancesAPI) AvailableVolumes(arg0 context.Context, arg1 string, arg2 *ah.ListOptions) ([]ah.Volume, *ah.Meta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailableVolumes", arg0, arg1, arg2)
	ret0, _ := ret[0].([]ah.Volume)
	ret1, _ := ret[1].(*ah.Meta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AvailableVolumes indicates an expected call of AvailableVolumes.
func (mr *MockInstancesAPIMockRecorder) AvailableVolumes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailableVolumes", reflect.TypeOf((*MockInstancesAPI)(nil).AvailableVolumes), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockInstancesAPI) Create(arg0 context.Context, arg1 *ah.InstanceCreateRequest) (*ah.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*ah.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockInstancesAPIMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockInstancesAPI)(nil).Create), arg0, arg1)
}

// CreateBackup mocks base method.
func (m *MockInstancesAPI) CreateBackup(arg0 context.Context, arg1, arg2 string) (*ah.InstanceAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBackup", arg0, arg1, arg2)
	ret0, _ := ret[0].(*ah.InstanceAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBackup indicates an expected call of CreateBackup.
func (mr *MockInstancesAPIMockRecorder) CreateBackup(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBackup", reflect.TypeOf((*MockInstancesAPI)(nil).CreateBackup), arg0, arg1, arg2)
}

// Destroy mocks base method.
func (m *MockInstancesAPI) Destroy(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockInstancesAPIMockRecorder) Destroy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockInstancesAPI)(nil).Destroy), arg0, arg1)
}

// DetachVolume mocks base method.
func (m *MockInstancesAPI) DetachVolume(arg0 context.Context, arg1, arg2 string) (*ah.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(*ah.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetachVolume indicates an expected call of DetachVolume.
func (mr *MockInstancesAPIMockRecorder) DetachVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachVolume", reflect.TypeOf((*MockInstancesAPI)(nil).DetachVolume), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockInstancesAPI) Get(arg0 context.Context, arg1 string) (*ah.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*ah.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInstancesAPIMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInstancesAPI)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockInstancesAPI) List(arg0 context.Context, arg1 *ah.ListOptions) ([]ah.Instance, *ah.Meta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]ah.Instance)
	ret1, _ := ret[1].(*ah.Meta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockInstancesAPIMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInstancesAPI)(nil).List), arg0, arg1)
}

// PowerOff mocks base method.
func (m *MockInstancesAPI) PowerOff(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PowerOff", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PowerOff indicates an expected call of PowerOff.
func (mr *MockInstancesAPIMockRecorder) PowerOff(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PowerOff", reflect.TypeOf((*MockInstancesAPI)(nil).PowerOff), arg0, arg1)
}

// Rename mocks base method.
func (m *MockInstancesAPI) Rename(arg0 context.Context, arg1, arg2 string) (*ah.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rename", arg0, arg1, arg2)
	ret0, _ := ret[0].(*ah.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rename indicates an expected call of Rename.
func (mr *MockInstancesAPIMockRecorder) Rename(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rename", reflect.TypeOf((*MockInstancesAPI)(nil).Rename), arg0, arg1, arg2)
}

// SetPrimaryIP mocks base method.
func (m *MockInstancesAPI) SetPrimaryIP(arg0 context.Context, arg1, arg2 string) (*ah.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPrimaryIP", arg0, arg1, arg2)
	ret0, _ := ret[0].(*ah.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetPrimaryIP indicates an expected call of SetPrimaryIP.
func (mr *MockInstancesAPIMockRecorder) SetPrimaryIP(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPrimaryIP", reflect.TypeOf((*MockInstancesAPI)(nil).SetPrimaryIP), arg0, arg1, arg2)
}

// Shutdown mocks base method.
func (m *MockInstancesAPI) Shutdown(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockInstancesAPIMockRecorder) Shutdown(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockInstancesAPI)(nil).Shutdown), arg0, arg1)
}

// Upgrade mocks base method.
func (m *MockInstancesAPI) Upgrade(arg0 context.Context, arg1 string, arg2 *ah.InstanceUpgradeRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upgrade indicates an expected call of Upgrade.
func (mr *MockInstancesAPIMockRecorder) Upgrade(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockInstancesAPI)(nil).Upgrade), arg0, arg1, arg2)
}
