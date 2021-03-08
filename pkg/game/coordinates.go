package game

// Coordinate represent the position of an object on the board
type Coordinate struct {
	X int
	Y int
}

// Equal checks whether the given coordinate is same or not
func (c *Coordinate) Equal(d Coordinate) bool {
	return c.X == d.X && c.Y == d.Y
}
