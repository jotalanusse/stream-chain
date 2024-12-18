package keeper_test

import (
	"fmt"
	"testing"

	errorsmod "cosmossdk.io/errors"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/constants"
	keepertest "github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/keeper"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/testutil/nullify"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/names/keeper"
	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/names/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	// firstValidNameId is the first valid name ID after the reserved `nameId=0` for Jota.
	firstValidNameId = uint32(1)
)

// createNNames creates n test names with id 1 to n (0 is reserved for Jota)
func createNNames(
	t *testing.T,
	ctx sdk.Context,
	keeper *keeper.Keeper,
	n int,
) ([]types.Name, error) {
	items := make([]types.Name, n)

	for i := range items {
		name, err := keeper.CreateName(
			ctx,
			uint32(i+1),               // NameId
			fmt.Sprintf("Name-%v", i), // Name
		)
		if err != nil {
			return items, err
		}

		items[i] = name
	}

	return items, nil
}

func TestCreateName_InvalidJotaName(t *testing.T) {
	ctx, keeper, _, _, _, _ := keepertest.NamesKeepers(t, true)

	// Throws error when creating a name with id 0 that's not Jota.
	_, err := keeper.CreateName(
		ctx,
		0,
		"Ema", // name
	)
	require.ErrorIs(t, err, types.ErrJotaMustBeNameZero)

	// Does not create an name.
	require.Len(t, keeper.GetAllNames(ctx), 0)

	// Throws error when creating name Jota with id other than 0.
	_, err = keeper.CreateName(
		ctx,
		1,
		constants.NameJota.Name, // name
	)
	require.ErrorIs(t, err, types.ErrJotaMustBeNameZero)

	// Does not create a name.
	require.Len(t, keeper.GetAllNames(ctx), 0)
}

func TestCreateName_NameAlreadyExists(t *testing.T) {
	ctx, keeper, _, _, _, _ := keepertest.NamesKeepers(t, true)

	_, err := keeper.CreateName(
		ctx,
		firstValidNameId,
		"Solal", // name
	)
	require.NoError(t, err)

	// Create a new name with identical name
	_, err = keeper.CreateName(
		ctx,
		2,
		"Solal", // name
	)
	require.EqualError(t, err, errorsmod.Wrap(types.ErrNameNameAlreadyExists, "Solal").Error())

	// Create a new name with the same ID
	_, err = keeper.CreateName(
		ctx,
		firstValidNameId,
		"Solal-COPY", // name
	)
	require.ErrorIs(t, err, types.ErrNameIdAlreadyExists)
}

func TestGetName_Success(t *testing.T) {
	ctx, keeper, _, _, _, _ := keepertest.NamesKeepers(t, true)
	items, err := createNNames(t, ctx, keeper, 10)
	require.NoError(t, err)

	for _, item := range items {
		rst, exists := keeper.GetName(ctx,
			item.Id,
		)
		require.True(t, exists)
		require.Equal(t,
			nullify.Fill(&item), //nolint:staticcheck
			nullify.Fill(&rst),  //nolint:staticcheck
		)
	}
}

func TestGetName_NotFound(t *testing.T) {
	ctx, keeper, _, _, _, _ := keepertest.NamesKeepers(t, true)
	_, exists := keeper.GetName(ctx,
		uint32(0),
	)
	require.False(t, exists)
}

func TestGetAllNames_Success(t *testing.T) {
	ctx, keeper, _, _, _, _ := keepertest.NamesKeepers(t, true)
	items, err := createNNames(t, ctx, keeper, 10)
	require.NoError(t, err)

	require.ElementsMatch(t,
		nullify.Fill(items),                   //nolint:staticcheck
		nullify.Fill(keeper.GetAllNames(ctx)), //nolint:staticcheck
	)
}
