package file

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrScanBasePathNotDir = fmt.Errorf("basePath 不是一个目录")

func NewErrScanBasePathNotDir() error {
	return errors.WithStack(ErrScanBasePathNotDir)
}
