package day4

import "testing"

func TestFindXmas(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]rune
		expected int
	}{
		{
			name: "given-example",
			grid: [][]rune{
				[]rune("MMMSXXMASM"),
				[]rune("MSAMXMSMSA"),
				[]rune("AMXSXMAAMM"),
				[]rune("MSAMASMSMX"),
				[]rune("XMASAMXAMM"),
				[]rune("XXAMMXXAMA"),
				[]rune("SMSMSASXSS"),
				[]rune("SAXAMASAAA"),
				[]rune("MAMMMXMMMM"),
				[]rune("MXMXAXMASX"),
			},
			expected: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FindXmasUni(tt.grid)
			if actual != tt.expected {
				t.Errorf("%s: got %v, expected %v", tt.name, actual, tt.expected)
			}
		})
	}
}
