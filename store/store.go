package store

import (
	"database/sql"

	"github.com/angeldhakal/tv-tracker/models"
)

type Store struct {
	conn *sql.DB
}

func NewStore() Storer {
	return &Store{
		conn: models.Connect(),
	}
}

type Storer interface {
	UserStorer
	TokenStorer
}
