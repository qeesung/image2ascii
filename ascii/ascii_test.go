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
	assertions.False(newOptions.Reverse, "Default reverse option should be false")
	assertions.Equal(" .,:;i1tfLCG08@", string(newOptions.Pixels), "Default pixels should be  .,:;i1tfLCG08@")
}

// TestMergeOptions test merge the options
func TestMergeOptions(t *testing.T) {
	assertions := assert.New(t)
	options1 := NewOptions()
	options2 := NewOptions()
	options2.Colored = false
	options1.mergeOptions(&options2)
	assertions.False(options1.Reverse, "Merged reverse option should be false")
	assertions.False(options1.Colored, "Merged colored option should be false")
}

// TestConvertPixelToASCIIWhiteColor convert a white image pixel to ascii string
func TestConvertPixelToASCIIWhiteColor(t *testing.T) {
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
	convertedChar := ConvertPixelToASCII(pixel, &defaultOptions)
	lastPixelChar := defaultOptions.Pixels[len(defaultOptions.Pixels)-1]
	assertions.Equal(convertedChar, string([]byte{lastPixelChar}),
		fmt.Sprintf("White color chould be converted to %s", string([]byte{lastPixelChar})))

	defaultOptions.Colored = false
	defaultOptions.Reverse = true
	convertedChar = ConvertPixelToASCII(pixel, &defaultOptions)
	firstPixelChar := reverse(defaultOptions.Pixels)[0]
	assertions.Equal(convertedChar, string([]byte{firstPixelChar}),
		fmt.Sprintf("Reverse white color chould be converted to %s", string([]byte{firstPixelChar})))
}

// TestConvertPixelToASCIIBlackColor convert a white image pixel to ascii string
func TestConvertPixelToASCIIBlackColor(t *testing.T) {
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
	convertedChar := ConvertPixelToASCII(pixel, &defaultOptions)
	firstPixelChar := defaultOptions.Pixels[0]
	assertions.Equal(convertedChar, string([]byte{firstPixelChar}),
		fmt.Sprintf("Black color chould be converted to %s", string([]byte{firstPixelChar})))

	defaultOptions.Colored = false
	defaultOptions.Reverse = true
	convertedChar = ConvertPixelToASCII(pixel, &defaultOptions)
	lastPixelChar := reverse(defaultOptions.Pixels)[len(defaultOptions.Pixels)-1]
	assertions.Equal(convertedChar, string([]byte{lastPixelChar}),
		fmt.Sprintf("Reverse Black color chould be converted to %s", string([]byte{lastPixelChar})))
}

// TestReverseSlice test reverse a slice
func TestReverseSlice(t *testing.T) {
	s := []byte{1, 2, 3, 4, 5}
	reversedSlice := reverse(s)
	expectedReversedSlice := []byte{5, 4, 3, 2, 1}
	assert.True(t, reflect.DeepEqual(reversedSlice, expectedReversedSlice),
		fmt.Sprintf("%+v reversed should equal to %+v", s, expectedReversedSlice))

	s = []byte{1, 2, 3, 4}
	reversedSlice = reverse(s)
	expectedReversedSlice = []byte{4, 3, 2, 1}
	assert.True(t, reflect.DeepEqual(reversedSlice, expectedReversedSlice),
		fmt.Sprintf("%+v reversed should equal to %+v", s, expectedReversedSlice))
}
