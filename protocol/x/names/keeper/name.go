package keeper

import (
	"sort"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	indexerevents "github.com/StreamFinance-Protocol/stream-chain/protocol/indexer/events"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/indexer/indexer_manager"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/lib"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/names/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreateName(
	ctx sdk.Context,
	nameId uint32,
	nameName string,
) (types.Name, error) {
	if prevName, exists := k.GetName(ctx, nameId); exists {
		return types.Name{}, errorsmod.Wrapf(
			types.ErrNameIdAlreadyExists,
			"previous name = %v",
			prevName,
		)
	}

	if nameId == types.NameJota.Id {
		// Ensure nameId zero is always Jota. This is a protocol-wide invariant.
		if nameName != types.NameJota.Name {
			return types.Name{}, types.ErrJotaMustBeNameZero
		}
	}

	// Ensure Jota is not created with a non-zero nameId. This is a protocol-wide invariant.
	if nameId != types.NameJota.Id && nameName == types.NameJota.Name {
		return types.Name{}, types.ErrJotaMustBeNameZero
	}

	// Ensure the name is unique versus existing names.
	allNames := k.GetAllNames(ctx)
	for _, name := range allNames {
		if name.Name == nameName {
			return types.Name{}, errorsmod.Wrap(types.ErrNameNameAlreadyExists, nameName)
		}
	}

	// Create the name
	name := types.Name{
		Id:   nameId,
		Name: nameName,
	}

	// Store the new name
	k.setName(ctx, name)

	k.GetIndexerEventManager().AddTxnEvent(
		ctx,
		indexerevents.SubtypeName,
		indexerevents.NameEventVersion,
		indexer_manager.GetBytes(
			indexerevents.NewNameCreateEvent(
				nameId,
				nameName,
			),
		),
	)

	return name, nil
}

func (k Keeper) setName(
	ctx sdk.Context,
	name types.Name,
) {
	b := k.cdc.MustMarshal(&name)
	nameStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NameKeyPrefix))
	nameStore.Set(lib.Uint32ToKey(name.Id), b)
}

func (k Keeper) GetName(
	ctx sdk.Context,
	id uint32,
) (val types.Name, exists bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NameKeyPrefix))

	b := store.Get(lib.Uint32ToKey(id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllNames(
	ctx sdk.Context,
) (list []types.Name) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.NameKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Name
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Id < list[j].Id
	})

	return list
}
