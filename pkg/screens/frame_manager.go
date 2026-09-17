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

type FrameManager struct {
	_innerFrame
	fps frames.FPS

	*_mu
	writingCond    *sync.Cond
	changingCond   *sync.Cond
	needsRendering bool
}

func NewFrame(inner frames.Frame, fps frames.FPS) FrameManager {
	mu := &sync.Mutex{}
	return FrameManager{
		_innerFrame:  inner,
		fps:          fps,
		_mu:          mu,
		writingCond:  sync.NewCond(mu),
		changingCond: sync.NewCond(mu),
	}
}

func (f *FrameManager) WaitForChanging() {
	for !f.needsRendering {
		f.writingCond.Wait()
	}
}

func (f *FrameManager) SignalToWrite() {
	f.needsRendering = true
	f.writingCond.Signal()
}

func (f *FrameManager) WaitForWriting() {
	for f.needsRendering {
		f.changingCond.Wait()
	}
}

func (f *FrameManager) SignalToChange() {
	f.needsRendering = false
	f.changingCond.Signal()
}

func (f FrameManager) Sleep() {
	time.Sleep(f.fps.Duration())
}

func (f FrameManager) FPS() frames.FPS {
	return f.fps
}

func (f *FrameManager) SetFPS(fps frames.FPS) {
	f.fps = fps
}
