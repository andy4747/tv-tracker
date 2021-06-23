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

func GetUserByToken(token string) (models.Users, error) {
	userStore := NewUserStore()
	tokenStore := NewTokenStore()

	tokens, err := tokenStore.GetTokenByToken(token)
	if err != nil {
		return models.Users{}, err
	}
	user, err := userStore.GetUser(tokens.UserID)
	if err != nil {
		return models.Users{}, err
	}
	return user, nil
}
