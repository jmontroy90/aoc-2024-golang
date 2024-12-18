package day13

import (
	"github.com/jmontroy90/aoc-2024/util"
	"testing"
)

func TestFindCheapestButtonPresses(t *testing.T) {
	tests := []struct {
		name     string
		machine  Machine
		expected int
	}{
		{"sample", Machine{ButtonA: util.XY{94, 34}, ButtonB: util.XY{22, 67}, PrizeAt: util.XY{8400, 5400}}, 280},
		{"not-possible", Machine{ButtonA: util.XY{26, 66}, ButtonB: util.XY{67, 21}, PrizeAt: util.XY{12748, 12176}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.machine.FindCheapestButtonPresses()
			if actual != tt.expected {
				t.Errorf("%v: expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}

func TestFindCheapestWithMath(t *testing.T) {
	tests := []struct {
		name     string
		machine  Machine
		expected int
	}{
		{"sample", Machine{ButtonA: util.XY{94, 34}, ButtonB: util.XY{22, 67}, PrizeAt: util.XY{8400, 5400}}, 280},
		{"sample-2", Machine{ButtonA: util.XY{26, 66}, ButtonB: util.XY{67, 21}, PrizeAt: util.XY{12748, 12176}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.machine.FindCheapestWithMath()
			if actual != tt.expected {
				t.Errorf("%v: expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}

func TestFindCheapestWitLinearAlgebra(t *testing.T) {
	tests := []struct {
		name     string
		machine  Machine
		expected int
	}{
		{"sample", Machine{ButtonA: util.XY{94, 34}, ButtonB: util.XY{22, 67}, PrizeAt: util.XY{8400, 5400}}, 280},
		{"sample-2", Machine{ButtonA: util.XY{26, 66}, ButtonB: util.XY{67, 21}, PrizeAt: util.XY{12748, 12176}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.machine.FindCheapestWithLinearAlgebra()
			if actual != tt.expected {
				t.Errorf("%v: expected %v, got %v", tt.name, tt.expected, actual)
			}
		})
	}
}
