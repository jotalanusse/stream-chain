package keeper

import (
	"encoding/json"
	"fmt"
	"math/big"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingpool/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		bankKeeper types.BankKeeper
		storeKey   storetypes.StoreKey

		// the addresses capable of executing MsgSetPoolParams message.
		authorities map[string]struct{}
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	bankKeeper types.BankKeeper,
	storeKey storetypes.StoreKey,

) *Keeper {
	return &Keeper{
		cdc:        cdc,
		bankKeeper: bankKeeper,
		storeKey:   storeKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With(log.ModuleKey, fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) InitializeForGenesis(ctx sdk.Context) {
}

func (k Keeper) getPoolParamsStore(
	ctx sdk.Context,
) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PoolParamsKeyPrefix))
}

// CreateEpochInfo creates a new EpochInfo.
// Return an error if the epoch fails validation, if the epoch Id already exists.
func (k Keeper) CreatePoolParams(ctx sdk.Context, poolParams types.PoolParams) error {

	// Perform stateless validation on the provided `EpochInfo`.
	internalParams, err := poolParams.Validate()
	if err != nil {
		return err
	}

	err = internalParams.ApplyDecimalConversions()
	if err != nil {
		return err
	}

	// Check if identifier already exists
	if _, found := k.GetPoolParams(ctx, internalParams.TokenDenom); found {
		return errorsmod.Wrapf(types.ErrPoolParamsAlreadyExists, "poolParams.TokenDenom already exists (%s)", poolParams.TokenDenom)
	}

	err = k.setPoolParams(ctx, internalParams)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) GetLastUpdatedTime(
	ctx sdk.Context,
	tokenDenom string,
) (timestamp *big.Int, found bool) {
	store := k.getLastUpdatedTimeStore(ctx)

	b := store.Get([]byte(tokenDenom))

	if b == nil {
		return nil, false
	}

	timestamp = new(big.Int).SetBytes(b)
	return timestamp, true
}

func (k Keeper) SetLastUpdatedTime(ctx sdk.Context, tokenDenom string, timestamp *big.Int) {
	store := k.getLastUpdatedTimeStore(ctx)

	bz := timestamp.Bytes()
	store.Set([]byte(tokenDenom), bz)
}

func (k Keeper) getLastUpdatedTimeStore(
	ctx sdk.Context,
) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.LastUpdatedTimePrefix))
}

func (k Keeper) GetCumulativeInterestRate(
	ctx sdk.Context,
	tokenDenom string,
) (rate *big.Int, found bool) {
	store := k.getCumulativeInterestRateStore(ctx)

	b := store.Get([]byte(tokenDenom))

	if b == nil {
		return nil, false
	}

	rate = new(big.Int).SetBytes(b)
	return rate, true
}

func (k Keeper) SetCumulativeInterestRate(ctx sdk.Context, tokenDenom string, rate *big.Int) {
	store := k.getCumulativeInterestRateStore(ctx)

	bz := rate.Bytes()
	store.Set([]byte(tokenDenom), bz)
}

func (k Keeper) getCumulativeInterestRateStore(
	ctx sdk.Context,
) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CumulativeInterestRatePrefix))
}

func (k Keeper) GetTotalBorrowed(ctx sdk.Context, tokenDenom string) (amount *big.Int, found bool) {
	store := k.getTotalBorrowedStore(ctx)

	b := store.Get([]byte(tokenDenom))

	if b == nil {
		return nil, false
	}

	amount = new(big.Int).SetBytes(b)
	return amount, true
}

func (k Keeper) SetTotalBorrowed(ctx sdk.Context, tokenDenom string, amount *big.Int) {
	store := k.getTotalBorrowedStore(ctx)

	bz := amount.Bytes()
	store.Set([]byte(tokenDenom), bz)
}

func (k Keeper) getTotalBorrowedStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.TotalBorrowedPrefix))
}

func (k Keeper) GetCurrentBorrowAPY(ctx sdk.Context, tokenDenom string) (rate *big.Int, found bool) {
	store := k.getCurrentBorrowAPYStore(ctx)

	b := store.Get([]byte(tokenDenom))

	if b == nil {
		return nil, false
	}

	rate = new(big.Int).SetBytes(b)
	return rate, true
}

func (k Keeper) SetCurrentBorrowAPY(ctx sdk.Context, tokenDenom string, rate *big.Int) {
	store := k.getCurrentBorrowAPYStore(ctx)

	bz := rate.Bytes()
	store.Set([]byte(tokenDenom), bz)
}

func (k Keeper) getCurrentBorrowAPYStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CurrentBorrowAPYEighteenDecimalsPrefix))
}

func (k Keeper) GetLastUpdatedTotalLiquidity(ctx sdk.Context, tokenDenom string) (amount *big.Int, found bool) {
	store := k.getLastUpdatedTotalLiquidityStore(ctx)

	b := store.Get([]byte(tokenDenom))

	if b == nil {
		return nil, false
	}

	amount = new(big.Int).SetBytes(b)
	return amount, true
}

func (k Keeper) SetLastUpdatedTotalLiquidity(ctx sdk.Context, tokenDenom string, amount *big.Int) {
	store := k.getLastUpdatedTotalLiquidityStore(ctx)

	bz := amount.Bytes()
	store.Set([]byte(tokenDenom), bz)
}

func (k Keeper) getLastUpdatedTotalLiquidityStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.LastUpdatedTotalLiquidityPrefix))
}

func (k Keeper) GetPoolParams(
	ctx sdk.Context,
	tokenDenom string,
) (val types.InternalPoolParams, found bool) {
	store := k.getPoolParamsStore(ctx)

	b := store.Get([]byte(tokenDenom))

	if b == nil {
		return types.InternalPoolParams{}, false
	}

	err := json.Unmarshal(b, &val)
	if err != nil {
		return val, false
	}

	return val, true
}

func (k Keeper) setPoolParams(ctx sdk.Context, poolParams types.InternalPoolParams) error {
	store := k.getPoolParamsStore(ctx)

	// Marshal the struct to JSON
	b, err := json.Marshal(poolParams)
	if err != nil {
		return err
	}

	store.Set([]byte(poolParams.TokenDenom), b)
	return nil
}

// GetAllPoolParams returns all poolParams
func (k Keeper) GetAllPoolParams(ctx sdk.Context) (list []types.InternalPoolParams) {
	store := k.getPoolParamsStore(ctx)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.InternalPoolParams
		err := json.Unmarshal(iterator.Value(), &val)
		if err != nil {
			continue
		}
		list = append(list, val)
	}

	return
}
