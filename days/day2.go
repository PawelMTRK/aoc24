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

func (r *Report) Without(n int) Report {
	levels := make([]int, 0)
	for i, v := range r.levels {
		if i == n {
			continue
		}
		levels = append(levels, v)
	}
	return Report{levels: levels}
}

func (r *Report) IsSafe() (bool, int) {
	last := r.levels[0]
	direction := 0
	for i, l := range r.levels[1:] {
		diff := l - last

		if direction == -1 && 0 < diff {
			return false, i
		} else if direction == 1 && diff < 0 {
			return false, i
		}

		// no int abs function workaround
		// and direction detection
		if diff < 0 {
			diff = -diff
			direction = -1
		} else if diff > 0 {
			direction = 1
		}

		if 3 < diff || diff < 1 {
			return false, i
		}

		last = l
	}
	return true, 0
}

func Day2() {
	res, _ := os.ReadFile("input2.txt")
	stringRes := string(res)
	lines := strings.Split(stringRes, "\n")
	safeN := 0
	for _, l := range lines {
		report := NewReport(l)
		isSafe, _ := report.IsSafe()
		if isSafe {
			safeN++
		} else {
			// TODO problem dampener answer is too low
			for j := range report.levels {
				correctedReport := report.Without(j)
				isSafeAgain, _ := correctedReport.IsSafe()
				if isSafeAgain {
					safeN++
					break
				}
			}
		}
	}
	fmt.Println("Safe reports:", safeN)
}
