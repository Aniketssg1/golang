package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

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

// Function to log memory usage into the file
func logMemUsage(file *os.File, buffer *bytes.Buffer, message string) {
	buffer.WriteString("\n" + message + "\n")
	printMemUsage(buffer)
	file.Write(buffer.Bytes())
	buffer.Reset()
}

func main() {
	// Create a file to log memory usage
	memoryUsageFile, err := os.Create("memoryUsageFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer memoryUsageFile.Close()

	buffer := bytes.Buffer{}

	// Log initial memory usage
	logMemUsage(memoryUsageFile, &buffer, "Memory usage initially")

	// Allocate memory in a loop
	for i := 0; i < 10; i++ {
		s := make([]byte, 1<<20) // 1 MiB
		_ = s
		time.Sleep(time.Second)
	}

	// Log memory usage after allocation
	logMemUsage(memoryUsageFile, &buffer, "Memory usage after large slice allocation")

	// Free up memory and force garbage collection
	runtime.GC()

	// Log memory usage after garbage collection
	logMemUsage(memoryUsageFile, &buffer, "Memory usage after garbage collection")
}
