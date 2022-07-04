package db

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wilorios/simple_bank/internal/entities"
)

func TestCreateAccount(t *testing.T) {
	account := entities.Account{
		Owner:       "Wilor",
		Balance:     100,
		CryptoMoney: "BITCOIN",
	}

	accountReturned, err := testEntities.CreateAccount(&account)
	require.NoError(t, err)
	require.NotEmpty(t, accountReturned)

	require.NotZero(t, accountReturned.Id)
	require.NotZero(t, accountReturned.CreatedAt)
}
