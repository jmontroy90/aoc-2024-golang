package day6

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed testGrid.txt
var testGrid string

func TestPart1(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		{name: "basic", expected: 41},
	}
	var grid [][]rune
	for _, s := range strings.Split(testGrid, "\n") {
		grid = append(grid, []rune(s))
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Part1(NewGrid(grid))
			if actual != tt.expected {
				t.Errorf("%s: got %v, expected %v", tt.name, actual, tt.expected)
			}
		})
	}
}

//go:embed testBlockGrid.txt
var testBlockGrid string

func TestPart2(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		{name: "basic", expected: 6},
	}
	var grid [][]rune
	for _, s := range strings.Split(testBlockGrid, "\n") {
		grid = append(grid, []rune(s))
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Part2(NewGrid(grid))
			if actual != tt.expected {
				t.Errorf("%s: got %v, expected %v", tt.name, actual, tt.expected)
			}
		})
	}
}
