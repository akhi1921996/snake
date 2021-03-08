package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeadFunction(t *testing.T) {
	c := Coordinate{
		X: 1,
		Y: 8,
	}
	s := NewSnake(c)
	assert.True(t, c.Equal(s.Head()))
}

func TestFeedFunc(t *testing.T) {
	c := Coordinate{
		X: 1,
		Y: 8,
	}
	s := NewSnake(c)
	l := s.Length()
	s.Feed()
	l++
	assert.Equal(t, s.Length(), l, "Snake length should increase after feeding")
}

func TestIsBodyOnPos(t *testing.T) {
	c := make([]Coordinate, 5)
	for i := 0; i < 5; i++ {
		c[i] = Coordinate{
			X: i,
			Y: i,
		}
	}
	s := &Snake{
		length: 5,
		body:   c,
	}
	var tests = []struct {
		name           string
		c              Coordinate
		expectedOnBody bool
	}{
		{
			name: "Test OnBody true",
			c: Coordinate{
				X: 1,
				Y: 1,
			},
			expectedOnBody: true,
		},
		{
			name: "Test OnBody false",
			c: Coordinate{
				X: 1,
				Y: 2,
			},
			expectedOnBody: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isOnBody := s.IsBodyOnPosition(tt.c)
			assert.Equal(t, isOnBody, tt.expectedOnBody, "ExpectedOnBody and isOnBody not same")
		})
	}
}

func TestNewHeadPos(t *testing.T) {
	c := Coordinate{
		X: 1,
		Y: 1,
	}
	s := NewSnake(c)
	var tests = []struct {
		name               string
		direction          string
		expectedNewHeadPos Coordinate
	}{
		{
			name:      "Test East Direction",
			direction: "E",
			expectedNewHeadPos: Coordinate{
				X: 2,
				Y: 1,
			},
		},
		{
			name:      "Test West Direction",
			direction: "W",
			expectedNewHeadPos: Coordinate{
				X: 0,
				Y: 1,
			},
		},
		{
			name:      "Test North Direction",
			direction: "N",
			expectedNewHeadPos: Coordinate{
				X: 1,
				Y: 0,
			},
		},
		{
			name:      "Test South Direction",
			direction: "S",
			expectedNewHeadPos: Coordinate{
				X: 1,
				Y: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := s.NewHeadPos(tt.direction)
			assert.Equal(t, h, tt.expectedNewHeadPos,
				"ExpectedHeadPos should be equal to Actual head position")
		})
	}
}

func TestPositionSnakeFunc(t *testing.T) {
	c := Coordinate{
		X: 1,
		Y: 1,
	}
	var tests = []struct {
		name                   string
		coordinate             Coordinate
		Feeded                 bool
		expectedFreeCoordinate Coordinate
		expectedSnakeLength    int
	}{
		{
			name: "Test Snake not feeded",
			coordinate: Coordinate{
				X: 2,
				Y: 1,
			},
			Feeded: false,
			expectedFreeCoordinate: Coordinate{
				X: 1,
				Y: 1,
			},
			expectedSnakeLength: 1,
		},
		{
			name: "Test Snake feeded",
			coordinate: Coordinate{
				X: 2,
				Y: 1,
			},
			Feeded: true,
			expectedFreeCoordinate: Coordinate{
				X: -1,
				Y: -1,
			},
			expectedSnakeLength: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSnake(c)
			if tt.Feeded {
				s.Feed()
			}
			h := s.PositionSnakeAndReturnFreeCoordinate(tt.coordinate)
			assert.Equal(t, h, tt.expectedFreeCoordinate,
				"ExpectedFreeCoordinate should be equal to Actual freed coordinate")
			assert.Equal(t, s.Length(), tt.expectedSnakeLength,
				"ExpectedSnakeLength should be equal to Actual snake length")
		})
	}
}
