package lineuplist

import "time"

// Festival is a music festival
type Festival struct {
	ID        string // format uuid v4
	Name      string `json:"name" db:"name"`
	Lineup    []Artist
	StartDate time.Time `json:"startDate" db:"start_date"`
	EndDate   time.Time `json:"endDate" db:"end_date"`
	Country   string    `json:"country"`
	State     string    `json:"state"`
	City      string    `json:"city"`
}

// FestivalStorage is an interface
// for saving, and loading a festival.
type FestivalStorage interface {
	Save(Festival) (Festival, error)
	LoadAll() ([]Festival, error)
	Load(string) (Festival, error)
	FromArtist(string) ([]Festival, error)
}

// Artist is a musician or band
type Artist struct {
	ID   string // format uuid v4
	Name string
}

// ArtistStorage is an interface
// for saving, and loading an artist.
type ArtistStorage interface {
	Save(Artist) (Artist, error)
	LoadAll() ([]Artist, error)
	Load(string) (Artist, error)
	FromFestival(string) ([]Artist, error)
}
