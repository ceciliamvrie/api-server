package lineuplist

type Festival struct {
	Name string
}

type FestivalStorage interface {
	Save(Festival) error
	LoadAll() ([]Festival, error)
	Load(string) ([]Festival, error)
}
