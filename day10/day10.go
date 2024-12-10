package day10

import (
	"fmt"
	"github.com/jmontroy90/aoc-2024/util"
)

const (
	InputPath = "day10/input.txt"
)

func Runner() error {
	grid, err := util.NewGridFromFile(InputPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 10 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(grid))
	fmt.Printf("Part 2: %v\n", Part2(grid))
	return nil
}

func Part1(grid *util.Grid) int {
	var trailheads []util.XY
	for grid.Scan('0') {
		trailheads = append(trailheads, grid.CurrentPosition())
	}
	tc := NewTrailCounter(CountPeaks)
	return tc.CountHikingTrails(grid, trailheads)
}

func Part2(grid *util.Grid) int {
	var trailheads []util.XY
	for grid.Scan('0') {
		trailheads = append(trailheads, grid.CurrentPosition())
	}
	tc := NewTrailCounter(CountTrails)
	return tc.CountHikingTrails(grid, trailheads)
}

const (
	Trailhead = 0
	Peak      = 9
)

type CountMethod int

const (
	CountPeaks CountMethod = iota
	CountTrails
)

type TrailCounter struct {
	visited     map[util.XY]struct{}
	countMethod CountMethod
}

func NewTrailCounter(method CountMethod) TrailCounter {
	tc := TrailCounter{countMethod: method}
	tc.ResetVisited()
	return tc
}

func (tc *TrailCounter) ResetVisited() {
	tc.visited = make(map[util.XY]struct{})
}

func (tc *TrailCounter) CountHikingTrails(grid *util.Grid, trailheads []util.XY) int {
	var totalVisited int
	for _, th := range trailheads {
		if tc.countMethod == CountPeaks {
			tc.ResetVisited()
		}
		for _, step := range th.NextSteps() {
			totalVisited += tc.countRecurseFn(grid, step, Trailhead)
		}
	}
	return totalVisited
}

func (tc *TrailCounter) countRecurseFn(grid *util.Grid, newPos util.XY, startHeight int) int {
	nextHeight, _ := grid.GetInt(newPos)
	if nextHeight == Peak && startHeight == Peak-1 {
		return tc.count(newPos)
	}
	if nextHeight == startHeight+1 {
		return tc.countRecurseFn(grid, newPos.Add(util.XY{1, 0}), nextHeight) +
			tc.countRecurseFn(grid, newPos.Add(util.XY{0, 1}), nextHeight) +
			tc.countRecurseFn(grid, newPos.Add(util.XY{-1, 0}), nextHeight) +
			tc.countRecurseFn(grid, newPos.Add(util.XY{0, -1}), nextHeight)
	}
	return 0
}

func (tc *TrailCounter) count(newPos util.XY) int {
	if tc.countMethod == CountTrails {
		// CountTrails, e.g. Part 2 logic
		return 1
	}
	// CountPeaks, e.g. Part 1 logic
	if _, ok := tc.visited[newPos]; !ok {
		tc.visited[newPos] = struct{}{}
		return 1
	}
	return 0
}
