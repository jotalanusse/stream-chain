package keeper

// import (
// 	"fmt"

// 	"github.com/StreamFinance-Protocol/stream-chain/protocol/x/lending/types"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// )

// // DepositType defines the type of deposit
// type DepositType string

// const (
// 	// BTCDeposit represents a deposit type of BTC
// 	BTCDeposit DepositType = "BTC"
// 	// Future deposit types can be added here
// )

// // Deposit represents a deposit into a lending account
// type Deposit struct {
// 	Amount sdk.Coin
// 	Type   DepositType
// }

// // ValidateDeposit checks if the deposit is valid based on the wallet and deposit type
// func (k Keeper) ValidateDeposit(ctx sdk.Context, deposit Deposit) error {
// 	// For now, only BTC deposits are supported
// 	if deposit.Type != BTCDeposit {
// 		return fmt.Errorf("unsupported deposit type: %s", deposit.Type)
// 	}

// 	// Additional validation logic can be added here

// 	return nil
// }

// // ProcessDeposit handles depositing money into a user's account
// func (k Keeper) ProcessDeposit(ctx sdk.Context, depositorAddress string, deposit Deposit) error {
// 	// Validate the deposit
// 	if err := k.ValidateDeposit(ctx, deposit); err != nil {
// 		return err
// 	}

// 	// Retrieve the lending account
// 	account, exists := k.GetLendAccount(ctx, depositorAddress)
// 	if !exists {
// 		return fmt.Errorf("account does not exist for address: %s", depositorAddress)
// 	}

// 	// Update the account balance
// 	updatedBalance, err := k.updateAccountBalance(account, deposit)
// 	if err != nil {
// 		return err
// 	}
// 	account.Balance = updatedBalance

// 	// Store the updated account
// 	k.SetLendingAccount(ctx, account)

// 	// Log the deposit
// 	fmt.Printf("Processed deposit: %s %s\n", deposit.Amount.String(), deposit.Type)

// 	return nil
// }

// // updateAccountBalance updates the balance of the lending account after a deposit
// func (k Keeper) updateAccountBalance(account types.LendingAccount, deposit Deposit) ([]*sdk.Coin, error) {
// 	// This function should add the deposit amount to the account's balance
// 	// Assuming sdk.Coin has a method to safely add amounts of the same denom
// 	for i, coin := range account.Balance {
// 		if coin.Denom == deposit.Amount.Denom {
// 			updatedAmount := coin.Add(deposit.Amount)
// 			account.Balance[i] = &updatedAmount
// 			return account.Balance, nil
// 		}
// 	}
// 	// If the deposit's denom wasn't found in the balance, append it
// 	account.Balance = append(account.Balance, &deposit.Amount)

// 	return account.Balance, nil
// }

// Simple function that returns 1
func getOne() int {
	return 1
}
