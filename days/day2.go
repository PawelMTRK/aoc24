package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Report struct {
	levels []int
}

func NewReport(line string) Report {
	levels := make([]int, 0)
	for _, v := range strings.Split(line, " ") {
		n, _ := strconv.Atoi(v)
		levels = append(levels, n)
	}
	return Report{levels: levels}
}

func (r *Report) IsSafe() bool {
	last := r.levels[0]
	for i, l := range r.levels {
		// HACK skip comparing first element with itself
		if i == 0 {
			continue
		}
		diff := l - last
		// apparently there is no int abs function :c
		if diff < 0 {
			diff = -diff
		}
		if 3 < diff || diff < 1 {
			return false
		}
		last = l
	}
	return true
}

func Day2() {
	res, _ := os.ReadFile("input2.txt")
	stringRes := string(res)
	lines := strings.Split(stringRes, "\n")

	for _, l := range lines {
		report := NewReport(l)
		// fmt.Println(report)
		fmt.Println(report.IsSafe())
	}
}
