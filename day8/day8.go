package day8

import (
	"bufio"
	"fmt"
	"os"
)

const (
	InputPath = "day8/input.txt"
)

func Runner() error {
	gridRaw, err := load(InputPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 8 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(NewGrid(gridRaw)))
	fmt.Printf("Part 2: %v\n", Part2(NewGrid(gridRaw)))
	return nil
}

const (
	OOB      = '\x10'
	Empty    = '.'
	Antinode = '#'
)

type Grid struct {
	raw  [][]rune // this copy for resets is a little weird
	grid [][]rune
}

func NewGrid(raw [][]rune) *Grid {
	return &Grid{
		raw:  raw, // for resets
		grid: newGridFromRaw(raw),
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
	if g.IsOOB(pos) {
		return OOB
	}
	return g.grid[pos.Y][pos.X]
}

func (g *Grid) IsOOB(pos XY) bool {
	if pos.X >= len(g.grid[0]) || pos.Y >= len(g.grid) || pos.X < 0 || pos.Y < 0 {
		return true
	}
	return false
}

func (g *Grid) Scan(r rune) (XY, bool) {
	for y, line := range g.grid {
		for x := range line {
			p := XY{X: x, Y: y}
			if g.Get(p) == r {
				return p, true
			}
		}
	}
	return XY{}, false
}

func (g *Grid) Set(pos XY, r rune) bool {
	if g.IsOOB(pos) {
		return false
	}
	g.grid[pos.Y][pos.X] = r
	return true
}

// GetSet returns the existing rune at `pos` if found, plus true if `r` is set successfully, false if not (OOB).
func (g *Grid) GetSet(pos XY, r rune) (rune, bool) {
	if g.IsOOB(pos) {
		return OOB, false // choosing a silent semantic for now
	}
	if existing := g.Get(pos); existing != Empty {
		g.grid[pos.Y][pos.X] = r
		return existing, true
	}
	g.grid[pos.Y][pos.X] = r
	return Empty, true
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

func (g *Grid) Clear() {
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

func (p XY) Subtract(that XY) XY {
	return XY{X: p.X - that.X, Y: p.Y - that.Y}
}

func (p XY) Invert() XY {
	return XY{X: -p.X, Y: -p.Y}
}

func (p XY) Multiply(n int) XY {
	return XY{X: p.X * n, Y: p.Y * n}
}

func Part1(grid *Grid) int {
	var total int
	am := buildAntennaMap(grid)
	for _, locs := range am {
		for ix, loc := range locs {
			for iix := ix + 1; iix < len(locs); iix++ {
				diff := loc.Subtract(locs[iix])
				total += SetAndAddUniqueAntinode(grid, loc.Add(diff))
				total += SetAndAddUniqueAntinode(grid, locs[iix].Subtract(diff))
			}
		}
	}
	return total
}

func Part2(grid *Grid) int {
	var total int
	am := buildAntennaMap(grid)
	for _, locs := range am {
		for ix, loc := range locs {
			for iix := ix + 1; iix < len(locs); iix++ {
				total += SetAndAddUniqueAntinode(grid, loc)
				total += SetAndAddUniqueAntinode(grid, locs[iix])
				diff := loc.Subtract(locs[iix])
				n := 1
				newLoc := loc.Add(diff.Multiply(n))
				for !grid.IsOOB(newLoc) {
					total += SetAndAddUniqueAntinode(grid, newLoc)
					n++
					newLoc = loc.Add(diff.Multiply(n))
				}
				n = 1
				newLoc = locs[iix].Subtract(diff.Multiply(n))
				for !grid.IsOOB(newLoc) {
					total += SetAndAddUniqueAntinode(grid, newLoc)
					n++
					newLoc = locs[iix].Subtract(diff.Multiply(n))
				}
			}
		}
	}
	return total
}

func SetAndAddUniqueAntinode(g *Grid, loc XY) int {
	// This is just for robust counting as we go, avoiding duplicates.
	if existing, ok := g.GetSet(loc, Antinode); ok && existing != Antinode {
		return 1
	}
	return 0
}

func buildAntennaMap(grid *Grid) map[rune][]XY {
	am := make(map[rune][]XY)
	for y, line := range grid.grid {
		for x := range line {
			p := XY{X: x, Y: y}
			if r := grid.Get(p); r != Empty {
				if poses, ok := am[r]; ok {
					am[r] = append(poses, p)
				} else {
					am[r] = []XY{p}
				}
			}
		}
	}
	return am
}

func load(filePath string) ([][]rune, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	var grid [][]rune
	for scanner.Scan() {
		row := []rune(scanner.Text())
		grid = append(grid, row)
	}
	return grid, nil
}
