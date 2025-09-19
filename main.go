package main

import (
	"fmt"
	"io"
	"os"

	"github.com/mininit/imgbase/utils"
)

var (
	Version    = ""
	validFlags = map[string]bool{
		"-v": true,
		"--version": true,
	}
)

func main() {
	// Handle flags
	if len(os.Args) > 1 && os.Args[1] != "-" && os.Args[1][0] == '-' {
		if !validFlags[os.Args[1]] {
			fmt.Fprintf(os.Stderr, "imgbase: illegal option %s\n", os.Args[1])
			usage()
		}
		if os.Args[1] == "-v" || os.Args[1] == "--version" {
			if Version == "" {
				Version = "unknown (built from source)"
			}
			fmt.Println("imgbase version", Version)
			return
		}
	}

	// If no args, read stdin
	if len(os.Args) < 2 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading stdin: ", err)
			os.Exit(1)
		}
		out(data)
		return
	}

	// Process all file arguments
	for _, filename := range os.Args[1:] {
		var data []byte
		var err error

		if filename == "-" {
			data, err = io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error reading stdin: ", err)
				os.Exit(1)
			}
		} else {
			data, err = os.ReadFile(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "imgbase: %s: No such file or directory\n", filename)
				continue
			}
		}

		out(data)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: imgbase [-v | --version] [files ...]")
	os.Exit(1)
}

func out(data []byte) {
	dataURI, err := utils.EncodeImageBytesToBase64DataURI(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(dataURI)
}
