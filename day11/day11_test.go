package day11

import (
	_ "embed"
	"reflect"
	"testing"
)

func TestCountStonesAfterBlinks(t *testing.T) {
	type input struct {
		stone   int
		nBlinks int
	}
	tests := []struct {
		name     string
		input    input
		expected int
	}{
		{"zero", input{0, 1}, 1},
		{"one", input{1, 1}, 1},
		{"10", input{10, 1}, 2},
		{"99", input{99, 1}, 2},
		{"999", input{999, 1}, 1},
		{"17-1", input{17, 1}, 2},
		{"17-2", input{17, 2}, 2},
		{"17-3", input{17, 3}, 3},
		{"17-4", input{17, 4}, 6},
		{"17-5", input{17, 5}, 8},
		{"17-6", input{17, 6}, 15},
		{"89741-25", input{89741, 25}, 14139},
		// 2024 - 20 24 - 2 0 2 4 - 4048 1 4048 8096
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CountStonesAfterBlinks(tt.input.stone, tt.input.nBlinks)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("%s:\n\t\texpected:\t\t%v\n\t\tgot\t\t\t\t%v", tt.name, tt.expected, actual)
			}
		})
	}
}

func TestCountStonesAfterBlinksCached(t *testing.T) {
	type input struct {
		stone   int
		nBlinks int
	}
	tests := []struct {
		name     string
		input    input
		expected int
	}{
		{"17-1", input{17, 1}, 2},
		{"17-2", input{17, 2}, 2},
		{"17-3", input{17, 3}, 3},
		{"17-4", input{17, 4}, 6},
		{"17-5", input{17, 5}, 8},
		{"17-6", input{17, 6}, 15},
		{"17-10", input{17, 10}, 71},
		{"89741-25", input{89741, 25}, 14139},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CountStonesAfterBlinksCached(tt.input.stone, tt.input.nBlinks) //, tt.input.stone, tt.input.nBlinks, 0)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("%s:\n\t\texpected:\t\t%v\n\t\tgot\t\t\t\t%v", tt.name, tt.expected, actual)
			}
		})
	}
}

func BenchmarkCountStonesAfterBlinks(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CountStonesAfterBlinks(89741, 25)
	}
}

func BenchmarkCountStonesAfterBlinksCached(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CountStonesAfterBlinksCached(89741, 25)
	}
}
