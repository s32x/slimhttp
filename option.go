package slimhttp

// An Option is a step we would like to take surrounding
// the execution of an Endpoint function
type Option int

const (
	// OptionURLSigning enables URL signature verification
	OptionURLSigning Option = iota

	// OptionCache enables Caching using the cache on the
	// Router
	OptionCache
)
