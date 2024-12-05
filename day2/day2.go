package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	InputPath = "day2/input.tsv"
)

func Runner() error {
	reports, err := loader(InputPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 2 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(reports))
	fmt.Printf("Part 2: %v\n", Part2(reports))
	return nil
}

func Part1(reports [][]int) int {
	var safe int
	for _, report := range reports {
		if ok, _ := IsReportSafe(report); ok {
			safe++
		}
	}
	return safe
}

func Part2(reports [][]int) int {
	var safe int
	for _, report := range reports {
		if IsDampenedReportSafe(report) {
			safe++
		}
	}
	return safe
}

// IsReportSafe returns whether the report is safer, and if not, the index where a problem was encountered.
func IsReportSafe(report []int) (bool, int) {
	last := report[0]
	lastDiff := 0
	for ix := 1; ix < len(report); ix++ {
		n := report[ix]
		diff := n - last
		if !isValidStep(diff, lastDiff) {
			return false, ix
		}
		lastDiff = n - last
		last = n
	}
	return true, -1
}

// IsDampenedReportSafe uses IsReportsafe and tries sliced permutations
// I was undercounting for a while before I got the full scope of the problem!
// lesson learned - think on all affected elements; I just focused on the obvious
// "remove one here right where the problem was encountered" case.
func IsDampenedReportSafe(report []int) bool {
	// at any given point, we have three elements in play:
	//	- report[ix], for n + diff
	//	- report[ix - 1], for last + diff
	//	- report[ix - 2], for lastDiff
	// so we need to try removing each of those elements to see if the new list is safe.
	// nix (e.g. "new index") goes up to the trouble point (not inclusive indexing), then starts after it
	// then decrements to try each prior one, up to three times total and not going out of bounds.
	if ok, ix := IsReportSafe(report); !ok {
		l := make([]int, len(report)-1)
		for nix := ix; nix >= 0 && nix-ix <= 2; nix-- {
			//l := append(report[:nix], report[nix+1:]...) // this is functionally WRONG
			//l := slices.Concat(report[:nix], report[nix+1:]) // this is unnecessary allocs
			copy(l, report[:nix])
			copy(l[nix:], report[nix+1:])
			if safe, _ := IsReportSafe(l); safe {
				return true
			}
		}
		return false
	}
	return true
}

func isValidStep(diff, lastDiff int) bool {
	if diff*lastDiff < 0 {
		return false // sign flipped
	}
	if abs(diff) > 3 || abs(diff) < 1 {
		return false
	}
	return true
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func loader(filePath string) ([][]int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	var lines [][]int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		iLine := make([]int, len(line))
		for ix, s := range line {
			iLine[ix], _ = strconv.Atoi(s)
		}
		lines = append(lines, iLine)
	}
	return lines, nil
}
