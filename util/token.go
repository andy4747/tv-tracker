package util

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateTokenUUID() string {
	id := uuid.New()
	token := strings.Replace(id.String(), "-", "", -1)
	return token
}
