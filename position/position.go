package position

import intmath "aoc2019/mymath"

type Pos struct {
	X, Y int
}

func (p *Pos) ManhattanDistance(other Pos) int {
	return intmath.Abs(other.X-p.X) +
		intmath.Abs(other.Y-p.Y)
}
