package ascii

// Options convert pixel to raw char
type Options struct {
	Pixels  []byte
	Reverse bool
	Colored bool
}

// DefaultOptions that contains the default pixels
var DefaultOptions = Options{
	Pixels:  []byte(" .,:;i1tfLCG08@"),
	Reverse: false,
	Colored: true,
}

// NewOptions create a new convert option
func NewOptions() Options {
	return DefaultOptions
}

// mergeOptions merge two options
func (options *Options) mergeOptions(newOptions *Options) {
	options.Pixels = append([]byte{}, newOptions.Pixels...)
	options.Reverse = newOptions.Reverse
	options.Colored = newOptions.Colored
}
