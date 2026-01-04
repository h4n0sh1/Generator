# Generator
A project to hold multiple populating "engines"

# File Generator : go-files

A Go script that generates multiple files with random content and random sizes.

## Usage

```bash
go run generator.go <number_of_files> <max_size_kb>
```

### Parameters

- `number_of_files`: Number of files to generate (must be a positive integer)
- `max_size_kb`: Maximum size of each file in KB (minimum is 1KB)

Each generated file will have a random size between 1KB and the specified maximum size.

## Examples

Generate 5 files with sizes between 1KB and 10KB:
```bash
go run generator.go 5 10
```

Generate 100 files with sizes between 1KB and 50KB:
```bash
go run generator.go 100 50
```

## Output

The script creates files named `file_1.txt`, `file_2.txt`, etc., in the current directory. Each file is populated with random alphanumeric characters until the desired size is reached.
