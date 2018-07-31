// Code generated by mockery v1.0.0. DO NOT EDIT.
package p2p

import mock "github.com/stretchr/testify/mock"
import net "github.com/libp2p/go-libp2p-net"
import protocol "github.com/libp2p/go-libp2p-protocol"
import time "time"

// Stream is an autogenerated mock type for the Stream type
type Stream struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Stream) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Conn provides a mock function with given fields:
func (_m *Stream) Conn() net.Conn {
	ret := _m.Called()

	var r0 net.Conn
	if rf, ok := ret.Get(0).(func() net.Conn); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Conn)
		}
	}

	return r0
}

// Protocol provides a mock function with given fields:
func (_m *Stream) Protocol() protocol.ID {
	ret := _m.Called()

	var r0 protocol.ID
	if rf, ok := ret.Get(0).(func() protocol.ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(protocol.ID)
	}

	return r0
}

// Read provides a mock function with given fields: p
func (_m *Stream) Read(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Reset provides a mock function with given fields:
func (_m *Stream) Reset() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDeadline provides a mock function with given fields: _a0
func (_m *Stream) SetDeadline(_a0 time.Time) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetProtocol provides a mock function with given fields: _a0
func (_m *Stream) SetProtocol(_a0 protocol.ID) {
	_m.Called(_a0)
}

// SetReadDeadline provides a mock function with given fields: _a0
func (_m *Stream) SetReadDeadline(_a0 time.Time) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetWriteDeadline provides a mock function with given fields: _a0
func (_m *Stream) SetWriteDeadline(_a0 time.Time) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(time.Time) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Write provides a mock function with given fields: p
func (_m *Stream) Write(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}