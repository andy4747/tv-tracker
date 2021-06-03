package store

import (
	"testing"

	"github.com/angeldhakal/tv-tracker/models"
	"github.com/angeldhakal/tv-tracker/util"
)

var Repo = NewStore()

func TestCreateUser(t *testing.T) {
	password := "root"
	hashedPassword, _ := util.HashPassword(password)
	arg := CreateUserParams{
		CreatedAt: util.GetCurrentDate(),
		Email:     "angel@gmail.com",
		Username:  "andy",
		Password:  hashedPassword,
	}
	user, err := Repo.CreateUser(arg)
	if err != nil {
		t.Error(err)
	}
	if user == (models.Users{}) {
		t.Error(err)
	}

	if arg.CreatedAt != user.CreatedAt {
		t.Errorf("created_at date doesn't matches err: %v\n", err)
	} else if arg.Email != user.Email {
		t.Errorf("Email doesn't matches err: %v\n", err)
	} else if arg.Username != user.Username {
		t.Errorf("Username doesn't matches err: %v\n", err)
	} else if err := util.ComparePassword(password, user.Password); err != nil {
		t.Error(err)
	} else if user.ID == 0 {
		t.Errorf("id must not be zero err: %v\n", err)
	}
}
