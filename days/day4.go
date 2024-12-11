package days

import (
	"aoc24/util"
	"fmt"
	"os"
	"strings"
	"sync"
)

type WordSearch struct {
	puzzle []string
	total  int
	m      sync.Mutex
	wg     sync.WaitGroup
}

func NewWordSearch(s string) *WordSearch {
	lines := strings.Split(s, "\n")
	ws := WordSearch{puzzle: lines, total: 0}
	return &ws
}

func (w *WordSearch) isXMASorSAMX(x, m, a, s byte) bool {
	// X - 88, 77 - M, 64 - A, 83 - S
	if x == 88 && m == 77 && a == 65 && s == 83 {
		return true
	} else if x == 83 && m == 65 && a == 77 && s == 88 {
		return true
	}
	return false
}

func (w *WordSearch) isExMAS(row1, row2, row3 string) bool {
	if row2[1] != 65 {
		return false
	}
	// 77 - M, 83 - S
	// find any of the 4 combinations
	// M M  S S  M S  S M
	//  A    A    A    A
	// S S  M M  M S  S M
	return util.Any(
		util.Compare3(row1[0], row1[2], 77) && util.Compare3(row3[0], row3[2], 83),
		util.Compare3(row1[0], row1[2], 83) && util.Compare3(row3[0], row3[2], 77),
		util.Compare3(row1[0], row3[0], 83) && util.Compare3(row1[2], row3[2], 77),
		util.Compare3(row1[0], row3[0], 77) && util.Compare3(row1[2], row3[2], 83),
	)
}

func (w *WordSearch) FindXMASByDirection(dir string) {
	switch dir {
	case "-":
		for r := range len(w.puzzle) {
			for c := range len(w.puzzle[r]) - 3 {
				if w.isXMASorSAMX(w.puzzle[r][c], w.puzzle[r][c+1], w.puzzle[r][c+2], w.puzzle[r][c+3]) {
					w.m.Lock()
					w.total++
					w.m.Unlock()
				}
			}
		}
	case "|":
		for r := range len(w.puzzle) - 3 {
			for c := range len(w.puzzle[r]) {
				if w.isXMASorSAMX(w.puzzle[r][c], w.puzzle[r+1][c], w.puzzle[r+2][c], w.puzzle[r+3][c]) {
					w.m.Lock()
					w.total++
					w.m.Unlock()
				}
			}
		}
	case "\\":
		for r := range len(w.puzzle) - 3 {
			for c := range len(w.puzzle[r]) - 3 {
				if w.isXMASorSAMX(w.puzzle[r][c], w.puzzle[r+1][c+1], w.puzzle[r+2][c+2], w.puzzle[r+3][c+3]) {
					w.m.Lock()
					w.total++
					w.m.Unlock()
				}
			}
		}
	case "/":
		for r := range len(w.puzzle) - 3 {
			for c := range len(w.puzzle[r]) - 3 {
				if w.isXMASorSAMX(w.puzzle[r][c+3], w.puzzle[r+1][c+2], w.puzzle[r+2][c+1], w.puzzle[r+3][c]) {
					w.m.Lock()
					w.total++
					w.m.Unlock()
				}
			}
		}
	}
	w.wg.Done()
}

func (w *WordSearch) FindExMAS() {
	for r := range len(w.puzzle) - 2 {
		for c := range len(w.puzzle[r]) - 2 {
			if w.isExMAS(w.puzzle[r][c:c+3], w.puzzle[r+1][c:c+3], w.puzzle[r+2][c:c+3]) {
				w.total++
			}
		}
	}
}

func Day4() {
	res, _ := os.ReadFile("input4.txt")
	stringRes := string(res)
	wordSearch := NewWordSearch(stringRes)
	wordSearch.wg.Add(4)
	go wordSearch.FindXMASByDirection("-")
	go wordSearch.FindXMASByDirection("|")
	go wordSearch.FindXMASByDirection("\\")
	go wordSearch.FindXMASByDirection("/")
	wordSearch.wg.Wait()
	fmt.Println("Total number of XMAS:", wordSearch.total)
	wordSearch.total = 0
	wordSearch.FindExMAS()
	fmt.Println("Total number of X-MAS:", wordSearch.total)
}
