// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	domain "falabella/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// BeerRepository is an autogenerated mock type for the BeerRepository type
type BeerRepository struct {
	mock.Mock
}

// CreateBeer provides a mock function with given fields: beer
func (_m *BeerRepository) CreateBeer(beer *domain.Beer) (int, error) {
	ret := _m.Called(beer)

	var r0 int
	if rf, ok := ret.Get(0).(func(*domain.Beer) int); ok {
		r0 = rf(beer)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Beer) error); ok {
		r1 = rf(beer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *BeerRepository) FindAll() ([]domain.Beer, error) {
	ret := _m.Called()

	var r0 []domain.Beer
	if rf, ok := ret.Get(0).(func() []domain.Beer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Beer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: id
func (_m *BeerRepository) FindByID(id int) (*domain.Beer, error) {
	ret := _m.Called(id)

	var r0 *domain.Beer
	if rf, ok := ret.Get(0).(func(int) *domain.Beer); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Beer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: name
func (_m *BeerRepository) FindByName(name string) (bool, error) {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
