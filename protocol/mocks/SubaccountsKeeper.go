// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	big "math/big"
	rand "math/rand"

	mock "github.com/stretchr/testify/mock"

	subaccountstypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/subaccounts/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// SubaccountsKeeper is an autogenerated mock type for the SubaccountsKeeper type
type SubaccountsKeeper struct {
	mock.Mock
}

// CanUpdateSubaccounts provides a mock function with given fields: ctx, updates, updateType
func (_m *SubaccountsKeeper) CanUpdateSubaccounts(ctx types.Context, updates []subaccountstypes.Update, updateType subaccountstypes.UpdateType) (bool, []subaccountstypes.UpdateResult, error) {
	ret := _m.Called(ctx, updates, updateType)

	if len(ret) == 0 {
		panic("no return value specified for CanUpdateSubaccounts")
	}

	var r0 bool
	var r1 []subaccountstypes.UpdateResult
	var r2 error
	if rf, ok := ret.Get(0).(func(types.Context, []subaccountstypes.Update, subaccountstypes.UpdateType) (bool, []subaccountstypes.UpdateResult, error)); ok {
		return rf(ctx, updates, updateType)
	}
	if rf, ok := ret.Get(0).(func(types.Context, []subaccountstypes.Update, subaccountstypes.UpdateType) bool); ok {
		r0 = rf(ctx, updates, updateType)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, []subaccountstypes.Update, subaccountstypes.UpdateType) []subaccountstypes.UpdateResult); ok {
		r1 = rf(ctx, updates, updateType)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]subaccountstypes.UpdateResult)
		}
	}

	if rf, ok := ret.Get(2).(func(types.Context, []subaccountstypes.Update, subaccountstypes.UpdateType) error); ok {
		r2 = rf(ctx, updates, updateType)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DepositFundsFromAccountToSubaccount provides a mock function with given fields: ctx, fromAccount, toSubaccountId, assetId, amount
func (_m *SubaccountsKeeper) DepositFundsFromAccountToSubaccount(ctx types.Context, fromAccount types.AccAddress, toSubaccountId subaccountstypes.SubaccountId, assetId uint32, amount *big.Int) error {
	ret := _m.Called(ctx, fromAccount, toSubaccountId, assetId, amount)

	if len(ret) == 0 {
		panic("no return value specified for DepositFundsFromAccountToSubaccount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, types.AccAddress, subaccountstypes.SubaccountId, uint32, *big.Int) error); ok {
		r0 = rf(ctx, fromAccount, toSubaccountId, assetId, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllSubaccount provides a mock function with given fields: ctx
func (_m *SubaccountsKeeper) GetAllSubaccount(ctx types.Context) []subaccountstypes.Subaccount {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllSubaccount")
	}

	var r0 []subaccountstypes.Subaccount
	if rf, ok := ret.Get(0).(func(types.Context) []subaccountstypes.Subaccount); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]subaccountstypes.Subaccount)
		}
	}

	return r0
}

// GetNegativeTncSubaccountSeenAtBlock provides a mock function with given fields: ctx, perpetualId
func (_m *SubaccountsKeeper) GetNegativeTncSubaccountSeenAtBlock(ctx types.Context, perpetualId uint32) (uint32, bool, error) {
	ret := _m.Called(ctx, perpetualId)

	if len(ret) == 0 {
		panic("no return value specified for GetNegativeTncSubaccountSeenAtBlock")
	}

	var r0 uint32
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(types.Context, uint32) (uint32, bool, error)); ok {
		return rf(ctx, perpetualId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint32) uint32); ok {
		r0 = rf(ctx, perpetualId)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint32) bool); ok {
		r1 = rf(ctx, perpetualId)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(types.Context, uint32) error); ok {
		r2 = rf(ctx, perpetualId)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetNetCollateralAndMarginRequirements provides a mock function with given fields: ctx, update
func (_m *SubaccountsKeeper) GetNetCollateralAndMarginRequirements(ctx types.Context, update subaccountstypes.Update) (*big.Int, *big.Int, *big.Int, error) {
	ret := _m.Called(ctx, update)

	if len(ret) == 0 {
		panic("no return value specified for GetNetCollateralAndMarginRequirements")
	}

	var r0 *big.Int
	var r1 *big.Int
	var r2 *big.Int
	var r3 error
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.Update) (*big.Int, *big.Int, *big.Int, error)); ok {
		return rf(ctx, update)
	}
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.Update) *big.Int); ok {
		r0 = rf(ctx, update)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, subaccountstypes.Update) *big.Int); ok {
		r1 = rf(ctx, update)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	if rf, ok := ret.Get(2).(func(types.Context, subaccountstypes.Update) *big.Int); ok {
		r2 = rf(ctx, update)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*big.Int)
		}
	}

	if rf, ok := ret.Get(3).(func(types.Context, subaccountstypes.Update) error); ok {
		r3 = rf(ctx, update)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetRandomSubaccount provides a mock function with given fields: ctx, _a1
func (_m *SubaccountsKeeper) GetRandomSubaccount(ctx types.Context, _a1 *rand.Rand) (subaccountstypes.Subaccount, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetRandomSubaccount")
	}

	var r0 subaccountstypes.Subaccount
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, *rand.Rand) (subaccountstypes.Subaccount, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *rand.Rand) subaccountstypes.Subaccount); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(subaccountstypes.Subaccount)
	}

	if rf, ok := ret.Get(1).(func(types.Context, *rand.Rand) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubaccount provides a mock function with given fields: ctx, id
func (_m *SubaccountsKeeper) GetSubaccount(ctx types.Context, id subaccountstypes.SubaccountId) subaccountstypes.Subaccount {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetSubaccount")
	}

	var r0 subaccountstypes.Subaccount
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId) subaccountstypes.Subaccount); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(subaccountstypes.Subaccount)
	}

	return r0
}

