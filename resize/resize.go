// The resize package resize the image to expected size
// base on the ratio, for the most matched display
package resize

import (
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"os"
	"runtime"
)

// Resize the convert to expected size base on the ratio
func ScaleImage(image image.Image, ratio float64) (newImage image.Image) {
	sz := image.Bounds()
	newWidth := int(float64(sz.Max.X) * ratio)
	newHeight := int(float64(sz.Max.Y) * ratio * charWidth())

	newImage = resize.Resize(uint(newWidth), uint(newHeight), image, resize.Lanczos3)
	out, err := os.Create("test_resized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, newImage, nil)
	return
}

// Get the terminal char width on different system
func charWidth() float64 {
	if isWindows() {
		return 0.714
	} else {
		return 0.625
	}
}

// Check if current system is windows
func isWindows() bool {
	return runtime.GOOS == "windows"
}
