package main

import (
	"fmt"
	"log"
	"os"
	"runtime/metrics"
)

// func main() {
// 	s := []metrics.Sample{{Name: "/gc/stack/starting-size:bytes"}, {Value: metrics.Value{}}}
// 	metrics.Read(s)
// 	fmt.Printf("%d\n", s[1].Value)
// }

func main() {
	descriptions := metrics.All()

	file, err := os.Create("description.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, d := range descriptions {
		_, err := file.WriteString(fmt.Sprintf("Name: %s\nDescription: %s\nKind: %v\n\n", d.Name, d.Description, d.Kind))
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	fmt.Println("Descriptions have been written to description.txt")
}
