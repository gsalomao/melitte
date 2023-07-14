// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	akira "github.com/gsalomao/akira"
	mock "github.com/stretchr/testify/mock"
)

// MockOnServerStartHook is an autogenerated mock type for the OnServerStartHook type
type MockOnServerStartHook struct {
	mock.Mock
}

type MockOnServerStartHook_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOnServerStartHook) EXPECT() *MockOnServerStartHook_Expecter {
	return &MockOnServerStartHook_Expecter{mock: &_m.Mock}
}

// Name provides a mock function with given fields:
func (_m *MockOnServerStartHook) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockOnServerStartHook_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockOnServerStartHook_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockOnServerStartHook_Expecter) Name() *MockOnServerStartHook_Name_Call {
	return &MockOnServerStartHook_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockOnServerStartHook_Name_Call) Run(run func()) *MockOnServerStartHook_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockOnServerStartHook_Name_Call) Return(_a0 string) *MockOnServerStartHook_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOnServerStartHook_Name_Call) RunAndReturn(run func() string) *MockOnServerStartHook_Name_Call {
	_c.Call.Return(run)
	return _c
}

// OnServerStart provides a mock function with given fields: s
func (_m *MockOnServerStartHook) OnServerStart(s *akira.Server) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*akira.Server) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockOnServerStartHook_OnServerStart_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnServerStart'
type MockOnServerStartHook_OnServerStart_Call struct {
	*mock.Call
}

// OnServerStart is a helper method to define mock.On call
//   - s *akira.Server
func (_e *MockOnServerStartHook_Expecter) OnServerStart(s interface{}) *MockOnServerStartHook_OnServerStart_Call {
	return &MockOnServerStartHook_OnServerStart_Call{Call: _e.mock.On("OnServerStart", s)}
}

func (_c *MockOnServerStartHook_OnServerStart_Call) Run(run func(s *akira.Server)) *MockOnServerStartHook_OnServerStart_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*akira.Server))
	})
	return _c
}

func (_c *MockOnServerStartHook_OnServerStart_Call) Return(_a0 error) *MockOnServerStartHook_OnServerStart_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOnServerStartHook_OnServerStart_Call) RunAndReturn(run func(*akira.Server) error) *MockOnServerStartHook_OnServerStart_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOnServerStartHook creates a new instance of MockOnServerStartHook. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOnServerStartHook(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOnServerStartHook {
	mock := &MockOnServerStartHook{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
