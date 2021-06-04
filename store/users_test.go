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
	hashedPassword, err := util.RandomHashedPassword()
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
	require.Equal(t, arg.Password, user.Password)

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

func TestGetUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := Repo.GetUserByEmail(user1.Email)

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
	user := createRandomUser(t)

	hashedPass, err := util.RandomHashedPassword()
	require.NoError(t, err)

	arg := UpdateUserParams{
		UpdatedAt: sql.NullString{
			String: util.GetCurrentDate(),
			Valid:  true,
		},
		Email:    util.RandomEmail(),
		Username: util.RandomUsername(),
		Password: hashedPass,
		ID:       user.ID,
	}
	//checking is updated_at is valid or not
	require.True(t, arg.UpdatedAt.Valid)

	user1, err := Repo.UpdateUser(arg)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, arg.Email, user1.Email)
	require.Equal(t, arg.Username, user1.Username)
	require.Equal(t, arg.UpdatedAt.String, user1.UpdatedAt.String)

	//checking if updated_at of updated user is valid or not
	require.True(t, user1.UpdatedAt.Valid)

	require.NotZero(t, user1.ID)

}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)

	err := Repo.DeleteUser(user.ID)
	require.NoError(t, err)
}

func TestDeleteUserByEmail(t *testing.T) {
	user := createRandomUser(t)

	err := Repo.DeleteUserByEmail(user.Email)
	require.NoError(t, err)
}

func TestListUsers(t *testing.T) {
	for i := 0; i <= 10; i++ {
		createRandomUser(t)
	}
	users, err := Repo.ListUsers()
	require.NoError(t, err)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

}
