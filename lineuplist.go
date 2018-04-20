package lineuplist

import "time"

// Festival is a music festival
type Festival struct {
	Name      string    `json:"name" db:"name"`
	StartDate time.Time `json:"startDate" db:"start_date"`
	EndDate   time.Time `json:"endDate" db:"end_date"`
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`
}

// FestivalStorage is an interface
// for saving, and loading a festival.
type FestivalStorage interface {
	Save(Festival) error
	LoadAll() ([]Festival, error)
	Load(string) ([]Festival, error)
}
