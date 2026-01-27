package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

const (
	minSizeKB = 1
	charset   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// FileGenerator defines the interface for generating different file types
type FileGenerator interface {
	Generate(sizeBytes int) ([]byte, error)
	Extension() string
}

// SupportedExtensions returns list of all supported file extensions
func SupportedExtensions() []string {
	return []string{"txt", "csv", "json", "xml", "html", "md", "log", "pdf", "docx", "xlsx", "png"}
}

// NewGenerator returns the appropriate generator for the given extension
func NewGenerator(ext string) FileGenerator {
	switch strings.ToLower(ext) {
	case "txt":
		return &TxtGenerator{}
	case "csv":
		return &CsvGenerator{}
	case "json":
		return &JsonGenerator{}
	case "xml":
		return &XmlGenerator{}
	case "html":
		return &HtmlGenerator{}
	case "md":
		return &MarkdownGenerator{}
	case "log":
		return &LogGenerator{}
	case "pdf":
		return &PdfGenerator{}
	case "docx":
		return &DocxGenerator{}
	case "xlsx":
		return &XlsxGenerator{}
	case "png":
		return &PngGenerator{}
	default:
		return &TxtGenerator{}
	}
}

// randomString generates a random string of the specified length
func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.IntN(len(charset))]
	}
	return string(b)
}

// randomWord generates a random word of 3-10 characters
func randomWord() string {
	return randomString(3 + rand.IntN(8))
}

// randomSentence generates a random sentence with 5-15 words
func randomSentence() string {
	wordCount := 5 + rand.IntN(11)
	words := make([]string, wordCount)
	for i := range words {
		words[i] = randomWord()
	}
	return strings.Join(words, " ") + "."
}

// randomParagraph generates a random paragraph with 3-7 sentences
func randomParagraph() string {
	sentenceCount := 3 + rand.IntN(5)
	sentences := make([]string, sentenceCount)
	for i := range sentences {
		sentences[i] = randomSentence()
	}
	return strings.Join(sentences, " ")
}

func main() {
	// Check command line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: generator <number_of_files> <max_size_kb> [extensions]")
		fmt.Println("  number_of_files: Total number of files to generate")
		fmt.Println("  max_size_kb: Maximum size of each file in KB (minimum is 1KB)")
		fmt.Println("  extensions: Comma-separated list of extensions (optional)")
		fmt.Printf("  Supported extensions: %s\n", strings.Join(SupportedExtensions(), ", "))
		fmt.Println("\nExample: generator 100 100 txt,csv,json")
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

	// Parse extensions (use all if not specified)
	extensions := SupportedExtensions()
	if len(os.Args) >= 4 {
		extensions = strings.Split(os.Args[3], ",")
		for i := range extensions {
			extensions[i] = strings.TrimSpace(extensions[i])
		}
	}

	// Generate files with random extensions
	fmt.Printf("Generating %d files with random sizes between %d KB and %d KB...\n",
		numFiles, minSizeKB, maxSizeKB)

	for i := 1; i <= numFiles; i++ {
		// Pick a random extension
		ext := extensions[rand.IntN(len(extensions))]
		generator := NewGenerator(ext)

		// Generate random size between minSizeKB and maxSizeKB
		var fileSizeKB int
		if maxSizeKB == minSizeKB {
			fileSizeKB = minSizeKB
		} else {
			fileSizeKB = minSizeKB + rand.IntN(maxSizeKB-minSizeKB+1)
		}
		fileSizeBytes := fileSizeKB * 1024

		// Create filename
		filename := fmt.Sprintf("file_%d.%s", i, generator.Extension())

		// Generate content using the appropriate generator
		content, err := generator.Generate(fileSizeBytes)
		if err != nil {
			fmt.Printf("Error generating content for %s: %v\n", filename, err)
			continue
		}

		// Write file
		err = os.WriteFile(filename, content, 0644)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("Created %s (size: %d KB)\n", filename, len(content)/1024)
	}

	fmt.Println("\nFile generation completed!")
}
