// The ascii package can convert a image pixel to a raw char
// base on it's RGBA value, in another word, input a image pixel
// output a raw char ascii.
package ascii

import (
	"github.com/aybabtme/rgbterm"
	"image/color"
	"math"
	"reflect"
)

// Convert a pixel to a ASCII char string
func ConvertPixelToASCII(pixel color.Color, options *Options) string {
	defaultOptions := NewOptions()
	defaultOptions.mergeOptions(options)

	if defaultOptions.Reverse {
		defaultOptions.Pixels = reverse(defaultOptions.Pixels)
	}

	r := reflect.ValueOf(pixel).FieldByName("R").Uint()
	g := reflect.ValueOf(pixel).FieldByName("G").Uint()
	b := reflect.ValueOf(pixel).FieldByName("B").Uint()
	a := reflect.ValueOf(pixel).FieldByName("A").Uint()
	value := intensity(r, g, b, a)

	// Choose the char
	precision := float64(255 * 3 / (len(options.Pixels) - 1))
	rawChar := options.Pixels[roundValue(float64(value)/precision)]
	if options.Colored {
		return decorateWithColor(r, g, b, rawChar)
	}
	return string([]byte{rawChar})
}

func roundValue(value float64) int {
	return int(math.Floor(value + 0.5))
}

func reverse(numbers []byte) []byte {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func intensity(r, g, b, a uint64) uint64 {
	return (r + g + b) * a / 255
}

func decorateWithColor(r, g, b uint64, rawChar byte) string {
	coloredChar := rgbterm.FgString(string([]byte{rawChar}), uint8(r), uint8(g), uint8(b))
	return coloredChar
}
