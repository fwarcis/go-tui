package cells

import (
	"fmt"

	"github.com/fwarcis/go-tui/internal/errs"
)

func newError(function string, format string, a ...any) error {
	return errs.Errorf("cells."+function, format, a...)
}

type InvalidPosError struct {
	Current int
	Max     int
	Axis    rune
}

func NewInvalidPosError(axis rune, cur int, max int) InvalidPosError {
	return InvalidPosError{Current: cur, Axis: axis}
}

func (e InvalidPosError) Error() string {
	return fmt.Sprintf(
		"pos %c out of range [%d] with max %d",
		e.Axis, e.Current, e.Max,
	)
}
