package main

import (
	"fmt"
	"unsafe"
)

// Function to show bytes in memory, from location start to start+n
func showMemRep(start unsafe.Pointer, n int) {
	bytes := *(*[]byte)(unsafe.Pointer(&struct {
		ptr unsafe.Pointer
		len int
	}{start, n}))

	for _, b := range bytes {
		fmt.Printf("%.2x ", b)
	}
	fmt.Println()
}

// Main function to call above function for 0x01234567
func main() {
	i := 0x01234567
	showMemRep(unsafe.Pointer(&i), int(unsafe.Sizeof(i)))
	main1()
	fmt.Scanln() // Wait for user input to exit
}
