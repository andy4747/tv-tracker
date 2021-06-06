package store

import (
	"database/sql"

	"github.com/angeldhakal/tv-tracker/models"
	"github.com/angeldhakal/tv-tracker/util"
)

type CreateTokenParams struct {
	CreatedAt string `json:"created_at"`
	Token     string `json:"token"`
	UserID    int64  `json:"user_id"`
}

type UpdateTokenParams struct {
	UpdatedAt sql.NullString `json:"updated_at"`
	Token     string         `json:"token"`
	UserID    int64          `json:"user_id"`
	ID        int64          `json:"id"`
}

type TokenStorer interface {
	GetToken(int64) (models.Tokens, error)
	GetTokenByToken(string) (models.Tokens, error)
	GetTokenByUser(int64) (models.Tokens, error)
	ListTokens() ([]models.Tokens, error)
	CreateToken(CreateTokenParams) (models.Tokens, error)
	UpdateToken(UpdateTokenParams) (models.Tokens, error)
	DeleteToken(int64) error
}

func (db *Store) GetToken(userID int64) (models.Tokens, error) {
	getToken := `SELECT id, created_at, updated_at, token, user_id
	FROM tokens
	WHERE id = $1
	`
	row := db.conn.QueryRow(getToken, userID)
	var token models.Tokens
	err := row.Scan(
		&token.ID,
		&token.CreatedAt,
		&token.UpdatedAt,
		&token.Token,
		&token.UserID,
	)
	return token, err
}

func (db *Store) GetTokenByToken(token string) (models.Tokens, error) {
	getTokenByToken := `SELECT id, created_at, updated_at, token, user_id
	FROM tokens
	WHERE token = $1
	`
	row := db.conn.QueryRow(getTokenByToken, token)
	var retrievedToken models.Tokens
	err := row.Scan(
		&retrievedToken.ID,
		&retrievedToken.CreatedAt,
		&retrievedToken.UpdatedAt,
		&retrievedToken.Token,
		&retrievedToken.UserID,
	)
	return retrievedToken, err
}

func (db *Store) GetTokenByUser(userID int64) (models.Tokens, error) {
	getTokenByUser := `SELECT id, created_at, updated_at, token, user_id
	FROM tokens
	WHERE user_id = $1
	`
	row := db.conn.QueryRow(getTokenByUser, userID)
	var token models.Tokens
	err := row.Scan(
		&token.ID,
		&token.CreatedAt,
		&token.UpdatedAt,
		&token.Token,
		&token.UserID,
	)
	return token, err
}

func (db *Store) ListTokens() ([]models.Tokens, error) {
	listTokens := `SELECT id, created_at, updated_at, token, user_id
	FROM tokens
	`
	rows, err := db.conn.Query(listTokens)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Tokens
	for rows.Next() {
		var token models.Tokens
		if err := rows.Scan(
			&token.ID,
			&token.CreatedAt,
			&token.UpdatedAt,
			&token.Token,
			&token.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, token)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (db *Store) CreateToken(tokenParams CreateTokenParams) (models.Tokens, error) {
	createToken := `INSERT INTO tokens (created_at, token, user_id)
	VALUES ($1, $2, $3)
	RETURNING id, created_at, updated_at, token, user_id
	`
	if tokenParams.CreatedAt == "" {
		tokenParams.CreatedAt = util.GetCurrentDate()
	}
	row := db.conn.QueryRow(createToken,
		tokenParams.CreatedAt,
		tokenParams.Token,
		tokenParams.UserID,
	)
	var token models.Tokens
	err := row.Scan(
		&token.ID,
		&token.CreatedAt,
		&token.UpdatedAt,
		&token.Token,
		&token.UserID,
	)
	return token, err
}

func (db *Store) UpdateToken(tokenParams UpdateTokenParams) (models.Tokens, error) {
	updateToken := `UPDATE tokens
	SET updated_at = $1, token = $2, user_id = $3
	WHERE id = $4
	RETURNING id, created_at, updated_at, token, user_id
	`
	if !tokenParams.UpdatedAt.Valid {
		tokenParams.UpdatedAt = sql.NullString{
			String: util.GetCurrentDate(),
			Valid:  true,
		}
	}
	row := db.conn.QueryRow(updateToken,
		tokenParams.UpdatedAt,
		tokenParams.Token,
		tokenParams.UserID,
		tokenParams.ID,
	)
	var token models.Tokens
	err := row.Scan(
		&token.ID,
		&token.CreatedAt,
		&token.UpdatedAt,
		&token.Token,
		&token.UserID,
	)
	return token, err
}

func (db *Store) DeleteToken(userID int64) error {
	deleteToken := `DELETE FROM tokens
	WHERE id = $1
	`
	_, err := db.conn.Exec(deleteToken, userID)
	return err
}
