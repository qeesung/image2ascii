package terminal

import (
	"errors"
	"github.com/mattn/go-isatty"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
	"os"
	"runtime"
)

const (
	charWidthWindows = 0.714
	charWidthOther   = 0.5
)

func NewTerminalAccessor() Terminal {
	return Accessor{}
}

// Terminal get the terminal basic information
type Terminal interface {
	CharWidth() float64
	ScreenSize() (width, height int, err error)
	IsWindows() bool
}

type Accessor struct {
}

func (accessor Accessor) CharWidth() float64 {
	if accessor.IsWindows() {
		return charWidthWindows
	}
	return charWidthOther
}

func (accessor Accessor) IsWindows() bool {
	return runtime.GOOS == "windows"
}

func (accessor Accessor) ScreenSize() (newWidth, newHeight int, err error) {
	if !isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		return 0, 0,
			errors.New("can not detect the terminal")
	}

	x, _ := terminal.Width()
	y, _ := terminal.Height()

	return int(x), int(y), nil
}
