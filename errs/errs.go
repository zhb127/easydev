package errs

import "github.com/pkg/errors"

var ErrInterrupt = errors.New("^C")
