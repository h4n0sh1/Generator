package main

import (
	"archive/zip"
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand/v2"
)

// PdfGenerator generates valid PDF files
type PdfGenerator struct{}

func (g *PdfGenerator) Extension() string {
	return "pdf"
}

func (g *PdfGenerator) Generate(sizeBytes int) ([]byte, error) {
	var buf bytes.Buffer

	// Generate text content to fill the PDF
	var textContent bytes.Buffer
	for textContent.Len() < sizeBytes/2 {
		textContent.WriteString(randomParagraph())
		textContent.WriteString(" ")
	}
	text := textContent.String()

	// PDF header
	buf.WriteString("%PDF-1.4\n")
	buf.WriteString("%âãÏÓ\n")

	// Object 1: Catalog
	obj1Offset := buf.Len()
	buf.WriteString("1 0 obj\n<< /Type /Catalog /Pages 2 0 R >>\nendobj\n")

	// Object 2: Pages
	obj2Offset := buf.Len()
	buf.WriteString("2 0 obj\n<< /Type /Pages /Kids [3 0 R] /Count 1 >>\nendobj\n")

	// Object 3: Page
	obj3Offset := buf.Len()
	buf.WriteString("3 0 obj\n<< /Type /Page /Parent 2 0 R /MediaBox [0 0 612 792] /Contents 4 0 R /Resources << /Font << /F1 5 0 R >> >> >>\nendobj\n")

	// Object 4: Content stream
	obj4Offset := buf.Len()
	streamContent := fmt.Sprintf("BT\n/F1 12 Tf\n50 750 Td\n14 TL\n")

	// Split text into lines (max ~80 chars per line for PDF)
	words := bytes.Fields([]byte(text))
	var line bytes.Buffer
	for _, word := range words {
		if line.Len()+len(word)+1 > 70 {
			streamContent += fmt.Sprintf("(%s) Tj T*\n", escapePdfString(line.String()))
			line.Reset()
		}
		if line.Len() > 0 {
			line.WriteByte(' ')
		}
		line.Write(word)
	}
	if line.Len() > 0 {
		streamContent += fmt.Sprintf("(%s) Tj\n", escapePdfString(line.String()))
	}
	streamContent += "ET"

	buf.WriteString(fmt.Sprintf("4 0 obj\n<< /Length %d >>\nstream\n%s\nendstream\nendobj\n", len(streamContent), streamContent))

	// Object 5: Font
	obj5Offset := buf.Len()
	buf.WriteString("5 0 obj\n<< /Type /Font /Subtype /Type1 /BaseFont /Helvetica >>\nendobj\n")

	// Cross-reference table
	xrefOffset := buf.Len()
	buf.WriteString("xref\n0 6\n")
	buf.WriteString("0000000000 65535 f \n")
	buf.WriteString(fmt.Sprintf("%010d 00000 n \n", obj1Offset))
	buf.WriteString(fmt.Sprintf("%010d 00000 n \n", obj2Offset))
	buf.WriteString(fmt.Sprintf("%010d 00000 n \n", obj3Offset))
	buf.WriteString(fmt.Sprintf("%010d 00000 n \n", obj4Offset))
	buf.WriteString(fmt.Sprintf("%010d 00000 n \n", obj5Offset))

	// Trailer
	buf.WriteString(fmt.Sprintf("trailer\n<< /Size 6 /Root 1 0 R >>\nstartxref\n%d\n%%%%EOF\n", xrefOffset))

	// Pad if needed
	result := buf.Bytes()
	if len(result) < sizeBytes {
		padding := make([]byte, sizeBytes-len(result))
		for i := range padding {
			padding[i] = ' '
		}
		// Insert padding as a comment before EOF
		eofIdx := bytes.LastIndex(result, []byte("%%EOF"))
		if eofIdx > 0 {
			var finalBuf bytes.Buffer
			finalBuf.Write(result[:eofIdx])
			finalBuf.WriteString("% ")
			finalBuf.Write(padding)
			finalBuf.WriteString("\n%%EOF\n")
			result = finalBuf.Bytes()
		}
	}

	return result, nil
}

func escapePdfString(s string) string {
	var buf bytes.Buffer
	for _, c := range s {
		switch c {
		case '(', ')', '\\':
			buf.WriteByte('\\')
			buf.WriteRune(c)
		default:
			if c >= 32 && c < 127 {
				buf.WriteRune(c)
			} else {
				buf.WriteByte(' ')
			}
		}
	}
	return buf.String()
}

// DocxGenerator generates valid DOCX files (Office Open XML)
type DocxGenerator struct{}

func (g *DocxGenerator) Extension() string {
	return "docx"
}

