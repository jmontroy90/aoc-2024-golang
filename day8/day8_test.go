package day8

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed testGrid.txt
var testGrid string

func TestIsTrueEquation(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		{name: "sample", expected: 14},
	}

	var grid [][]rune
	for _, s := range strings.Split(testGrid, "\n") {
		grid = append(grid, []rune(s))
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Part1(NewGrid(grid))
			if actual != tt.expected {
				t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}
