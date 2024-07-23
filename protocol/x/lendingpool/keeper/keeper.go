package keeper

import (
	"fmt"

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
	err := poolParams.Validate()
	if err != nil {
		return err
	}

	err = poolParams.ApplyDecimalConversions()
	if err != nil {
		return err
	}

	// Check if identifier already exists
	if _, found := k.GetPoolParams(ctx, poolParams.TokenDenom); found {
		return errorsmod.Wrapf(types.ErrPoolParamsAlreadyExists, "poolParams.TokenDenom already exists (%s)", poolParams.TokenDenom)
	}

	k.setPoolParams(ctx, poolParams)
	return nil
}

func (k Keeper) GetPoolParams(
	ctx sdk.Context,
	tokenDenom string,
) (val types.PoolParams, found bool) {
	store := k.getPoolParamsStore(ctx)

	b := store.Get([]byte(tokenDenom))

	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) setPoolParams(ctx sdk.Context, poolParams types.PoolParams) {
	store := k.getPoolParamsStore(ctx)
	b := k.cdc.MustMarshal(&poolParams)
	store.Set([]byte(poolParams.TokenDenom), b)
}

// GetAllPoolParams returns all poolParams
func (k Keeper) GetAllPoolParams(ctx sdk.Context) (list []types.PoolParams) {
	store := k.getPoolParamsStore(ctx)
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PoolParams
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
