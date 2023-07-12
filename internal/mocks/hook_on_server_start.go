// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	"github.com/gsalomao/akira"
	"github.com/stretchr/testify/mock"
)

// MockHookOnServerStart is an autogenerated mock type for the HookOnServerStart type
type MockHookOnServerStart struct {
	mock.Mock
}

type MockHookOnServerStart_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHookOnServerStart) EXPECT() *MockHookOnServerStart_Expecter {
	return &MockHookOnServerStart_Expecter{mock: &_m.Mock}
}

// Name provides a mock function with given fields:
func (_m *MockHookOnServerStart) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockHookOnServerStart_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockHookOnServerStart_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockHookOnServerStart_Expecter) Name() *MockHookOnServerStart_Name_Call {
	return &MockHookOnServerStart_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockHookOnServerStart_Name_Call) Run(run func()) *MockHookOnServerStart_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockHookOnServerStart_Name_Call) Return(_a0 string) *MockHookOnServerStart_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHookOnServerStart_Name_Call) RunAndReturn(run func() string) *MockHookOnServerStart_Name_Call {
	_c.Call.Return(run)
	return _c
}

// OnServerStart provides a mock function with given fields: s
func (_m *MockHookOnServerStart) OnServerStart(s *akira.Server) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*akira.Server) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockHookOnServerStart_OnServerStart_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnServerStart'
type MockHookOnServerStart_OnServerStart_Call struct {
	*mock.Call
}

// OnServerStart is a helper method to define mock.On call
//   - s *akira.Server
func (_e *MockHookOnServerStart_Expecter) OnServerStart(s interface{}) *MockHookOnServerStart_OnServerStart_Call {
	return &MockHookOnServerStart_OnServerStart_Call{Call: _e.mock.On("OnServerStart", s)}
}

func (_c *MockHookOnServerStart_OnServerStart_Call) Run(run func(s *akira.Server)) *MockHookOnServerStart_OnServerStart_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*akira.Server))
	})
	return _c
}

func (_c *MockHookOnServerStart_OnServerStart_Call) Return(_a0 error) *MockHookOnServerStart_OnServerStart_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHookOnServerStart_OnServerStart_Call) RunAndReturn(run func(*akira.Server) error) *MockHookOnServerStart_OnServerStart_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockHookOnServerStart creates a new instance of MockHookOnServerStart. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHookOnServerStart(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHookOnServerStart {
	mock := &MockHookOnServerStart{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
