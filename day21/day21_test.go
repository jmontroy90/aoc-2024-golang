package day21

import (
	"fmt"
	"reflect"
	"slices"
	"testing"

	"github.com/jmontroy90/aoc-2024/util"
)

func TestDirectionsForDigit(t *testing.T) {
	type args struct {
		start  rune
		target rune
	}
	tests := []struct {
		name string
		args args
		want [][]util.Direction
	}{
		{"1-2", args{'1', '2'}, [][]util.Direction{{util.Right}}},
		{"1-3", args{'1', '3'}, [][]util.Direction{{util.Right, util.Right}}},
		{"1-4", args{'1', '4'}, [][]util.Direction{{util.Up}}},
		{"1-5", args{'1', '5'}, [][]util.Direction{{util.Up, util.Right}, {util.Right, util.Up}}},
		{"1-6", args{'1', '6'}, [][]util.Direction{
			{util.Up, util.Right, util.Right},
			{util.Right, util.Up, util.Right},
			{util.Right, util.Right, util.Up},
		}},
		{"1-9", args{'1', '9'}, [][]util.Direction{
			{util.Up, util.Up, util.Right, util.Right},
			{util.Right, util.Up, util.Up, util.Right},
			{util.Right, util.Right, util.Up, util.Up},
			{util.Up, util.Right, util.Right, util.Up},
			{util.Up, util.Right, util.Up, util.Right},
			{util.Right, util.Up, util.Right, util.Up},
		}},
		{"9-1", args{'9', '1'}, [][]util.Direction{
			{util.Down, util.Down, util.Left, util.Left},
			{util.Left, util.Down, util.Down, util.Left},
			{util.Left, util.Left, util.Down, util.Down},
			{util.Down, util.Left, util.Left, util.Down},
			{util.Down, util.Left, util.Down, util.Left},
			{util.Left, util.Down, util.Left, util.Down},
		}},
		{"3-7", args{'3', '7'}, [][]util.Direction{
			{util.Up, util.Up, util.Left, util.Left},
			{util.Up, util.Left, util.Up, util.Left},
			{util.Up, util.Left, util.Left, util.Up},
			{util.Left, util.Up, util.Up, util.Left},
			{util.Left, util.Up, util.Left, util.Up},
			{util.Left, util.Left, util.Up, util.Up},
		}},
		{"0-9", args{'0', '9'}, [][]util.Direction{
			{util.Up, util.Up, util.Up, util.Right},
			{util.Right, util.Up, util.Up, util.Up},
			{util.Up, util.Right, util.Up, util.Up},
			{util.Up, util.Up, util.Right, util.Up},
		}},
		{"A-9", args{'A', '9'}, [][]util.Direction{{util.Up, util.Up, util.Up}}},
		// 10 total for our level of duplication - 5! / (3! * 2!) == 120 / (6 * 2) == 10
		{"A-7", args{'A', '7'}, [][]util.Direction{
			{util.Up, util.Up, util.Up, util.Left, util.Left},
			{util.Up, util.Up, util.Left, util.Up, util.Left},
			{util.Up, util.Up, util.Left, util.Left, util.Up},
			{util.Up, util.Left, util.Up, util.Up, util.Left},
			{util.Up, util.Left, util.Up, util.Left, util.Up},
			{util.Up, util.Left, util.Left, util.Up, util.Up},
			{util.Left, util.Up, util.Up, util.Up, util.Left},
			{util.Left, util.Up, util.Up, util.Left, util.Up},
			{util.Left, util.Up, util.Left, util.Up, util.Up},
			{util.Left, util.Left, util.Up, util.Up, util.Up},
		}},
		{"7-A", args{'7', 'A'}, [][]util.Direction{
			{util.Down, util.Down, util.Down, util.Right, util.Right},
			{util.Down, util.Down, util.Right, util.Down, util.Right},
			{util.Down, util.Down, util.Right, util.Right, util.Down},
			{util.Down, util.Right, util.Down, util.Down, util.Right},
			{util.Down, util.Right, util.Down, util.Right, util.Down},
			{util.Down, util.Right, util.Right, util.Down, util.Down},
			{util.Right, util.Down, util.Down, util.Down, util.Right},
			{util.Right, util.Down, util.Down, util.Right, util.Down},
			{util.Right, util.Down, util.Right, util.Down, util.Down},
			{util.Right, util.Right, util.Down, util.Down, util.Down},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirectionsForDigit(tt.args.start, tt.args.target); !reflect.DeepEqual(dirSet(got), dirSet(tt.want)) {
				t.Errorf("DirectionsForDigit() = %v\nwant %v", got, tt.want)
			}
		})
	}
}

