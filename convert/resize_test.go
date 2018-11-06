package convert

import (
	"fmt"
	terminal2 "github.com/qeesung/image2ascii/terminal"
	"github.com/qeesung/image2ascii/terminal/mocks"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// TestScaleImageWithFixedHeight test scale the image by fixed height
func TestScaleImageWithFixedHeight(t *testing.T) {
	handler := NewResizeHandler()
	assertions := assert.New(t)
	imageFilePath := "testdata/cat_2000x1500.jpg"
	img, err := OpenImageFile(imageFilePath)
	assertions.True(img != nil)
	assertions.True(err == nil)

	options := DefaultOptions
	options.Colored = false
	options.FixedHeight = 100

	scaledImage := handler.ScaleImage(img, &options)
	sz := scaledImage.Bounds()
	oldSz := img.Bounds()
	assertions.Equal(100, sz.Max.Y, "scaled image height should be 100")
	assertions.Equal(oldSz.Max.X, sz.Max.X, "scaled image width should be changed")
}

// TestScaleImageWithFixedWidth test scale the image by fixed width
func TestScaleImageWithFixedWidth(t *testing.T) {
	handler := NewResizeHandler()
	assertions := assert.New(t)
	imageFilePath := "testdata/cat_2000x1500.jpg"
	img, err := OpenImageFile(imageFilePath)
	assertions.True(img != nil)
	assertions.True(err == nil)

	options := DefaultOptions
	options.Colored = false
	options.FixedWidth = 200

	scaledImage := handler.ScaleImage(img, &options)
	sz := scaledImage.Bounds()
	oldSz := img.Bounds()
	assertions.Equal(oldSz.Max.Y, sz.Max.Y, "scaled image height should be changed")
	assertions.Equal(200, sz.Max.X, "scaled image width should be 200")
}

// TestScaleImageWithFixedWidthHeight test scale the image by fixed width
func TestScaleImageWithFixedWidthHeight(t *testing.T) {
	handler := NewResizeHandler()
	assertions := assert.New(t)
	imageFilePath := "testdata/cat_2000x1500.jpg"
	img, err := OpenImageFile(imageFilePath)
	assertions.True(img != nil)
	assertions.True(err == nil)

	options := DefaultOptions
	options.Colored = false
	options.FixedWidth = 200
	options.FixedHeight = 100

	scaledImage := handler.ScaleImage(img, &options)
	sz := scaledImage.Bounds()
	assertions.Equal(100, sz.Max.Y, "scaled image height should be 100")
	assertions.Equal(200, sz.Max.X, "scaled image width should be 200")
}

// TestScaleImageByRatio test scale image by ratio
func TestScaleImageByRatio(t *testing.T) {
	handler := NewResizeHandler()
	terminal := terminal2.NewTerminalAccessor()
	assertions := assert.New(t)
	imageFilePath := "testdata/cat_2000x1500.jpg"
	img, err := OpenImageFile(imageFilePath)
	assertions.True(img != nil)
	assertions.True(err == nil)

	options := DefaultOptions
	options.Colored = false
	options.Ratio = 0.5

	scaledImage := handler.ScaleImage(img, &options)
	sz := scaledImage.Bounds()
	oldSz := img.Bounds()
	expectedHeight := int(float64(oldSz.Max.Y) * 0.5 * terminal.CharWidth())
	expectedWidth := int(float64(oldSz.Max.X) * 0.5)
	assertions.Equal(expectedHeight, sz.Max.Y, fmt.Sprintf("scaled image height should be %d", expectedHeight))
	assertions.Equal(expectedWidth, sz.Max.X, fmt.Sprintf("scaled image width should be %d", expectedHeight))
}

// TestCalcFitSize test calc the fit size
func TestCalcFitSize(t *testing.T) {
	handler := ImageResizeHandler{
		terminal: terminal2.NewTerminalAccessor(),
	}
	fitSizeTests := []struct {
		width         int
		height        int
		toBeFitWidth  int
		toBeFitHeight int
		fitWidth      int
		fitHeight     int
	}{
		{width: 100, height: 80, toBeFitWidth: 50, toBeFitHeight: 120, fitWidth: 66, fitHeight: 80},
		{width: 100, height: 80, toBeFitWidth: 120, toBeFitHeight: 50, fitWidth: 100, fitHeight: 20},
	}
	for _, tt := range fitSizeTests {
		t.Run(fmt.Sprintf("%d, %d -> %d, %d",
			tt.width, tt.height, tt.toBeFitWidth, tt.toBeFitHeight), func(t *testing.T) {
			fitWidth, fitHeight := handler.CalcFitSize(
				float64(tt.width),
				float64(tt.height),
				float64(tt.toBeFitWidth),
				float64(tt.toBeFitHeight))
			if fitWidth != tt.fitWidth || fitHeight != tt.fitHeight {
				t.Errorf("%d, %d -> %d, %d should be %d, %d, but get %d, %d",
					tt.width, tt.height, tt.toBeFitWidth,
					tt.toBeFitHeight, tt.fitWidth, tt.fitHeight,
					fitWidth, fitHeight)
			}
		})
	}
}

func TestFitTheTerminalScreenSize(t *testing.T) {
	assertions := assert.New(t)
	// mock
	terminalMock := mocks.Terminal{}
	terminalMock.On("IsWindows").Return(false)
	terminalMock.On("CharWidth").Return(0.5)
	terminalMock.On("ScreenSize").Return(100, 80, nil)
	handler := ImageResizeHandler{
		terminal: &terminalMock,
	}
	initResizeResolver(&handler)

	imageFilePath := "testdata/cat_2000x1500.jpg"
	img, err := OpenImageFile(imageFilePath)
	if err != nil {
		log.Fatal("open image file " + imageFilePath + " failed")
	}

	assertions.False(terminalMock.IsWindows())
	assertions.Equal(terminalMock.CharWidth(), 0.5)
	screenWidth, screenHeight, err := terminalMock.ScreenSize()
	assertions.Equal(screenWidth, 100)
	assertions.Equal(screenHeight, 80)
	assertions.True(err == nil)

	options := DefaultOptions
	options.Colored = false
	options.FitScreen = true
	scaledImage := handler.ScaleImage(img, &options)
	sz := scaledImage.Bounds()
	assertions.Equal(100, sz.Max.X)
	assertions.Equal(37, sz.Max.Y)
}

func TestStretchTheTerminalScreenSize(t *testing.T) {
	assertions := assert.New(t)
	// mock
	terminalMock := mocks.Terminal{}
	terminalMock.On("IsWindows").Return(false)
	terminalMock.On("CharWidth").Return(0.5)
	terminalMock.On("ScreenSize").Return(100, 80, nil)
	handler := ImageResizeHandler{
		terminal: &terminalMock,
	}
	initResizeResolver(&handler)

	imageFilePath := "testdata/cat_2000x1500.jpg"
	img, err := OpenImageFile(imageFilePath)
	if err != nil {
		log.Fatal("open image file " + imageFilePath + " failed")
	}

	assertions.False(terminalMock.IsWindows())
	assertions.Equal(terminalMock.CharWidth(), 0.5)
	screenWidth, screenHeight, err := terminalMock.ScreenSize()
	assertions.Equal(screenWidth, 100)
	assertions.Equal(screenHeight, 80)
	assertions.True(err == nil)

	options := DefaultOptions
	options.Colored = false
	options.StretchedScreen = true
	options.FitScreen = false
	scaledImage := handler.ScaleImage(img, &options)
	sz := scaledImage.Bounds()
	assertions.Equal(100, sz.Max.X)
	assertions.Equal(80, sz.Max.Y)
}

// ExampleScaleImage is scale image example
func ExampleImageResizeHandler_ScaleImage() {
	handler := NewResizeHandler()
	imageFilePath := "testdata/cat_2000x1500.jpg"
	img, err := OpenImageFile(imageFilePath)
	if err != nil {
		log.Fatal("open image file " + imageFilePath + " failed")
	}

	options := DefaultOptions
	options.Colored = false
	options.FixedWidth = 200
	options.FixedHeight = 100

	scaledImage := handler.ScaleImage(img, &options)
	sz := scaledImage.Bounds()
	fmt.Print(sz.Max.X, sz.Max.Y)
	// output: 200 100
}

// BenchmarkScaleImage benchmark scale big image
func BenchmarkScaleBigImage(b *testing.B) {
	handler := NewResizeHandler()
	imageFilePath := "testdata/cat_2000x1500.jpg"
	img, err := OpenImageFile(imageFilePath)
	if err != nil {
		log.Fatalf("Open image file %s failed", imageFilePath)
	}

	options := DefaultOptions
	options.Colored = false
	options.FitScreen = false
	options.FixedHeight = 100
	options.FixedWidth = 100

	for i := 0; i < b.N; i++ {
		_ = handler.ScaleImage(img, &options)
	}
}

// BenchmarkScaleSmallImage benchmark scale small image
func BenchmarkScaleSmallImage(b *testing.B) {
	handler := NewResizeHandler()
	imageFilePath := "testdata/husky_200x200.jpg"
	img, err := OpenImageFile(imageFilePath)
	if err != nil {
		log.Fatalf("Open image file %s failed : %s", imageFilePath, err.Error())
	}

	options := DefaultOptions
	options.Colored = false
	options.FitScreen = false
	options.FixedHeight = 100
	options.FixedWidth = 100

	for i := 0; i < b.N; i++ {
		_ = handler.ScaleImage(img, &options)
	}
}
