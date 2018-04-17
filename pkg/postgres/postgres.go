package postgres

import (
	"github.com/jmoiron/sqlx"
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

	err := fs.Select(&fests, "SELECT name FROM festival")
	if err != nil {
		return []lineuplist.Festival{{}}, err
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
func (fs *FestivalStorage) Save(fest lineuplist.Festival) error {
	_, err := fs.Exec("INSERT INTO festival(name) VALUES($1)", fest.Name)
	if err != nil {
		return err
	}
	return nil
}
