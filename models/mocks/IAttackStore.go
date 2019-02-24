// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "vegeta-server/models"

// IAttackStore is an autogenerated mock type for the IAttackStore type
type IAttackStore struct {
	mock.Mock
}

// Add provides a mock function with given fields: _a0
func (_m *IAttackStore) Add(_a0 models.AttackDetails) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.AttackDetails) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: _a0
func (_m *IAttackStore) Delete(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *IAttackStore) GetAll() []models.AttackDetails {
	ret := _m.Called()

	var r0 []models.AttackDetails
	if rf, ok := ret.Get(0).(func() []models.AttackDetails); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.AttackDetails)
		}
	}

	return r0
}

// GetByID provides a mock function with given fields: _a0
func (_m *IAttackStore) GetByID(_a0 string) (models.AttackDetails, error) {
	ret := _m.Called(_a0)

	var r0 models.AttackDetails
	if rf, ok := ret.Get(0).(func(string) models.AttackDetails); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.AttackDetails)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *IAttackStore) Update(_a0 string, _a1 models.AttackDetails) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, models.AttackDetails) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}