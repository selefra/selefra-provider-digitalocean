package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockStorageService struct {
	ctrl     *gomock.Controller
	recorder *MockStorageServiceMockRecorder
}

type MockStorageServiceMockRecorder struct {
	mock *MockStorageService
}

func NewMockStorageService(ctrl *gomock.Controller) *MockStorageService {
	mock := &MockStorageService{ctrl: ctrl}
	mock.recorder = &MockStorageServiceMockRecorder{mock}
	return mock
}

func (m *MockStorageService) EXPECT() *MockStorageServiceMockRecorder {
	return m.recorder
}

func (m *MockStorageService) ListVolumes(arg0 context.Context, arg1 *godo.ListVolumeParams) ([]godo.Volume, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.ListVolumes, arg0, arg1)
	ret0, _ := ret[0].([]godo.Volume)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockStorageServiceMockRecorder) ListVolumes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.ListVolumes, reflect.TypeOf((*MockStorageService)(nil).ListVolumes), arg0, arg1)
}
