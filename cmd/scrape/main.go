package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	migpg "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/techmexdev/lineuplist/pkg/postgres"
	"github.com/techmexdev/lineuplist/pkg/scraper"
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

	db := postgres.NewFestivalStorage(os.Getenv("PG_DSN"))

	for _, f := range fests {
		err = db.Save(f)
		if err != nil {
			log.Print(err)
		}
	}

	log.Println("retrieving stored festivals...")
	storedFests, err := db.LoadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range storedFests {
		log.Printf("   %v", f)
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
