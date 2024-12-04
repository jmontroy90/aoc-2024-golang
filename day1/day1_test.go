package day1

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		name     string
		l, r     []int
		expected int
	}{
		{"negative", []int{2, 3, 4}, []int{3, 4, 5}, 3},
		{"positive", []int{3, 4, 5}, []int{1, 2, 3}, 6},
		{"unordered", []int{4, 1, 7}, []int{1, 9, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.l, tt.r)
			if got != tt.expected {
				t.Errorf("%s: want %d, got %d", tt.name, tt.expected, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name     string
		l, r     []int
		expected int
	}{
		{"each-appears-once", []int{1, 2, 3}, []int{1, 2, 3}, 6},
		{"each-appears-twice", []int{1, 2, 3}, []int{1, 1, 2, 2, 3, 3}, 12},
		{"some-missing", []int{1, 2, 3}, []int{1, 1, 3, 3}, 8},
		{"unordered", []int{3, 1, 2}, []int{1, 3, 1, 3}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.l, tt.r)
			if got != tt.expected {
				t.Errorf("%s: want %d, got %d", tt.name, tt.expected, got)
			}
		})
	}
}
