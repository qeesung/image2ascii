package ascii

// Convert options
type Options struct {
	Pixels  []byte
	Reverse bool
	Colored bool
}

// Default options
var DefaultOptions = Options{
	Pixels:  []byte(" .,:;i1tfLCG08@"),
	Reverse: false,
	Colored: true,
}

// Create a new options
func NewOptions() Options {
	return DefaultOptions
}

// Merge options
func (options *Options) mergeOptions(newOptions *Options) {
	options.Pixels = append([]byte{}, newOptions.Pixels...)
	options.Reverse = newOptions.Reverse
	options.Colored = newOptions.Colored
}
