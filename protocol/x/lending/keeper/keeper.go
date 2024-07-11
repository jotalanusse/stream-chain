package keeper

import (
	"fmt"

	cosmoslog "cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc             codec.BinaryCodec
		storeKey        storetypes.StoreKey
		bankKeeper      types.BankKeeper
		blockTimeKeeper types.BlockTimeKeeper
		ics4Wrapper     types.ICS4Wrapper

		// the addresses capable of executing MsgSetLimitParams message.
		authorities map[string]struct{}
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	bankKeeper types.BankKeeper,
	blockTimeKeeper types.BlockTimeKeeper,
	ics4Wrapper types.ICS4Wrapper,
	authorities []string,
) *Keeper {
	return &Keeper{
		cdc:             cdc,
		storeKey:        storeKey,
		bankKeeper:      bankKeeper,
		blockTimeKeeper: blockTimeKeeper,
		ics4Wrapper:     ics4Wrapper,
		authorities:     lib.UniqueSliceToSet(authorities),
	}
}

// Takes the context, the address making the deposit, and the amount to deposit.
func (k Keeper) ProcessDeposit(ctx sdk.Context, depositorAddress string, amount sdk.Coin) error {
	// Use the bankKeeper to add the amount to the module's account
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, depositorAddress, types.ModuleName, sdk.NewCoins(amount))
	if err != nil {
		return err
	}
	// Log the deposit
	k.Logger(ctx).Info("Processed deposit", "depositor", depositorAddress, "amount", amount.String())

	return nil
}

func (k Keeper) HasAuthority(authority string) bool {
	_, ok := k.authorities[authority]
	return ok
}

func (k Keeper) Logger(ctx sdk.Context) cosmoslog.Logger {
	return ctx.Logger().With(cosmoslog.ModuleKey, fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) InitializeForGenesis(ctx sdk.Context) {}
