package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/angeldhakal/tv-tracker/models"
	"github.com/angeldhakal/tv-tracker/repository"
	"github.com/angeldhakal/tv-tracker/util"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	repo repository.UserRepository
}

func NewUserHandler() userHandler {
	return userHandler{
		repo: repository.NewUserRepository(),
	}
}

func (h *userHandler) Logout(ctx *gin.Context) {
	tokenFromRequest := ctx.GetHeader("token")
	userToken, err := models.Connect().GetTokenByToken(context.Background(), tokenFromRequest)
	if err != nil {
		log.Printf("token not found err: %v", err)
		ctx.JSON(http.StatusNotFound, "token not found")
		return
	}
	err = models.Connect().DeleteToken(context.Background(), userToken.ID)
	if err != nil {
		log.Printf("some error occurred err: %v", err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}
	ctx.JSON(http.StatusOK, "Successfully Logged Out")
}

func (h *userHandler) Login(ctx *gin.Context) {
	var credential util.LoginCredentials
	if err := ctx.ShouldBindJSON(&credential); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "enter valid json")
		return
	}
	if credential.Email == "" && credential.Password == "" {
		ctx.JSON(http.StatusNotFound, "please enter valid credentials")
		return
	}
	user, err := h.repo.GetUserByEmail(credential.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, fmt.Sprintf("Email \"%s\" not registered", credential.Email))
			return
		}
		ctx.JSON(http.StatusNotFound, "new error occured")
		return
	}
	//verify user credentials
	err = util.ComparePassword(credential.Password, user.Password)
	if err != nil {
		log.Printf("passwords don't match\n")
		ctx.JSON(http.StatusNotFound, "please enter valid password")
		return
	}

	var token string
	userToken, err := models.Connect().GetTokenByUser(context.Background(), user.ID)
	token = userToken.Token
	if err != nil {
		if err == sql.ErrNoRows {
			generatedToken := util.GenerateTokenUUID()
			fmt.Println(generatedToken)

			//saving session token in the db
			savedToken, err := models.Connect().CreateToken(context.Background(), models.CreateTokenParams{
				CreatedAt: util.GetCurrentDate(),
				Token:     generatedToken,
				UserID:    user.ID,
			})
			if err != nil {
				log.Printf("couldn't save the row to tokens table err: %v\n", err)
				ctx.JSON(http.StatusInternalServerError, "")
				return
			}
			fmt.Println("Token Saved")
			token = savedToken.Token
			log.Printf("no token row existed so created one")
		} else {
			log.Printf("some error occured err: %v\n", err)
			ctx.JSON(http.StatusInternalServerError, "some error occureed")
			return
		}
	}

	// create a token cookie
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(12 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(ctx.Writer, cookie)
	ctx.JSON(http.StatusOK, token)
}

func (h *userHandler) Signup(ctx *gin.Context) {
	var credential util.RegistrationCredentials
	if err := ctx.ShouldBindJSON(&credential); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "enter valid json")
		return
	}

	// verify the passwords
	if credential.Password1 != credential.Password2 || credential.Password1 == "" && credential.Password2 == "" {
		ctx.JSON(http.StatusBadRequest, "enter valid passwords")
		return
	}

	var user models.Users
	user.Email = credential.Email
	user.Username = credential.Username

	//hashing the password
	hashedPassword, err := util.HashPassword(credential.Password2)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "couldn't hash the password")
		return
	}

	user.Password = hashedPassword

	//add user credential in the database
	addedUser, err := h.repo.CreateUser(user)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, "Email already taken.")
		return
	}
	ctx.JSON(http.StatusOK, addedUser)
}
