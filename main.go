package main

import (
	"github.com/jmontroy90/aoc-2024/day1"
	"github.com/jmontroy90/aoc-2024/day2"
)

func main() {
	if err := day1.Runner(); err != nil {
		panic(err)
	}
	if err := day2.Runner(); err != nil {
		panic(err)
	}
}
