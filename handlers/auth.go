package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/angeldhakal/tv-tracker/store"
	"github.com/angeldhakal/tv-tracker/util"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userRepo  store.UserTracker
	tokenRepo store.TokenTracker
}

func NewUserHandler() userHandler {
	return userHandler{
		userRepo:  store.NewUserStore(),
		tokenRepo: store.NewTokenStore(),
	}
}

func (h *userHandler) Logout(ctx *gin.Context) {
	tokenFromRequest := ctx.GetHeader("token")
	userToken, err := h.tokenRepo.GetTokenByToken(tokenFromRequest)
	if err != nil {
		log.Printf("token not found err: %v", err)
		ctx.JSON(http.StatusNotFound, "token not found")
		return
	}
	err = h.tokenRepo.DeleteToken(userToken.ID)
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
	user, err := h.userRepo.GetUserByEmail(credential.Email)
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
	userToken, err := h.tokenRepo.GetTokenByUser(user.ID)
	token = userToken.Token
	if err != nil {
		if err == sql.ErrNoRows {
			generatedToken := util.GenerateTokenUUID()
			fmt.Println(generatedToken)

			//saving session token in the db
			savedToken, err := h.tokenRepo.CreateToken(store.CreateTokenParams{
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

	var userParam store.CreateUserParams
	userParam.Email = credential.Email
	userParam.Username = credential.Username

	//hashing the password
	hashedPassword, err := util.HashPassword(credential.Password2)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "couldn't hash the password")
		return
	}

	userParam.Password = hashedPassword

	//add user credential in the database
	addedUser, err := h.userRepo.CreateUser(userParam)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, "Email already taken.")
		return
	}
	ctx.JSON(http.StatusOK, addedUser)
}

func (h *userHandler) DeleteUser(ctx *gin.Context) {
	var credential util.LoginCredentials
	if err := ctx.ShouldBindJSON(&credential); err != nil {
		log.Printf("Invalid JSON entered")
		ctx.JSON(http.StatusUnprocessableEntity, "enter valid JSON data")
		return
	}
	//validating credentials
	if credential.Password == "" || credential.Email == "" {
		log.Printf("enter valid data")
		ctx.JSON(http.StatusBadRequest, "enter valid email and password")
		return
	}
	retrievedUser, err := h.userRepo.GetUserByEmail(credential.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found with %s as email", credential.Email)
			ctx.JSON(http.StatusNotFound, fmt.Sprintf("No user found with %s as email", credential.Email))
			return
		} else {
			log.Printf("some server error occurred err: %v", err)
			ctx.JSON(http.StatusInternalServerError, "somer server error occurred")
			return
		}
	}
	err = util.ComparePassword(credential.Password, retrievedUser.Password)
	if err != nil {
		log.Printf("password don't match\n")
		ctx.JSON(http.StatusBadRequest, "password don't match. try again...")
		return
	}
	err = h.userRepo.DeleteUser(retrievedUser.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("user doesn't exist")
			ctx.JSON(http.StatusNotFound, "user doesn't exists")
			return
		} else {
			log.Printf("some server error occurred err: %v", err)
			ctx.JSON(http.StatusInternalServerError, "somer server error occurred")
			return
		}
	}
	ctx.JSON(http.StatusOK, "Successfully Deleted")
}
