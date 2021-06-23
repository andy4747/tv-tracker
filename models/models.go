package models

import (
	"database/sql"
)

type Status string

const (
	Completed Status = "c"
	Waiting   Status = "w"
	Plan      Status = "p"
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

type Movies struct {
	ID            int64          `json:"id"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     sql.NullString `json:"updated_at"`
	UserID        int64          `json:"user_id"`
	Name          string         `json:"name"`
	Status        Status         `json:"status"`
	CurrentLength int64          `json:"current_length"`
	Language      string         `json:"language"`
}

var EmptyObject = make(map[string]string)

var EmptyMovieList = []Movies{}
