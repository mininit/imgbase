# imgbase

**imgbase** is a simple command-line tool that converts image files to Base64-encoded data URIs.  
It’s useful for embedding small images directly in HTML, CSS, or JSON without external file references.

## Features

- Converts images (PNG, JPG, GIF, SVG) to Base64 data URLs
- Reads from files or stdin
- Outputs directly to terminal or file

## Installation

### Package Manager

```bash
# macOS or Linux
brew install mininit/tap/imgbase
```

## Usage

`imgbase [file ...]`
- Reads one or more image files
- Use - to read from stdin
- No arguments defaults to reading from stdin

### Examples

Convert a single image and print to terminal:

`imgbase path/to/image.png`

Convert multiple images:

`imgbase logo.png icon.jpg banner.gif`

Read from stdin (pipe an image):

`cat image.png | imgbase`

`curl -s https://example.com/image.jpg | imgbase`

Save the output to a file:

`imgbase logo.png > logo.txt`

Check version:

`imgbase --version`

## Tips for Embedding
### HTML

```html
<img src="data:image/png;base64,iVBORw0KGgoA...">
```

### CSS

```css
.logo {
    background-image: url("data:image/svg+xml;base64,PHN2ZyB4bWxu...");
}
```

### JSON

```json
{
    "icon": "data:image/jpeg;base64,/9j/4QGKRXhp..."
}
```

You can directly copy the output from imgbase into your HTML, CSS, or JSON files.
