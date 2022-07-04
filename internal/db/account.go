package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/wilorios/simple_bank/internal/entities"
)

type AccountDB struct {
	*sqlx.DB
}

func (a *AccountDB) GetAccount(id int64) (entities.Account, error) {
	var ea entities.Account
	err := a.Get(&ea, `SELECT * FROM accounts WHERE id = $1`, id)
	if err != nil {
		return entities.Account{}, err
	}
	return ea, nil
}

func (a *AccountDB) ListAccounts(limit int, offset int) ([]entities.Account, error) {
	var ea []entities.Account
	err := a.Select(&ea, `SELECT id, owner, balance, crypto_money, created_at FROM accounts ORDER BY id LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return []entities.Account{}, fmt.Errorf("error getting accounts: %w", err)
	}
	return ea, nil
}

func (a *AccountDB) CreateAccount(input entities.Account) (entities.Account, error) {
	var ea entities.Account
	err := a.Get(&ea, `INSERT INTO accounts (owner, balance, crypto_money) VALUES ($1, $2, $3) RETURNING id, owner, balance, crypto_money, created_at`, input.Owner, input.Balance, input.CryptoMoney)
	if err != nil {
		return entities.Account{}, fmt.Errorf("error inserting account: %w", err)
	}
	return ea, nil
}

func (a *AccountDB) UpdateAccount(input entities.Account) (entities.Account, error) {
	var ea entities.Account
	err := a.Get(&ea, `UPDATE accounts SET balance = $1 WHERE id=$2 RETURNING id, owner, balance, crypto_money, created_at`, input.Balance, input.Id)
	if err != nil {
		return entities.Account{}, fmt.Errorf("error updating account: %w", err)
	}
	return ea, nil
}

func (a *AccountDB) DeleteAccount(id int64) error {
	_, err := a.Exec(`DELETE FROM accounts WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting account: %w", err)
	}
	return nil
}
