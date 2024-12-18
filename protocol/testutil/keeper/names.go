package keeper

import (
	"testing"

	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/gogoproto/proto"

	indexerevents "github.com/StreamFinance-Protocol/stream-chain/protocol/indexer/events"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/indexer/indexer_manager"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/mocks"

	storetypes "cosmossdk.io/store/types"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/constants"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/names/keeper"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/names/types"
	priceskeeper "github.com/StreamFinance-Protocol/stream-chain/protocol/x/prices/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// CreateJotaName creates Jota in the names module for tests.
func CreateJotaName(ctx sdk.Context, namesKeeper *keeper.Keeper) error {
	_, err := namesKeeper.CreateName(
		ctx,
		constants.NameJota.Id,
		constants.NameJota.Name,
	)
	return err
}

func CreateNonJotaName(ctx sdk.Context, namesKeeper *keeper.Keeper) error {
	_, err := namesKeeper.CreateName(
		ctx,
		constants.NameSolal.Id,
		constants.NameSolal.Name,
	)
	return err
}

func NamesKeepers(
	t testing.TB,
	msgSenderEnabled bool,
) (
	ctx sdk.Context,
	keeper *keeper.Keeper,
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
		keeper, storeKey = createNamesKeeper(stateStore, db, cdc, pricesKeeper, transientStoreKey, msgSenderEnabled)

		return []GenesisInitializer{pricesKeeper, keeper}
	})
	// Mock time provider response for market creation.
	mockTimeProvider.On("Now").Return(constants.TimeT)
	return ctx, keeper, pricesKeeper, accountKeeper, bankKeeper, storeKey
}

func createNamesKeeper(
	stateStore storetypes.CommitMultiStore,
	db *dbm.MemDB,
	cdc *codec.ProtoCodec,
	pk *priceskeeper.Keeper,
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
		pk,
		mockIndexerEventsManager,
	)

	return k, storeKey
}

// GetNameCreateEventsFromIndexerBlock returns the name create events in the
// Indexer Block event Kafka message.
func GetNameCreateEventsFromIndexerBlock(
	ctx sdk.Context,
	keeper *keeper.Keeper,
) []*indexerevents.NameCreateEventV1 {
	var nameEvents []*indexerevents.NameCreateEventV1
	block := keeper.GetIndexerEventManager().ProduceBlock(ctx)
	if block == nil {
		return nameEvents
	}
	for _, event := range block.Events {
		if event.Subtype != indexerevents.SubtypeName {
			continue
		}
		var nameEvent indexerevents.NameCreateEventV1
		err := proto.Unmarshal(event.DataBytes, &nameEvent)
		if err != nil {
			panic(err)
		}
		nameEvents = append(nameEvents, &nameEvent)
	}
	return nameEvents
}