// SetNegativeTncSubaccountSeenAtBlock provides a mock function with given fields: ctx, perpetualId, blockHeight
func (_m *SubaccountsKeeper) SetNegativeTncSubaccountSeenAtBlock(ctx types.Context, perpetualId uint32, blockHeight uint32) error {
	ret := _m.Called(ctx, perpetualId, blockHeight)

	if len(ret) == 0 {
		panic("no return value specified for SetNegativeTncSubaccountSeenAtBlock")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, uint32, uint32) error); ok {
		r0 = rf(ctx, perpetualId, blockHeight)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetSubaccount provides a mock function with given fields: ctx, subaccount
func (_m *SubaccountsKeeper) SetSubaccount(ctx types.Context, subaccount subaccountstypes.Subaccount) {
	_m.Called(ctx, subaccount)
}

// TransferFundsFromSubaccountToSubaccount provides a mock function with given fields: ctx, senderSubaccountId, recipientSubaccountId, assetId, quantums
func (_m *SubaccountsKeeper) TransferFundsFromSubaccountToSubaccount(ctx types.Context, senderSubaccountId subaccountstypes.SubaccountId, recipientSubaccountId subaccountstypes.SubaccountId, assetId uint32, quantums *big.Int) error {
	ret := _m.Called(ctx, senderSubaccountId, recipientSubaccountId, assetId, quantums)

	if len(ret) == 0 {
		panic("no return value specified for TransferFundsFromSubaccountToSubaccount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, subaccountstypes.SubaccountId, uint32, *big.Int) error); ok {
		r0 = rf(ctx, senderSubaccountId, recipientSubaccountId, assetId, quantums)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSubaccounts provides a mock function with given fields: ctx, updates, updateType
func (_m *SubaccountsKeeper) UpdateSubaccounts(ctx types.Context, updates []subaccountstypes.Update, updateType subaccountstypes.UpdateType) (bool, []subaccountstypes.UpdateResult, error) {
	ret := _m.Called(ctx, updates, updateType)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSubaccounts")
	}

	var r0 bool
	var r1 []subaccountstypes.UpdateResult
	var r2 error
	if rf, ok := ret.Get(0).(func(types.Context, []subaccountstypes.Update, subaccountstypes.UpdateType) (bool, []subaccountstypes.UpdateResult, error)); ok {
		return rf(ctx, updates, updateType)
	}
	if rf, ok := ret.Get(0).(func(types.Context, []subaccountstypes.Update, subaccountstypes.UpdateType) bool); ok {
		r0 = rf(ctx, updates, updateType)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, []subaccountstypes.Update, subaccountstypes.UpdateType) []subaccountstypes.UpdateResult); ok {
		r1 = rf(ctx, updates, updateType)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]subaccountstypes.UpdateResult)
		}
	}

	if rf, ok := ret.Get(2).(func(types.Context, []subaccountstypes.Update, subaccountstypes.UpdateType) error); ok {
		r2 = rf(ctx, updates, updateType)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// WithdrawFundsFromSubaccountToAccount provides a mock function with given fields: ctx, fromSubaccountId, toAccount, assetId, amount
func (_m *SubaccountsKeeper) WithdrawFundsFromSubaccountToAccount(ctx types.Context, fromSubaccountId subaccountstypes.SubaccountId, toAccount types.AccAddress, assetId uint32, amount *big.Int) error {
	ret := _m.Called(ctx, fromSubaccountId, toAccount, assetId, amount)

	if len(ret) == 0 {
		panic("no return value specified for WithdrawFundsFromSubaccountToAccount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, types.AccAddress, uint32, *big.Int) error); ok {
		r0 = rf(ctx, fromSubaccountId, toAccount, assetId, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSubaccountsKeeper creates a new instance of SubaccountsKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSubaccountsKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *SubaccountsKeeper {
	mock := &SubaccountsKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
