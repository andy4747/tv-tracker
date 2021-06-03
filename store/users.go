package store

import (
	"database/sql"

	"github.com/angeldhakal/tv-tracker/models"
	"github.com/angeldhakal/tv-tracker/util"
)

type CreateUserParams struct {
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type UpdateUserParams struct {
	UpdatedAt sql.NullString `json:"updated_at"`
	Email     string         `json:"email"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	ID        int64          `json:"id"`
}

type UserStorer interface {
	GetUser(int64) (models.Users, error)
	GetUserByEmail(string) (models.Users, error)
	ListUsers() ([]models.Users, error)
	CreateUser(CreateUserParams) (models.Users, error)
	UpdateUser(UpdateUserParams) (models.Users, error)
	DeleteUser(int64) error
}

func (db *Store) GetUser(id int64) (models.Users, error) {
	getUserQuery := `SELECT id, created_at, updated_at, email, username, password
	FROM users
	WHERE id = $1`

	row := db.conn.QueryRow(getUserQuery, id)
	var user models.Users
	err := row.Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Email,
		&user.Username,
		&user.Password,
	)
	return user, err
}

func (db *Store) GetUserByEmail(email string) (models.Users, error) {
	getUserByEmailQuery := `SELECT id, created_at, updated_at, email, username, password
	FROM users
	WHERE email = $1`
	row := db.conn.QueryRow(getUserByEmailQuery, email)
	var user models.Users
	err := row.Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Email,
		&user.Username,
		&user.Password,
	)
	return user, err
}

func (db *Store) ListUsers() ([]models.Users, error) {
	listUsersQuery := `SELECT id, created_at, updated_at, email, username, password
	FROM users;`
	rows, err := db.conn.Query(listUsersQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Users
	for rows.Next() {
		var user models.Users
		if err := rows.Scan(
			&user.ID,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Email,
			&user.Username,
			&user.Password,
		); err != nil {
			return nil, err
		}
		items = append(items, user)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (db *Store) CreateUser(userParams CreateUserParams) (models.Users, error) {
	createUserQuery := `INSERT INTO users (created_at, email, username, password)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, updated_at, email, username, password
	`
	if userParams.CreatedAt == "" {
		userParams.CreatedAt = util.GetCurrentDate()
	}
	row := db.conn.QueryRow(createUserQuery,
		userParams.CreatedAt,
		userParams.Email,
		userParams.Username,
		userParams.Password,
	)
	var user models.Users
	err := row.Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Email,
		&user.Username,
		&user.Password,
	)
	return user, err
}

func (db *Store) UpdateUser(userParams UpdateUserParams) (models.Users, error) {
	updateUserQuery := `UPDATE users
	SET updated_at = $1, email = $2, username = $3, password = $4
	WHERE id = $5
	RETURNING id, created_at, updated_at, email, username, password
	`
	if !userParams.UpdatedAt.Valid {
		userParams.UpdatedAt = sql.NullString{
			String: util.GetCurrentDate(),
			Valid:  true,
		}
	}
	row := db.conn.QueryRow(updateUserQuery,
		userParams.UpdatedAt,
		userParams.Email,
		userParams.Username,
		userParams.Password,
		userParams.ID,
	)
	var user models.Users
	err := row.Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Email,
		&user.Username,
		&user.Password,
	)
	return user, err
}

func (db Store) DeleteUser(id int64) error {
	deleteUserQuery := `DELETE FROM users
	WHERE id = $1
	`
	_, err := db.conn.Exec(deleteUserQuery, id)
	return err
}
