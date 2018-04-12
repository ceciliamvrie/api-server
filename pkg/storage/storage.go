package storage

import "github.com/techmexdev/lineuplist/pkg/model"

type Storage interface {
	InsertFests([]model.Festival) error
}
