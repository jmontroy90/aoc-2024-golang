package day16

import (
	_ "embed"
	"github.com/jmontroy90/aoc-2024/util"
	"testing"
)

func TestPart1(t *testing.T) {
	grid, _ := util.NewGridFromFile("testdata/testGrid.txt")
	grid2, _ := util.NewGridFromFile("testdata/testGrid2.txt")
	t.Run("small", func(t *testing.T) {
		expected := 7036
		start, _ := grid.ScanOnce('S')
		got := FindPathBFS(grid, start, util.Right)
		if got != expected {
			t.Errorf("small: expected %v, got %v", expected, got)
		}
	})
	t.Run("bigger", func(t *testing.T) {
		expected := 11048
		start, _ := grid2.ScanOnce('S')
		got := FindPathBFS(grid2, start, util.Right)
		if got != expected {
			t.Errorf("small: expected %v, got %v", expected, got)
		}
	})
}

func TestPart2(t *testing.T) {
	grid, _ := util.NewGridFromFile("testdata/testGrid.txt")
	grid2, _ := util.NewGridFromFile("testdata/testGrid2.txt")
	t.Run("small", func(t *testing.T) {
		expected := 45
		start, _ := grid.ScanOnce('S')
		end, _ := grid.ScanOnce('E')
		got := FindPathWithMemory(grid, start, util.Right, end)
		if got != expected {
			t.Errorf("small: expected %v, got %v", expected, got)
		}
	})
	t.Run("bigger", func(t *testing.T) {
		expected := 64
		start, _ := grid2.ScanOnce('S')
		end, _ := grid2.ScanOnce('E')
		got := FindPathWithMemory(grid2, start, util.Right, end)
		if got != expected {
			t.Errorf("small: expected %v, got %v", expected, got)
		}
	})
}
