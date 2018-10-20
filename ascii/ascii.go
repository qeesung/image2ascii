// The ascii package can convert a image pixel to a raw char
// base on it's RGBA value, in another word, input a image pixel
// output a raw char ascii.
package ascii

import (
	"image/color"
	"math"
	"reflect"
)

// Convert a pixel to a ASCII char string
func ConvertPixelToASCII(pixel color.Color, options *Options) string {
	defaultOptions := NewOptions()
	defaultOptions.mergeOptions(options)

	if defaultOptions.reverse {
		defaultOptions.pixels = reverse(defaultOptions.pixels)
	}

	r := reflect.ValueOf(pixel).FieldByName("R").Uint()
	g := reflect.ValueOf(pixel).FieldByName("G").Uint()
	b := reflect.ValueOf(pixel).FieldByName("B").Uint()
	a := reflect.ValueOf(pixel).FieldByName("A").Uint()
	value := intensity(r, g, b, a)

	// Choose the char
	precision := float64(255 * 3 / (len(options.pixels) - 1))
	rawChar := options.pixels[roundValue(float64(value)/precision)]
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
