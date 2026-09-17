package frames

import "github.com/fwarcis/go-tui/internal/errs"

func newError(function string, format string, a ...any) error {
	return errs.Errorf("frames."+function, format, a...)
}
