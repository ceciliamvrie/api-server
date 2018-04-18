package lineuplist

// Festival is a music festival
type Festival struct {
	Name     string `json:"name"`
	Date     string `json:"date"`
	Location string `json:"location"`
}

// FestivalStorage is an interface
// for saving, and loading a festival.
type FestivalStorage interface {
	Save(Festival) error
	LoadAll() ([]Festival, error)
	Load(string) ([]Festival, error)
}
