package storage

import "github.com/techmexdev/lineuplist/pkg/model"

type Storage interface {
	GetFests() ([]model.Festival, error)
	InsertFests([]model.Festival) error
}
