// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	"github.com/gsalomao/akira/packet"
	"github.com/stretchr/testify/mock"
)

// MockPacketDecodable is an autogenerated mock type for the PacketDecodable type
type MockPacketDecodable struct {
	mock.Mock
}

type MockPacketDecodable_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPacketDecodable) EXPECT() *MockPacketDecodable_Expecter {
	return &MockPacketDecodable_Expecter{mock: &_m.Mock}
}

// Decode provides a mock function with given fields: buf, header
func (_m *MockPacketDecodable) Decode(buf []byte, header packet.FixedHeader) (int, error) {
	ret := _m.Called(buf, header)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, packet.FixedHeader) (int, error)); ok {
		return rf(buf, header)
	}
	if rf, ok := ret.Get(0).(func([]byte, packet.FixedHeader) int); ok {
		r0 = rf(buf, header)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte, packet.FixedHeader) error); ok {
		r1 = rf(buf, header)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPacketDecodable_Decode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decode'
type MockPacketDecodable_Decode_Call struct {
	*mock.Call
}

// Decode is a helper method to define mock.On call
//   - buf []byte
//   - header packet.FixedHeader
func (_e *MockPacketDecodable_Expecter) Decode(buf interface{}, header interface{}) *MockPacketDecodable_Decode_Call {
	return &MockPacketDecodable_Decode_Call{Call: _e.mock.On("Decode", buf, header)}
}

func (_c *MockPacketDecodable_Decode_Call) Run(run func(buf []byte, header packet.FixedHeader)) *MockPacketDecodable_Decode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(packet.FixedHeader))
	})
	return _c
}

func (_c *MockPacketDecodable_Decode_Call) Return(n int, err error) *MockPacketDecodable_Decode_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *MockPacketDecodable_Decode_Call) RunAndReturn(run func([]byte, packet.FixedHeader) (int, error)) *MockPacketDecodable_Decode_Call {
	_c.Call.Return(run)
	return _c
}

// Size provides a mock function with given fields:
func (_m *MockPacketDecodable) Size() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// MockPacketDecodable_Size_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Size'
type MockPacketDecodable_Size_Call struct {
	*mock.Call
}

// Size is a helper method to define mock.On call
func (_e *MockPacketDecodable_Expecter) Size() *MockPacketDecodable_Size_Call {
	return &MockPacketDecodable_Size_Call{Call: _e.mock.On("Size")}
}

func (_c *MockPacketDecodable_Size_Call) Run(run func()) *MockPacketDecodable_Size_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPacketDecodable_Size_Call) Return(_a0 int) *MockPacketDecodable_Size_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPacketDecodable_Size_Call) RunAndReturn(run func() int) *MockPacketDecodable_Size_Call {
	_c.Call.Return(run)
	return _c
}

// Type provides a mock function with given fields:
func (_m *MockPacketDecodable) Type() packet.Type {
	ret := _m.Called()

	var r0 packet.Type
	if rf, ok := ret.Get(0).(func() packet.Type); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(packet.Type)
	}

	return r0
}

// MockPacketDecodable_Type_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Type'
type MockPacketDecodable_Type_Call struct {
	*mock.Call
}

// Type is a helper method to define mock.On call
func (_e *MockPacketDecodable_Expecter) Type() *MockPacketDecodable_Type_Call {
	return &MockPacketDecodable_Type_Call{Call: _e.mock.On("Type")}
}

func (_c *MockPacketDecodable_Type_Call) Run(run func()) *MockPacketDecodable_Type_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPacketDecodable_Type_Call) Return(_a0 packet.Type) *MockPacketDecodable_Type_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPacketDecodable_Type_Call) RunAndReturn(run func() packet.Type) *MockPacketDecodable_Type_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPacketDecodable creates a new instance of MockPacketDecodable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPacketDecodable(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPacketDecodable {
	mock := &MockPacketDecodable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
