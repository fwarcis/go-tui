package screens

import (
	"errors"

	"github.com/fwarcis/go-tui/pkg/frames"
)

type WriterResponse[RespValue any] struct {
	Value RespValue
	Err   error
}

var ErrFrameBusy = errors.New("frame busy")

func (s *Screen[WriterRespValue]) LoopRendering(
	nextFrame func(it Iteration[WriterRespValue]) bool,
) error {
	if !s.loopsRendering.CompareAndSwap(false, true) {
		return newError(
			"Screen.LoopRendering",
			"%w", ErrFrameBusy,
		)
	}

	for s.loopsRendering.Load() {
		s.frame.Lock()
		s.frame.WaitForChanging()

		s.frame.Sleep()
		respVal, err := s.w.WriteFrame(s.frame)
		it := NewIteration(&s.frame, respVal, err)
		s.loopsRendering.Store(nextFrame(it))

		s.frame.SignalToChange()
		s.frame.Unlock()
	}

	return nil
}

type Iteration[WriterRespValue any] struct {
	Resp  WriterResponse[WriterRespValue]
	Frame frames.Immutable
}

func NewIteration[WriterRespValue any](
	frame frames.Immutable,
	val WriterRespValue,
	err error,
) Iteration[WriterRespValue] {
	return Iteration[WriterRespValue]{
		Frame: frame,
		Resp: WriterResponse[WriterRespValue]{
			Value: val,
			Err:   err,
		},
	}
}
