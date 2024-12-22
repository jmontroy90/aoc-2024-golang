package day14

import (
	"github.com/jmontroy90/aoc-2024/util"
	"testing"
)

func TestRobot_Step(t *testing.T) {
	tests := []struct {
		name         string
		robot        Robot
		gridX, gridY int
		steps        int
		expected     util.XY
	}{
		{"sample-1", Robot{CurrPos: util.XY{X: 2, Y: 4}, Velocity: util.XY{X: 2, Y: -3}}, 11, 7, 1, util.XY{4, 1}},
		{"sample-5", Robot{CurrPos: util.XY{X: 2, Y: 4}, Velocity: util.XY{X: 2, Y: -3}}, 11, 7, 5, util.XY{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.robot.Step(tt.steps, tt.gridX, tt.gridY)
			if actual != tt.expected {
				t.Errorf("%v: expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		name         string
		robots       []Robot
		gridX, gridY int
		steps        int
		expected     int
	}{
		{
			"sample",
			[]Robot{
				{CurrPos: util.XY{X: 0, Y: 4}, Velocity: util.XY{X: 3, Y: -3}},
				{CurrPos: util.XY{X: 6, Y: 3}, Velocity: util.XY{X: -1, Y: -3}},
				{CurrPos: util.XY{X: 10, Y: 3}, Velocity: util.XY{X: -1, Y: 2}},
				{CurrPos: util.XY{X: 2, Y: 0}, Velocity: util.XY{X: 2, Y: -1}},
				{CurrPos: util.XY{X: 0, Y: 0}, Velocity: util.XY{X: 1, Y: 3}},
				{CurrPos: util.XY{X: 3, Y: 0}, Velocity: util.XY{X: -2, Y: -2}},
				{CurrPos: util.XY{X: 7, Y: 6}, Velocity: util.XY{X: -1, Y: -3}},
				{CurrPos: util.XY{X: 3, Y: 0}, Velocity: util.XY{X: -1, Y: -2}},
				{CurrPos: util.XY{X: 9, Y: 3}, Velocity: util.XY{X: 2, Y: 3}},
				{CurrPos: util.XY{X: 7, Y: 3}, Velocity: util.XY{X: -1, Y: 2}},
				{CurrPos: util.XY{X: 2, Y: 4}, Velocity: util.XY{X: 2, Y: -3}},
				{CurrPos: util.XY{X: 9, Y: 5}, Velocity: util.XY{X: -3, Y: -3}},
			},
			11, 7, 100, 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Part1(tt.robots, tt.gridX, tt.gridY)
			if actual != tt.expected {
				t.Errorf("%v: expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}
