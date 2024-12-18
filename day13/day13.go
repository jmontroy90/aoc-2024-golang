package day13

import (
	"bufio"
	"fmt"
	"github.com/jmontroy90/aoc-2024/util"
	"math"
	"os"
	"regexp"
	"strconv"
)

const (
	InputPath = "day13/input.txt"
)

func Runner() error {
	machines, err := loadMachines(InputPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 13 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(machines))
	fmt.Printf("Part 2: %v\n", Part2(machines))
	return nil
}

func Part1(machines []Machine) int {
	var total int
	for _, m := range machines {
		total += m.FindCheapestWithMath()
	}
	return total
}

func Part2(machines []Machine) int {
	var total int
	for _, m := range machines {
		m.PrizeAt.X += ErrorCorrection
		m.PrizeAt.Y += ErrorCorrection
		total += m.FindCheapestWithLinearAlgebra()
	}
	return total
}

type Machine struct {
	ButtonA util.XY // This is kinda a semantic overload on this type, but that's ok.
	ButtonB util.XY
	PrizeAt util.XY
}

const (
	ACost           int = 3
	BCost           int = 1
	ErrorCorrection     = 10000000000000
)

// Very dumb brute force, but because it's multiplication I think that's ok.
// UPDATE: not okay when we have so many more branches in part 2!
func (m Machine) FindCheapestButtonPresses() int {
	cheapest := 100*ACost + 100*BCost // The most it could be.
	var foundCheapest bool
	for i1 := 0; i1 <= 100; i1++ {
		for i2 := 0; i2 <= 100; i2++ {
			if cost := m.FindCostFor(i1, i2); cost > 0 && cost < cheapest {
				cheapest = cost
				foundCheapest = true
			}
		}
	}
	if !foundCheapest {
		return 0
	}
	return cheapest
}

func (m Machine) FindCostFor(nPressesA, nPressesB int) int {
	xA := nPressesA * m.ButtonA.X
	yA := nPressesA * m.ButtonA.Y
	xB := nPressesB * m.ButtonB.X
	yB := nPressesB * m.ButtonB.Y
	if xA+xB != m.PrizeAt.X || yA+yB != m.PrizeAt.Y {
		return -1
	}
	return ACost*nPressesA + BCost*nPressesB
}

func (m Machine) FindCheapestWithMath() int {
	cheapest := ErrorCorrection*ACost + ErrorCorrection*BCost // Kinda just a big number.
	var foundCheapest bool
	// Let's get rid of a loop.
	for nPressesA := 0; nPressesA <= 100; nPressesA++ {
		if nPressesA*ACost > cheapest {
			break // We know the rest of the candidates will only be more expensive.
		}
		nPressesB := m.FindBPressesFor(nPressesA) // Our candidate.
		if nPressesB == -1 {
			continue // No possible candidate.
		}
		if nPressesB*m.ButtonB.Y+nPressesA*m.ButtonA.Y != m.PrizeAt.Y {
			continue // Candidate's y-coordinate doesn't work.
		}
		// Everything works - check if it's cheaper than a previous one.
		if cost := ACost*nPressesA + BCost*nPressesB; cost < cheapest && cost > 0 {
			cheapest = cost
			foundCheapest = true
		}
	}
	if !foundCheapest {
		return 0
	}
	return cheapest
}

func (m Machine) FindBPressesFor(nPressesA int) int {
	xA := nPressesA * m.ButtonA.X
	nPressesB := float32(m.PrizeAt.X-xA) / float32(m.ButtonB.X)
	if math.Mod(float64(nPressesB), 1) != 0 {
		return -1
	}
	return int(nPressesB)
}

// FindCheapestWithLinearAlgebra produces the correct minimal - and ONLY - solution to a system of two equations with two variables.
// I followed this Youtube video for this: https://www.youtube.com/watch?v=-5J-DAsWuJc
// I feel okay with the fact that I didn't think to use linear algebra to help me. Fun lesson.
func (m Machine) FindCheapestWithLinearAlgebra() int {
	nPressesA := float64(m.PrizeAt.X*m.ButtonB.Y-m.PrizeAt.Y*m.ButtonB.X) / float64(m.ButtonA.X*m.ButtonB.Y-m.ButtonA.Y*m.ButtonB.X)
	nPressesB := (float64(m.PrizeAt.X) - float64(m.ButtonA.X)*nPressesA) / float64(m.ButtonB.X)
	if math.Mod(nPressesA, 1.0) != 0 || math.Mod(nPressesB, 1.0) != 0 {
		return 0
	}
	return ACost*int(nPressesA) + BCost*int(nPressesB)
}

var (
	rButtonLine = regexp.MustCompile("Button \\w: X\\+(?P<x>\\d+), Y\\+(?P<y>\\d+)")
	rPrizeLine  = regexp.MustCompile("Prize: X=(?P<x>\\d+), Y=(?P<y>\\d+)")
)

func loadMachines(fileName string) ([]Machine, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewScanner(f)
	var line int
	var machine Machine
	var machines []Machine
	for buf.Scan() {
		text := buf.Text()
		switch line {
		case 0:
			submatches := rButtonLine.FindStringSubmatch(text)
			x, _ := strconv.Atoi(submatches[1])
			y, _ := strconv.Atoi(submatches[2])
			machine.ButtonA = util.XY{x, y}
		case 1:
			submatches := rButtonLine.FindStringSubmatch(text)
			x, _ := strconv.Atoi(submatches[1])
			y, _ := strconv.Atoi(submatches[2])
			machine.ButtonB = util.XY{x, y}
		case 2:
			submatches := rPrizeLine.FindStringSubmatch(text)
			x, _ := strconv.Atoi(submatches[1])
			y, _ := strconv.Atoi(submatches[2])
			machine.PrizeAt = util.XY{x, y}
		case 3:
			machines = append(machines, machine)
		}
		line = (line + 1) % 4
	}
	return append(machines, machine), nil // last one needs to not be cut off
}
