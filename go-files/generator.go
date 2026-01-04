package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

const (
	minSizeKB = 1
	charset   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func main() {
	// Check command line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run generator.go <number_of_files> <max_size_kb>")
		fmt.Println("  number_of_files: Number of files to generate")
		fmt.Println("  max_size_kb: Maximum size of each file in KB (minimum is 1KB)")
		os.Exit(1)
	}

	// Parse arguments
	numFiles, err := strconv.Atoi(os.Args[1])
	if err != nil || numFiles <= 0 {
		fmt.Printf("Error: Invalid number of files '%s'. Must be a positive integer.\n", os.Args[1])
		os.Exit(1)
	}

	maxSizeKB, err := strconv.Atoi(os.Args[2])
	if err != nil || maxSizeKB < minSizeKB {
		fmt.Printf("Error: Invalid max size '%s'. Must be at least %d KB.\n", os.Args[2], minSizeKB)
		os.Exit(1)
	}

	// Generate files
	fmt.Printf("Generating %d files with random sizes between %d KB and %d KB...\n", numFiles, minSizeKB, maxSizeKB)

	for i := 1; i <= numFiles; i++ {
		// Generate random size between minSizeKB and maxSizeKB
		var fileSizeKB int
		if maxSizeKB == minSizeKB {
			fileSizeKB = minSizeKB
		} else {
			fileSizeKB = minSizeKB + rand.IntN(maxSizeKB-minSizeKB+1)
		}
		fileSizeBytes := fileSizeKB * 1024

		// Create filename
		filename := fmt.Sprintf("file_%d.txt", i)

		// Generate random content
		content := make([]byte, fileSizeBytes)
		for j := 0; j < fileSizeBytes; j++ {
			content[j] = charset[rand.IntN(len(charset))]
		}

		// Write file
		err := os.WriteFile(filename, content, 0644)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("Created %s (size: %d KB)\n", filename, fileSizeKB)
	}

	fmt.Println("File generation completed!")
}
