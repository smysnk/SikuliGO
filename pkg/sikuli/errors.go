package sikuli

import "errors"

var (
	ErrFindFailed         = errors.New("sikuli: find failed")
	ErrTimeout            = errors.New("sikuli: timeout")
	ErrInvalidTarget      = errors.New("sikuli: invalid target")
	ErrBackendUnsupported = errors.New("sikuli: backend unsupported")
)

