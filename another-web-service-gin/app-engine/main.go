package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func exampleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Hi, your api is reachable")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to %s port\n", port)
	}

	router := gin.New()

	router.GET("/exampleHandler", exampleHandler)

	log.Printf("Listening on port %s \n", port)

	router.Run(":" + port)
}
