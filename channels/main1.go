package main

import (
	"fmt"
	"sync"
)

func main1() {
	ch := make(chan string, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		wg.Done()
		msg1 := "Hi, Welcome to "
		msg2 := "MKCL"
		fmt.Printf("Message sent: %s %s\n", msg1, msg2)
		ch <- msg1
		ch <- msg2
	}()

	go func() {
		//time.Sleep(time.Second * 1)
		wg.Done()
		fmt.Printf("Message received: %s %s\n", <-ch, <-ch)
	}()

	wg.Wait()
}
