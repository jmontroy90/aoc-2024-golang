package day6

import (
	"bufio"
	"fmt"
	"os"
)

const (
	InputPath = "day6/input.txt"
)

func Runner() error {
	grid, err := loadGrid(InputPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 6 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(NewGrid(grid)))
	fmt.Printf("Part 2: %v\n", Part2(NewGrid(grid)))
	return nil
}

const (
	OOB       = '\x10'
	Traversed = 'X'
	Block     = '#'
	Start     = '^'
)

type Grid struct {
	raw            [][]rune // this copy for resets is a little
	grid           [][]rune
	traveledSpots  int
	seenGuardCount map[Guard]int
}

func NewGrid(raw [][]rune) *Grid {
	return &Grid{
		raw:            raw, // for resets
		grid:           newGridFromRaw(raw),
		seenGuardCount: make(map[Guard]int),
	}
}

func newGridFromRaw(raw [][]rune) [][]rune {
	grid := make([][]rune, len(raw))
	for ix, g := range raw {
		newG := make([]rune, len(g))
		copy(newG, g)
		grid[ix] = newG
	}
	return grid
}

func (g *Grid) Get(pos XY) rune {
	if pos.X >= len(g.grid[0]) || pos.Y >= len(g.grid) || pos.X < 0 || pos.Y < 0 {
		return OOB
	}
	return g.grid[pos.Y][pos.X]
}

func (g *Grid) Scan(r rune) (XY, bool) {
	for y, line := range g.grid {
		for x := range line {
			p := XY{x, y}
			if g.Get(p) == r {
				return p, true
			}
		}
	}
	return XY{}, false
}

func (g *Grid) MarkTravel(guard Guard) {
	if r := g.Get(guard.Position); r != Traversed {
		g.Set(guard.Position, Traversed)
		g.traveledSpots++
		g.seenGuardCount[guard] = 1
	} else {
		g.seenGuardCount[guard] += 1
	}
}

func (g *Grid) Set(pos XY, r rune) {
	g.grid[pos.Y][pos.X] = r
}

func (g *Grid) Print() {
	for y := range g.grid {
		for x := range g.grid[y] {
			fmt.Printf("%c", g.grid[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *Grid) BeenHereBefore(guard Guard) bool {
	if c, ok := g.seenGuardCount[guard]; ok {
		return c > 1
	}
	return false
}

func (g *Grid) Clear() {
	g.seenGuardCount = make(map[Guard]int)
	g.traveledSpots = 0
	g.grid = newGridFromRaw(g.raw)
}

type XY struct {
	X int
	Y int
}

func (p XY) Print() {
	fmt.Printf("(%d,%d)\n", p.X+1, p.Y+1)
}

func (p XY) Add(that XY) XY {
	return XY{X: p.X + that.X, Y: p.Y + that.Y}
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (d Direction) ToStep() XY {
	var pos XY
	switch d {
	case Up:
		pos = XY{0, -1}
	case Right:
		pos = XY{1, 0}
	case Down:
		pos = XY{0, 1}
	case Left:
		pos = XY{-1, 0}
	}
	return pos
}

func (d Direction) Next() Direction {
	return (d + 1) % 4
}

type Guard struct {
	Direction Direction
	Position  XY
}

func (g *Guard) NextStep() XY {
	return g.Position.Add(g.Direction.ToStep())
}

func (g *Guard) TakeStep() {
	g.Position = g.NextStep()
}

func Part1(grid *Grid) int {
	startPos, ok := grid.Scan(Start)
	if !ok {
		return -1
	}
	g := Guard{Direction: Up, Position: startPos}
	grid.MarkTravel(g)
	for {
		for grid.Get(g.NextStep()) != Block {
			if grid.Get(g.NextStep()) == OOB {
				return grid.traveledSpots
			}
			g.TakeStep()
			grid.MarkTravel(g)
		}
		g.Direction = g.Direction.Next()
	}
}

func Part2(grid *Grid) int {
	startPos, ok := grid.Scan(Start)
	if !ok {
		return -1
	}
	g := Guard{Direction: Up, Position: startPos}
	var loopCount int
	for y, row := range grid.grid {
		for x := range row {
			grid.Set(XY{x, y}, Block)
			if hasLoop(grid, g) {
				loopCount++
			}
			grid.Set(XY{x, y}, '.')
			grid.Clear()
		}
	}
	return loopCount
}

// TODO: This logic is probably reconcilable and collapsible with Part1.
func hasLoop(grid *Grid, g Guard) bool {
	for {
		for grid.Get(g.NextStep()) != Block {
			if grid.Get(g.NextStep()) == OOB {
				return false
			}
			if grid.BeenHereBefore(g) {
				return true
			}
			g.TakeStep()
			grid.MarkTravel(g)
		}
		g.Direction = g.Direction.Next()
	}
}

func loadGrid(filePath string) ([][]rune, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	var grid [][]rune
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return grid, nil
}
