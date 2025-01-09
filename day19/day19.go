package day19

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	InputTowelsPath  = "day19/input_towels.txt"
	InputDesignsPath = "day19/input_designs.txt"
)

func Runner() error {
	towels, err := loadTowels(InputTowelsPath)
	if err != nil {
		return err
	}
	designs, err := loadDesigns(InputDesignsPath)
	if err != nil {
		return err
	}
	fmt.Printf("\n:: DAY 19 ::\n")
	fmt.Printf("Part 1: %v\n", Part1(towels, designs))
	fmt.Printf("Part 2: %v\n", Part2(towels, designs))
	return nil
}

// I think this is a trie problem, but let's try it without first. Might also be dynamic.
func Part1(towels, designs []string) int {
	var total int
	for _, design := range designs {
		if IsDesignPossible(towels, design, nil) {
			total++
		}
	}
	return total
}

func IsDesignPossible(towels []string, design string, acc []string) bool {
	var possible bool
	for _, towel := range towels {
		if len(design) == 0 {
			return true
		}
		if strings.HasPrefix(design, towel) {
			possible = possible || IsDesignPossible(towels, strings.TrimPrefix(design, towel), append(acc, towel))
		}
	}
	return possible
}

func Part2(towels, designs []string) int {
	var total int
	for _, design := range designs {
		total += NumPossibleDesigns(towels, design, make(map[string]int))
	}
	return total
}

func NumPossibleDesigns(towels []string, design string, seen map[string]int) int {
	if c, ok := seen[design]; ok {
		return c
	}
	var total int
	for _, towel := range towels {
		if len(design) == 0 {
			return 1
		}
		if strings.HasPrefix(design, towel) {
			total += NumPossibleDesigns(towels, strings.TrimPrefix(design, towel), seen)
		}
	}
	seen[design] = total
	return total
}

func loadTowels(inputPath string) ([]string, error) {
	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewScanner(f)
	buf.Scan()
	return strings.Split(buf.Text(), ", "), nil
}

func loadDesigns(inputPath string) ([]string, error) {
	f, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewScanner(f)
	var designs []string
	for buf.Scan() {
		designs = append(designs, buf.Text())
	}
	return designs, nil
}
