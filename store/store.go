package store

import (
	"database/sql"

	"github.com/angeldhakal/tv-tracker/models"
)

type Store struct {
	conn *sql.DB
}

func NewStore() Store {
	return Store{
		conn: models.Connect(),
	}
}

type Storer interface {
	GetUser(int64) (models.Users, error)
	GetUserByEmail(string) (models.Users, error)
	ListUsers() ([]models.Users, error)
	CreateUser(CreateUserParams) (models.Users, error)
	UpdateUser(UpdateUserParams) (models.Users, error)
	DeleteUser(int64) error
}
