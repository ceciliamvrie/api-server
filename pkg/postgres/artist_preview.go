package postgres

import (
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"github.com/techmexdev/lineuplist"
)

func NewArtistPreviewStorage(dsn string) *ArtistPreviewStorage {
	return &ArtistPreviewStorage{sqlx.MustConnect("postgres", dsn)}
}

// ArtistPreviewStorage implements lineuplist.ArtistPreviewStorage
type ArtistPreviewStorage struct {
	*sqlx.DB
}

func (db *ArtistPreviewStorage) Save(ap lineuplist.ArtistPreview) (lineuplist.ArtistPreview, error) {
	id, err := uuid.NewV4()
	ap.ID = id.String()
	if err != nil {
		return lineuplist.ArtistPreview{}, err
	}

	_, err = db.Exec("INSERT INTO artist(id, name) VALUES($1, $2)", ap.ID, ap.Name)
	if err != nil {
		return lineuplist.ArtistPreview{}, err
	}

	return ap, nil
}

func (db *ArtistPreviewStorage) LoadAll() ([]lineuplist.ArtistPreview, error) {
	aStore := ArtistStorage{db.DB}

	aa, err := aStore.LoadAll()
	if err != nil {
		return []lineuplist.ArtistPreview{}, err
	}

	var aps []lineuplist.ArtistPreview
	for _, a := range aa {
		aps = append(aps, lineuplist.ArtistPreview{
			ID:   a.ID,
			Name: a.Name,
		})
	}

	return aps, nil
}

func (db *ArtistPreviewStorage) Load(name string) (lineuplist.ArtistPreview, error) {
	aStore := ArtistStorage{db.DB}

	a, err := aStore.Load(name)
	if err != nil {
		return lineuplist.ArtistPreview{}, err
	}

	return lineuplist.ArtistPreview{
		ID:   a.ID,
		Name: a.Name,
	}, nil
}

func (db *ArtistPreviewStorage) FromFestival(fest string) ([]lineuplist.ArtistPreview, error) {
	aStore := ArtistStorage{db.DB}

	aa, err := aStore.FromFestival(fest)
	if err != nil {
		return []lineuplist.ArtistPreview{}, err
	}

	var aps []lineuplist.ArtistPreview
	for _, a := range aa {
		aps = append(aps, lineuplist.ArtistPreview{
			ID:   a.ID,
			Name: a.Name,
		})
	}

	return aps, nil
}
