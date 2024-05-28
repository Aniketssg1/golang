package main

import (
	"fmt"
	"unsafe"
)

func main1() {
	i := uint32(1)
	c := (*byte)(unsafe.Pointer(&i))

	if *c != 0 {
		fmt.Println("Little endian")
	} else {
		fmt.Println("Big endian")
	}
}
