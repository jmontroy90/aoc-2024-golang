package day16

import (
	"fmt"
	"github.com/jmontroy90/aoc-2024/util"
	"slices"
)

const (
	InputPath = "day16/input.txt"
)

func Runner() error {
	grid, err := util.NewGridFromFile(InputPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 16 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(grid))
	grid.Reset()
	fmt.Printf("Part 2: %v\n", Part2(grid))
	return nil
}

func Part1(grid *util.Grid) int {
	start, _ := grid.ScanOnce('S')
	return FindPathBFS(grid, start, util.Right)
}

func Part2(grid *util.Grid) int {
	start, _ := grid.ScanOnce('S')
	end, _ := grid.ScanOnce('E')
	return FindPathWithMemory(grid, start, util.Right, end)
}

type Step struct {
	pos util.XY
	dir util.Direction
}

type StepScore struct {
	pos   util.XY
	dir   util.Direction
	score int
}

func FindPathBFS(grid *util.Grid, pos, dir util.XY) int {
	seen := make(map[Step]struct{})
	queue := []StepScore{{pos, dir, 0}}
	var minScore int
	for len(queue) != 0 {
		ss := queue[0]
		queue = queue[1:]
		step := Step{ss.pos, ss.dir}
		if _, ok := seen[step]; ok {
			continue
		}
		//e, _ := grid.GetSet(ss.pos, 'X')
		//grid.Print()
		//grid.Set(ss.pos, e)
		seen[step] = struct{}{}
		if grid.Get(step.pos) == 'E' {
			minScore = ss.score
			break
		}
		nexts := []StepScore{
			{step.pos.Add(step.dir), step.dir, ss.score + 1},
			{step.pos.Add(step.dir.TurnClockwise()), step.dir.TurnClockwise(), ss.score + 1001},
			{step.pos.Add(step.dir.TurnCounterClockwise()), step.dir.TurnCounterClockwise(), ss.score + 1001},
		}
		for _, next := range nexts {
			if grid.Get(next.pos) != '#' {
				queue = append(queue, next)
			}
		}
		slices.SortStableFunc(queue, func(a, b StepScore) int {
			if a.score < b.score {
				return -1
			} else if a.score == b.score {
				return 0
			}
			return 1
		})
	}
	return minScore
}

type StepScorePath struct {
	pos   util.XY
	dir   util.Direction
	score int
	path  int
}

type StepPath struct {
	pos  util.XY
	dir  util.Direction
	path int
}

func FindPathWithMemory(grid *util.Grid, start, dir util.XY, end util.XY) int {
	seen := make(map[Step][]int)
	seenScore := make(map[StepPath]int)
	queue := []StepScorePath{{start, dir, 0, 0}}
	pathTracker := make(map[int][]util.XY)
	pathTracker[0] = []util.XY{start}
	pathScores := make(map[int]int)
	minScore := 1000000000 // some very big number we can be less than
	var latestPath int
	for len(queue) != 0 {
		ss := queue[0]
		queue = queue[1:]
		s := Step{ss.pos, ss.dir}
		if paths, ok := seen[s]; ok {
			var goodNewPath bool
			for _, path := range paths {
				score := seenScore[StepPath{s.pos, s.dir, path}]
				if ss.score <= score {
					goodNewPath = true
					break
				}
			}
			if !goodNewPath {
				continue
			}
		}
		seen[s] = append(seen[s], ss.path)
		seenScore[StepPath{ss.pos, ss.dir, ss.path}] = ss.score
		pathTracker[ss.path] = append(pathTracker[ss.path], ss.pos)
		if ss.pos == end {
			if ss.score <= minScore {
				minScore = ss.score
			}
			pathScores[ss.path] = ss.score
			continue
		}
		var addPaths bool
		nexts := []StepScorePath{
			{ss.pos.Add(ss.dir), ss.dir, ss.score + 1, ss.path},
			{ss.pos.Add(ss.dir.TurnClockwise()), ss.dir.TurnClockwise(), ss.score + 1001, ss.path},
			{ss.pos.Add(ss.dir.TurnCounterClockwise()), ss.dir.TurnCounterClockwise(), ss.score + 1001, ss.path},
		}
		for _, next := range nexts {
			if grid.Get(next.pos) != '#' {
				if addPaths {
					latestPath++
					newPath := make([]util.XY, len(pathTracker[ss.path]))
					copy(newPath, pathTracker[ss.path])
					pathTracker[latestPath] = newPath
					next.path = latestPath
				}
				queue = append(queue, next)
				addPaths = true
			}
		}
		slices.SortStableFunc(queue, func(a, b StepScorePath) int {
			if a.score < b.score {
				return -1
			} else if a.score == b.score {
				return 0
			}
			return 1
		})
	}
	benchLocations := make(map[util.XY]struct{})
	for pathID, score := range pathScores {
		if score == minScore {
			path := pathTracker[pathID]
			for _, p := range path {
				benchLocations[p] = struct{}{}
			}
		}
	}
	return len(benchLocations)
}

func printPath(grid *util.Grid, pathTracker map[int][]util.XY, path int) {
	for _, pos := range pathTracker[path] {
		grid.Set(pos, rune(path)+'A')
	}
	grid.Print()
	for _, pos := range pathTracker[path] {
		grid.Set(pos, util.Empty)
	}
}
