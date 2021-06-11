package store

import (
	"database/sql"
	"testing"
	"time"

	"github.com/angeldhakal/tv-tracker/models"
	"github.com/angeldhakal/tv-tracker/util"
	"github.com/stretchr/testify/require"
)

func createRandomMovie(t *testing.T) models.Movies {
	user := createRandomUser(t)

	arg := CreateMovieParams{
		CreatedAt:     util.GetCurrentDate(),
		UserID:        user.ID,
		Name:          util.RandomMovie(),
		Status:        models.Completed,
		CurrentLength: 0,
	}

	movie, err := Repo.CreateMovie(arg)

	require.NoError(t, err)
	require.NotEmpty(t, movie)

	require.Equal(t, arg.CreatedAt, movie.CreatedAt)
	require.Equal(t, arg.UserID, movie.UserID)
	require.Equal(t, arg.Name, movie.Name)
	require.Equal(t, arg.Status, movie.Status)
	require.Equal(t, arg.CurrentLength, movie.CurrentLength)

	require.NotZero(t, movie.ID)
	require.NotZero(t, movie.UserID)

	return movie
}

func TestCreateMovie(t *testing.T) {
	createRandomMovie(t)
}

func TestGetMovie(t *testing.T) {
	movie := createRandomMovie(t)

	movie1, err := Repo.GetMovie(movie.ID)
	require.NoError(t, err)
	require.NotEmpty(t, movie1)

	require.Equal(t, movie.CreatedAt, movie1.CreatedAt)
	require.Equal(t, movie.UserID, movie1.UserID)
	require.Equal(t, movie.Name, movie1.Name)
	require.Equal(t, movie.Status, movie1.Status)
	require.Equal(t, movie.CurrentLength, movie1.CurrentLength)

}

func TestGetMovieByUser(t *testing.T) {
	for i := 0; i <= 5; i++ {
		createRandomMovie(t)
	}
	movies, err := Repo.ListMovies()
	require.NoError(t, err)

	for _, movie := range movies {
		require.NotEmpty(t, movie)
	}
}

func TestUpdateMovie(t *testing.T) {
	movie := createRandomMovie(t)

	arg := UpdateMovieParams{
		UpdatedAt: sql.NullString{
			String: util.GetCurrentDate(),
			Valid:  true,
		},
		Name:          util.RandomMovie(),
		Status:        models.Waiting,
		CurrentLength: int64(time.Minute * 63),
		ID:            movie.ID,
	}

	//checking is updated_at is valid or not
	require.True(t, arg.UpdatedAt.Valid)

	movie1, err := Repo.UpdateMovie(arg)
	require.NoError(t, err)
	require.NotEmpty(t, movie1)
	require.NotZero(t, movie1.ID)
	require.True(t, movie1.UpdatedAt.Valid)

	require.Equal(t, arg.UpdatedAt.String, movie1.UpdatedAt.String)
	require.Equal(t, arg.Name, movie1.Name)
	require.Equal(t, arg.Status, movie1.Status)
	require.Equal(t, arg.CurrentLength, movie1.CurrentLength)
}

func TestDeleteMovie(t *testing.T) {
	movie := createRandomMovie(t)

	err := Repo.DeleteMovie(movie.ID)
	require.NoError(t, err)
}

func TestListMovies(t *testing.T) {
	for i := 0; i <= 5; i++ {
		createRandomMovie(t)
	}
	movies, err := Repo.ListMovies()
	require.NoError(t, err)

	for _, movie := range movies {
		require.NotEmpty(t, movie)
	}
}
