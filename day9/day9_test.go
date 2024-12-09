package day9

import (
	_ "embed"
	"reflect"
	"testing"
)

func TestExpandDiskMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []rune
		expected []int
	}{
		{name: "sample", input: []rune("12345"), expected: []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}},
		{name: "big-sample", input: []rune("2333133121414131402"), expected: []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9}},
		{name: "double-digit", input: []rune("233313312141413140233"), expected: []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9, -1, -1, -1, 10, 10, 10}},
		{name: "edge", input: []rune("1010101010101010101010"), expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ExpandDiskMap(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("%s:\n\t\texpected:\t\t%v\n\t\tgot\t\t\t\t%v", tt.name, tt.expected, actual)
			}
		})
	}
}

func TestCompact(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "sample", input: []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}, expected: []int{0, 2, 2, 1, 1, 1, 2, 2, 2, -1, -1, -1, -1, -1, -1}},
		{
			name:     "big-sample",
			input:    []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9},
			expected: []int{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		},
		{
			name:     "double-digit",
			input:    []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9, -1, -1, -1, 10, 10, 10},
			expected: []int{0, 0, 10, 10, 10, 1, 1, 1, 9, 9, 8, 2, 8, 8, 8, 3, 3, 3, 7, 4, 4, 7, 5, 5, 5, 5, 7, 6, 6, 6, 6, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Compact(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		name     string
		input    []rune
		expected int
	}{
		{name: "sample", input: []rune("12345"), expected: 60},
		{name: "big-sample", input: []rune("2333133121414131402"), expected: 1928},
		{name: "edge", input: []rune("1010101010101010101010"), expected: 385},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Part1(tt.input)
			if actual != tt.expected {
				t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}

func TestCompactBlocks(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "big-sample",
			input:    []int{0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9},
			expected: []int{0, 0, 9, 9, 2, 1, 1, 1, 7, 7, 7, -1, 4, 4, -1, 3, 3, 3, -1, -1, -1, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, -1, -1, -1, -1, 8, 8, 8, 8, -1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CompactBlocks(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}
