package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate"
	migpg "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/lib/pq"
	"github.com/techmexdev/lineuplist/pkg/handler"
)

func main() {
	goEnv := os.Getenv("GO_ENV")
	dsn := os.Getenv("PG_DSN")

	var options handler.Options
	if goEnv == "PROD" {
		options = handler.Options{Log: false}
	} else {
		options = handler.Options{Log: true}
	}

	migrateDB()
	router := handler.New(dsn, options)

	if goEnv == "PROD" {
		log.Println("Starting server at port 80...")
		http.ListenAndServeTLS(":80", "server.crt", "server.key", router)
	} else {
		log.Println("Starting server at localhost:3000...")
		log.Fatal(http.ListenAndServe(":3000", router))
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
