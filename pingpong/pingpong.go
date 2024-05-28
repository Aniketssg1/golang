package main

import (
	"fmt"
	"time"
)

func pingpong(numOfPings int) {
	pings := make(chan struct{})
	pongs := make(chan struct{})

	go pinger(pings, pongs, numOfPings)
	go ponger(pings, pongs)
}

func pinger(pings, pongs chan struct{}, numOfPings int) {
	go func() {
		gapBetween := time.Millisecond * 50
		for i := 1; i <= numOfPings; i++ {
			fmt.Printf("sending ping %v\n", i)
			pings <- struct{}{}
			time.Sleep(gapBetween)
			gapBetween *= 2
		}
		close(pings)
	}()
	i := 0
	for range pongs {
		fmt.Println("got pong", i)
		i++
	}
	fmt.Println("pongs done")
}

func ponger(pings, pongs chan struct{}) {
	i := 0
	for range pings {
		fmt.Printf("got ping %v, sending pong %v\n", i, i)
		pongs <- struct{}{}
		i++
	}
	fmt.Println("pings done")
	close(pongs)
}

func test(numOfPings int) {
	fmt.Print("Game Starting============================================")
	pingpong(numOfPings)
	fmt.Print("Game Over================================================")
}

func main() {
	test(4)
}
