package cells

import "github.com/fwarcis/go-tui/pkg/common/sizes"

type WidthSize = sizes.Size[[]Cell, Cell]

func NewWidthSize(len, nToCap int) (WidthSize, error) {
	size, err := sizes.NewSize[[]Cell](len, nToCap)
	if err != nil {
		return WidthSize{}, newError(
			"NewWidthSize", "%w", err,
		)
	}

	return size, nil
}

type HeightSize = sizes.Size[[]Row, Row]

func NewHeightSize(len, nToCap int) (HeightSize, error) {
	size, err := sizes.NewSize[[]Row](len, nToCap)
	if err != nil {
		return HeightSize{}, newError(
			"NewHeightSize", "%w", err,
		)
	}
	return size, nil
}

type Grid struct {
	rows []Row
}

func NewGridByRows(rows []Row) Grid {
	return Grid{rows: rows}
}

func NewGrid(
	width WidthSize,
	height HeightSize,
) Grid {
	rows := make(
		[]Row, height.Len()+height.NToCap(),
	)
	for y := range rows {
		rows[y] = NewRow(width)
	}
	return NewGridByRows(
		rows[:height.Len()],
	)
}

func (g Grid) Copy() Grid {
	newRows := make([]Row, len(g.rows), cap(g.rows))
	for y := range g.rows {
		newRows[y] = g.rows[y].Copy()
	}
	return NewGridByRows(newRows)
}

func (g Grid) Cell(x, y int) (Cell, error) {
	if y < 0 || y >= len(g.rows) {
		return Cell{}, newError(
			"Grid.Cell", "%w",
			NewInvalidPosError('y', y, len(g.rows)),
		)
	}
	return g.rows[y].Cell(x)
}

func (g Grid) SetCell(x, y int, cell Cell) error {
	if y < 0 || y >= len(g.rows) {
		return newError(
			"Grid.SetCell", "%w",
			NewInvalidPosError('y', y, len(g.rows)),
		)
	}
	return g.rows[y].SetCell(x, cell)
}

func (g Grid) Width() int {
	if cap(g.rows) == 0 {
		return 0
	}
	return g.rows[:1][0].Len()
}

func (g Grid) Height() int {
	return len(g.rows)
}

func (g Grid) SetWidth(size WidthSize) {
	if cap(g.rows) == 0 {
		return
	}
	allRows := g.rows[:cap(g.rows)]
	for y := range allRows {
		allRows[y].SetSize(size)
	}
}

func (g *Grid) SetHeight(size HeightSize) {
	g.rows = size.Apply(g.rows)
}

func (g Grid) Cap() int {
	return cap(g.rows)
}
