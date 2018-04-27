package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
	"github.com/techmexdev/lineuplist"
)

// NewArtistStorage returns a ArtistStorage postgres implementation.
func NewArtistStorage(dsn string) *ArtistStorage {
	return &ArtistStorage{sqlx.MustConnect("postgres", dsn)}
}

// ArtistStorage implements lineuplist.ArtistStorage.
type ArtistStorage struct {
	*sqlx.DB
}

// LoadAll returns all stored artists.
func (db *ArtistStorage) LoadAll() ([]lineuplist.Artist, error) {
	var aa []lineuplist.Artist

	q := "SELECT * FROM artist"

	err := db.Select(&aa, q)
	if err != nil {
		return []lineuplist.Artist{}, err
	}

	return aa, nil
}

// Load returns the first stored festival with that name.
func (db *ArtistStorage) Load(name string) (lineuplist.Artist, error) {
	var a lineuplist.Artist

	err := db.Get(&a, "SELECT * FROM artist WHERE name = $1", name)
	if err != nil {
		return lineuplist.Artist{}, err
	}

	fpStore := FestivalPreviewStorage{db.DB}
	a.Festivals, err = fpStore.FromArtist(name)
	if err != nil {
		return lineuplist.Artist{}, err
	}

	return a, nil
}

// Save inserts the festival in the databdbe if not exists, or returns it
// if it does.
func (db *ArtistStorage) Save(a lineuplist.Artist) (lineuplist.Artist, error) {
	id, err := uuid.NewV4()
	a.ID = id.String()
	if err != nil {
		return lineuplist.Artist{}, err
	}

	_, err = db.Exec("INSERT INTO artist(id, name) VALUES($1, $2) ON CONFLICT DO NOTHING", a.ID, a.Name)
	if err != nil {
		storedA, err := db.Load(a.Name)
		if err != nil {
			return lineuplist.Artist{}, err
		}
		return storedA, nil
	}

	return a, nil
}

func (db *ArtistStorage) FromFestival(fest string) ([]lineuplist.Artist, error) {
	var aa []lineuplist.Artist

	q := `SELECT * FROM artist WHERE id IN (
  	SELECT artist_id FROM festival_artist WHERE festival_id=(
  		SELECT id FROM festival WHERE name=$1))`

	err := db.Select(&aa, q, fest)
	if err != nil {
		return []lineuplist.Artist{}, err
	}

	return aa, nil
}
