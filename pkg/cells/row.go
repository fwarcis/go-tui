package cells

import "github.com/fwarcis/go-tui/pkg/common/sizes"

type Row struct {
	cells []Cell
}

func NewRowByCells(cells []Cell) Row {
	return Row{cells: cells}
}

func NewRow(size sizes.Size[[]Cell, Cell]) Row {
	return Row{
		cells: make(
			[]Cell,
			size.Len(),
			size.Len()+size.NToCap(),
		),
	}
}

func (r Row) Copy() Row {
	newCells := make([]Cell, len(r.cells), cap(r.cells))
	copy(newCells, r.cells)
	return NewRowByCells(newCells)
}

func (r Row) Cell(x int) (Cell, error) {
	if x < 0 || x >= len(r.cells) {
		return Cell{}, newError(
			"Row.Cell", "%w",
			NewInvalidPosError(
				'x', x, len(r.cells),
			),
		)
	}
	return r.cells[x], nil
}

func (r *Row) SetCell(x int, cell Cell) error {
	if x < 0 || x >= len(r.cells) {
		return newError(
			"Row.SetCell", "%w",
			NewInvalidPosError(
				'x', x, len(r.cells),
			),
		)
	}
	r.cells[x] = cell
	return nil
}

func (r Row) Len() int {
	return len(r.cells)
}

func (r Row) Cap() int {
	return cap(r.cells)
}

func (r *Row) SetSize(
	new sizes.Size[[]Cell, Cell],
) {
	r.cells = new.Apply(r.cells)
}
