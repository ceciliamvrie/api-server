package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
	"github.com/techmexdev/lineuplist"
)

// NewFestivalStorage returns a FestivalStorage postgres implementation.
func NewFestivalStorage(dsn string) *FestivalStorage {
	return &FestivalStorage{sqlx.MustConnect("postgres", dsn)}
}

// FestivalStorage implements lineuplist.FestivalStorage
type FestivalStorage struct {
	*sqlx.DB
}

// LoadAll returns all stored festivals
func (db *FestivalStorage) LoadAll() ([]lineuplist.Festival, error) {
	var ff []lineuplist.Festival

	err := db.Select(&ff, "SELECT * FROM festival")
	if err != nil {
		return []lineuplist.Festival{}, err
	}

	aStore := ArtistStorage{db.DB}
	for i := range ff {
		aa, err := aStore.FromFestival(ff[i].Name)
		if err != nil {
			return []lineuplist.Festival{}, err
		}
		ff[i].Lineup = aa
	}

	return ff, nil
}

// Load returns all stored festivals
// with that name.
func (db *FestivalStorage) Load(name string) (lineuplist.Festival, error) {
	var f lineuplist.Festival

	err := db.Get(&f, "SELECT id, name FROM festival WHERE name = $1", name)
	if err != nil {
		return lineuplist.Festival{}, err
	}

	aStore := &ArtistStorage{db.DB}
	aa, err := aStore.FromFestival(f.Name)
	if err != nil {
		return lineuplist.Festival{}, err
	}

	f.Lineup = aa

	return f, nil
}

// Save inserts the festival in the database if it doesn't exist,
// and retrieves it if it does.
func (db *FestivalStorage) Save(f lineuplist.Festival) (lineuplist.Festival, error) {
	q := "INSERT INTO festival(id, name, start_date, end_date, country, state, city)" +
		"VALUES($1, $2, $3, $4, $5, $6, $7)"

	id, err := uuid.NewV4()
	f.ID = id.String()
	if err != nil {
		return lineuplist.Festival{}, err
	}

	_, err = db.Exec(q, f.ID, f.Name, f.StartDate, f.EndDate, f.Country, f.State, f.City)
	if err != nil {
		return lineuplist.Festival{}, err
	}

	aStore := ArtistStorage{db.DB}

	for _, a := range f.Lineup {
		var storedA lineuplist.Artist

		storedA, err = aStore.Load(a.Name)
		if err != nil {
			storedA, err = aStore.Save(a)
			if err != nil {
				return lineuplist.Festival{}, err
			}
		}

		festArtID, err := uuid.NewV4()
		if err != nil {
			return lineuplist.Festival{}, err
		}

		_, err = db.Exec(`INSERT INTO festival_artist(id, festival_id, artist_id)
			VALUES($1, $2, $3)`, festArtID.String(), f.ID, storedA.ID)
		if err != nil {
			return lineuplist.Festival{}, err
		}
	}

	return f, nil
}

func (db *FestivalStorage) FromArtist(artist string) ([]lineuplist.Festival, error) {
	var ff []lineuplist.Festival

	q := `SELECT name, start_date, end_date, country, state, city
		FROM festival WHERE id IN (
  		SELECT festival_id FROM festival_artist WHERE artist_id=(
  			SELECT id FROM artist WHERE name=$1))`
	err := db.Select(&ff, q, artist)
	if err != nil {
		return []lineuplist.Festival{}, err
	}

	return ff, nil
}
