package keeper

import (
	"fmt"
	"math/big"

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
		cdc               codec.BinaryCodec
		storeKey          storetypes.StoreKey
		lendingPoolKeeper types.LendingPoolKeeper
		subaccountsKeeper types.SubaccountsKeeper
		assetsKeeper      types.AssetsKeeper
		pricesKeeper      types.PricesKeeper

		// the addresses capable of executing MsgSetPoolParams message.
		authorities map[string]struct{}
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	lendingPoolKeeper types.LendingPoolKeeper,
	subaccountsKeeper types.SubaccountsKeeper,
	assetsKeeper types.AssetsKeeper,
	pricesKeeper types.PricesKeeper,
	authorities []string,
) *Keeper {
	return &Keeper{
		cdc:               cdc,
		storeKey:          storeKey,
		lendingPoolKeeper: lendingPoolKeeper,
		subaccountsKeeper: subaccountsKeeper,
		assetsKeeper:      assetsKeeper,
		pricesKeeper:      pricesKeeper,
		authorities:       lib.UniqueSliceToSet(authorities),
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

func (k Keeper) GetLendingAccount(ctx sdk.Context, managerName string, lendingAccountId uint32) (account types.LendingAccount, found bool) {
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

func (k Keeper) SetLendingAccount(ctx sdk.Context, managerName string, lendingAccountId uint32, account types.LendingAccount) error {
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

func (k Keeper) getLendingInterfaceStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.LendingInterfaceKeyPrefix))
}

func (k Keeper) GetLendingInterface(ctx sdk.Context, managerName string) (lendingInterface types.LendingInterface, found bool) {
	store := k.getLendingInterfaceStore(ctx)
	b := store.Get([]byte(managerName))

	if b == nil {
		return types.LendingInterface{}, false
	}

	err := k.cdc.Unmarshal(b, &lendingInterface)
	if err != nil {
		return types.LendingInterface{}, false
	}

	return lendingInterface, true
}

func (k Keeper) SetLendingInterface(ctx sdk.Context, lendingInterface types.LendingInterface) error {
	store := k.getLendingInterfaceStore(ctx)
	managerName := lendingInterface.ManagerName

	// Check if a LendingInterface with the same manager name already exists
	if _, found := k.GetLendingInterface(ctx, managerName); found {
		return fmt.Errorf("lending interface with manager name %s already exists", managerName)
	}

	b, err := k.cdc.Marshal(&lendingInterface)
	if err != nil {
		return err
	}

	store.Set([]byte(managerName), b)
	return nil
}

func (k Keeper) getLendingManagerStore(ctx sdk.Context) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.LendingManagerKeyPrefix))
}

func (k Keeper) GetLendingManager(ctx sdk.Context, managerName string) (manager types.LendingManager, found bool) {
	store := k.getLendingManagerStore(ctx)
	b := store.Get([]byte(managerName))

	if b == nil {
		return types.LendingManager{}, false
	}

	err := k.cdc.Unmarshal(b, &manager)
	if err != nil {
		return types.LendingManager{}, false
	}

	return manager, true
}

func (k Keeper) SetLendingManager(ctx sdk.Context, manager types.LendingManager) error {

	asset, found := k.assetsKeeper.GetAsset(ctx, manager.AssetId)
	if !found {
		return fmt.Errorf("asset with id %d not found", manager.AssetId)
	}

	_, found = k.lendingPoolKeeper.GetPoolParams(ctx, asset.Denom)
	if !found {
		return fmt.Errorf("pool with base denom %s not found", asset.Denom)
	}

	store := k.getLendingManagerStore(ctx)
	managerName := manager.Name

	// Check if a LendingManager with the same name already exists
	if _, found := k.GetLendingManager(ctx, managerName); found {
		return fmt.Errorf("lending manager with name %s already exists", managerName)
	}

	b, err := k.cdc.Marshal(&manager)
	if err != nil {
		return err
	}

	store.Set([]byte(managerName), b)
	return nil
}

func (k Keeper) CheckLendingAccountExistsAndAddToAddressMapping(ctx sdk.Context, managerName string, lendingAccountId uint32, address string) error {
	if lendingAccountId >= subaccounttypes.MaxSubaccountIdNumber {
		return fmt.Errorf("lendingAccountId out of range")
	}

	store := k.getLendingManagerStore(ctx)
	b := store.Get([]byte(managerName))

	if b == nil {
		return fmt.Errorf("lending manager with name %s not found", managerName)
	}

	var manager types.LendingManager
	err := k.cdc.Unmarshal(b, &manager)
	if err != nil {
		return err
	}

	if _, exists := manager.AccountNumberToAddress[lendingAccountId]; exists {
		return fmt.Errorf("account number %d is already in use", lendingAccountId)
	}

	// Update the account_number_to_address map
	manager.AccountNumberToAddress[lendingAccountId] = address

	// Marshal the updated manager back to bytes
	b, err = k.cdc.Marshal(&manager)
	if err != nil {
		return err
	}

	// Store the updated manager
	store.Set([]byte(managerName), b)
	return nil
}

