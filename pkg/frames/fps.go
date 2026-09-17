package frames

import (
	"fmt"
	"time"
)

type FPS int

func NewFPS(val int) (FPS, error) {
	if val < 0 {
		return 0, newError(
			"NewFPS", "%w",
			NewNegativeFPSError(val),
		)
	}
	return FPS(val), nil
}

type NegativeFPSError struct {
	Value int
}

func NewNegativeFPSError(val int) NegativeFPSError {
	return NegativeFPSError{Value: val}
}

func (e NegativeFPSError) Error() string {
	return fmt.Sprintf(
		"fps out of range [%d]", e.Value,
	)
}

func (f FPS) Duration() time.Duration {
	if f == 0 {
		return time.Duration(-1)
	}
	return time.Second / time.Duration(f)
}

func (f FPS) NewTicker() *time.Ticker {
	if f == 0 {
		return nil
	}
	return time.NewTicker(f.Duration())
}
