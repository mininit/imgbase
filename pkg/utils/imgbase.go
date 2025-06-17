package utils

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// EncodeImageToBase64DataURL reads the given file and returns a data URL
func EncodeImageToBase64DataURL(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	sniffLen := 512
	if len(data) < sniffLen {
		sniffLen = len(data)
	}

	contentType := http.DetectContentType(data[:sniffLen])

	// Special case for SVG
	if strings.Contains(string(data[:sniffLen]), "<svg") {
		contentType = "image/svg+xml"
	}

	encoded := base64.StdEncoding.EncodeToString(data)
	return fmt.Sprintf("data:%s;base64,%s", contentType, encoded), nil
}