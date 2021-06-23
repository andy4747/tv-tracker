package routes

import (
	"net/http"
	"time"

	"github.com/angeldhakal/tv-tracker/handlers"
	"github.com/angeldhakal/tv-tracker/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var AllowedOrigins = []string{"http://localhost:3000/"}

var AllowedMethods = []string{"GET", "POST"}

var AllowedHeaders = []string{"Content-Type"}

var ExposedHeaders = []string{"Content-Length"}

func Home(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Home Path")
}

func MainRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     AllowedOrigins,
		AllowMethods:     AllowedMethods,
		AllowHeaders:     AllowedHeaders,
		ExposeHeaders:    ExposedHeaders,
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/", Home)

	apiRoutes := router.Group("/api")
	//user routes block
	{
		authHandler := handlers.NewUserHandler()
		authRoutes := apiRoutes.Group("/auth")
		authRoutes.POST("/register", authHandler.Signup)
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.GET("/logout", authHandler.Logout)
		authRoutes.POST("/de-register", authHandler.DeleteUser)
		authRoutes.GET("/user", middlewares.Protected(), authHandler.GetUser)
	}

	{
		movieHandler := handlers.NewMovieHandler()
		movieTrackRoutes := apiRoutes.Group("/track/movies")
		movieTrackRoutes.Use(middlewares.Protected())
		movieTrackRoutes.POST("/create", movieHandler.CreateMovie)
		movieTrackRoutes.PUT("/update/:id", movieHandler.UpdateMovie)
		movieTrackRoutes.PATCH("/update/:id", movieHandler.PatchMovie)
		movieTrackRoutes.GET("/get/:id", movieHandler.GetMovie)
		movieTrackRoutes.GET("/get-movies", movieHandler.GetMovies)
		movieTrackRoutes.DELETE("/delete/:id", movieHandler.GetMovies)
	}
	return router
}
