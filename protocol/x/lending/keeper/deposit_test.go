package keeper_test

// import (
// 	"testing"

// 	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/keeper"
// 	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/gogo/protobuf/codec"
// 	"github.com/stretchr/testify/require"
// )

// func TestDeposit(t *testing.T) {
// 	// Set up the keeper and context
// 	cdc := codec.New()
// 	storeKey := sdk.NewKVStoreKey(types.StoreKey)
// 	k := keeper.NewKeeper(cdc, storeKey)
// 	ctx := sdk.NewContext(nil, sdk.Header{}, false, nil)

// 	// Define a depositor and amount
// 	depositor := sdk.AccAddress("address1")
// 	amount := sdk.NewCoin("token", sdk.NewInt(1000))

// 	// Call the Deposit function
// 	err := k.Deposit(ctx, depositor, amount)
// 	require.NoError(t, err)

// 	// Retrieve and check the stored balance
// 	balanceBytes, found := k.getDeposits(ctx, depositor.String())
// 	require.True(t, found)

// 	var balance sdk.Coins
// 	k.cdc.MustUnmarshal(balanceBytes, &balance)
// 	require.Equal(t, sdk.NewCoins(amount), balance)
// }

func getValue() int {
	return 1
}
