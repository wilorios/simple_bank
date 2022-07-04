package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/wilorios/simple_bank/internal/entities"
)

type AccountDB struct {
	*sqlx.DB
}

func (a *AccountDB) GetAccount(id int) (entities.Account, error) {
	var ea entities.Account
	err := a.Get(&ea, `SELECT * FROM accounts WHERE id = $1`, id)
	if err != nil {
		return entities.Account{}, fmt.Errorf("error getting account: %w", err)
	}
	return ea, nil
}

func (a *AccountDB) GetListAccounts() ([]entities.Account, error) {
	var ea []entities.Account
	err := a.Select(&ea, `SELECT * FROM accounts`)
	if err != nil {
		return []entities.Account{}, fmt.Errorf("error getting accounts: %w", err)
	}
	return ea, nil
}

func (a *AccountDB) CreateAccount(ea *entities.Account) error {
	err := a.Get(a, `INSERT INTO accounts VALUES ($1, $2, $3, $4) RETURNING *`, ea.Id, ea.Owner, ea.Balance, ea.CryptoMoney)
	if err != nil {
		return fmt.Errorf("error inserting account: %w", err)
	}
	return nil
}

func (a *AccountDB) UpdateAccount(ea *entities.Account) error {
	err := a.Get(a, `UPDATE INTO accounts SET balance = $1 WHERE id=$2 RETURNING *`, ea.Balance, ea.Id)
	if err != nil {
		return fmt.Errorf("error updating account: %w", err)
	}
	return nil
}

func (a *AccountDB) DeleteAccount(id int) error {
	_, err := a.Exec(`DELETE FROM accounts WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting account: %w", err)
	}
	return nil
}