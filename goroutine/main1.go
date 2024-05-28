package main

import (
	"fmt"
	"time"
)

func display(msg string) {
	for i := 0; i < 6; i++ {
		fmt.Print(msg, "\n")
	}
}

func main1() {
	go display("Welcome")
	time.Sleep(time.Millisecond * 1000)
	display("Aniket")
}
