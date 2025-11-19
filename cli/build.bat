@echo off
REM Define output names for different platforms
set LINUX_OUTPUT=remote-cli-linux
set MAC_OUTPUT=remote-cli-mac
set WINDOWS_OUTPUT=remote-cli.exe

REM Build for Linux
set GOOS=linux
set GOARCH=amd64
go build -o %LINUX_OUTPUT% cmd/main.go

REM Build for macOS
set GOOS=darwin
set GOARCH=arm64
go build -o %MAC_OUTPUT% cmd/main.go

REM Build for Windows
set GOOS=windows
set GOARCH=amd64
go build -o %WINDOWS_OUTPUT% cmd/main.go

REM Check for build success
if %ERRORLEVEL% equ 0 (
    echo Build successful!
    echo Linux binary: %LINUX_OUTPUT%
    echo macOS binary: %MAC_OUTPUT%
    echo Windows binary: %WINDOWS_OUTPUT%
) else (
    echo Build failed. Please check the error messages.
)
