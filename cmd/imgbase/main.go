package main

import (
	"fmt"
	"os"

	"github.com/mininit/imgbase/pkg/utils"
)

func main() {
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
