// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	akira "github.com/gsalomao/akira"
	mock "github.com/stretchr/testify/mock"
)

// MockOnServerStartFailedHook is an autogenerated mock type for the OnServerStartFailedHook type
type MockOnServerStartFailedHook struct {
	mock.Mock
}

type MockOnServerStartFailedHook_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOnServerStartFailedHook) EXPECT() *MockOnServerStartFailedHook_Expecter {
	return &MockOnServerStartFailedHook_Expecter{mock: &_m.Mock}
}

// Name provides a mock function with given fields:
func (_m *MockOnServerStartFailedHook) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockOnServerStartFailedHook_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockOnServerStartFailedHook_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockOnServerStartFailedHook_Expecter) Name() *MockOnServerStartFailedHook_Name_Call {
	return &MockOnServerStartFailedHook_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockOnServerStartFailedHook_Name_Call) Run(run func()) *MockOnServerStartFailedHook_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockOnServerStartFailedHook_Name_Call) Return(_a0 string) *MockOnServerStartFailedHook_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOnServerStartFailedHook_Name_Call) RunAndReturn(run func() string) *MockOnServerStartFailedHook_Name_Call {
	_c.Call.Return(run)
	return _c
}

// OnServerStartFailed provides a mock function with given fields: s, err
func (_m *MockOnServerStartFailedHook) OnServerStartFailed(s *akira.Server, err error) {
	_m.Called(s, err)
}

// MockOnServerStartFailedHook_OnServerStartFailed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OnServerStartFailed'
type MockOnServerStartFailedHook_OnServerStartFailed_Call struct {
	*mock.Call
}

// OnServerStartFailed is a helper method to define mock.On call
//   - s *akira.Server
//   - err error
func (_e *MockOnServerStartFailedHook_Expecter) OnServerStartFailed(s interface{}, err interface{}) *MockOnServerStartFailedHook_OnServerStartFailed_Call {
	return &MockOnServerStartFailedHook_OnServerStartFailed_Call{Call: _e.mock.On("OnServerStartFailed", s, err)}
}

func (_c *MockOnServerStartFailedHook_OnServerStartFailed_Call) Run(run func(s *akira.Server, err error)) *MockOnServerStartFailedHook_OnServerStartFailed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*akira.Server), args[1].(error))
	})
	return _c
}

func (_c *MockOnServerStartFailedHook_OnServerStartFailed_Call) Return() *MockOnServerStartFailedHook_OnServerStartFailed_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockOnServerStartFailedHook_OnServerStartFailed_Call) RunAndReturn(run func(*akira.Server, error)) *MockOnServerStartFailedHook_OnServerStartFailed_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOnServerStartFailedHook creates a new instance of MockOnServerStartFailedHook. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOnServerStartFailedHook(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOnServerStartFailedHook {
	mock := &MockOnServerStartFailedHook{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
