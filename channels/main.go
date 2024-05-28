package main

import (
	"fmt"
)

func main() {
	bufferedChannel := make(chan string, 2)
	bufferedChannel <- "Welcome"
	bufferedChannel <- "Aniket"
	//bufferedChannel <- "Ghage"
	fmt.Print(<-bufferedChannel, " ", <-bufferedChannel, "\n")

	//main1()

	//main2()

	//main3(&testing.T{})

	//main4()
}
