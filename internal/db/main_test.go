package db

import (
	"log"
	"os"
	"testing"
)

var testEntities *Entities

func TestMain(m *testing.M) {
	//TODO load config
	dsn := "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	var err error
	testEntities, err = NewDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}
