package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"github.com/techmexdev/lineuplist/pkg/model"
)

// Postgres is an implementation of Storage
type Postgres struct {
	*sqlx.DB
}

// New returns a pointer to a pg connection
func New(dsn string) *Postgres {
	return &Postgres{sqlx.MustConnect("postgres", dsn)}
}

func (db *Postgres) InsertFests(fests []model.Festival) error {
	q := "INSERT INTO festival(id, name) VALUES ($1, $2)"

	stmt, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range fests {
		id, err := uuid.NewV4()
		if err != nil {
			return err
		}

		_, err = stmt.Exec(id, f.Name)
		if err != nil {
			log.Printf("could not insert festival - %#v: %s", f, err.Error())
		}
	}
	return nil
}
