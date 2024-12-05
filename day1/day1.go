package day1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	InputPath = "day1/input.tsv"
)

func Runner() error {
	l, r, err := Loader(InputPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 1 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(l, r))
	fmt.Printf("Part 2: %v\n", Part2(l, r))
	return nil
}

func Part1(l, r []int) int {
	slices.Sort(l)
	slices.Sort(r)
	var distance int
	for ix := range l {
		distance += int(math.Abs(float64(l[ix]) - float64(r[ix])))
	}
	return distance
}

func Part2(l, r []int) int {
	counts := make(map[int]int, len(r))
	for _, n := range r {
		counts[n] += 1
	}
	var distance int
	for _, n := range l {
		distance += n * counts[n]
	}
	return distance
}

func Loader(filePath string) ([]int, []int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	scanner := bufio.NewScanner(f)
	var left, right []int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		li, _ := strconv.Atoi(line[0])
		ri, _ := strconv.Atoi(line[1])
		left = append(left, li)
		right = append(right, ri)
	}
	return left, right, nil
}
