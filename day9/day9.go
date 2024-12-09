package day9

import (
	"bufio"
	"fmt"
	"os"
)

const (
	InputPath = "day9/input.txt"
)

func Runner() error {
	rs, err := load(InputPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 9 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(rs))
	fmt.Printf("Part 2: %v\n", Part2(rs))
	return nil
}

func Part1(rs []rune) int {
	expanded := ExpandDiskMap(rs)
	c := Compact(expanded)
	return GenerateChecksum(c)
}

func Part2(rs []rune) int {
	expanded := ExpandDiskMap(rs)
	c := CompactBlocks(expanded)
	return GenerateChecksum(c)
}

// ExpandDiskMap converts alternating digits to ID numbers and free spaces.
// We need to handle ID numbers as numbers, not runes.
func ExpandDiskMap(rs []rune) []int {
	var expanded []int
	for ix, count := range rs {
		n := -1
		if ix%2 == 0 { // block
			n = ix / 2
		}
		var adds []int
		for iix := 0; iix < int(count-'0'); iix++ {
			adds = append(adds, n)
		}
		expanded = append(expanded, adds...)
	}
	return expanded
}

func Compact(rs []int) []int {
	var lIx int
	rIx := len(rs) - 1
	for {
		lr := rs[lIx]
		for lr != -1 && lIx < rIx {
			lIx++
			lr = rs[lIx]
		}
		rr := rs[rIx]
		for rr == -1 && rIx > lIx {
			rIx--
			rr = rs[rIx]
		}
		if lIx >= rIx {
			break // post-increment / decrement
		}
		rs[lIx], rs[rIx] = rs[rIx], rs[lIx]

	}
	return rs
}

func CompactBlocks(rs []int) []int {
	var blockSize int
	var lIx, lScanIx int
	rIx, rScanIx := len(rs)-1, len(rs)-1
	for {
		rIx = rScanIx
		if rIx <= 0 {
			// we've already scanned here, we're done
			break
		}
		if rs[rIx] == -1 { // not at a block
			rIx--
			rScanIx--
			continue
		}
		for rScanIx >= 0 && rs[rScanIx] == rs[rIx] {
			rScanIx--
		}
		blockSize = rIx - rScanIx
		// scan for space per block
		lIx = 0
		lScanIx = 0
		for {
			lIx = lScanIx       // move up left scanner index
			if lIx >= rScanIx { // if we start scanning our original block itself...
				break // ...the block must stay where it is
			}
			if rs[lIx] != -1 {
				lIx++
				lScanIx++
				continue
			}
			for rs[lScanIx] == -1 {
				lScanIx++
			}
			if lScanIx-lIx < blockSize {
				continue // does not fit in this spot
			}
			// block fits, put it in
			copy(rs[lIx:lIx+blockSize], rs[rScanIx+1:rIx+1]) // off by one yay
			copy(rs[rScanIx+1:rIx+1], FillSlice(-1, blockSize))
			break
		}
	}
	return rs
}

func FillSlice(n, c int) []int {
	out := make([]int, c)
	for i := 0; i < c; i++ {
		out[i] = n
	}
	return out
}

func GenerateChecksum(is []int) int {
	var checksum int
	for ix, i := range is {
		if i == -1 {
			continue
		}
		checksum += ix * i
	}
	return checksum
}

func load(filePath string) ([]rune, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	return []rune(scanner.Text()), nil
}
