package day12

import (
	_ "embed"
	"encoding/json"
	"github.com/jmontroy90/aoc-2024/util"
	"github.com/sebdah/goldie/v2"
	"testing"
)

type testScanRegion struct {
	name           string
	gridFile       string
	regionRune     rune
	startPos       util.XY
	expectedPoses  int
	expectedFences int
}

func TestScanRegion(t *testing.T) {
	tests := []testScanRegion{
		{"simple-A", "simple.txt", 'A', util.XY{0, 0}, 4, 10},
		{"simple-B", "simple.txt", 'B', util.XY{0, 1}, 4, 8},
		{"simple-C", "simple.txt", 'C', util.XY{2, 1}, 4, 10},
	}
	g := goldie.New(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := NewRegionScanner()
			grid, _ := util.NewGridFromFile("testdata/input/" + tt.gridFile)
			rs.ScanRegion(grid, tt.startPos, tt.regionRune)
			assertStats(t, g, tt, rs)
			// kinda a backup check for goldie goofs
			if len(rs.posVisited) != tt.expectedPoses {
				t.Errorf("len(posVisited): expected %v, got %v", tt.expectedPoses, len(rs.posVisited))
			}
			if len(rs.fencePlacements) != tt.expectedFences {
				t.Errorf("len(fencePlacements): expected %v, got %v", tt.expectedFences, len(rs.fencePlacements))
			}
		})
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		name     string
		gridFile string
		expected int
	}{
		{"simple", "simple.txt", 140},
		{"contained", "contained.txt", 772},
		{"complex", "complex.txt", 1930},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid, _ := util.NewGridFromFile("testdata/input/" + tt.gridFile)
			actual := Part1(grid)
			if actual != tt.expected {
				t.Errorf("TestPart1: expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

func TestCalculateSides(t *testing.T) {
	tests := []struct {
		name       string
		gridFile   string
		startPos   util.XY
		regionRune rune
		expected   int
	}{
		{"simple-A", "simple.txt", util.XY{0, 0}, 'A', 4},
		{"simple-B", "simple.txt", util.XY{0, 1}, 'B', 4},
		{"simple-C", "simple.txt", util.XY{2, 1}, 'C', 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := NewRegionScanner()
			grid, _ := util.NewGridFromFile("testdata/input/" + tt.gridFile)
			rs.ScanRegion(grid, tt.startPos, tt.regionRune)
			actual := rs.CalculateSides()
			if actual != tt.expected {
				t.Errorf("TestCalculateSides: expected %v, got %v", tt.expected, actual)
			}
		})
	}

}

func assertStats(t *testing.T, g *goldie.Goldie, tt testScanRegion, rs *RegionScanner) {
	bsPosVisited, err := json.MarshalIndent(rs.posVisited, "", " ")
	if err != nil {
		t.Errorf("err marshaling posVisited: %v", err)
	}
	g.Assert(t, tt.name+"-posVisited", bsPosVisited)
	bsFencePlacements, err := json.MarshalIndent(rs.fencePlacements, "", " ")
	if err != nil {
		t.Errorf("err marshaling fencePlacements: %v", err)
	}
	g.Assert(t, tt.name+"-fencePlacements", bsFencePlacements)
}