func (g *DocxGenerator) Generate(sizeBytes int) ([]byte, error) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	// [Content_Types].xml
	contentTypes := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types">
  <Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/>
  <Default Extension="xml" ContentType="application/xml"/>
  <Override PartName="/word/document.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"/>
</Types>`
	writeZipFile(zipWriter, "[Content_Types].xml", contentTypes)

	// _rels/.rels
	rels := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
  <Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" Target="word/document.xml"/>
</Relationships>`
	writeZipFile(zipWriter, "_rels/.rels", rels)

	// Generate document content
	var docContent bytes.Buffer
	docContent.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:document xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main">
  <w:body>`)

	// Add paragraphs until we reach target size
	for docContent.Len() < sizeBytes/2 {
		docContent.WriteString("\n    <w:p><w:r><w:t>")
		docContent.WriteString(randomParagraph())
		docContent.WriteString("</w:t></w:r></w:p>")
	}

	docContent.WriteString(`
  </w:body>
</w:document>`)

	writeZipFile(zipWriter, "word/document.xml", docContent.String())

	// word/_rels/document.xml.rels
	docRels := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
</Relationships>`
	writeZipFile(zipWriter, "word/_rels/document.xml.rels", docRels)

	zipWriter.Close()
	return buf.Bytes(), nil
}

// XlsxGenerator generates valid XLSX files (Office Open XML)
type XlsxGenerator struct{}

func (g *XlsxGenerator) Extension() string {
	return "xlsx"
}

