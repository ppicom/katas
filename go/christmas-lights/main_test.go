package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid_TurnOn(t *testing.T) {
	testCases := []struct {
		gridSize Coordinate
		turnOn   [2]Coordinate
		expected int
	}{
		{Coordinate{9, 9},
			[2]Coordinate{{0, 0}, {0, 9}}, 10},
		{Coordinate{999, 999},
			[2]Coordinate{{0, 0}, {999, 999}}, 1000000},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Turn on %v, %v in grid of %v", tc.turnOn[0], tc.turnOn[1], tc.gridSize), func(t *testing.T) {
			grid := NewGrid(tc.gridSize)

			grid.TurnOn(tc.turnOn[0], tc.turnOn[1])
			n := grid.LightsOn()

			assert.Equal(t, tc.expected, n)
		})
	}
}

func TestGrid_TurnOff(t *testing.T) {
	testCases := []struct {
		gridSize Coordinate
		turnOn   [2]Coordinate
		expected int
	}{
		{Coordinate{9, 9},
			[2]Coordinate{{0, 0}, {0, 9}}, 90},
		{Coordinate{999, 999},
			[2]Coordinate{{0, 0}, {999, 999}}, 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Turn on %v, %v in grid of %v", tc.turnOn[0], tc.turnOn[1], tc.gridSize), func(t *testing.T) {
			grid := NewGrid(tc.gridSize)
			grid.TurnOn(Coordinate{}, tc.gridSize) // turn all lights on

			grid.TurnOff(tc.turnOn[0], tc.turnOn[1])
			n := grid.LightsOn()

			assert.Equal(t, tc.expected, n)
		})
	}
}

func TestGrid_Toggle(t *testing.T) {
	grid := NewGrid(Coordinate{9, 9})
	grid.TurnOn(Coordinate{}, Coordinate{0, 9})
	assert.Equal(t, 10, grid.LightsOn())

	grid.Toggle(Coordinate{}, Coordinate{9, 9})

	assert.Equal(t, 90, grid.LightsOn())
}
