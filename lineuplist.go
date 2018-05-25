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
	LoadAll(category string) ([]Festival, error)
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

// FestivalPreviewStorage is an interface
// for saving, and loading a festival.
type FestivalPreviewStorage interface {
	Save(FestivalPreview) (FestivalPreview, error)
	LoadAll(category string) ([]FestivalPreview, error)
	Load(string) (FestivalPreview, error)
	FromArtist(string) ([]FestivalPreview, error)
}

// Artist is a musician or band.
type Artist struct {
	ID          string            `json:"-"`
	Name        string            `json:"name"`
	ImgSrc      string            `json:"imgSrc" db:"img_src"`
	ExternalURL string            `json:"externalURL" db:"externalUrl"`
	Popularity  int               `json:"popularity"`
	Followers   int               `json:"followers"`
	Genres      []string          `json:"genres"`
	TopTracks   []Track           `json:"topTracks" db:"top_tracks"`
	Albums      []Album           `json:"albums"`
	Related     []ArtistPreview   `json:"relatedArtists" db:"related_artist"`
	Festivals   []FestivalPreview `json:"festivals"`
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
	ID         string `json:"-"`
	Name       string `json:"name"`
	ImgSrc     string `json:"imgSrc"`
	Popularity int    `json:"popularity"`
}

// ArtistPreviewStorage is an interface
// for saving, and loading an artist preview.
type ArtistPreviewStorage interface {
	Save(ArtistPreview) (ArtistPreview, error)
	LoadAll() ([]ArtistPreview, error)
	Load(string) (ArtistPreview, error)
	FromFestival(string) ([]ArtistPreview, error)
}

// Track is an artist's song.
type Track struct {
	ID          string `json:"-"`
	Name        string `json:"name"`
	ExternalURL string `json:"externalUrl"`
	Album       `json:"album"`
}

// Album is an artist's music album
type Album struct {
	ID          string `json:"-"`
	Name        string `json:"name"`
	ImgSrc      string `json:"imgSrc"`
	ExternalURL string `json:"externalUrl"`
}
