package main

import (
	"bytes"
	"fmt"
	"math/rand/v2"
	"strings"
)

// TxtGenerator generates plain text files
type TxtGenerator struct{}

func (g *TxtGenerator) Extension() string {
	return "txt"
}

func (g *TxtGenerator) Generate(sizeBytes int) ([]byte, error) {
	var buf bytes.Buffer
	for buf.Len() < sizeBytes {
		buf.WriteString(randomParagraph())
		buf.WriteString("\n\n")
	}
	return buf.Bytes()[:sizeBytes], nil
}

// CsvGenerator generates CSV files
type CsvGenerator struct{}

func (g *CsvGenerator) Extension() string {
	return "csv"
}

func (g *CsvGenerator) Generate(sizeBytes int) ([]byte, error) {
	var buf bytes.Buffer
	// Write header
	buf.WriteString("id,name,email,department,salary\n")

	id := 1
	for buf.Len() < sizeBytes {
		name := randomWord() + " " + randomWord()
		email := randomWord() + "@" + randomWord() + ".com"
		dept := randomWord()
		salary := 30000 + rand.IntN(70000)
		buf.WriteString(fmt.Sprintf("%d,%s,%s,%s,%d\n", id, name, email, dept, salary))
		id++
	}
	return buf.Bytes()[:sizeBytes], nil
}

// JsonGenerator generates JSON files
type JsonGenerator struct{}

func (g *JsonGenerator) Extension() string {
	return "json"
}

func (g *JsonGenerator) Generate(sizeBytes int) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[\n")

	first := true
	for buf.Len() < sizeBytes-10 {
		if !first {
			buf.WriteString(",\n")
		}
		first = false

		buf.WriteString("  {\n")
		buf.WriteString(fmt.Sprintf("    \"id\": %d,\n", rand.IntN(100000)))
		buf.WriteString(fmt.Sprintf("    \"name\": \"%s\",\n", randomWord()+" "+randomWord()))
		buf.WriteString(fmt.Sprintf("    \"email\": \"%s@%s.com\",\n", randomWord(), randomWord()))
		buf.WriteString(fmt.Sprintf("    \"active\": %t,\n", rand.IntN(2) == 1))
		buf.WriteString(fmt.Sprintf("    \"score\": %d,\n", rand.IntN(100)))
		buf.WriteString(fmt.Sprintf("    \"description\": \"%s\"\n", randomSentence()))
		buf.WriteString("  }")
	}

	buf.WriteString("\n]")
	return buf.Bytes(), nil
}

// XmlGenerator generates XML files
type XmlGenerator struct{}

func (g *XmlGenerator) Extension() string {
	return "xml"
}

func (g *XmlGenerator) Generate(sizeBytes int) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
	buf.WriteString("<records>\n")

	for buf.Len() < sizeBytes-20 {
		buf.WriteString("  <record>\n")
		buf.WriteString(fmt.Sprintf("    <id>%d</id>\n", rand.IntN(100000)))
		buf.WriteString(fmt.Sprintf("    <name>%s</name>\n", randomWord()+" "+randomWord()))
		buf.WriteString(fmt.Sprintf("    <email>%s@%s.com</email>\n", randomWord(), randomWord()))
		buf.WriteString(fmt.Sprintf("    <description>%s</description>\n", randomSentence()))
		buf.WriteString("  </record>\n")
	}

	buf.WriteString("</records>")
	return buf.Bytes(), nil
}

// HtmlGenerator generates HTML files
type HtmlGenerator struct{}

func (g *HtmlGenerator) Extension() string {
	return "html"
}

func (g *HtmlGenerator) Generate(sizeBytes int) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("<!DOCTYPE html>\n<html>\n<head>\n")
	buf.WriteString("  <meta charset=\"UTF-8\">\n")
	buf.WriteString(fmt.Sprintf("  <title>%s</title>\n", randomWord()))
	buf.WriteString("</head>\n<body>\n")

	for buf.Len() < sizeBytes-20 {
		buf.WriteString(fmt.Sprintf("  <h2>%s</h2>\n", randomSentence()))
		buf.WriteString(fmt.Sprintf("  <p>%s</p>\n", randomParagraph()))
		buf.WriteString("  <ul>\n")
		for i := 0; i < 3+rand.IntN(5); i++ {
			buf.WriteString(fmt.Sprintf("    <li>%s</li>\n", randomSentence()))
		}
		buf.WriteString("  </ul>\n")
	}

	buf.WriteString("</body>\n</html>")
	return buf.Bytes(), nil
}

// MarkdownGenerator generates Markdown files
type MarkdownGenerator struct{}

func (g *MarkdownGenerator) Extension() string {
	return "md"
}

func (g *MarkdownGenerator) Generate(sizeBytes int) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("# %s\n\n", strings.Title(randomWord()+" "+randomWord())))

	for buf.Len() < sizeBytes {
		buf.WriteString(fmt.Sprintf("## %s\n\n", randomSentence()))
		buf.WriteString(randomParagraph() + "\n\n")

		// Add a list
		buf.WriteString("### Key Points\n\n")
		for i := 0; i < 3+rand.IntN(4); i++ {
			buf.WriteString(fmt.Sprintf("- %s\n", randomSentence()))
		}
		buf.WriteString("\n")

		// Add a code block
		buf.WriteString("```\n")
		buf.WriteString(randomSentence() + "\n")
		buf.WriteString("```\n\n")
	}

	return buf.Bytes()[:sizeBytes], nil
}

// LogGenerator generates log files
type LogGenerator struct{}

func (g *LogGenerator) Extension() string {
	return "log"
}

func (g *LogGenerator) Generate(sizeBytes int) ([]byte, error) {
	levels := []string{"DEBUG", "INFO", "WARN", "ERROR"}
	var buf bytes.Buffer

	timestamp := 1704067200 // 2024-01-01 00:00:00
	for buf.Len() < sizeBytes {
		level := levels[rand.IntN(len(levels))]
		ts := timestamp + rand.IntN(86400)
		buf.WriteString(fmt.Sprintf("[%d] [%s] %s: %s\n", ts, level, randomWord(), randomSentence()))
		timestamp += rand.IntN(60)
	}

	return buf.Bytes()[:sizeBytes], nil
}
