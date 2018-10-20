package ascii

// Convert options
type Options struct {
	pixels  []byte
	reverse bool
	colored bool
	bg      bool
	fg      bool
}

// Default options
var DefaultOptions = Options{
	pixels:  []byte(" .,:;i1tfLCG08@"),
	reverse: false,
	colored: true,
	bg:      false,
	fg:      true,
}

// Create a new options
func NewOptions() Options {
	return DefaultOptions
}

// Merge options
func (options *Options) mergeOptions(newOptions *Options) {
	options.pixels = append([]byte{}, newOptions.pixels...)
	options.reverse = newOptions.reverse
	options.colored = newOptions.colored
	options.bg = newOptions.bg
	options.fg = newOptions.fg
}
