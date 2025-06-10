package day21

import (
	"fmt"
	"github.com/jmontroy90/aoc-2024/util"
	"slices"
	"strconv"
	"strings"
)

const (
	InputPath = "day21/input.txt"
)

func Runner() error {
	_, err := util.NewGridFromFile(InputPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 21 ::\n")
	fmt.Printf("Part 1: %v\n", Part1([]string{"029A", "980A", "179A", "456A", "379A"}))
	return nil
}

func Part1(codes []string) int {
	var total int
	for _, code := range codes {
		ss := FindShortestSeq(code)
		codeNum, _ := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(code, "0"), "A"))
		fmt.Printf("shortest sequence: %d\ncode number: %d\n", ss, codeNum)
		total += ss * codeNum
	}
	return total
}

func FindShortestSeq(code string) int {
	inputs := DetermineInputs(code)
	adj := BuildAdjacency(BuildKeypadGrid())
	fullPaths := GenCombo(inputs, nil)
	var newFullPaths [][]string
	for _, fp := range fullPaths {
		newFullPaths = append(newFullPaths, GenCombo(GenPath(fp, adj), nil)...)
	}
	var finalPaths [][]string
	for _, fp := range newFullPaths {
		finalPaths = append(finalPaths, GenCombo(GenPath(fp, adj), nil)...)
	}

	//for _, fp := range finalPaths {
	//	if len(strings.Join(fp, ""))
	//}
	//sequenceLenCount := make(map[int]int)
	//stepCount := make(map[int]int)
	shortestSeq := int(1e8) // BIG NUMBAH
	for _, fp := range finalPaths {
		pathLen := len(strings.Join(fp, ""))
		if pathLen < shortestSeq {
			shortestSeq = pathLen
		}
		if pathLen == shortestSeq {
			fmt.Println(strings.Join(fp, ""))
		}
		//sequenceLenCount[pathLen]++
		//stepCount[len(fp)]++
	}
	return shortestSeq
}

func printPath(path []string) {
	//for _, p := range path {
	var sb strings.Builder
	for _, ps := range path {
		sb.WriteString(ps)
	}
	fmt.Println(sb.String())
	sb.Reset()
}

func GenPath(path []string, adj map[rune][]Position) [][]string {
	curr := 'A'
	var out [][]string
	for _, pathStep := range path {
		for _, end := range []rune(pathStep) {
			pathOpts := DirectionsForKeypad(curr, end, adj)
			out = append(out, pathOpts)
			curr = end
		}
	}
	return out
}

func asSortedStrings(paths [][]string) []string {
	var allPathsStr []string
	for _, ap := range paths {
		var path string
		for _, p := range ap {
			path += p
		}
		allPathsStr = append(allPathsStr, path)
	}
	slices.Sort(allPathsStr)
	return allPathsStr
}

func GenCombo(paths [][]string, curr []string) [][]string {
	if len(paths) == 0 {
		return [][]string{curr}
	}
	var allPaths [][]string
	for _, opt := range paths[0] {
		next := make([]string, len(curr)+1)
		copy(next, curr)
		next[len(next)-1] = opt
		newPaths := GenCombo(paths[1:], next)
		allPaths = append(allPaths, newPaths...)
	}
	return allPaths
}

func DetermineInputs(sequence string) [][]string {
	start := 'A'
	var allDirs [][]string
	for _, r := range []rune(sequence) {
		end := r
		dirs := DirectionsForDigit(start, end)
		var dirStrings []string
		for _, dir := range dirs {
			var dirString string
			for _, d := range dir {
				dirString += string(d.ToRune())
			}
			dirString += "A"
			dirStrings = append(dirStrings, dirString)
		}
		allDirs = append(allDirs, dirStrings)
		start = r
	}
	return allDirs
}

func BuildKeypadGrid() *util.Grid {
	grid := util.NewGridFromDim(3, 2)
	grid.Set(util.XY{1, 0}, '^')
	grid.Set(util.XY{2, 0}, 'A')
	grid.Set(util.XY{0, 1}, '<')
	grid.Set(util.XY{1, 1}, 'v')
	grid.Set(util.XY{2, 1}, '>')
	grid.Set(util.XY{0, 0}, '#')
	return grid
}

