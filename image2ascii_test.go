package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseEmptyFilenameOptions(t *testing.T) {
	assertions := assert.New(t)
	imageFilename = ""
	_, err := parseOptions()
	assertions.True(err != nil)
}

func TestParseOptions(t *testing.T) {
	assertions := assert.New(t)
	imageFilename = "filename"
	ratio = 0.5
	fitScreen = false
	colored = false
	fixedHeight = 100
	fixedWidth = 100
	opt, err := parseOptions()
	assertions.True(err == nil)
	assertions.Equal(ratio, opt.Ratio)
	assertions.False(fitScreen)
	assertions.False(colored)
	assertions.Equal(fixedWidth, opt.FixedWidth)
	assertions.Equal(fixedHeight, opt.FixedHeight)
}

func TestParseUsage(t *testing.T) {
	usage()
}
