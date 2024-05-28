package main

import "fmt"

func multiplyByThree(arr []int, ch chan int) {
	for _, e := range arr {
		ch <- e * 3
	}
}

func main5() {
	fmt.Println("Just about to start exectuing a goroutine")
	arr := []int{1, 2, 3, 4, 5}
	ch := make(chan int, len(arr))
	go multiplyByThree(arr, ch)
	for i := 0; i < len(arr); i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println("\nFinished executing the goroutine")
}
