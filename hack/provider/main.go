package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var checksumMap = map[string]string{
	"##CHECKSUM_LINUX_AMD64##":   "devpod-provider-sprites-linux-amd64",
	"##CHECKSUM_LINUX_ARM64##":   "devpod-provider-sprites-linux-arm64",
	"##CHECKSUM_DARWIN_AMD64##":  "devpod-provider-sprites-darwin-amd64",
	"##CHECKSUM_DARWIN_ARM64##":  "devpod-provider-sprites-darwin-arm64",
	"##CHECKSUM_WINDOWS_AMD64##": "devpod-provider-sprites-windows-amd64.exe",
}

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "usage: %s <version> <release-dir> <template>\n", os.Args[0])
		os.Exit(1)
	}

	version := os.Args[1]
	releaseDir := os.Args[2]
	templatePath := os.Args[3]

	content, err := os.ReadFile(templatePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read template: %v\n", err)
		os.Exit(1)
	}

	result := strings.ReplaceAll(string(content), "##VERSION##", version)

	for placeholder, binary := range checksumMap {
		data, err := os.ReadFile(filepath.Join(releaseDir, binary))
		if err != nil {
			fmt.Fprintf(os.Stderr, "read binary %s: %v\n", binary, err)
			os.Exit(1)
		}
		checksum := fmt.Sprintf("%x", sha256.Sum256(data))
		result = strings.ReplaceAll(result, placeholder, checksum)
	}

	outPath := filepath.Join(releaseDir, "provider.yaml")
	if err := os.WriteFile(outPath, []byte(result), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "write provider.yaml: %v\n", err)
		os.Exit(1)
	}
}
