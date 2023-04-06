// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// StructWithTag is an autogenerated mock type for the StructWithTag type
type StructWithTag struct {
	mock.Mock
}

type StructWithTag_Expecter struct {
	mock *mock.Mock
}

func (_m *StructWithTag) EXPECT() *StructWithTag_Expecter {
	return &StructWithTag_Expecter{mock: &_m.Mock}
}

// MethodA provides a mock function with given fields: v
func (_m *StructWithTag) MethodA(v *struct {
	FieldA int `json:"field_a"`
	FieldB int `json:"field_b" xml:"field_b"`
}) *struct {
	FieldC int `json:"field_c"`
	FieldD int `json:"field_d" xml:"field_d"`
} {
	ret := _m.Called(v)

	var r0 *struct {
		FieldC int `json:"field_c"`
		FieldD int `json:"field_d" xml:"field_d"`
	}
	if rf, ok := ret.Get(0).(func(*struct {
		FieldA int `json:"field_a"`
		FieldB int `json:"field_b" xml:"field_b"`
	}) *struct {
		FieldC int `json:"field_c"`
		FieldD int `json:"field_d" xml:"field_d"`
	}); ok {
		r0 = rf(v)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*struct {
				FieldC int `json:"field_c"`
				FieldD int `json:"field_d" xml:"field_d"`
			})
		}
	}

	return r0
}

// StructWithTag_MethodA_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MethodA'
type StructWithTag_MethodA_Call struct {
	*mock.Call
}

// MethodA is a helper method to define mock.On call
//   - v *struct{FieldA int `json:"field_a"`;FieldB int `json:"field_b" xml:"field_b"`}
func (_e *StructWithTag_Expecter) MethodA(v interface{}) *StructWithTag_MethodA_Call {
	return &StructWithTag_MethodA_Call{Call: _e.mock.On("MethodA", v)}
}

func (_c *StructWithTag_MethodA_Call) Run(run func(v *struct {
	FieldA int `json:"field_a"`
	FieldB int `json:"field_b" xml:"field_b"`
})) *StructWithTag_MethodA_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*struct {
			FieldA int `json:"field_a"`
			FieldB int `json:"field_b" xml:"field_b"`
		}))
	})
	return _c
}

func (_c *StructWithTag_MethodA_Call) Return(_a0 *struct {
	FieldC int `json:"field_c"`
	FieldD int `json:"field_d" xml:"field_d"`
}) *StructWithTag_MethodA_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StructWithTag_MethodA_Call) RunAndReturn(run func(*struct {
	FieldA int `json:"field_a"`
	FieldB int `json:"field_b" xml:"field_b"`
}) *struct {
	FieldC int `json:"field_c"`
	FieldD int `json:"field_d" xml:"field_d"`
}) *StructWithTag_MethodA_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewStructWithTag interface {
	mock.TestingT
	Cleanup(func())
}

// NewStructWithTag creates a new instance of StructWithTag. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStructWithTag(t mockConstructorTestingTNewStructWithTag) *StructWithTag {
	mock := &StructWithTag{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
