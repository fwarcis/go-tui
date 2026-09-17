package screens

import (
	"io"

	"github.com/charmbracelet/x/ansi"
)

type Cursor struct {
	rw io.ReadWriter
}

func NewCursor(rw io.ReadWriter) Cursor {
	return Cursor{rw: rw}
}

func (c *Cursor) Show() (written int, err error) {
	return c.rw.Write([]byte(ansi.ShowCursor))
}

func (c *Cursor) Hide() (written int, err error) {
	return c.rw.Write([]byte(ansi.HideCursor))
}

func (c *Cursor) SetPosition(x, y int) (written int, err error) {
	return c.rw.Write([]byte(ansi.CursorPosition(x, y)))
}

func (c *Cursor) Home() (written int, err error) {
	return c.rw.Write([]byte(ansi.CursorHomePosition))
}

func (c *Cursor) Up(n int) (written int, err error) {
	return c.rw.Write([]byte(ansi.CursorUp(n)))
}

func (c *Cursor) Down(n int) (written int, err error) {
	return c.rw.Write([]byte(ansi.CursorDown(n)))
}

func (c *Cursor) Forward(n int) (written int, err error) {
	return c.rw.Write([]byte(ansi.CursorForward(n)))
}

func (c *Cursor) Backward(n int) (written int, err error) {
	return c.rw.Write([]byte(ansi.CursorBackward(n)))
}

func (c *Cursor) Save() (written int, err error) {
	return c.rw.Write([]byte(ansi.SaveCursor))
}

func (c *Cursor) Restore() (written int, err error) {
	return c.rw.Write([]byte(ansi.RestoreCursor))
}
