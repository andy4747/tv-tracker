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
	Language      string        `json:"lang"`
}

type UpdateMovieParams struct {
	UpdatedAt     sql.NullString `json:"updated_at"`
	Name          string         `json:"name"`
	Status        models.Status  `json:"status"`
	CurrentLength int64          `json:"current_length"`
	Language      string         `json:"lang"`
	ID            int64          `json:"id"`
}

type MovieTracker interface {
	GetMovie(int64) (models.Movies, error)
	GetMoviesByUser(int64) ([]models.Movies, error)
	CreateMovie(CreateMovieParams) (models.Movies, error)
	UpdateMovie(UpdateMovieParams) (models.Movies, error)
	UpdateName(UpdateMovieParams) (models.Movies, error)
	UpdateStatus(UpdateMovieParams) (models.Movies, error)
	UpdateLanguage(UpdateMovieParams) (models.Movies, error)
	UpdateCurrentLength(UpdateMovieParams) (models.Movies, error)
	DeleteMovie(int64) error
	ListMovies() ([]models.Movies, error)
}

type movieStore struct {
	conn *sql.DB
}

func (db *movieStore) GetMovie(userID int64) (models.Movies, error) {
	getMovieQuery := `SELECT id, created_at, updated_at, name, status, current_length, user_id, language FROM movies WHERE id=$1;
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
		&movie.Language,
	)
	return movie, err
}

func (db *movieStore) GetMoviesByUser(userID int64) ([]models.Movies, error) {
	getMovieQuery := `SELECT id, created_at, updated_at, name, status, current_length, user_id, language FROM movies WHERE user_id=$1;
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
			&movie.Language,
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

func (db *movieStore) ListMovies() ([]models.Movies, error) {
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

func (db *movieStore) CreateMovie(movieParams CreateMovieParams) (models.Movies, error) {
	createMovie := `INSERT INTO movies (created_at, user_id, name, status, current_length, language)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, created_at, updated_at, user_id, name, status, current_length
	, language;
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
		movieParams.Language,
	)
	var movie models.Movies
	err := row.Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.UserID,
		&movie.Name,
		&movie.Status,
		&movie.CurrentLength,
		&movie.Language,
	)
	return movie, err
}

func (db *movieStore) UpdateMovie(movieParams UpdateMovieParams) (models.Movies, error) {
	updateMovies := `UPDATE movies
	SET updated_at = $1, name = $2, status = $3, current_length = $4
	WHERE id = $5
	RETURNING id, created_at, updated_at, user_id, name, status,current_length, language;
	`
	if !movieParams.UpdatedAt.Valid {
		movieParams.UpdatedAt = sql.NullString{
			String: util.GetCurrentDate(),
			Valid:  true,
		}
	}
	row := db.conn.QueryRow(updateMovies,
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
		&movie.UserID,
		&movie.Name,
		&movie.Status,
		&movie.CurrentLength,
		&movie.Language,
	)
	return movie, err
}

func (db *movieStore) UpdateName(movieParams UpdateMovieParams) (models.Movies, error) {
	updateNameQuery := `UPDATE movies
	SET updated_at = $1, name = $2
	WHERE id = $3
	RETURNING id, created_at, updated_at, user_id, name, status,current_length, language;
	`
	if !movieParams.UpdatedAt.Valid {
		movieParams.UpdatedAt = sql.NullString{
			String: util.GetCurrentDate(),
			Valid:  true,
		}
	}
	row := db.conn.QueryRow(updateNameQuery,
		movieParams.UpdatedAt,
		movieParams.Name,
		movieParams.ID,
	)
	var movie models.Movies
	err := row.Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.UserID,
		&movie.Name,
		&movie.Status,
		&movie.CurrentLength,
		&movie.Language,
	)
	return movie, err
}

func (db *movieStore) UpdateLanguage(movieParams UpdateMovieParams) (models.Movies, error) {
	updateNameQuery := `UPDATE movies
	SET updated_at = $1, language = $2
	WHERE id = $3
	RETURNING id, created_at, updated_at, user_id, name, status,current_length, language;
	`
	if !movieParams.UpdatedAt.Valid {
		movieParams.UpdatedAt = sql.NullString{
			String: util.GetCurrentDate(),
			Valid:  true,
		}
	}
	row := db.conn.QueryRow(updateNameQuery,
		movieParams.UpdatedAt,
		movieParams.Language,
		movieParams.ID,
	)
	var movie models.Movies
	err := row.Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.UserID,
		&movie.Name,
		&movie.Status,
		&movie.CurrentLength,
		&movie.Language,
	)
	return movie, err
}

func (db *movieStore) UpdateStatus(movieParams UpdateMovieParams) (models.Movies, error) {
	updateStatusQuery := `UPDATE movies
	SET updated_at = $1, status = $2
	WHERE id = $3
	RETURNING id, created_at, updated_at, user_id, name, status,current_length, language;
	`
	if !movieParams.UpdatedAt.Valid {
		movieParams.UpdatedAt = sql.NullString{
			String: util.GetCurrentDate(),
			Valid:  true,
		}
	}
	row := db.conn.QueryRow(updateStatusQuery,
		movieParams.UpdatedAt,
		movieParams.Status,
		movieParams.ID,
	)
	var movie models.Movies
	err := row.Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.UserID,
		&movie.Name,
		&movie.Status,
		&movie.CurrentLength,
		&movie.Language,
	)
	return movie, err
}

func (db *movieStore) UpdateCurrentLength(movieParams UpdateMovieParams) (models.Movies, error) {
	updateCurrentLengthQuery := `UPDATE movies
	SET updated_at = $1, current_length = $2
	WHERE id = $3
	RETURNING id, created_at, updated_at, user_id, name, status, current_length, language;
	`
	if !movieParams.UpdatedAt.Valid {
		movieParams.UpdatedAt = sql.NullString{
			String: util.GetCurrentDate(),
			Valid:  true,
		}
	}
	row := db.conn.QueryRow(updateCurrentLengthQuery,
		movieParams.UpdatedAt,
		movieParams.CurrentLength,
		movieParams.ID,
	)
	var movie models.Movies
	err := row.Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.UpdatedAt,
		&movie.UserID,
		&movie.Name,
		&movie.Status,
		&movie.CurrentLength,
		&movie.Language,
	)
	return movie, err
}

func (db *movieStore) DeleteMovie(userID int64) error {
	deleteToken := `DELETE FROM movies
	WHERE id = $1
	`
	_, err := db.conn.Exec(deleteToken, userID)
	return err
}
