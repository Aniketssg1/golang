package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime"
)

type car struct {
	price int
	color string
}

// Function to convert bytes to MiB
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// Function to print memory usage into the buffer
func printMemUsage(buffer *bytes.Buffer) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Fprintf(buffer, "Alloc = %vMiB ", bToMb(m.Alloc))
	fmt.Fprintf(buffer, "TotalAlloc = %vMiB ", bToMb(m.TotalAlloc))
	fmt.Fprintf(buffer, "Sys = %vMiB ", bToMb(m.Sys))
	fmt.Fprintf(buffer, "NumGC = %v ", m.NumGC)
}

func logMemoryProfile(file *os.File, buffer *bytes.Buffer, message string) {
	buffer.WriteString("\n" + message + "\n")
	printMemUsage(buffer)
	file.Write(buffer.Bytes())
	buffer.Reset()
}

func main() {
	memoryProfileFile, err := os.Create("memoryProfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer memoryProfileFile.Close()

	buffer := bytes.Buffer{}
	logMemoryProfile(memoryProfileFile, &buffer, "Memory usage initially")

	car := car{
		price: 1000,
		color: "White",
	}
	logMemoryProfile(memoryProfileFile, &buffer, "Memory usage after allocating a car struct")

	fmt.Println(car)
	logMemoryProfile(memoryProfileFile, &buffer, "Memory usage after printing the car struct")

	runtime.GC()
	logMemoryProfile(memoryProfileFile, &buffer, "Memory usage after garbage collection")
}
