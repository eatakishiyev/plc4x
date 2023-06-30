/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.30.16. DO NOT EDIT.

package plc4go

import mock "github.com/stretchr/testify/mock"

// MockWithDiscoveryOption is an autogenerated mock type for the WithDiscoveryOption type
type MockWithDiscoveryOption struct {
	mock.Mock
}

type MockWithDiscoveryOption_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWithDiscoveryOption) EXPECT() *MockWithDiscoveryOption_Expecter {
	return &MockWithDiscoveryOption_Expecter{mock: &_m.Mock}
}

// isDiscoveryOption provides a mock function with given fields:
func (_m *MockWithDiscoveryOption) isDiscoveryOption() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockWithDiscoveryOption_isDiscoveryOption_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'isDiscoveryOption'
type MockWithDiscoveryOption_isDiscoveryOption_Call struct {
	*mock.Call
}

// isDiscoveryOption is a helper method to define mock.On call
func (_e *MockWithDiscoveryOption_Expecter) isDiscoveryOption() *MockWithDiscoveryOption_isDiscoveryOption_Call {
	return &MockWithDiscoveryOption_isDiscoveryOption_Call{Call: _e.mock.On("isDiscoveryOption")}
}

func (_c *MockWithDiscoveryOption_isDiscoveryOption_Call) Run(run func()) *MockWithDiscoveryOption_isDiscoveryOption_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockWithDiscoveryOption_isDiscoveryOption_Call) Return(_a0 bool) *MockWithDiscoveryOption_isDiscoveryOption_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWithDiscoveryOption_isDiscoveryOption_Call) RunAndReturn(run func() bool) *MockWithDiscoveryOption_isDiscoveryOption_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWithDiscoveryOption creates a new instance of MockWithDiscoveryOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWithDiscoveryOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWithDiscoveryOption {
	mock := &MockWithDiscoveryOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
