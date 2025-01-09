package day18

import (
	"bufio"
	"fmt"
	"github.com/jmontroy90/aoc-2024/util"
	"os"
	"strconv"
	"strings"
)

const (
	InputPath = "day18/input.txt"
)

func Runner() error {
	coords, err := load(InputPath)
	if err != nil {
		return err
	}
	grid := util.NewGridFromDim(71, 71)
	fmt.Printf("\n:: DAY 18 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(grid, coords))
	fmt.Printf("Part 2: %v\n", Part2(grid, coords))
	return nil
}

func Part1(grid *util.Grid, coords []util.XY) int {
	for _, coord := range coords[:1024] {
		grid.Set(coord, '#')
	}
	return BFS(grid, util.XY{0, 0}, util.XY{70, 70}, true)
}

// This is dumb brute-force, and it does work because our input is pretty small.
// A binary search would help for a bigger sample.
func Part2(grid *util.Grid, coords []util.XY) string {
	grid.Reset()
	// We know this is passable.
	for _, coord := range coords[:1024] {
		grid.Set(coord, '#')
	}
	for _, coord := range coords[1024:] {
		grid.Set(coord, '#')
		res := BFS(grid, util.XY{0, 0}, util.XY{70, 70}, false)
		if res == -1 {
			return fmt.Sprintf("%v,%v", coord.X, coord.Y)
		}
	}
	return "-1"
}

type POSLayer struct {
	pos   util.XY
	layer int
}

func BFS(grid *util.Grid, start, end util.XY, mark bool) int {
	queue := []POSLayer{{start, 0}}
	seen := make(map[util.XY]bool)
	seen[start] = true
	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		if mark { // just for them visuals
			marker := rune(p.layer + '0')
			if p.pos == start {
				marker = 'S'
			} else if p.pos == end {
				marker = 'E'
			}
			grid.Set(p.pos, marker)
		}
		if p.pos == end {
			return p.layer
		}
		for _, step := range util.NextDirections() {
			n := p.pos.Add(step)
			if grid.Get(n) != util.Empty {
				continue
			}
			if !seen[n] {
				seen[n] = true
				queue = append(queue, POSLayer{pos: n, layer: p.layer + 1})
			}
		}
	}
	return -1
}

func load(inputPath string) ([]util.XY, error) {
	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewScanner(f)
	var coords []util.XY
	for buf.Scan() {
		out := strings.Split(buf.Text(), ",")
		x, _ := strconv.Atoi(out[0])
		y, _ := strconv.Atoi(out[1])
		coords = append(coords, util.XY{x, y})
	}
	return coords, nil
}
