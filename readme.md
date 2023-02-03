# goafm

A Go library for parsing Adobe Font Metrics (AFM) files. Supports the full specification.

Adobe Font Metrics (AFM) files provide information about the size, position, and spacing of characters in a font. They are used by software applications to correctly display text using a specific font. The `goafm` library provides a convenient way to parse AFM files and extract the information they contain.

## Features

Parsing & extraction of the following information:
- Font metrics
- Character metrics
- Kerning pairs
- Track kerning
- Composites Characters

## Installation

Use `go get` to install the library:

```bash
go get github.com/Sett17/goafm
```

## Usage

`ParseFile` takes in a `filename` and returns a `FontMetric` pointer and an `error`.

`Parse` takes in a `byte slice` and returns a `FontMetric` pointer and an `error`.

```go
package main

import (
  "fmt"

  "github.com/Sett17/goafm"
)

func main() {
  font, err := goafm.ParseFile("font.afm")
  if err != nil {
    panic(err)
  }
  fmt.Printf("%#v", font)
}
```