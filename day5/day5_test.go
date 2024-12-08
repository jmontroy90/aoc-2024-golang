package day5

import (
	"reflect"
	"testing"
)

var (
	testOrderings = [][2]int{
		{47, 53},
		{97, 13},
		{97, 61},
		{97, 47},
		{75, 29},
		{61, 13},
		{75, 53},
		{29, 13},
		{97, 29},
		{53, 29},
		{61, 53},
		{97, 53},
		{61, 29},
		{47, 13},
		{75, 47},
		{97, 75},
		{47, 61},
		{75, 61},
		{47, 29},
		{75, 13},
		{53, 13},
	}
	testRL = createOrderMap(testOrderings)
)

func TestIsOrderingValid(t *testing.T) {
	tests := []struct {
		name     string
		update   []int
		expected bool
	}{
		{"one", []int{75, 47, 61, 53, 29}, true},
		{"two", []int{97, 61, 53, 29, 13}, true},
		{"three", []int{75, 29, 13}, true},
		{"four", []int{75, 97, 47, 61, 53}, false},
		{"five", []int{61, 13, 29}, false},
		{"six", []int{97, 13, 75, 29, 47}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := IsUpdateValid(tt.update, testRL)
			if actual != tt.expected {
				t.Errorf("%s: got %v, expected %v", tt.name, actual, tt.expected)
			}
		})
	}
}

func TestReorderUpdate(t *testing.T) {
	tests := []struct {
		name     string
		update   []int
		expected []int
	}{
		{"one", []int{75, 97, 47, 61, 53}, []int{97, 75, 47, 61, 53}},
		{"two", []int{61, 13, 29}, []int{61, 29, 13}},
		{"three", []int{97, 13, 75, 29, 47}, []int{97, 75, 47, 29, 13}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ReorderUpdate(tt.update, testRL)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("%s: got %v, expected %v", tt.name, actual, tt.expected)
			}
		})
	}
}
