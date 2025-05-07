package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: imgbase <file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	sniffLen := 512
	if len(data) < sniffLen {
		sniffLen = len(data)
	}

	contentType := http.DetectContentType(data[:sniffLen])

	if strings.Contains(string(data[:sniffLen]), "<svg") {
		contentType = "image/svg+xml"
	}

	// Encode to Base64
	encoded := base64.StdEncoding.EncodeToString(data)

	fmt.Printf("data:%s;base64,%s\n", contentType, encoded)

}
