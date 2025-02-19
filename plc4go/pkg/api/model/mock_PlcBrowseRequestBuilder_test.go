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

// Code generated by mockery v2.42.2. DO NOT EDIT.

package model

import mock "github.com/stretchr/testify/mock"

// MockPlcBrowseRequestBuilder is an autogenerated mock type for the PlcBrowseRequestBuilder type
type MockPlcBrowseRequestBuilder struct {
	mock.Mock
}

type MockPlcBrowseRequestBuilder_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcBrowseRequestBuilder) EXPECT() *MockPlcBrowseRequestBuilder_Expecter {
	return &MockPlcBrowseRequestBuilder_Expecter{mock: &_m.Mock}
}

// AddQuery provides a mock function with given fields: name, query
func (_m *MockPlcBrowseRequestBuilder) AddQuery(name string, query string) PlcBrowseRequestBuilder {
	ret := _m.Called(name, query)

	if len(ret) == 0 {
		panic("no return value specified for AddQuery")
	}

	var r0 PlcBrowseRequestBuilder
	if rf, ok := ret.Get(0).(func(string, string) PlcBrowseRequestBuilder); ok {
		r0 = rf(name, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcBrowseRequestBuilder)
		}
	}

	return r0
}

// MockPlcBrowseRequestBuilder_AddQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddQuery'
type MockPlcBrowseRequestBuilder_AddQuery_Call struct {
	*mock.Call
}

// AddQuery is a helper method to define mock.On call
//   - name string
//   - query string
func (_e *MockPlcBrowseRequestBuilder_Expecter) AddQuery(name interface{}, query interface{}) *MockPlcBrowseRequestBuilder_AddQuery_Call {
	return &MockPlcBrowseRequestBuilder_AddQuery_Call{Call: _e.mock.On("AddQuery", name, query)}
}

func (_c *MockPlcBrowseRequestBuilder_AddQuery_Call) Run(run func(name string, query string)) *MockPlcBrowseRequestBuilder_AddQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockPlcBrowseRequestBuilder_AddQuery_Call) Return(_a0 PlcBrowseRequestBuilder) *MockPlcBrowseRequestBuilder_AddQuery_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcBrowseRequestBuilder_AddQuery_Call) RunAndReturn(run func(string, string) PlcBrowseRequestBuilder) *MockPlcBrowseRequestBuilder_AddQuery_Call {
	_c.Call.Return(run)
	return _c
}

// Build provides a mock function with given fields:
func (_m *MockPlcBrowseRequestBuilder) Build() (PlcBrowseRequest, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Build")
	}

	var r0 PlcBrowseRequest
	var r1 error
	if rf, ok := ret.Get(0).(func() (PlcBrowseRequest, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() PlcBrowseRequest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcBrowseRequest)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPlcBrowseRequestBuilder_Build_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Build'
type MockPlcBrowseRequestBuilder_Build_Call struct {
	*mock.Call
}

// Build is a helper method to define mock.On call
func (_e *MockPlcBrowseRequestBuilder_Expecter) Build() *MockPlcBrowseRequestBuilder_Build_Call {
	return &MockPlcBrowseRequestBuilder_Build_Call{Call: _e.mock.On("Build")}
}

func (_c *MockPlcBrowseRequestBuilder_Build_Call) Run(run func()) *MockPlcBrowseRequestBuilder_Build_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcBrowseRequestBuilder_Build_Call) Return(_a0 PlcBrowseRequest, _a1 error) *MockPlcBrowseRequestBuilder_Build_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPlcBrowseRequestBuilder_Build_Call) RunAndReturn(run func() (PlcBrowseRequest, error)) *MockPlcBrowseRequestBuilder_Build_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPlcBrowseRequestBuilder) String() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for String")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcBrowseRequestBuilder_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPlcBrowseRequestBuilder_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPlcBrowseRequestBuilder_Expecter) String() *MockPlcBrowseRequestBuilder_String_Call {
	return &MockPlcBrowseRequestBuilder_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPlcBrowseRequestBuilder_String_Call) Run(run func()) *MockPlcBrowseRequestBuilder_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcBrowseRequestBuilder_String_Call) Return(_a0 string) *MockPlcBrowseRequestBuilder_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcBrowseRequestBuilder_String_Call) RunAndReturn(run func() string) *MockPlcBrowseRequestBuilder_String_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcBrowseRequestBuilder creates a new instance of MockPlcBrowseRequestBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcBrowseRequestBuilder(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcBrowseRequestBuilder {
	mock := &MockPlcBrowseRequestBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
