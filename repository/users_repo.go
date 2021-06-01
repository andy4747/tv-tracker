package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/angeldhakal/tv-tracker/models"
	"github.com/angeldhakal/tv-tracker/util"
)

type UserRepository interface {
	GetUser(int64) (models.Users, error)
	GetUserByEmail(string) (models.Users, error)
	ListUsers() ([]models.Users, error)
	CreateUser(models.Users) (models.Users, error)
	UpdateUser(models.Users) (models.Users, error)
	DeleteUser(int64) error
}

type userRepo struct {
	conn *models.Queries
}

func NewUserRepository() UserRepository {
	return &userRepo{
		conn: models.Connect(),
	}
}

func (db *userRepo) GetUser(id int64) (user models.Users, err error) {
	user, err = db.conn.GetUser(context.Background(), id)
	if err != nil {
		log.Println("userRepo GetUser: ", err.Error())
		return models.Users{}, err
	}
	return user, nil
}

func (db *userRepo) GetUserByEmail(email string) (user models.Users, err error) {
	user, err = db.conn.GetUserByEmail(context.Background(), email)
	if err != nil {
		log.Println("userRepo GetUserByEmail: ", err.Error())
		return models.Users{}, err
	}
	return user, nil
}

func (db *userRepo) ListUsers() (users []models.Users, err error) {
	users, err = db.conn.ListUsers(context.Background())
	if err != nil {
		log.Println("userRepo ListUsers: ", err.Error())
		return []models.Users{}, err
	}
	return users, nil
}

func (db *userRepo) CreateUser(user models.Users) (createdUser models.Users, err error) {
	userParams := models.CreateUserParams{
		CreatedAt: util.GetCurrentDate(),
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
	}
	createdUser, err = db.conn.CreateUser(context.Background(), userParams)
	if err != nil {
		log.Println("userRepo CreateUser: ", err.Error())
		return models.Users{}, err
	}
	return createdUser, nil
}

func (db *userRepo) UpdateUser(user models.Users) (updatedUser models.Users, err error) {
	userParams := models.UpdateUserParams{
		ID: user.ID,
		UpdatedAt: sql.NullString{
			String: util.GetCurrentDate(),
			Valid:  true,
		},
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}
	updatedUser, err = db.conn.UpdateUser(context.Background(), userParams)
	if err != nil {
		log.Println("userRepo UpdateUser: ", err.Error())
		return models.Users{}, err
	}
	return updatedUser, nil
}

func (db userRepo) DeleteUser(id int64) error {
	err := db.conn.DeleteUser(context.Background(), id)
	if err != nil {
		log.Println("userRepo DeleteUser: ", err.Error())
		return err
	}
	return nil
}