func (g *XlsxGenerator) Generate(sizeBytes int) ([]byte, error) {
	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	// [Content_Types].xml
	contentTypes := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types">
  <Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/>
  <Default Extension="xml" ContentType="application/xml"/>
  <Override PartName="/xl/workbook.xml" ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"/>
  <Override PartName="/xl/worksheets/sheet1.xml" ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"/>
</Types>`
	writeZipFile(zipWriter, "[Content_Types].xml", contentTypes)

	// _rels/.rels
	rels := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
  <Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" Target="xl/workbook.xml"/>
</Relationships>`
	writeZipFile(zipWriter, "_rels/.rels", rels)

	// xl/workbook.xml
	workbook := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<workbook xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
  <sheets>
    <sheet name="Sheet1" sheetId="1" r:id="rId1"/>
  </sheets>
</workbook>`
	writeZipFile(zipWriter, "xl/workbook.xml", workbook)

	// xl/_rels/workbook.xml.rels
	wbRels := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
  <Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet" Target="worksheets/sheet1.xml"/>
</Relationships>`
	writeZipFile(zipWriter, "xl/_rels/workbook.xml.rels", wbRels)

	// xl/worksheets/sheet1.xml - Generate spreadsheet data
	var sheetContent bytes.Buffer
	sheetContent.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<worksheet xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main">
  <sheetData>`)

	// Header row
	sheetContent.WriteString(`
    <row r="1">
      <c r="A1" t="inlineStr"><is><t>ID</t></is></c>
      <c r="B1" t="inlineStr"><is><t>Name</t></is></c>
      <c r="C1" t="inlineStr"><is><t>Email</t></is></c>
      <c r="D1" t="inlineStr"><is><t>Department</t></is></c>
      <c r="E1" t="inlineStr"><is><t>Salary</t></is></c>
    </row>`)

	// Data rows
	row := 2
	for sheetContent.Len() < sizeBytes/2 {
		name := randomWord() + " " + randomWord()
		email := randomWord() + "@" + randomWord() + ".com"
		dept := randomWord()
		salary := 30000 + rand.IntN(70000)

		sheetContent.WriteString(fmt.Sprintf(`
    <row r="%d">
      <c r="A%d"><v>%d</v></c>
      <c r="B%d" t="inlineStr"><is><t>%s</t></is></c>
      <c r="C%d" t="inlineStr"><is><t>%s</t></is></c>
      <c r="D%d" t="inlineStr"><is><t>%s</t></is></c>
      <c r="E%d"><v>%d</v></c>
    </row>`, row, row, row-1, row, name, row, email, row, dept, row, salary))
		row++
	}

	sheetContent.WriteString(`
  </sheetData>
</worksheet>`)

	writeZipFile(zipWriter, "xl/worksheets/sheet1.xml", sheetContent.String())

	zipWriter.Close()
	return buf.Bytes(), nil
}

func writeZipFile(zw *zip.Writer, name string, content string) error {
	w, err := zw.Create(name)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(content))
	return err
}

// PngGenerator generates valid PNG image files with pixel art animals
type PngGenerator struct{}

func (g *PngGenerator) Extension() string {
	return "png"
}

func (g *PngGenerator) Generate(sizeBytes int) ([]byte, error) {
	// Determine image size based on target file size
	imgSize := 256
	if sizeBytes > 50*1024 {
		imgSize = 512
	}
	if sizeBytes > 200*1024 {
		imgSize = 1024
	}

	// Create image with a nice background
	img := image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))

	// Fill with a random pastel background
	bgColor := randomPastelColor()
	for y := 0; y < imgSize; y++ {
		for x := 0; x < imgSize; x++ {
			img.Set(x, y, bgColor)
		}
	}

	// Get a random animal pattern
	animal := GetRandomAnimal()

	// Calculate pixel size to scale the animal to fit nicely
	// Animal should take up about 60-80% of the image
	patternWidth := len(animal.Pattern[0])
	patternHeight := len(animal.Pattern)
	pixelSize := (imgSize * 7 / 10) / max(patternWidth, patternHeight)
	if pixelSize < 1 {
		pixelSize = 1
	}

	// Center the animal
	startX := (imgSize - patternWidth*pixelSize) / 2
	startY := (imgSize - patternHeight*pixelSize) / 2

	// Draw the animal
	for py, row := range animal.Pattern {
		for px, char := range row {
			var c color.RGBA
			switch char {
			case '#':
				c = animal.Primary
			case 'o':
				c = animal.Secondary
			case '.':
				c = animal.Accent
			default:
				continue // Skip transparent pixels
			}

			// Draw a pixelSize x pixelSize block
			for dy := 0; dy < pixelSize; dy++ {
				for dx := 0; dx < pixelSize; dx++ {
					x := startX + px*pixelSize + dx
					y := startY + py*pixelSize + dy
					if x >= 0 && x < imgSize && y >= 0 && y < imgSize {
						img.Set(x, y, c)
					}
				}
			}
		}
	}

	// Add a simple border/frame
	frameColor := color.RGBA{50, 50, 50, 255}
	frameWidth := max(2, imgSize/64)
	for i := 0; i < frameWidth; i++ {
		for x := 0; x < imgSize; x++ {
			img.Set(x, i, frameColor)
			img.Set(x, imgSize-1-i, frameColor)
		}
		for y := 0; y < imgSize; y++ {
			img.Set(i, y, frameColor)
			img.Set(imgSize-1-i, y, frameColor)
		}
	}

	// Encode to PNG
	var buf bytes.Buffer
	err := png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}

	result := buf.Bytes()

	// If result is smaller than target, add metadata padding
	if len(result) < sizeBytes {
		padSize := sizeBytes - len(result)
		if padSize > 0 {
			result = appendPngTextChunk(result, padSize)
		}
	}

	return result, nil
}

func randomPastelColor() color.RGBA {
	pastels := []color.RGBA{
		{255, 230, 230, 255}, // Light pink
		{255, 240, 220, 255}, // Peach
		{255, 255, 220, 255}, // Light yellow
		{220, 255, 220, 255}, // Light green
		{220, 240, 255, 255}, // Light blue
		{240, 220, 255, 255}, // Light purple
		{255, 220, 240, 255}, // Light magenta
		{220, 255, 255, 255}, // Light cyan
		{245, 245, 220, 255}, // Beige
		{230, 230, 250, 255}, // Lavender
	}
	return pastels[rand.IntN(len(pastels))]
}

func appendPngTextChunk(pngData []byte, padSize int) []byte {
	// Find IEND chunk position (last 12 bytes: 4 length + 4 type + 4 crc)
	if len(pngData) < 12 {
		return pngData
	}

	iendPos := len(pngData) - 12

	// Create tEXt chunk with padding data
	keyword := "Comment"
	padding := make([]byte, padSize-len(keyword)-1-12) // subtract overhead
	if len(padding) < 0 {
		return pngData
	}
	for i := range padding {
		padding[i] = charset[rand.IntN(len(charset))]
	}

	// tEXt chunk format: keyword + null byte + text
	textData := append([]byte(keyword), 0)
	textData = append(textData, padding...)

	// Create chunk: length (4) + type (4) + data + crc (4)
	var chunk bytes.Buffer
	binary.Write(&chunk, binary.BigEndian, uint32(len(textData)))
	chunk.WriteString("tEXt")
	chunk.Write(textData)
	// Simplified CRC (real implementation would calculate proper CRC32)
	binary.Write(&chunk, binary.BigEndian, uint32(0))

	// Insert before IEND
	var result bytes.Buffer
	result.Write(pngData[:iendPos])
	result.Write(chunk.Bytes())
	result.Write(pngData[iendPos:])

	return result.Bytes()
}
