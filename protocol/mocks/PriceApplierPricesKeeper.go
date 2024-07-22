// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	pricestypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/prices/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// PriceApplierPricesKeeper is an autogenerated mock type for the PriceApplierPricesKeeper type
type PriceApplierPricesKeeper struct {
	mock.Mock
}

// GetAllMarketParams provides a mock function with given fields: ctx
func (_m *PriceApplierPricesKeeper) GetAllMarketParams(ctx types.Context) []pricestypes.MarketParam {
	ret := _m.Called(ctx)

	var r0 []pricestypes.MarketParam
	if rf, ok := ret.Get(0).(func(types.Context) []pricestypes.MarketParam); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pricestypes.MarketParam)
		}
	}

	return r0
}

// GetMarketParam provides a mock function with given fields: ctx, id
func (_m *PriceApplierPricesKeeper) GetMarketParam(ctx types.Context, id uint32) (pricestypes.MarketParam, bool) {
	ret := _m.Called(ctx, id)

	var r0 pricestypes.MarketParam
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, uint32) (pricestypes.MarketParam, bool)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint32) pricestypes.MarketParam); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(pricestypes.MarketParam)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint32) bool); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// PerformStatefulPriceUpdateValidation provides a mock function with given fields: ctx, marketPriceUpdates, performNonDeterministicValidation
func (_m *PriceApplierPricesKeeper) PerformStatefulPriceUpdateValidation(ctx types.Context, marketPriceUpdates *pricestypes.MarketPriceUpdates, performNonDeterministicValidation bool) error {
	ret := _m.Called(ctx, marketPriceUpdates, performNonDeterministicValidation)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *pricestypes.MarketPriceUpdates, bool) error); ok {
		r0 = rf(ctx, marketPriceUpdates, performNonDeterministicValidation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMarketPrice provides a mock function with given fields: ctx, update
func (_m *PriceApplierPricesKeeper) UpdateMarketPrice(ctx types.Context, update *pricestypes.MarketPriceUpdates_MarketPriceUpdate) error {
	ret := _m.Called(ctx, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *pricestypes.MarketPriceUpdates_MarketPriceUpdate) error); ok {
		r0 = rf(ctx, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPriceApplierPricesKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewPriceApplierPricesKeeper creates a new instance of PriceApplierPricesKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPriceApplierPricesKeeper(t mockConstructorTestingTNewPriceApplierPricesKeeper) *PriceApplierPricesKeeper {
	mock := &PriceApplierPricesKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}