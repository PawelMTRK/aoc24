package main

import (
	"aoc24/days"
	"fmt"
)

func RunDay(n int, dayfunc func()) {
	fmt.Println("-*", n)
	dayfunc()
}

func main() {
	fmt.Println("-_* AOC 2024 *_-")
	// RunDay(1, days.Day1)
	// RunDay(2, days.Day2)
	RunDay(3, days.Day3)
}
