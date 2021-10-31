package database

import (
	"database/sql"
	"log"

	"github.com/gustialfian/go-simple-payment/pkg/postgresql"
)

// New register database service
func New(con string) (*sql.DB, error) {
	db, err := postgresql.New(con)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	if err := postgresql.Migration(db); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	if err := postgresql.Seed(db); err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return db, nil
}
