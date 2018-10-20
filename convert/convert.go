// The convert package convert a image to ascii string or matrix
package convert

import (
	"bytes"
	"github.com/qeesung/image2asicc/ascii"
	"github.com/qeesung/image2asicc/resize"
	"image"
	"image/color"
)

type Options struct {
	Ratio float64
}

// Convert a image to ascii matrix
func Image2ASCIIMatrix(image image.Image, options *Options) []string {
	// Resize the convert first
	newImage := resize.ScaleImage(image, options.Ratio)
	sz := newImage.Bounds()
	newWidth := sz.Max.Y
	newHeight := sz.Max.X
	rawCharValues := make([]string, int(newWidth*newHeight))
	for i := 0; i < int(newWidth); i++ {
		for j := 0; j < int(newHeight); j++ {
			pixel := color.NRGBAModel.Convert(newImage.At(j, i))
			// Convert the pixel to ascii char
			pixelConvertOptions := ascii.NewOptions()
			rawChar := ascii.ConvertPixelToASCII(pixel, &pixelConvertOptions)
			rawCharValues = append(rawCharValues, rawChar)
		}
		rawCharValues = append(rawCharValues, "\n")
	}
	return rawCharValues
}

// Convert a image to ascii matrix, then concat the matrix value
// to a long string for easy display
func Image2ASCIIString(image image.Image, options *Options) string {
	convertedPixelASCII := Image2ASCIIMatrix(image, options)
	var buffer bytes.Buffer

	for i := 0; i < len(convertedPixelASCII); i++ {
		buffer.WriteString(convertedPixelASCII[i])
	}
	return buffer.String()
}
