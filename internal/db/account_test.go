package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/wilorios/simple_bank/internal/entities"
	"github.com/wilorios/simple_bank/internal/util"
)

func createRandomAccount(t *testing.T) entities.Account {
	input := entities.Account{
		Owner:       util.RandomOwner(),
		Balance:     util.RandomBalance(),
		CryptoMoney: util.RandomCryptoMoney(),
	}

	account, err := testEntities.CreateAccount(input)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.NotZero(t, account.Id)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)

	account2, err := testEntities.GetAccount(account.Id)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account.Id, account2.Id)
	require.Equal(t, account.Owner, account2.Owner)
	require.Equal(t, account.Balance, account2.Balance)
	require.Equal(t, account.CryptoMoney, account2.CryptoMoney)
	require.WithinDuration(t, account.CreatedAt, account2.CreatedAt, time.Second)
}
