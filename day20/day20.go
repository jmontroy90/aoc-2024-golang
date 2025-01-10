package day20

import (
	"fmt"
	"github.com/jmontroy90/aoc-2024/util"
)

const (
	InputPath = "day20/input.txt"
)

func Runner() error {
	grid, err := util.NewGridFromFile(InputPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 20 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(grid, 100))
	fmt.Printf("Part 2: %v\n", Part2(grid, 100))
	return nil
}

func Part1(grid *util.Grid, threshold int) int {
	poses, _ := Traverse(grid)
	var totalSavers int
	for pos, sc := range poses {
		for _, dir := range util.NextDirections() {
			next1, next2 := pos.Add(dir), pos.Add(dir).Add(dir)
			if grid.Get(next1) == '#' && (grid.Get(next2) == util.Empty || grid.Get(next2) == 'E') {
				newSC, ok := poses[next2]
				if !ok {
					panic("WHAT YOU DO")
				}
				saved := newSC - sc - 2 // our cheat takes two steps
				if saved >= threshold {
					totalSavers++
				}
			}
		}
	}
	return totalSavers
}

func Part2(grid *util.Grid, threshold int) int {
	poses, _ := Traverse(grid)
	var totalSavers int
	for pos, sc := range poses {
		totalSavers += TryCheat(grid, poses, threshold, pos, sc)
	}
	return totalSavers
}

type PosStep struct {
	pos  util.XY
	step int
}

type StartEnd struct {
	start, end util.XY
}

// This is just a slightly-tweaked BFS, I think.
func TryCheat(grid *util.Grid, poses map[util.XY]int, threshold int, initPos util.XY, initScore int) int {
	queue := []PosStep{{initPos, 0}}
	processed := make(map[util.XY]bool)
	processed[initPos] = true
	processedCheat := make(map[StartEnd]bool)
	var savers int
	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		if p.step == 20 {
			break
		}
		for _, dir := range util.NextDirections() {
			next := p.pos.Add(dir)
			if processed[next] {
				continue
			}
			if grid.Get(next) == '#' {
				processed[next] = true
				queue = append(queue, PosStep{next, p.step + 1})
			}
			if grid.Get(next) == util.Empty || grid.Get(next) == 'E' || grid.Get(next) == 'S' {
				processed[next] = true
				newSC, ok := poses[next]
				if !ok {
					panic("ohno")
				}
				saved := newSC - initScore - p.step
				if saved >= threshold && !processedCheat[StartEnd{initPos, next}] {
					processedCheat[StartEnd{initPos, next}] = true
					savers++
				}
				queue = append(queue, PosStep{next, p.step + 1})
			}
		}
	}
	return savers
}

func Traverse(grid *util.Grid) (map[util.XY]int, int) {
	start, _ := grid.ScanOnce('S')
	end, _ := grid.ScanOnce('E')
	currPos := start
	poses := make(map[util.XY]int)
	stepCount := 1
	for {
		poses[currPos] = stepCount
		for _, step := range util.NextDirections() {
			next := currPos.Add(step)
			if _, ok := poses[next]; ok {
				continue
			}
			if currPos == end {
				return poses, stepCount
			}
			if grid.Get(next) != '#' {
				currPos = next
				stepCount++
				break
			}
		}
	}
}
