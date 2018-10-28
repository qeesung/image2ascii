package convert

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// TestOpenImageFile test open different type image file
func TestOpenImageFile(t *testing.T) {
	assertions := assert.New(t)
	jpgFilename := "testdata/jpg_sample_image.jpg"
	openedImage, err := OpenImageFile(jpgFilename)
	assertions.True(err == nil, "jpg image format should be supported")
	assertions.True(openedImage != nil, "opened jpg file should not be nil")

	pngFilename := "testdata/png_sample_image.png"
	openedImage, err = OpenImageFile(pngFilename)
	assertions.True(err == nil, "png image format should be supported")
	assertions.True(openedImage != nil, "opened jpg file should not be nil")

	notSupported := "testdata/not_supported_sample_image"
	openedImage, err = OpenImageFile(notSupported)
	assertions.True(err != nil, "should not open unsupported image")
	assertions.True(openedImage == nil, "not supported image should be nil")
}

// TestOpenNotExistsFile test open a not exists file
func TestOpenNotExistsFile(t *testing.T) {
	assertions := assert.New(t)
	_, err := OpenImageFile("not exists")
	assertions.True(err != nil)
}

// TestImage2ASCIIMatrix test convert a image to ascii matrix
func TestImage2ASCIIMatrix(t *testing.T) {
	imageTests := []struct {
		imageFilename string
		asciiMatrix   []string
	}{
		{"testdata/3x3_black.png", []string{
			" ", " ", " ", "\n",
			" ", " ", " ", "\n",
			" ", " ", " ", "\n"}},
		{"testdata/3x3_white.png", []string{
			"@", "@", "@", "\n",
			"@", "@", "@", "\n",
			"@", "@", "@", "\n"}},
		{"testdata/8x3_multi_colors.png", []string{
			"L", "0", "L", "t", "0", "t", "G", "0", "\n",
			"i", "L", "t", "1", "L", "1", "f", "L", "\n",
			"i", "f", "i", "i", ";", "L", ";", "t", "\n",
		}},
	}

	for _, tt := range imageTests {
		t.Run(tt.imageFilename, func(t *testing.T) {
			convertOptions := DefaultOptions
			convertOptions.FitScreen = false
			convertOptions.Colored = false

			matrix := ImageFile2ASCIIMatrix(tt.imageFilename, &convertOptions)
			if !reflect.DeepEqual(matrix, tt.asciiMatrix) {
				t.Errorf("image %s convert expected to %+v, but get %+v",
					tt.imageFilename, tt.asciiMatrix, matrix)
			}
		})
	}
}

func TestImageFile2ASCIIString(t *testing.T) {
	imageTests := []struct {
		imageFilename string
		asciiString   string
	}{
		{"testdata/3x3_black.png", "   \n   \n   \n"},
		{"testdata/3x3_white.png", "@@@\n@@@\n@@@\n"},
		{"testdata/8x3_multi_colors.png", "L0Lt0tG0\niLt1L1fL\nifii;L;t\n"},
	}

	for _, tt := range imageTests {
		t.Run(tt.imageFilename, func(t *testing.T) {
			convertOptions := DefaultOptions
			convertOptions.FitScreen = false
			convertOptions.Colored = false

			charString := ImageFile2ASCIIString(tt.imageFilename, &convertOptions)
			if charString != tt.asciiString {
				t.Errorf("image %s convert expected to %+v, but get %+v",
					tt.imageFilename, tt.asciiString, charString)
			}
		})
	}
}

func TestImage2ReversedASCIIString(t *testing.T) {
	imageTests := []struct {
		imageFilename string
		asciiString   string
	}{
		{"testdata/3x3_white.png", "   \n   \n   \n"},
		{"testdata/3x3_black.png", "@@@\n@@@\n@@@\n"},
	}

	for _, tt := range imageTests {
		t.Run(tt.imageFilename, func(t *testing.T) {
			convertOptions := DefaultOptions
			convertOptions.FitScreen = false
			convertOptions.Colored = false
			convertOptions.Reversed = true

			charString := ImageFile2ASCIIString(tt.imageFilename, &convertOptions)
			if charString != tt.asciiString {
				t.Errorf("image %s convert expected to %+v, but get %+v",
					tt.imageFilename, tt.asciiString, charString)
			}
		})
	}
}

// BenchmarkBigImage2ASCIIMatrix benchmark convert big image to ascii
func BenchmarkBigImage2ASCIIMatrix(b *testing.B) {
	convertOptions := DefaultOptions
	convertOptions.FitScreen = false
	convertOptions.Colored = false
	convertOptions.ExpectedWidth = 200
	convertOptions.ExpectedHeight = 200

	for i := 0; i < b.N; i++ {
		_ = ImageFile2ASCIIMatrix("testdata/cat_2000x1500.jpg", &convertOptions)
	}
}

// BenchmarkSmallImage2ASCIIMatrix benchmark convert small image to ascii
func BenchmarkSmallImage2ASCIIMatrix(b *testing.B) {
	convertOptions := DefaultOptions
	convertOptions.FitScreen = false
	convertOptions.Colored = false
	convertOptions.ExpectedWidth = 200
	convertOptions.ExpectedHeight = 200

	for i := 0; i < b.N; i++ {
		_ = ImageFile2ASCIIMatrix("testdata/husky_200x200.jpg", &convertOptions)
	}
}

// ExampleImage2ASCIIMatrix is example
func ExampleImage2ASCISString() {
	imageFilename := "testdata/3x3_white.png"
	convertOptions := DefaultOptions
	convertOptions.FitScreen = false
	convertOptions.Colored = false
	asciiString := ImageFile2ASCIIString(imageFilename, &convertOptions)
	fmt.Println(asciiString)
	/* Output:
@@@
@@@
@@@
	 */
}
