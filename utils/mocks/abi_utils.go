// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	io "io"

	abi "github.com/ethereum/go-ethereum/accounts/abi"

	mock "github.com/stretchr/testify/mock"
)

// ABIUtils is an autogenerated mock type for the ABIUtils type
type ABIUtils struct {
	mock.Mock
}

// Pack provides a mock function with given fields: _a0, _a1, _a2
func (_m *ABIUtils) Pack(_a0 abi.ABI, _a1 string, _a2 ...interface{}) ([]byte, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(abi.ABI, string, ...interface{}) []byte); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(abi.ABI, string, ...interface{}) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Parse provides a mock function with given fields: _a0
func (_m *ABIUtils) Parse(_a0 io.Reader) (abi.ABI, error) {
	ret := _m.Called(_a0)

	var r0 abi.ABI
	if rf, ok := ret.Get(0).(func(io.Reader) abi.ABI); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(abi.ABI)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Reader) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}