package screens

import (
	"sync/atomic"

	"github.com/fwarcis/go-tui/pkg/cells"
	"github.com/fwarcis/go-tui/pkg/frames"
)

type FrameWriter[RespValue any] interface {
	WriteFrame(frame frames.Immutable) (val RespValue, err error)
}

type Screen[WriterRespValue any] struct {
	w FrameWriter[WriterRespValue]

	frame Frame

	loopsRendering atomic.Bool
}

func NewScreen[WriterRespValue any](
	w FrameWriter[WriterRespValue],
	commitingCellsCap int,
	width cells.WidthSize,
	height cells.HeightSize,
	fps frames.FPS,
) Screen[WriterRespValue] {
	return Screen[WriterRespValue]{
		w: w,
		frame: NewFrame(frames.New(
			cells.NewGrid(width, height),
			commitingCellsCap,
		), fps),
	}
}
