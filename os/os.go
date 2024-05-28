package main

import (
	"fmt"
	"os"
)

func main() {
	id2 := os.Geteuid()
	fmt.Printf("id: %T, %v", id2, id2)
}
