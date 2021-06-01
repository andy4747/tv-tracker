package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/angeldhakal/tv-tracker/models"
	"github.com/angeldhakal/tv-tracker/util"
)

type TokenRepository interface {
	GetToken(int64) (models.Tokens, error)
	GetTokenBytoken(int64) (models.Tokens, error)
	ListTokens() ([]models.Tokens, error)
	CreateToken(models.Tokens) (models.Tokens, error)
	UpdateToken(models.Tokens) (models.Tokens, error)
	DeleteToken(int64) error
}

type tokenRepo struct {
	conn *models.Queries
}

func NewTokenRepository() TokenRepository {
	return &tokenRepo{
		conn: models.Connect(),
	}
}

func (db *tokenRepo) GetToken(userID int64) (models.Tokens, error) {
	token, err := db.conn.GetToken(context.Background(), userID)
	if err != nil {
		log.Printf("couldn't get the token")
		return models.Tokens{}, err
	}
	return token, nil
}

func (db *tokenRepo) GetTokenBytoken(userID int64) (models.Tokens, error) {
	token, err := db.conn.GetTokenByUser(context.Background(), userID)
	if err != nil {
		log.Printf("couldn't get token for token")
		return models.Tokens{}, err
	}
	return token, nil
}

func (db *tokenRepo) ListTokens() ([]models.Tokens, error) {
	tokens, err := db.conn.ListTokens(context.Background())
	if err != nil {
		log.Printf("couldn't get tokens")
		return []models.Tokens{}, err
	}
	return tokens, nil
}

func (db *tokenRepo) CreateToken(token models.Tokens) (models.Tokens, error) {
	params := models.CreateTokenParams{
		CreatedAt: util.GetCurrentDate(),
		Token:     token.Token,
		UserID:    token.UserID,
	}
	createdToken, err := db.conn.CreateToken(context.Background(), params)
	if err != nil {
		log.Println("couldn't create token")
		return models.Tokens{}, err
	}
	return createdToken, nil
}

func (db *tokenRepo) UpdateToken(token models.Tokens) (models.Tokens, error) {
	params := models.UpdateTokenParams{
		UpdatedAt: sql.NullString{
			String: util.GetCurrentDate(),
			Valid:  true,
		},
		Token:  token.Token,
		UserID: token.UserID,
	}
	updateToken, err := db.conn.UpdateToken(context.Background(), params)
	if err != nil {
		log.Println("couldn't create token")
		return models.Tokens{}, err
	}
	return updateToken, nil
}

func (db *tokenRepo) DeleteToken(userID int64) error {
	err := db.conn.DeleteToken(context.Background(), userID)
	if err != nil {
		log.Println("userRepo DeleteUser: ", err.Error())
		return err
	}
	return nil
}
