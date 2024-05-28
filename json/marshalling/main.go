package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

type Hooman struct {
	Name    string   `json:"name" validate:"required"`
	Age     int      `json:"age" validate:"required,min=18,max=99,number"`
	Email   string   `json:"email" validate:"required,email"`
	Phone   string   `json:"phone" validate:"required"`
	Hobbies []string `json:"hobbies,omitempty"`
}

type (
	CustomTime struct {
		time.Time
	}
	Baby struct {
		BirthDate CustomTime `json:"birth_date"`
		Name      string     `json:"name"`
		Gender    string     `json:"gender"`
	}
)

func marshal(in any) []byte {
	out, err := json.MarshalIndent(in, "", "	")
	if err != nil {
		log.Fatalf("Error: Unable to marshall!! \n %v", err)
	}
	fmt.Println(json.Valid(out))

	return out
}

func (b Baby) MarshalJSON() ([]byte, error) {
	type Alias Baby
	return json.Marshal(&struct {
		BirthDate string `json:"birth_date"`
		*Alias
	}{
		BirthDate: b.BirthDate.Format("02-01-2006"),
		Alias:     (*Alias)(&b),
	})
}

func main() {
	first := marshal(14)
	second := marshal("Hello world")
	third := marshal([]float32{1.66, 6.86, 10.1})
	fourth := marshal(map[string]int{"num": 15, "other": 17})
	fmt.Printf(
		"first: %s\nsecond: %s\nthird: %s\nfourth: %s\n",
		first,
		second,
		third,
		fourth,
	)

	hooman := Hooman{
		Email: "johnjones@email.com",
		Name:  "John Jones",
		Age:   26,
		Phone: "89910119",
		Hobbies: []string{
			"Swimming",
			"Badminton",
		},
	}

	marshaledHooman := marshal(hooman)

	if err := validator.New().Struct(hooman); err != nil {
		log.Printf("Error: %v\n", err)
	}

	fmt.Println(string(marshaledHooman))

	baby := Baby{
		Name:   "johnny",
		Gender: "male",
		BirthDate: CustomTime{
			time.Date(2023, 1, 1, 12, 0, 0, 0, time.Now().Location()),
		},
	}
	jsonedBaby := marshal(baby)

	fmt.Println(string(jsonedBaby))
}
