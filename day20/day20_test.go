package day20

import (
	"testing"

	"github.com/jmontroy90/aoc-2024/util"
)

func TestPart1(t *testing.T) {
	grid, _ := util.NewGridFromFile("testdata/small.txt")
	type args struct {
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{4}, 30},
		{"example", args{2}, 44},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid.Reset()
			if got := Part1(grid, tt.args.threshold); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	grid, _ := util.NewGridFromFile("testdata/small.txt")
	type args struct {
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"76", args{76}, 3},
		{"74", args{74}, 7},
		{"72", args{72}, 29},
		{"70", args{70}, 41},
		{"68", args{68}, 55},
		{"66", args{66}, 67},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid.Reset()
			if got := Part2(grid, tt.args.threshold); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
