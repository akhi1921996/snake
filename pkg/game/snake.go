package game

// Snake holds the attributes of the snake in the game
type Snake struct {
	body   []Coordinate
	length int
}

// NewSnake returns a snake with its head at given coordinate
func NewSnake(c Coordinate) *Snake {
	return &Snake{
		body: []Coordinate{
			c,
		},
		length: 1,
	}
}

// Head returns the current coordinate of the snake head
func (s *Snake) Head() Coordinate {
	return s.body[len(s.body)-1]
}

// Length returns the current length of the snake
func (s *Snake) Length() int {
	return s.length
}

// Feed will increase the length of the snake by 1
func (s *Snake) Feed() {
	s.length++
}

// IsBodyOnPosition returns if the snake has its body on the given coordinate position
func (s *Snake) IsBodyOnPosition(pos Coordinate) bool {
	for _, coordinate := range s.body {
		if pos.X == coordinate.X && pos.Y == coordinate.Y {
			return true
		}
	}
	return false
}

// NewHeadPos calculates and returns the new position is snake moves in the given direction
func (s *Snake) NewHeadPos(direction string) Coordinate {
	h := s.Head()
	switch direction {
	case "E":
		h.X++
	case "W":
		h.X--
	case "S":
		h.Y++
	case "N":
		h.Y--
	}
	return h
}

// PositionSnakeAndReturnFreeCoordinate will position snake head on the given coordinate and
// will adjust snake body. While adjusting snake body, if previously occupies coordinate is freed,
// then that coordinate will be returned. Else a dummy Coordinate{X:-1,Y:-1} is returned.
func (s *Snake) PositionSnakeAndReturnFreeCoordinate(newHeadPos Coordinate) Coordinate {
	s.body = append(s.body, newHeadPos)
	c := Coordinate{
		X: -1,
		Y: -1,
	}
	if s.Length() < len(s.body) {
		c = s.body[0]
		s.body = s.body[1:]
	}
	return c
}
