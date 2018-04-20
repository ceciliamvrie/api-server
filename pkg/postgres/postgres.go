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

// FestivalStorage implements lineuplist.FestivalStorage
type FestivalStorage struct {
	*sqlx.DB
}

// LoadAll returns all stored festivals
func (fs *FestivalStorage) LoadAll() ([]lineuplist.Festival, error) {
	var fests []lineuplist.Festival
	q := "SELECT name, start_date, end_date, country, state, city FROM festival"

	err := fs.Select(&fests, q)
	if err != nil {
		return []lineuplist.Festival{}, err
	}

	return fests, nil
}

// Load returns all stored festivals
// with that name.
func (fs *FestivalStorage) Load(name string) ([]lineuplist.Festival, error) {
	var fests []lineuplist.Festival
	q := "SELECT name FROM festival WHERE name =" + name

	err := fs.Select(&fests, q)
	if err != nil {
		return []lineuplist.Festival{{}}, err
	}

	return fests, nil
}

// Save inserts the festival in the database.
func (fs *FestivalStorage) Save(f lineuplist.Festival) error {
	q := "INSERT INTO festival(id, name, start_date, end_date, country, state, city)" +
		"VALUES($1, $2, $3, $4, $5, $6, $7)"

	id, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
	}

	_, err = fs.Exec(
		q, id, f.Name, f.StartDate, f.EndDate,
		f.Country, f.State, f.City,
	)
	if err != nil {
		return err
	}
	return nil
}
