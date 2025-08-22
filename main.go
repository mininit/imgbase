package main

import (
	"fmt"
	"os"

	"github.com/mininit/imgbase/utils"
)

// Version is set at build time with -ldflags "-X main.Version=..."
var Version = "dev"

func main() {
	// Only handle --version as the sole argument
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Println("imgbase", Version)
		return
	}

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: imgbase <file>")
		os.Exit(1)
	}

	dataURL, err := utils.EncodeImageToBase64DataURL(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(dataURL)
}
