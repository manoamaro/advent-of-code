package grid

import "fmt"

type Cell [2]int

type Dir Cell

var (
	Up        = Dir{-1, 0}
	UpRight   = Dir{-1, 1}
	Right     = Dir{0, 1}
	RightDown = Dir{1, 1}
	Down      = Dir{1, 0}
	DownLeft  = Dir{1, -1}
	Left      = Dir{0, -1}
	LeftUp    = Dir{-1, -1}
)

var CardinalDirections = [4]Dir{Up, Left, Down, Right}
var DiagonalDirections = [4]Dir{UpRight, RightDown, DownLeft, LeftUp}
var AllDirections = [8]Dir{Up, UpRight, Right, RightDown, Down, DownLeft, Left, LeftUp}

func NewCell(x, y int) Cell {
	return Cell{x, y}
}

func (c *Cell) Move(d Dir) Cell {
	return Cell{c[0] + d[0], c[1] + d[1]}
}

func (d Dir) RotateCW() Dir {
	return Dir{d[1], -d[0]}
}

func (d Dir) RotateCCW() Dir {
	return Dir{-d[1], d[0]}
}

func (c *Cell) String() string {
	return fmt.Sprintf("(%d, %d)", c[0], c[1])
}

func NewDir(x, y int) Dir {
	return Dir{x, y}
}
