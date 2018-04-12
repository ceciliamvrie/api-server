package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	migpg "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/techmexdev/lineuplist/pkg/scraper"
	"github.com/techmexdev/lineuplist/pkg/storage/postgres"
)

func main() {
	err := migrateDB()
	if err != nil {
		log.Fatal(err)
	}

	fests, err := scraper.GetFestivals()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("scraped %v festivals", len(fests))

	db := postgres.New(os.Getenv("PG_DSN"))
	err = db.InsertFests(fests)
	if err != nil {
		log.Fatal(err)
	}
}

func migrateDB() error {
	db, err := sql.Open("postgres", os.Getenv("PG_DSN"))
	if err != nil {
		return err
	}

	driver, err := migpg.WithInstance(db, &migpg.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

	m.Up()
	return nil
}
