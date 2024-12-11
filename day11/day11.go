package day11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	InputPath = "day11/input.txt"
)

func Runner() error {
	stones, err := load(InputPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 11 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(stones, 25))
	fmt.Printf("Part 2: %v\n", Part2(stones, 75))
	return nil
}

func Part1(stones []int, nBlinks int) int {
	var total int
	for _, stone := range stones {
		total += CountStonesAfterBlinks(stone, nBlinks)
	}
	return total
}

func Part2(stones []int, nBlinks int) int {
	var total int
	for _, stone := range stones {
		total += CountStonesAfterBlinksCached(stone, nBlinks) //, stone, nBlinks, 0)
	}
	return total
}

// CountStonesAfterBlinks doesn't care about the numbers for output. I bet Part 2 will make me care about the numbers.
func CountStonesAfterBlinks(stone int, nBlinks int) int {
	if nBlinks == 0 {
		return 1
	}
	if stone == 0 {
		return CountStonesAfterBlinks(1, nBlinks-1)
	} else if ss := strconv.Itoa(stone); len(ss)%2 == 0 { // even
		ns1, _ := strconv.Atoi(ss[:len(ss)/2])
		ns2, _ := strconv.Atoi(ss[len(ss)/2:])
		return CountStonesAfterBlinks(ns1, nBlinks-1) + CountStonesAfterBlinks(ns2, nBlinks-1)
	} else { // odd
		return CountStonesAfterBlinks(stone*2024, nBlinks-1)
	}
}

type StoneBlinkResult struct {
	Value     int
	NumBlinks int
}

var blinkCacher = make(map[StoneBlinkResult]int)

// CountStonesAfterBlinksCached does the algorithm without blasting off into space due to stacks.
// We do this simply via a cache. Memoization patterns here will be interesting to look at.
func CountStonesAfterBlinksCached(stone int, nBlinks int) int {
	if c, ok := blinkCacher[StoneBlinkResult{Value: stone, NumBlinks: nBlinks}]; ok {
		return c // early-exit for caching
	}
	if nBlinks == 0 {
		return 1
	}
	if stone == 0 {
		return CountStonesAfterBlinksCached(1, nBlinks-1)
	} else if ss := strconv.Itoa(stone); len(ss)%2 == 0 { // even
		ns1, _ := strconv.Atoi(ss[:len(ss)/2])
		ns2, _ := strconv.Atoi(ss[len(ss)/2:])
		created := CountStonesAfterBlinksCached(ns1, nBlinks-1) + CountStonesAfterBlinksCached(ns2, nBlinks-1)
		blinkCacher[StoneBlinkResult{Value: stone, NumBlinks: nBlinks}] = created
		return created
	} else { // odd
		created := CountStonesAfterBlinksCached(stone*2024, nBlinks-1)
		blinkCacher[StoneBlinkResult{Value: stone, NumBlinks: nBlinks}] = created
		return created
	}
}

func load(filePath string) ([]int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	ss := strings.Split(scanner.Text(), " ")
	stones := make([]int, len(ss))
	for ix := range ss {
		i, _ := strconv.Atoi(ss[ix])
		stones[ix] = i
	}
	return stones, nil
}
