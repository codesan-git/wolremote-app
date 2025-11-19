#!/bin/bash

# Define output names for different platforms
LINUX_OUTPUT="remote-cli-linux"
MAC_OUTPUT="remote-cli-mac"
WINDOWS_OUTPUT="remote-cli.exe"

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o $LINUX_OUTPUT cmd/main.go

# Build for macOS (Apple Silicon processor)
GOOS=darwin GOARCH=arm64 go build -o $MAC_OUTPUT cmd/main.go

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o $WINDOWS_OUTPUT cmd/main.go

# Check for build success
if [ $? -eq 0 ]; then
    echo "Build successful!"
    echo "Linux binary: $LINUX_OUTPUT"
    echo "macOS binary: $MAC_OUTPUT"
    echo "Windows binary: $WINDOWS_OUTPUT"
else
    echo "Build failed. Please check the error messages."
fi
