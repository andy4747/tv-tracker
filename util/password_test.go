package util

import (
	"fmt"
	"testing"
)

func TestHashedPassword(t *testing.T) {
	password := "root"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("Password Cannot Be Hashed err: %v\n", err)
	}

	if err == fmt.Errorf("cannot hash the password") {
		t.Errorf("Password Cannot Be Hashed err: %v\n", err)
	}

	if hashedPassword == "" {
		t.Errorf("Password Cannot Be Hashed err: %v\n", err)
	}

	err = ComparePassword(password, hashedPassword)
	if err != nil {
		t.Errorf("Wrong Password err: %v\n", err)
	}
}
