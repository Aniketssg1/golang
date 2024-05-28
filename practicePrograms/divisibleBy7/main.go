package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var numbers []string
	var numbersBytesBuffer bytes.Buffer
	var numbersStringBuilder strings.Builder
	for i := 2000; i <= 3200; i++ {
		if i%7 == 0 && i%5 != 0 {
			numbers = append(numbers, strconv.Itoa(i))
			numbersBytesBuffer.WriteString(strconv.Itoa(i) + " ")
			numbersStringBuilder.WriteString(strconv.Itoa(i) + " ")
		}
	}
	// numbersString := strings.Join(numbers, ", ")
	// fmt.Println(numbersString)
	//fmt.Println(numbersBytesBuffer.String())
	fmt.Println(numbersStringBuilder.String())
}
