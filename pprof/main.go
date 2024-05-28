package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	log.Println("Booting on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
