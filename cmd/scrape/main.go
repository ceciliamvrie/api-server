package main

import (
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/techmexdev/lineuplist/pkg/postgres"
	"github.com/techmexdev/lineuplist/pkg/scraper"
)

func main() {
	log.Println("migrating database...")
	postgres.MigrateUp("file://migrations", os.Getenv("PG_DSN"))

	log.Println("scraping festivals...")
	fests, err := scraper.Festivals()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("scraped %v festivals\n\n", len(fests))

	fStore := postgres.NewFestivalStorage(os.Getenv("PG_DSN"))

	for _, f := range fests {
		_, err = fStore.Save(f)
		if err != nil {
			log.Printf("error saving %s: %s", f.Name, err)
		}
	}

}
