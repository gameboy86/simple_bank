package db

import (
	"context"
	"database/sql"
	"simple_bank/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
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
    account_create := createRandomAccount(t)
    account_get, err := testQueries.GetAccount(context.Background(), account_create.ID)
    require.NoError(t, err)
    require.NotEmpty(t, account_get)

    require.Equal(t, account_create.ID, account_get.ID)
    require.Equal(t, account_create.Owner, account_get.Owner)
    require.Equal(t, account_create.Balance, account_get.Balance)
    require.Equal(t, account_create.Currency, account_get.Currency)
    require.WithinDuration(t, account_create.CreatedAt, account_get.CreatedAt, time.Second)
}

func TestUpdateAccou(t *testing.T) {
    account_create := createRandomAccount(t)
    arg := UpdateAccountParams{
        ID: account_create.ID,
        Balance: utils.RandomMoney(),
    }
    account_updated, err := testQueries.UpdateAccount(context.Background(), arg)
    require.NoError(t, err)
    require.NotEmpty(t, account_updated)
    require.Equal(t, account_create.ID, account_updated.ID)
    require.Equal(t, account_create.Owner, account_updated.Owner)
    require.Equal(t, arg.Balance, account_updated.Balance)
    require.Equal(t, account_create.Currency, account_updated.Currency)
    require.Equal(t, account_create.CreatedAt, account_updated.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
    account_create := createRandomAccount(t)
    err := testQueries.DeleteAccount(context.Background(), account_create.ID)
    require.NoError(t, err)
    account_get, err := testQueries.GetAccount(context.Background(), account_create.ID)
    require.Error(t, err)
    require.EqualError(t, err, sql.ErrNoRows.Error())
    require.Empty(t, account_get)
}

func TestListAccounts(t *testing.T) {
    for i:=0;i<10;i++ {
        createRandomAccount(t)
    }
    arg := ListAccountsParams{
        Limit: 5,
        Offset: 5,
    }
    accounts, err := testQueries.ListAccounts(context.Background(), arg)
    require.NoError(t, err)
    require.Len(t, accounts, 5)
    for _, account := range accounts {
        require.NotEmpty(t, account)
    }
}
