package day15

import (
	"bufio"
	"fmt"
	"github.com/jmontroy90/aoc-2024/util"
	"os"
	"slices"
)

const (
	InputGridPath  = "day15/input_grid.txt"
	InputMovesPath = "day15/input_moves.txt"
)

func Runner() error {
	grid, err := util.NewGridFromFile(InputGridPath)
	if err != nil {
		return err
	}
	moves, err := loadMoves(InputMovesPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 15 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(grid, moves))
	grid.Reset()
	_, sum := Part2(grid, moves)
	fmt.Printf("Part 2: %v\n", sum)
	return nil
}

func Part1(grid *util.Grid, moves []util.XY) int {
	pos, _ := grid.ScanOnce('@')
	for _, m := range moves {
		//grid.Print()
		//fmt.Printf("move: %v\n", string(xyToDir(m)))
		pos = Step(grid, pos, m, '@')
	}
	var sum int
	for grid.Scan('O') {
		sum += 100*grid.CurrentPosition().Y + grid.CurrentPosition().X
	}
	return sum
}

func Part2(grid *util.Grid, moves []util.XY) (*util.Grid, int) {
	wide := WidenGrid(grid)
	pos, _ := wide.ScanOnce('@')
	//wide.Print()
	for _, m := range moves {
		//fmt.Printf("move: %v\n", string(xyToDir(m)))
		pos = StepBox(wide, pos, m, '@')
		//wide.Print()
	}
	var sum int
	for wide.Scan('[') {
		sum += 100*wide.CurrentPosition().Y + wide.CurrentPosition().X
	}
	return wide, sum
}

func WidenGrid(orig *util.Grid) *util.Grid {
	wide := util.NewGridFromDim(len(orig.Grid[0])*2, len(orig.Grid))
	for rIx := range orig.Grid {
		for cIx := range orig.Grid[0] {
			switch c := orig.Grid[rIx][cIx]; c {
			case '#', util.Empty:
				wide.Grid[rIx][cIx*2] = c
				wide.Grid[rIx][cIx*2+1] = c
			case 'O':
				wide.Grid[rIx][cIx*2] = '['
				wide.Grid[rIx][cIx*2+1] = ']'
			case '@':
				wide.Grid[rIx][cIx*2] = '@'
				wide.Grid[rIx][cIx*2+1] = util.Empty
			}
		}
	}
	return wide
}

func Step(grid *util.Grid, pos util.XY, move util.XY, char rune) util.XY {
	nextPos := pos.Add(move)
	switch nextChar := grid.Get(nextPos); nextChar {
	case util.Empty: // trivial case
		grid.Set(nextPos, char)
		grid.Set(pos, util.Empty)
		return nextPos
	case 'O':
		return Push(grid, pos, move)
	default: // case '#', e.g. blocked
		return pos
	}
}

func Push(grid *util.Grid, pos, move util.XY) util.XY {
	orig := pos
	var pushed []util.XY
	var canPush bool
	for {
		next := pos.Add(move)
		c := grid.Get(next)
		if c == 'O' {
			pushed = append(pushed, next)
			pos = next
			continue
		} else if c == '#' {
			break
		} else if c == util.Empty {
			canPush = true
			break
		}
	}
	if !canPush {
		return orig
	}
	for _, p := range pushed {
		grid.Set(p.Add(move), grid.Get(p))
	}
	grid.Set(orig, util.Empty)
	grid.Set(orig.Add(move), '@')
	return orig.Add(move)
}

func StepBox(grid *util.Grid, pos util.XY, move util.XY, char rune) util.XY {
	nextPos := pos.Add(move)
	switch nextChar := grid.Get(nextPos); nextChar {
	case util.Empty: // trivial case
		grid.Set(nextPos, char)
		grid.Set(pos, util.Empty)
		return nextPos
	case ']', '[':
		return PushBox(grid, pos, move)
	default: // case '#', e.g. blocked
		return pos
	}
}

func PushBox(grid *util.Grid, orig, move util.XY) util.XY {
	var finalPos util.XY
	switch d := xyToDir(move); d {
	case '^', 'v': // check for multi-box expanding pushes
		finalPos = pushBoxUD(grid, orig, move)
	case '<', '>': // check for single-box pushes
		finalPos = pushBoxLR(grid, orig, move)
	}
	return finalPos
}

