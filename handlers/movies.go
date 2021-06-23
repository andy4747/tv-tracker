package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/angeldhakal/tv-tracker/models"
	"github.com/angeldhakal/tv-tracker/store"
	"github.com/gin-gonic/gin"
)

type movieHandler struct {
	movieRepo store.MovieTracker
}

func NewMovieHandler() movieHandler {
	return movieHandler{
		movieRepo: store.NewMovieStore(),
	}
}

func (h *movieHandler) TrackMovie(ctx *gin.Context) {
	userID := ctx.GetHeader("user_id")
	userToken := ctx.GetHeader("user_token")
	dict := make(map[string]string)
	dict["token"] = userToken
	dict["id"] = userID
	ctx.JSON(200, dict)
}

func (h *movieHandler) CreateMovie(ctx *gin.Context) {
	userID, err := strconv.ParseInt(ctx.GetHeader("user_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Couldn't parse user")
		return
	}
	params := store.CreateMovieParams{}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "enter valid json")
		return
	}
	params.UserID = userID
	if params.Language == "" {
		params.Language = "English"
	}
	if params.Status == "w" {
		if params.CurrentLength == 0 {
			ctx.JSON(http.StatusBadRequest, "please enter current time in movie as well")
			return
		} else if params.CurrentLength < 0 {
			ctx.JSON(http.StatusBadRequest, "current time in movie should be positive")
			return
		}
	}
	movie, err := h.movieRepo.CreateMovie(params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "enter valid json")
		return
	}
	ctx.JSON(http.StatusOK, movie)
}

func (h *movieHandler) PatchMovie(ctx *gin.Context) {
	movieID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "please enter a valid integer as id")
		return
	}
	_, err = h.movieRepo.GetMovie(movieID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "no movie to update")
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, "some error occurred")
			return
		}
	}
	params := store.UpdateMovieParams{}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "enter valid json")
		return
	}
	params.ID = movieID
	if params.Name != "" {
		_, err := h.movieRepo.UpdateName(params)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, "cannot update name")
			return
		}
	}
	if params.Status != "" {
		if params.Status == "w" {
			if params.CurrentLength == 0 {
				ctx.JSON(http.StatusBadRequest, "please enter current time in movie as well")
				return
			} else if params.CurrentLength < 0 {
				ctx.JSON(http.StatusBadRequest, "current time in movie should be positive")
				return
			} else {
				_, err := h.movieRepo.UpdateStatus(params)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, "cannot update status")
					return
				}
				_, err = h.movieRepo.UpdateCurrentLength(params)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, "cannot update current length")
					return
				}
			}
		} else if params.Status == "p" || params.Status == "c" {
			_, err := h.movieRepo.UpdateStatus(params)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, "cannot update status")
				return
			}
		}
	}

	if params.Language != "" {
		_, err := h.movieRepo.UpdateLanguage(params)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "cannot update language")
			return
		}
	}

	movie, err := h.movieRepo.GetMovie(params.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, "no movie to update")
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, "some error occurred")
			return
		}
	}
	ctx.JSON(http.StatusOK, movie)
}

func (h *movieHandler) GetMovie(ctx *gin.Context) {
	movieID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "please enter a valid integer as id")
		return
	}
	movie, err := h.movieRepo.GetMovie(movieID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, models.EmptyObject)
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, "some error occurred")
			return
		}
	}
	ctx.JSON(http.StatusOK, movie)
}

func (h *movieHandler) GetMovies(ctx *gin.Context) {
	userID, err := strconv.ParseInt(ctx.GetHeader("user_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Couldn't parse user")
		return
	}
	movies, err := h.movieRepo.GetMoviesByUser(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, models.EmptyMovieList)
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, "some error occurred")
			return
		}
	}
	if len(movies) == 0 {
		ctx.JSON(http.StatusNotFound, models.EmptyMovieList)
		return
	}
	ctx.JSON(http.StatusOK, movies)
}

func (h *movieHandler) DeleteMovie(ctx *gin.Context) {
	movieID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "please enter a valid integer as id")
		return
	}
	err = h.movieRepo.DeleteMovie(movieID)
	if err != nil {
		if err == sql.ErrNoRows {
			msg := fmt.Sprintf("no movie with %d found", movieID)
			ctx.JSON(http.StatusNotFound, msg)
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, "some unknown error occurred")
			return
		}
	}
	ctx.JSON(http.StatusNoContent, "")
}

func (h *movieHandler) UpdateMovie(ctx *gin.Context) {
	movieID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "please enter a valid integer as id")
		return
	}
	param := store.UpdateMovieParams{}
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "enter valid data")
		return
	}
	param.ID = movieID
	movie, err := h.movieRepo.UpdateMovie(param)
	if err != nil {
		if err == sql.ErrNoRows {
			msg := fmt.Sprintf("no movie with %d found", movieID)
			ctx.JSON(http.StatusNotFound, msg)
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, "some unknown error occurred")
			return
		}
	}
	ctx.JSON(http.StatusOK, movie)
}
