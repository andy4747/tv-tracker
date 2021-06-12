package store

import (
	"github.com/angeldhakal/tv-tracker/models"
)

func NewUserStore() UserTracker {
	return &userStore{
		conn: models.Connect(),
	}
}

func NewTokenStore() TokenTracker {
	return &tokenStore{
		conn: models.Connect(),
	}
}

func NewMovieStore() MovieTracker {
	return &movieStore{
		conn: models.Connect(),
	}
}
