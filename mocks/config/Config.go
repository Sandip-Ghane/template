// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Config is an autogenerated mock type for the Config type
type Config struct {
	mock.Mock
}

// GetBool provides a mock function with given fields: key
func (_m *Config) GetBool(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetInt provides a mock function with given fields: key
func (_m *Config) GetInt(key string) int {
	ret := _m.Called(key)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetString provides a mock function with given fields: key
func (_m *Config) GetString(key string) string {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetStringSlice provides a mock function with given fields: key
func (_m *Config) GetStringSlice(key string) []string {
	ret := _m.Called(key)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Init provides a mock function with given fields:
func (_m *Config) Init() {
	_m.Called()
}