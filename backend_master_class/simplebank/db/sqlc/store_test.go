package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccountWithBalance(t, 200, "EUR")
	account2 := createRandomAccountWithBalance(t, 200, "EUR")

	fmt.Println(">> before:", account1.Balance, account2.Balance)
	// run a concurrent transfer transactions
	n := 5
	amount := int64(10)

	// Channel is designed to connect concurrent Go routines.
	// Allows to sync with each other without explicit locking.
	// Use `make` keyword to create the channel

	// One channel to receive the errors
	errs := make(chan error)
	// One other channel to receive the transfer results
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		// Log context values:
		// txName := fmt.Sprintf("tx %d", i+1)
		// Run concurrent go routines
		go func() {
			// Log context values:
			// ctx := context.WithValue(context.Background(), txKey, txName)
			ctx := context.Background()
			result, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			// Send errors to the `errs` channel
			errs <- err
			// Send transaction results to the `results` channel
			results <- result
		}()
	}

	// Check the results
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		// Receive errors from `errs` channel
		err := <-errs
		require.NoError(t, err)

		// Receive transaction results from `results` channel
		result := <-results
		require.NotEmpty(t, result)

		// Check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// Check entry from
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		// Check entry to
		fromTo := result.ToEntry
		require.NotEmpty(t, fromTo)
		require.Equal(t, account2.ID, fromTo.AccountID)
		require.Equal(t, amount, fromTo.Amount)
		require.NotZero(t, fromTo.ID)

		_, err = store.GetEntry(context.Background(), fromTo.ID)
		require.NoError(t, err)

		// Check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		// Check accounts' balance
		fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0) // amount, 2 * amount, 3 * amount, ..., n * amount

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// Check final updated balances
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)
}

func TestTransferTDeadlock(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccountWithBalance(t, 200, "EUR")
	account2 := createRandomAccountWithBalance(t, 200, "EUR")

	fmt.Println(">> before:", account1.Balance, account2.Balance)

	n := 10
	amount := int64(10)
	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		// Receive errors from `errs` channel
		err := <-errs
		require.NoError(t, err)
	}

	// Check final updated balances
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)
}
