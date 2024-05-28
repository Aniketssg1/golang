package main

import "fmt"

func main2() {
	counterChannel := make(chan int)
	done := make(chan struct{})

	for i := 0; i < 10; i++ {
		go func() {
			counterChannel <- 1
		}()
	}

	go func() {
		counter := 0
		for {
			select {
			case <-counterChannel:
				counter++
				fmt.Print(counter, " ")
			case <-done:
				return
			}
		}
	}()

	close(done)
}
