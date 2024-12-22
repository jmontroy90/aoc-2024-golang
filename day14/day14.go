package day14

import (
	"bufio"
	"fmt"
	"github.com/jmontroy90/aoc-2024/util"
	"os"
	"regexp"
	"strconv"
)

const (
	InputPath = "day14/input.txt"
)

func Runner() error {
	robots, err := loadRobots(InputPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 14 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(robots, GridX, GridY))
	fmt.Printf("Part 2: %v\n", Part2(robots, GridX, GridY))
	return nil
}

func Part1(robots []Robot, gridX, gridY int) int {
	finalPosLookup := make(map[int]int)
	for _, r := range robots {
		finalPos := r.Step(100, gridX, gridY)
		finalPosLookup[determineQuadrant(finalPos, gridX, gridY)]++
	}
	total := 1
	for quadrant, score := range finalPosLookup {
		if quadrant == -1 {
			continue
		}
		total *= score
	}
	return total
}

// Whooo boy this is ugly. Basically we just look for an iteration that has a lot of robots in a single row.
// The Christmas tree pops out from there.
func Part2(robots []Robot, gridX, gridY int) int {
	robotsCopy := make([]Robot, len(robots))
	copy(robotsCopy, robots)
	grid := util.NewGridFromDim(gridX, gridY)
	var maxInRow, stepsWithMax int
	// 10000 is arbitrarily chosen based on the magnitude of other answers I've seen; basically, I cheated.
	for steps := 1; steps < 10000; steps++ {
		grid.Clear()
		if steps%1000 == 0 {
			fmt.Printf("steps = %v\n", steps)
		}
		for _, robot := range robotsCopy {
			pos := robot.Step(steps, gridX, gridY)
			curr := grid.Get(pos)
			if curr == '.' {
				if ok := grid.Set(pos, '1'); !ok {
					panic("something happened while creating grid!")
				}
			} else {
				iCurr, _ := strconv.Atoi(string(curr))
				n := rune(iCurr+1) + '0'
				grid.Set(pos, n)
			}
		}
		for rowIx, row := range grid.Grid {
			for ix, r := range row {
				if r != util.Empty {
					var thisRow int
					ix2 := ix + 1
					if ix2 >= gridX {
						continue
					}
					r2 := row[ix2]
					for r2 == r {
						thisRow++
						ix2++
						if ix2 >= gridX {
							break
						}
						r2 = row[ix2]
					}
					if thisRow > maxInRow {
						fmt.Printf("new high of %v, steps == %v and row = %v\n", thisRow, steps, rowIx)
						maxInRow = thisRow
						stepsWithMax = steps
					}
				}
			}
		}
	}
	return stepsWithMax
}

// Key:
//
//	0 == top left
//	1 == bottom right
//	2 == top right
//	3 == bottom left
func determineQuadrant(pos util.XY, gridX, gridY int) int {
	switch {
	case pos.X < (gridX)/2 && pos.Y < (gridY)/2:
		return 0
	case pos.X > (gridX)/2 && pos.Y > (gridY)/2:
		return 1
	case pos.X > (gridX)/2 && pos.Y < (gridY)/2:
		return 2
	case pos.X < (gridX)/2 && pos.Y > (gridY)/2:
		return 3
	default:
		return -1
	}
}

const (
	GridX = 101
	GridY = 103
)

type Robot struct {
	CurrPos  util.XY
	Velocity util.XY // Once again, overloaded type.
}

// Awkward signature with grid stuff overloaded in there.
func (r *Robot) Step(n int, gridX, gridY int) util.XY {
	if n == 0 {
		return r.CurrPos
	}
	r.CurrPos = r.CurrPos.Add(r.Velocity)
	r.wrapCurrPos(gridX, gridY)
	return r.Step(n-1, gridX, gridY)
}

func (r *Robot) wrapCurrPos(xGrid, yGrid int) {
	r.CurrPos.X = mod(r.CurrPos.X, xGrid)
	r.CurrPos.Y = mod(r.CurrPos.Y, yGrid)
}

func mod(a, b int) int {
	return (a%b + b) % b
}

var (
	rRobot = regexp.MustCompile("p=(?P<posX>\\d+),(?P<posY>\\d+) v=(?P<vX>-?\\d+),(?P<vY>-?\\d+)")
)

func loadRobots(fileName string) ([]Robot, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewScanner(f)
	var robots []Robot
	for buf.Scan() {
		text := buf.Text()
		ms := rRobot.FindStringSubmatch(text)
		posX, _ := strconv.Atoi(ms[1])
		posY, _ := strconv.Atoi(ms[2])
		vX, _ := strconv.Atoi(ms[3])
		vY, _ := strconv.Atoi(ms[4])
		robots = append(robots, Robot{CurrPos: util.XY{X: posX, Y: posY}, Velocity: util.XY{X: vX, Y: vY}})
	}
	return robots, nil
}