func (k Keeper) GetLendingAccountAddress(ctx sdk.Context, managerName string, lendingAccountId uint32) (address string, found bool) {
	manager, found := k.GetLendingManager(ctx, managerName)
	if !found {
		return "", false
	}

	address, found = manager.AccountNumberToAddress[lendingAccountId]
	return address, found
}

func (k Keeper) RemoveFromLendingAccountToAddressMapping(ctx sdk.Context, managerName string, lendingAccountId uint32, address string) error {

	if lendingAccountId >= subaccounttypes.MaxSubaccountIdNumber {
		return fmt.Errorf("lendingAccountId out of range")
	}

	store := k.getLendingManagerStore(ctx)
	b := store.Get([]byte(managerName))

	if b == nil {
		return fmt.Errorf("lending manager with name %s not found", managerName)
	}

	var manager types.LendingManager
	err := k.cdc.Unmarshal(b, &manager)
	if err != nil {
		return err
	}

	if _, exists := manager.AccountNumberToAddress[lendingAccountId]; !exists {
		return fmt.Errorf("account number %d is not in use", lendingAccountId)
	}

	// Update the account_number_to_address map
	delete(manager.AccountNumberToAddress, lendingAccountId)

	// Marshal the updated manager back to bytes
	b, err = k.cdc.Marshal(&manager)
	if err != nil {
		return err
	}

	// Store the updated manager
	store.Set([]byte(managerName), b)
	return nil
}

func (k Keeper) CheckAddressNotInUseAndAddToSybilResistanceMapping(ctx sdk.Context, managerName string, address string) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.LendingManagerSybilResistanceKeyPrefix+":"+managerName))
	b := store.Get([]byte(address))

	if b != nil {
		return fmt.Errorf("address %s already exists in sybil resistance mapping", address)
	}

	store.Set([]byte(address), []byte{1})
	return nil
}

func (k Keeper) RemoveFromSybilResistanceMapping(ctx sdk.Context, managerName string, address string) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.LendingManagerSybilResistanceKeyPrefix+":"+managerName))
	b := store.Get([]byte(address))

	if b == nil {
		return fmt.Errorf("address %s not found in sybil resistance mapping", address)
	}

	store.Delete([]byte(address))
	return nil
}

func (k Keeper) SetTotalDebt(ctx sdk.Context, managerName string, totalDebt *big.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.TotalDebtKeyPrefix+":"+managerName))
	store.Set([]byte("total_debt"), totalDebt.Bytes())
}

func (k Keeper) GetTotalDebt(ctx sdk.Context, managerName string) (*big.Int, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.TotalDebtKeyPrefix+":"+managerName))
	b := store.Get([]byte("total_debt"))
	if b == nil {
		return nil, fmt.Errorf("total debt not found for manager %s", managerName)
	}
	totalDebt := new(big.Int).SetBytes(b)
	return totalDebt, nil
}

func (k Keeper) SetLastBorrowedInBlock(ctx sdk.Context, managerName string, borrowedInBlock *big.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BorrowedInBlockKeyPrefix+":"+managerName))
	store.Set([]byte("borrowed_in_block"), borrowedInBlock.Bytes())
}

func (k Keeper) GetLastBorrowedInBlock(ctx sdk.Context, managerName string) (*big.Int, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BorrowedInBlockKeyPrefix+":"+managerName))
	b := store.Get([]byte("borrowed_in_block"))
	if b == nil {
		return nil, fmt.Errorf("borrowed in block not found for manager %s", managerName)
	}
	borrowedInBlock := new(big.Int).SetBytes(b)
	return borrowedInBlock, nil
}

func (k Keeper) SetTotalBorrowedLastBlock(ctx sdk.Context, managerName string, blockLastBorrowed *big.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BlockLastBorrowedKeyPrefix+":"+managerName))
	store.Set([]byte("block_last_borrowed"), blockLastBorrowed.Bytes())
}

func (k Keeper) GetTotalBorrowedLastBlock(ctx sdk.Context, managerName string) (*big.Int, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.BlockLastBorrowedKeyPrefix+":"+managerName))
	b := store.Get([]byte("block_last_borrowed"))
	if b == nil {
		return nil, fmt.Errorf("block last borrowed not found for manager %s", managerName)
	}
	blockLastBorrowed := new(big.Int).SetBytes(b)
	return blockLastBorrowed, nil
}
