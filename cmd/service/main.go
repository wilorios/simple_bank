package main

import (
	"fmt"
	"log"

	"github.com/wilorios/simple_bank/internal/db"
)

func main() {
	dsn := "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"

	entities, err := db.NewDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	accounts, err := entities.ListAccounts(5, 0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("accounts ", accounts)
}
