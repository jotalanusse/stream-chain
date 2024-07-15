package keeper

// import (
//     sdk "github.com/cosmos/cosmos-sdk/types"
//     "github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
// )

// // Deposit handles the logic for depositing an amount
// func (k Keeper) Deposit(ctx sdk.Context, depositor sdk.AccAddress, amount sdk.Coin) error {
//     depositorKey := depositor.String()

//     // Retrieve existing balance
//     var existingBalance sdk.Coins
//     balanceBytes, found := k.getDeposits(ctx, depositorKey)
//     if found {
//         k.cdc.MustUnmarshal(balanceBytes, &existingBalance)
//     }

//     // Add the deposit amount to the existing balance
//     newBalance := existingBalance.Add(amount)

//     // Save the new balance
//     balanceBytes = k.cdc.MustMarshal(&newBalance)
//     k.setDeposits(ctx, depositorKey, balanceBytes)

//     return nil
// }

func someFunctionName() int {
	return 1
}
