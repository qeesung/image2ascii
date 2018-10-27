// Package convert can convert a image to ascii string or matrix
package convert

import (
	"bytes"
	"github.com/qeesung/image2ascii/ascii"
	"image"
	"image/color"
	// Support decode jpeg image
	_ "image/jpeg"
	// Support deocde the png image
	_ "image/png"
	"log"
	"os"
)

// Options to convert the image to ASCII
type Options struct {
	Ratio          float64
	ExpectedWidth  int
	ExpectedHeight int
	FitScreen      bool
	Colored        bool
}

// DefaultOptions for convert image
var DefaultOptions = Options{
	Ratio:          1,
	ExpectedWidth:  -1,
	ExpectedHeight: -1,
	FitScreen:      true,
	Colored:        true,
}

// Image2ASCIIMatrix converts a image to ASCII matrix
func Image2ASCIIMatrix(image image.Image, imageConvertOptions *Options) []string {
	// Resize the convert first
	newImage := ScaleImage(image, imageConvertOptions)
	sz := newImage.Bounds()
	newWidth := sz.Max.Y
	newHeight := sz.Max.X
	rawCharValues := make([]string, 0, int(newWidth*newHeight+newWidth))
	for i := 0; i < int(newWidth); i++ {
		for j := 0; j < int(newHeight); j++ {
			pixel := color.NRGBAModel.Convert(newImage.At(j, i))
			// Convert the pixel to ascii char
			pixelConvertOptions := ascii.NewOptions()
			pixelConvertOptions.Colored = imageConvertOptions.Colored
			rawChar := ascii.ConvertPixelToASCII(pixel, &pixelConvertOptions)
			rawCharValues = append(rawCharValues, rawChar)
		}
		rawCharValues = append(rawCharValues, "\n")
	}
	return rawCharValues
}

// Image2ASCIIString converts a image to ascii matrix, and the join the matrix to a string
func Image2ASCIIString(image image.Image, options *Options) string {
	convertedPixelASCII := Image2ASCIIMatrix(image, options)
	var buffer bytes.Buffer

	for i := 0; i < len(convertedPixelASCII); i++ {
		buffer.WriteString(convertedPixelASCII[i])
	}
	return buffer.String()
}

// ImageFile2ASCIIMatrix converts a image file to ascii matrix
func ImageFile2ASCIIMatrix(imageFilename string, option *Options) []string {
	img, err := OpenImageFile(imageFilename)
	if err != nil {
		log.Fatal("open image failed : " + err.Error())
	}
	return Image2ASCIIMatrix(img, option)
}

// ImageFile2ASCIIString converts a image file to ascii string
func ImageFile2ASCIIString(imageFilename string, option *Options) string {
	img, err := OpenImageFile(imageFilename)
	if err != nil {
		log.Fatal("open image failed : " + err.Error())
	}
	return Image2ASCIIString(img, option)
}

// OpenImageFile open a image and return a image object
func OpenImageFile(imageFilename string) (image.Image, error) {
	f, err := os.Open(imageFilename)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	return img, nil
}
