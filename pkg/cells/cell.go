package cells

import (
	"fmt"
)

type Cell struct {
	Foreground Foreground
	Background Background
	Char       string
	Styles     Style
}

func (c Cell) String() string {
	return fmt.Sprintf(
		"Cell{Char: %q, Foreground: %v, Background: %v, Styles: %v}",
		c.Char, c.Foreground, c.Background, c.Styles,
	)
}

var MaxCellANSILen = (MaxForegroundANSILen +
	MaxBackgroundANSILen +
	MaxStylesANSILen)

func (c Cell) ANSI() string {
	buf := make([]byte, MaxCellANSILen+len(c.Char))
	copied := c.CopyText(buf)
	return string(buf[:copied])
}

func (c Cell) AppendText(b []byte) ([]byte, error) {
	b = append(b, c.Foreground.ANSI()[:]...)
	b = append(b, c.Background.ANSI()[:]...)
	b = append(b, c.Styles.ANSI()[:]...)
	return append(b, c.Char[:]...), nil
}

func (c Cell) CopyText(b []byte) (copied int) {
	copied = copy(b, c.Foreground.ANSI())
	copied += copy(b[copied:], c.Background.ANSI())
	copied += copy(b[copied:], c.Styles.ANSI())
	copied += copy(b[copied:], c.Char)
	return copied
}
