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
	log.Println("migrating database...")
	migrateDB()

	log.Println("scraping festivals...")
	fests, err := scraper.Festivals()
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

}

func migrateDB() {
	db, err := sql.Open("postgres", os.Getenv("PG_DSN"))
	if err != nil {
		log.Fatal(err)
	}

	driver, err := migpg.WithInstance(db, &migpg.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	m.Up()
}
