package day2

import "testing"

func TestIsReportSafe(t *testing.T) {
	tests := []struct {
		name     string
		report   []int
		expected bool
	}{
		{"safe", []int{7, 6, 4, 2, 1}, true},
		{"unsafe", []int{1, 2, 7, 8, 9}, false},
		{"unsafe", []int{9, 7, 6, 2, 1}, false},
		{"unsafe", []int{1, 3, 2, 4, 5}, false},
		{"unsafe", []int{8, 6, 4, 4, 1}, false},
		{"safe", []int{1, 3, 6, 7, 9}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, _ := IsReportSafe(tt.report)
			if actual != tt.expected {
				t.Errorf("%s: want %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}

func TestIsDampenedReportSafe(t *testing.T) {
	tests := []struct {
		name     string
		report   []int
		expected bool
	}{
		{"safe", []int{7, 6, 4, 2, 1}, true},
		{"unsafe", []int{1, 2, 7, 8, 9}, false},
		{"unsafe", []int{9, 7, 6, 2, 1}, false},
		{"safe-with-dampen", []int{1, 3, 2, 4, 5}, true},
		{"safe-with-dampen-2", []int{8, 6, 4, 4, 1}, true},
		{"safe", []int{1, 3, 6, 7, 9}, true},
		{"unsafe", []int{42, 39, 40, 41, 42, 43, 46, 43}, false},
		{"safe-with-removal", []int{80, 82, 81, 82, 83, 85, 88}, true},
		{"safe-with-removal", []int{35, 38, 39, 41, 44, 47, 50, 54}, true},
		{"safe-with-removal", []int{58, 56, 58, 59, 60, 62, 63, 65}, true},
		{"safe-with-removal", []int{63, 63, 64, 65, 67, 70, 72}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsDampenedReportSafe(tt.report)
			if actual != tt.expected {
				t.Errorf("%s: want %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}
