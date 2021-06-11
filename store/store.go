package store

import (
	"database/sql"

	"github.com/angeldhakal/tv-tracker/models"
)

type Store struct {
	conn *sql.DB
}

func NewStore() Tracker {
	return &Store{
		conn: models.Connect(),
	}
}

type Tracker interface {
	UserTracker
	TokenTracker
	MovieTracker
}
