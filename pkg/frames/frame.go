package frames

import (
	"errors"

	"github.com/charmbracelet/x/ansi"
	"github.com/fwarcis/go-tui/internal/asrts"
	"github.com/fwarcis/go-tui/pkg/cells"
	"github.com/fwarcis/go-tui/pkg/common/arrs"
	"github.com/fwarcis/go-tui/pkg/common/sizes"
)

type Mutable interface {
	SetWidth(size cells.WidthSize)
	SetHeight(size cells.HeightSize)
	SetCell(x, y int, cell cells.Cell) error
}

type Immutable interface {
	Width() int
	Height() int
	Cell(x, y int) (cells.Cell, error)
	ChangedCells() arrs.Immutable[PosCell]
}

type _grid = cells.Grid

type Frame struct {
	_grid

	commitingData commitingData

	notes string
}

func New(grid cells.Grid, commitingCellsCap int) Frame {
	return Frame{
		_grid: grid,
		commitingData: newCommitingData(
			commitingCellsCap,
			asrts.MayBeNil(cells.NewWidthSize(grid.Width(), 0)),
			asrts.MayBeNil(cells.NewHeightSize(grid.Height(), 0)),
		),
	}
}

type commitingData struct {
	PreviousCells    []PosCell
	CurrentCells     []PosCell
	LastWidth        cells.WidthSize
	LastHeight       cells.HeightSize
	LastChangesCount int
}

func newCommitingData(
	cellsCap int,
	width cells.WidthSize,
	height cells.HeightSize,
) commitingData {
	return commitingData{
		PreviousCells: make([]PosCell, 0, cellsCap),
		CurrentCells:  make([]PosCell, 0, cellsCap),
		LastWidth:     width,
		LastHeight:    height,
	}
}

type PosCell struct {
	cells.Cell
	Pos sizes.Pos2
}

func NewPosCell(x, y int, cell cells.Cell) PosCell {
	return PosCell{
		Pos:  sizes.Pos2{X: x, Y: y},
		Cell: cell,
	}
}

func (c PosCell) AppendText(b []byte) ([]byte, error) {
	b = append(b, ansi.CursorPosition(
		c.Pos.X, c.Pos.Y,
	)[:]...)
	b, err := c.Cell.AppendText(b)
	return b, err
}

func (c PosCell) CopyText(b []byte) (copied int) {
	copied = copy(b, ansi.CursorPosition(c.Pos.X, c.Pos.Y))
	copied += c.Cell.CopyText(b[copied:])
	return copied
}

func (f Frame) ChangedCells() arrs.Immutable[PosCell] {
	return arrs.NewImmutable(
		f.commitingData.PreviousCells,
	)
}

func (f *Frame) Commit() error {
	f._grid.SetHeight(f.commitingData.LastHeight)
	f._grid.SetWidth(f.commitingData.LastWidth)

	var err error
	for _, cell := range f.commitingData.PreviousCells {
		err = errors.Join(err, f._grid.SetCell(
			cell.Pos.X, cell.Pos.Y, cell.Cell,
		))
	}
	if err != nil {
		return err
	}

	prevCells := f.commitingData.PreviousCells
	f.commitingData.PreviousCells = f.commitingData.CurrentCells
	f.commitingData.CurrentCells = prevCells[:0]

	return nil
}

func (f *Frame) SetCell(x, y int, new cells.Cell) error {
	_, err := f.Cell(x, y)
	if err != nil {
		return err
	}

	f.commitingData.CurrentCells = append(f.commitingData.CurrentCells,
		NewPosCell(x, y, new))

	return nil
}

func (f *Frame) SetWidth(size cells.WidthSize) {
	f.commitingData.LastWidth = size
}

func (f *Frame) SetHeight(size cells.HeightSize) {
	f.commitingData.LastHeight = size
}

func (f Frame) Copy() Frame {
	return New(
		f._grid.Copy(),
		cap(f.commitingData.PreviousCells),
	)
}
