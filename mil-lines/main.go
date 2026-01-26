package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a new file
	file, err := os.Create("file.go")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create a buffered writer
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Write 1 million lines
	totalLines := 1_000_000
	for i := 1; i <= totalLines; i++ {
		// Write line with "txt" at the end
		_, err := writer.WriteString(fmt.Sprintf("Line %d with txt\n", i))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		// Progress indicator (optional)
		if i%100000 == 0 {
			fmt.Printf("Written %d lines...\n", i)
		}
	}

	fmt.Printf("Successfully wrote %d lines to file.go\n", totalLines)
}