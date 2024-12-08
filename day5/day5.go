package day5

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"strconv"
)

const (
	OrderingsPath = "day5/input_1.csv"
	UpdatesPath   = "day5/input_2.csv"
)

func Runner() error {
	orderings, err := loadOrderings(OrderingsPath)
	if err != nil {
		return err
	}
	updates, err := loadUpdates(UpdatesPath)
	if err != nil {
		return err
	}
	rlOrder := createOrderMap(orderings)
	fmt.Printf("\n:: DAY 5 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(updates, rlOrder))
	fmt.Printf("Part 2: %v\n", Part2(updates, rlOrder))
	return nil
}

func Part1(orderings [][]int, rlOrder map[int][]int) int {
	var sumOfMiddle int
	for _, ordering := range orderings {
		if IsUpdateValid(ordering, rlOrder) {
			sumOfMiddle += ordering[(len(ordering)-1)/2]
		}
	}
	return sumOfMiddle
}

func IsUpdateValid(update []int, rl map[int][]int) bool {
	for ix, o := range update {
		if v, ok := rl[o]; ok {
			for iix := ix; iix < len(update); iix++ {
				for _, mustBeBefore := range v {
					if update[iix] == mustBeBefore {
						return false
					}
				}
			}
		}
	}
	return true
}

func Part2(orderings [][]int, rlOrder map[int][]int) int {
	var sumOfMiddle int
	for _, ordering := range orderings {
		if IsUpdateValid(ordering, rlOrder) {
			continue
		}
		reordered := make([]int, len(ordering))
		copy(reordered, ordering)
		// each reordering is sorta greedy / local in scope, and so might disorder other things
		// but over time, they converge, yay
		// some stats on how many passes this takes might be interesting
		for !IsUpdateValid(reordered, rlOrder) {
			reordered = ReorderUpdate(reordered, rlOrder)
		}
		sumOfMiddle += reordered[(len(reordered)-1)/2]
	}
	return sumOfMiddle
}

func ReorderUpdate(update []int, rl map[int][]int) []int {
	for ix := range update {
		swapIx := ix
		if v, ok := rl[update[ix]]; ok {
			for iix := ix; iix < len(update); iix++ {
				for _, mustBeBefore := range v {
					if update[iix] == mustBeBefore {
						update = slices.Insert(update, swapIx, mustBeBefore)
						swapIx += 1
						update = slices.Delete(update, iix+1, iix+2)
					}
				}
			}
		}
	}
	return update
}

func createOrderMap(orderings [][2]int) map[int][]int {
	rl := make(map[int][]int, len(orderings))
	for _, ordering := range orderings {
		if v, ok := rl[ordering[1]]; ok {
			v = append(v, ordering[0])
			rl[ordering[1]] = v
		} else {
			rl[ordering[1]] = []int{ordering[0]}
		}
	}
	return rl
}

func loadOrderings(filePath string) ([][2]int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	creader := csv.NewReader(f)
	creader.Comma = '|'
	var orderings [][2]int
	rec, readErr := creader.Read()
	for readErr == nil {
		l, _ := strconv.Atoi(rec[0])
		r, _ := strconv.Atoi(rec[1])
		orderings = append(orderings, [2]int{l, r})
		rec, readErr = creader.Read()
	}
	return orderings, nil
}

func loadUpdates(filePath string) ([][]int, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	creader := csv.NewReader(f)
	creader.FieldsPerRecord = -1

	var updates [][]int
	rec, readErr := creader.Read()
	for readErr == nil {
		update := make([]int, len(rec))
		for ix := range rec {
			i, _ := strconv.Atoi(rec[ix])
			update[ix] = i
		}
		updates = append(updates, update)
		rec, readErr = creader.Read()
	}
	return updates, nil
}