type Position struct {
	Key       rune
	Direction util.Direction
}

// BFS to build adjacency
func BuildAdjacency(grid *util.Grid) map[rune][]Position {
	adjacency := make(map[rune][]Position)
	start := util.XY{2, 0}
	processed := make(map[util.XY]bool)
	queue := []util.XY{start}
	for len(queue) != 0 {
		pos := queue[0]
		queue = queue[1:]
		if processed[pos] {
			continue
		}
		for _, dir := range util.NextDirections() {
			next := pos.Add(dir)
			if grid.Get(next) != util.OOB && grid.Get(next) != '#' {
				adjacency[grid.Get(pos)] = append(adjacency[grid.Get(pos)], Position{grid.Get(next), dir})
				queue = append(queue, next)
			}
		}
		processed[pos] = true
	}
	return adjacency
}

// DFS to explore all paths between start and end (could probably BFS too).
func DirectionsForKeypad(start, end rune, adjacency map[rune][]Position) []string {
	var paths []string
	seen := make(map[rune]bool)
	var dfs func(node Position, path []rune)
	dfs = func(node Position, path []rune) {
		path = append(path, node.Direction.ToRune()) // This creates a new slice for this frame; backtracking for free!
		seen[node.Key] = true
		if node.Key == end {
			pathCopy := make([]rune, len(path))
			copy(pathCopy, path)
			paths = append(paths, string(pathCopy[1:])+"A") // the first direction added is bogus, and we need to press 'A'
		}
		for _, neighbor := range adjacency[node.Key] {
			if !seen[neighbor.Key] {
				dfs(neighbor, path)
			}
		}
		seen[node.Key] = false // We're done here, other paths can look at this.
	}
	dfs(Position{start, util.XY{}}, nil)
	return filterShortestPaths(paths)
}

func filterShortestPaths(paths []string) []string {
	minLen := 100000000 // big enough lol
	for _, path := range paths {
		if len(path) < minLen {
			minLen = len(path)
		}
	}
	var shortestPaths []string
	for _, path := range paths {
		if len(path) == minLen {
			shortestPaths = append(shortestPaths, path)
		}
	}
	return shortestPaths
}

// We kinda cheat and consider the whole thing to be a 4x3 grid of numbers.
// The movements don't care about the actual value, just their "distance" away from each other.
// So '0' -> 2, 'A' -> 3, '1' -> 4
// This is a "numerical" solution - building an adjacency list is probably much much cleaner,
// because we'd have no row / col magic + permutations would just pop out.
func DirectionsForDigit(start rune, end rune) [][]util.Direction {
	var out []util.Direction
	var si int
	switch start {
	case '0':
		si = 1
	case 'A':
		si = 2
	default:
		si = int(start-'0') + 2
	}
	var ei int
	switch end {
	case '0':
		ei = 1
	case 'A':
		ei = 2
	default:
		ei = int(end-'0') + 2
	}
	sPos := util.XY{si % 3, si / 3}
	ePos := util.XY{ei % 3, ei / 3}
	threes := abs(sPos.Y - ePos.Y)
	ones := abs(sPos.X - ePos.X)
	for i := 0; i < threes; i++ {
		if sPos.Y > ePos.Y {
			out = append(out, util.Down)
		} else {
			out = append(out, util.Up)
		}
	}
	for i := 0; i < ones; i++ {
		if sPos.X > ePos.X {
			out = append(out, util.Left)
		} else {
			out = append(out, util.Right)
		}
	}
	// get permutations
	var ixes []int
	for i := 0; i < len(out); i++ {
		ixes = append(ixes, i)
	}
	ixPerms := util.Permutations(ixes)
	if (ones > 0 && threes == 0) || (threes > 0 && ones == 0) {
		ixPerms = ixPerms[:1] // handles the "just go up / right / etc." case
	}
	var perms [][]util.XY
	permDupe := make(map[string]bool)
	for _, ixPerm := range ixPerms {
		var perm []util.XY
		var permString string
		for _, ix := range ixPerm {
			permString += string(out[ix].ToRune())
			perm = append(perm, out[ix])
		}
		if !permDupe[permString] {
			perms = append(perms, perm)
			permDupe[permString] = true
		}
	}
	return perms
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
