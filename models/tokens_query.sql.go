// Code generated by sqlc. DO NOT EDIT.
// source: tokens_query.sql

package models

import (
	"context"
	"database/sql"
)

const createToken = `-- name: CreateToken :one
INSERT INTO tokens (created_at, token, user_id)
VALUES ($1, $2, $3)
RETURNING id, created_at, updated_at, token, user_id
`

type CreateTokenParams struct {
	CreatedAt string `json:"created_at"`
	Token     string `json:"token"`
	UserID    int64  `json:"user_id"`
}

func (q *Queries) CreateToken(ctx context.Context, arg CreateTokenParams) (Tokens, error) {
	row := q.db.QueryRowContext(ctx, createToken, arg.CreatedAt, arg.Token, arg.UserID)
	var i Tokens
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Token,
		&i.UserID,
	)
	return i, err
}

const deleteToken = `-- name: DeleteToken :exec
DELETE FROM tokens
WHERE id = $1
`

func (q *Queries) DeleteToken(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteToken, id)
	return err
}

const getToken = `-- name: GetToken :one
SELECT id, created_at, updated_at, token, user_id
FROM tokens
WHERE id = $1
`

func (q *Queries) GetToken(ctx context.Context, id int64) (Tokens, error) {
	row := q.db.QueryRowContext(ctx, getToken, id)
	var i Tokens
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Token,
		&i.UserID,
	)
	return i, err
}

const getTokenByToken = `-- name: GetTokenByToken :one
SELECT id, created_at, updated_at, token, user_id
FROM tokens
WHERE token = $1
`

func (q *Queries) GetTokenByToken(ctx context.Context, token string) (Tokens, error) {
	row := q.db.QueryRowContext(ctx, getTokenByToken, token)
	var i Tokens
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Token,
		&i.UserID,
	)
	return i, err
}

const getTokenByUser = `-- name: GetTokenByUser :one
SELECT id, created_at, updated_at, token, user_id
FROM tokens
WHERE user_id = $1
`

func (q *Queries) GetTokenByUser(ctx context.Context, userID int64) (Tokens, error) {
	row := q.db.QueryRowContext(ctx, getTokenByUser, userID)
	var i Tokens
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Token,
		&i.UserID,
	)
	return i, err
}

const listTokens = `-- name: ListTokens :many
SELECT id, created_at, updated_at, token, user_id
FROM tokens
`

func (q *Queries) ListTokens(ctx context.Context) ([]Tokens, error) {
	rows, err := q.db.QueryContext(ctx, listTokens)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tokens
	for rows.Next() {
		var i Tokens
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Token,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateToken = `-- name: UpdateToken :one
UPDATE tokens
SET updated_at = $1, token = $2, user_id = $3
WHERE id = $4
RETURNING id, created_at, updated_at, token, user_id
`

type UpdateTokenParams struct {
	UpdatedAt sql.NullString `json:"updated_at"`
	Token     string         `json:"token"`
	UserID    int64          `json:"user_id"`
	ID        int64          `json:"id"`
}

func (q *Queries) UpdateToken(ctx context.Context, arg UpdateTokenParams) (Tokens, error) {
	row := q.db.QueryRowContext(ctx, updateToken,
		arg.UpdatedAt,
		arg.Token,
		arg.UserID,
		arg.ID,
	)
	var i Tokens
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Token,
		&i.UserID,
	)
	return i, err
}