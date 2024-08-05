package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func getAlbums(c *gin.Context) {
	response, err := http.Get("https://cdn.shriramfinance.in/")
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}

	defer response.Body.Close()

	var albums []album
	err = json.NewDecoder(response.Body).Decode(&albums)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)

	router.Run(":8080")
}
