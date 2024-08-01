// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	aggregator "github.com/StreamFinance-Protocol/stream-chain/protocol/app/ve/aggregator"

	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// VoteAggregator is an autogenerated mock type for the VoteAggregator type
type VoteAggregator struct {
	mock.Mock
}

// AggregateDaemonVEIntoFinalPrices provides a mock function with given fields: ctx, votes
func (_m *VoteAggregator) AggregateDaemonVEIntoFinalPrices(ctx types.Context, votes []aggregator.Vote) (map[string]*big.Int, error) {
	ret := _m.Called(ctx, votes)

	if len(ret) == 0 {
		panic("no return value specified for AggregateDaemonVEIntoFinalPrices")
	}

	var r0 map[string]*big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, []aggregator.Vote) (map[string]*big.Int, error)); ok {
		return rf(ctx, votes)
	}
	if rf, ok := ret.Get(0).(func(types.Context, []aggregator.Vote) map[string]*big.Int); ok {
		r0 = rf(ctx, votes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, []aggregator.Vote) error); ok {
		r1 = rf(ctx, votes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPriceForValidator provides a mock function with given fields: validator
func (_m *VoteAggregator) GetPriceForValidator(validator types.ConsAddress) map[string]*big.Int {
	ret := _m.Called(validator)

	if len(ret) == 0 {
		panic("no return value specified for GetPriceForValidator")
	}

	var r0 map[string]*big.Int
	if rf, ok := ret.Get(0).(func(types.ConsAddress) map[string]*big.Int); ok {
		r0 = rf(validator)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*big.Int)
		}
	}

	return r0
}

// NewVoteAggregator creates a new instance of VoteAggregator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVoteAggregator(t interface {
	mock.TestingT
	Cleanup(func())
}) *VoteAggregator {
	mock := &VoteAggregator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}