func dirSet(dirs [][]util.XY) map[string]struct{} {
	set := make(map[string]struct{}, len(dirs))
	for _, dir := range dirs {
		var dirString string
		for ix := range dir {
			dirString += string(dir[ix].ToRune())
		}
		set[dirString] = struct{}{}
	}
	return set
}
func sortDirections(dirs [][]util.XY) [][]util.XY {
	slices.SortFunc(dirs, func(a, b []util.XY) int {
		for ix := range a {
			if a[ix] == b[ix] {
				continue
			}
			if a[ix].X < b[ix].X || a[ix].Y < b[ix].Y {
				return -1
			} else {
				return 1
			}
		}
		return 0
	})
	return dirs
}

func TestDetermineInputs(t *testing.T) {
	type args struct {
		sequence string
	}
	tests := []struct {
		name string
		args args
		want [][]rune
	}{
		{"sample", args{"029A"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DetermineInputs(tt.args.sequence); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetermineInputs() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO: goldie
func TestBuildKeypadGrid(t *testing.T) {
	wantKeypad, _ := util.NewGridFromFile("testdata/keypad.txt")
	tests := []struct {
		name string
		want *util.Grid
	}{
		{"test", wantKeypad},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want.Print()
			got := BuildKeypadGrid()
			got.Print()
			if !reflect.DeepEqual(got.Checksum(), tt.want.Checksum()) {
				t.Errorf("BuildKeypadGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildAdjacency(t *testing.T) {
	keypad, _ := util.NewGridFromFile("testdata/keypad.txt")
	type args struct {
		grid *util.Grid
	}
	tests := []struct {
		name string
		args args
		want map[rune][]rune
	}{
		{"sample", args{keypad}, map[rune][]rune{
			'A': {'>', '^'},
			'>': {'v', 'A'},
			'v': {'>', '<', '^'},
			'<': {'v'},
			'^': {'A', 'v'},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildAdjacency(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildAdjacency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirectionsForKeypad(t *testing.T) {
	keypad, _ := util.NewGridFromFile("testdata/keypad.txt")
	type args struct {
		start     rune
		end       rune
		adjacency map[rune][]Position
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{"sample", args{'A', '<', BuildAdjacency(keypad)}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirectionsForKeypad(tt.args.start, tt.args.end, tt.args.adjacency); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DirectionsForKeypad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenCombo(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"layer2", args{[][]string{{"v<<A", "<v<A"}, {">>^A", ">^>A"}, {"<A"}, {">A", "^>vA"}}}, nil},
		{"layer3", args{[][]string{
			{"v<<A", "<v<A"},
			{">>^A", ">^>A"},
			{"<A"},
			{">A"},
			{"vA"},
			{"<^A", "^<A"},
			{"A"},
			{">A"},
			{"v<A", "<vA"},
			{"A"},
			{"A"},
			{">^A", "^>A"},
		}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenCombo(tt.args.paths, nil)
			for _, p := range asSortedStrings(got) {
				fmt.Println(p)
			}
			//fmt.Println(allPathsStr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestGenPath(t *testing.T) {
//	type args struct {
//		paths [][]string
//	}
//	tests := []struct {
//		name string
//		args args
//		want []string
//	}{
//		{"layer2", args{[][]string{{"<A"}, {"^A"}, {">^^A"}, {"vvvA"}}}, nil},
//		{"layer3", args{[][]string{{"v<<A"}, {">>^A"}, {"<A"}, {">A"}, {"vA"}, {"<^A"}, {"A"}, {">A"}, {"<vA"}, {"A"}, {"A"}, {">^A"}}}, nil},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got := GenPath(tt.args.paths, BuildAdjacency(BuildKeypadGrid()))
//			for _, p := range asSortedStrings(GenCombo(got, nil)) {
//				fmt.Println(p)
//			}
//			//fmt.Println(allPathsStr)
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Part1() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestPart1(t *testing.T) {
	type args struct {
		codes []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample", args{[]string{"029A", "980A", "179A", "456A", "379A"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.codes); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
