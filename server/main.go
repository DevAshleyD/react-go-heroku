package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve the frontend from build under /
	router.Use(static.Serve("/", static.LocalFile("./client/build/", true)))

	// CORS middleware -- TAKE THIS OUT IN PRODUCTION
	// The * is a security vulnerability
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Serve the api under the route group /api/*
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Pong!",
			})
		})
	}

	// Start the app
	router.Run(":5555")
}