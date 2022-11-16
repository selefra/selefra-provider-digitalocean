package mocks

import (
	context "context"
	reflect "reflect"

	"github.com/selefra/selefra-provider-digitalocean/constants"

	godo "github.com/digitalocean/godo"
	gomock "github.com/golang/mock/gomock"
)

type MockLoadBalancersService struct {
	ctrl     *gomock.Controller
	recorder *MockLoadBalancersServiceMockRecorder
}

type MockLoadBalancersServiceMockRecorder struct {
	mock *MockLoadBalancersService
}

func NewMockLoadBalancersService(ctrl *gomock.Controller) *MockLoadBalancersService {
	mock := &MockLoadBalancersService{ctrl: ctrl}
	mock.recorder = &MockLoadBalancersServiceMockRecorder{mock}
	return mock
}

func (m *MockLoadBalancersService) EXPECT() *MockLoadBalancersServiceMockRecorder {
	return m.recorder
}

func (m *MockLoadBalancersService) List(arg0 context.Context, arg1 *godo.ListOptions) ([]godo.LoadBalancer, *godo.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, constants.List, arg0, arg1)
	ret0, _ := ret[0].([]godo.LoadBalancer)
	ret1, _ := ret[1].(*godo.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockLoadBalancersServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, constants.List, reflect.TypeOf((*MockLoadBalancersService)(nil).List), arg0, arg1)
}
