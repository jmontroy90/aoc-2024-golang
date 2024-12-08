package day7

import (
	_ "embed"
	"testing"
)

func TestIsTrueEquation(t *testing.T) {
	tests := []struct {
		name     string
		eq       Equation
		expected bool
	}{
		{"one", Equation{[]int{10, 19}, 190}, true},
		{"two", Equation{[]int{81, 40, 27}, 3267}, true},
		{"three", Equation{[]int{17, 5}, 83}, false},
		{"four", Equation{[]int{15, 6}, 156}, false},
		{"five", Equation{[]int{6, 8, 6, 15}, 7290}, false},
		{"six", Equation{[]int{16, 10, 13}, 161011}, false},
		{"seven", Equation{[]int{17, 8, 14}, 192}, false},
		{"eight", Equation{[]int{9, 7, 18, 13}, 21037}, false},
		{"nine", Equation{[]int{11, 6, 16, 20}, 292}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsTrueEquation(tt.eq, []Op{Add, Mult})
			if actual != tt.expected {
				t.Errorf("%s: expected %t, got %t", tt.name, tt.expected, actual)
			}
		})
	}
}

func TestIsTrueEquationWithConcat(t *testing.T) {
	tests := []struct {
		name     string
		eq       Equation
		expected bool
	}{
		{"one", Equation{[]int{10, 19}, 190}, true},
		{"two", Equation{[]int{81, 40, 27}, 3267}, true},
		{"three", Equation{[]int{17, 5}, 83}, false},
		{"four", Equation{[]int{15, 6}, 156}, true},
		{"five", Equation{[]int{6, 8, 6, 15}, 7290}, true},
		{"six", Equation{[]int{16, 10, 13}, 161011}, false},
		{"seven", Equation{[]int{17, 8, 14}, 192}, true},
		{"eight", Equation{[]int{9, 7, 18, 13}, 21037}, false},
		{"nine", Equation{[]int{11, 6, 16, 20}, 292}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsTrueEquation(tt.eq, []Op{Add, Mult, Concat})
			if actual != tt.expected {
				t.Errorf("%s: expected %t, got %t", tt.name, tt.expected, actual)
			}
		})
	}
}
