package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HostSwitch maps host names to http.Handlers
type HostSwitch map[string]http.Handler

// ServeHTTP method to handle requests based on the host
func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check if a http.Handler is registered for the given host.
	// If yes, use it to handle the request.
	if handler := hs[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		// Handle host names for which no handler is registered
		http.Error(w, "Forbidden", http.StatusForbidden) // Or Redirect?
	}
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	router2 := gin.Default()

	// Define the Index handler
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome!")
	})

	router2.GET("/by", func(ctx *gin.Context) {
		c.String(http.StatusOK, "By")
	})

	// Define the Hello handler
	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s!", name)
	})

	// Make a new HostSwitch and insert the router (our http handler)
	// for example.com and port 12345
	hs := make(HostSwitch)
	hs["example.com:12345"] = router

	// Use the HostSwitch to listen and serve on port 12345
	log.Fatal(http.ListenAndServe(":12345", hs))
}
