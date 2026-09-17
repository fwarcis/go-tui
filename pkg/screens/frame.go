package screens

import (
	"sync"
	"time"

	"github.com/fwarcis/go-tui/pkg/frames"
)

type (
	_mu         = sync.Mutex
	_innerFrame = frames.Frame
)

type Frame struct {
	_innerFrame
	fps frames.FPS

	*_mu
	writingCond    *sync.Cond
	changingCond   *sync.Cond
	needsRendering bool
}

func NewFrame(inner frames.Frame, fps frames.FPS) Frame {
	mu := &sync.Mutex{}
	return Frame{
		_innerFrame:  inner,
		fps:          fps,
		_mu:          mu,
		writingCond:  sync.NewCond(mu),
		changingCond: sync.NewCond(mu),
	}
}

func (f *Frame) WaitForChanging() {
	for !f.needsRendering {
		f.writingCond.Wait()
	}
}

func (f *Frame) SignalToWrite() {
	f.needsRendering = true
	f.writingCond.Signal()
}

func (f *Frame) WaitForWriting() {
	for f.needsRendering {
		f.changingCond.Wait()
	}
}

func (f *Frame) SignalToChange() {
	f.needsRendering = false
	f.changingCond.Signal()
}

func (f Frame) Sleep() {
	time.Sleep(f.fps.Duration())
}

func (f Frame) FPS() frames.FPS {
	return f.fps
}

func (f *Frame) SetFPS(fps frames.FPS) {
	f.fps = fps
}
