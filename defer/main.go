package main

import "fmt"

func main() {
	defer fmt.Print("World")
	fmt.Print("Hello ")
}
