package main

import (
	"days/days"
	"fmt"
)

func RunDay(n int, dayfunc func()) {
	fmt.Println("-*", n)
	dayfunc()
}

func main() {
	fmt.Println("-_* AOC 2024 *_-")
	RunDay(1, days.Day1)
	RunDay(2, days.Day2)
}
