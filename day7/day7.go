package day7

import (
	"bufio"
	"fmt"
	"github.com/jmontroy90/aoc-2024/util"
	"os"
	"strconv"
	"strings"
)

const (
	InputPath = "day7/input.txt"
)

func Runner() error {
	eqs, err := load(InputPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 7 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(eqs))
	fmt.Printf("Part 2: %v\n", Part2(eqs))
	return nil
}

type Op int

const (
	Add Op = iota
	Mult
	Concat
)

type Equation struct {
	Numbers   []int
	TestValue int
}

func Part1(eqs []Equation) int {
	var total int
	for _, eq := range eqs {
		if IsTrueEquation(eq, []Op{Add, Mult}) {
			total += eq.TestValue
		}
	}
	return total
}

func Part2(eqs []Equation) int {
	var total int
	for _, eq := range eqs {
		if IsTrueEquation(eq, []Op{Add, Mult, Concat}) {
			total += eq.TestValue
		}
	}
	return total
}

func IsTrueEquation(eq Equation, opTypes []Op) bool {
	// TODO: this is weird, I would do something else here, like use consistent types throughout.
	opInts := make([]int, len(opTypes))
	for ix := range opTypes {
		opInts[ix] = int(opTypes[ix])
	}
	gen := util.CartesianProduct(opInts, len(eq.Numbers))
	for {
		opsCartesian := gen()
		if len(opsCartesian) == 0 {
			break
		}
		var res int
		for ix, op := range opsCartesian {
			if res > eq.TestValue {
				break // short-circuit to prune some branches
			}
			if op == 0 {
				res += eq.Numbers[ix]
			} else if op == 1 {
				res *= eq.Numbers[ix]
			} else if op == 2 {
				res, _ = strconv.Atoi(fmt.Sprintf("%d%d", res, eq.Numbers[ix]))
			}
		}
		if res == eq.TestValue {
			return true
		}
	}
	return false
}

func load(filePath string) ([]Equation, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	var eqs []Equation
	for scanner.Scan() {
		res := strings.Split(scanner.Text(), ": ")
		var is []int
		for _, n := range strings.Split(res[1], " ") {
			i, _ := strconv.Atoi(n)
			is = append(is, i)
		}
		tv, _ := strconv.Atoi(res[0])
		eqs = append(eqs, Equation{Numbers: is, TestValue: tv})
	}
	return eqs, nil
}
