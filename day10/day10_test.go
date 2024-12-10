package day10

import (
	_ "embed"
	"github.com/jmontroy90/aoc-2024/util"
	"reflect"
	"testing"
)

func TestPart1(t *testing.T) {
	grid2, _ := util.NewGridFromFile("testdata/twoScore.txt")
	grid4, _ := util.NewGridFromFile("testdata/fourScore.txt")
	grid3, _ := util.NewGridFromFile("testdata/threeScore.txt")
	grid36, _ := util.NewGridFromFile("testdata/36Score.txt")

	tests := []struct {
		name     string
		input    *util.Grid
		expected int
	}{
		{name: "2-score", input: grid2, expected: 2},
		{name: "4-score", input: grid4, expected: 4},
		{name: "3-score", input: grid3, expected: 3},
		{name: "36-score", input: grid36, expected: 36},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Part1(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("%s:\n\t\texpected:\t\t%v\n\t\tgot\t\t\t\t%v", tt.name, tt.expected, actual)
			}
		})
	}
}
