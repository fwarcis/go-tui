package sizes

import "fmt"

type Pos2 struct {
	X, Y int
}

func (p Pos2) String() string {
	return fmt.Sprintf(
		"Pos2{X: %d, Y: %d}",
		p.X, p.Y,
	)
}
