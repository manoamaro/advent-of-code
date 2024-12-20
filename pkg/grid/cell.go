package grid

type Cell [2]int

type Dir Cell

var (
	Up    = Dir{-1, 0}
	Down  = Dir{1, 0}
	Left  = Dir{0, -1}
	Right = Dir{0, 1}
)

func (c Cell) Move(d Dir) Cell {
	return Cell{c[0] + d[0], c[1] + d[1]}
}

func (d Dir) RotateCW() Dir {
	return Dir{d[1], -d[0]}
}

func (d Dir) RotateCCW() Dir {
	return Dir{-d[1], d[0]}
}
