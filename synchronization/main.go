package main

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	var mu sync.Mutex
	var counter int
	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			fmt.Print(counter, " ")
			mu.Unlock()
		}()
	}

	wg.Wait()
}

func main() {
	counterChannel := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
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
			case <-time.After(time.Second * 1):
				close(done)
				return
			}
		}
	}()

	wg.Wait()
}
