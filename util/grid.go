package util

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Grid struct {
	raw     [][]rune // this copy for resets is a little weird
	Grid    [][]rune
	scanPos XY
}

const (
	OOB   = '\x10'
	Empty = '.'
)

func NewGridFromFile(filePath string) (*Grid, error) {
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
	return NewGrid(grid), nil
}

func NewGrid(raw [][]rune) *Grid {
	return &Grid{
		raw:     raw, // for resets
		Grid:    newGridFromRaw(raw),
		scanPos: XY{-1, 0},
	}
}

func NewGridFromDim(x, y int) *Grid {
	var grid [][]rune
	for i := 0; i < y; i++ {
		row := make([]rune, x)
		for ix := 0; ix < x; ix++ {
			row[ix] = Empty
		}
		grid = append(grid, row)
	}
	return NewGrid(grid)
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
	return g.Grid[pos.Y][pos.X]
}

var (
	ErrCannotConvertToInt = errors.New("ErrCannotConvertToInt")
)

func (g *Grid) GetInt(pos XY) (int, error) {
	r := g.Get(pos)
	if r < 48 || r > 57 {
		return -1, ErrCannotConvertToInt
	}
	return int(r - '0'), nil
}

func (g *Grid) IsOOB(pos XY) bool {
	if pos.X >= len(g.Grid[0]) || pos.Y >= len(g.Grid) || pos.X < 0 || pos.Y < 0 {
		return true
	}
	return false
}

var (
	equalScanFunc    = func(a, b rune) bool { return a == b }
	notEqualScanFunc = func(a, b rune) bool { return a != b && a != OOB }
)

func (g *Grid) ScanOnce(r rune) (XY, bool) {
	return g.scanNext(r, XY{0, 0}, equalScanFunc)
}

// ScanOnceForNot finds the first instance of anything other than
func (g *Grid) ScanOnceForNot(r rune) (XY, bool) {
	return g.scanNext(r, XY{0, 0}, notEqualScanFunc)
}

func (g *Grid) Scan(r rune) bool {
	found, ok := g.scanNext(r, g.scanPos.Add(XY{1, 0}), equalScanFunc)
	if !ok {
		g.scanPos = XY{-1, 0} // reset
		return false
	}
	g.scanPos = found
	return true
}

func (g *Grid) ScanForNot(r rune) bool {
	found, ok := g.scanNext(r, g.scanPos.Add(XY{1, 0}), notEqualScanFunc)
	if !ok {
		g.scanPos = XY{-1, 0} // reset
		return false
	}
	g.scanPos = found
	return true
}

func (g *Grid) CurrentPosition() XY {
	return g.scanPos
}

func (g *Grid) scanNext(r rune, from XY, comparisonFunc func(a, b rune) bool) (XY, bool) {
	firstScan := true
	for y := from.Y; y < len(g.Grid); y++ {
		for x := 0; x < len(g.Grid[0]); x++ {
			if firstScan { // first scan pos
				x = from.X
				firstScan = false
			}
			p := XY{X: x, Y: y}
			if comparisonFunc(g.Get(p), r) {
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
	g.Grid[pos.Y][pos.X] = r
	return true
}

// GetSet returns the existing rune at `pos` if found, plus true if `r` is set successfully, false if not (OOB).
func (g *Grid) GetSet(pos XY, r rune) (rune, bool) {
	if g.IsOOB(pos) {
		return OOB, false // choosing a silent semantic for now
	}
	if existing := g.Get(pos); existing != Empty {
		g.Grid[pos.Y][pos.X] = r
		return existing, true
	}
	g.Grid[pos.Y][pos.X] = r
	return Empty, true
}

func (g *Grid) Print() {
	for y := range g.Grid {
		for x := range g.Grid[y] {
			fmt.Printf("%c", g.Grid[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *Grid) PrintString() string {
	var sb strings.Builder
	for y := range g.Grid {
		for x := range g.Grid[y] {
			sb.WriteRune(g.Grid[y][x])
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (g *Grid) Checksum() string {
	var sb strings.Builder
	for _, row := range g.Grid {
		for _, r := range row {
			sb.WriteRune(r)
		}
	}
	hash := sha256.Sum256([]byte(sb.String()))
	return hex.EncodeToString(hash[:])
}
func (g *Grid) Reset() {
	g.Grid = newGridFromRaw(g.raw)
}

type XY struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p XY) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("(%v,%v)", p.X, p.Y)), nil
}

func (p XY) Formatted() string {
	return fmt.Sprintf("(%d,%d)", p.X+1, p.Y+1)
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

func (p XY) NextSteps() []XY {
	return []XY{
		p.Add(Right), // right
		p.Add(Down),  // down
		p.Add(Left),  // left
		p.Add(Up),    // up
	}
}

func (p XY) NextDirections() []Direction {
	return []Direction{Right, Down, Left, Up}
}

func NextDirections() []Direction {
	return []Direction{Right, Down, Left, Up}
}

// TODO: This is imprecise semantically and maybe weird.
type Direction = XY

var (
	Right = XY{1, 0}
	Down  = XY{0, 1}
	Left  = XY{-1, 0}
	Up    = XY{0, -1}
)

func (d Direction) TurnClockwise() Direction {
	var nd Direction
	switch d {
	case Right:
		nd = Down
	case Down:
		nd = Left
	case Left:
		nd = Up
	case Up:
		nd = Right
	}
	return nd
}

func (d Direction) TurnCounterClockwise() Direction {
	var nd Direction
	switch d {
	case Right:
		nd = Up
	case Down:
		nd = Right
	case Left:
		nd = Down
	case Up:
		nd = Left
	}
	return nd
}

func (d Direction) Name() string {
	var n string
	switch d {
	case Right:
		n = "Right"
	case Down:
		n = "Down"
	case Left:
		n = "Left"
	case Up:
		n = "Up"
	}
	return n
}

func (d Direction) ToRune() rune {
	switch d {
	case Up:
		return '^'
	case Left:
		return '<'
	case Right:
		return '>'
	default: // case Down:
		return 'v'
	}
}
