package main

import (
	"fmt"
	"packages/package2"

	tinytime "github.com/wagslane/go-tinytime"
)

func main() {
	fmt.Println("Hi I am from package 1")
	simpleInterest := package2.CalculateSimpleInterest(10000, 10, 10)
	fmt.Printf("Simple Interest: %g\n", simpleInterest)

	car := package2.Car{
		Name:  "Porche",
		Model: "911 GT",
	}
	fmt.Print(car, "\n")

	tt := tinytime.New(1585750374)

	fmt.Print(tt)

}
