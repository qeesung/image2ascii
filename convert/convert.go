// Package convert can convert a image to ascii string or matrix
package convert

import (
	"bytes"
	"github.com/qeesung/image2ascii/ascii"
	"image"
	"image/color"
	"runtime"
	// Support decode jpeg image
	_ "image/jpeg"
	// Support deocde the png image
	_ "image/png"
	"log"
	"os"
)

type Pixel struct {
	i              int
	j              int
	color          color.Color
	convertedChars string
}

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
	newWidth := sz.Max.X
	newHeight := sz.Max.Y
	rawCharValues := make([]string, 0, int(newWidth*newHeight+newHeight))
	for i := 0; i < int(newHeight); i++ {
		for j := 0; j < int(newWidth); j++ {
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

func Image2ASCIIMatrixParallel(image image.Image, imageConvertOptions *Options) []string {
	// Resize the convert first
	newImage := ScaleImage(image, imageConvertOptions)
	sz := newImage.Bounds()
	newWidth := sz.Max.X
	newHeight := sz.Max.Y
	cap := int(newWidth*newHeight + newHeight)
	rawCharValues := make([]string, cap, cap)

	// set the new line string
	for i := 0; i < newHeight; i++ {
		newLineIndex := (i+1)*(newWidth+1) - 1
		rawCharValues[newLineIndex] = "\n"
	}

	pixelChannel := make(chan Pixel, newWidth)
	charChannel := make(chan Pixel, newWidth)
	defer close(pixelChannel)
	defer close(charChannel)

	// schedule the convert task
	go func() {
		for i := 0; i < int(newHeight); i++ {
			for j := 0; j < int(newWidth); j++ {
				pixelColor := color.NRGBAModel.Convert(newImage.At(j, i))
				pixelChannel <- Pixel{
					i:              i,
					j:              j,
					color:          pixelColor,
					convertedChars: "",
				}
			}
		}
	}()

	// convert the pixel
	pixelConvertOptions := ascii.NewOptions()
	pixelConvertOptions.Colored = imageConvertOptions.Colored
	coreCount := runtime.NumCPU()
	for i := 0; i < coreCount; i++ {
		go func() {
			for pixel := range pixelChannel {
				rawChar := ascii.ConvertPixelToASCII(pixel.color, &pixelConvertOptions)
				pixel.convertedChars = rawChar
				charChannel <- pixel
			}
		}()
	}

	// get the results
	for k := 0; k < int(newWidth)*int(newHeight); k++ {
		convertedPixel := <-charChannel
		i, j := convertedPixel.i, convertedPixel.j
		index := i*int(newWidth+1) + j
		rawCharValues[index] = convertedPixel.convertedChars
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

func Image2ASCIIStringParallel(image image.Image, options *Options) string {
	convertedPixelASCII := Image2ASCIIMatrixParallel(image, options)
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

// ImageFile2ASCIIMatrix converts a image file to ascii matrix
func ImageFile2ASCIIMatrixParallel(imageFilename string, option *Options) []string {
	img, err := OpenImageFile(imageFilename)
	if err != nil {
		log.Fatal("open image failed : " + err.Error())
	}
	return Image2ASCIIMatrixParallel(img, option)
}

// ImageFile2ASCIIString converts a image file to ascii string
func ImageFile2ASCIIString(imageFilename string, option *Options) string {
	img, err := OpenImageFile(imageFilename)
	if err != nil {
		log.Fatal("open image failed : " + err.Error())
	}
	return Image2ASCIIString(img, option)
}

// ImageFile2ASCIIString converts a image file to ascii string
func ImageFile2ASCIIStringParallel(imageFilename string, option *Options) string {
	img, err := OpenImageFile(imageFilename)
	if err != nil {
		log.Fatal("open image failed : " + err.Error())
	}
	return Image2ASCIIStringParallel(img, option)
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
