package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoordinateEqual(t *testing.T) {
	var tests = []struct {
		name             string
		firstCoordinate  Coordinate
		secondCoordinate Coordinate
		expectedEqual    bool
	}{
		{
			name: "Equal Test",
			firstCoordinate: Coordinate{
				X: 2,
				Y: 3,
			},
			secondCoordinate: Coordinate{
				X: 2,
				Y: 3,
			},
			expectedEqual: true,
		},
		{
			name: "Not Equal Test",
			firstCoordinate: Coordinate{
				X: 2,
				Y: 4,
			},
			secondCoordinate: Coordinate{
				X: 2,
				Y: 3,
			},
			expectedEqual: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			equal := tt.firstCoordinate.Equal(tt.secondCoordinate)
			assert.Equal(t, equal, tt.expectedEqual, "Equal and Expected Equal should be same")
		})
	}
}