// This is basically the same logic as Part 1, but with a two-pos box. With some types, they could be the same code.
func pushBoxLR(grid *util.Grid, orig, move util.XY) util.XY {
	currPos := orig
	var pushes []util.XY
	var canPush bool
	for {
		p1 := currPos.Add(move)
		p2 := p1.Add(move)
		n1 := grid.Get(p1)
		n2 := grid.Get(p2)
		if !isBox(n1, n2) {
			continue
		}
		p3 := p2.Add(move)                     // place to push onto
		if n3 := grid.Get(p3); isBoxSide(n3) { // trivial
			pushes = append(pushes, p1, p2)
			currPos = p2
			continue
		} else if n3 == util.Empty {
			canPush = true
			pushes = append(pushes, p1, p2)
			break
		} else { // if n3 == '#'
			break
		}
	}
	if !canPush {
		return orig
	}
	slices.Reverse(pushes)
	for _, p := range pushes {
		grid.Set(p.Add(move), grid.Get(p))
	}
	grid.Set(orig, util.Empty)
	grid.Set(orig.Add(move), '@')
	return orig.Add(move)
}

// So much logic that could be consolidated and cleaned up. Probably there's data structures here that like immediately solve the problems with deduping and ordering I had. Well, it works.
func pushBoxUD(grid *util.Grid, orig, move util.XY) util.XY {
	var pushes []util.XY
	pushers := map[util.XY]struct{}{orig: {}}
	seen := make(map[util.XY]struct{})
	var noPush bool
PUSHLOOP:
	for {
		newPushers := make(map[util.XY]struct{})
		for pusher := range pushers {
			next := pusher.Add(move)
			switch c := grid.Get(next); c {
			case '[':
				if _, ok := seen[next]; !ok {
					pushes = append(pushes, next)
					newPushers[next] = struct{}{}
					seen[next] = struct{}{}
				}
				r := next.Add(util.XY{1, 0})
				if _, ok := seen[r]; !ok {
					pushes = append(pushes, r)
					newPushers[r] = struct{}{}
					seen[r] = struct{}{}
				}
			case ']':
				if _, ok := seen[next]; !ok {
					pushes = append(pushes, next)
					newPushers[next] = struct{}{}
					seen[next] = struct{}{}
				}
				l := next.Add(util.XY{-1, 0})
				if _, ok := seen[l]; !ok {
					pushes = append(pushes, l)
					newPushers[l] = struct{}{}
					seen[l] = struct{}{}
				}
			case util.Empty:
				continue
			case '#':
				noPush = true
				break PUSHLOOP
			}
		}
		if len(newPushers) == 0 {
			break // do it!!!
		}
		pushers = newPushers
	}
	if noPush {
		return orig
	}
	slices.Reverse(pushes)
	for _, p := range pushes {
		grid.Set(p.Add(move), grid.Get(p))
		grid.Set(p, util.Empty)
	}
	grid.Set(orig, util.Empty)
	grid.Set(orig.Add(move), '@')
	return orig.Add(move)
}

// This only works if all boxes are well-formed, which they will be since we control that.
func isBox(n1, n2 rune) bool {
	return isBoxSide(n1) && isBoxSide(n2)
}

func isBoxSide(c rune) bool {
	return c == '[' || c == ']'
}

func loadMoves(fileName string) ([]util.XY, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewScanner(f)
	var moves []util.XY
	for buf.Scan() {
		moves = append(moves, dirsToXYs(buf.Text())...)
	}
	return moves, nil
}

func dirsToXYs(dirs string) []util.XY {
	moves := make([]util.XY, len(dirs))
	for ix, dir := range []rune(dirs) {
		moves[ix] = dirToXY(dir)
	}
	return moves
}

func dirToXY(dir rune) util.XY {
	switch dir {
	case '^':
		return util.XY{0, -1}
	case '<':
		return util.XY{-1, 0}
	case '>':
		return util.XY{1, 0}
	default: // case 'v':
		return util.XY{0, 1}
	}
}

func xyToDir(xy util.XY) rune {
	switch xy {
	case util.XY{0, -1}:
		return '^'
	case util.XY{-1, 0}:
		return '<'
	case util.XY{1, 0}:
		return '>'
	default: // case 'v':
		return 'v'
	}
}
