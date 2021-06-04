package store

import (
	"database/sql"
	"testing"

	"github.com/angeldhakal/tv-tracker/util"
	"github.com/stretchr/testify/require"
)

var Repo = NewStore()

func TestCreateUser(t *testing.T) {
	password := "root"

	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	arg := CreateUserParams{
		CreatedAt: util.GetCurrentDate(),
		Email:     "angel@gmail.com",
		Username:  "andy",
		Password:  hashedPassword,
	}

	//creating a user
	user, err := Repo.CreateUser(arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.CreatedAt, user.CreatedAt)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Username, user.Username)

	//if password and hashed password don't match an error is thrown
	err = util.ComparePassword(password, user.Password)
	require.NoError(t, err)

	require.NotZero(t, user.ID)

}

func TestUpdateUser(t *testing.T) {
	password := "mradhakal"
	hashedPass, err := util.HashPassword(password)
	require.NoError(t, err)

	arg := UpdateUserParams{
		UpdatedAt: sql.NullString{
			String: util.GetCurrentDate(),
			Valid:  true,
		},
		Email:    "angeldhakal@gmail.com",
		Username: "anton",
		Password: hashedPass,
		ID:       1,
	}
	//checking is updated_at is valid or not
	require.True(t, arg.UpdatedAt.Valid)

	user, err := Repo.UpdateUser(arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.UpdatedAt.String, user.UpdatedAt.String)

	//checking if updated_at of updated user is valid or not
	require.True(t, user.UpdatedAt.Valid)

	require.NotZero(t, user.ID)

}
