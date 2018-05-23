package lineuplist

import "time"

// Festival is a music festival.
type Festival struct {
	ID        string          `json:"-"`
	Name      string          `json:"name"`
	Lineup    []ArtistPreview `json:"lineup"`
	ImgSrc    string          `json:"imgSrc" db:"img_src"`
	StartDate time.Time       `json:"startDate" db:"start_date"`
	EndDate   time.Time       `json:"endDate" db:"end_date"`
	Country   string          `json:"country"`
	State     string          `json:"state"`
	City      string          `json:"city"`
}

// FestivalStorage is an interface
// for saving, and loading a festival.
type FestivalStorage interface {
	Save(Festival) (Festival, error)
	LoadAll() ([]Festival, error)
	Load(string) (Festival, error)
	FromArtist(string) ([]Festival, error)
}

// FestivalPreview is a minimal representation
// of a Festival.
type FestivalPreview struct {
	ID     string `json:"-"`
	Name   string `json:"name"`
	ImgSrc string `json:"imgSrc"`
}

// FestivalStorage is an interface
// for saving, and loading a festival.
type FestivalPreviewStorage interface {
	Save(FestivalPreview) (FestivalPreview, error)
	LoadAll() ([]FestivalPreview, error)
	Load(string) (FestivalPreview, error)
	FromArtist(string) ([]FestivalPreview, error)
}

// Artist is a musician or band.
type Artist struct {
	ID        string            `json:"-"`
	Name      string            `json:"name"`
	Festivals []FestivalPreview `json:"festivals"`
}

// ArtistStorage is an interface
// for saving, and loading an artist.
type ArtistStorage interface {
	Save(Artist) (Artist, error)
	LoadAll() ([]Artist, error)
	Load(string) (Artist, error)
	FromFestival(string) ([]Artist, error)
}

// ArtistPreview is a minimal representation
// of an Artist.
type ArtistPreview struct {
	ID   string `json:"-"`
	Name string `json:"name"`
}

// ArtistPreviewStorage is an interface
// for saving, and loading an artist preview.
type ArtistPreviewStorage interface {
	Save(ArtistPreview) (ArtistPreview, error)
	LoadAll() ([]ArtistPreview, error)
	Load(string) (ArtistPreview, error)
	FromFestival(string) ([]ArtistPreview, error)
}
