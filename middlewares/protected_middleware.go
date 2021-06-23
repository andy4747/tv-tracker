package middlewares

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/angeldhakal/tv-tracker/store"
	"github.com/gin-gonic/gin"
)

func Protected() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userToken, err := ctx.Cookie("token")
		if userToken == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "Login to access this resource")
			return
		} else if len(userToken) != 32 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "Login to access this resource")
			return
		}
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "No Cookie Found")
			return
		}
		tokenStore := store.NewTokenStore()
		token, err := tokenStore.GetTokenByToken(userToken)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.AbortWithStatusJSON(http.StatusNotFound, "Login to access this resource")
				return
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, "some server error occured")
				return
			}
		}
		user_id := strconv.FormatInt(token.UserID, 10)
		ctx.Request.Header.Add("user_token", token.Token)
		ctx.Request.Header.Add("user_id", user_id)
		ctx.Next()
	}
}
