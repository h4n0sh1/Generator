# Generator

A multi-format file generator with support for text, documents, and images.

## Quick Start

Pre-built binaries are available in `dist/`:
- `generator.exe` - Windows 11
- `generator` - macOS (Intel & Apple Silicon)

## Usage

```bash
generator <number_of_files> <max_size_kb> [extensions]
```

### Parameters

- `number_of_files`: Total number of files to generate
- `max_size_kb`: Maximum file size in KB (min: 1KB)
- `extensions`: Comma-separated list (optional, defaults to all)

### Supported Formats

| Text | Documents | Binary |
|------|-----------|--------|
| txt, csv, json, xml, html, md, log | pdf, docx, xlsx | png (pixel art animals!) |

## Examples

```bash
# Generate 100 files with random extensions, up to 100KB each
generator 100 100

# Generate 50 files, only txt/csv/json formats
generator 50 50 txt,csv,json

# Generate 20 document/image files
generator 20 200 pdf,docx,png
```

## Build

```bash
make          # Build for Windows & macOS
make windows  # Windows only
make mac      # macOS only
make clean    # Clean build artifacts
```
