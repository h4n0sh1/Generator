.PHONY: all clean windows mac build-dir

# Default target builds for both platforms
all: windows mac

# Create build directory
build-dir:
	@mkdir -p build

# Build for Windows 11 (64-bit)
windows: build-dir
	@echo "Building for Windows 11 (64-bit)..."
	cd go-files && GOOS=windows GOARCH=amd64 go build -o ../build/generator.exe generator.go
	@echo "Build succeeded. Output: build/generator.exe"

# Build for macOS (Universal binary for Intel and Apple Silicon)
mac: build-dir
	@echo "Building for macOS (Universal binary)..."
	cd go-files && GOOS=darwin GOARCH=amd64 go build -o ../build/generator-amd64 generator.go
	cd go-files && GOOS=darwin GOARCH=arm64 go build -o ../build/generator-arm64 generator.go
	lipo -create -output build/generator build/generator-amd64 build/generator-arm64
	@rm build/generator-amd64 build/generator-arm64
	@echo "Build succeeded. Output: build/generator"

# Build for macOS Intel only (if universal binary is not needed)
mac-intel: build-dir
	@echo "Building for macOS (Intel only)..."
	cd go-files && GOOS=darwin GOARCH=amd64 go build -o ../build/generator generator.go
	@echo "Build succeeded. Output: build/generator"

# Build for macOS Apple Silicon only (if universal binary is not needed)
mac-arm: build-dir
	@echo "Building for macOS (Apple Silicon only)..."
	cd go-files && GOOS=darwin GOARCH=arm64 go build -o ../build/generator generator.go
	@echo "Build succeeded. Output: build/generator"

# Clean build artifacts
clean:
	@echo "Cleaning build directory..."
	@rm -rf build
	@echo "Clean complete."
