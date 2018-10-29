package ascii

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"image/color"
	"reflect"
	"testing"
)

// TestNewOptions create options with default values
func TestNewOptions(t *testing.T) {
	newOptions := NewOptions()
	assertions := assert.New(t)
	assertions.True(newOptions.Colored, "Default colored option should be true")
	assertions.False(newOptions.Reversed, "Default reverse option should be false")
	assertions.Equal(" .,:;i1tfLCG08@", string(newOptions.Pixels), "Default pixels should be  .,:;i1tfLCG08@")
}

// TestMergeOptions test merge the options
func TestMergeOptions(t *testing.T) {
	assertions := assert.New(t)
	options1 := NewOptions()
	options2 := NewOptions()
	options2.Colored = false
	options1.mergeOptions(&options2)
	assertions.False(options1.Reversed, "Merged reverse option should be false")
	assertions.False(options1.Colored, "Merged colored option should be false")
}

// TestConvertPixelToASCIIWhiteColor convert a white image pixel to ascii string
func TestConvertPixelToASCIIWhiteColor(t *testing.T) {
	converter := NewPixelConverter()
	assertions := assert.New(t)
	r, g, b, a := uint8(255), uint8(255), uint8(255), uint8(255)
	pixel := color.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}

	defaultOptions := NewOptions()
	defaultOptions.Colored = false
	convertedChar := converter.ConvertPixelToASCII(pixel, &defaultOptions)
	lastPixelChar := defaultOptions.Pixels[len(defaultOptions.Pixels)-1]
	assertions.Equal(convertedChar, string([]byte{lastPixelChar}),
		fmt.Sprintf("White color chould be converted to %s", string([]byte{lastPixelChar})))

	defaultOptions.Colored = false
	defaultOptions.Reversed = true
	convertedChar = converter.ConvertPixelToASCII(pixel, &defaultOptions)
	firstPixelChar := defaultOptions.Pixels[0]
	assertions.Equal(convertedChar, string([]byte{firstPixelChar}),
		fmt.Sprintf("Reversed white color chould be converted to %s", string([]byte{firstPixelChar})))
}

// TestConvertPixelToASCIIBlackColor convert a white image pixel to ascii string
func TestConvertPixelToASCIIBlackColor(t *testing.T) {
	converter := NewPixelConverter()
	assertions := assert.New(t)
	r, g, b, a := uint8(0), uint8(0), uint8(0), uint8(0)
	pixel := color.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}

	defaultOptions := NewOptions()
	defaultOptions.Colored = false
	convertedChar := converter.ConvertPixelToASCII(pixel, &defaultOptions)
	firstPixelChar := defaultOptions.Pixels[0]
	assertions.Equal(convertedChar, string([]byte{firstPixelChar}),
		fmt.Sprintf("Black color chould be converted to %s", string([]byte{firstPixelChar})))

	defaultOptions.Colored = false
	defaultOptions.Reversed = true
	convertedChar = converter.ConvertPixelToASCII(pixel, &defaultOptions)
	lastPixelChar := defaultOptions.Pixels[len(defaultOptions.Pixels)-1]
	assertions.Equal(convertedChar, string([]byte{lastPixelChar}),
		fmt.Sprintf("Reversed Black color chould be converted to %s", string([]byte{lastPixelChar})))
}

func TestColoredASCIIChar(t *testing.T) {
	converter := NewPixelConverter()
	assertions := assert.New(t)
	r, g, b, a := uint8(123), uint8(123), uint8(123), uint8(255)
	pixel := color.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}
	defaultOptions := NewOptions()
	defaultOptions.Colored = true
	coloredChar := converter.ConvertPixelToASCII(pixel, &defaultOptions)
	assertions.True(len(coloredChar) > 1)
}

// TestReverseSlice test reverse a slice
func TestReverseSlice(t *testing.T) {
	converter := PixelASCIIConverter{}
	s := []byte{1, 2, 3, 4, 5}
	reversedSlice := converter.reverse(s)
	expectedReversedSlice := []byte{5, 4, 3, 2, 1}
	assert.True(t, reflect.DeepEqual(reversedSlice, expectedReversedSlice),
		fmt.Sprintf("%+v reversed should equal to %+v", s, expectedReversedSlice))

	s = []byte{1, 2, 3, 4}
	reversedSlice = converter.reverse(s)
	expectedReversedSlice = []byte{4, 3, 2, 1}
	assert.True(t, reflect.DeepEqual(reversedSlice, expectedReversedSlice),
		fmt.Sprintf("%+v reversed should equal to %+v", s, expectedReversedSlice))
}

// ExampleConvertPixelToASCII is a example convert pixel to ascii char
func ExamplePixelASCIIConverter_ConvertPixelToASCII() {
	converter := NewPixelConverter()
	// Create the pixel
	r, g, b, a := uint8(255), uint8(255), uint8(255), uint8(255)
	pixel := color.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}

	// Create the convert options
	defaultOptions := NewOptions()
	defaultOptions.Colored = false
	convertedChar := converter.ConvertPixelToASCII(pixel, &defaultOptions)
	fmt.Println(convertedChar)
	// Output: @
}
