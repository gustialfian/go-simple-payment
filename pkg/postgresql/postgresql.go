package postgresql

import (
	"database/sql"
	"log"
)

// New return new *sql.DB instance
func New(con string) (*sql.DB, error) {
	db, err := sql.Open("postgres", con)
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("psql new: %v", err.Error())
		log.Println("db down")
		return nil, err
	}
	log.Println("db up")

	return db, nil
}

// Migration make table and seed the data
func Migration(db *sql.DB) error {
	log.Println("Migrating...")
	query := `
	drop table if exists users;
	create table if not exists users (
		id integer,
		username varchar(50),
		password varchar(50),
		amount integer
	);`

	_, err := db.Query(query)
	if err != nil {
		return err
	}

	return nil
}

// Seed the dummy data
func Seed(db *sql.DB) error {
	log.Println("Seeding...")
	query := `
	insert into users (id, username, password, amount) 
	values (1, 'foo', 'secret', 100)
	, (2, 'bar', 'secret', 100)
	, (3, 'baz', 'secret', 100)`

	_, err := db.Query(query)
	if err != nil {
		return err
	}

	return nil
}
