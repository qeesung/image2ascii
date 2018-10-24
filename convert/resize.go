package convert

import (
	"errors"
	"github.com/mattn/go-isatty"
	"github.com/nfnt/resize"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
	"image"
	"log"
	"os"
	"runtime"
)

// ScaleImage resize the convert to expected size base on the convert options
func ScaleImage(image image.Image, options *Options) (newImage image.Image) {
	sz := image.Bounds()
	ratio := options.Ratio
	newHeight := sz.Max.Y
	newWidth := sz.Max.X

	if options.ExpectedWidth != -1 {
		newWidth = options.ExpectedWidth
	}

	if options.ExpectedHeight != -1 {
		newHeight = options.ExpectedHeight
	}

	// use the ratio the scale the image
	if options.ExpectedHeight == -1 && options.ExpectedWidth == -1 && ratio != 1 {
		newWidth = int(float64(sz.Max.X) * ratio)
		newHeight = int(float64(sz.Max.Y) * ratio * charWidth())
	}

	// fit the screen
	// get the fit the screen size
	if ratio == 1 &&
		options.ExpectedWidth == -1 &&
		options.ExpectedHeight == -1 &&
		options.FitScreen {
		screenWidth, screenHeight, err := getTerminalScreenSize()
		if err != nil {
			log.Fatal(err)
		}
		newWidth = int(screenWidth)
		newHeight = int(screenHeight)
	}

	newImage = resize.Resize(uint(newWidth), uint(newHeight), image, resize.Lanczos3)
	return
}

// charWidth get the terminal char width on different system
func charWidth() float64 {
	if isWindows() {
		return 0.714
	}
	return 0.5
}

// isWindows check if current system is windows
func isWindows() bool {
	return runtime.GOOS == "windows"
}

// getTerminalScreenSize get the current terminal screen size
func getTerminalScreenSize() (newWidth, newHeight uint, err error) {
	if !isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		return 0, 0,
			errors.New("can not detect the terminal, please disable the '-s fitScreen' option")
	}

	x, _ := terminal.Width()
	y, _ := terminal.Height()

	return x, y, nil
}
