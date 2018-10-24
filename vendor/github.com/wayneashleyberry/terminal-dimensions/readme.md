[![GoDoc](https://godoc.org/github.com/wayneashleyberry/terminal-dimensions?status.svg)](https://godoc.org/github.com/wayneashleyberry/terminal-dimensions)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/terminal-dimensions)](https://goreportcard.com/report/github.com/wayneashleyberry/terminal-dimensions)
[![Build Status](https://travis-ci.org/wayneashleyberry/terminal-dimensions.svg?branch=master)](https://travis-ci.org/wayneashleyberry/terminal-dimensions)
[![Coverage Status](https://coveralls.io/repos/github/wayneashleyberry/terminal-dimensions/badge.svg?branch=master)](https://coveralls.io/github/wayneashleyberry/terminal-dimensions?branch=master)

```sh
go get github.com/wayneashleyberry/terminal-dimensions
```

```go
package main

import (
	"fmt"

	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

func main() {
	x, _ := terminal.Width()
	y, _ := terminal.Height()
	fmt.Printf("Terminal is %d wide and %d high", x, y)
}
```
