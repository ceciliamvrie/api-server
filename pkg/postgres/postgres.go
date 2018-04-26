package postgres

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate"
	migpg "github.com/golang-migrate/migrate/database/postgres"
)

// MigrateUp applies all up migrations to a pg db.
func MigrateUp(dsn string) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := migpg.WithInstance(db, &migpg.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		log.Println("error applying up migrations: ", err)
	}

}

// MigrateDown applies all down migrations to a pg db.
func MigrateDown(dsn string) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := migpg.WithInstance(db, &migpg.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Down()
	if err != nil {
		log.Println("error applying down migrations: ", err)
	}

}
