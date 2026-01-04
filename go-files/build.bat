@echo off
REM Build Go project and generate PE32 Windows executable


REM Ensure build output directory exists and is empty
if exist build rmdir /s /q build
mkdir build

REM Build the Go project for Windows 32-bit (PE32)
set GOOS=windows
set GOARCH=386
go build -o build\generator.exe generator.go

if %ERRORLEVEL% neq 0 (
    echo Build failed.
    exit /b %ERRORLEVEL%
) else (
    echo Build succeeded. Output: build\generator.exe
)
