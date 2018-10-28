package terminal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTerminalAccessor(t *testing.T) {
	assertions := assert.New(t)
	accessor := NewTerminalAccessor()
	assertions.True(accessor != nil)
}

func TestAccessor_CharWidth(t *testing.T) {
	assertions := assert.New(t)
	accessor := NewTerminalAccessor()
	charWidth := accessor.CharWidth()
	if accessor.IsWindows() {
		assertions.Equal(charWidthWindows, charWidth)
	} else {
		assertions.Equal(charWidthOther, charWidth)
	}
}

func TestAccessor_ScreenSize(t *testing.T) {
	assertions := assert.New(t)
	accessor := NewTerminalAccessor()
	_, _, err := accessor.ScreenSize()
	assertions.True(err != nil)
}
