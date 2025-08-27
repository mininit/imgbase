package utils

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

// EncodeImageBytesToBase64DataURI encodes raw image bytes to a data URI
func EncodeImageBytesToBase64DataURI(data []byte) (string, error) {
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
