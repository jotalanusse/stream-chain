package keeper

import (
	"fmt"

	"cosmossdk.io/log"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lendingaccount/types"
	subaccounttypes "github.com/StreamFinance-Protocol/stream-chain/protocol/x/subaccounts/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey

		// the addresses capable of executing MsgSetPoolParams message.
		authorities map[string]struct{}
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	authorities []string,
) *Keeper {
	return &Keeper{
		cdc:         cdc,
		storeKey:    storeKey,
		authorities: lib.UniqueSliceToSet(authorities),
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With(log.ModuleKey, fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) InitializeForGenesis(ctx sdk.Context) {
}

func (k Keeper) getLendingAccountStore(ctx sdk.Context, managerName string) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.LendingAccountKeyPrefix+":"+managerName))
}

func (k Keeper) GetLendingAccount(ctx sdk.Context, managerName string, lendingAccountId uint64) (account types.LendingAccount, found bool) {
	if lendingAccountId >= subaccounttypes.MaxSubaccountIdNumber {
		return types.LendingAccount{}, false
	}

	store := k.getLendingAccountStore(ctx, managerName)
	b := store.Get([]byte(fmt.Sprintf("%d", lendingAccountId)))

	if b == nil {
		return types.LendingAccount{}, false
	}

	err := k.cdc.Unmarshal(b, &account)
	if err != nil {
		return types.LendingAccount{}, false
	}

	return account, true
}

func (k Keeper) SetLendingAccount(ctx sdk.Context, managerName string, lendingAccountId uint64, account types.LendingAccount) error {
	if lendingAccountId >= subaccounttypes.MaxSubaccountIdNumber {
		return fmt.Errorf("lendingAccountId out of range")
	}

	store := k.getLendingAccountStore(ctx, managerName)
	b, err := k.cdc.Marshal(&account)
	if err != nil {
		return err
	}

	store.Set([]byte(fmt.Sprintf("%d", lendingAccountId)), b)
	return nil
}
