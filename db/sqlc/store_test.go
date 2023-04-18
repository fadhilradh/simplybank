package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	store := NewStore(testDB)

	account1 := CreateRandomAccount(t)
	time.Sleep(time.Second * 1)
	account2 := CreateRandomAccount(t)

	// run concurrent transfer TX
	n := 30
	amount := int64(10_000)

	errsChan := make(chan error)
	resultsChan := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errsChan <- err
			resultsChan <- result

		}()
	}

	for i := 0; i < n; i++ {
		err := <-errsChan
		require.NoError(t, err)

		result := <-resultsChan
		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.NotZero(t, toEntry.ID)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)
	}
}
