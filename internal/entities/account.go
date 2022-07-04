package entities

import "time"

type Account struct {
	Id          int       `db:"id"`
	Owner       string    `db:"owner"`
	Balance     int       `db:"balance"`
	CryptoMoney string    `db:"crypto_money"`
	CreatedAt   time.Time `db:"created_at"`
}

type IAccount interface {
	GetAccount(id int) (Account, error)
	GetListAccounts() ([]Account, error)
	CreateAccount(a *Account) error
	UpdateAccount(a *Account) error
	DeleteAccount(id int) error
}
