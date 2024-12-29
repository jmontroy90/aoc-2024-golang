package day15

import (
	_ "embed"
	"github.com/jmontroy90/aoc-2024/util"
	"testing"
)

var (
	largeMoves = "<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^"
)

func TestPart1(t *testing.T) {
	small, _ := util.NewGridFromFile("testdata/testSmallGrid.txt")
	large, _ := util.NewGridFromFile("testdata/testLargeGrid.txt")

	tests := []struct {
		name     string
		grid     *util.Grid
		moves    string
		expected int
	}{
		{"small", small, "<^^>>>vv<v>>v<<", 2028},
		{
			"large",
			large,
			largeMoves,
			10092,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.grid.Reset()
			actual := Part1(tt.grid, dirsToXYs(tt.moves))
			if actual != tt.expected {
				t.Errorf("%v: expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}

func TestWidenGrid(t *testing.T) {
	wide, _ := util.NewGridFromFile("testdata/testWidenGrid.txt")
	expected, _ := util.NewGridFromFile("testdata/testWidenGrid_Output.txt")
	t.Run("WidenGrid", func(t *testing.T) {
		actual := WidenGrid(wide)
		if expected.PrintString() != actual.PrintString() {
			t.Errorf("WidenGrid: grids not the same!\nExpected:\n%v\nGot:\n%v\n", expected.PrintString(), actual.PrintString())
		}
	})
}

func TestPart2(t *testing.T) {
	orig, _ := util.NewGridFromFile("testdata/testLargeGrid.txt")
	expectedGrid, _ := util.NewGridFromFile("testdata/testLargeWidenGrid_Output.txt")
	expected := 9021

	t.Run("Part2", func(t *testing.T) {
		actualGrid, actual := Part2(orig, dirsToXYs(largeMoves))
		if expectedGrid.PrintString() != actualGrid.PrintString() {
			t.Errorf("Part2: grids not the same!\nExpected:\n%v\nGot:\n%v\n", expectedGrid.PrintString(), actualGrid.PrintString())
		}
		if actual != expected {
			t.Errorf("%v: expected %v, got %v", "widen", expected, actual)
		}
	})
}
