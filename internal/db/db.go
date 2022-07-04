package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB(dataSourceName string) (*Entities, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	errPing := db.Ping()
	if errPing != nil {
		return nil, fmt.Errorf("error connecting to database: %w ", errPing)
	}
	return &Entities{
		AccountDB: &AccountDB{DB: db},
	}, nil

}

type Entities struct {
	*AccountDB
}
