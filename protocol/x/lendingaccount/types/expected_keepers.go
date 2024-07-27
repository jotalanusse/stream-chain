package types

import (
	"math/big"
	"math/rand"

	assettypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/assets/types"
	lendingpooltypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingpool/types"
	satypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/subaccounts/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//nolint:staticcheck
type SubaccountsKeeper interface {
	GetAllSubaccount(ctx sdk.Context) (list []satypes.Subaccount)
	GetRandomSubaccount(ctx sdk.Context, rand *rand.Rand) (satypes.Subaccount, error)
	GetNetCollateralAndMarginRequirements(
		ctx sdk.Context,
		update satypes.Update,
	) (
		bigNetCollateral *big.Int,
		bigInitialMargin *big.Int,
		bigMaintenanceMargin *big.Int,
		err error,
	)
	CanUpdateSubaccounts(
		ctx sdk.Context,
		updates []satypes.Update,
		updateType satypes.UpdateType,
	) (
		success bool,
		successPerUpdate []satypes.UpdateResult,
		err error,
	)
	UpdateSubaccounts(
		ctx sdk.Context,
		updates []satypes.Update,
		updateType satypes.UpdateType,
	) (
		success bool,
		successPerUpdate []satypes.UpdateResult,
		err error,
	)
	DepositFundsFromAccountToSubaccount(
		ctx sdk.Context,
		fromAccount sdk.AccAddress,
		toSubaccountId satypes.SubaccountId,
		assetId uint32,
		amount *big.Int,
	) (err error)
	WithdrawFundsFromSubaccountToAccount(
		ctx sdk.Context,
		fromSubaccountId satypes.SubaccountId,
		toAccount sdk.AccAddress,
		assetId uint32,
		amount *big.Int,
	) (err error)
	TransferFundsFromSubaccountToSubaccount(
		ctx sdk.Context,
		senderSubaccountId satypes.SubaccountId,
		recipientSubaccountId satypes.SubaccountId,
		assetId uint32,
		quantums *big.Int,
	) (err error)
	SetSubaccount(ctx sdk.Context, subaccount satypes.Subaccount)
	GetSubaccount(
		ctx sdk.Context,
		id satypes.SubaccountId,
	) (val satypes.Subaccount)
}

type LendingPoolKeeper interface {
	GetPoolParams(
		ctx sdk.Context,
		tokenDenom string,
	) (val lendingpooltypes.InternalPoolParams, found bool)
	CalculateCumulativeBorrowIndex(ctx sdk.Context, tokenDenom string) (newBorrowIndex *big.Int, err error)
	LendToCreditAccount(ctx sdk.Context, amount *big.Int, creditAccount string, tokenDenom string) error
}

type AssetsKeeper interface {
	GetAsset(ctx sdk.Context, id uint32) (val assettypes.Asset, exists bool)
}

type PricesKeeper interface {
	GetUsdValue(
		ctx sdk.Context,
		assetId uint32,
		bigQuantums *big.Int,
	) (usdValue *big.Int, err error)
}
