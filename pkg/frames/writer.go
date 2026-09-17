package frames

import (
	"io"
)

type Writer struct {
	w   io.Writer
	buf []byte
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w:   w,
		buf: make([]byte, 200000),
	}
}

func (w *Writer) WriteFrame(frame Immutable) (int, error) {
	var copied int
	for cell := range frame.ChangedCells().Elements {
		copied += cell.CopyText(w.buf[copied:])
	}
	return w.w.Write(w.buf[:copied])
}
