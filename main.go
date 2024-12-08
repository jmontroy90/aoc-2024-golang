package main

import (
	"github.com/jmontroy90/aoc-2024/day1"
	"github.com/jmontroy90/aoc-2024/day2"
	"github.com/jmontroy90/aoc-2024/day3"
	"github.com/jmontroy90/aoc-2024/day4"
	"github.com/jmontroy90/aoc-2024/day5"
	"github.com/jmontroy90/aoc-2024/day6"
	"github.com/jmontroy90/aoc-2024/day7"
)

func main() {
	if err := day7.Runner(); err != nil {
		panic(err)
	}
	if err := day6.Runner(); err != nil {
		panic(err)
	}
	if err := day5.Runner(); err != nil {
		panic(err)
	}
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
