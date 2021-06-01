package routes

import (
	"github.com/angeldhakal/tv-tracker/handlers"
	"github.com/gin-gonic/gin"
)

func MainRouter() *gin.Engine {
	router := gin.Default()
	apiRoutes := router.Group("/api")
	//user routes block
	{
		authHandler := handlers.NewUserHandler()
		authRoutes := apiRoutes.Group("/auth")
		authRoutes.POST("/register", authHandler.Signup)
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.GET("/logout", authHandler.Logout)
	}
	return router
}
