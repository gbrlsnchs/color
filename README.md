# color (wrapper for ANSI escape codes)

[![Sourcegraph](https://sourcegraph.com/github.com/gbrlsnchs/color/-/badge.svg)](https://sourcegraph.com/github.com/gbrlsnchs/color?badge)
[![GoDoc](https://godoc.org/github.com/gbrlsnchs/color?status.svg)](https://godoc.org/github.com/gbrlsnchs/color)
[![Minimal version](https://img.shields.io/badge/minimal%20version-go1.10%2B-5272b4.svg)](https://golang.org/doc/go1.10)

## About
This package is a wrapper for coloring text using [ANSI escape codes](https://en.wikipedia.org/wiki/ANSI_escape_code).

## Usage
Full documentation [here](https://godoc.org/github.com/gbrlsnchs/color).

### Installing
#### Go 1.10
`vgo get -u github.com/gbrlsnchs/color`
#### Go 1.11 or after
`go get -u github.com/gbrlsnchs/color`

### Importing
```go
import (
	// ...

	"github.com/gbrlsnchs/color"
)
```

### Coloring text
```go
colorRed := color.New(color.CodeFgRed)
log.Print(colorRed.Wrap("Hello, world!"))
log.Print(colorRed.Wrapf("Hello, %s!", "world")
log.Print(colorRed.Wrapln("Hello, world!")
```

### Disabling color support
```go
colorYellow := color.New(color.CodeFgYellow)
colorYellow.SetDisabled(true)
log.Print(colorYellow.Wrap("This won't be colored...")
```

### Adding more attributes
```go
colorBoldGreen := color.New(color.CodeBold, color.CodeFgGreen)
log.Print(colorBoldGreen.Wrap("This is a mix of bold and green text!")

colorUnderlinedMagenta := color.New(color.CodeFgMagenta)
colorUnderlinedMagenta.Add(color.CodeUnderline)
log.Print(colorUnderlinedMagenta.Wrap("This is a mix of underlined and magenta text!")
```

## Contributing
### How to help
- For bugs and opinions, please [open an issue](https://github.com/gbrlsnchs/color/issues/new)
- For pushing changes, please [open a pull request](https://github.com/gbrlsnchs/color/compare)
