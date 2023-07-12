// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	"github.com/gsalomao/akira"
	"github.com/stretchr/testify/mock"

	"net"
)

// MockOnConnectionFunc is an autogenerated mock type for the OnConnectionFunc type
type MockOnConnectionFunc struct {
	mock.Mock
}

type MockOnConnectionFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOnConnectionFunc) EXPECT() *MockOnConnectionFunc_Expecter {
	return &MockOnConnectionFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *MockOnConnectionFunc) Execute(_a0 akira.Listener, _a1 net.Conn) {
	_m.Called(_a0, _a1)
}

// MockOnConnectionFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockOnConnectionFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 akira.Listener
//   - _a1 net.Conn
func (_e *MockOnConnectionFunc_Expecter) Execute(_a0 interface{}, _a1 interface{}) *MockOnConnectionFunc_Execute_Call {
	return &MockOnConnectionFunc_Execute_Call{Call: _e.mock.On("Execute", _a0, _a1)}
}

func (_c *MockOnConnectionFunc_Execute_Call) Run(run func(_a0 akira.Listener, _a1 net.Conn)) *MockOnConnectionFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(akira.Listener), args[1].(net.Conn))
	})
	return _c
}

func (_c *MockOnConnectionFunc_Execute_Call) Return() *MockOnConnectionFunc_Execute_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockOnConnectionFunc_Execute_Call) RunAndReturn(run func(akira.Listener, net.Conn)) *MockOnConnectionFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOnConnectionFunc creates a new instance of MockOnConnectionFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOnConnectionFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOnConnectionFunc {
	mock := &MockOnConnectionFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
