package day12

import (
	"fmt"
	"github.com/jmontroy90/aoc-2024/util"
	"strings"
)

const (
	InputPath = "day12/input.txt"
)

func Runner() error {
	grid, err := util.NewGridFromFile(InputPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 12 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(grid))
	grid.Clear()
	fmt.Printf("Part 2: %v\n", Part2(grid))
	return nil
}

func Part1(grid *util.Grid) int {
	rs := NewRegionScanner()
	var total int
	for grid.ScanForNot(Scanned) {
		rr := grid.Get(grid.CurrentPosition())
		rs.ScanRegion(grid, grid.CurrentPosition(), rr)
		regionTotal := len(rs.posVisited) * len(rs.fencePlacements)
		total += regionTotal
		rs.ResetStats()
	}
	return total
}

func Part2(grid *util.Grid) int {
	rs := NewRegionScanner()
	var total int
	for grid.ScanForNot(Scanned) {
		rr := grid.Get(grid.CurrentPosition())
		rs.ScanRegion(grid, grid.CurrentPosition(), rr)
		sides := rs.CalculateSides()
		regionTotal := len(rs.posVisited) * sides
		total += regionTotal
		rs.ResetStats()
	}
	return total
}

const (
	Scanned = '-'
)

type Fence struct {
	POS       util.XY
	Direction util.Direction
}

func (f Fence) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("(%v,%v)@%v", f.POS.X, f.POS.Y, strings.ToUpper(f.Direction.Name()))), nil
}

// TODO: This is not the best naming structure / division of labor
type RegionScanner struct {
	posVisited      map[util.XY]struct{}
	fencePlacements map[Fence]struct{}
}

func NewRegionScanner() *RegionScanner {
	return &RegionScanner{
		posVisited:      make(map[util.XY]struct{}),
		fencePlacements: make(map[Fence]struct{}),
	}
}

func (rs *RegionScanner) ResetStats() {
	rs.posVisited = make(map[util.XY]struct{})
	rs.fencePlacements = make(map[Fence]struct{})
}

func (rs *RegionScanner) ScanRegion(grid *util.Grid, pos util.XY, regionRune rune) {
	rs.posVisited[pos] = struct{}{} // i am in here
	grid.Set(pos, Scanned)          // we still need to do this to help our next iteration find where to start from
	for _, dir := range pos.NextDirections() {
		if nextPos := pos.Add(dir); grid.Get(nextPos) == regionRune { // so not something else OR OOB
			rs.ScanRegion(grid, nextPos, regionRune) // take steps until we fail
		} else if _, ok := rs.posVisited[nextPos]; !ok {
			// we haven't been here, AND it's not our region - must be something to fence!
			rs.fencePlacements[Fence{POS: pos, Direction: dir}] = struct{}{} // put up fences
		}
	}
}

func (rs *RegionScanner) CalculateSides() int {
	var total int
	for fence := range rs.fencePlacements {
		total += rs.findSide(fence)
	}
	return total
}

func (rs *RegionScanner) findSide(fence Fence) int {
	switch d := fence.Direction; d {
	case util.Up, util.Down:
		rs.deleteSide(fence, []util.XY{util.Right, util.Left}, d)
	case util.Left, util.Right:
		rs.deleteSide(fence, []util.XY{util.Up, util.Down}, d)
	}
	return 1
}

func (rs *RegionScanner) deleteSide(fence Fence, scanDirs []util.XY, fenceDir util.XY) {
	delete(rs.fencePlacements, fence)
	for _, sd := range scanDirs {
		nextFence := Fence{POS: fence.POS.Add(sd), Direction: fenceDir}
		if _, ok := rs.fencePlacements[nextFence]; ok {
			rs.deleteSide(nextFence, scanDirs, fenceDir)
		}
	}
}
