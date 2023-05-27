package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// tun a concurrent transfer transactions
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
		// Run concurrent go routines
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
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

		// TODO: Check accounts' balance
	}
}
