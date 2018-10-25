package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseEmptyFilenameOptions(t *testing.T) {
	assertions := assert.New(t)
	imageFilename = ""
	_, err :=parseOptions()
	assertions.True(err != nil)
}

func TestParseOptions(t *testing.T) {
	assertions := assert.New(t)
	imageFilename = "filename"
	ratio = 0.5
	fitScreen = false
	colored = false
	expectedHeight = 100
	expectedWidth = 100
	opt, err :=parseOptions()
	assertions.True(err == nil)
	assertions.Equal(ratio, opt.Ratio)
	assertions.False(fitScreen)
	assertions.False(colored)
	assertions.Equal(expectedWidth, opt.ExpectedWidth)
	assertions.Equal(expectedHeight, opt.ExpectedHeight)
}

func TestParseUsage(t *testing.T) {
	usage()
}
