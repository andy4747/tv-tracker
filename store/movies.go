package store

import (
	"database/sql"

	"github.com/angeldhakal/tv-tracker/models"
	"github.com/angeldhakal/tv-tracker/util"
)

type CreateMovieParams struct {
	CreatedAt     string        `json:"created_at"`
	UserID        int64         `json:"user_id"`
	Name          string        `json:"name"`
	Status        models.Status `json:"status"`
	CurrentLength int64         `json:"current_length"`
}

type UpdateMovieParams struct {
	UpdatedAt     sql.NullString `json:"updated_at"`
	Name          string         `json:"name"`
	Status        models.Status  `json:"status"`
	CurrentLength int64          `json:"current_length"`
	ID            int64          `json:"id"`
}

type MovieTracker interface {
	GetMovie(int64) (models.Movies, error)
	GetMoviesByUser(int64) ([]models.Movies, error)
	CreateMovie(CreateMovieParams) (models.Movies, error)
	UpdateMovie(UpdateMovieParams) (models.Movies, error)
	DeleteMovie(int64) error
	ListMovies() ([]models.Movies, error)
}

func (db *Store) GetMovie(userID int64) (models.Movies, error) {
	getMovieQuery := `SELECT id, created_at, updated_at, name, status, current_length, user_id FROM movies WHERE id=$1;
	`
	row := db.conn.QueryRow(getMovieQuery, userID)
	var movie models.Movies
	err := row.Scan(&movie.ID,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.Name,
		&movie.Status,
		&movie.CurrentLength,
		&movie.UserID,
	)
	return movie, err
}

func (db *Store) GetMoviesByUser(userID int64) ([]models.Movies, error) {
	getMovieQuery := `SELECT id, created_at, updated_at, name, status, current_length, user_id FROM movies WHERE user_id=$1;
	`
	rows, err := db.conn.Query(getMovieQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Movies
	for rows.Next() {
		var movie models.Movies
		if err := rows.Scan(&movie.ID,
			&movie.CreatedAt,
			&movie.UpdatedAt,
			&movie.Name,
			&movie.Status,
			&movie.CurrentLength,
			&movie.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, movie)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, err
}

func (db *Store) ListMovies() ([]models.Movies, error) {
	listTokens := `SELECT id, created_at, updated_at, name, status, current_length, user_id FROM movies;
	`
	rows, err := db.conn.Query(listTokens)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Movies
	for rows.Next() {
		var movie models.Movies
		if err := rows.Scan(
			&movie.ID,
			&movie.CreatedAt,
			&movie.UpdatedAt,
			&movie.Name,
			&movie.Status,
			&movie.CurrentLength,
			&movie.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, movie)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (db *Store) CreateMovie(movieParams CreateMovieParams) (models.Movies, error) {
	createMovie := `INSERT INTO movies (created_at, user_id, name, status, current_length)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at, updated_at, user_id, name, status, current_length
	, language, total_length;
	`
	if movieParams.CreatedAt == "" {
		movieParams.CreatedAt = util.GetCurrentDate()
	}
	row := db.conn.QueryRow(createMovie,
		movieParams.CreatedAt,
		movieParams.UserID,
		movieParams.Name,
		movieParams.Status,
		movieParams.CurrentLength,
	)
	var movie models.Movies
	err := row.Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.Name,
		&movie.Status,
		&movie.CurrentLength,
		&movie.Language,
		&movie.TotalLength,
	)
	return movie, err
}

func (db *Store) UpdateMovie(movieParams UpdateMovieParams) (models.Movies, error) {
	updateToken := `UPDATE movies
	SET updated_at = $1, name = $2, status = $3, current_length = $4,
	WHERE id = $5
	RETURNING id, created_at, updated_at, user_id, name, status,current_length, language, total_length;
	`
	if !movieParams.UpdatedAt.Valid {
		movieParams.UpdatedAt = sql.NullString{
			String: util.GetCurrentDate(),
			Valid:  true,
		}
	}
	row := db.conn.QueryRow(updateToken,
		movieParams.UpdatedAt,
		movieParams.Name,
		movieParams.Status,
		movieParams.CurrentLength,
		movieParams.ID,
	)
	var movie models.Movies
	err := row.Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.Name,
		&movie.Status,
		&movie.CurrentLength,
		&movie.Language,
		&movie.TotalLength,
	)
	return movie, err
}

func (db *Store) DeleteMovie(userID int64) error {
	deleteToken := `DELETE FROM movies
	WHERE id = $1
	`
	_, err := db.conn.Exec(deleteToken, userID)
	return err
}
