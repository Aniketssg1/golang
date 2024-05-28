package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/go-playground/validator/v10"
	"github.com/sanity-io/litter"
)

type (
	BirthDate time.Time
	Baby      struct {
		BirthDate BirthDate `json:"birth_date"`
		Name      string    `json:"name"`
		Gender    string    `json:"gender"`
	}
)

type Dog struct {
	Breed          string `json:"breed" validate:"required"`
	FavoriteTreats string `json:"FavoriteTreats" validate:"required"`
	NAME           string `json:"Name" validate:"required"`
	Age            int    `json:"Age" validate:"required,min=0,max=99"`
}

type (
	Hooman struct {
		Address Address
		Name    string
		Pets    []Pet
		Age     int
	}
	Address struct {
		Line1  string
		Line2  string
		Postal string
	}
	Pet struct {
		Name  string
		Kind  string
		Color string
		Age   int
	}
)

func (bd *BirthDate) UnmarshalJSON(input []byte) error {
	value := strings.Trim(string(input), `"`)

	t, err := dateparse.ParseAny(value)
	if err != nil {
		t, err = time.Parse("2006/01/02", value)
		if err != nil {
			return err
		}
	}

	*bd = BirthDate(t)

	return nil
}

func (bd BirthDate) String() string {
	return time.Time(bd).Format("02-01-2006")
}

func main() {
	// Input of type JSON
	const input = `{
		"Name" : "Max",
		"Breed" : "Golden Retriever",
		"Age" : 30,
		"FavoriteTreats" : "Bone"
	}`

	// Check whether input is valid JSON or not
	if !json.Valid([]byte(input)) {
		log.Fatalln("false")
	} else {
		log.Println("json.Valid says valid json")
	}

	// Declare a dog struct
	var dog Dog

	// Unmarshall the input json to dog go data structure
	if err := json.Unmarshal([]byte(input), &dog); err != nil {
		log.Fatalf("Unable to unmarshall the json!!! \n Error: %v", err)
	}

	// Validate the json using validator
	if err := validator.New().Struct(dog); err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Print the dog data structure
	fmt.Println(dog.NAME, "\n", dog.Breed, "\n", dog.Age, "\n", dog.FavoriteTreats)

	// Read json data from a file
	data, err := os.ReadFile("./inputData.json")
	if err != nil {
		log.Fatalf("Error: Unable to read file\n %v", err)
	}

	// Declare a hooman
	var hooman Hooman

	// Unmarshall the input json read from the file to hooman object
	if err := json.Unmarshal(data, &hooman); err != nil {
		log.Fatalf("Error: Unable to unmarshall data\n %v", err)
	}

	// Print the hooman object
	fmt.Println(hooman)

	// Read json data from a file
	// babyData, err := os.ReadFile("./babyData.json")
	// if err != nil {
	// 	log.Fatalf("Error: Unable to read file\n %v", err)
	// }

	var babies []Baby
	// if err := json.Unmarshal(babyData, &babies); err != nil {
	// 	log.Fatalf("Error: Unable to unmarshall data\n %v", err)
	// }
	// fmt.Println(babies)

	babyFile, err := os.Open("./babyData.json")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	if err := json.NewDecoder(babyFile).Decode(&babies); err != nil {
		log.Fatalf("Error: %v", err)
	}
	litter.Dump(babies)
}
