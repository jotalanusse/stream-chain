package keeper

import (
	"testing"

	dbm "github.com/cosmos/cosmos-db"

	"github.com/StreamFinance-Protocol/stream-chain/protocol/indexer/indexer_manager"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/mocks"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/constants"

	storetypes "cosmossdk.io/store/types"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/keeper"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
	priceskeeper "github.com/StreamFinance-Protocol/stream-chain/protocol/x/prices/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// LendingKeepers initializes the context and keepers for lending module tests.
func LendingKeepers(
	t testing.TB,
	msgSenderEnabled bool,
) (
	ctx sdk.Context,
	lendingKeeper *keeper.Keeper,
	pricesKeeper *priceskeeper.Keeper,
	accountKeeper *authkeeper.AccountKeeper,
	bankKeeper *bankkeeper.BaseKeeper,
	storeKey storetypes.StoreKey,
) {
	var mockTimeProvider *mocks.TimeProvider
	ctx = initKeepers(t, func(
		db *dbm.MemDB,
		registry codectypes.InterfaceRegistry,
		cdc *codec.ProtoCodec,
		stateStore storetypes.CommitMultiStore,
		transientStoreKey storetypes.StoreKey,
	) []GenesisInitializer {
		// Define necessary keepers here for unit tests
		pricesKeeper, _, _, _, mockTimeProvider = createPricesKeeper(stateStore, db, cdc, transientStoreKey)
		accountKeeper, _ = createAccountKeeper(stateStore, db, cdc, registry)
		bankKeeper, _ = createBankKeeper(stateStore, db, cdc, accountKeeper)
		lendingKeeper, storeKey = createLendingKeeper(stateStore, db, cdc, transientStoreKey, msgSenderEnabled)

		return []GenesisInitializer{pricesKeeper, lendingKeeper}
	})

	// Mock time provider response for market creation.
	mockTimeProvider.On("Now").Return(constants.TimeT)

	return ctx, lendingKeeper, pricesKeeper, accountKeeper, bankKeeper, storeKey
}

func createLendingKeeper(
	stateStore storetypes.CommitMultiStore,
	db *dbm.MemDB,
	cdc *codec.ProtoCodec,
	transientStoreKey storetypes.StoreKey,
	msgSenderEnabled bool,
) (*keeper.Keeper, storetypes.StoreKey) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)

	mockMsgSender := &mocks.IndexerMessageSender{}
	mockMsgSender.On("Enabled").Return(msgSenderEnabled)
	mockIndexerEventsManager := indexer_manager.NewIndexerEventManager(mockMsgSender, transientStoreKey, true)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		mockIndexerEventsManager,
	)

	return k, storeKey
}

//Additional functions required for keeper implementation for lending

// // CreateUsdcLendingAccount creates a sample USDC lending account for testing purposes.
// func CreateUsdcLendingAccount(ctx sdk.Context, lendingKeeper *keeper.Keeper) error {
// 	_, err := lendingKeeper.CreateLendingAccount(
// 		ctx,
// 		constants.Usdc.Address,
// 		map[string]interface{}{
// 			"symbol": constants.Usdc.Symbol,
// 		},
// 	)
// 	return err
// }

// // GetLendingAccountCreateEventsFromIndexerBlock returns the lending account create events in the
// // Indexer Block event Kafka message.
// func GetLendingAccountCreateEventsFromIndexerBlock(
// 	ctx sdk.Context,
// 	keeper *keeper.Keeper,
// ) []*indexerevents.LendingAccountCreateEventV1 {
// 	var lendingAccountEvents []*indexerevents.LendingAccountCreateEventV1
// 	block := keeper.GetIndexerEventManager().ProduceBlock(ctx)
// 	if block == nil {
// 		return lendingAccountEvents
// 	}
// 	for _, event := range block.Events {
// 		if event.Subtype != indexerevents.SubtypeLending {
// 			continue
// 		}
// 		unmarshaler := common.UnmarshalerImpl{}
// 		var lendingAccountEvent indexerevents.LendingAccountCreateEventV1
// 		err := unmarshaler.Unmarshal(event.DataBytes, &lendingAccountEvent)
// 		if err != nil {
// 			panic(err)
// 		}
// 		lendingAccountEvents = append(lendingAccountEvents, &lendingAccountEvent)
// 	}
// 	return lendingAccountEvents
// }
