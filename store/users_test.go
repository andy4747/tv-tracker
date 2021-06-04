package store

import (
	"database/sql"
	"testing"

	"github.com/angeldhakal/tv-tracker/models"
	"github.com/angeldhakal/tv-tracker/util"
	"github.com/stretchr/testify/require"
)

var Repo = NewStore()

func createRandomUser(t *testing.T) models.Users {
	password := "root"
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	arg := CreateUserParams{
		CreatedAt: util.GetCurrentDate(),
		Email:     util.RandomEmail(),
		Username:  util.RandomUsername(),
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

	return user
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := Repo.GetUser(user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	//testing if all the fields of user1 is equal to user2
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)

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
