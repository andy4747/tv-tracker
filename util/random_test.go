package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	var max int64 = 20
	var min int64 = 10
	number := RandomInt(min, max)
	require.LessOrEqual(t, number, max)
	require.GreaterOrEqual(t, number, min)
	require.NotZero(t, number)
	require.Positive(t, number)
}

func TestRandomString(t *testing.T) {
	strLen := 5
	value := RandomString(strLen)
	require.Equal(t, strLen, len(value))
	require.NotZero(t, value)
	require.NotZero(t, len(value))
}

func TestRandomEmail(t *testing.T) {
	email := RandomEmail()
	splittedStr := strings.Split(email, "@")
	name := splittedStr[0]
	domain := splittedStr[1]

	require.NotEmpty(t, name)
	require.NotEmpty(t, domain)

	require.Contains(t, []string{"gmail.com", "outlook.com", "icloud.com"}, domain)

	require.Len(t, name, 5)

}
