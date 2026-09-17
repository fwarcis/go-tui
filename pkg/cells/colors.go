package cells

import (
	"image/color"

	"github.com/charmbracelet/x/ansi"
)

var MaxForegroundANSILen = len(
	ansi.NewStyle().BackgroundColor(color.White).String(),
)

type Foreground struct {
	color.Color
}

func NewForeground(inner color.Color) Foreground {
	return Foreground{Color: inner}
}

func (c Foreground) ANSI() string {
	if c.Color == nil {
		return ""
	}
	return ansi.NewStyle().ForegroundColor(c.Color).String()
}

var MaxBackgroundANSILen = len(
	ansi.NewStyle().BackgroundColor(color.White).String(),
)

type Background struct {
	color.Color
}

func NewBackground(inner color.Color) Background {
	return Background{Color: inner}
}

func (c Background) ANSI() string {
	if c.Color == nil {
		return ""
	}
	return ansi.NewStyle().BackgroundColor(c.Color).String()
}
