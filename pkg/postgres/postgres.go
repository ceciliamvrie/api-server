package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
	"github.com/techmexdev/lineuplist"
)

// FestivalStorage is a postgres implementation
// of lineuplist.FestivalStorage

// NewFestivalStorage returns a FestivalStorage postgres implementation.
func NewFestivalStorage(dsn string) *FestivalStorage {
	return &FestivalStorage{sqlx.MustConnect("postgres", dsn)}
}

type FestivalStorage struct {
	*sqlx.DB
}

// LoadAll returns all stored festivals
func (fs *FestivalStorage) LoadAll() ([]lineuplist.Festival, error) {
	var fests []lineuplist.Festival

	err := fs.Select(&fests, "SELECT name, date, location FROM festival")
	if err != nil {
		return []lineuplist.Festival{}, err
	}

	return fests, nil
}

// Load returns all stored festivals
// with that name.
func (fs *FestivalStorage) Load(name string) ([]lineuplist.Festival, error) {
	var fests []lineuplist.Festival

	err := fs.Select(&fests, "SELECT name FROM festival WHERE name ="+name)
	if err != nil {
		return []lineuplist.Festival{{}}, err
	}

	return fests, nil
}

// Save inserts the festival in the database.
func (fs *FestivalStorage) Save(f lineuplist.Festival) error {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	_, err = fs.Exec("INSERT INTO festival(id, name, date, location)"+
		"VALUES($1, $2, $3, $4)", id, f.Name, f.Date, f.Location)
	if err != nil {
		return err
	}
	return nil
}
