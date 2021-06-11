package store

import (
	"testing"

	"github.com/angeldhakal/tv-tracker/models"
	"github.com/angeldhakal/tv-tracker/util"
	"github.com/stretchr/testify/require"
)

func createRandomToken(t *testing.T) models.Tokens {
	user := createRandomUser(t)

	arg := CreateTokenParams{
		CreatedAt: util.GetCurrentDate(),
		Token:     util.GenerateTokenUUID(),
		UserID:    user.ID,
	}

	token, err := Repo.CreateToken(arg)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.Equal(t, arg.CreatedAt, token.CreatedAt)
	require.Equal(t, arg.Token, token.Token)
	require.Equal(t, arg.UserID, token.UserID)

	require.NotZero(t, token.UserID)

	return token

}

func TestCreateToken(t *testing.T) {
	createRandomToken(t)
}

func TestGetToken(t *testing.T) {
	token := createRandomToken(t)

	token1, err := Repo.GetToken(token.ID)
	require.NoError(t, err)
	require.NotEmpty(t, token1)

	require.Equal(t, token.ID, token1.ID)
	require.Equal(t, token.Token, token1.Token)
	require.Equal(t, token.UserID, token1.UserID)
	require.Equal(t, token.CreatedAt, token1.CreatedAt)
	require.Equal(t, token.UpdatedAt.String, token1.UpdatedAt.String)
	require.Equal(t, token.UpdatedAt.Valid, token1.UpdatedAt.Valid)
}

func TestGetTokenByToken(t *testing.T) {
	token := createRandomToken(t)

	token1, err := Repo.GetTokenByToken(token.Token)
	require.NoError(t, err)
	require.NotEmpty(t, token1)

	require.Equal(t, token.ID, token1.ID)
	require.Equal(t, token.Token, token1.Token)
	require.Equal(t, token.UserID, token1.UserID)
	require.Equal(t, token.CreatedAt, token1.CreatedAt)
	require.Equal(t, token.UpdatedAt.String, token1.UpdatedAt.String)
	require.Equal(t, token.UpdatedAt.Valid, token1.UpdatedAt.Valid)
}

func TestGetTokenByUser(t *testing.T) {
	token := createRandomToken(t)

	token1, err := Repo.GetTokenByUser(token.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, token1)

	require.Equal(t, token.ID, token1.ID)
	require.Equal(t, token.Token, token1.Token)
	require.Equal(t, token.UserID, token1.UserID)
	require.Equal(t, token.CreatedAt, token1.CreatedAt)
	require.Equal(t, token.UpdatedAt.String, token1.UpdatedAt.String)
	require.Equal(t, token.UpdatedAt.Valid, token1.UpdatedAt.Valid)
}

func TestDeleteToken(t *testing.T) {
	token := createRandomToken(t)

	err := Repo.DeleteToken(token.ID)
	require.NoError(t, err)
}

func TestListTokens(t *testing.T) {
	for i := 0; i <= 5; i++ {
		createRandomToken(t)
	}
	tokens, err := Repo.ListTokens()
	require.NoError(t, err)

	for _, token := range tokens {
		require.NotEmpty(t, token)
	}
}
