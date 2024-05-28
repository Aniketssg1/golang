package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// file, err := os.Open("test1.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	data := make([]byte, 100)
	// 	count, err := file.Read(data)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	} else {
	// 		fmt.Printf("read: %d bytes; content: %s\n", count, string(data))
	// 	}
	// }
	// er := os.Remove("test.txt")
	// if er != nil {
	// 	log.Fatal(er)
	// }

	if _, err1 := os.Stat("test.txt"); err1 != nil {
		log.Fatal(err1)
	} else {
		fmt.Println("file exists")
	}

	if _, err2 := os.Stat("test1.txt"); os.IsExist(err2) {
		log.Fatal(err2)
	}
}
