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

	if options.FixedWidth != -1 {
		newWidth = options.FixedWidth
	}

	if options.FixedHeight != -1 {
		newHeight = options.FixedHeight
	}

	// use the ratio the scale the image
	if options.FixedHeight == -1 && options.FixedWidth == -1 && ratio != 1 {
		newWidth = ScaleWidthByRatio(float64(sz.Max.X), ratio)
		newHeight = ScaleHeightByRatio(float64(sz.Max.Y), ratio)
	}

	// fit the screen
	if ratio == 1 &&
		options.FixedWidth == -1 &&
		options.FixedHeight == -1 &&
		options.FitScreen {
		fitWidth, fitHeight, err := CalcProportionalFittingScreenSize(image)
		if err != nil {
			log.Fatal(err)
		}
		newWidth = int(fitWidth)
		newHeight = int(fitHeight)
	}

	//Stretch the picture to overspread the terminal
	if ratio == 1 &&
		options.FixedWidth == -1 &&
		options.FixedHeight == -1 &&
		!options.FitScreen &&
		options.StretchedScreen {
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

// CalcProportionalFittingScreenSize proportional scale the image
// so that the terminal can just show the picture.
func CalcProportionalFittingScreenSize(image image.Image) (newWidth, newHeight int, err error) {
	if !isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		return 0, 0,
			errors.New("can not detect the terminal, please disable the '-s fitScreen' option")
	}

	screenWidth, _ := terminal.Width()
	screenHeight, _ := terminal.Height()
	sz := image.Bounds()
	newWidth, newHeight = CalcFitSize(
		float64(screenWidth),
		float64(screenHeight),
		float64(sz.Max.X),
		float64(sz.Max.Y))
	return
}

// CalcFitSizeRatio through the given length and width,
// the computation can match the optimal scaling ratio of the length and width.
// In other words, it is able to give a given size rectangle to contain pictures
// Either match the width first, then scale the length equally,
// or match the length first, then scale the height equally.
// More detail please check the example
func CalcFitSizeRatio(width, height, imageWidth, imageHeight float64) (ratio float64) {
	ratio = 1.0
	// try to fit the height
	ratio = height / imageHeight
	scaledWidth := imageWidth * ratio / charWidth()
	if scaledWidth < width {
		return ratio / charWidth()
	}

	// try to fit the width
	ratio = width / imageWidth
	return ratio
}

// CalcFitSize through the given length and width ,
// Calculation is able to match the length and width of
// the specified size, and is proportional scaling.
func CalcFitSize(width, height, toBeFitWidth, toBeFitHeight float64) (fitWidth, fitHeight int) {
	ratio := CalcFitSizeRatio(width, height, toBeFitWidth, toBeFitHeight)
	fitWidth = ScaleWidthByRatio(toBeFitWidth, ratio)
	fitHeight = ScaleHeightByRatio(toBeFitHeight, ratio)
	return
}

func ScaleWidthByRatio(width float64, ratio float64) int {
	return int(width * ratio)
}

func ScaleHeightByRatio(height float64, ratio float64) int {
	return int(height * ratio * charWidth())
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
