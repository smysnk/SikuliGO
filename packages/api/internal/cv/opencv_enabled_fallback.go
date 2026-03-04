//go:build !opencv

package cv

// OpenCVEnabled reports whether this binary was built with the "opencv" build tag.
func OpenCVEnabled() bool {
	return false
}
