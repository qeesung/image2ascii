package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/qeesung/image2ascii/convert"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

var imageFilename string
var ratio float64
var expectedWidth int
var expectedHeight int
var fitScreen bool
var colored bool

func init() {
	flag.StringVar(&imageFilename, "f", "", "Image filename to be convert")
	flag.Float64Var(&ratio, "r", 1, "Ratio to scale the image, ignored when use -w or -g")
	flag.IntVar(&expectedWidth, "w", -1, "Expected image width, -1 for image default width")
	flag.IntVar(&expectedHeight, "g", -1, "Expected image height, -1 for image default height")
	flag.BoolVar(&fitScreen, "s", true, "Fit the terminal screen, ignored when use -w, -g, -r")
	flag.BoolVar(&colored, "c", true, "Colored the ascii when output to the terminal")
	flag.Usage = usage
}

func main() {
	flag.Parse()
	if convertOptions, err := parseOptions(); err == nil {
		fmt.Print(convert.ImageFile2ASCIIString(imageFilename, convertOptions))
	} else {
		usage()
	}
}

func parseOptions() (*convert.Options, error) {
	if imageFilename == "" {
		return nil, errors.New("image file should not be empty")
	}
	// config  the options
	convertOptions := &convert.Options{
		Ratio:          ratio,
		ExpectedHeight: expectedHeight,
		ExpectedWidth:  expectedWidth,
		FitScreen:      fitScreen,
		Colored:        colored,
	}
	return convertOptions, nil
}

func usage() {
	fmt.Fprintf(os.Stderr, `image2ascii version: image2ascii/1.0.0 
>> HomePage: https://github.com/qeesung/image2ascii
>> Issue   : https://github.com/qeesung/image2ascii/issues
>> Author  : qeesung
Usage: image2ascii [-s] -f <filename> -r <ratio> -w <width> -g <height>

Options:
`)
	flag.PrintDefaults()
}
