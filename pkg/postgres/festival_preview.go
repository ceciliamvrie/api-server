package postgres

import (
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	"github.com/techmexdev/lineuplist"
)

func NewFestivalPreviewStorage(dsn string) *FestivalPreviewStorage {
	return &FestivalPreviewStorage{sqlx.MustConnect("postgres", dsn)}
}

// FestivalPreviewStorage implements lineuplist.FestivalPreviewStorage
type FestivalPreviewStorage struct {
	*sqlx.DB
}

func (db *FestivalPreviewStorage) Save(fp lineuplist.FestivalPreview) (lineuplist.FestivalPreview, error) {
	id, err := uuid.NewV4()
	fp.ID = id.String()
	if err != nil {
		return lineuplist.FestivalPreview{}, err
	}

	_, err = db.Exec("INSERT INTO festival(id, name) VALUES($1, $2)", fp.ID, fp.Name)
	if err != nil {
		return lineuplist.FestivalPreview{}, err
	}

	return fp, nil
}

func (db *FestivalPreviewStorage) LoadAll() ([]lineuplist.FestivalPreview, error) {
	fStore := FestivalStorage{db.DB}

	ff, err := fStore.LoadAll()
	if err != nil {
		return []lineuplist.FestivalPreview{}, err
	}

	var fps []lineuplist.FestivalPreview
	for _, f := range ff {
		fps = append(fps, lineuplist.FestivalPreview{
			ID:     f.ID,
			Name:   f.Name,
			ImgSrc: f.ImgSrc,
		})
	}

	return fps, nil
}

func (db *FestivalPreviewStorage) Load(name string) (lineuplist.FestivalPreview, error) {
	fStore := FestivalStorage{db.DB}

	f, err := fStore.Load(name)
	if err != nil {
		return lineuplist.FestivalPreview{}, err
	}

	return lineuplist.FestivalPreview{
		ID:     f.ID,
		Name:   f.Name,
		ImgSrc: f.ImgSrc,
	}, nil
}

func (db *FestivalPreviewStorage) FromArtist(name string) ([]lineuplist.FestivalPreview, error) {
	fStore := FestivalStorage{db.DB}

	ff, err := fStore.FromArtist(name)
	if err != nil {
		return []lineuplist.FestivalPreview{}, err
	}

	var fps []lineuplist.FestivalPreview
	for _, f := range ff {
		fps = append(fps, lineuplist.FestivalPreview{
			ID:   f.ID,
			Name: f.Name,
		})
	}

	return fps, nil
}
