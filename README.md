# imgbase

**imgbase** is a simple command-line tool that converts image files to Base64-encoded data URIs. It's useful for embedding small images directly in HTML, CSS, or JSON without external file references.

## Features

- Converts images (e.g. PNG, JPG, GIF, SVG) to Base64 data URLs
- Outputs directly to terminal or file

## Installation

You can install using Homebrew:

```bash
brew install mininit/tap/imgbase
```


Or build from source using Go:

```bash
git clone https://github.com/mininit/imgbase.git
cd imgbase
go build -o imgbase
```

## Usage
`imgbase path/to/image.png`

Example Output:
`data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAA...`

To save the output to a file:
`imgbase logo.png > logo.txt`