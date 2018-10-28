package ascii

// Options convert pixel to raw char
type Options struct {
	Pixels   []byte
	Reversed bool
	Colored  bool
}

// DefaultOptions that contains the default pixels
var DefaultOptions = Options{
	Pixels:   []byte(" .,:;i1tfLCG08@"),
	Reversed: false,
	Colored:  true,
}

// NewOptions create a new convert option
func NewOptions() Options {
	newOptions := Options{}
	newOptions.mergeOptions(&DefaultOptions)
	return newOptions
}

// mergeOptions merge two options
func (options *Options) mergeOptions(newOptions *Options) {
	options.Pixels = append([]byte{}, newOptions.Pixels...)
	options.Reversed = newOptions.Reversed
	options.Colored = newOptions.Colored
}
