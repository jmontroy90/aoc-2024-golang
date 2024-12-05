package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	InputPath = "day3/input.txt"
)

func Runner() error {
	lines, err := loader(InputPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 3 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(lines))
	fmt.Printf("Part 2: %v\n", Part2(lines))
	return nil
}

func Part1(lines []string) int {
	var total int
	for _, line := range lines {
		total += SumMultiplications(line)
	}
	return total
}

func SumMultiplications(line string) int {
	reg := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	var total int
	for _, match := range reg.FindAllStringSubmatch(line, -1) {
		i1, _ := strconv.Atoi(match[1])
		i2, _ := strconv.Atoi(match[2])
		total += i1 * i2
	}
	return total
}

func Part2(lines []string) int {
	var catLine strings.Builder
	for _, line := range lines {
		catLine.WriteString(line)
	}
	return SumWithStop(catLine.String())
}

type changePoint struct {
	active bool // true == "do", false == "don't"
	ix     int  // the index of the final position of the change point, after which we scan
}

type scanRange struct {
	start, end int
	isDone     bool
}

func SumWithStop(line string) int {
	rDo := regexp.MustCompile("do\\(\\)")
	rDont := regexp.MustCompile("don't\\(\\)")
	dos := rDo.FindAllStringIndex(line, -1)
	donts := rDont.FindAllStringIndex(line, -1)
	changePoints := make([]changePoint, 1, len(dos)+len(donts)+1)
	changePoints[0] = changePoint{active: true, ix: 0}
	for _, do := range dos {
		changePoints = append(changePoints, changePoint{
			active: true,
			ix:     do[1],
		})
	}
	for _, dont := range donts {
		changePoints = append(changePoints, changePoint{
			active: false,
			ix:     dont[0],
		})
	}
	sort.SliceStable(changePoints, func(i, j int) bool {
		return changePoints[i].ix < changePoints[j].ix
	})
	isActive := changePoints[0].active // we always start true
	var scanRanges []scanRange
	var currentRange scanRange
	for _, cp := range changePoints {
		switch {
		case !isActive && cp.active: // turn on
			isActive = true
			currentRange.start = cp.ix
		case isActive && !cp.active: // turn off
			isActive = false
			currentRange.end = cp.ix
			currentRange.isDone = true
			scanRanges = append(scanRanges, currentRange)
		}
	}
	var catLine strings.Builder
	for _, sr := range scanRanges {
		catLine.WriteString(line[sr.start:sr.end])
	}
	return SumMultiplications(catLine.String())
}

func loader(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
