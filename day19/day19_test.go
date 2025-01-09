package day19

import (
	_ "embed"
	"testing"
)

func TestIsDesignPossible(t *testing.T) {
	towels := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	type args struct {
		towels []string
		design string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"brwrr", args{towels, "brwrr"}, true},
		{"bggr", args{towels, "bggr"}, true},
		{"gbbr", args{towels, "gbbr"}, true},
		{"rrbgbr", args{towels, "rrbgbr"}, true},
		{"ubwu", args{towels, "ubwu"}, false},
		{"bwurrg", args{towels, "bwurrg"}, true},
		{"brgr", args{towels, "brgr"}, true},
		{"bbrgwb", args{towels, "bbrgwb"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDesignPossible(tt.args.towels, tt.args.design, nil); got != tt.want {
				t.Errorf("IsDesignPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumPossibleDesigns(t *testing.T) {
	towels := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	type args struct {
		towels []string
		design string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"brwrr", args{towels, "brwrr"}, 2},
		{"bggr", args{towels, "bggr"}, 1},
		{"gbbr", args{towels, "gbbr"}, 4},
		{"rrbgbr", args{towels, "rrbgbr"}, 6},
		{"ubwu", args{towels, "ubwu"}, 0},
		{"bwurrg", args{towels, "bwurrg"}, 1},
		{"brgr", args{towels, "brgr"}, 2},
		{"bbrgwb", args{towels, "bbrgwb"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumPossibleDesigns(tt.args.towels, tt.args.design, make(map[string]int)); got != tt.want {
				t.Errorf("NumPossibleDesigns() = %v, want %v", got, tt.want)
			}
		})
	}
}
