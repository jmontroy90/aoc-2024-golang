package main

import (
	"github.com/jmontroy90/aoc-2024/day1"
	"github.com/jmontroy90/aoc-2024/day2"
	"github.com/jmontroy90/aoc-2024/day3"
	"github.com/jmontroy90/aoc-2024/day4"
)

func main() {
	if err := day4.Runner(); err != nil {
		panic(err)
	}
	if err := day3.Runner(); err != nil {
		panic(err)
	}
	if err := day2.Runner(); err != nil {
		panic(err)
	}
	if err := day1.Runner(); err != nil {
		panic(err)
	}
}
