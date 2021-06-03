package util

import (
	"strings"
	"testing"
)

func TestGenerateTokenUUID(t *testing.T) {
	token := GenerateTokenUUID()
	if token == "" {
		t.Errorf("token not generated output: %v\n", token)
	} else if strings.Contains(token, "-") {
		t.Errorf("invalid token generated output: %v\n", token)
	} else if strings.Contains(token, " ") {
		t.Errorf("invalid token generated output: %v\n", token)
	}
}
