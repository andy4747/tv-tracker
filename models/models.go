package models

import (
	"database/sql"
)

type Tokens struct {
	ID        int64          `json:"id"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
	Token     string         `json:"token"`
	UserID    int64          `json:"user_id"`
}

type Users struct {
	ID        int64          `json:"id"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at"`
	Email     string         `json:"email"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
}
