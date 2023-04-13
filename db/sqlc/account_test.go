package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	f := faker.New()
	arg := CreateAccountParams{
		OwnerName: f.Person().FirstName(),
		Currency: "IDR",
		Balance: f.Int64Between(10000, 1000000000),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.OwnerName, account.OwnerName)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.OwnerName, account2.OwnerName)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

	require.NotZero(t, account2.ID)
	require.NotZero(t, account2.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	f := faker.New()
	
	arg := UpdateAccountParams{
		ID: account1.ID,
		Balance: f.Int64Between(10_000, 1_000_000_000),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, arg.ID, updatedAccount.ID)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, account1.OwnerName, updatedAccount.OwnerName)
	require.Equal(t, account1.Currency, updatedAccount.Currency)
	require.WithinDuration(t, account1.CreatedAt, updatedAccount.CreatedAt, time.Second)

	require.NotZero(t, updatedAccount.ID)
	require.NotZero(t, updatedAccount.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	acc := createRandomAccount(t)

	_, err := testQueries.DeleteAccount(context.Background(), acc.ID)
	require.NoError(t, err)

	deletedAcc, err := testQueries.GetAccount(context.Background(), acc.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedAcc)
	require.Zero(t, deletedAcc.ID)
}