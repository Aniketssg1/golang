package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Dog struct {
	Breed         string `json:"breed" validate:"required"`
	FavoriteTreat string `json:"FavoriteTreats" validate:"required"`
	Name          string `json:"Name" validate:"required"`
	Age           int    `json:"Age" validate:"required,min=0,max=99"`
}

func handler(rw http.ResponseWriter, request *http.Request) {
	responseData := map[string]string{"message": "Hello your server is live"}

	responseJson, err := json.Marshal(responseData)
	if err != nil {
		http.Error(rw, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	//rw.Header().Set("Content-Type", "application/json")

	rw.Write(responseJson)
}

func marshalHooman(rw http.ResponseWriter, request *http.Request) {
	newDog := Dog{
		Breed:         "Poodle",
		Age:           15,
		Name:          "Chloe",
		FavoriteTreat: "Dog Sticks",
	}
	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(newDog); err != nil {
		log.Fatalf("Unable to encode due to %s\n", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/hooman", marshalHooman)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
