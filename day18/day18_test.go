package day18

import (
	_ "embed"
	"github.com/jmontroy90/aoc-2024/util"
	"testing"
)

func TestBFS(t *testing.T) {
	grid, _ := util.NewGridFromFile("testdata/small.txt")
	expected := 22
	t.Run("small", func(t *testing.T) {
		got := BFS(grid, util.XY{0, 0}, util.XY{6, 6}, true)
		if got != expected {
			t.Errorf("small, expected %v, got %v", expected, got)
		}
		grid.Print()
	})
}
