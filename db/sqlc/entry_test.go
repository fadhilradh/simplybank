package db

import (
	"context"
	"testing"

	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	f := faker.New()
	account := CreateRandomAccount(t)
	testEntry := CreateEntryParams{
		AccountID: account.ID,
		Amount: f.Int64Between(10_000, 1_000_000_000),
	}

	entry, err := testQueries.CreateEntry(context.Background(), testEntry)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, account.ID, entry.AccountID)
	require.Equal(t, entry.Amount, testEntry.Amount)
	require.WithinDuration(t, account.CreatedAt, entry.CreatedAt, time.Second)

}